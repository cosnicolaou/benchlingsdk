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

// RnaOligosUnarchive The request body for unarchiving RNA Oligos. 
type RnaOligosUnarchive struct {
	RnaOligoIds []string `json:"rnaOligoIds"`
}

// NewRnaOligosUnarchive instantiates a new RnaOligosUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaOligosUnarchive(rnaOligoIds []string) *RnaOligosUnarchive {
	this := RnaOligosUnarchive{}
	this.RnaOligoIds = rnaOligoIds
	return &this
}

// NewRnaOligosUnarchiveWithDefaults instantiates a new RnaOligosUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaOligosUnarchiveWithDefaults() *RnaOligosUnarchive {
	this := RnaOligosUnarchive{}
	return &this
}

// GetRnaOligoIds returns the RnaOligoIds field value
func (o *RnaOligosUnarchive) GetRnaOligoIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RnaOligoIds
}

// GetRnaOligoIdsOk returns a tuple with the RnaOligoIds field value
// and a boolean to check if the value has been set.
func (o *RnaOligosUnarchive) GetRnaOligoIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RnaOligoIds, true
}

// SetRnaOligoIds sets field value
func (o *RnaOligosUnarchive) SetRnaOligoIds(v []string) {
	o.RnaOligoIds = v
}

func (o RnaOligosUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rnaOligoIds"] = o.RnaOligoIds
	}
	return json.Marshal(toSerialize)
}

type NullableRnaOligosUnarchive struct {
	value *RnaOligosUnarchive
	isSet bool
}

func (v NullableRnaOligosUnarchive) Get() *RnaOligosUnarchive {
	return v.value
}

func (v *NullableRnaOligosUnarchive) Set(val *RnaOligosUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaOligosUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaOligosUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaOligosUnarchive(val *RnaOligosUnarchive) *NullableRnaOligosUnarchive {
	return &NullableRnaOligosUnarchive{value: val, isSet: true}
}

func (v NullableRnaOligosUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaOligosUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


