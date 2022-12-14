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

// WorkflowTaskCreatedEventAllOf struct for WorkflowTaskCreatedEventAllOf
type WorkflowTaskCreatedEventAllOf struct {
	EventType *string `json:"eventType,omitempty"`
	WorkflowTask *WorkflowTask `json:"workflowTask,omitempty"`
}

// NewWorkflowTaskCreatedEventAllOf instantiates a new WorkflowTaskCreatedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskCreatedEventAllOf() *WorkflowTaskCreatedEventAllOf {
	this := WorkflowTaskCreatedEventAllOf{}
	return &this
}

// NewWorkflowTaskCreatedEventAllOfWithDefaults instantiates a new WorkflowTaskCreatedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskCreatedEventAllOfWithDefaults() *WorkflowTaskCreatedEventAllOf {
	this := WorkflowTaskCreatedEventAllOf{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *WorkflowTaskCreatedEventAllOf) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskCreatedEventAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *WorkflowTaskCreatedEventAllOf) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *WorkflowTaskCreatedEventAllOf) SetEventType(v string) {
	o.EventType = &v
}

// GetWorkflowTask returns the WorkflowTask field value if set, zero value otherwise.
func (o *WorkflowTaskCreatedEventAllOf) GetWorkflowTask() WorkflowTask {
	if o == nil || isNil(o.WorkflowTask) {
		var ret WorkflowTask
		return ret
	}
	return *o.WorkflowTask
}

// GetWorkflowTaskOk returns a tuple with the WorkflowTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskCreatedEventAllOf) GetWorkflowTaskOk() (*WorkflowTask, bool) {
	if o == nil || isNil(o.WorkflowTask) {
    return nil, false
	}
	return o.WorkflowTask, true
}

// HasWorkflowTask returns a boolean if a field has been set.
func (o *WorkflowTaskCreatedEventAllOf) HasWorkflowTask() bool {
	if o != nil && !isNil(o.WorkflowTask) {
		return true
	}

	return false
}

// SetWorkflowTask gets a reference to the given WorkflowTask and assigns it to the WorkflowTask field.
func (o *WorkflowTaskCreatedEventAllOf) SetWorkflowTask(v WorkflowTask) {
	o.WorkflowTask = &v
}

func (o WorkflowTaskCreatedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.WorkflowTask) {
		toSerialize["workflowTask"] = o.WorkflowTask
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskCreatedEventAllOf struct {
	value *WorkflowTaskCreatedEventAllOf
	isSet bool
}

func (v NullableWorkflowTaskCreatedEventAllOf) Get() *WorkflowTaskCreatedEventAllOf {
	return v.value
}

func (v *NullableWorkflowTaskCreatedEventAllOf) Set(val *WorkflowTaskCreatedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskCreatedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskCreatedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskCreatedEventAllOf(val *WorkflowTaskCreatedEventAllOf) *NullableWorkflowTaskCreatedEventAllOf {
	return &NullableWorkflowTaskCreatedEventAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskCreatedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskCreatedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


