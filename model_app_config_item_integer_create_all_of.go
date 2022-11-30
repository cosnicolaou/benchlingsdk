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

// AppConfigItemIntegerCreateAllOf struct for AppConfigItemIntegerCreateAllOf
type AppConfigItemIntegerCreateAllOf struct {
	Type *string `json:"type,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewAppConfigItemIntegerCreateAllOf instantiates a new AppConfigItemIntegerCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemIntegerCreateAllOf() *AppConfigItemIntegerCreateAllOf {
	this := AppConfigItemIntegerCreateAllOf{}
	return &this
}

// NewAppConfigItemIntegerCreateAllOfWithDefaults instantiates a new AppConfigItemIntegerCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemIntegerCreateAllOfWithDefaults() *AppConfigItemIntegerCreateAllOf {
	this := AppConfigItemIntegerCreateAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppConfigItemIntegerCreateAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigItemIntegerCreateAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppConfigItemIntegerCreateAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppConfigItemIntegerCreateAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AppConfigItemIntegerCreateAllOf) GetValue() int32 {
	if o == nil || isNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigItemIntegerCreateAllOf) GetValueOk() (*int32, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AppConfigItemIntegerCreateAllOf) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *AppConfigItemIntegerCreateAllOf) SetValue(v int32) {
	o.Value = &v
}

func (o AppConfigItemIntegerCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemIntegerCreateAllOf struct {
	value *AppConfigItemIntegerCreateAllOf
	isSet bool
}

func (v NullableAppConfigItemIntegerCreateAllOf) Get() *AppConfigItemIntegerCreateAllOf {
	return v.value
}

func (v *NullableAppConfigItemIntegerCreateAllOf) Set(val *AppConfigItemIntegerCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemIntegerCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemIntegerCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemIntegerCreateAllOf(val *AppConfigItemIntegerCreateAllOf) *NullableAppConfigItemIntegerCreateAllOf {
	return &NullableAppConfigItemIntegerCreateAllOf{value: val, isSet: true}
}

func (v NullableAppConfigItemIntegerCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemIntegerCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

