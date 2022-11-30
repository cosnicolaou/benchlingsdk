/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
)

// AppConfigItemIntegerUpdate struct for AppConfigItemIntegerUpdate
type AppConfigItemIntegerUpdate struct {
	Type string `json:"type"`
	Value int32 `json:"value"`
}

// NewAppConfigItemIntegerUpdate instantiates a new AppConfigItemIntegerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemIntegerUpdate(type_ string, value int32) *AppConfigItemIntegerUpdate {
	this := AppConfigItemIntegerUpdate{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAppConfigItemIntegerUpdateWithDefaults instantiates a new AppConfigItemIntegerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemIntegerUpdateWithDefaults() *AppConfigItemIntegerUpdate {
	this := AppConfigItemIntegerUpdate{}
	return &this
}

// GetType returns the Type field value
func (o *AppConfigItemIntegerUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemIntegerUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfigItemIntegerUpdate) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AppConfigItemIntegerUpdate) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemIntegerUpdate) GetValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AppConfigItemIntegerUpdate) SetValue(v int32) {
	o.Value = v
}

func (o AppConfigItemIntegerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemIntegerUpdate struct {
	value *AppConfigItemIntegerUpdate
	isSet bool
}

func (v NullableAppConfigItemIntegerUpdate) Get() *AppConfigItemIntegerUpdate {
	return v.value
}

func (v *NullableAppConfigItemIntegerUpdate) Set(val *AppConfigItemIntegerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemIntegerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemIntegerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemIntegerUpdate(val *AppConfigItemIntegerUpdate) *NullableAppConfigItemIntegerUpdate {
	return &NullableAppConfigItemIntegerUpdate{value: val, isSet: true}
}

func (v NullableAppConfigItemIntegerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemIntegerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

