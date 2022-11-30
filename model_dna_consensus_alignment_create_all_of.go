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

// DnaConsensusAlignmentCreateAllOf struct for DnaConsensusAlignmentCreateAllOf
type DnaConsensusAlignmentCreateAllOf struct {
	NewSequence *DnaConsensusAlignmentCreateAllOfNewSequence `json:"newSequence,omitempty"`
	SequenceId *string `json:"sequenceId,omitempty"`
}

// NewDnaConsensusAlignmentCreateAllOf instantiates a new DnaConsensusAlignmentCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaConsensusAlignmentCreateAllOf() *DnaConsensusAlignmentCreateAllOf {
	this := DnaConsensusAlignmentCreateAllOf{}
	return &this
}

// NewDnaConsensusAlignmentCreateAllOfWithDefaults instantiates a new DnaConsensusAlignmentCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaConsensusAlignmentCreateAllOfWithDefaults() *DnaConsensusAlignmentCreateAllOf {
	this := DnaConsensusAlignmentCreateAllOf{}
	return &this
}

// GetNewSequence returns the NewSequence field value if set, zero value otherwise.
func (o *DnaConsensusAlignmentCreateAllOf) GetNewSequence() DnaConsensusAlignmentCreateAllOfNewSequence {
	if o == nil || isNil(o.NewSequence) {
		var ret DnaConsensusAlignmentCreateAllOfNewSequence
		return ret
	}
	return *o.NewSequence
}

// GetNewSequenceOk returns a tuple with the NewSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreateAllOf) GetNewSequenceOk() (*DnaConsensusAlignmentCreateAllOfNewSequence, bool) {
	if o == nil || isNil(o.NewSequence) {
    return nil, false
	}
	return o.NewSequence, true
}

// HasNewSequence returns a boolean if a field has been set.
func (o *DnaConsensusAlignmentCreateAllOf) HasNewSequence() bool {
	if o != nil && !isNil(o.NewSequence) {
		return true
	}

	return false
}

// SetNewSequence gets a reference to the given DnaConsensusAlignmentCreateAllOfNewSequence and assigns it to the NewSequence field.
func (o *DnaConsensusAlignmentCreateAllOf) SetNewSequence(v DnaConsensusAlignmentCreateAllOfNewSequence) {
	o.NewSequence = &v
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise.
func (o *DnaConsensusAlignmentCreateAllOf) GetSequenceId() string {
	if o == nil || isNil(o.SequenceId) {
		var ret string
		return ret
	}
	return *o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreateAllOf) GetSequenceIdOk() (*string, bool) {
	if o == nil || isNil(o.SequenceId) {
    return nil, false
	}
	return o.SequenceId, true
}

// HasSequenceId returns a boolean if a field has been set.
func (o *DnaConsensusAlignmentCreateAllOf) HasSequenceId() bool {
	if o != nil && !isNil(o.SequenceId) {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given string and assigns it to the SequenceId field.
func (o *DnaConsensusAlignmentCreateAllOf) SetSequenceId(v string) {
	o.SequenceId = &v
}

func (o DnaConsensusAlignmentCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NewSequence) {
		toSerialize["newSequence"] = o.NewSequence
	}
	if !isNil(o.SequenceId) {
		toSerialize["sequenceId"] = o.SequenceId
	}
	return json.Marshal(toSerialize)
}

type NullableDnaConsensusAlignmentCreateAllOf struct {
	value *DnaConsensusAlignmentCreateAllOf
	isSet bool
}

func (v NullableDnaConsensusAlignmentCreateAllOf) Get() *DnaConsensusAlignmentCreateAllOf {
	return v.value
}

func (v *NullableDnaConsensusAlignmentCreateAllOf) Set(val *DnaConsensusAlignmentCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaConsensusAlignmentCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaConsensusAlignmentCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaConsensusAlignmentCreateAllOf(val *DnaConsensusAlignmentCreateAllOf) *NullableDnaConsensusAlignmentCreateAllOf {
	return &NullableDnaConsensusAlignmentCreateAllOf{value: val, isSet: true}
}

func (v NullableDnaConsensusAlignmentCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaConsensusAlignmentCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

