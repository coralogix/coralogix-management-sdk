// Copyright 2026 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "testing"

// connectorListReadPolicy reads the Connector of testdata/lists.yaml from its List: no operation of
// the contract is used to read one connector, and Update and Delete take the id in the item path.
func connectorListReadPolicy() Policy {
	return Policy{
		Existing: true, GetFromList: true, IDProperty: "id",
		Operations: OperationIDs{Get: "ConnectorsService_ListConnectors"},
	}
}

// A GetFromList resource whose Update takes the id in the item path, the path of Delete.
func TestGetFromListUpdateIDInPath(t *testing.T) {
	doc := listsDoc(t)
	policy := connectorListReadPolicy()
	if report := ValidateWithPolicy(doc, "Connector", OperationIDs{}, policy); len(report) != 0 {
		t.Fatalf("issues:\n%s", report.Error())
	}
	r, err := BuildWithPolicy(doc, "Connector", OperationIDs{}, policy)
	if err != nil {
		t.Fatal(err)
	}
	if !r.GetFromList || r.List == nil || r.List.OperationID != "ConnectorsService_ListConnectors" {
		t.Fatalf("Get is not the List: %+v", r)
	}
	if r.IDParam != "id" || r.IDField != "id" {
		t.Fatalf("id parameter %q, id field %q", r.IDParam, r.IDField)
	}
	if r.Update.Path != "/connectors/{id}" || r.Delete.Path != "/connectors/{id}" || r.Create.Path != "/connectors" {
		t.Fatalf("paths: create %s, update %s, delete %s", r.Create.Path, r.Update.Path, r.Delete.Path)
	}
}
