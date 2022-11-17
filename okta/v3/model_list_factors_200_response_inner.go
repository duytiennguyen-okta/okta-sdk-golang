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

// model_oneof.mustache
// ListFactors200ResponseInner - struct for ListFactors200ResponseInner
type ListFactors200ResponseInner struct {
	CallUserFactor             *CallUserFactor
	CustomHotpUserFactor       *CustomHotpUserFactor
	EmailUserFactor            *EmailUserFactor
	HardwareUserFactor         *HardwareUserFactor
	PushUserFactor             *PushUserFactor
	SecurityQuestionUserFactor *SecurityQuestionUserFactor
	SmsUserFactor              *SmsUserFactor
	TokenUserFactor            *TokenUserFactor
	TotpUserFactor             *TotpUserFactor
	U2fUserFactor              *U2fUserFactor
	WebAuthnUserFactor         *WebAuthnUserFactor
	WebUserFactor              *WebUserFactor
}

// CallUserFactorAsListFactors200ResponseInner is a convenience function that returns CallUserFactor wrapped in ListFactors200ResponseInner
func CallUserFactorAsListFactors200ResponseInner(v *CallUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		CallUserFactor: v,
	}
}

// CustomHotpUserFactorAsListFactors200ResponseInner is a convenience function that returns CustomHotpUserFactor wrapped in ListFactors200ResponseInner
func CustomHotpUserFactorAsListFactors200ResponseInner(v *CustomHotpUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		CustomHotpUserFactor: v,
	}
}

// EmailUserFactorAsListFactors200ResponseInner is a convenience function that returns EmailUserFactor wrapped in ListFactors200ResponseInner
func EmailUserFactorAsListFactors200ResponseInner(v *EmailUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		EmailUserFactor: v,
	}
}

// HardwareUserFactorAsListFactors200ResponseInner is a convenience function that returns HardwareUserFactor wrapped in ListFactors200ResponseInner
func HardwareUserFactorAsListFactors200ResponseInner(v *HardwareUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		HardwareUserFactor: v,
	}
}

// PushUserFactorAsListFactors200ResponseInner is a convenience function that returns PushUserFactor wrapped in ListFactors200ResponseInner
func PushUserFactorAsListFactors200ResponseInner(v *PushUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		PushUserFactor: v,
	}
}

// SecurityQuestionUserFactorAsListFactors200ResponseInner is a convenience function that returns SecurityQuestionUserFactor wrapped in ListFactors200ResponseInner
func SecurityQuestionUserFactorAsListFactors200ResponseInner(v *SecurityQuestionUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		SecurityQuestionUserFactor: v,
	}
}

// SmsUserFactorAsListFactors200ResponseInner is a convenience function that returns SmsUserFactor wrapped in ListFactors200ResponseInner
func SmsUserFactorAsListFactors200ResponseInner(v *SmsUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		SmsUserFactor: v,
	}
}

// TokenUserFactorAsListFactors200ResponseInner is a convenience function that returns TokenUserFactor wrapped in ListFactors200ResponseInner
func TokenUserFactorAsListFactors200ResponseInner(v *TokenUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		TokenUserFactor: v,
	}
}

// TotpUserFactorAsListFactors200ResponseInner is a convenience function that returns TotpUserFactor wrapped in ListFactors200ResponseInner
func TotpUserFactorAsListFactors200ResponseInner(v *TotpUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		TotpUserFactor: v,
	}
}

// U2fUserFactorAsListFactors200ResponseInner is a convenience function that returns U2fUserFactor wrapped in ListFactors200ResponseInner
func U2fUserFactorAsListFactors200ResponseInner(v *U2fUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		U2fUserFactor: v,
	}
}

// WebAuthnUserFactorAsListFactors200ResponseInner is a convenience function that returns WebAuthnUserFactor wrapped in ListFactors200ResponseInner
func WebAuthnUserFactorAsListFactors200ResponseInner(v *WebAuthnUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		WebAuthnUserFactor: v,
	}
}

// WebUserFactorAsListFactors200ResponseInner is a convenience function that returns WebUserFactor wrapped in ListFactors200ResponseInner
func WebUserFactorAsListFactors200ResponseInner(v *WebUserFactor) ListFactors200ResponseInner {
	return ListFactors200ResponseInner{
		WebUserFactor: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListFactors200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CallUserFactor'
	if jsonDict["factorType"] == "CallUserFactor" {
		// try to unmarshal JSON data into CallUserFactor
		err = json.Unmarshal(data, &dst.CallUserFactor)
		if err == nil {
			return nil // data stored in dst.CallUserFactor, return on the first match
		} else {
			dst.CallUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as CallUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustomHotpUserFactor'
	if jsonDict["factorType"] == "CustomHotpUserFactor" {
		// try to unmarshal JSON data into CustomHotpUserFactor
		err = json.Unmarshal(data, &dst.CustomHotpUserFactor)
		if err == nil {
			return nil // data stored in dst.CustomHotpUserFactor, return on the first match
		} else {
			dst.CustomHotpUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as CustomHotpUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EmailUserFactor'
	if jsonDict["factorType"] == "EmailUserFactor" {
		// try to unmarshal JSON data into EmailUserFactor
		err = json.Unmarshal(data, &dst.EmailUserFactor)
		if err == nil {
			return nil // data stored in dst.EmailUserFactor, return on the first match
		} else {
			dst.EmailUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as EmailUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'HardwareUserFactor'
	if jsonDict["factorType"] == "HardwareUserFactor" {
		// try to unmarshal JSON data into HardwareUserFactor
		err = json.Unmarshal(data, &dst.HardwareUserFactor)
		if err == nil {
			return nil // data stored in dst.HardwareUserFactor, return on the first match
		} else {
			dst.HardwareUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as HardwareUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PushUserFactor'
	if jsonDict["factorType"] == "PushUserFactor" {
		// try to unmarshal JSON data into PushUserFactor
		err = json.Unmarshal(data, &dst.PushUserFactor)
		if err == nil {
			return nil // data stored in dst.PushUserFactor, return on the first match
		} else {
			dst.PushUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as PushUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurityQuestionUserFactor'
	if jsonDict["factorType"] == "SecurityQuestionUserFactor" {
		// try to unmarshal JSON data into SecurityQuestionUserFactor
		err = json.Unmarshal(data, &dst.SecurityQuestionUserFactor)
		if err == nil {
			return nil // data stored in dst.SecurityQuestionUserFactor, return on the first match
		} else {
			dst.SecurityQuestionUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as SecurityQuestionUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SmsUserFactor'
	if jsonDict["factorType"] == "SmsUserFactor" {
		// try to unmarshal JSON data into SmsUserFactor
		err = json.Unmarshal(data, &dst.SmsUserFactor)
		if err == nil {
			return nil // data stored in dst.SmsUserFactor, return on the first match
		} else {
			dst.SmsUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as SmsUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TokenUserFactor'
	if jsonDict["factorType"] == "TokenUserFactor" {
		// try to unmarshal JSON data into TokenUserFactor
		err = json.Unmarshal(data, &dst.TokenUserFactor)
		if err == nil {
			return nil // data stored in dst.TokenUserFactor, return on the first match
		} else {
			dst.TokenUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as TokenUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TotpUserFactor'
	if jsonDict["factorType"] == "TotpUserFactor" {
		// try to unmarshal JSON data into TotpUserFactor
		err = json.Unmarshal(data, &dst.TotpUserFactor)
		if err == nil {
			return nil // data stored in dst.TotpUserFactor, return on the first match
		} else {
			dst.TotpUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as TotpUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'U2fUserFactor'
	if jsonDict["factorType"] == "U2fUserFactor" {
		// try to unmarshal JSON data into U2fUserFactor
		err = json.Unmarshal(data, &dst.U2fUserFactor)
		if err == nil {
			return nil // data stored in dst.U2fUserFactor, return on the first match
		} else {
			dst.U2fUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as U2fUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WebAuthnUserFactor'
	if jsonDict["factorType"] == "WebAuthnUserFactor" {
		// try to unmarshal JSON data into WebAuthnUserFactor
		err = json.Unmarshal(data, &dst.WebAuthnUserFactor)
		if err == nil {
			return nil // data stored in dst.WebAuthnUserFactor, return on the first match
		} else {
			dst.WebAuthnUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as WebAuthnUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WebUserFactor'
	if jsonDict["factorType"] == "WebUserFactor" {
		// try to unmarshal JSON data into WebUserFactor
		err = json.Unmarshal(data, &dst.WebUserFactor)
		if err == nil {
			return nil // data stored in dst.WebUserFactor, return on the first match
		} else {
			dst.WebUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as WebUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'call'
	if jsonDict["factorType"] == "call" {
		// try to unmarshal JSON data into CallUserFactor
		err = json.Unmarshal(data, &dst.CallUserFactor)
		if err == nil {
			return nil // data stored in dst.CallUserFactor, return on the first match
		} else {
			dst.CallUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as CallUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'email'
	if jsonDict["factorType"] == "email" {
		// try to unmarshal JSON data into EmailUserFactor
		err = json.Unmarshal(data, &dst.EmailUserFactor)
		if err == nil {
			return nil // data stored in dst.EmailUserFactor, return on the first match
		} else {
			dst.EmailUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as EmailUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'hotp'
	if jsonDict["factorType"] == "hotp" {
		// try to unmarshal JSON data into CustomHotpUserFactor
		err = json.Unmarshal(data, &dst.CustomHotpUserFactor)
		if err == nil {
			return nil // data stored in dst.CustomHotpUserFactor, return on the first match
		} else {
			dst.CustomHotpUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as CustomHotpUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'push'
	if jsonDict["factorType"] == "push" {
		// try to unmarshal JSON data into PushUserFactor
		err = json.Unmarshal(data, &dst.PushUserFactor)
		if err == nil {
			return nil // data stored in dst.PushUserFactor, return on the first match
		} else {
			dst.PushUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as PushUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'question'
	if jsonDict["factorType"] == "question" {
		// try to unmarshal JSON data into SecurityQuestionUserFactor
		err = json.Unmarshal(data, &dst.SecurityQuestionUserFactor)
		if err == nil {
			return nil // data stored in dst.SecurityQuestionUserFactor, return on the first match
		} else {
			dst.SecurityQuestionUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as SecurityQuestionUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'sms'
	if jsonDict["factorType"] == "sms" {
		// try to unmarshal JSON data into SmsUserFactor
		err = json.Unmarshal(data, &dst.SmsUserFactor)
		if err == nil {
			return nil // data stored in dst.SmsUserFactor, return on the first match
		} else {
			dst.SmsUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as SmsUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token'
	if jsonDict["factorType"] == "token" {
		// try to unmarshal JSON data into TokenUserFactor
		err = json.Unmarshal(data, &dst.TokenUserFactor)
		if err == nil {
			return nil // data stored in dst.TokenUserFactor, return on the first match
		} else {
			dst.TokenUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as TokenUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token:hardware'
	if jsonDict["factorType"] == "token:hardware" {
		// try to unmarshal JSON data into HardwareUserFactor
		err = json.Unmarshal(data, &dst.HardwareUserFactor)
		if err == nil {
			return nil // data stored in dst.HardwareUserFactor, return on the first match
		} else {
			dst.HardwareUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as HardwareUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token:hotp'
	if jsonDict["factorType"] == "token:hotp" {
		// try to unmarshal JSON data into CustomHotpUserFactor
		err = json.Unmarshal(data, &dst.CustomHotpUserFactor)
		if err == nil {
			return nil // data stored in dst.CustomHotpUserFactor, return on the first match
		} else {
			dst.CustomHotpUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as CustomHotpUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'token:software:totp'
	if jsonDict["factorType"] == "token:software:totp" {
		// try to unmarshal JSON data into TotpUserFactor
		err = json.Unmarshal(data, &dst.TotpUserFactor)
		if err == nil {
			return nil // data stored in dst.TotpUserFactor, return on the first match
		} else {
			dst.TotpUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as TotpUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'u2f'
	if jsonDict["factorType"] == "u2f" {
		// try to unmarshal JSON data into U2fUserFactor
		err = json.Unmarshal(data, &dst.U2fUserFactor)
		if err == nil {
			return nil // data stored in dst.U2fUserFactor, return on the first match
		} else {
			dst.U2fUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as U2fUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'web'
	if jsonDict["factorType"] == "web" {
		// try to unmarshal JSON data into WebUserFactor
		err = json.Unmarshal(data, &dst.WebUserFactor)
		if err == nil {
			return nil // data stored in dst.WebUserFactor, return on the first match
		} else {
			dst.WebUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as WebUserFactor: %s", err.Error())
		}
	}

	// check if the discriminator value is 'webauthn'
	if jsonDict["factorType"] == "webauthn" {
		// try to unmarshal JSON data into WebAuthnUserFactor
		err = json.Unmarshal(data, &dst.WebAuthnUserFactor)
		if err == nil {
			return nil // data stored in dst.WebAuthnUserFactor, return on the first match
		} else {
			dst.WebAuthnUserFactor = nil
			return fmt.Errorf("Failed to unmarshal ListFactors200ResponseInner as WebAuthnUserFactor: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListFactors200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.CallUserFactor != nil {
		return json.Marshal(&src.CallUserFactor)
	}

	if src.CustomHotpUserFactor != nil {
		return json.Marshal(&src.CustomHotpUserFactor)
	}

	if src.EmailUserFactor != nil {
		return json.Marshal(&src.EmailUserFactor)
	}

	if src.HardwareUserFactor != nil {
		return json.Marshal(&src.HardwareUserFactor)
	}

	if src.PushUserFactor != nil {
		return json.Marshal(&src.PushUserFactor)
	}

	if src.SecurityQuestionUserFactor != nil {
		return json.Marshal(&src.SecurityQuestionUserFactor)
	}

	if src.SmsUserFactor != nil {
		return json.Marshal(&src.SmsUserFactor)
	}

	if src.TokenUserFactor != nil {
		return json.Marshal(&src.TokenUserFactor)
	}

	if src.TotpUserFactor != nil {
		return json.Marshal(&src.TotpUserFactor)
	}

	if src.U2fUserFactor != nil {
		return json.Marshal(&src.U2fUserFactor)
	}

	if src.WebAuthnUserFactor != nil {
		return json.Marshal(&src.WebAuthnUserFactor)
	}

	if src.WebUserFactor != nil {
		return json.Marshal(&src.WebUserFactor)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListFactors200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CallUserFactor != nil {
		return obj.CallUserFactor
	}

	if obj.CustomHotpUserFactor != nil {
		return obj.CustomHotpUserFactor
	}

	if obj.EmailUserFactor != nil {
		return obj.EmailUserFactor
	}

	if obj.HardwareUserFactor != nil {
		return obj.HardwareUserFactor
	}

	if obj.PushUserFactor != nil {
		return obj.PushUserFactor
	}

	if obj.SecurityQuestionUserFactor != nil {
		return obj.SecurityQuestionUserFactor
	}

	if obj.SmsUserFactor != nil {
		return obj.SmsUserFactor
	}

	if obj.TokenUserFactor != nil {
		return obj.TokenUserFactor
	}

	if obj.TotpUserFactor != nil {
		return obj.TotpUserFactor
	}

	if obj.U2fUserFactor != nil {
		return obj.U2fUserFactor
	}

	if obj.WebAuthnUserFactor != nil {
		return obj.WebAuthnUserFactor
	}

	if obj.WebUserFactor != nil {
		return obj.WebUserFactor
	}

	// all schemas are nil
	return nil
}

type NullableListFactors200ResponseInner struct {
	value *ListFactors200ResponseInner
	isSet bool
}

func (v NullableListFactors200ResponseInner) Get() *ListFactors200ResponseInner {
	return v.value
}

func (v *NullableListFactors200ResponseInner) Set(val *ListFactors200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListFactors200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListFactors200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFactors200ResponseInner(val *ListFactors200ResponseInner) *NullableListFactors200ResponseInner {
	return &NullableListFactors200ResponseInner{value: val, isSet: true}
}

func (v NullableListFactors200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFactors200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
