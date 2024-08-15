# Coralogix Management SDK for Rust

The Rust SDK is an abstraction over the protobuf files in the root of the repository. Generating the .rs files happens during the build and the output is stored in the target directory (locally). Refer to the examples directory for how to connect to the individual services.

Use cargo to build and test:

`cargo build` - builds the SDK

`cd examples && cargo test` -  runs all examples

`cd examples && cargo test -p archive-logs` -  runs the example `archive-logs`