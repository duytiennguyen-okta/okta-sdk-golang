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
)

// MultifactorEnrollmentPolicyAuthenticatorSettings struct for MultifactorEnrollmentPolicyAuthenticatorSettings
type MultifactorEnrollmentPolicyAuthenticatorSettings struct {
	Enroll               *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll `json:"enroll,omitempty"`
	Key                  *MultifactorEnrollmentPolicyAuthenticatorType           `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicyAuthenticatorSettings MultifactorEnrollmentPolicyAuthenticatorSettings

// NewMultifactorEnrollmentPolicyAuthenticatorSettings instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicyAuthenticatorSettings() *MultifactorEnrollmentPolicyAuthenticatorSettings {
	this := MultifactorEnrollmentPolicyAuthenticatorSettings{}
	return &this
}

// NewMultifactorEnrollmentPolicyAuthenticatorSettingsWithDefaults instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicyAuthenticatorSettingsWithDefaults() *MultifactorEnrollmentPolicyAuthenticatorSettings {
	this := MultifactorEnrollmentPolicyAuthenticatorSettings{}
	return &this
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetEnroll() MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll {
	if o == nil || o.Enroll == nil {
		var ret MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetEnrollOk() (*MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll and assigns it to the Enroll field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) SetEnroll(v MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) {
	o.Enroll = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetKey() MultifactorEnrollmentPolicyAuthenticatorType {
	if o == nil || o.Key == nil {
		var ret MultifactorEnrollmentPolicyAuthenticatorType
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetKeyOk() (*MultifactorEnrollmentPolicyAuthenticatorType, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given MultifactorEnrollmentPolicyAuthenticatorType and assigns it to the Key field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) SetKey(v MultifactorEnrollmentPolicyAuthenticatorType) {
	o.Key = &v
}

func (o MultifactorEnrollmentPolicyAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) UnmarshalJSON(bytes []byte) (err error) {
	varMultifactorEnrollmentPolicyAuthenticatorSettings := _MultifactorEnrollmentPolicyAuthenticatorSettings{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicyAuthenticatorSettings)
	if err == nil {
		*o = MultifactorEnrollmentPolicyAuthenticatorSettings(varMultifactorEnrollmentPolicyAuthenticatorSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMultifactorEnrollmentPolicyAuthenticatorSettings struct {
	value *MultifactorEnrollmentPolicyAuthenticatorSettings
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettings) Get() *MultifactorEnrollmentPolicyAuthenticatorSettings {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettings) Set(val *MultifactorEnrollmentPolicyAuthenticatorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicyAuthenticatorSettings(val *MultifactorEnrollmentPolicyAuthenticatorSettings) *NullableMultifactorEnrollmentPolicyAuthenticatorSettings {
	return &NullableMultifactorEnrollmentPolicyAuthenticatorSettings{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
