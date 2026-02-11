//go:build jsonhcl

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

	"github.com/PaesslerAG/jsonpath"
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
		if strings.Contains(v, "\n") {
			return fmt.Sprintf("<<EOF\n%s\nEOF", v)
		}
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
			quotedKey := key
			if strings.Contains(key, ".") {
				quotedKey = fmt.Sprintf(`"%s"`, key)
			}
			lines = append(lines, fmt.Sprintf(`  %s = %s`, quotedKey, formatValue(val)))
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
		return fmt.Sprintf("  %s = %s", key, formatValue(v))
	case []interface{}:
		if len(v) > 0 {
			if _, isSimple := v[0].(string); isSimple {
				return fmt.Sprintf("  %s = %s", key, formatValue(v))
			}
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
		return fmt.Sprintf("  %s = []", key)
	default:
		return fmt.Sprintf("  %s = %s", key, formatValue(v))
	}
}

// handleArchiveRetentions applies special logic for archive_retentions resource
// The first retention is the default retention and can't have a name set
func handleArchiveRetentions(resourceMap map[string]interface{}) map[string]interface{} {
	nameResult, err := jsonpath.Get("$.retentions[0].name", resourceMap)
	if err != nil || nameResult == nil {
		return resourceMap // No first retention or no name field
	}

	result := make(map[string]interface{})
	for k, v := range resourceMap {
		result[k] = v
	}

	retentions, err := jsonpath.Get("$.retentions", resourceMap)
	if err != nil {
		return resourceMap
	}

	if retentionSlice, ok := retentions.([]interface{}); ok && len(retentionSlice) > 0 {
		newRetentions := make([]interface{}, len(retentionSlice))
		copy(newRetentions, retentionSlice)

		if firstRetention, ok := newRetentions[0].(map[string]interface{}); ok {
			newFirstRetention := make(map[string]interface{})
			for key, value := range firstRetention {
				if key != "name" {
					newFirstRetention[key] = value
				}
			}
			newRetentions[0] = newFirstRetention
		}

		result["retentions"] = newRetentions
	}

	return result
}

// handleDataTableFilters ensures data_table filters have required metric attribute
func handleDataTableFilters(resourceMap map[string]interface{}) map[string]interface{} {
	filtersResult, err := jsonpath.Get("$..data_table.query.metrics.filters", resourceMap)
	if err != nil {
		return resourceMap // No data_table filters found
	}

	jsonBytes, err := json.Marshal(resourceMap)
	if err != nil {
		return resourceMap
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return resourceMap
	}

	switch filters := filtersResult.(type) {
	case []interface{}:
		processFilterArray(filters)
	case [][]interface{}:
		for _, filterArray := range filters {
			processFilterArray(filterArray)
		}
	}

	return result
}

// processFilterArray processes a single filter array to ensure metric attributes exist
func processFilterArray(filters []interface{}) {
	for _, filter := range filters {
		if filterMap, ok := filter.(map[string]interface{}); ok {
			if _, hasMetric := filterMap["metric"]; !hasMetric {
				if label, hasLabel := filterMap["label"]; hasLabel {
					if labelStr, ok := label.(string); ok {
						filterMap["metric"] = labelStr
					} else {
						filterMap["metric"] = "vector"
					}
				} else {
					filterMap["metric"] = "vector"
				}
			}
		}
	}
}

// Alert group_by rules are validated against Schema V2:
//   internal/provider/alerts/alert_schema/v2.go (group_by attr + type_definition)
//   internal/provider/alerts/alert_schema/common.go (GroupByValidator, logsRatioGroupByForSchema)
//
// handleAlertGroupBy fixes group_by related validation issues for alert resources.
//
// Case 1: GroupByValidator (common.go ValidateList) - group_by is not allowed for
//   logs_immediate, logs_new_value, tracing_immediate. We remove group_by when type_definition has any of these.
// Case 2: logsRatioGroupByForSchema (common.go) - group_by_for has AlsoRequires(group_by).
//   We remove group_by_for when group_by is missing to satisfy the provider.
func handleAlertGroupBy(resourceMap map[string]interface{}) map[string]interface{} {
	typeDefinition, ok := resourceMap["type_definition"].(map[string]interface{})
	if !ok {
		return resourceMap
	}

	unsupportedTypes := []string{"logs_immediate", "logs_new_value", "tracing_immediate"}
	for _, alertType := range unsupportedTypes {
		if _, exists := typeDefinition[alertType]; exists {
			delete(resourceMap, "group_by")
			return resourceMap
		}
	}

	if logsRatio, ok := typeDefinition["logs_ratio_threshold"].(map[string]interface{}); ok {
		if _, hasGroupByFor := logsRatio["group_by_for"]; hasGroupByFor {
			if _, hasGroupBy := resourceMap["group_by"]; !hasGroupBy {
				delete(logsRatio, "group_by_for")
			}
		}
	}

	return resourceMap
}

// ensureDescription adds an empty description when missing.
// This works around a Coralogix provider bug where null description vs empty string
// causes "Provider produced inconsistent result after apply" errors.
// Used for coralogix_alert and coralogix_events2metric.
func ensureDescription(resourceMap map[string]interface{}) map[string]interface{} {
	if _, hasDescription := resourceMap["description"]; !hasDescription {
		resourceMap["description"] = ""
	}
	return resourceMap
}

// defaultHistogramBuckets is used when the provider requires buckets but the API/state
// returns histogram with enable=false and no buckets. The coralogix provider v2.2.3
// marks histogram.buckets as required whenever a histogram block exists.
var defaultHistogramBuckets = []interface{}{0.0}

// handleEvents2MetricHistogramBuckets adds default buckets to histogram blocks that
// are missing them. The Coralogix provider requires buckets to be set whenever a
// histogram aggregation block exists, even when enable=false.
func handleEvents2MetricHistogramBuckets(resourceMap map[string]interface{}) map[string]interface{} {
	metricFields, ok := resourceMap["metric_fields"].(map[string]interface{})
	if !ok {
		return resourceMap
	}

	for _, fieldVal := range metricFields {
		fieldMap, ok := fieldVal.(map[string]interface{})
		if !ok {
			continue
		}
		aggregations, ok := fieldMap["aggregations"].(map[string]interface{})
		if !ok {
			continue
		}
		histogram, ok := aggregations["histogram"].(map[string]interface{})
		if !ok {
			continue
		}
		if _, hasBuckets := histogram["buckets"]; !hasBuckets {
			histogram["buckets"] = defaultHistogramBuckets
		}
	}

	return resourceMap
}

// formatBlockBody outputs HCL attribute lines for a map (used inside blocks).
func formatBlockBody(m map[string]interface{}) string {
	var lines []string
	for k, v := range m {
		lines = append(lines, processBlock(k, v))
	}
	return strings.Join(lines, "\n")
}

// formatRuleTypeBlocks emits repeated HCL blocks for a rule type (e.g. "block", "replace").
// listVal is a []interface{} of config maps; each becomes "blockName { ... }".
func formatRuleTypeBlocks(blockName string, listVal []interface{}) string {
	var blocks []string
	for _, item := range listVal {
		elem, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		body := formatBlockBody(elem)
		blocks = append(blocks, "  "+blockName+" {\n"+body+"\n  }")
	}
	return strings.Join(blocks, "\n")
}

// formatRulesBlocks emits "rules { ... }" blocks. Each element of rulesList is a map
// with one key (the rule type, e.g. "block", "parse_json_field") and value []interface{}.
func formatRulesBlocks(rulesList []interface{}) string {
	var blocks []string
	for _, ruleItem := range rulesList {
		ruleMap, ok := ruleItem.(map[string]interface{})
		if !ok {
			continue
		}
		for blockName, blockVal := range ruleMap {
			listVal, ok := blockVal.([]interface{})
			if !ok || len(listVal) == 0 {
				continue
			}
			inner := formatRuleTypeBlocks(blockName, listVal)
			blocks = append(blocks, "  rules {\n"+inner+"\n  }")
			break // one rule type per rule map
		}
	}
	return strings.Join(blocks, "\n")
}

// formatRuleSubgroupsBlocks emits "rule_subgroups { ... }" blocks for coralogix_rules_group.
// Each subgroup is a map with attributes (order, active, etc.) and "rules" (list).
func formatRuleSubgroupsBlocks(ruleSubgroups interface{}) string {
	list, ok := ruleSubgroups.([]interface{})
	if !ok || len(list) == 0 {
		return processBlock("rule_subgroups", ruleSubgroups)
	}
	var blocks []string
	for _, subgroupItem := range list {
		subgroup, ok := subgroupItem.(map[string]interface{})
		if !ok {
			continue
		}
		var lines []string
		for k, v := range subgroup {
			if k == "rules" {
				rulesList, ok := v.([]interface{})
				if ok {
					lines = append(lines, formatRulesBlocks(rulesList))
				} else {
					lines = append(lines, processBlock(k, v))
				}
			} else {
				lines = append(lines, processBlock(k, v))
			}
		}
		blocks = append(blocks, "  rule_subgroups {\n"+strings.Join(lines, "\n")+"\n  }")
	}
	return strings.Join(blocks, "\n")
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
								if resourceType == "coralogix_archive_retentions" {
									resourceMap = handleArchiveRetentions(resourceMap)
								}
								if resourceType == "coralogix_dashboard" {
									resourceMap = handleDataTableFilters(resourceMap)
								}
								if resourceType == "coralogix_alert" {
									resourceMap = handleAlertGroupBy(resourceMap)
									resourceMap = ensureDescription(resourceMap)
								}
								if resourceType == "coralogix_events2metric" {
									resourceMap = ensureDescription(resourceMap)
									resourceMap = handleEvents2MetricHistogramBuckets(resourceMap)
								}

								terraformLines = append(terraformLines, fmt.Sprintf(`resource "%s" "%s" {`, resourceType, resourceName))
								for key, val := range resourceMap {
									if resourceType == "coralogix_rules_group" && key == "rule_subgroups" {
										terraformLines = append(terraformLines, formatRuleSubgroupsBlocks(val))
									} else {
										terraformLines = append(terraformLines, processBlock(key, val))
									}
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
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	terraformOutput := generateTerraform(parsed)

	err = os.WriteFile(tfFilePath, []byte(terraformOutput), 0644)
	if err != nil {
		log.Fatalf("Failed to write Terraform file: %v", err)
	}

	fmt.Printf("Terraform configuration written to %s\n", tfFilePath)
}
