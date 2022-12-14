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

// WorkflowTasksUnarchive struct for WorkflowTasksUnarchive
type WorkflowTasksUnarchive struct {
	WorkflowTaskIds []string `json:"workflowTaskIds"`
}

// NewWorkflowTasksUnarchive instantiates a new WorkflowTasksUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTasksUnarchive(workflowTaskIds []string) *WorkflowTasksUnarchive {
	this := WorkflowTasksUnarchive{}
	this.WorkflowTaskIds = workflowTaskIds
	return &this
}

// NewWorkflowTasksUnarchiveWithDefaults instantiates a new WorkflowTasksUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTasksUnarchiveWithDefaults() *WorkflowTasksUnarchive {
	this := WorkflowTasksUnarchive{}
	return &this
}

// GetWorkflowTaskIds returns the WorkflowTaskIds field value
func (o *WorkflowTasksUnarchive) GetWorkflowTaskIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WorkflowTaskIds
}

// GetWorkflowTaskIdsOk returns a tuple with the WorkflowTaskIds field value
// and a boolean to check if the value has been set.
func (o *WorkflowTasksUnarchive) GetWorkflowTaskIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WorkflowTaskIds, true
}

// SetWorkflowTaskIds sets field value
func (o *WorkflowTasksUnarchive) SetWorkflowTaskIds(v []string) {
	o.WorkflowTaskIds = v
}

func (o WorkflowTasksUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowTaskIds"] = o.WorkflowTaskIds
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTasksUnarchive struct {
	value *WorkflowTasksUnarchive
	isSet bool
}

func (v NullableWorkflowTasksUnarchive) Get() *WorkflowTasksUnarchive {
	return v.value
}

func (v *NullableWorkflowTasksUnarchive) Set(val *WorkflowTasksUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTasksUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTasksUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTasksUnarchive(val *WorkflowTasksUnarchive) *NullableWorkflowTasksUnarchive {
	return &NullableWorkflowTasksUnarchive{value: val, isSet: true}
}

func (v NullableWorkflowTasksUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTasksUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


