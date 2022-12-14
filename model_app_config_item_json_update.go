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

// AppConfigItemJsonUpdate struct for AppConfigItemJsonUpdate
type AppConfigItemJsonUpdate struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewAppConfigItemJsonUpdate instantiates a new AppConfigItemJsonUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemJsonUpdate(type_ string, value string) *AppConfigItemJsonUpdate {
	this := AppConfigItemJsonUpdate{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAppConfigItemJsonUpdateWithDefaults instantiates a new AppConfigItemJsonUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemJsonUpdateWithDefaults() *AppConfigItemJsonUpdate {
	this := AppConfigItemJsonUpdate{}
	return &this
}

// GetType returns the Type field value
func (o *AppConfigItemJsonUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemJsonUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfigItemJsonUpdate) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AppConfigItemJsonUpdate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemJsonUpdate) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AppConfigItemJsonUpdate) SetValue(v string) {
	o.Value = v
}

func (o AppConfigItemJsonUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemJsonUpdate struct {
	value *AppConfigItemJsonUpdate
	isSet bool
}

func (v NullableAppConfigItemJsonUpdate) Get() *AppConfigItemJsonUpdate {
	return v.value
}

func (v *NullableAppConfigItemJsonUpdate) Set(val *AppConfigItemJsonUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemJsonUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemJsonUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemJsonUpdate(val *AppConfigItemJsonUpdate) *NullableAppConfigItemJsonUpdate {
	return &NullableAppConfigItemJsonUpdate{value: val, isSet: true}
}

func (v NullableAppConfigItemJsonUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemJsonUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


