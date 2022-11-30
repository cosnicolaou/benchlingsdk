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

// WorkflowSampleList struct for WorkflowSampleList
type WorkflowSampleList struct {
	Samples []WorkflowSample `json:"samples,omitempty"`
}

// NewWorkflowSampleList instantiates a new WorkflowSampleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSampleList() *WorkflowSampleList {
	this := WorkflowSampleList{}
	return &this
}

// NewWorkflowSampleListWithDefaults instantiates a new WorkflowSampleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSampleListWithDefaults() *WorkflowSampleList {
	this := WorkflowSampleList{}
	return &this
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *WorkflowSampleList) GetSamples() []WorkflowSample {
	if o == nil || isNil(o.Samples) {
		var ret []WorkflowSample
		return ret
	}
	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSampleList) GetSamplesOk() ([]WorkflowSample, bool) {
	if o == nil || isNil(o.Samples) {
    return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *WorkflowSampleList) HasSamples() bool {
	if o != nil && !isNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given []WorkflowSample and assigns it to the Samples field.
func (o *WorkflowSampleList) SetSamples(v []WorkflowSample) {
	o.Samples = v
}

func (o WorkflowSampleList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSampleList struct {
	value *WorkflowSampleList
	isSet bool
}

func (v NullableWorkflowSampleList) Get() *WorkflowSampleList {
	return v.value
}

func (v *NullableWorkflowSampleList) Set(val *WorkflowSampleList) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSampleList) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSampleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSampleList(val *WorkflowSampleList) *NullableWorkflowSampleList {
	return &NullableWorkflowSampleList{value: val, isSet: true}
}

func (v NullableWorkflowSampleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSampleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


