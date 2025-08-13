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

//! Build script for the cx-sdk crate.
use std::process::{
    Command,
    Stdio,
};

const DEFAULT_VERSION_STRING: &str = "unknown_version";

#[deprecated(
    since = "1.8.0",
    note = "The SDK will receive a major update in Fall 2025. Use https://docs.coralogix.com/ for minimal changes.  "
)]
fn main() -> Result<(), Box<dyn std::error::Error>> {
    let rust_compiler_binary = std::env::var("RUSTC").unwrap_or_else(|_| "rustc".to_string());
    let rustc_version_command = Command::new(rust_compiler_binary)
        .arg("--version")
        .stdout(Stdio::piped())
        .spawn();

    if rustc_version_command.is_err() {
        println!("cargo:rustc-env=RUSTC_VERSION={DEFAULT_VERSION_STRING}");
        return Ok(());
    }

    let rustc_version_output = rustc_version_command.unwrap().wait_with_output();

    if rustc_version_output.is_err() {
        println!("cargo:rustc-env=RUSTC_VERSION={DEFAULT_VERSION_STRING}");
        return Ok(());
    }

    let rust_version_string = String::from_utf8(rustc_version_output.unwrap().stdout);

    if rust_version_string.is_err() {
        println!("cargo:rustc-env=RUSTC_VERSION={DEFAULT_VERSION_STRING}");
        return Ok(());
    }

    println!(
        "cargo:rustc-env=RUSTC_VERSION={}",
        rust_version_string.unwrap()
    );
    Ok(())
}
