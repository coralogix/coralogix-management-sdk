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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/stretchr/testify/assert"
)

func TestConnectors(t *testing.T) {
	t.Skip("Connectors are not supported in the current version of the SDK")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewNotificationsClient(creator)
	createRes, err := c.CreateConnector(context.Background(), &cxsdk.CreateConnectorRequest{
		Connector: &cxsdk.Connector{
			Type:        cxsdk.ConnectorTypeGenericHTTPS,
			Name:        "TestConnector",
			Description: "This is the connector to use for Notification Center testing.",
			ConnectorConfigs: []*cxsdk.ConnectorConfig{
				{
					OutputSchemaId: "default",
					Fields: []*cxsdk.ConnectorConfigField{
						{FieldName: "url", Template: "https://example@coralogix.com"},
						{FieldName: "method", Template: "post"},
						{FieldName: "additionalHeaders", Template: "{}"},
						{FieldName: "additionalBodyFields", Template: "{}"},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	connectorId := createRes.Connector.Id
	connector, err := c.GetConnector(context.Background(), &cxsdk.GetConnectorRequest{
		Id: *connectorId,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, connector.Connector.Name, "TestConnector")

	_, err = c.DeleteConnector(context.Background(), &cxsdk.DeleteConnectorRequest{
		Id: *connectorId,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestPresets(t *testing.T) {
	t.Skip("Presets are not supported in the current version of the SDK")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewNotificationsClient(creator)
	presetType := cxsdk.PresetTypeCustom
	parentId := "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
	createRes, err := c.CreateCustomPreset(context.Background(), &cxsdk.CreateCustomPresetRequest{
		Preset: &cxsdk.Preset{
			Name:        "TestPreset",
			Description: "This is the preset to use for Notification Center testing.",
			PresetType:  &presetType,
			EntityType:  "alerts",
			Parent: &cxsdk.Preset{
				Id: &parentId,
			},
			ConnectorType: cxsdk.ConnectorTypeGenericHTTPS,
			ConfigOverrides: []*cxsdk.ConfigOverrides{
				{
					ConditionType: &cxsdk.ConditionType{
						Condition: &cxsdk.ConditionTypeMatchEntityType{
							MatchEntityType: &cxsdk.MatchEntityTypeCondition{
								EntityType: "alerts",
							},
						},
					},
					OutputSchemaId: "default",
					MessageConfig: &cxsdk.MessageConfig{
						Fields: []*cxsdk.MessageConfigField{
							{
								Template:  "{ \"test\": \"test\" }",
								FieldName: "headers",
							},
							{
								Template:  "{ \"test\": \"test\" }",
								FieldName: "body",
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	presetId := createRes.Preset.Id
	preset, err := c.GetPreset(context.Background(), &cxsdk.GetPresetRequest{
		Identifier: &cxsdk.PresetIdentifier{
			Value: &cxsdk.PresetIdentifierIDValue{
				Id: *presetId,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, preset.Preset.Name, "TestPreset")

	_, err = c.DeleteCustomPreset(context.Background(), &cxsdk.DeleteCustomPresetRequest{
		Identifier: &cxsdk.PresetIdentifier{
			Value: &cxsdk.PresetIdentifierIDValue{
				Id: *presetId,
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
