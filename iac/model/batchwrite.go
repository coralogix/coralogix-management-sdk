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
	"errors"
	"fmt"
	"strings"
)

// batchWriteLocation is the line of the behavior-overrides file that says the writes take arrays.
const batchWriteLocation = "behavior-overrides.yaml:api.batchWrite"

// readBatchBodies reads the Create and Update bodies of a BatchWrite resource. Each is an array of
// one request component that accepts one item, and the Operation names the component. The rest of
// the model reads the item where it reads a body (Policy.bodyOf).
func (r *Resource) readBatchBodies(ops map[verb]foundOp) error {
	if !r.Policy.BatchWrite {
		return nil
	}
	if !r.Policy.Existing {
		return errors.New("api.batchWrite is only for existing resources")
	}
	if arrayItems(bodyProxy(ops[opCreate].op)) == nil && arrayItems(bodyProxy(ops[opUpdate].op)) == nil {
		return &contractError{code: "OVERRIDE_UNUSED", location: batchWriteLocation,
			message:     "Neither the Create nor the Update body is an array.",
			remediation: "Delete api.batchWrite."}
	}
	writes := map[verb]*Operation{opCreate: &r.Create, opUpdate: &r.Update}
	for _, v := range []verb{opCreate, opUpdate} {
		o := writes[v]
		location := "paths." + strings.ToLower(string(v)) + "." + o.OperationID + ".requestBody"
		notArray := func(message string) error {
			return &contractError{code: "BATCH_WRITE_BODY_NOT_ARRAY", location: location, message: message,
				remediation: "Describe a resource whose Create and Update both take an array of one request component, or delete api.batchWrite."}
		}
		body := bodyProxy(ops[v].op)
		items := arrayItems(body)
		if items == nil {
			return notArray(fmt.Sprintf("The %s body is not an array. api.batchWrite needs both writes to take an array.", v))
		}
		component, err := componentName(unwrapRef(items))
		if err != nil {
			return notArray(fmt.Sprintf("The items of the %s body are not a component.", v))
		}
		if s := mustSchema(body); s.MinItems != nil && *s.MinItems > 1 || s.MaxItems != nil && *s.MaxItems < 1 {
			return &contractError{code: "BATCH_WRITE_ONE_ITEM_REFUSED", location: location,
				message:     fmt.Sprintf("The %s body does not accept an array of one item. A tool writes one resource at a time.", v),
				remediation: "Accept one item in the array body."}
		}
		o.Body, o.BodyItems = component, true
	}
	return nil
}
