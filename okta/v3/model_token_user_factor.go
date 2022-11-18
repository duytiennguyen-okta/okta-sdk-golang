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
	"reflect"
	"strings"
)

// TokenUserFactor struct for TokenUserFactor
type TokenUserFactor struct {
	UserFactor
	Profile              *TokenUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenUserFactor TokenUserFactor

// NewTokenUserFactor instantiates a new TokenUserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenUserFactor() *TokenUserFactor {
	this := TokenUserFactor{}
	return &this
}

// NewTokenUserFactorWithDefaults instantiates a new TokenUserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenUserFactorWithDefaults() *TokenUserFactor {
	this := TokenUserFactor{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *TokenUserFactor) GetProfile() TokenUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret TokenUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUserFactor) GetProfileOk() (*TokenUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *TokenUserFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given TokenUserFactorProfile and assigns it to the Profile field.
func (o *TokenUserFactor) SetProfile(v TokenUserFactorProfile) {
	o.Profile = &v
}

func (o TokenUserFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenUserFactor) UnmarshalJSON(bytes []byte) (err error) {
	type TokenUserFactorWithoutEmbeddedStruct struct {
		Profile *TokenUserFactorProfile `json:"profile,omitempty"`
	}

	varTokenUserFactorWithoutEmbeddedStruct := TokenUserFactorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTokenUserFactorWithoutEmbeddedStruct)
	if err == nil {
		varTokenUserFactor := _TokenUserFactor{}
		varTokenUserFactor.Profile = varTokenUserFactorWithoutEmbeddedStruct.Profile
		*o = TokenUserFactor(varTokenUserFactor)
	} else {
		return err
	}

	varTokenUserFactor := _TokenUserFactor{}

	err = json.Unmarshal(bytes, &varTokenUserFactor)
	if err == nil {
		o.UserFactor = varTokenUserFactor.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenUserFactor struct {
	value *TokenUserFactor
	isSet bool
}

func (v NullableTokenUserFactor) Get() *TokenUserFactor {
	return v.value
}

func (v *NullableTokenUserFactor) Set(val *TokenUserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenUserFactor(val *TokenUserFactor) *NullableTokenUserFactor {
	return &NullableTokenUserFactor{value: val, isSet: true}
}

func (v NullableTokenUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
