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

// WorkflowTaskGroupsArchive struct for WorkflowTaskGroupsArchive
type WorkflowTaskGroupsArchive struct {
	Reason WorkflowTaskGroupArchiveReason `json:"reason"`
	WorkflowTaskGroupIds []string `json:"workflowTaskGroupIds"`
}

// NewWorkflowTaskGroupsArchive instantiates a new WorkflowTaskGroupsArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskGroupsArchive(reason WorkflowTaskGroupArchiveReason, workflowTaskGroupIds []string) *WorkflowTaskGroupsArchive {
	this := WorkflowTaskGroupsArchive{}
	this.Reason = reason
	this.WorkflowTaskGroupIds = workflowTaskGroupIds
	return &this
}

// NewWorkflowTaskGroupsArchiveWithDefaults instantiates a new WorkflowTaskGroupsArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskGroupsArchiveWithDefaults() *WorkflowTaskGroupsArchive {
	this := WorkflowTaskGroupsArchive{}
	return &this
}

// GetReason returns the Reason field value
func (o *WorkflowTaskGroupsArchive) GetReason() WorkflowTaskGroupArchiveReason {
	if o == nil {
		var ret WorkflowTaskGroupArchiveReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupsArchive) GetReasonOk() (*WorkflowTaskGroupArchiveReason, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *WorkflowTaskGroupsArchive) SetReason(v WorkflowTaskGroupArchiveReason) {
	o.Reason = v
}

// GetWorkflowTaskGroupIds returns the WorkflowTaskGroupIds field value
func (o *WorkflowTaskGroupsArchive) GetWorkflowTaskGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WorkflowTaskGroupIds
}

// GetWorkflowTaskGroupIdsOk returns a tuple with the WorkflowTaskGroupIds field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupsArchive) GetWorkflowTaskGroupIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WorkflowTaskGroupIds, true
}

// SetWorkflowTaskGroupIds sets field value
func (o *WorkflowTaskGroupsArchive) SetWorkflowTaskGroupIds(v []string) {
	o.WorkflowTaskGroupIds = v
}

func (o WorkflowTaskGroupsArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["workflowTaskGroupIds"] = o.WorkflowTaskGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskGroupsArchive struct {
	value *WorkflowTaskGroupsArchive
	isSet bool
}

func (v NullableWorkflowTaskGroupsArchive) Get() *WorkflowTaskGroupsArchive {
	return v.value
}

func (v *NullableWorkflowTaskGroupsArchive) Set(val *WorkflowTaskGroupsArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskGroupsArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskGroupsArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskGroupsArchive(val *WorkflowTaskGroupsArchive) *NullableWorkflowTaskGroupsArchive {
	return &NullableWorkflowTaskGroupsArchive{value: val, isSet: true}
}

func (v NullableWorkflowTaskGroupsArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskGroupsArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


