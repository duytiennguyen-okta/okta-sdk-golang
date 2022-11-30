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

// PasswordPolicyAuthenticationProviderType the model 'PasswordPolicyAuthenticationProviderType'
type PasswordPolicyAuthenticationProviderType string

// List of PasswordPolicyAuthenticationProviderType
const (
	PASSWORDPOLICYAUTHENTICATIONPROVIDERTYPE_ACTIVE_DIRECTORY PasswordPolicyAuthenticationProviderType = "ACTIVE_DIRECTORY"
	PASSWORDPOLICYAUTHENTICATIONPROVIDERTYPE_ANY              PasswordPolicyAuthenticationProviderType = "ANY"
	PASSWORDPOLICYAUTHENTICATIONPROVIDERTYPE_LDAP             PasswordPolicyAuthenticationProviderType = "LDAP"
	PASSWORDPOLICYAUTHENTICATIONPROVIDERTYPE_OKTA             PasswordPolicyAuthenticationProviderType = "OKTA"
)

// All allowed values of PasswordPolicyAuthenticationProviderType enum
var AllowedPasswordPolicyAuthenticationProviderTypeEnumValues = []PasswordPolicyAuthenticationProviderType{
	"ACTIVE_DIRECTORY",
	"ANY",
	"LDAP",
	"OKTA",
}

func (v *PasswordPolicyAuthenticationProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PasswordPolicyAuthenticationProviderType(value)
	for _, existing := range AllowedPasswordPolicyAuthenticationProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PasswordPolicyAuthenticationProviderType", value)
}

// NewPasswordPolicyAuthenticationProviderTypeFromValue returns a pointer to a valid PasswordPolicyAuthenticationProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPasswordPolicyAuthenticationProviderTypeFromValue(v string) (*PasswordPolicyAuthenticationProviderType, error) {
	ev := PasswordPolicyAuthenticationProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PasswordPolicyAuthenticationProviderType: valid values are %v", v, AllowedPasswordPolicyAuthenticationProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PasswordPolicyAuthenticationProviderType) IsValid() bool {
	for _, existing := range AllowedPasswordPolicyAuthenticationProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PasswordPolicyAuthenticationProviderType value
func (v PasswordPolicyAuthenticationProviderType) Ptr() *PasswordPolicyAuthenticationProviderType {
	return &v
}

type NullablePasswordPolicyAuthenticationProviderType struct {
	value *PasswordPolicyAuthenticationProviderType
	isSet bool
}

func (v NullablePasswordPolicyAuthenticationProviderType) Get() *PasswordPolicyAuthenticationProviderType {
	return v.value
}

func (v *NullablePasswordPolicyAuthenticationProviderType) Set(val *PasswordPolicyAuthenticationProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyAuthenticationProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyAuthenticationProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyAuthenticationProviderType(val *PasswordPolicyAuthenticationProviderType) *NullablePasswordPolicyAuthenticationProviderType {
	return &NullablePasswordPolicyAuthenticationProviderType{value: val, isSet: true}
}

func (v NullablePasswordPolicyAuthenticationProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyAuthenticationProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
