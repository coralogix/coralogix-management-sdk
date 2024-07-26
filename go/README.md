# Coralogix Management SDK for Go

The Golang SDK for Coralogix uses (mostly) GRPC for interacting with the SaaS platform. Use `make` (install if required) to build and test.

## Important Makefile Targets

- `make build` - builds the binary in the default architecture 
- `make proto-clean` - clean the Go files created from protobuf
- `make proto-compile` - compile the Go files from protobuf
- `make test` - run examples/tests
