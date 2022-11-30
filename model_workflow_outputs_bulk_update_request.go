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

// WorkflowOutputsBulkUpdateRequest struct for WorkflowOutputsBulkUpdateRequest
type WorkflowOutputsBulkUpdateRequest struct {
	WorkflowOutputs []WorkflowOutputBulkUpdate `json:"workflowOutputs,omitempty"`
}

// NewWorkflowOutputsBulkUpdateRequest instantiates a new WorkflowOutputsBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOutputsBulkUpdateRequest() *WorkflowOutputsBulkUpdateRequest {
	this := WorkflowOutputsBulkUpdateRequest{}
	return &this
}

// NewWorkflowOutputsBulkUpdateRequestWithDefaults instantiates a new WorkflowOutputsBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOutputsBulkUpdateRequestWithDefaults() *WorkflowOutputsBulkUpdateRequest {
	this := WorkflowOutputsBulkUpdateRequest{}
	return &this
}

// GetWorkflowOutputs returns the WorkflowOutputs field value if set, zero value otherwise.
func (o *WorkflowOutputsBulkUpdateRequest) GetWorkflowOutputs() []WorkflowOutputBulkUpdate {
	if o == nil || isNil(o.WorkflowOutputs) {
		var ret []WorkflowOutputBulkUpdate
		return ret
	}
	return o.WorkflowOutputs
}

// GetWorkflowOutputsOk returns a tuple with the WorkflowOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOutputsBulkUpdateRequest) GetWorkflowOutputsOk() ([]WorkflowOutputBulkUpdate, bool) {
	if o == nil || isNil(o.WorkflowOutputs) {
    return nil, false
	}
	return o.WorkflowOutputs, true
}

// HasWorkflowOutputs returns a boolean if a field has been set.
func (o *WorkflowOutputsBulkUpdateRequest) HasWorkflowOutputs() bool {
	if o != nil && !isNil(o.WorkflowOutputs) {
		return true
	}

	return false
}

// SetWorkflowOutputs gets a reference to the given []WorkflowOutputBulkUpdate and assigns it to the WorkflowOutputs field.
func (o *WorkflowOutputsBulkUpdateRequest) SetWorkflowOutputs(v []WorkflowOutputBulkUpdate) {
	o.WorkflowOutputs = v
}

func (o WorkflowOutputsBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WorkflowOutputs) {
		toSerialize["workflowOutputs"] = o.WorkflowOutputs
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowOutputsBulkUpdateRequest struct {
	value *WorkflowOutputsBulkUpdateRequest
	isSet bool
}

func (v NullableWorkflowOutputsBulkUpdateRequest) Get() *WorkflowOutputsBulkUpdateRequest {
	return v.value
}

func (v *NullableWorkflowOutputsBulkUpdateRequest) Set(val *WorkflowOutputsBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOutputsBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOutputsBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOutputsBulkUpdateRequest(val *WorkflowOutputsBulkUpdateRequest) *NullableWorkflowOutputsBulkUpdateRequest {
	return &NullableWorkflowOutputsBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableWorkflowOutputsBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOutputsBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


