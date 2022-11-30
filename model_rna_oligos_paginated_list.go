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

// RnaOligosPaginatedList struct for RnaOligosPaginatedList
type RnaOligosPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	RnaOligos []RnaOligo `json:"rnaOligos,omitempty"`
}

// NewRnaOligosPaginatedList instantiates a new RnaOligosPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaOligosPaginatedList() *RnaOligosPaginatedList {
	this := RnaOligosPaginatedList{}
	return &this
}

// NewRnaOligosPaginatedListWithDefaults instantiates a new RnaOligosPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaOligosPaginatedListWithDefaults() *RnaOligosPaginatedList {
	this := RnaOligosPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *RnaOligosPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligosPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *RnaOligosPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *RnaOligosPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetRnaOligos returns the RnaOligos field value if set, zero value otherwise.
func (o *RnaOligosPaginatedList) GetRnaOligos() []RnaOligo {
	if o == nil || isNil(o.RnaOligos) {
		var ret []RnaOligo
		return ret
	}
	return o.RnaOligos
}

// GetRnaOligosOk returns a tuple with the RnaOligos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligosPaginatedList) GetRnaOligosOk() ([]RnaOligo, bool) {
	if o == nil || isNil(o.RnaOligos) {
    return nil, false
	}
	return o.RnaOligos, true
}

// HasRnaOligos returns a boolean if a field has been set.
func (o *RnaOligosPaginatedList) HasRnaOligos() bool {
	if o != nil && !isNil(o.RnaOligos) {
		return true
	}

	return false
}

// SetRnaOligos gets a reference to the given []RnaOligo and assigns it to the RnaOligos field.
func (o *RnaOligosPaginatedList) SetRnaOligos(v []RnaOligo) {
	o.RnaOligos = v
}

func (o RnaOligosPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.RnaOligos) {
		toSerialize["rnaOligos"] = o.RnaOligos
	}
	return json.Marshal(toSerialize)
}

type NullableRnaOligosPaginatedList struct {
	value *RnaOligosPaginatedList
	isSet bool
}

func (v NullableRnaOligosPaginatedList) Get() *RnaOligosPaginatedList {
	return v.value
}

func (v *NullableRnaOligosPaginatedList) Set(val *RnaOligosPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaOligosPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaOligosPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaOligosPaginatedList(val *RnaOligosPaginatedList) *NullableRnaOligosPaginatedList {
	return &NullableRnaOligosPaginatedList{value: val, isSet: true}
}

func (v NullableRnaOligosPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaOligosPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

