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

	archiveLogs "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v2"
)

// SetTargetRequest is a request to set the archive logs target.
type SetTargetRequest = archiveLogs.S3TargetServiceSetTargetRequest

// SetTargetResponse is a request to set the archive logs target to IBM COS.
type SetTargetResponse = archiveLogs.S3TargetServiceSetTargetResponse

// GetTargetResponse is a response to get the archive logs target.
type GetTargetResponse = archiveLogs.S3TargetServiceGetTargetResponse

// SetS3TargetRequest is an S3 target spec.
type SetS3TargetRequest = archiveLogs.S3TargetServiceSetTargetRequest_S3

// Target is a target for storing archive logs.
type Target = archiveLogs.S3TargetSpec


const archiveLogsFeatureGroupID = "logs"

// RPC names.
const (
	ArchiveLogsGetTargetRPC      = archiveLogs.TargetService_GetTarget_FullMethodName
	ArchiveLogsSetTargetRPC      = archiveLogs.TargetService_SetTarget_FullMethodName
	ArchiveLogsValidateTargetRPC = archiveLogs.TargetService_ValidateTarget_FullMethodName
)

// IbmBucketType values.
const (
	IbmBucketTypeUnspecified = archiveLogs.IbmBucketType_IBM_BUCKET_TYPE_UNSPECIFIED
	IbmBucketTypeExternal    = archiveLogs.IbmBucketType_IBM_BUCKET_TYPE_EXTERNAL
	IbmBucketTypeInternal    = archiveLogs.IbmBucketType_IBM_BUCKET_TYPE_INTERNAL
)

// ArchiveLogsClient is a client for the Coralogix Archive Logs API.
type ArchiveLogsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Update updates the archive logs target.
func (c ArchiveLogsClient) Update(ctx context.Context, req *SetTargetRequest) (*SetTargetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveLogs.NewS3TargetServiceClient(conn)

	response, err := client.SetTarget(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ArchiveLogsSetTargetRPC, archiveLogsFeatureGroupID)
	}
	return response, nil
}

// Get gets the archive logs target.
func (c ArchiveLogsClient) Get(ctx context.Context) (*GetTargetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveLogs.NewS3TargetServiceClient(conn)

	response, err := client.GetTarget(callProperties.Ctx, &S3TargetServiceGetTargetRequest{}, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ArchiveLogsGetTargetRPC, archiveLogsFeatureGroupID)
	}
	return response, nil
}

// NewArchiveLogsClient creates a new archive logs client.
func NewArchiveLogsClient(c *CallPropertiesCreator) *ArchiveLogsClient {
	return &ArchiveLogsClient{callPropertiesCreator: c}
}
