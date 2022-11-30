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

// DnaAlignmentsPaginatedList struct for DnaAlignmentsPaginatedList
type DnaAlignmentsPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	DnaAlignments []DnaAlignmentSummary `json:"dnaAlignments,omitempty"`
}

// NewDnaAlignmentsPaginatedList instantiates a new DnaAlignmentsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaAlignmentsPaginatedList() *DnaAlignmentsPaginatedList {
	this := DnaAlignmentsPaginatedList{}
	return &this
}

// NewDnaAlignmentsPaginatedListWithDefaults instantiates a new DnaAlignmentsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaAlignmentsPaginatedListWithDefaults() *DnaAlignmentsPaginatedList {
	this := DnaAlignmentsPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *DnaAlignmentsPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignmentsPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *DnaAlignmentsPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *DnaAlignmentsPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetDnaAlignments returns the DnaAlignments field value if set, zero value otherwise.
func (o *DnaAlignmentsPaginatedList) GetDnaAlignments() []DnaAlignmentSummary {
	if o == nil || isNil(o.DnaAlignments) {
		var ret []DnaAlignmentSummary
		return ret
	}
	return o.DnaAlignments
}

// GetDnaAlignmentsOk returns a tuple with the DnaAlignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignmentsPaginatedList) GetDnaAlignmentsOk() ([]DnaAlignmentSummary, bool) {
	if o == nil || isNil(o.DnaAlignments) {
    return nil, false
	}
	return o.DnaAlignments, true
}

// HasDnaAlignments returns a boolean if a field has been set.
func (o *DnaAlignmentsPaginatedList) HasDnaAlignments() bool {
	if o != nil && !isNil(o.DnaAlignments) {
		return true
	}

	return false
}

// SetDnaAlignments gets a reference to the given []DnaAlignmentSummary and assigns it to the DnaAlignments field.
func (o *DnaAlignmentsPaginatedList) SetDnaAlignments(v []DnaAlignmentSummary) {
	o.DnaAlignments = v
}

func (o DnaAlignmentsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.DnaAlignments) {
		toSerialize["dnaAlignments"] = o.DnaAlignments
	}
	return json.Marshal(toSerialize)
}

type NullableDnaAlignmentsPaginatedList struct {
	value *DnaAlignmentsPaginatedList
	isSet bool
}

func (v NullableDnaAlignmentsPaginatedList) Get() *DnaAlignmentsPaginatedList {
	return v.value
}

func (v *NullableDnaAlignmentsPaginatedList) Set(val *DnaAlignmentsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaAlignmentsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaAlignmentsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaAlignmentsPaginatedList(val *DnaAlignmentsPaginatedList) *NullableDnaAlignmentsPaginatedList {
	return &NullableDnaAlignmentsPaginatedList{value: val, isSet: true}
}

func (v NullableDnaAlignmentsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaAlignmentsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

