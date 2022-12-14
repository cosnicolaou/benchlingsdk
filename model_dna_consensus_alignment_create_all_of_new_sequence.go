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

// DnaConsensusAlignmentCreateAllOfNewSequence struct for DnaConsensusAlignmentCreateAllOfNewSequence
type DnaConsensusAlignmentCreateAllOfNewSequence struct {
	FolderId *string `json:"folderId,omitempty"`
}

// NewDnaConsensusAlignmentCreateAllOfNewSequence instantiates a new DnaConsensusAlignmentCreateAllOfNewSequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaConsensusAlignmentCreateAllOfNewSequence() *DnaConsensusAlignmentCreateAllOfNewSequence {
	this := DnaConsensusAlignmentCreateAllOfNewSequence{}
	return &this
}

// NewDnaConsensusAlignmentCreateAllOfNewSequenceWithDefaults instantiates a new DnaConsensusAlignmentCreateAllOfNewSequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaConsensusAlignmentCreateAllOfNewSequenceWithDefaults() *DnaConsensusAlignmentCreateAllOfNewSequence {
	this := DnaConsensusAlignmentCreateAllOfNewSequence{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DnaConsensusAlignmentCreateAllOfNewSequence) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreateAllOfNewSequence) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DnaConsensusAlignmentCreateAllOfNewSequence) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *DnaConsensusAlignmentCreateAllOfNewSequence) SetFolderId(v string) {
	o.FolderId = &v
}

func (o DnaConsensusAlignmentCreateAllOfNewSequence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	return json.Marshal(toSerialize)
}

type NullableDnaConsensusAlignmentCreateAllOfNewSequence struct {
	value *DnaConsensusAlignmentCreateAllOfNewSequence
	isSet bool
}

func (v NullableDnaConsensusAlignmentCreateAllOfNewSequence) Get() *DnaConsensusAlignmentCreateAllOfNewSequence {
	return v.value
}

func (v *NullableDnaConsensusAlignmentCreateAllOfNewSequence) Set(val *DnaConsensusAlignmentCreateAllOfNewSequence) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaConsensusAlignmentCreateAllOfNewSequence) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaConsensusAlignmentCreateAllOfNewSequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaConsensusAlignmentCreateAllOfNewSequence(val *DnaConsensusAlignmentCreateAllOfNewSequence) *NullableDnaConsensusAlignmentCreateAllOfNewSequence {
	return &NullableDnaConsensusAlignmentCreateAllOfNewSequence{value: val, isSet: true}
}

func (v NullableDnaConsensusAlignmentCreateAllOfNewSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaConsensusAlignmentCreateAllOfNewSequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


