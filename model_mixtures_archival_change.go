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

// MixturesArchivalChange IDs of all mixtures that were archived or unarchived. 
type MixturesArchivalChange struct {
	MixtureIds []string `json:"mixtureIds,omitempty"`
}

// NewMixturesArchivalChange instantiates a new MixturesArchivalChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixturesArchivalChange() *MixturesArchivalChange {
	this := MixturesArchivalChange{}
	return &this
}

// NewMixturesArchivalChangeWithDefaults instantiates a new MixturesArchivalChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixturesArchivalChangeWithDefaults() *MixturesArchivalChange {
	this := MixturesArchivalChange{}
	return &this
}

// GetMixtureIds returns the MixtureIds field value if set, zero value otherwise.
func (o *MixturesArchivalChange) GetMixtureIds() []string {
	if o == nil || isNil(o.MixtureIds) {
		var ret []string
		return ret
	}
	return o.MixtureIds
}

// GetMixtureIdsOk returns a tuple with the MixtureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixturesArchivalChange) GetMixtureIdsOk() ([]string, bool) {
	if o == nil || isNil(o.MixtureIds) {
    return nil, false
	}
	return o.MixtureIds, true
}

// HasMixtureIds returns a boolean if a field has been set.
func (o *MixturesArchivalChange) HasMixtureIds() bool {
	if o != nil && !isNil(o.MixtureIds) {
		return true
	}

	return false
}

// SetMixtureIds gets a reference to the given []string and assigns it to the MixtureIds field.
func (o *MixturesArchivalChange) SetMixtureIds(v []string) {
	o.MixtureIds = v
}

func (o MixturesArchivalChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MixtureIds) {
		toSerialize["mixtureIds"] = o.MixtureIds
	}
	return json.Marshal(toSerialize)
}

type NullableMixturesArchivalChange struct {
	value *MixturesArchivalChange
	isSet bool
}

func (v NullableMixturesArchivalChange) Get() *MixturesArchivalChange {
	return v.value
}

func (v *NullableMixturesArchivalChange) Set(val *MixturesArchivalChange) {
	v.value = val
	v.isSet = true
}

func (v NullableMixturesArchivalChange) IsSet() bool {
	return v.isSet
}

func (v *NullableMixturesArchivalChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixturesArchivalChange(val *MixturesArchivalChange) *NullableMixturesArchivalChange {
	return &NullableMixturesArchivalChange{value: val, isSet: true}
}

func (v NullableMixturesArchivalChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixturesArchivalChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


