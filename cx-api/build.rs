fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed=../Cargo.lock");
    println!("cargo:rerun-if-changed=protofetch.toml");
    println!("cargo:rerun-if-changed=protofetch.lock");

    #[cfg(feature = "gcp")]
    gcp_proto()?;

    Ok(())
}

#[cfg(feature = "gcp")]
fn gcp_proto() -> Result<(), Box<dyn std::error::Error>> {
    use protofetch::{LockMode, Protofetch};

    Protofetch::builder()
        .try_build()?
        .fetch(if is_ci::cached() {
            LockMode::Locked
        } else {
            LockMode::Update
        })?;
    tonic_build::configure()
        .include_file("resource_manager_grpc.rs")
        .compile_well_known_types(false)
        .build_server(false)
        .build_client(true)
        .compile(
            &["target/proto/google/cloud/resourcemanager/v3/projects.proto"],
            &["target/proto"],
        )?;
    Ok(())
}