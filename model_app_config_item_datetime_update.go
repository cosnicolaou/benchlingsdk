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

// AppConfigItemDatetimeUpdate struct for AppConfigItemDatetimeUpdate
type AppConfigItemDatetimeUpdate struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewAppConfigItemDatetimeUpdate instantiates a new AppConfigItemDatetimeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemDatetimeUpdate(type_ string, value string) *AppConfigItemDatetimeUpdate {
	this := AppConfigItemDatetimeUpdate{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAppConfigItemDatetimeUpdateWithDefaults instantiates a new AppConfigItemDatetimeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemDatetimeUpdateWithDefaults() *AppConfigItemDatetimeUpdate {
	this := AppConfigItemDatetimeUpdate{}
	return &this
}

// GetType returns the Type field value
func (o *AppConfigItemDatetimeUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDatetimeUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfigItemDatetimeUpdate) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AppConfigItemDatetimeUpdate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDatetimeUpdate) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AppConfigItemDatetimeUpdate) SetValue(v string) {
	o.Value = v
}

func (o AppConfigItemDatetimeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemDatetimeUpdate struct {
	value *AppConfigItemDatetimeUpdate
	isSet bool
}

func (v NullableAppConfigItemDatetimeUpdate) Get() *AppConfigItemDatetimeUpdate {
	return v.value
}

func (v *NullableAppConfigItemDatetimeUpdate) Set(val *AppConfigItemDatetimeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemDatetimeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemDatetimeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemDatetimeUpdate(val *AppConfigItemDatetimeUpdate) *NullableAppConfigItemDatetimeUpdate {
	return &NullableAppConfigItemDatetimeUpdate{value: val, isSet: true}
}

func (v NullableAppConfigItemDatetimeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemDatetimeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

