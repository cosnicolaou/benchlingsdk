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

// RnaSequencesUnarchive The request body for unarchiving RNA sequences. 
type RnaSequencesUnarchive struct {
	RnaSequenceIds []string `json:"rnaSequenceIds"`
}

// NewRnaSequencesUnarchive instantiates a new RnaSequencesUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaSequencesUnarchive(rnaSequenceIds []string) *RnaSequencesUnarchive {
	this := RnaSequencesUnarchive{}
	this.RnaSequenceIds = rnaSequenceIds
	return &this
}

// NewRnaSequencesUnarchiveWithDefaults instantiates a new RnaSequencesUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaSequencesUnarchiveWithDefaults() *RnaSequencesUnarchive {
	this := RnaSequencesUnarchive{}
	return &this
}

// GetRnaSequenceIds returns the RnaSequenceIds field value
func (o *RnaSequencesUnarchive) GetRnaSequenceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RnaSequenceIds
}

// GetRnaSequenceIdsOk returns a tuple with the RnaSequenceIds field value
// and a boolean to check if the value has been set.
func (o *RnaSequencesUnarchive) GetRnaSequenceIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RnaSequenceIds, true
}

// SetRnaSequenceIds sets field value
func (o *RnaSequencesUnarchive) SetRnaSequenceIds(v []string) {
	o.RnaSequenceIds = v
}

func (o RnaSequencesUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rnaSequenceIds"] = o.RnaSequenceIds
	}
	return json.Marshal(toSerialize)
}

type NullableRnaSequencesUnarchive struct {
	value *RnaSequencesUnarchive
	isSet bool
}

func (v NullableRnaSequencesUnarchive) Get() *RnaSequencesUnarchive {
	return v.value
}

func (v *NullableRnaSequencesUnarchive) Set(val *RnaSequencesUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaSequencesUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaSequencesUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaSequencesUnarchive(val *RnaSequencesUnarchive) *NullableRnaSequencesUnarchive {
	return &NullableRnaSequencesUnarchive{value: val, isSet: true}
}

func (v NullableRnaSequencesUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaSequencesUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


