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

// DropdownOptionCreate struct for DropdownOptionCreate
type DropdownOptionCreate struct {
	Name string `json:"name"`
}

// NewDropdownOptionCreate instantiates a new DropdownOptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropdownOptionCreate(name string) *DropdownOptionCreate {
	this := DropdownOptionCreate{}
	this.Name = name
	return &this
}

// NewDropdownOptionCreateWithDefaults instantiates a new DropdownOptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropdownOptionCreateWithDefaults() *DropdownOptionCreate {
	this := DropdownOptionCreate{}
	return &this
}

// GetName returns the Name field value
func (o *DropdownOptionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DropdownOptionCreate) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DropdownOptionCreate) SetName(v string) {
	o.Name = v
}

func (o DropdownOptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDropdownOptionCreate struct {
	value *DropdownOptionCreate
	isSet bool
}

func (v NullableDropdownOptionCreate) Get() *DropdownOptionCreate {
	return v.value
}

func (v *NullableDropdownOptionCreate) Set(val *DropdownOptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDropdownOptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDropdownOptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropdownOptionCreate(val *DropdownOptionCreate) *NullableDropdownOptionCreate {
	return &NullableDropdownOptionCreate{value: val, isSet: true}
}

func (v NullableDropdownOptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropdownOptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


