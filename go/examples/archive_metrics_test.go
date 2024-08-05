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

package examples

// import (
// 	"context"
// 	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
// 	"github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/metrics"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestArchiveMetrics(t *testing.T) {

// 	region, err := cxsdk.CoralogixRegionFromEnv()
// 	assert.Nil(t, err)
// 	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
// 	assert.Nil(t, err)
// 	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
// 	c := cxsdk.NewArchiveMetricsClient(creator)

// 	_, configureTenantError := c.UpdateArchiveMetrics(context.Background(), &metrics.ConfigureTenantRequest{
// 		RetentionPolicy: nil,
// 		StorageConfig: &metrics.ConfigureTenantRequest_S3{
// 			S3: &metrics.S3Config{
// 				Bucket: "coralogix-c4c-eu2-prometheus-data",
// 				Region: "eu-west-1",
// 			},
// 		},
// 	})

// 	assert.Nil(t, configureTenantError)

// 	_, getTenantError := c.GetArchiveMetrics(context.Background())
// 	assert.Nil(t, getTenantError)
// }
