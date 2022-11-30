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

// DeprecatedEntitySchemaAllOf struct for DeprecatedEntitySchemaAllOf
type DeprecatedEntitySchemaAllOf struct {
	Type *string `json:"type,omitempty"`
}

// NewDeprecatedEntitySchemaAllOf instantiates a new DeprecatedEntitySchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedEntitySchemaAllOf() *DeprecatedEntitySchemaAllOf {
	this := DeprecatedEntitySchemaAllOf{}
	return &this
}

// NewDeprecatedEntitySchemaAllOfWithDefaults instantiates a new DeprecatedEntitySchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedEntitySchemaAllOfWithDefaults() *DeprecatedEntitySchemaAllOf {
	this := DeprecatedEntitySchemaAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeprecatedEntitySchemaAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedEntitySchemaAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeprecatedEntitySchemaAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeprecatedEntitySchemaAllOf) SetType(v string) {
	o.Type = &v
}

func (o DeprecatedEntitySchemaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedEntitySchemaAllOf struct {
	value *DeprecatedEntitySchemaAllOf
	isSet bool
}

func (v NullableDeprecatedEntitySchemaAllOf) Get() *DeprecatedEntitySchemaAllOf {
	return v.value
}

func (v *NullableDeprecatedEntitySchemaAllOf) Set(val *DeprecatedEntitySchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedEntitySchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedEntitySchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedEntitySchemaAllOf(val *DeprecatedEntitySchemaAllOf) *NullableDeprecatedEntitySchemaAllOf {
	return &NullableDeprecatedEntitySchemaAllOf{value: val, isSet: true}
}

func (v NullableDeprecatedEntitySchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedEntitySchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


