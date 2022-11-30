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

// DnaSequencesArchivalChange IDs of all items that were archived or unarchived, grouped by resource type. This includes the IDs of DNA sequences along with any IDs of batches that were archived / unarchived. 
type DnaSequencesArchivalChange struct {
	BatchIds []string `json:"batchIds,omitempty"`
	DnaSequenceIds []string `json:"dnaSequenceIds,omitempty"`
}

// NewDnaSequencesArchivalChange instantiates a new DnaSequencesArchivalChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaSequencesArchivalChange() *DnaSequencesArchivalChange {
	this := DnaSequencesArchivalChange{}
	return &this
}

// NewDnaSequencesArchivalChangeWithDefaults instantiates a new DnaSequencesArchivalChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaSequencesArchivalChangeWithDefaults() *DnaSequencesArchivalChange {
	this := DnaSequencesArchivalChange{}
	return &this
}

// GetBatchIds returns the BatchIds field value if set, zero value otherwise.
func (o *DnaSequencesArchivalChange) GetBatchIds() []string {
	if o == nil || isNil(o.BatchIds) {
		var ret []string
		return ret
	}
	return o.BatchIds
}

// GetBatchIdsOk returns a tuple with the BatchIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequencesArchivalChange) GetBatchIdsOk() ([]string, bool) {
	if o == nil || isNil(o.BatchIds) {
    return nil, false
	}
	return o.BatchIds, true
}

// HasBatchIds returns a boolean if a field has been set.
func (o *DnaSequencesArchivalChange) HasBatchIds() bool {
	if o != nil && !isNil(o.BatchIds) {
		return true
	}

	return false
}

// SetBatchIds gets a reference to the given []string and assigns it to the BatchIds field.
func (o *DnaSequencesArchivalChange) SetBatchIds(v []string) {
	o.BatchIds = v
}

// GetDnaSequenceIds returns the DnaSequenceIds field value if set, zero value otherwise.
func (o *DnaSequencesArchivalChange) GetDnaSequenceIds() []string {
	if o == nil || isNil(o.DnaSequenceIds) {
		var ret []string
		return ret
	}
	return o.DnaSequenceIds
}

// GetDnaSequenceIdsOk returns a tuple with the DnaSequenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequencesArchivalChange) GetDnaSequenceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.DnaSequenceIds) {
    return nil, false
	}
	return o.DnaSequenceIds, true
}

// HasDnaSequenceIds returns a boolean if a field has been set.
func (o *DnaSequencesArchivalChange) HasDnaSequenceIds() bool {
	if o != nil && !isNil(o.DnaSequenceIds) {
		return true
	}

	return false
}

// SetDnaSequenceIds gets a reference to the given []string and assigns it to the DnaSequenceIds field.
func (o *DnaSequencesArchivalChange) SetDnaSequenceIds(v []string) {
	o.DnaSequenceIds = v
}

func (o DnaSequencesArchivalChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BatchIds) {
		toSerialize["batchIds"] = o.BatchIds
	}
	if !isNil(o.DnaSequenceIds) {
		toSerialize["dnaSequenceIds"] = o.DnaSequenceIds
	}
	return json.Marshal(toSerialize)
}

type NullableDnaSequencesArchivalChange struct {
	value *DnaSequencesArchivalChange
	isSet bool
}

func (v NullableDnaSequencesArchivalChange) Get() *DnaSequencesArchivalChange {
	return v.value
}

func (v *NullableDnaSequencesArchivalChange) Set(val *DnaSequencesArchivalChange) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaSequencesArchivalChange) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaSequencesArchivalChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaSequencesArchivalChange(val *DnaSequencesArchivalChange) *NullableDnaSequencesArchivalChange {
	return &NullableDnaSequencesArchivalChange{value: val, isSet: true}
}

func (v NullableDnaSequencesArchivalChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaSequencesArchivalChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

