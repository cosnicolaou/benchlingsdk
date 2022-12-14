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

// WorkflowOutputsArchive struct for WorkflowOutputsArchive
type WorkflowOutputsArchive struct {
	Reason WorkflowOutputArchiveReason `json:"reason"`
	WorkflowOutputIds []string `json:"workflowOutputIds"`
}

// NewWorkflowOutputsArchive instantiates a new WorkflowOutputsArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOutputsArchive(reason WorkflowOutputArchiveReason, workflowOutputIds []string) *WorkflowOutputsArchive {
	this := WorkflowOutputsArchive{}
	this.Reason = reason
	this.WorkflowOutputIds = workflowOutputIds
	return &this
}

// NewWorkflowOutputsArchiveWithDefaults instantiates a new WorkflowOutputsArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOutputsArchiveWithDefaults() *WorkflowOutputsArchive {
	this := WorkflowOutputsArchive{}
	return &this
}

// GetReason returns the Reason field value
func (o *WorkflowOutputsArchive) GetReason() WorkflowOutputArchiveReason {
	if o == nil {
		var ret WorkflowOutputArchiveReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *WorkflowOutputsArchive) GetReasonOk() (*WorkflowOutputArchiveReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *WorkflowOutputsArchive) SetReason(v WorkflowOutputArchiveReason) {
	o.Reason = v
}

// GetWorkflowOutputIds returns the WorkflowOutputIds field value
func (o *WorkflowOutputsArchive) GetWorkflowOutputIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WorkflowOutputIds
}

// GetWorkflowOutputIdsOk returns a tuple with the WorkflowOutputIds field value
// and a boolean to check if the value has been set.
func (o *WorkflowOutputsArchive) GetWorkflowOutputIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WorkflowOutputIds, true
}

// SetWorkflowOutputIds sets field value
func (o *WorkflowOutputsArchive) SetWorkflowOutputIds(v []string) {
	o.WorkflowOutputIds = v
}

func (o WorkflowOutputsArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["workflowOutputIds"] = o.WorkflowOutputIds
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowOutputsArchive struct {
	value *WorkflowOutputsArchive
	isSet bool
}

func (v NullableWorkflowOutputsArchive) Get() *WorkflowOutputsArchive {
	return v.value
}

func (v *NullableWorkflowOutputsArchive) Set(val *WorkflowOutputsArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOutputsArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOutputsArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOutputsArchive(val *WorkflowOutputsArchive) *NullableWorkflowOutputsArchive {
	return &NullableWorkflowOutputsArchive{value: val, isSet: true}
}

func (v NullableWorkflowOutputsArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOutputsArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


