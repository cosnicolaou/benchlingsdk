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

// DnaOligosPaginatedList struct for DnaOligosPaginatedList
type DnaOligosPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	DnaOligos []DnaOligo `json:"dnaOligos,omitempty"`
}

// NewDnaOligosPaginatedList instantiates a new DnaOligosPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaOligosPaginatedList() *DnaOligosPaginatedList {
	this := DnaOligosPaginatedList{}
	return &this
}

// NewDnaOligosPaginatedListWithDefaults instantiates a new DnaOligosPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaOligosPaginatedListWithDefaults() *DnaOligosPaginatedList {
	this := DnaOligosPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *DnaOligosPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaOligosPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *DnaOligosPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *DnaOligosPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetDnaOligos returns the DnaOligos field value if set, zero value otherwise.
func (o *DnaOligosPaginatedList) GetDnaOligos() []DnaOligo {
	if o == nil || isNil(o.DnaOligos) {
		var ret []DnaOligo
		return ret
	}
	return o.DnaOligos
}

// GetDnaOligosOk returns a tuple with the DnaOligos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaOligosPaginatedList) GetDnaOligosOk() ([]DnaOligo, bool) {
	if o == nil || isNil(o.DnaOligos) {
    return nil, false
	}
	return o.DnaOligos, true
}

// HasDnaOligos returns a boolean if a field has been set.
func (o *DnaOligosPaginatedList) HasDnaOligos() bool {
	if o != nil && !isNil(o.DnaOligos) {
		return true
	}

	return false
}

// SetDnaOligos gets a reference to the given []DnaOligo and assigns it to the DnaOligos field.
func (o *DnaOligosPaginatedList) SetDnaOligos(v []DnaOligo) {
	o.DnaOligos = v
}

func (o DnaOligosPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.DnaOligos) {
		toSerialize["dnaOligos"] = o.DnaOligos
	}
	return json.Marshal(toSerialize)
}

type NullableDnaOligosPaginatedList struct {
	value *DnaOligosPaginatedList
	isSet bool
}

func (v NullableDnaOligosPaginatedList) Get() *DnaOligosPaginatedList {
	return v.value
}

func (v *NullableDnaOligosPaginatedList) Set(val *DnaOligosPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaOligosPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaOligosPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaOligosPaginatedList(val *DnaOligosPaginatedList) *NullableDnaOligosPaginatedList {
	return &NullableDnaOligosPaginatedList{value: val, isSet: true}
}

func (v NullableDnaOligosPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaOligosPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


