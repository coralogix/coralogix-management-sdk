//! Build script for the cx-sdk crate.
use std::process::{Command, Stdio};

const DEFAULT_VERSION_STRING: &str = "unknown_version";

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let rust_compiler_binary = std::env::var("RUSTC").unwrap_or_else(|_| "rustc".to_string());
    let rustc_version_command = Command::new(rust_compiler_binary)
        .arg("--version")
        .stdout(Stdio::piped())
        .spawn();

    if rustc_version_command.is_err() {
        println!("cargo:rustc-env=RUSTC_VERSION={}", DEFAULT_VERSION_STRING);
        return Ok(());
    }

    let rustc_version_output = rustc_version_command.unwrap().wait_with_output();

    if rustc_version_output.is_err() {
        println!("cargo:rustc-env=RUSTC_VERSION={}", DEFAULT_VERSION_STRING);
        return Ok(());
    }

    let rust_version_string = String::from_utf8(rustc_version_output.unwrap().stdout);

    if rust_version_string.is_err() {
        println!("cargo:rustc-env=RUSTC_VERSION={}", DEFAULT_VERSION_STRING);
        return Ok(());
    }

    println!(
        "cargo:rustc-env=RUSTC_VERSION={}",
        rust_version_string.unwrap()
    );
    Ok(())
}
