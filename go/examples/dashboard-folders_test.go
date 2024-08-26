// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
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
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestDashboardFolders(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewDashboardsFoldersClient(creator)

	id := uuid.New().String()

	dashboardFolder := cxsdk.DashboardFolder{
		Id:   wrapperspb.String(id),
		Name: wrapperspb.String(id),
	}

	_, creationError := c.Create(context.Background(), &cxsdk.CreateDashboardFolderRequest{
		Folder: &dashboardFolder,
	})

	assert.Nil(t, creationError)

	_, updateError := c.Replace(context.Background(), &v1.ReplaceDashboardFolderRequest{
		Folder: &cxsdk.DashboardFolder{
			Id:   wrapperspb.String(id),
			Name: wrapperspb.String("updated " + id),
		},
	})

	assert.Nil(t, updateError)

	getDashboardFolderResponse, retrievalError := c.Get(context.Background(), &v1.GetDashboardFolderRequest{
		FolderId: wrapperspb.String(id),
	})

	assert.Nil(t, retrievalError)
	assert.Equal(t, "updated "+id, getDashboardFolderResponse.Folder.Name.Value)

	_, listError := c.List(context.Background())

	assert.Nil(t, listError)

	_, deletionError := c.Delete(context.Background(), &v1.DeleteDashboardFolderRequest{
		FolderId: wrapperspb.String(id),
	})

	assert.Nil(t, deletionError)

}
