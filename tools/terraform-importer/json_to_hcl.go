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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// formatValue formats values for Terraform, handling multi-line strings with heredoc syntax.
func formatValue(value interface{}) string {
	switch v := value.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64, int:
		return fmt.Sprintf("%v", v)
	case string:
		// Check for multi-line strings
		if strings.Contains(v, "\n") {
			return fmt.Sprintf("<<EOF\n%s\nEOF", v)
		}
		// Escape special characters for HCL
		escaped := strings.ReplaceAll(v, `\`, `\\`)
		escaped = strings.ReplaceAll(escaped, `"`, `\"`)
		return fmt.Sprintf(`"%s"`, escaped)
	case []interface{}:
		var items []string
		for _, item := range v {
			items = append(items, formatValue(item))
		}
		return fmt.Sprintf("[%s]", strings.Join(items, ", "))
	case map[string]interface{}:
		var lines []string
		for key, val := range v {
			lines = append(lines, fmt.Sprintf(`  %s = %s`, key, formatValue(val)))
		}
		return fmt.Sprintf("{\n%s\n}", strings.Join(lines, "\n"))
	default:
		return `""`
	}
}

// processBlock processes key-value pairs, handling lists and maps.
func processBlock(key string, value interface{}) string {
	switch v := value.(type) {
	case map[string]interface{}:
		// Map: Output as an attribute (`key = { ... }`)
		return fmt.Sprintf("  %s = %s", key, formatValue(v))
	case []interface{}:
		if len(v) > 0 {
			if _, isSimple := v[0].(string); isSimple {
				// List of simple values
				return fmt.Sprintf("  %s = %s", key, formatValue(v))
			}

			// List of maps: Format as a list of maps
			if _, isMap := v[0].(map[string]interface{}); isMap {
				var items []string
				for _, item := range v {
					if itemMap, ok := item.(map[string]interface{}); ok {
						items = append(items, formatValue(itemMap))
					}
				}
				return fmt.Sprintf("  %s = [%s]", key, strings.Join(items, ", "))
			}
		}
		// Empty list
		return fmt.Sprintf("  %s = []", key)
	default:
		// Scalar value: Output as an attribute
		return fmt.Sprintf("  %s = %s", key, formatValue(v))
	}
}

// handleArchiveRetentions applies special logic for archive_retentions resource
// The first retention is the default retention and can't have a name set
func handleArchiveRetentions(resourceMap map[string]interface{}) map[string]interface{} {
	if retentions, ok := resourceMap["retentions"].([]interface{}); ok && len(retentions) > 0 {
		// Check if the first retention has a name and remove it
		if firstRetention, ok := retentions[0].(map[string]interface{}); ok {
			if _, hasName := firstRetention["name"]; hasName {
				// Create a copy of the first retention without the name field
				newFirstRetention := make(map[string]interface{})
				for key, value := range firstRetention {
					if key != "name" {
						newFirstRetention[key] = value
					}
				}

				// Create a new retentions slice with the modified first retention
				newRetentions := make([]interface{}, len(retentions))
				newRetentions[0] = newFirstRetention
				copy(newRetentions[1:], retentions[1:])

				// Create a copy of the resource map with the new retentions
				newResourceMap := make(map[string]interface{})
				for key, value := range resourceMap {
					if key == "retentions" {
						newResourceMap[key] = newRetentions
					} else {
						newResourceMap[key] = value
					}
				}
				return newResourceMap
			}
		}
	}
	return resourceMap
}

// generateTerraform converts parsed JSON into a Terraform HCL configuration.
func generateTerraform(jsonData map[string]interface{}) string {
	var terraformLines []string

	if resources, ok := jsonData["resource"].(map[string]interface{}); ok {
		for resourceType, resourceData := range resources {
			if resourceMap, ok := resourceData.(map[string]interface{}); ok {
				for resourceName, resourceList := range resourceMap {
					if resourceArray, ok := resourceList.([]interface{}); ok {
						for _, resource := range resourceArray {
							if resourceMap, ok := resource.(map[string]interface{}); ok {
								// Apply special handling for archive_retentions
								if resourceType == "coralogix_archive_retentions" {
									resourceMap = handleArchiveRetentions(resourceMap)
								}

								terraformLines = append(terraformLines, fmt.Sprintf(`resource "%s" "%s" {`, resourceType, resourceName))
								for key, val := range resourceMap {
									terraformLines = append(terraformLines, processBlock(key, val))
								}
								terraformLines = append(terraformLines, "}\n")
							}
						}
					}
				}
			}
		}
	}

	return strings.Join(terraformLines, "\n")
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run json_to_hcl.go <path_to_json_file> <path_to_output_file>")
	}

	jsonFilePath := os.Args[1]
	tfFilePath := os.Args[2]

	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal("Failed to read JSON file: %v\n", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		log.Fatal("Failed to parse JSON: %v\n", err)
	}

	terraformOutput := generateTerraform(parsed)

	err = os.WriteFile(tfFilePath, []byte(terraformOutput), 0644)
	if err != nil {
		log.Fatal("Failed to write Terraform file: %v\n", err)
	}

	fmt.Printf("Terraform configuration written to %s\n", tfFilePath)
}
