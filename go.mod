module github.com/coralogix/coralogix-management-sdk

go 1.24.0

require (
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.0
	github.com/stretchr/testify v1.10.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grpc-ecosystem/grpc-gateway/v2 => github.com/coralogix/grpc-gateway/v2 v2.0.0-20250821080545-ab097d7805fa
