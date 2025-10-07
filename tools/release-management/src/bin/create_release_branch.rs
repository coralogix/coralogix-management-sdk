// Copyright 2025 Coralogix Ltd.
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

use clap::{Parser, command};
use git_cmd::Repo;
use lazy_static::lazy_static;
use toml_edit::DocumentMut;

lazy_static! {
    static ref LTS_MONTHS: Vec<u32> = vec![6, 12];
}

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long, default_value = "../..")]
    root: String,

    #[arg(short, long, default_value = ".")]
    git_dir: String,

    #[arg(short = 'c', long, default_value = "rust")]
    rust_root: String,

    #[arg(short = 't', long, default_value = "go")]
    go_root: String,
}

#[tokio::main]
async fn main() -> eyre::Result<()> {
    let args = Args::parse();
    let tag_string = std::env::var("GITHUB_REF").expect("GITHUB_REF must be set");
    let tag = GitTag::from_str(&tag_string)?;
    if tag.is_lts() && !tag.is_patch() {
        let cargo_toml_path = PathBuf::from(&args.root)
            .join(&args.rust_root)
            .join("Cargo.toml");
        let go_constants_path = PathBuf::from(&args.root)
            .join(&args.go_root)
            .join("constants.go");

        let cargo_toml_string = tokio::fs::read_to_string(&cargo_toml_path).await?;
        let go_constants_source_code_string = tokio::fs::read_to_string(&go_constants_path).await?;
        check_cargo_version_is_the_same_as_tag(&cargo_toml_string, &tag)?;
        check_go_version_is_the_same_as_tag(&go_constants_source_code_string, &tag)?;

        let branch_name = tag_string;
        let repo = Repo::new(&args.git_dir)
            .unwrap_or_else(|_| panic!("Couldn't find git repository at {}", args.git_dir));

        repo.checkout_new_branch(&branch_name)
            .expect("Failed to create branch");
        repo.push(&branch_name).expect("Failed to push");
    }
    Ok(())
}

fn check_cargo_version_is_the_same_as_tag(
    cargo_toml_string: &str,
    tag: &GitTag,
) -> eyre::Result<()> {
    let cargo_toml = cargo_toml_string.parse::<DocumentMut>()?;
    if cargo_toml["workspace"]["package"]["version"].to_string() != tag.to_string() {
        return Err(eyre::Error::msg(
            "Cargo.toml version is not the same as the tag",
        ));
    }
    Ok(())
}

fn check_go_version_is_the_same_as_tag(
    go_constants_source_code_string: &str,
    tag: &GitTag,
) -> eyre::Result<()> {
    let vanilla_sdk_version_string = format!("vanillaSdkVersion = \"{}\"", tag.to_string());
    if !go_constants_source_code_string.contains(&vanilla_sdk_version_string) {
        return Err(eyre::Error::msg("Go version is not the same as the tag"));
    }
    Ok(())
}

struct GitTag {
    major: u32,
    minor: u32,
    patch: u32,
}

impl GitTag {
    fn from_str(s: &str) -> eyre::Result<Self> {
        let version = s[1..].split('.').collect::<Vec<&str>>();
        if version.len() != 3 {
            return Err(eyre::Error::msg("Invalid tag format"));
        }
        Ok(GitTag {
            major: version[0].parse()?,
            minor: version[1].parse()?,
            patch: version[2].parse()?,
        })
    }

    fn is_lts(&self) -> bool {
        LTS_MONTHS.contains(&self.minor)
    }

    fn is_patch(&self) -> bool {
        self.patch > 0
    }

    fn to_string(&self) -> String {
        format!("{}.{}.{}", self.major, self.minor, self.patch)
    }
}
