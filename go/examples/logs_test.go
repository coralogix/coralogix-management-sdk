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
	"os"
	"strings"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/stretchr/testify/assert"
)

func TestArchiveLogs(t *testing.T) {
	logsBucket := os.Getenv("LOGS_BUCKET")
	awsRegion := os.Getenv("AWS_REGION")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewArchiveLogsClient(creator)
	_, setTargetError := c.Update(context.Background(), &cxsdk.SetTargetRequest{
		TargetSpec: &cxsdk.SetS3TargetRequest{
			S3: &cxsdk.Target{
				Bucket: logsBucket,
				Region: &awsRegion,
			},
		},
	})
	assertNilAndPrintError(t, setTargetError)

	_, getTargetError := c.Get(context.Background())

	assertNilAndPrintError(t, getTargetError)
}

func TestDataSets(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	c := cxsdk.NewDataSetClient(creator)

	raw, e := os.ReadFile("date-to-day-of-the-week.csv")
	assertNilAndPrintError(t, e)
	textual := string(raw)
	name := uuid.NewString()
	data, e := c.Create(context.Background(), &cxsdk.CreateDataSetRequest{
		Name:        wrapperspb.String(strings.ReplaceAll(name, "-", "")),
		Description: wrapperspb.String("My custom enrichment description"),
		File: &cxsdk.File{
			Name:      wrapperspb.String("date-to-day-of-the-week"),
			Extension: wrapperspb.String("csv"),
			Content: &cxsdk.FileTextual{
				Textual: wrapperspb.String(textual),
			},
		},
	})
	assertNilAndPrintError(t, e)

	updated, e := c.Update(context.Background(), &cxsdk.UpdateDataSetRequest{
		CustomEnrichmentId: wrapperspb.UInt32(data.CustomEnrichment.Id),
		Description:        wrapperspb.String("My updated enrichment description"),
		File: &cxsdk.File{
			Name:      wrapperspb.String("date-to-day-of-the-week"),
			Extension: wrapperspb.String("csv"),
			Content: &cxsdk.FileTextual{
				Textual: wrapperspb.String(textual),
			},
		},
	})
	assertNilAndPrintError(t, e)
	fetched, e := c.Get(context.Background(), &cxsdk.GetDataSetRequest{
		Id: wrapperspb.UInt32(updated.CustomEnrichment.Id),
	})
	assertNilAndPrintError(t, e)

	assert.Equal(t, updated.CustomEnrichment.Description, fetched.CustomEnrichment.Description)
	assert.Equal(t, data.CustomEnrichment.Version+1, fetched.CustomEnrichment.Version)
	c.Delete(context.Background(), &cxsdk.DeleteDataSetRequest{
		CustomEnrichmentId: wrapperspb.UInt32(fetched.CustomEnrichment.Id),
	})
}
