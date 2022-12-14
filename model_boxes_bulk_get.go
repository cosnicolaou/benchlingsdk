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

// BoxesBulkGet struct for BoxesBulkGet
type BoxesBulkGet struct {
	Boxes []Box `json:"boxes,omitempty"`
}

// NewBoxesBulkGet instantiates a new BoxesBulkGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoxesBulkGet() *BoxesBulkGet {
	this := BoxesBulkGet{}
	return &this
}

// NewBoxesBulkGetWithDefaults instantiates a new BoxesBulkGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxesBulkGetWithDefaults() *BoxesBulkGet {
	this := BoxesBulkGet{}
	return &this
}

// GetBoxes returns the Boxes field value if set, zero value otherwise.
func (o *BoxesBulkGet) GetBoxes() []Box {
	if o == nil || isNil(o.Boxes) {
		var ret []Box
		return ret
	}
	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxesBulkGet) GetBoxesOk() ([]Box, bool) {
	if o == nil || isNil(o.Boxes) {
    return nil, false
	}
	return o.Boxes, true
}

// HasBoxes returns a boolean if a field has been set.
func (o *BoxesBulkGet) HasBoxes() bool {
	if o != nil && !isNil(o.Boxes) {
		return true
	}

	return false
}

// SetBoxes gets a reference to the given []Box and assigns it to the Boxes field.
func (o *BoxesBulkGet) SetBoxes(v []Box) {
	o.Boxes = v
}

func (o BoxesBulkGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Boxes) {
		toSerialize["boxes"] = o.Boxes
	}
	return json.Marshal(toSerialize)
}

type NullableBoxesBulkGet struct {
	value *BoxesBulkGet
	isSet bool
}

func (v NullableBoxesBulkGet) Get() *BoxesBulkGet {
	return v.value
}

func (v *NullableBoxesBulkGet) Set(val *BoxesBulkGet) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxesBulkGet) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxesBulkGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxesBulkGet(val *BoxesBulkGet) *NullableBoxesBulkGet {
	return &NullableBoxesBulkGet{value: val, isSet: true}
}

func (v NullableBoxesBulkGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoxesBulkGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


