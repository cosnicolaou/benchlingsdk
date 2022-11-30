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

// DatetimeAppConfigItemAllOf struct for DatetimeAppConfigItemAllOf
type DatetimeAppConfigItemAllOf struct {
	Type *string `json:"type,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewDatetimeAppConfigItemAllOf instantiates a new DatetimeAppConfigItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatetimeAppConfigItemAllOf() *DatetimeAppConfigItemAllOf {
	this := DatetimeAppConfigItemAllOf{}
	return &this
}

// NewDatetimeAppConfigItemAllOfWithDefaults instantiates a new DatetimeAppConfigItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatetimeAppConfigItemAllOfWithDefaults() *DatetimeAppConfigItemAllOf {
	this := DatetimeAppConfigItemAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DatetimeAppConfigItemAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatetimeAppConfigItemAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DatetimeAppConfigItemAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DatetimeAppConfigItemAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatetimeAppConfigItemAllOf) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatetimeAppConfigItemAllOf) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *DatetimeAppConfigItemAllOf) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *DatetimeAppConfigItemAllOf) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *DatetimeAppConfigItemAllOf) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *DatetimeAppConfigItemAllOf) UnsetValue() {
	o.Value.Unset()
}

func (o DatetimeAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDatetimeAppConfigItemAllOf struct {
	value *DatetimeAppConfigItemAllOf
	isSet bool
}

func (v NullableDatetimeAppConfigItemAllOf) Get() *DatetimeAppConfigItemAllOf {
	return v.value
}

func (v *NullableDatetimeAppConfigItemAllOf) Set(val *DatetimeAppConfigItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDatetimeAppConfigItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDatetimeAppConfigItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatetimeAppConfigItemAllOf(val *DatetimeAppConfigItemAllOf) *NullableDatetimeAppConfigItemAllOf {
	return &NullableDatetimeAppConfigItemAllOf{value: val, isSet: true}
}

func (v NullableDatetimeAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatetimeAppConfigItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


