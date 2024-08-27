// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

use std::path::PathBuf;

use clap::Parser;
use git_cmd::Repo;
use protofetch::Protofetch;
use time::format_description;
use toml_edit::{value, DocumentMut};

const PROTOFETCH_META_KEYS: [&str; 3] = ["name", "description", "proto_out_dir"];

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long, default_value = "protofetch.toml")]
    protofetch_descriptor: String,

    #[arg(short = 'l', long, default_value = "protofetch.lock")]
    protofetch_lockfile: String,

    #[arg(short, long)]
    root: String,

    #[arg(short, long, default_value = ".")]
    git_dir: String,

    #[arg(short = 'd', long, default_value = "proto")]
    protos_dir: String,

    #[arg(short = 'o', long)]
    github_repo: String,

    #[arg(short, long, default_value = "true")]
    skip_hashes: bool,
}

#[tokio::main]
async fn main() -> eyre::Result<()> {
    // Setup
    let subscriber = tracing_subscriber::FmtSubscriber::new();
    tracing::subscriber::set_global_default(subscriber)?;

    let args = Args::parse();

    let (owner, repo_name) = args
        .github_repo
        .split_once('/')
        .expect("Invalid value for github_repo (use owner/repo format)");

    let root = PathBuf::from(&args.root);

    let mut protofetch_path = root.clone();
    protofetch_path.push(args.protofetch_descriptor);

    let mut proto_dir = root.clone();
    proto_dir.push(args.protos_dir);

    let mut protolock_path = root.clone();
    protolock_path.push(args.protofetch_lockfile);

    tracing::info!(?protofetch_path, ?protolock_path, ?proto_dir);
    let github_token = std::env::var("GITHUB_TOKEN").expect("A GITHUB_TOKEN must be set");
    let reader_token = std::env::var("GH_READER_TOKEN")
        .expect("A GH_READER_TOKEN must be set for reading the organization repositories");

    // Parse protofetch & fetch releases

    let toml = tokio::fs::read_to_string(&protofetch_path).await?;
    let mut protofetch_descriptor = toml
        .parse::<DocumentMut>()
        .expect("Protofetch descriptor not found");

    let authed_github = octocrab::OctocrabBuilder::default()
        .personal_token(reader_token)
        .build()?;
    let mut changes = false;
    let mut pr_message = String::new();

    for (k, val) in protofetch_descriptor.iter_mut() {
        let name = k.get();

        // skip other keys
        if PROTOFETCH_META_KEYS.contains(&name) {
            continue;
        }

        let url = val["url"].as_str().expect("URL must be a string");
        let (owner, repo) = owner_repo_from_url(url).expect("Invalid URL");
        let revision = val["revision"].as_str().expect("Revision must be a string");

        let release = authed_github
            .repos(owner, repo)
            .releases()
            .get_latest()
            .await;
        match release {
            Ok(release) => {
                // heuristics to skip hash updates since they might be newer than a release tag
                let skip_hash = revision.len() == 40 && args.skip_hashes;
                tracing::info!(%url, %revision, %release.tag_name, %skip_hash);

                if skip_hash {
                    continue;
                }

                if revision != release.tag_name {
                    pr_message
                        .push_str(&format!("{}: {} -> {}\n", name, revision, release.tag_name));
                    val["revision"] = value(release.tag_name);
                    changes = true;
                }
            }
            Err(e) => {
                // continue on error
                tracing::error!(%name, %e, %url);
            }
        }
    }
    // No changes
    if !changes {
        tracing::info!("No changes");
        return Ok(());
    }

    // Protofetch updating
    tokio::fs::write(&protofetch_path, protofetch_descriptor.to_string()).await?;
    tokio::fs::remove_dir_all(&proto_dir).await?;

    let protos = Protofetch::builder()
        .module_file_name(&protofetch_path)
        .output_directory_name(&proto_dir)
        .lock_file_name(&protolock_path)
        .try_build()
        .expect("Protofetch failed");

    protos
        .fetch(protofetch::LockMode::Recreate)
        .expect("Protofetch failed");

    // Git operations
    let repo = Repo::new(&args.git_dir)
        .unwrap_or_else(|_| panic!("Couldn't find git repository at {}", args.git_dir));
    let format = format_description::parse("[year]-[month]-[day]").expect("Failed to parse format");
    let branch_name = format!(
        "update-{}",
        time::OffsetDateTime::now_utc().format(&format)?
    );
    let changes = repo.changes(|_| true).expect("Failed to add changes");
    repo.checkout_new_branch(&branch_name)
        .expect("Failed to create branch");
    tracing::info!(%branch_name, ?changes);

    repo.add(&[
        protofetch_path.into_os_string().to_str().unwrap(),
        protolock_path.into_os_string().to_str().unwrap(),
        proto_dir.into_os_string().to_str().unwrap(),
    ])
    .expect("Failed to add files");
    repo.commit("chore: Update protofetch.toml")
        .expect("Failed to commit");
    repo.push(&branch_name).expect("Failed to push");

    // Create PR
    let authed_github = octocrab::OctocrabBuilder::default()
        .personal_token(github_token)
        .build()?;

    tracing::info!(%owner, %repo_name, ?branch_name, ?pr_message);
    authed_github
        .pulls(owner, repo_name)
        .create("Updating API descriptions", branch_name, "master")
        .body(pr_message)
        .send()
        .await?;

    Ok(())
}

#[tracing::instrument(level = "debug")]
fn owner_repo_from_url(url: &str) -> Option<(&str, &str)> {
    let url = url
        .trim_end_matches(".git")
        .trim_start_matches("github.com/");
    let mut parts = url.split('/').rev();
    let repo = parts.next()?;
    let owner = parts.next()?;
    Some((owner, repo))
}
