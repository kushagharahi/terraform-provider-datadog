/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// ApplicationKeyResponse TODO.
type ApplicationKeyResponse struct {
	ApplicationKey *ApplicationKey `json:"application_key,omitempty"`
}

// NewApplicationKeyResponse instantiates a new ApplicationKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyResponse() *ApplicationKeyResponse {
	this := ApplicationKeyResponse{}
	return &this
}

// NewApplicationKeyResponseWithDefaults instantiates a new ApplicationKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyResponseWithDefaults() *ApplicationKeyResponse {
	this := ApplicationKeyResponse{}
	return &this
}

// GetApplicationKey returns the ApplicationKey field value if set, zero value otherwise.
func (o *ApplicationKeyResponse) GetApplicationKey() ApplicationKey {
	if o == nil || o.ApplicationKey == nil {
		var ret ApplicationKey
		return ret
	}
	return *o.ApplicationKey
}

// GetApplicationKeyOk returns a tuple with the ApplicationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyResponse) GetApplicationKeyOk() (*ApplicationKey, bool) {
	if o == nil || o.ApplicationKey == nil {
		return nil, false
	}
	return o.ApplicationKey, true
}

// HasApplicationKey returns a boolean if a field has been set.
func (o *ApplicationKeyResponse) HasApplicationKey() bool {
	if o != nil && o.ApplicationKey != nil {
		return true
	}

	return false
}

// SetApplicationKey gets a reference to the given ApplicationKey and assigns it to the ApplicationKey field.
func (o *ApplicationKeyResponse) SetApplicationKey(v ApplicationKey) {
	o.ApplicationKey = &v
}

func (o ApplicationKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationKey != nil {
		toSerialize["application_key"] = o.ApplicationKey
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationKeyResponse struct {
	value *ApplicationKeyResponse
	isSet bool
}

func (v NullableApplicationKeyResponse) Get() *ApplicationKeyResponse {
	return v.value
}

func (v *NullableApplicationKeyResponse) Set(val *ApplicationKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyResponse(val *ApplicationKeyResponse) *NullableApplicationKeyResponse {
	return &NullableApplicationKeyResponse{value: val, isSet: true}
}

func (v NullableApplicationKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}