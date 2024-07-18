package cxsdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// CallPropertiesCreator is a struct that creates CallProperties objects.
type CallPropertiesCreator struct {
	targetURL string
	apiKey    string
	//allowRetry bool
}

// CallProperties is a struct that holds the context, connection, and call options for a gRPC call.
type CallProperties struct {
	Ctx         context.Context
	Connection  *grpc.ClientConn
	CallOptions []grpc.CallOption
}

// GetCallProperties returns a new CallProperties object. It essentially prepares the context, connection, and call options for a gRPC call.
func (c CallPropertiesCreator) GetCallProperties(ctx context.Context) (*CallProperties, error) {
	ctx = createAuthContext(ctx, c.apiKey)

	conn, err := createSecureConnection(c.targetURL)
	if err != nil {
		return nil, err
	}

	callOptions := createCallOptions()

	return &CallProperties{Ctx: ctx, Connection: conn, CallOptions: callOptions}, nil
}

func createCallOptions() []grpc.CallOption {
	var callOptions []grpc.CallOption
	callOptions = append(callOptions, grpc_retry.WithMax(5))
	callOptions = append(callOptions, grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)))
	return callOptions
}

func createSecureConnection(targetURL string) (*grpc.ClientConn, error) {
	return grpc.NewClient(targetURL,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
}

func createAuthContext(ctx context.Context, apiKey string) context.Context {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", apiKey)})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

// NewCallPropertiesCreator creates a new CallPropertiesCreator object.
func NewCallPropertiesCreator(targetURL, apiKey string) *CallPropertiesCreator {
	return &CallPropertiesCreator{
		targetURL: targetURL,
		apiKey:    apiKey,
	}
}
