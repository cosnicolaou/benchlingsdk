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

// RnaOligosArchive The request body for archiving RNA Oligos. 
type RnaOligosArchive struct {
	Reason EntityArchiveReason `json:"reason"`
	RnaOligoIds []string `json:"rnaOligoIds"`
}

// NewRnaOligosArchive instantiates a new RnaOligosArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaOligosArchive(reason EntityArchiveReason, rnaOligoIds []string) *RnaOligosArchive {
	this := RnaOligosArchive{}
	this.Reason = reason
	this.RnaOligoIds = rnaOligoIds
	return &this
}

// NewRnaOligosArchiveWithDefaults instantiates a new RnaOligosArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaOligosArchiveWithDefaults() *RnaOligosArchive {
	this := RnaOligosArchive{}
	return &this
}

// GetReason returns the Reason field value
func (o *RnaOligosArchive) GetReason() EntityArchiveReason {
	if o == nil {
		var ret EntityArchiveReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RnaOligosArchive) GetReasonOk() (*EntityArchiveReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RnaOligosArchive) SetReason(v EntityArchiveReason) {
	o.Reason = v
}

// GetRnaOligoIds returns the RnaOligoIds field value
func (o *RnaOligosArchive) GetRnaOligoIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RnaOligoIds
}

// GetRnaOligoIdsOk returns a tuple with the RnaOligoIds field value
// and a boolean to check if the value has been set.
func (o *RnaOligosArchive) GetRnaOligoIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RnaOligoIds, true
}

// SetRnaOligoIds sets field value
func (o *RnaOligosArchive) SetRnaOligoIds(v []string) {
	o.RnaOligoIds = v
}

func (o RnaOligosArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["rnaOligoIds"] = o.RnaOligoIds
	}
	return json.Marshal(toSerialize)
}

type NullableRnaOligosArchive struct {
	value *RnaOligosArchive
	isSet bool
}

func (v NullableRnaOligosArchive) Get() *RnaOligosArchive {
	return v.value
}

func (v *NullableRnaOligosArchive) Set(val *RnaOligosArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaOligosArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaOligosArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaOligosArchive(val *RnaOligosArchive) *NullableRnaOligosArchive {
	return &NullableRnaOligosArchive{value: val, isSet: true}
}

func (v NullableRnaOligosArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaOligosArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


