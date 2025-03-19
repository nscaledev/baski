/*
Copyright 2025 Nscale.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package data

import (
	"errors"
	"testing"
)

// TestGetNestedField enables the ability to step through a map[string]interface{} to get sub keys
func TestGetNestedField(t *testing.T) {
	testMap := map[string]interface{}{
		"this": map[string]interface{}{
			"is": map[string]interface{}{
				"a": map[string]interface{}{
					"test": map[string]interface{}{
						"item": "result1",
					},
				},
				"another": map[string]interface{}{
					"test": map[string]string{
						"item": "result2",
					},
				},
			},
		},
	}
	tests := []struct {
		Name     string
		Keys     []string
		Expected string
		Error    error
	}{
		{
			Name:     "Test map[string]interface",
			Keys:     []string{"this", "is", "a", "test", "item"},
			Expected: "result1",
			Error:    nil,
		},
		{
			Name:     "Test map[string]string",
			Keys:     []string{"this", "is", "another", "test", "item"},
			Expected: "",
			Error:    errors.New("key item not found or not a map\n"),
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := GetNestedField(testMap, test.Keys...)

			if err != nil {
				if test.Error != nil {
					if test.Error.Error() != err.Error() {
						t.Errorf("expected error \"%s\" got \"%s\"\n", test.Error.Error(), err.Error())
					}
					return
				} else {
					t.Error(err)
					return
				}
			}

			if res != test.Expected {
				t.Errorf("expected %s got %s\n", test.Expected, res)
			}
		})
	}

}
