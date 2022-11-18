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

// WsFederationApplicationSettingsAllOf struct for WsFederationApplicationSettingsAllOf
type WsFederationApplicationSettingsAllOf struct {
	App                  *WsFederationApplicationSettingsApplication `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WsFederationApplicationSettingsAllOf WsFederationApplicationSettingsAllOf

// NewWsFederationApplicationSettingsAllOf instantiates a new WsFederationApplicationSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsFederationApplicationSettingsAllOf() *WsFederationApplicationSettingsAllOf {
	this := WsFederationApplicationSettingsAllOf{}
	return &this
}

// NewWsFederationApplicationSettingsAllOfWithDefaults instantiates a new WsFederationApplicationSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsFederationApplicationSettingsAllOfWithDefaults() *WsFederationApplicationSettingsAllOf {
	this := WsFederationApplicationSettingsAllOf{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *WsFederationApplicationSettingsAllOf) GetApp() WsFederationApplicationSettingsApplication {
	if o == nil || o.App == nil {
		var ret WsFederationApplicationSettingsApplication
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettingsAllOf) GetAppOk() (*WsFederationApplicationSettingsApplication, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *WsFederationApplicationSettingsAllOf) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given WsFederationApplicationSettingsApplication and assigns it to the App field.
func (o *WsFederationApplicationSettingsAllOf) SetApp(v WsFederationApplicationSettingsApplication) {
	o.App = &v
}

func (o WsFederationApplicationSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WsFederationApplicationSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWsFederationApplicationSettingsAllOf := _WsFederationApplicationSettingsAllOf{}

	err = json.Unmarshal(bytes, &varWsFederationApplicationSettingsAllOf)
	if err == nil {
		*o = WsFederationApplicationSettingsAllOf(varWsFederationApplicationSettingsAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWsFederationApplicationSettingsAllOf struct {
	value *WsFederationApplicationSettingsAllOf
	isSet bool
}

func (v NullableWsFederationApplicationSettingsAllOf) Get() *WsFederationApplicationSettingsAllOf {
	return v.value
}

func (v *NullableWsFederationApplicationSettingsAllOf) Set(val *WsFederationApplicationSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWsFederationApplicationSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWsFederationApplicationSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWsFederationApplicationSettingsAllOf(val *WsFederationApplicationSettingsAllOf) *NullableWsFederationApplicationSettingsAllOf {
	return &NullableWsFederationApplicationSettingsAllOf{value: val, isSet: true}
}

func (v NullableWsFederationApplicationSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWsFederationApplicationSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
