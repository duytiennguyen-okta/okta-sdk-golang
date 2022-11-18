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

// AppLink struct for AppLink
type AppLink struct {
	AppAssignmentId      *string `json:"appAssignmentId,omitempty"`
	AppInstanceId        *string `json:"appInstanceId,omitempty"`
	AppName              *string `json:"appName,omitempty"`
	CredentialsSetup     *bool   `json:"credentialsSetup,omitempty"`
	Hidden               *bool   `json:"hidden,omitempty"`
	Id                   *string `json:"id,omitempty"`
	Label                *string `json:"label,omitempty"`
	LinkUrl              *string `json:"linkUrl,omitempty"`
	LogoUrl              *string `json:"logoUrl,omitempty"`
	SortOrder            *int32  `json:"sortOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppLink AppLink

// NewAppLink instantiates a new AppLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppLink() *AppLink {
	this := AppLink{}
	return &this
}

// NewAppLinkWithDefaults instantiates a new AppLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppLinkWithDefaults() *AppLink {
	this := AppLink{}
	return &this
}

// GetAppAssignmentId returns the AppAssignmentId field value if set, zero value otherwise.
func (o *AppLink) GetAppAssignmentId() string {
	if o == nil || o.AppAssignmentId == nil {
		var ret string
		return ret
	}
	return *o.AppAssignmentId
}

// GetAppAssignmentIdOk returns a tuple with the AppAssignmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetAppAssignmentIdOk() (*string, bool) {
	if o == nil || o.AppAssignmentId == nil {
		return nil, false
	}
	return o.AppAssignmentId, true
}

// HasAppAssignmentId returns a boolean if a field has been set.
func (o *AppLink) HasAppAssignmentId() bool {
	if o != nil && o.AppAssignmentId != nil {
		return true
	}

	return false
}

// SetAppAssignmentId gets a reference to the given string and assigns it to the AppAssignmentId field.
func (o *AppLink) SetAppAssignmentId(v string) {
	o.AppAssignmentId = &v
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *AppLink) GetAppInstanceId() string {
	if o == nil || o.AppInstanceId == nil {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || o.AppInstanceId == nil {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *AppLink) HasAppInstanceId() bool {
	if o != nil && o.AppInstanceId != nil {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *AppLink) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AppLink) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AppLink) HasAppName() bool {
	if o != nil && o.AppName != nil {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AppLink) SetAppName(v string) {
	o.AppName = &v
}

// GetCredentialsSetup returns the CredentialsSetup field value if set, zero value otherwise.
func (o *AppLink) GetCredentialsSetup() bool {
	if o == nil || o.CredentialsSetup == nil {
		var ret bool
		return ret
	}
	return *o.CredentialsSetup
}

// GetCredentialsSetupOk returns a tuple with the CredentialsSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetCredentialsSetupOk() (*bool, bool) {
	if o == nil || o.CredentialsSetup == nil {
		return nil, false
	}
	return o.CredentialsSetup, true
}

// HasCredentialsSetup returns a boolean if a field has been set.
func (o *AppLink) HasCredentialsSetup() bool {
	if o != nil && o.CredentialsSetup != nil {
		return true
	}

	return false
}

// SetCredentialsSetup gets a reference to the given bool and assigns it to the CredentialsSetup field.
func (o *AppLink) SetCredentialsSetup(v bool) {
	o.CredentialsSetup = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *AppLink) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *AppLink) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *AppLink) SetHidden(v bool) {
	o.Hidden = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppLink) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppLink) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppLink) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AppLink) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AppLink) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AppLink) SetLabel(v string) {
	o.Label = &v
}

// GetLinkUrl returns the LinkUrl field value if set, zero value otherwise.
func (o *AppLink) GetLinkUrl() string {
	if o == nil || o.LinkUrl == nil {
		var ret string
		return ret
	}
	return *o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetLinkUrlOk() (*string, bool) {
	if o == nil || o.LinkUrl == nil {
		return nil, false
	}
	return o.LinkUrl, true
}

// HasLinkUrl returns a boolean if a field has been set.
func (o *AppLink) HasLinkUrl() bool {
	if o != nil && o.LinkUrl != nil {
		return true
	}

	return false
}

// SetLinkUrl gets a reference to the given string and assigns it to the LinkUrl field.
func (o *AppLink) SetLinkUrl(v string) {
	o.LinkUrl = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AppLink) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AppLink) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AppLink) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AppLink) GetSortOrder() int32 {
	if o == nil || o.SortOrder == nil {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetSortOrderOk() (*int32, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AppLink) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AppLink) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o AppLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppAssignmentId != nil {
		toSerialize["appAssignmentId"] = o.AppAssignmentId
	}
	if o.AppInstanceId != nil {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}
	if o.AppName != nil {
		toSerialize["appName"] = o.AppName
	}
	if o.CredentialsSetup != nil {
		toSerialize["credentialsSetup"] = o.CredentialsSetup
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LinkUrl != nil {
		toSerialize["linkUrl"] = o.LinkUrl
	}
	if o.LogoUrl != nil {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppLink) UnmarshalJSON(bytes []byte) (err error) {
	varAppLink := _AppLink{}

	err = json.Unmarshal(bytes, &varAppLink)
	if err == nil {
		*o = AppLink(varAppLink)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appAssignmentId")
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appName")
		delete(additionalProperties, "credentialsSetup")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "linkUrl")
		delete(additionalProperties, "logoUrl")
		delete(additionalProperties, "sortOrder")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppLink struct {
	value *AppLink
	isSet bool
}

func (v NullableAppLink) Get() *AppLink {
	return v.value
}

func (v *NullableAppLink) Set(val *AppLink) {
	v.value = val
	v.isSet = true
}

func (v NullableAppLink) IsSet() bool {
	return v.isSet
}

func (v *NullableAppLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppLink(val *AppLink) *NullableAppLink {
	return &NullableAppLink{value: val, isSet: true}
}

func (v NullableAppLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
