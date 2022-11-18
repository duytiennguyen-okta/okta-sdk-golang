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

// AuthorizationServerPolicyRuleActionsAllOf struct for AuthorizationServerPolicyRuleActionsAllOf
type AuthorizationServerPolicyRuleActionsAllOf struct {
	Token                *TokenAuthorizationServerPolicyRuleAction `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleActionsAllOf AuthorizationServerPolicyRuleActionsAllOf

// NewAuthorizationServerPolicyRuleActionsAllOf instantiates a new AuthorizationServerPolicyRuleActionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleActionsAllOf() *AuthorizationServerPolicyRuleActionsAllOf {
	this := AuthorizationServerPolicyRuleActionsAllOf{}
	return &this
}

// NewAuthorizationServerPolicyRuleActionsAllOfWithDefaults instantiates a new AuthorizationServerPolicyRuleActionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleActionsAllOfWithDefaults() *AuthorizationServerPolicyRuleActionsAllOf {
	this := AuthorizationServerPolicyRuleActionsAllOf{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleActionsAllOf) GetToken() TokenAuthorizationServerPolicyRuleAction {
	if o == nil || o.Token == nil {
		var ret TokenAuthorizationServerPolicyRuleAction
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleActionsAllOf) GetTokenOk() (*TokenAuthorizationServerPolicyRuleAction, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleActionsAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given TokenAuthorizationServerPolicyRuleAction and assigns it to the Token field.
func (o *AuthorizationServerPolicyRuleActionsAllOf) SetToken(v TokenAuthorizationServerPolicyRuleAction) {
	o.Token = &v
}

func (o AuthorizationServerPolicyRuleActionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRuleActionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyRuleActionsAllOf := _AuthorizationServerPolicyRuleActionsAllOf{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleActionsAllOf)
	if err == nil {
		*o = AuthorizationServerPolicyRuleActionsAllOf(varAuthorizationServerPolicyRuleActionsAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyRuleActionsAllOf struct {
	value *AuthorizationServerPolicyRuleActionsAllOf
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleActionsAllOf) Get() *AuthorizationServerPolicyRuleActionsAllOf {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleActionsAllOf) Set(val *AuthorizationServerPolicyRuleActionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleActionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleActionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleActionsAllOf(val *AuthorizationServerPolicyRuleActionsAllOf) *NullableAuthorizationServerPolicyRuleActionsAllOf {
	return &NullableAuthorizationServerPolicyRuleActionsAllOf{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleActionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleActionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
