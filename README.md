[![](https://img.shields.io/crates/v/cx-sdk)](https://crates.io/crates/cx-sdk)
![](https://img.shields.io/crates/dv/cx-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/coralogix/coralogix-management-sdk.svg)](https://pkg.go.dev/github.com/coralogix/coralogix-management-sdk)

# 🪵 The Coralogix Management SDK 

🌟 [Master Docs](https://coralogix.github.io/coralogix-management-sdk) 🌟

## REST and OpenAPI

Our APIs are also avaialable over REST through a variety of OpenAPI endpoints. Read more over at the [OpenAPI Docs](https://docs.coralogix.com)

*These endpoints follow the same versioning schemes mentioned below.*

# Versioning

The SDK follows a rolling release schedule with one LTS version once a year (in June). Revisions allow additional fixes that add hotfixes for immediate bugs, security fixes, etc. The remaining issue fixes are rolled up in the following month’s release. The exception is the LTS release which will contain relevant fixes. 

### LTS

Every year in June the x.6.x version is considered LTS and will be backwards compatible for a year. This is the version recommended for any stable product (Terraform Provider, …) and will be maintained without major (API-related) changes.

### Examples

Here are a few examples of (anticipated) versions:

- 0.12.0 → Year 0 (= 2024), month (12), revision (0)
- 1.1.0 → January 2025 without any additional revisions
- 1.6.10 → June 2025 (LTS) with 10 revisions

# Building

## Prerequisites

- `protoc` - for compiling protobuf files
- `Make` - for Go development

The file `proto-toolchain-versions.txt` contains the versions of the protobuf tools that you need to contribute to the SDK: please make sure these are the versions that you have on your system.

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

# Examples
You'll find some example use-cases of the SDK in the `examples` folder under each programming language. Please, keep in mind that `aaa` stands for "authentication, authorization and accounting", so you should look there if you're looking for an example related to one of those areas.

# Protobuf

The `proto/` directory contains all protobuf files for generating your own SDK if necessary. They are also the basis for all SDKs in this repository. 

# Contributing

We welcome contributions! Be it docs, code, or even just opening an issue will help improve this SDK for everyone. To learn more check out [CONTRIBUTING.md](). 

To add code to this repository you are required to sign our Contributor License Agreement. 

# License

[Apache-2](LICENSE)