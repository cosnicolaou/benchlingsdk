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

// WorkflowStageList struct for WorkflowStageList
type WorkflowStageList struct {
	WorkflowStages []WorkflowStage `json:"workflowStages,omitempty"`
}

// NewWorkflowStageList instantiates a new WorkflowStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStageList() *WorkflowStageList {
	this := WorkflowStageList{}
	return &this
}

// NewWorkflowStageListWithDefaults instantiates a new WorkflowStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStageListWithDefaults() *WorkflowStageList {
	this := WorkflowStageList{}
	return &this
}

// GetWorkflowStages returns the WorkflowStages field value if set, zero value otherwise.
func (o *WorkflowStageList) GetWorkflowStages() []WorkflowStage {
	if o == nil || isNil(o.WorkflowStages) {
		var ret []WorkflowStage
		return ret
	}
	return o.WorkflowStages
}

// GetWorkflowStagesOk returns a tuple with the WorkflowStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStageList) GetWorkflowStagesOk() ([]WorkflowStage, bool) {
	if o == nil || isNil(o.WorkflowStages) {
    return nil, false
	}
	return o.WorkflowStages, true
}

// HasWorkflowStages returns a boolean if a field has been set.
func (o *WorkflowStageList) HasWorkflowStages() bool {
	if o != nil && !isNil(o.WorkflowStages) {
		return true
	}

	return false
}

// SetWorkflowStages gets a reference to the given []WorkflowStage and assigns it to the WorkflowStages field.
func (o *WorkflowStageList) SetWorkflowStages(v []WorkflowStage) {
	o.WorkflowStages = v
}

func (o WorkflowStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WorkflowStages) {
		toSerialize["workflowStages"] = o.WorkflowStages
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowStageList struct {
	value *WorkflowStageList
	isSet bool
}

func (v NullableWorkflowStageList) Get() *WorkflowStageList {
	return v.value
}

func (v *NullableWorkflowStageList) Set(val *WorkflowStageList) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStageList) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStageList(val *WorkflowStageList) *NullableWorkflowStageList {
	return &NullableWorkflowStageList{value: val, isSet: true}
}

func (v NullableWorkflowStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


