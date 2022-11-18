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

// OktaSignOnPolicyFactorPromptMode the model 'OktaSignOnPolicyFactorPromptMode'
type OktaSignOnPolicyFactorPromptMode string

// List of OktaSignOnPolicyFactorPromptMode
const (
	OKTASIGNONPOLICYFACTORPROMPTMODE_ALWAYS  OktaSignOnPolicyFactorPromptMode = "ALWAYS"
	OKTASIGNONPOLICYFACTORPROMPTMODE_DEVICE  OktaSignOnPolicyFactorPromptMode = "DEVICE"
	OKTASIGNONPOLICYFACTORPROMPTMODE_SESSION OktaSignOnPolicyFactorPromptMode = "SESSION"
)

// All allowed values of OktaSignOnPolicyFactorPromptMode enum
var AllowedOktaSignOnPolicyFactorPromptModeEnumValues = []OktaSignOnPolicyFactorPromptMode{
	"ALWAYS",
	"DEVICE",
	"SESSION",
}

func (v *OktaSignOnPolicyFactorPromptMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OktaSignOnPolicyFactorPromptMode(value)
	for _, existing := range AllowedOktaSignOnPolicyFactorPromptModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OktaSignOnPolicyFactorPromptMode", value)
}

// NewOktaSignOnPolicyFactorPromptModeFromValue returns a pointer to a valid OktaSignOnPolicyFactorPromptMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOktaSignOnPolicyFactorPromptModeFromValue(v string) (*OktaSignOnPolicyFactorPromptMode, error) {
	ev := OktaSignOnPolicyFactorPromptMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OktaSignOnPolicyFactorPromptMode: valid values are %v", v, AllowedOktaSignOnPolicyFactorPromptModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OktaSignOnPolicyFactorPromptMode) IsValid() bool {
	for _, existing := range AllowedOktaSignOnPolicyFactorPromptModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OktaSignOnPolicyFactorPromptMode value
func (v OktaSignOnPolicyFactorPromptMode) Ptr() *OktaSignOnPolicyFactorPromptMode {
	return &v
}

type NullableOktaSignOnPolicyFactorPromptMode struct {
	value *OktaSignOnPolicyFactorPromptMode
	isSet bool
}

func (v NullableOktaSignOnPolicyFactorPromptMode) Get() *OktaSignOnPolicyFactorPromptMode {
	return v.value
}

func (v *NullableOktaSignOnPolicyFactorPromptMode) Set(val *OktaSignOnPolicyFactorPromptMode) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyFactorPromptMode) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyFactorPromptMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyFactorPromptMode(val *OktaSignOnPolicyFactorPromptMode) *NullableOktaSignOnPolicyFactorPromptMode {
	return &NullableOktaSignOnPolicyFactorPromptMode{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyFactorPromptMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyFactorPromptMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
