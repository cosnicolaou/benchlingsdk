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

// JsonAppConfigItemAllOf struct for JsonAppConfigItemAllOf
type JsonAppConfigItemAllOf struct {
	Type *string `json:"type,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewJsonAppConfigItemAllOf instantiates a new JsonAppConfigItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonAppConfigItemAllOf() *JsonAppConfigItemAllOf {
	this := JsonAppConfigItemAllOf{}
	return &this
}

// NewJsonAppConfigItemAllOfWithDefaults instantiates a new JsonAppConfigItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonAppConfigItemAllOfWithDefaults() *JsonAppConfigItemAllOf {
	this := JsonAppConfigItemAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JsonAppConfigItemAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonAppConfigItemAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JsonAppConfigItemAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JsonAppConfigItemAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JsonAppConfigItemAllOf) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JsonAppConfigItemAllOf) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *JsonAppConfigItemAllOf) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *JsonAppConfigItemAllOf) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *JsonAppConfigItemAllOf) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *JsonAppConfigItemAllOf) UnsetValue() {
	o.Value.Unset()
}

func (o JsonAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJsonAppConfigItemAllOf struct {
	value *JsonAppConfigItemAllOf
	isSet bool
}

func (v NullableJsonAppConfigItemAllOf) Get() *JsonAppConfigItemAllOf {
	return v.value
}

func (v *NullableJsonAppConfigItemAllOf) Set(val *JsonAppConfigItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonAppConfigItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonAppConfigItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonAppConfigItemAllOf(val *JsonAppConfigItemAllOf) *NullableJsonAppConfigItemAllOf {
	return &NullableJsonAppConfigItemAllOf{value: val, isSet: true}
}

func (v NullableJsonAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonAppConfigItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


