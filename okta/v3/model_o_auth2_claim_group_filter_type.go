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

// OAuth2ClaimGroupFilterType the model 'OAuth2ClaimGroupFilterType'
type OAuth2ClaimGroupFilterType string

// List of OAuth2ClaimGroupFilterType
const (
	OAUTH2CLAIMGROUPFILTERTYPE_CONTAINS    OAuth2ClaimGroupFilterType = "CONTAINS"
	OAUTH2CLAIMGROUPFILTERTYPE_EQUALS      OAuth2ClaimGroupFilterType = "EQUALS"
	OAUTH2CLAIMGROUPFILTERTYPE_REGEX       OAuth2ClaimGroupFilterType = "REGEX"
	OAUTH2CLAIMGROUPFILTERTYPE_STARTS_WITH OAuth2ClaimGroupFilterType = "STARTS_WITH"
)

// All allowed values of OAuth2ClaimGroupFilterType enum
var AllowedOAuth2ClaimGroupFilterTypeEnumValues = []OAuth2ClaimGroupFilterType{
	"CONTAINS",
	"EQUALS",
	"REGEX",
	"STARTS_WITH",
}

func (v *OAuth2ClaimGroupFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OAuth2ClaimGroupFilterType(value)
	for _, existing := range AllowedOAuth2ClaimGroupFilterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OAuth2ClaimGroupFilterType", value)
}

// NewOAuth2ClaimGroupFilterTypeFromValue returns a pointer to a valid OAuth2ClaimGroupFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOAuth2ClaimGroupFilterTypeFromValue(v string) (*OAuth2ClaimGroupFilterType, error) {
	ev := OAuth2ClaimGroupFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OAuth2ClaimGroupFilterType: valid values are %v", v, AllowedOAuth2ClaimGroupFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OAuth2ClaimGroupFilterType) IsValid() bool {
	for _, existing := range AllowedOAuth2ClaimGroupFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuth2ClaimGroupFilterType value
func (v OAuth2ClaimGroupFilterType) Ptr() *OAuth2ClaimGroupFilterType {
	return &v
}

type NullableOAuth2ClaimGroupFilterType struct {
	value *OAuth2ClaimGroupFilterType
	isSet bool
}

func (v NullableOAuth2ClaimGroupFilterType) Get() *OAuth2ClaimGroupFilterType {
	return v.value
}

func (v *NullableOAuth2ClaimGroupFilterType) Set(val *OAuth2ClaimGroupFilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClaimGroupFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClaimGroupFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClaimGroupFilterType(val *OAuth2ClaimGroupFilterType) *NullableOAuth2ClaimGroupFilterType {
	return &NullableOAuth2ClaimGroupFilterType{value: val, isSet: true}
}

func (v NullableOAuth2ClaimGroupFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClaimGroupFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
