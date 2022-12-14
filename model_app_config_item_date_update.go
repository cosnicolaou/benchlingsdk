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

// AppConfigItemDateUpdate struct for AppConfigItemDateUpdate
type AppConfigItemDateUpdate struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewAppConfigItemDateUpdate instantiates a new AppConfigItemDateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemDateUpdate(type_ string, value string) *AppConfigItemDateUpdate {
	this := AppConfigItemDateUpdate{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAppConfigItemDateUpdateWithDefaults instantiates a new AppConfigItemDateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemDateUpdateWithDefaults() *AppConfigItemDateUpdate {
	this := AppConfigItemDateUpdate{}
	return &this
}

// GetType returns the Type field value
func (o *AppConfigItemDateUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDateUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfigItemDateUpdate) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AppConfigItemDateUpdate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDateUpdate) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AppConfigItemDateUpdate) SetValue(v string) {
	o.Value = v
}

func (o AppConfigItemDateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemDateUpdate struct {
	value *AppConfigItemDateUpdate
	isSet bool
}

func (v NullableAppConfigItemDateUpdate) Get() *AppConfigItemDateUpdate {
	return v.value
}

func (v *NullableAppConfigItemDateUpdate) Set(val *AppConfigItemDateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemDateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemDateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemDateUpdate(val *AppConfigItemDateUpdate) *NullableAppConfigItemDateUpdate {
	return &NullableAppConfigItemDateUpdate{value: val, isSet: true}
}

func (v NullableAppConfigItemDateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemDateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


