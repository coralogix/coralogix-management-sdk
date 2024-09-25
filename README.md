[![](https://img.shields.io/crates/v/cx-sdk)](https://crates.io/crates/cx-sdk)
![](https://img.shields.io/crates/dv/cx-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/coralogix/coralogix-management-sdk.svg)](https://pkg.go.dev/github.com/coralogix/coralogix-management-sdk)

# 🪵 The Coralogix Management SDK 

🌟 [Master Docs](https://coralogix.github.io/coralogix-management-sdk) 🌟

# Building

## Prerequisites

- `protoc` - for compiling protobuf files
- `Make` - for Go development

Please refer to the individual languages on how to build the SDKs

## Configuration

The SDK can be configured using environment variables:

- `CORALOGIX_TEAM_API_KEY`: The API key that is used for all team-level interactions. Note that it has to have appropriate permissions. Read the [docs](https://coralogix.com/docs/api-keys/) for more information.
- `CORALOGIX_USER_API_KEY`: The API key that is used for all user-level interactions. Note that it has to have appropriate permissions. Read the [docs](https://coralogix.com/docs/api-keys/) for more information.
- `CORALGOIX_REGION`: The region/cluster to connect to as a shorthand (EU2, AP1, etc. read more [here](https://coralogix.com/docs/coralogix-domain/)). 

Furthermore, if you want to run the examples locally, you're going to to have set the following environment variables:
- `AWS_REGION`: The aws region that you wanna use in the examples, eg. `eu-north-1`.
- `METRICS_BUCKET`: The name of the S3 bucket that you want to use for the archive-metrics example.
- `LOGS_BUCKET`: The name of the S3 bucket that you want to use for the archive-logs example.

Please note that for all examples these variables have to be set to a valid cluster and API key.

# Protobuf

The `proto/` directory contains all protobuf files for generating your own SDK if necessary. They are also the basis for all SDKs in this repository. 

# Contributing

We welcome contributions! Be it docs, code, or even just opening an issue will help improve this SDK for everyone. To learn more check out [CONTRIBUTING.md](). 

To add code to this repository you are required to sign our Contributor License Agreement. 

# License

[Apache-2](LICENSE)