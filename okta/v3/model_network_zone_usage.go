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

// NetworkZoneUsage the model 'NetworkZoneUsage'
type NetworkZoneUsage string

// List of NetworkZoneUsage
const (
	NETWORKZONEUSAGE_BLOCKLIST NetworkZoneUsage = "BLOCKLIST"
	NETWORKZONEUSAGE_POLICY    NetworkZoneUsage = "POLICY"
)

// All allowed values of NetworkZoneUsage enum
var AllowedNetworkZoneUsageEnumValues = []NetworkZoneUsage{
	"BLOCKLIST",
	"POLICY",
}

func (v *NetworkZoneUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkZoneUsage(value)
	for _, existing := range AllowedNetworkZoneUsageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkZoneUsage", value)
}

// NewNetworkZoneUsageFromValue returns a pointer to a valid NetworkZoneUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkZoneUsageFromValue(v string) (*NetworkZoneUsage, error) {
	ev := NetworkZoneUsage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkZoneUsage: valid values are %v", v, AllowedNetworkZoneUsageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkZoneUsage) IsValid() bool {
	for _, existing := range AllowedNetworkZoneUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkZoneUsage value
func (v NetworkZoneUsage) Ptr() *NetworkZoneUsage {
	return &v
}

type NullableNetworkZoneUsage struct {
	value *NetworkZoneUsage
	isSet bool
}

func (v NullableNetworkZoneUsage) Get() *NetworkZoneUsage {
	return v.value
}

func (v *NullableNetworkZoneUsage) Set(val *NetworkZoneUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkZoneUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkZoneUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkZoneUsage(val *NetworkZoneUsage) *NullableNetworkZoneUsage {
	return &NullableNetworkZoneUsage{value: val, isSet: true}
}

func (v NullableNetworkZoneUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkZoneUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
