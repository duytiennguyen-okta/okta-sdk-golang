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

// PolicySubjectMatchType the model 'PolicySubjectMatchType'
type PolicySubjectMatchType string

// List of PolicySubjectMatchType
const (
	POLICYSUBJECTMATCHTYPE_CUSTOM_ATTRIBUTE  PolicySubjectMatchType = "CUSTOM_ATTRIBUTE"
	POLICYSUBJECTMATCHTYPE_EMAIL             PolicySubjectMatchType = "EMAIL"
	POLICYSUBJECTMATCHTYPE_USERNAME          PolicySubjectMatchType = "USERNAME"
	POLICYSUBJECTMATCHTYPE_USERNAME_OR_EMAIL PolicySubjectMatchType = "USERNAME_OR_EMAIL"
)

// All allowed values of PolicySubjectMatchType enum
var AllowedPolicySubjectMatchTypeEnumValues = []PolicySubjectMatchType{
	"CUSTOM_ATTRIBUTE",
	"EMAIL",
	"USERNAME",
	"USERNAME_OR_EMAIL",
}

func (v *PolicySubjectMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicySubjectMatchType(value)
	for _, existing := range AllowedPolicySubjectMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicySubjectMatchType", value)
}

// NewPolicySubjectMatchTypeFromValue returns a pointer to a valid PolicySubjectMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicySubjectMatchTypeFromValue(v string) (*PolicySubjectMatchType, error) {
	ev := PolicySubjectMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicySubjectMatchType: valid values are %v", v, AllowedPolicySubjectMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicySubjectMatchType) IsValid() bool {
	for _, existing := range AllowedPolicySubjectMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicySubjectMatchType value
func (v PolicySubjectMatchType) Ptr() *PolicySubjectMatchType {
	return &v
}

type NullablePolicySubjectMatchType struct {
	value *PolicySubjectMatchType
	isSet bool
}

func (v NullablePolicySubjectMatchType) Get() *PolicySubjectMatchType {
	return v.value
}

func (v *NullablePolicySubjectMatchType) Set(val *PolicySubjectMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySubjectMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySubjectMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySubjectMatchType(val *PolicySubjectMatchType) *NullablePolicySubjectMatchType {
	return &NullablePolicySubjectMatchType{value: val, isSet: true}
}

func (v NullablePolicySubjectMatchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySubjectMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
