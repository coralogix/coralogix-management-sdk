// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cxsdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// CallPropertiesCreator is a struct that creates CallProperties objects.
type CallPropertiesCreator struct {
	coraglogixRegion string
	teamsLevelAPIKey string
	userLevelAPIKey  string
	correlationID    string
	sdkVersion       string
	//allowRetry bool
}

// CallProperties is a struct that holds the context, connection, and call options for a gRPC call.
type CallProperties struct {
	Ctx         context.Context
	Connection  *grpc.ClientConn
	CallOptions []grpc.CallOption
}

// GetTeamsLevelCallProperties returns a new CallProperties object built from a team-level API key. It essentially prepares the context, connection, and call options for a gRPC call.
func (c CallPropertiesCreator) GetTeamsLevelCallProperties(ctx context.Context) (*CallProperties, error) {
	ctx = createContext(ctx, c.teamsLevelAPIKey, c.correlationID, c.sdkVersion)

	endpoint := CoralogixGrpcEndpointFromRegion(strings.ToLower(c.coraglogixRegion))
	conn, err := createSecureConnection(endpoint)
	if err != nil {
		return nil, err
	}

	callOptions := createCallOptions()

	return &CallProperties{Ctx: ctx, Connection: conn, CallOptions: callOptions}, nil
}

// GetUserLevelCallProperties returns a new CallProperties object built from a user-level API key. It essentially prepares the context, connection, and call options for a gRPC call.
func (c CallPropertiesCreator) GetUserLevelCallProperties(ctx context.Context) (*CallProperties, error) {
	ctx = createContext(ctx, c.userLevelAPIKey, c.correlationID, c.sdkVersion)

	endpoint := CoralogixGrpcEndpointFromRegion(strings.ToLower(c.coraglogixRegion))
	conn, err := createSecureConnection(endpoint)
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
	callOptions = append(callOptions, grpc.MaxCallRecvMsgSize(50*1024*1024)) // 50MB
	callOptions = append(callOptions, grpc.MaxCallSendMsgSize(50*1024*1024)) // 50MB
	return callOptions
}

func createSecureConnection(targetURL string) (*grpc.ClientConn, error) {
	// We cannot use grpc.NewClient because it doesn't work behind a proxy: https://github.com/grpc/grpc-go/releases/tag/v1.69.0
	return grpc.Dial(targetURL,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
}

func createContext(ctx context.Context, apiKey string, corrleationID string, sdkVersion string) context.Context {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", apiKey), sdkVersionHeaderName: sdkVersion, sdkLanguageHeaderName: "go", sdkGoVersionHeaderName: runtime.Version(), sdkCorrelationIDHeaderName: corrleationID})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

// NewCallPropertiesCreator creates a new CallPropertiesCreator object.
func NewCallPropertiesCreator(region string, authContext AuthContext) *CallPropertiesCreator {
	return &CallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("sdk-", vanillaSdkVersion),
	}
}

// NewCallPropertiesCreatorTerraform creates a new CallPropertiesCreator object, specifying which version of the Terraform Operator is being used.
func NewCallPropertiesCreatorTerraform(region string, authContext AuthContext, terraformProviderVersion string) *CallPropertiesCreator {
	return &CallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("terraform-", terraformProviderVersion),
	}
}

// NewCallPropertiesCreatorOperator creates a new CallPropertiesCreator object, specifying which version of the Operator Operator is being used.
func NewCallPropertiesCreatorOperator(region string, authContext AuthContext, cxOperator string) *CallPropertiesCreator {
	return &CallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("cxo-", cxOperator),
	}
}

// NewCallPropertiesCreatorCustomSdk creates a new CallPropertiesCreator object with a custom version.
func NewCallPropertiesCreatorCustomSdk(region string, authContext AuthContext, customSdkVersion string) *CallPropertiesCreator {
	return &CallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("custom-", customSdkVersion),
	}
}
