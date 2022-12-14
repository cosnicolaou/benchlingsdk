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

// GenericApiIdentifiedAppConfigItemAllOf struct for GenericApiIdentifiedAppConfigItemAllOf
type GenericApiIdentifiedAppConfigItemAllOf struct {
	Type *string `json:"type,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewGenericApiIdentifiedAppConfigItemAllOf instantiates a new GenericApiIdentifiedAppConfigItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericApiIdentifiedAppConfigItemAllOf() *GenericApiIdentifiedAppConfigItemAllOf {
	this := GenericApiIdentifiedAppConfigItemAllOf{}
	return &this
}

// NewGenericApiIdentifiedAppConfigItemAllOfWithDefaults instantiates a new GenericApiIdentifiedAppConfigItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericApiIdentifiedAppConfigItemAllOfWithDefaults() *GenericApiIdentifiedAppConfigItemAllOf {
	this := GenericApiIdentifiedAppConfigItemAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItemAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItemAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItemAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GenericApiIdentifiedAppConfigItemAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericApiIdentifiedAppConfigItemAllOf) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericApiIdentifiedAppConfigItemAllOf) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItemAllOf) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *GenericApiIdentifiedAppConfigItemAllOf) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *GenericApiIdentifiedAppConfigItemAllOf) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *GenericApiIdentifiedAppConfigItemAllOf) UnsetValue() {
	o.Value.Unset()
}

func (o GenericApiIdentifiedAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGenericApiIdentifiedAppConfigItemAllOf struct {
	value *GenericApiIdentifiedAppConfigItemAllOf
	isSet bool
}

func (v NullableGenericApiIdentifiedAppConfigItemAllOf) Get() *GenericApiIdentifiedAppConfigItemAllOf {
	return v.value
}

func (v *NullableGenericApiIdentifiedAppConfigItemAllOf) Set(val *GenericApiIdentifiedAppConfigItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericApiIdentifiedAppConfigItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericApiIdentifiedAppConfigItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericApiIdentifiedAppConfigItemAllOf(val *GenericApiIdentifiedAppConfigItemAllOf) *NullableGenericApiIdentifiedAppConfigItemAllOf {
	return &NullableGenericApiIdentifiedAppConfigItemAllOf{value: val, isSet: true}
}

func (v NullableGenericApiIdentifiedAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericApiIdentifiedAppConfigItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


