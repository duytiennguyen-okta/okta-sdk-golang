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

// PolicyPlatformType the model 'PolicyPlatformType'
type PolicyPlatformType string

// List of PolicyPlatformType
const (
	POLICYPLATFORMTYPE_ANY     PolicyPlatformType = "ANY"
	POLICYPLATFORMTYPE_DESKTOP PolicyPlatformType = "DESKTOP"
	POLICYPLATFORMTYPE_MOBILE  PolicyPlatformType = "MOBILE"
	POLICYPLATFORMTYPE_OTHER   PolicyPlatformType = "OTHER"
)

// All allowed values of PolicyPlatformType enum
var AllowedPolicyPlatformTypeEnumValues = []PolicyPlatformType{
	"ANY",
	"DESKTOP",
	"MOBILE",
	"OTHER",
}

func (v *PolicyPlatformType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyPlatformType(value)
	for _, existing := range AllowedPolicyPlatformTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyPlatformType", value)
}

// NewPolicyPlatformTypeFromValue returns a pointer to a valid PolicyPlatformType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyPlatformTypeFromValue(v string) (*PolicyPlatformType, error) {
	ev := PolicyPlatformType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyPlatformType: valid values are %v", v, AllowedPolicyPlatformTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyPlatformType) IsValid() bool {
	for _, existing := range AllowedPolicyPlatformTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyPlatformType value
func (v PolicyPlatformType) Ptr() *PolicyPlatformType {
	return &v
}

type NullablePolicyPlatformType struct {
	value *PolicyPlatformType
	isSet bool
}

func (v NullablePolicyPlatformType) Get() *PolicyPlatformType {
	return v.value
}

func (v *NullablePolicyPlatformType) Set(val *PolicyPlatformType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyPlatformType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyPlatformType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyPlatformType(val *PolicyPlatformType) *NullablePolicyPlatformType {
	return &NullablePolicyPlatformType{value: val, isSet: true}
}

func (v NullablePolicyPlatformType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyPlatformType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
