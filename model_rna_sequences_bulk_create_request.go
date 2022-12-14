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

// RnaSequencesBulkCreateRequest struct for RnaSequencesBulkCreateRequest
type RnaSequencesBulkCreateRequest struct {
	RnaSequences []RnaSequenceBulkCreate `json:"rnaSequences,omitempty"`
}

// NewRnaSequencesBulkCreateRequest instantiates a new RnaSequencesBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaSequencesBulkCreateRequest() *RnaSequencesBulkCreateRequest {
	this := RnaSequencesBulkCreateRequest{}
	return &this
}

// NewRnaSequencesBulkCreateRequestWithDefaults instantiates a new RnaSequencesBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaSequencesBulkCreateRequestWithDefaults() *RnaSequencesBulkCreateRequest {
	this := RnaSequencesBulkCreateRequest{}
	return &this
}

// GetRnaSequences returns the RnaSequences field value if set, zero value otherwise.
func (o *RnaSequencesBulkCreateRequest) GetRnaSequences() []RnaSequenceBulkCreate {
	if o == nil || isNil(o.RnaSequences) {
		var ret []RnaSequenceBulkCreate
		return ret
	}
	return o.RnaSequences
}

// GetRnaSequencesOk returns a tuple with the RnaSequences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaSequencesBulkCreateRequest) GetRnaSequencesOk() ([]RnaSequenceBulkCreate, bool) {
	if o == nil || isNil(o.RnaSequences) {
    return nil, false
	}
	return o.RnaSequences, true
}

// HasRnaSequences returns a boolean if a field has been set.
func (o *RnaSequencesBulkCreateRequest) HasRnaSequences() bool {
	if o != nil && !isNil(o.RnaSequences) {
		return true
	}

	return false
}

// SetRnaSequences gets a reference to the given []RnaSequenceBulkCreate and assigns it to the RnaSequences field.
func (o *RnaSequencesBulkCreateRequest) SetRnaSequences(v []RnaSequenceBulkCreate) {
	o.RnaSequences = v
}

func (o RnaSequencesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RnaSequences) {
		toSerialize["rnaSequences"] = o.RnaSequences
	}
	return json.Marshal(toSerialize)
}

type NullableRnaSequencesBulkCreateRequest struct {
	value *RnaSequencesBulkCreateRequest
	isSet bool
}

func (v NullableRnaSequencesBulkCreateRequest) Get() *RnaSequencesBulkCreateRequest {
	return v.value
}

func (v *NullableRnaSequencesBulkCreateRequest) Set(val *RnaSequencesBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaSequencesBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaSequencesBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaSequencesBulkCreateRequest(val *RnaSequencesBulkCreateRequest) *NullableRnaSequencesBulkCreateRequest {
	return &NullableRnaSequencesBulkCreateRequest{value: val, isSet: true}
}

func (v NullableRnaSequencesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaSequencesBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


