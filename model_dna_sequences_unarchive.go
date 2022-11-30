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

// DnaSequencesUnarchive The request body for unarchiving DNA sequences. 
type DnaSequencesUnarchive struct {
	DnaSequenceIds []string `json:"dnaSequenceIds"`
}

// NewDnaSequencesUnarchive instantiates a new DnaSequencesUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaSequencesUnarchive(dnaSequenceIds []string) *DnaSequencesUnarchive {
	this := DnaSequencesUnarchive{}
	this.DnaSequenceIds = dnaSequenceIds
	return &this
}

// NewDnaSequencesUnarchiveWithDefaults instantiates a new DnaSequencesUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaSequencesUnarchiveWithDefaults() *DnaSequencesUnarchive {
	this := DnaSequencesUnarchive{}
	return &this
}

// GetDnaSequenceIds returns the DnaSequenceIds field value
func (o *DnaSequencesUnarchive) GetDnaSequenceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnaSequenceIds
}

// GetDnaSequenceIdsOk returns a tuple with the DnaSequenceIds field value
// and a boolean to check if the value has been set.
func (o *DnaSequencesUnarchive) GetDnaSequenceIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DnaSequenceIds, true
}

// SetDnaSequenceIds sets field value
func (o *DnaSequencesUnarchive) SetDnaSequenceIds(v []string) {
	o.DnaSequenceIds = v
}

func (o DnaSequencesUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnaSequenceIds"] = o.DnaSequenceIds
	}
	return json.Marshal(toSerialize)
}

type NullableDnaSequencesUnarchive struct {
	value *DnaSequencesUnarchive
	isSet bool
}

func (v NullableDnaSequencesUnarchive) Get() *DnaSequencesUnarchive {
	return v.value
}

func (v *NullableDnaSequencesUnarchive) Set(val *DnaSequencesUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaSequencesUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaSequencesUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaSequencesUnarchive(val *DnaSequencesUnarchive) *NullableDnaSequencesUnarchive {
	return &NullableDnaSequencesUnarchive{value: val, isSet: true}
}

func (v NullableDnaSequencesUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaSequencesUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

