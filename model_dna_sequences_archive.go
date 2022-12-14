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

// DnaSequencesArchive The request body for archiving DNA sequences. 
type DnaSequencesArchive struct {
	DnaSequenceIds []string `json:"dnaSequenceIds"`
	Reason EntityArchiveReason `json:"reason"`
}

// NewDnaSequencesArchive instantiates a new DnaSequencesArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaSequencesArchive(dnaSequenceIds []string, reason EntityArchiveReason) *DnaSequencesArchive {
	this := DnaSequencesArchive{}
	this.DnaSequenceIds = dnaSequenceIds
	this.Reason = reason
	return &this
}

// NewDnaSequencesArchiveWithDefaults instantiates a new DnaSequencesArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaSequencesArchiveWithDefaults() *DnaSequencesArchive {
	this := DnaSequencesArchive{}
	return &this
}

// GetDnaSequenceIds returns the DnaSequenceIds field value
func (o *DnaSequencesArchive) GetDnaSequenceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnaSequenceIds
}

// GetDnaSequenceIdsOk returns a tuple with the DnaSequenceIds field value
// and a boolean to check if the value has been set.
func (o *DnaSequencesArchive) GetDnaSequenceIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DnaSequenceIds, true
}

// SetDnaSequenceIds sets field value
func (o *DnaSequencesArchive) SetDnaSequenceIds(v []string) {
	o.DnaSequenceIds = v
}

// GetReason returns the Reason field value
func (o *DnaSequencesArchive) GetReason() EntityArchiveReason {
	if o == nil {
		var ret EntityArchiveReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DnaSequencesArchive) GetReasonOk() (*EntityArchiveReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *DnaSequencesArchive) SetReason(v EntityArchiveReason) {
	o.Reason = v
}

func (o DnaSequencesArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnaSequenceIds"] = o.DnaSequenceIds
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableDnaSequencesArchive struct {
	value *DnaSequencesArchive
	isSet bool
}

func (v NullableDnaSequencesArchive) Get() *DnaSequencesArchive {
	return v.value
}

func (v *NullableDnaSequencesArchive) Set(val *DnaSequencesArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaSequencesArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaSequencesArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaSequencesArchive(val *DnaSequencesArchive) *NullableDnaSequencesArchive {
	return &NullableDnaSequencesArchive{value: val, isSet: true}
}

func (v NullableDnaSequencesArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaSequencesArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


