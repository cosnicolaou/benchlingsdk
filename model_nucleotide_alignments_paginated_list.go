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

// NucleotideAlignmentsPaginatedList struct for NucleotideAlignmentsPaginatedList
type NucleotideAlignmentsPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	Alignments []NucleotideAlignmentSummary `json:"alignments,omitempty"`
}

// NewNucleotideAlignmentsPaginatedList instantiates a new NucleotideAlignmentsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNucleotideAlignmentsPaginatedList() *NucleotideAlignmentsPaginatedList {
	this := NucleotideAlignmentsPaginatedList{}
	return &this
}

// NewNucleotideAlignmentsPaginatedListWithDefaults instantiates a new NucleotideAlignmentsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNucleotideAlignmentsPaginatedListWithDefaults() *NucleotideAlignmentsPaginatedList {
	this := NucleotideAlignmentsPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *NucleotideAlignmentsPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NucleotideAlignmentsPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *NucleotideAlignmentsPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *NucleotideAlignmentsPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAlignments returns the Alignments field value if set, zero value otherwise.
func (o *NucleotideAlignmentsPaginatedList) GetAlignments() []NucleotideAlignmentSummary {
	if o == nil || isNil(o.Alignments) {
		var ret []NucleotideAlignmentSummary
		return ret
	}
	return o.Alignments
}

// GetAlignmentsOk returns a tuple with the Alignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NucleotideAlignmentsPaginatedList) GetAlignmentsOk() ([]NucleotideAlignmentSummary, bool) {
	if o == nil || isNil(o.Alignments) {
    return nil, false
	}
	return o.Alignments, true
}

// HasAlignments returns a boolean if a field has been set.
func (o *NucleotideAlignmentsPaginatedList) HasAlignments() bool {
	if o != nil && !isNil(o.Alignments) {
		return true
	}

	return false
}

// SetAlignments gets a reference to the given []NucleotideAlignmentSummary and assigns it to the Alignments field.
func (o *NucleotideAlignmentsPaginatedList) SetAlignments(v []NucleotideAlignmentSummary) {
	o.Alignments = v
}

func (o NucleotideAlignmentsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.Alignments) {
		toSerialize["alignments"] = o.Alignments
	}
	return json.Marshal(toSerialize)
}

type NullableNucleotideAlignmentsPaginatedList struct {
	value *NucleotideAlignmentsPaginatedList
	isSet bool
}

func (v NullableNucleotideAlignmentsPaginatedList) Get() *NucleotideAlignmentsPaginatedList {
	return v.value
}

func (v *NullableNucleotideAlignmentsPaginatedList) Set(val *NucleotideAlignmentsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableNucleotideAlignmentsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableNucleotideAlignmentsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNucleotideAlignmentsPaginatedList(val *NucleotideAlignmentsPaginatedList) *NullableNucleotideAlignmentsPaginatedList {
	return &NullableNucleotideAlignmentsPaginatedList{value: val, isSet: true}
}

func (v NullableNucleotideAlignmentsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNucleotideAlignmentsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


