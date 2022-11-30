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

// SampleGroupStatusUpdate struct for SampleGroupStatusUpdate
type SampleGroupStatusUpdate struct {
	// The string id of the sample group
	SampleGroupId string `json:"sampleGroupId"`
	Status SampleGroupStatus `json:"status"`
}

// NewSampleGroupStatusUpdate instantiates a new SampleGroupStatusUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleGroupStatusUpdate(sampleGroupId string, status SampleGroupStatus) *SampleGroupStatusUpdate {
	this := SampleGroupStatusUpdate{}
	this.SampleGroupId = sampleGroupId
	this.Status = status
	return &this
}

// NewSampleGroupStatusUpdateWithDefaults instantiates a new SampleGroupStatusUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleGroupStatusUpdateWithDefaults() *SampleGroupStatusUpdate {
	this := SampleGroupStatusUpdate{}
	return &this
}

// GetSampleGroupId returns the SampleGroupId field value
func (o *SampleGroupStatusUpdate) GetSampleGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SampleGroupId
}

// GetSampleGroupIdOk returns a tuple with the SampleGroupId field value
// and a boolean to check if the value has been set.
func (o *SampleGroupStatusUpdate) GetSampleGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SampleGroupId, true
}

// SetSampleGroupId sets field value
func (o *SampleGroupStatusUpdate) SetSampleGroupId(v string) {
	o.SampleGroupId = v
}

// GetStatus returns the Status field value
func (o *SampleGroupStatusUpdate) GetStatus() SampleGroupStatus {
	if o == nil {
		var ret SampleGroupStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SampleGroupStatusUpdate) GetStatusOk() (*SampleGroupStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SampleGroupStatusUpdate) SetStatus(v SampleGroupStatus) {
	o.Status = v
}

func (o SampleGroupStatusUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sampleGroupId"] = o.SampleGroupId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSampleGroupStatusUpdate struct {
	value *SampleGroupStatusUpdate
	isSet bool
}

func (v NullableSampleGroupStatusUpdate) Get() *SampleGroupStatusUpdate {
	return v.value
}

func (v *NullableSampleGroupStatusUpdate) Set(val *SampleGroupStatusUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleGroupStatusUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleGroupStatusUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleGroupStatusUpdate(val *SampleGroupStatusUpdate) *NullableSampleGroupStatusUpdate {
	return &NullableSampleGroupStatusUpdate{value: val, isSet: true}
}

func (v NullableSampleGroupStatusUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleGroupStatusUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


