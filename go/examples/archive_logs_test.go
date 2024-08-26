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

import (
	"context"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk"
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v2"

	"github.com/stretchr/testify/assert"
)

func TestArchiveLogs(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewArchiveLogsClient(creator)
	s3Region := "eu-north-1"
	_, setTargetError := c.Update(context.Background(), &v2.SetTargetRequest{
		TargetSpec: &v2.SetTargetRequest_S3{
			S3: &v2.S3TargetSpec{
				Bucket: "yak-2-bucket",
				Region: &s3Region,
			},
		},
	})
	assert.Nil(t, setTargetError)

	_, getTargetError := c.Get(context.Background())

	assert.Nil(t, getTargetError)
}
