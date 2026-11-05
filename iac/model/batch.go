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

import (
	"fmt"
	"slices"
)

// readBatchOperations keeps the existing batch Get on the collection path. Delete still
// identifies one resource, and Update follows the normal body-id or path-id contract.
func (r *Resource) readBatchOperations(ops map[verb]foundOp) error {
	if !r.Policy.Existing {
		return fmt.Errorf("api.getByIDs is only for existing resources")
	}
	get := ops[opGet]
	if len(pathParams(get)) != 0 {
		return fmt.Errorf("api.getByIDs needs a Get without path parameters")
	}
	var found bool
	for _, p := range slices.Concat(get.item.Parameters, get.op.Parameters) {
		if p.In != "query" || p.Name != r.Policy.GetByIDs {
			continue
		}
		if found {
			return fmt.Errorf("api.getByIDs query parameter is duplicated")
		}
		found = true
		s, err := schemaOf(p.Schema)
		if err != nil || !slices.Equal(s.Type, []string{"array"}) || s.Items == nil || s.Items.A == nil {
			return fmt.Errorf("api.getByIDs needs an array query parameter")
		}
		item, err := schemaOf(s.Items.A)
		if err != nil || !slices.Equal(item.Type, []string{"string"}) {
			return fmt.Errorf("api.getByIDs needs string ids")
		}
	}
	if !found {
		return fmt.Errorf("api.getByIDs: no query parameter %q", r.Policy.GetByIDs)
	}
	id, err := idParam(ops[opDelete])
	if err != nil {
		return fmt.Errorf("api.getByIDs delete: %w", err)
	}
	r.IDParam = id
	if len(pathParams(ops[opCreate])) != 0 {
		return fmt.Errorf("create must have no path parameters")
	}
	item, err := r.itemOperations(ops)
	if err != nil {
		return err
	}
	for _, role := range item {
		other, err := idParam(ops[role])
		if err != nil || other != id {
			return fmt.Errorf("%s must use the same id path parameter as Delete", role)
		}
	}
	if !r.Policy.CustomMethods && ops[opDelete].path != get.path+"/{"+id+"}" {
		return fmt.Errorf("api.getByIDs Delete must use Get path/{%s}", id)
	}
	if get.path != ops[opCreate].path && !r.Policy.CustomMethods {
		return fmt.Errorf("api.getByIDs Get must use the Create path")
	}
	if err := r.readTargets(ops); err != nil {
		return err
	}
	if !r.Get.Response.Items {
		return fmt.Errorf("api.getByIDs needs a response array of %s", r.Name)
	}
	return nil
}
