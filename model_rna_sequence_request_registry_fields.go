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

// RnaSequenceRequestRegistryFields struct for RnaSequenceRequestRegistryFields
type RnaSequenceRequestRegistryFields struct {
	EntityRegistryId *string `json:"entityRegistryId,omitempty"`
}

// NewRnaSequenceRequestRegistryFields instantiates a new RnaSequenceRequestRegistryFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaSequenceRequestRegistryFields() *RnaSequenceRequestRegistryFields {
	this := RnaSequenceRequestRegistryFields{}
	return &this
}

// NewRnaSequenceRequestRegistryFieldsWithDefaults instantiates a new RnaSequenceRequestRegistryFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaSequenceRequestRegistryFieldsWithDefaults() *RnaSequenceRequestRegistryFields {
	this := RnaSequenceRequestRegistryFields{}
	return &this
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise.
func (o *RnaSequenceRequestRegistryFields) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSequenceRequestRegistryFields) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityRegistryId) {
    return nil, false
	}
	return o.EntityRegistryId, true
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *RnaSequenceRequestRegistryFields) HasEntityRegistryId() bool {
	if o != nil && !isNil(o.EntityRegistryId) {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given string and assigns it to the EntityRegistryId field.
func (o *RnaSequenceRequestRegistryFields) SetEntityRegistryId(v string) {
	o.EntityRegistryId = &v
}

func (o RnaSequenceRequestRegistryFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EntityRegistryId) {
		toSerialize["entityRegistryId"] = o.EntityRegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableRnaSequenceRequestRegistryFields struct {
	value *RnaSequenceRequestRegistryFields
	isSet bool
}

func (v NullableRnaSequenceRequestRegistryFields) Get() *RnaSequenceRequestRegistryFields {
	return v.value
}

func (v *NullableRnaSequenceRequestRegistryFields) Set(val *RnaSequenceRequestRegistryFields) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaSequenceRequestRegistryFields) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaSequenceRequestRegistryFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaSequenceRequestRegistryFields(val *RnaSequenceRequestRegistryFields) *NullableRnaSequenceRequestRegistryFields {
	return &NullableRnaSequenceRequestRegistryFields{value: val, isSet: true}
}

func (v NullableRnaSequenceRequestRegistryFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaSequenceRequestRegistryFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


