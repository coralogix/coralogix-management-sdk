# Coralogix Management SDK for Go

The Go SDK talks to Coralogix over REST.

Use `github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk` for API clients.
The `cxsdk` package in `go/` keeps the SCIM Users HTTP client and region helpers.

Use `make` to build and test.

# Important Makefile Targets

- `make build` - compiles the Go packages
- `make test` - run the remaining non-OpenAPI examples (SCIM Users)
- `make test-openapi` - run OpenAPI examples. To run a single example, use `go test -run <NAME_OF_TEST> ./openapi/examples`.
