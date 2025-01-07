// Copyright 2025 cm
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNilAndPrintError(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}
	assert.Nil(t, err)
}
