/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// FeatureStageState the model 'FeatureStageState'
type FeatureStageState string

// List of FeatureStageState
const (
	FEATURESTAGESTATE_CLOSED FeatureStageState = "CLOSED"
	FEATURESTAGESTATE_OPEN   FeatureStageState = "OPEN"
)

// All allowed values of FeatureStageState enum
var AllowedFeatureStageStateEnumValues = []FeatureStageState{
	"CLOSED",
	"OPEN",
}

func (v *FeatureStageState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeatureStageState(value)
	for _, existing := range AllowedFeatureStageStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeatureStageState", value)
}

// NewFeatureStageStateFromValue returns a pointer to a valid FeatureStageState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeatureStageStateFromValue(v string) (*FeatureStageState, error) {
	ev := FeatureStageState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeatureStageState: valid values are %v", v, AllowedFeatureStageStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeatureStageState) IsValid() bool {
	for _, existing := range AllowedFeatureStageStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureStageState value
func (v FeatureStageState) Ptr() *FeatureStageState {
	return &v
}

type NullableFeatureStageState struct {
	value *FeatureStageState
	isSet bool
}

func (v NullableFeatureStageState) Get() *FeatureStageState {
	return v.value
}

func (v *NullableFeatureStageState) Set(val *FeatureStageState) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureStageState) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureStageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureStageState(val *FeatureStageState) *NullableFeatureStageState {
	return &NullableFeatureStageState{value: val, isSet: true}
}

func (v NullableFeatureStageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureStageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
