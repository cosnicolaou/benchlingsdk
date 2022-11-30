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

// WorkflowTaskGroupAllOf struct for WorkflowTaskGroupAllOf
type WorkflowTaskGroupAllOf struct {
	// The method by which the workflow is executed
	ExecutionType *string `json:"executionType,omitempty"`
}

// NewWorkflowTaskGroupAllOf instantiates a new WorkflowTaskGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskGroupAllOf() *WorkflowTaskGroupAllOf {
	this := WorkflowTaskGroupAllOf{}
	return &this
}

// NewWorkflowTaskGroupAllOfWithDefaults instantiates a new WorkflowTaskGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskGroupAllOfWithDefaults() *WorkflowTaskGroupAllOf {
	this := WorkflowTaskGroupAllOf{}
	return &this
}

// GetExecutionType returns the ExecutionType field value if set, zero value otherwise.
func (o *WorkflowTaskGroupAllOf) GetExecutionType() string {
	if o == nil || isNil(o.ExecutionType) {
		var ret string
		return ret
	}
	return *o.ExecutionType
}

// GetExecutionTypeOk returns a tuple with the ExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupAllOf) GetExecutionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ExecutionType) {
    return nil, false
	}
	return o.ExecutionType, true
}

// HasExecutionType returns a boolean if a field has been set.
func (o *WorkflowTaskGroupAllOf) HasExecutionType() bool {
	if o != nil && !isNil(o.ExecutionType) {
		return true
	}

	return false
}

// SetExecutionType gets a reference to the given string and assigns it to the ExecutionType field.
func (o *WorkflowTaskGroupAllOf) SetExecutionType(v string) {
	o.ExecutionType = &v
}

func (o WorkflowTaskGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExecutionType) {
		toSerialize["executionType"] = o.ExecutionType
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskGroupAllOf struct {
	value *WorkflowTaskGroupAllOf
	isSet bool
}

func (v NullableWorkflowTaskGroupAllOf) Get() *WorkflowTaskGroupAllOf {
	return v.value
}

func (v *NullableWorkflowTaskGroupAllOf) Set(val *WorkflowTaskGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskGroupAllOf(val *WorkflowTaskGroupAllOf) *NullableWorkflowTaskGroupAllOf {
	return &NullableWorkflowTaskGroupAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


