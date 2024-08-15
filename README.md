# The Coralogix Management SDK

# Building

## Prerequisites

- `protoc` - for compiling protobuf files
- `Make` - for Go development

Please refer to the individual languages on how to build the SDKs

## Configuration

The SDK can be configured using environment variables:

- `CORALOGIX_API_KEY`: The API key that is used for all interactions. Note that it has to have appropriate permissions. Read the [docs](https://coralogix.com/docs/api-keys/) for more information
- `CORALGOIX_REGION`: The region/cluster to connect to as a shorthand (EU2, AP1, etc. read more [here](https://coralogix.com/docs/coralogix-domain/)). 

Please note that for all examples these variables have to be set to a valid cluster and API key.

# Docs

Check out the corresponding Github pages!

# Protobuf

The `proto/` directory contains all protobuf files for generating your own SDK if necessary. They are also the basis for all SDKs in this repository. 

# Contributing

We welcome contributions! Be it docs, code, or even just opening an issue will help improve this SDK for everyone. To learn more check out [CONTRIBUTING.md](). 

To add code to this repository you are required to sign our Contributor License Agreement. 

# License

[Apache-2](LICENSE)