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

// WorkflowOutputsUnarchive struct for WorkflowOutputsUnarchive
type WorkflowOutputsUnarchive struct {
	WorkflowOutputIds []string `json:"workflowOutputIds"`
}

// NewWorkflowOutputsUnarchive instantiates a new WorkflowOutputsUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOutputsUnarchive(workflowOutputIds []string) *WorkflowOutputsUnarchive {
	this := WorkflowOutputsUnarchive{}
	this.WorkflowOutputIds = workflowOutputIds
	return &this
}

// NewWorkflowOutputsUnarchiveWithDefaults instantiates a new WorkflowOutputsUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOutputsUnarchiveWithDefaults() *WorkflowOutputsUnarchive {
	this := WorkflowOutputsUnarchive{}
	return &this
}

// GetWorkflowOutputIds returns the WorkflowOutputIds field value
func (o *WorkflowOutputsUnarchive) GetWorkflowOutputIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WorkflowOutputIds
}

// GetWorkflowOutputIdsOk returns a tuple with the WorkflowOutputIds field value
// and a boolean to check if the value has been set.
func (o *WorkflowOutputsUnarchive) GetWorkflowOutputIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WorkflowOutputIds, true
}

// SetWorkflowOutputIds sets field value
func (o *WorkflowOutputsUnarchive) SetWorkflowOutputIds(v []string) {
	o.WorkflowOutputIds = v
}

func (o WorkflowOutputsUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowOutputIds"] = o.WorkflowOutputIds
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowOutputsUnarchive struct {
	value *WorkflowOutputsUnarchive
	isSet bool
}

func (v NullableWorkflowOutputsUnarchive) Get() *WorkflowOutputsUnarchive {
	return v.value
}

func (v *NullableWorkflowOutputsUnarchive) Set(val *WorkflowOutputsUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOutputsUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOutputsUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOutputsUnarchive(val *WorkflowOutputsUnarchive) *NullableWorkflowOutputsUnarchive {
	return &NullableWorkflowOutputsUnarchive{value: val, isSet: true}
}

func (v NullableWorkflowOutputsUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOutputsUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

