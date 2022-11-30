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

// AppConfigItemJsonCreateAllOf struct for AppConfigItemJsonCreateAllOf
type AppConfigItemJsonCreateAllOf struct {
	Type *string `json:"type,omitempty"`
	// The value of a json create object should be json parseable.
	Value *string `json:"value,omitempty"`
}

// NewAppConfigItemJsonCreateAllOf instantiates a new AppConfigItemJsonCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemJsonCreateAllOf() *AppConfigItemJsonCreateAllOf {
	this := AppConfigItemJsonCreateAllOf{}
	return &this
}

// NewAppConfigItemJsonCreateAllOfWithDefaults instantiates a new AppConfigItemJsonCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemJsonCreateAllOfWithDefaults() *AppConfigItemJsonCreateAllOf {
	this := AppConfigItemJsonCreateAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppConfigItemJsonCreateAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigItemJsonCreateAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppConfigItemJsonCreateAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppConfigItemJsonCreateAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AppConfigItemJsonCreateAllOf) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigItemJsonCreateAllOf) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AppConfigItemJsonCreateAllOf) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AppConfigItemJsonCreateAllOf) SetValue(v string) {
	o.Value = &v
}

func (o AppConfigItemJsonCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemJsonCreateAllOf struct {
	value *AppConfigItemJsonCreateAllOf
	isSet bool
}

func (v NullableAppConfigItemJsonCreateAllOf) Get() *AppConfigItemJsonCreateAllOf {
	return v.value
}

func (v *NullableAppConfigItemJsonCreateAllOf) Set(val *AppConfigItemJsonCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemJsonCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemJsonCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemJsonCreateAllOf(val *AppConfigItemJsonCreateAllOf) *NullableAppConfigItemJsonCreateAllOf {
	return &NullableAppConfigItemJsonCreateAllOf{value: val, isSet: true}
}

func (v NullableAppConfigItemJsonCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemJsonCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

