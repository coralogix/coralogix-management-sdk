use git_cmd::Repo;
use protofetch::Protofetch;
use time::format_description;
use toml_edit::{value, DocumentMut};

const PROTOFETCH_PATH: &str = "../../protofetch.toml";
const PROTOFETCH_PATH_LOCK: &str = "../../protofetch.lock";
const PROTOS_DIR: &str = "../../proto";

const PROTOFETCH_META_KEYS: [&str; 3] = ["name", "description", "proto_out_dir"];
const ALLOWED_CHANGES: [&str; 3] = [PROTOFETCH_PATH, PROTOFETCH_PATH_LOCK, PROTOS_DIR];

#[tokio::main]
async fn main() -> eyre::Result<()> {
    let toml = tokio::fs::read_to_string(PROTOFETCH_PATH).await?;
    let mut protofetch_descriptor = toml
        .parse::<DocumentMut>()
        .expect("Protofetch.toml not found");
    let github_token = std::env::var("GITHUB_TOKEN").expect("A GITHUB_TOKEN must be set");
    let authed_github = octocrab::OctocrabBuilder::default()
        .personal_token(github_token)
        .build()?;

    let mut changes = false;
    for (k, val) in protofetch_descriptor.iter_mut() {
        let name = k.get();

        if PROTOFETCH_META_KEYS.contains(&name) {
            continue;
        }

        let url = val["url"].as_str().expect("Value must be a string");
        let (owner, repo) = owner_repo_from_url(url).expect("Invalid URL");
        let revision = val["revision"].as_str().expect("Value must be a string");

        let release = authed_github
            .repos(owner, repo)
            .releases()
            .get_latest()
            .await;
        match release {
            Ok(release) => {
                if revision != release.tag_name {
                    val["revision"] = value(release.tag_name);
                    changes = true;
                }
            }
            Err(e) => {
                eprintln!("[{}] {}/releases returned an error: {:?} ", name, url, e);
            }
        }
    }
    // No changes
    if !changes {
        return Ok(());
    }
    tokio::fs::write(PROTOFETCH_PATH, protofetch_descriptor.to_string()).await?;
    tokio::fs::remove_dir_all(PROTOS_DIR).await?;

    let protos = Protofetch::builder()
        .module_file_name(PROTOFETCH_PATH)
        .try_build()
        .expect("Protofetch failed");

    protos
        .fetch(protofetch::LockMode::Recreate)
        .expect("Protofetch failed");

    let _ = match Repo::new(".") {
        Ok(repo) => {
            let format =
                format_description::parse("[year]-[month]-[day]").expect("Failed to parse format");
            let branch_name = format!(
                "update-{}",
                time::OffsetDateTime::now_utc().format(&format)?
            );
            let changes = repo.changes(|_| true).expect("Failed to add changes");
            repo.checkout_new_branch(&branch_name)
                .expect("Failed to create branch");
            repo.add(&ALLOWED_CHANGES).expect("Failed to add files");
            repo.commit("Update protofetch.toml")
                .expect("Failed to commit");
            repo.push(&branch_name).expect("Failed to push");
            authed_github
                .pulls("coralogix", "coralogix-management-sdk")
                .create("Updating API descriptions", branch_name, "master")
                .body("body")
                .send()
                .await?;
        }

        Err(e) => panic!("failed to open: {}", e),
    };
    Ok(())
}

fn owner_repo_from_url(url: &str) -> Option<(&str, &str)> {
    let url = url
        .trim_end_matches(".git")
        .trim_start_matches("github.com/");
    let mut parts = url.split('/').rev();
    let repo = parts.next()?;
    let owner = parts.next()?;
    Some((owner, repo))
}
