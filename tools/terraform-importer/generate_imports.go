//go:build !jsonhcl

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
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
)

var terraformEnvironmentAliasToSdkEnvironment = map[string]string{
	"APAC1":   "ap1",
	"AP1":     "ap1",
	"APAC2":   "ap2",
	"AP2":     "ap2",
	"APAC3":   "ap3",
	"AP3":     "ap3",
	"EUROPE1": "eu1",
	"EU1":     "eu1",
	"EUROPE2": "eu2",
	"EU2":     "eu2",
	"USA1":    "us1",
	"US1":     "us1",
	"USA2":    "us2",
	"US2":     "us2",
}

// Resource represents a resource in the Terraform state file
type Resource struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Instances []struct {
		Attributes struct {
			ID string `json:"id"`
		} `json:"attributes"`
	} `json:"instances"`
}

// TFState represents the structure of the Terraform state file
type TFState struct {
	Resources []Resource `json:"resources"`
}

// findStateFile searches for a .tfstate file in the specified folder
func findStateFile(folderPath string) (string, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return "", fmt.Errorf("error reading folder: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".tfstate" {
			return filepath.Join(folderPath, file.Name()), nil
		}
	}

	return "", errors.New("no .tfstate file found in the specified folder")
}

// generateImports reads a Terraform state file and generates an imports.tf file
func generateImportsFromState(tfstatePath string, outputPath string) error {
	tfstateData, err := os.ReadFile(tfstatePath)
	if err != nil {
		return fmt.Errorf("error reading tfstate file: %v", err)
	}

	var tfstate TFState
	err = json.Unmarshal(tfstateData, &tfstate)
	if err != nil {
		return fmt.Errorf("error parsing tfstate JSON: %v", err)
	}

	importsContent := ""

	for _, resource := range tfstate.Resources {
		if strings.HasPrefix(resource.Type, "coralogix_") {
			for _, instance := range resource.Instances {
				importsContent += fmt.Sprintf(`import {
  to = %s.%s
  id = "%s"
}

`, resource.Type, resource.Name, instance.Attributes.ID)
			}
		}
	}

	err = os.WriteFile(outputPath, []byte(importsContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing imports.tf file: %v", err)
	}

	return nil
}

type IdAndName struct {
	Id   string
	Name string
}

func main() {
	resourceType := flag.String("type", "", "Type of the resource to import")
	folderPath := flag.String("folder", "", "Path to the folder containing the .tfstate file")
	outputPath := flag.String("output", "imports.tf", "Path to the output file")
	flag.Parse()

	if *resourceType != "" {
		var idsAndNames []IdAndName
		apiKey := os.Getenv("CORALOGIX_API_KEY")
		region := strings.ToUpper(os.Getenv("CORALOGIX_ENV"))
		mappedRegion := terraformEnvironmentAliasToSdkEnvironment[region]

		switch *resourceType {
		case "alert":
			alertClient := cxsdk.NewAlertsClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			alerts, err := alertClient.List(context.Background(), &cxsdk.ListAlertDefsRequest{})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, alert := range alerts.GetAlertDefs() {
				alertName := convertToTerraformResourceName(alert.GetAlertDefProperties().GetName().GetValue())
				idsAndNames = append(idsAndNames, IdAndName{Id: alert.GetId().GetValue(), Name: alertName})
			}
		case "archive_logs":
			idsAndNames = append(idsAndNames, IdAndName{Id: "archive_logs", Name: "archive_logs"})
		case "archive_metrics":
			idsAndNames = append(idsAndNames, IdAndName{Id: "archive_metrics", Name: "archive_metrics"})
		case "archive_retentions":
			idsAndNames = append(idsAndNames, IdAndName{Id: "archive_retentions", Name: "archive_retentions"})
		case "custom_role":
			rolesClients := cxsdk.NewRolesClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			roles, err := rolesClients.List(context.Background(), &cxsdk.ListCustomRolesRequest{})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, role := range roles.GetRoles() {
				roleName := convertToTerraformResourceName(role.GetName())
				idsAndNames = append(idsAndNames, IdAndName{Id: strconv.Itoa(int(role.RoleId)), Name: roleName})
			}
		case "dashboard":
			dashboardClient := cxsdk.NewDashboardsClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			dashboards, err := dashboardClient.List(context.Background())
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, dashboard := range dashboards.GetItems() {
				dashboardName := convertToTerraformResourceName(dashboard.GetName().GetValue())
				idsAndNames = append(idsAndNames, IdAndName{Id: dashboard.GetId().GetValue(), Name: dashboardName})
			}
		case "dashboards_folder":
			dashboardsFolderClient := cxsdk.NewDashboardsFoldersClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			dashboardsFolders, err := dashboardsFolderClient.List(context.Background())
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, dashboardsFolder := range dashboardsFolders.GetFolder() {
				dashboardsFolderName := convertToTerraformResourceName(dashboardsFolder.GetName().GetValue())
				idsAndNames = append(idsAndNames, IdAndName{Id: dashboardsFolder.GetId().GetValue(), Name: dashboardsFolderName})
			}
		case "events2metric":
			events2metricsClient := cxsdk.NewEvents2MetricsClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			events2metrics, err := events2metricsClient.List(context.Background())
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, events2metric := range events2metrics.E2M {
				events2metricName := convertToTerraformResourceName(events2metric.GetName().GetValue())
				idsAndNames = append(idsAndNames, IdAndName{Id: events2metric.GetId().GetValue(), Name: events2metricName})
			}
		case "group":
			groupClient := cxsdk.NewGroupsClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			groups, err := groupClient.List(context.Background(), &cxsdk.GetTeamGroupsRequest{})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, group := range groups.GetGroups() {
				groupName := convertToTerraformResourceName(group.GetName())
				idsAndNames = append(idsAndNames, IdAndName{Id: strconv.Itoa(int(group.GroupId.Id)), Name: groupName})
			}
		case "recording_rules_groups_set":
			recordingRulesGroupsSetClient := cxsdk.NewRecordingRuleGroupSetsClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			recordingRulesGroupsSets, err := recordingRulesGroupsSetClient.List(context.Background())
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, recordingRulesGroupsSet := range recordingRulesGroupsSets.GetSets() {
				recordingRulesGroupsSetName := convertToTerraformResourceName(recordingRulesGroupsSet.GetName())
				idsAndNames = append(idsAndNames, IdAndName{Id: recordingRulesGroupsSet.GetId(), Name: recordingRulesGroupsSetName})
			}
		case "scope":
			scopesClient := cxsdk.NewScopesClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			scopes, err := scopesClient.List(context.Background(), &cxsdk.GetTeamScopesRequest{})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, scope := range scopes.GetScopes() {
				scopeName := convertToTerraformResourceName(scope.DisplayName)
				idsAndNames = append(idsAndNames, IdAndName{Id: scope.GetId(), Name: scopeName})
			}
		case "tco_policies_logs":
			idsAndNames = []IdAndName{{Id: " ", Name: "tco_policies_logs"}}
		case "tco_policies_traces":
			idsAndNames = []IdAndName{{Id: " ", Name: "tco_policies_traces"}}
		case "webhook":
			webhookClient := cxsdk.NewWebhooksClient(cxsdk.NewCallPropertiesCreator(mappedRegion, cxsdk.NewAuthContext(apiKey, apiKey)))
			webhooks, err := webhookClient.List(context.Background(), &cxsdk.ListAllOutgoingWebhooksRequest{})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			for _, webhook := range webhooks.GetDeployed() {
				webhookName := convertToTerraformResourceName(webhook.GetName().GetValue())
				idsAndNames = append(idsAndNames, IdAndName{Id: webhook.GetId().GetValue(), Name: webhookName})
			}
		default:
			log.Print("Error: Invalid resource type")
		}

		err := generateImportsFromIds(*resourceType, *outputPath, idsAndNames)
		if err != nil {
			log.Fatal("Error generating imports.tf: %v\n", err)
		}
	} else {
		if *folderPath == "" {
			log.Fatal("Error: Please provide a folder path using the -folder flag")
		}

		tfstatePath, err := findStateFile(*folderPath)
		if err != nil {
			log.Fatal("Error: %v\n", err)
		}

		err = generateImportsFromState(tfstatePath, *outputPath)
		if err != nil {
			fmt.Printf("Error generating imports.tf: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("`imports.tf` file has been generated at: %s\n", outputPath)
}

func convertToTerraformResourceName(input string) string {
	input = strings.ToLower(input)

	re := regexp.MustCompile(`[^a-z0-9_\-]`)
	input = re.ReplaceAllString(input, "_")

	re = regexp.MustCompile(`^[^a-z]+`)
	input = re.ReplaceAllString(input, "")

	re = regexp.MustCompile(`_+`)
	input = re.ReplaceAllString(input, "_")

	return input
}

func generateImportsFromIds(resourceType, outputFilePath string, idsAndNames []IdAndName) error {
	importsContent := ""
	nameCounts := make(map[string]int)
	uniqueNames := make(map[string]string)

	for _, idAndName := range idsAndNames {
		originalName := idAndName.Name
		name := originalName

		for {
			if _, exists := uniqueNames[name]; !exists {
				break
			}
			suffixPattern := regexp.MustCompile(`_(\d+)$`)
			match := suffixPattern.FindStringSubmatch(name)
			if len(match) > 0 {
				num, _ := strconv.Atoi(match[1])
				name = fmt.Sprintf("%s_%d", strings.TrimSuffix(name, fmt.Sprintf("_%d", num)), num+1)
			} else {
				name = fmt.Sprintf("%s_2", name)
			}
		}

		uniqueNames[name] = idAndName.Id
		nameCounts[originalName]++ // Increment the counter for the original name

		importsContent += fmt.Sprintf(`import {
  to = coralogix_%s.%s
  id = "%s"
}

`, resourceType, name, idAndName.Id)
	}

	err := os.WriteFile(outputFilePath, []byte(importsContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing imports.tf file: %v", err)
	}

	return nil
}
