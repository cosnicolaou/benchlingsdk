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

// WorkflowOutputsArchivalChange IDs of all items that were archived or unarchived, grouped by resource type 
type WorkflowOutputsArchivalChange struct {
	WorkflowOutputIds []string `json:"workflowOutputIds,omitempty"`
}

// NewWorkflowOutputsArchivalChange instantiates a new WorkflowOutputsArchivalChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOutputsArchivalChange() *WorkflowOutputsArchivalChange {
	this := WorkflowOutputsArchivalChange{}
	return &this
}

// NewWorkflowOutputsArchivalChangeWithDefaults instantiates a new WorkflowOutputsArchivalChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOutputsArchivalChangeWithDefaults() *WorkflowOutputsArchivalChange {
	this := WorkflowOutputsArchivalChange{}
	return &this
}

// GetWorkflowOutputIds returns the WorkflowOutputIds field value if set, zero value otherwise.
func (o *WorkflowOutputsArchivalChange) GetWorkflowOutputIds() []string {
	if o == nil || isNil(o.WorkflowOutputIds) {
		var ret []string
		return ret
	}
	return o.WorkflowOutputIds
}

// GetWorkflowOutputIdsOk returns a tuple with the WorkflowOutputIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOutputsArchivalChange) GetWorkflowOutputIdsOk() ([]string, bool) {
	if o == nil || isNil(o.WorkflowOutputIds) {
    return nil, false
	}
	return o.WorkflowOutputIds, true
}

// HasWorkflowOutputIds returns a boolean if a field has been set.
func (o *WorkflowOutputsArchivalChange) HasWorkflowOutputIds() bool {
	if o != nil && !isNil(o.WorkflowOutputIds) {
		return true
	}

	return false
}

// SetWorkflowOutputIds gets a reference to the given []string and assigns it to the WorkflowOutputIds field.
func (o *WorkflowOutputsArchivalChange) SetWorkflowOutputIds(v []string) {
	o.WorkflowOutputIds = v
}

func (o WorkflowOutputsArchivalChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WorkflowOutputIds) {
		toSerialize["workflowOutputIds"] = o.WorkflowOutputIds
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowOutputsArchivalChange struct {
	value *WorkflowOutputsArchivalChange
	isSet bool
}

func (v NullableWorkflowOutputsArchivalChange) Get() *WorkflowOutputsArchivalChange {
	return v.value
}

func (v *NullableWorkflowOutputsArchivalChange) Set(val *WorkflowOutputsArchivalChange) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOutputsArchivalChange) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOutputsArchivalChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOutputsArchivalChange(val *WorkflowOutputsArchivalChange) *NullableWorkflowOutputsArchivalChange {
	return &NullableWorkflowOutputsArchivalChange{value: val, isSet: true}
}

func (v NullableWorkflowOutputsArchivalChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOutputsArchivalChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


