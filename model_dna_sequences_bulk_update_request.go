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

// DnaSequencesBulkUpdateRequest struct for DnaSequencesBulkUpdateRequest
type DnaSequencesBulkUpdateRequest struct {
	DnaSequences []DnaSequenceBulkUpdate `json:"dnaSequences,omitempty"`
}

// NewDnaSequencesBulkUpdateRequest instantiates a new DnaSequencesBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaSequencesBulkUpdateRequest() *DnaSequencesBulkUpdateRequest {
	this := DnaSequencesBulkUpdateRequest{}
	return &this
}

// NewDnaSequencesBulkUpdateRequestWithDefaults instantiates a new DnaSequencesBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaSequencesBulkUpdateRequestWithDefaults() *DnaSequencesBulkUpdateRequest {
	this := DnaSequencesBulkUpdateRequest{}
	return &this
}

// GetDnaSequences returns the DnaSequences field value if set, zero value otherwise.
func (o *DnaSequencesBulkUpdateRequest) GetDnaSequences() []DnaSequenceBulkUpdate {
	if o == nil || isNil(o.DnaSequences) {
		var ret []DnaSequenceBulkUpdate
		return ret
	}
	return o.DnaSequences
}

// GetDnaSequencesOk returns a tuple with the DnaSequences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequencesBulkUpdateRequest) GetDnaSequencesOk() ([]DnaSequenceBulkUpdate, bool) {
	if o == nil || isNil(o.DnaSequences) {
    return nil, false
	}
	return o.DnaSequences, true
}

// HasDnaSequences returns a boolean if a field has been set.
func (o *DnaSequencesBulkUpdateRequest) HasDnaSequences() bool {
	if o != nil && !isNil(o.DnaSequences) {
		return true
	}

	return false
}

// SetDnaSequences gets a reference to the given []DnaSequenceBulkUpdate and assigns it to the DnaSequences field.
func (o *DnaSequencesBulkUpdateRequest) SetDnaSequences(v []DnaSequenceBulkUpdate) {
	o.DnaSequences = v
}

func (o DnaSequencesBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnaSequences) {
		toSerialize["dnaSequences"] = o.DnaSequences
	}
	return json.Marshal(toSerialize)
}

type NullableDnaSequencesBulkUpdateRequest struct {
	value *DnaSequencesBulkUpdateRequest
	isSet bool
}

func (v NullableDnaSequencesBulkUpdateRequest) Get() *DnaSequencesBulkUpdateRequest {
	return v.value
}

func (v *NullableDnaSequencesBulkUpdateRequest) Set(val *DnaSequencesBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaSequencesBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaSequencesBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaSequencesBulkUpdateRequest(val *DnaSequencesBulkUpdateRequest) *NullableDnaSequencesBulkUpdateRequest {
	return &NullableDnaSequencesBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableDnaSequencesBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaSequencesBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

