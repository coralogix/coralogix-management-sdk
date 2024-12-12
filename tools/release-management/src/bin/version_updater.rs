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
use time::format_description;
use toml_edit::DocumentMut;

const COMMIT_MESSAGE: &str = "chore: Update Rust & Go SDK versions";
const PR_TITLE: &str = "Update SDK versions";

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long, default_value = "../..")]
    root: String,

    #[arg(short, long, default_value = ".")]
    git_dir: String,

    #[arg(
        short = 'o',
        long,
        default_value = "coralogix/coralogix-management-sdk"
    )]
    github_repo: String,

    #[arg(short, long, short = 'v', default_value = "0.2.1")]
    sdk_version: String,

    #[arg(short = 'c', long, default_value = "rust")]
    rust_root: String,

    #[arg(short = 't', long, default_value = "go")]
    go_root: String,
}

#[tokio::main]
async fn main() -> eyre::Result<()> {
    let args = Args::parse();
    let cargo_toml_path = PathBuf::from(&args.root)
        .join(&args.rust_root)
        .join("Cargo.toml");
    let go_constants_path = PathBuf::from(&args.root)
        .join(&args.go_root)
        .join("constants.go");

    let cargo_toml_string = tokio::fs::read_to_string(&cargo_toml_path).await?;
    let go_constants_source_code_string = tokio::fs::read_to_string(&go_constants_path).await?;
    let updated_cargo_toml = set_rust_sdk_version(&args, &cargo_toml_string).await?;
    let updated_go_constants = set_go_sdk_version(&args, &go_constants_source_code_string).await?;
    tokio::fs::write(&cargo_toml_path, &updated_cargo_toml).await?;
    tokio::fs::write(&go_constants_path, updated_go_constants).await?;

    let (owner, repo_name) = args
        .github_repo
        .split_once('/')
        .expect("Invalid value for github_repo (use owner/repo format)");

    let github_token = std::env::var("GITHUB_TOKEN").expect("A GITHUB_TOKEN must be set");

    let repo = Repo::new(&args.git_dir)
        .unwrap_or_else(|_| panic!("Couldn't find git repository at {}", args.git_dir));
    let format = format_description::parse("[year]-[month]-[day]").expect("Failed to parse format");
    let branch_name = format!(
        "version-update-{}",
        time::OffsetDateTime::now_utc().format(&format)?
    );
    let changes = repo.changes(|_| true).expect("Failed to add changes");
    repo.checkout_new_branch(&branch_name)
        .expect("Failed to create branch");
    tracing::info!(%branch_name, ?changes);

    repo.add(&[
        go_constants_path.to_str().unwrap(),
        cargo_toml_path.to_str().unwrap(),
    ])
    .expect("Failed to add files");
    repo.commit(COMMIT_MESSAGE).expect("Failed to commit");
    repo.push(&branch_name).expect("Failed to push");

    // Create PR
    let authed_github = octocrab::OctocrabBuilder::default()
        .personal_token(github_token)
        .build()?;

    // tracing::info!(%owner, %repo_name, ?branch_name, ?pr_message);
    authed_github
        .pulls(owner, repo_name)
        .create(PR_TITLE, branch_name, "master")
        .body(COMMIT_MESSAGE)
        .send()
        .await?;

    Ok(())
}

async fn set_rust_sdk_version(args: &Args, cargo_toml_string: &str) -> eyre::Result<String> {
    let mut cargo_toml = cargo_toml_string.parse::<DocumentMut>()?;
    cargo_toml["workspace"]["package"]["version"] = toml_edit::value(args.sdk_version.clone());
    Ok(cargo_toml.to_string())
}

async fn set_go_sdk_version(
    args: &Args,
    go_constants_source_code_string: &str,
) -> eyre::Result<String> {
    Ok(go_constants_source_code_string.replace(
        go_constants_source_code_string
            .lines()
            .find(|line| line.contains("vanillaSdkVersion ="))
            .unwrap(),
        &format!(
            "const vanillaSdkVersion = \"coralogix-mgmt-sdk-{}\"",
            &args.sdk_version
        ),
    ))
}

#[cfg(test)]
mod test {

    #[tokio::test]
    async fn test_set_rust_sdk_version() {
        let initial_cargo_toml = r#"
        [workspace]
        members = ["cx-api", "cx-sdk"]
        exclude = [
            "examples/actions",
            "examples/alerts",
            "examples/custom-enrichments",
            "examples/dashboards",
            "examples/dataprime-query",
        ]
        resolver = "2"

        [workspace.package]
        version = "0.1.0"
        license = "Apache-2.0"
        edition = "2021"
        repository = "https://github.com/coralogix/coralogix-management-sdk"
        keywords = ["api-bindings", "coralogix", "log", "management", "sdk"]
        categories = ["api-bindings", "development-tools::debugging", "config"]
        authors = ["The YAK team at Coralogix"]
        readme = "README.md"
        homepage = "https://coralogix.com"
        "#;
        let expected_cargo_toml = r#"
        [workspace]
        members = ["cx-api", "cx-sdk"]
        exclude = [
            "examples/actions",
            "examples/alerts",
            "examples/custom-enrichments",
            "examples/dashboards",
            "examples/dataprime-query",
        ]
        resolver = "2"

        [workspace.package]
        version = "0.2.1"
        license = "Apache-2.0"
        edition = "2021"
        repository = "https://github.com/coralogix/coralogix-management-sdk"
        keywords = ["api-bindings", "coralogix", "log", "management", "sdk"]
        categories = ["api-bindings", "development-tools::debugging", "config"]
        authors = ["The YAK team at Coralogix"]
        readme = "README.md"
        homepage = "https://coralogix.com"
        "#;
        let args = super::Args {
            root: String::from("."),
            git_dir: String::from("."),
            github_repo: String::from("."),
            sdk_version: String::from("0.2.1"),
            rust_root: String::from("rust"),
            go_root: String::from("go"),
        };

        let updated_cargo_toml = super::set_rust_sdk_version(&args, initial_cargo_toml)
            .await
            .unwrap();

        assert!(updated_cargo_toml == expected_cargo_toml);
    }

    #[tokio::test]
    async fn test_set_go_sdk_version() {
        let initial_go_constants = r#"
package cxsdk

const sdkVersionHeaderName = "cx-sdk-version"
const sdkLanguageHeaderName = "cx-sdk-language"
const sdkGoVersionHeaderName = "cx-go-version"
const sdkCorrelationIDHeaderName = "cx-correlation-id"
const sdkVersion = "0.1.0"
        "#;
        let expected_go_constants = r#"
package cxsdk

const sdkVersionHeaderName = "cx-sdk-version"
const sdkLanguageHeaderName = "cx-sdk-language"
const sdkGoVersionHeaderName = "cx-go-version"
const sdkCorrelationIDHeaderName = "cx-correlation-id"
const sdkVersion = "0.2.1"
        "#;
        let args = super::Args {
            root: String::from("."),
            git_dir: String::from("."),
            github_repo: String::from("."),
            sdk_version: String::from("0.2.1"),
            go_root: String::from("go"),
            rust_root: String::from("rust"),
        };

        let updated_go_source_code = super::set_go_sdk_version(&args, initial_go_constants)
            .await
            .unwrap();

        assert!(updated_go_source_code == expected_go_constants);
    }
}
