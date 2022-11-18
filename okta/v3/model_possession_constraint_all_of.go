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

// PossessionConstraintAllOf struct for PossessionConstraintAllOf
type PossessionConstraintAllOf struct {
	DeviceBound          *string `json:"deviceBound,omitempty"`
	HardwareProtection   *string `json:"hardwareProtection,omitempty"`
	PhishingResistant    *string `json:"phishingResistant,omitempty"`
	UserPresence         *string `json:"userPresence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PossessionConstraintAllOf PossessionConstraintAllOf

// NewPossessionConstraintAllOf instantiates a new PossessionConstraintAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPossessionConstraintAllOf() *PossessionConstraintAllOf {
	this := PossessionConstraintAllOf{}
	return &this
}

// NewPossessionConstraintAllOfWithDefaults instantiates a new PossessionConstraintAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPossessionConstraintAllOfWithDefaults() *PossessionConstraintAllOf {
	this := PossessionConstraintAllOf{}
	return &this
}

// GetDeviceBound returns the DeviceBound field value if set, zero value otherwise.
func (o *PossessionConstraintAllOf) GetDeviceBound() string {
	if o == nil || o.DeviceBound == nil {
		var ret string
		return ret
	}
	return *o.DeviceBound
}

// GetDeviceBoundOk returns a tuple with the DeviceBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraintAllOf) GetDeviceBoundOk() (*string, bool) {
	if o == nil || o.DeviceBound == nil {
		return nil, false
	}
	return o.DeviceBound, true
}

// HasDeviceBound returns a boolean if a field has been set.
func (o *PossessionConstraintAllOf) HasDeviceBound() bool {
	if o != nil && o.DeviceBound != nil {
		return true
	}

	return false
}

// SetDeviceBound gets a reference to the given string and assigns it to the DeviceBound field.
func (o *PossessionConstraintAllOf) SetDeviceBound(v string) {
	o.DeviceBound = &v
}

// GetHardwareProtection returns the HardwareProtection field value if set, zero value otherwise.
func (o *PossessionConstraintAllOf) GetHardwareProtection() string {
	if o == nil || o.HardwareProtection == nil {
		var ret string
		return ret
	}
	return *o.HardwareProtection
}

// GetHardwareProtectionOk returns a tuple with the HardwareProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraintAllOf) GetHardwareProtectionOk() (*string, bool) {
	if o == nil || o.HardwareProtection == nil {
		return nil, false
	}
	return o.HardwareProtection, true
}

// HasHardwareProtection returns a boolean if a field has been set.
func (o *PossessionConstraintAllOf) HasHardwareProtection() bool {
	if o != nil && o.HardwareProtection != nil {
		return true
	}

	return false
}

// SetHardwareProtection gets a reference to the given string and assigns it to the HardwareProtection field.
func (o *PossessionConstraintAllOf) SetHardwareProtection(v string) {
	o.HardwareProtection = &v
}

// GetPhishingResistant returns the PhishingResistant field value if set, zero value otherwise.
func (o *PossessionConstraintAllOf) GetPhishingResistant() string {
	if o == nil || o.PhishingResistant == nil {
		var ret string
		return ret
	}
	return *o.PhishingResistant
}

// GetPhishingResistantOk returns a tuple with the PhishingResistant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraintAllOf) GetPhishingResistantOk() (*string, bool) {
	if o == nil || o.PhishingResistant == nil {
		return nil, false
	}
	return o.PhishingResistant, true
}

// HasPhishingResistant returns a boolean if a field has been set.
func (o *PossessionConstraintAllOf) HasPhishingResistant() bool {
	if o != nil && o.PhishingResistant != nil {
		return true
	}

	return false
}

// SetPhishingResistant gets a reference to the given string and assigns it to the PhishingResistant field.
func (o *PossessionConstraintAllOf) SetPhishingResistant(v string) {
	o.PhishingResistant = &v
}

// GetUserPresence returns the UserPresence field value if set, zero value otherwise.
func (o *PossessionConstraintAllOf) GetUserPresence() string {
	if o == nil || o.UserPresence == nil {
		var ret string
		return ret
	}
	return *o.UserPresence
}

// GetUserPresenceOk returns a tuple with the UserPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraintAllOf) GetUserPresenceOk() (*string, bool) {
	if o == nil || o.UserPresence == nil {
		return nil, false
	}
	return o.UserPresence, true
}

// HasUserPresence returns a boolean if a field has been set.
func (o *PossessionConstraintAllOf) HasUserPresence() bool {
	if o != nil && o.UserPresence != nil {
		return true
	}

	return false
}

// SetUserPresence gets a reference to the given string and assigns it to the UserPresence field.
func (o *PossessionConstraintAllOf) SetUserPresence(v string) {
	o.UserPresence = &v
}

func (o PossessionConstraintAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceBound != nil {
		toSerialize["deviceBound"] = o.DeviceBound
	}
	if o.HardwareProtection != nil {
		toSerialize["hardwareProtection"] = o.HardwareProtection
	}
	if o.PhishingResistant != nil {
		toSerialize["phishingResistant"] = o.PhishingResistant
	}
	if o.UserPresence != nil {
		toSerialize["userPresence"] = o.UserPresence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PossessionConstraintAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPossessionConstraintAllOf := _PossessionConstraintAllOf{}

	err = json.Unmarshal(bytes, &varPossessionConstraintAllOf)
	if err == nil {
		*o = PossessionConstraintAllOf(varPossessionConstraintAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "deviceBound")
		delete(additionalProperties, "hardwareProtection")
		delete(additionalProperties, "phishingResistant")
		delete(additionalProperties, "userPresence")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePossessionConstraintAllOf struct {
	value *PossessionConstraintAllOf
	isSet bool
}

func (v NullablePossessionConstraintAllOf) Get() *PossessionConstraintAllOf {
	return v.value
}

func (v *NullablePossessionConstraintAllOf) Set(val *PossessionConstraintAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePossessionConstraintAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePossessionConstraintAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePossessionConstraintAllOf(val *PossessionConstraintAllOf) *NullablePossessionConstraintAllOf {
	return &NullablePossessionConstraintAllOf{value: val, isSet: true}
}

func (v NullablePossessionConstraintAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePossessionConstraintAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
