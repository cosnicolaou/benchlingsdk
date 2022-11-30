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

// DnaOligosArchive The request body for archiving DNA Oligos. 
type DnaOligosArchive struct {
	DnaOligoIds []string `json:"dnaOligoIds"`
	Reason EntityArchiveReason `json:"reason"`
}

// NewDnaOligosArchive instantiates a new DnaOligosArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaOligosArchive(dnaOligoIds []string, reason EntityArchiveReason) *DnaOligosArchive {
	this := DnaOligosArchive{}
	this.DnaOligoIds = dnaOligoIds
	this.Reason = reason
	return &this
}

// NewDnaOligosArchiveWithDefaults instantiates a new DnaOligosArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaOligosArchiveWithDefaults() *DnaOligosArchive {
	this := DnaOligosArchive{}
	return &this
}

// GetDnaOligoIds returns the DnaOligoIds field value
func (o *DnaOligosArchive) GetDnaOligoIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnaOligoIds
}

// GetDnaOligoIdsOk returns a tuple with the DnaOligoIds field value
// and a boolean to check if the value has been set.
func (o *DnaOligosArchive) GetDnaOligoIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DnaOligoIds, true
}

// SetDnaOligoIds sets field value
func (o *DnaOligosArchive) SetDnaOligoIds(v []string) {
	o.DnaOligoIds = v
}

// GetReason returns the Reason field value
func (o *DnaOligosArchive) GetReason() EntityArchiveReason {
	if o == nil {
		var ret EntityArchiveReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DnaOligosArchive) GetReasonOk() (*EntityArchiveReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *DnaOligosArchive) SetReason(v EntityArchiveReason) {
	o.Reason = v
}

func (o DnaOligosArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnaOligoIds"] = o.DnaOligoIds
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableDnaOligosArchive struct {
	value *DnaOligosArchive
	isSet bool
}

func (v NullableDnaOligosArchive) Get() *DnaOligosArchive {
	return v.value
}

func (v *NullableDnaOligosArchive) Set(val *DnaOligosArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaOligosArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaOligosArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaOligosArchive(val *DnaOligosArchive) *NullableDnaOligosArchive {
	return &NullableDnaOligosArchive{value: val, isSet: true}
}

func (v NullableDnaOligosArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaOligosArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


