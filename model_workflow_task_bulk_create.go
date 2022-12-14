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

// WorkflowTaskBulkCreate struct for WorkflowTaskBulkCreate
type WorkflowTaskBulkCreate struct {
	// The id of the user assigned to the task
	AssigneeId *string `json:"assigneeId,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	// The date on which the task is scheduled to be executed
	ScheduledOn *string `json:"scheduledOn,omitempty"`
	// The workflow ID
	WorkflowTaskGroupId string `json:"workflowTaskGroupId"`
}

// NewWorkflowTaskBulkCreate instantiates a new WorkflowTaskBulkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskBulkCreate(workflowTaskGroupId string) *WorkflowTaskBulkCreate {
	this := WorkflowTaskBulkCreate{}
	this.WorkflowTaskGroupId = workflowTaskGroupId
	return &this
}

// NewWorkflowTaskBulkCreateWithDefaults instantiates a new WorkflowTaskBulkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskBulkCreateWithDefaults() *WorkflowTaskBulkCreate {
	this := WorkflowTaskBulkCreate{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *WorkflowTaskBulkCreate) GetAssigneeId() string {
	if o == nil || isNil(o.AssigneeId) {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkCreate) GetAssigneeIdOk() (*string, bool) {
	if o == nil || isNil(o.AssigneeId) {
    return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *WorkflowTaskBulkCreate) HasAssigneeId() bool {
	if o != nil && !isNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *WorkflowTaskBulkCreate) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *WorkflowTaskBulkCreate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkCreate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *WorkflowTaskBulkCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *WorkflowTaskBulkCreate) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetScheduledOn returns the ScheduledOn field value if set, zero value otherwise.
func (o *WorkflowTaskBulkCreate) GetScheduledOn() string {
	if o == nil || isNil(o.ScheduledOn) {
		var ret string
		return ret
	}
	return *o.ScheduledOn
}

// GetScheduledOnOk returns a tuple with the ScheduledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkCreate) GetScheduledOnOk() (*string, bool) {
	if o == nil || isNil(o.ScheduledOn) {
    return nil, false
	}
	return o.ScheduledOn, true
}

// HasScheduledOn returns a boolean if a field has been set.
func (o *WorkflowTaskBulkCreate) HasScheduledOn() bool {
	if o != nil && !isNil(o.ScheduledOn) {
		return true
	}

	return false
}

// SetScheduledOn gets a reference to the given string and assigns it to the ScheduledOn field.
func (o *WorkflowTaskBulkCreate) SetScheduledOn(v string) {
	o.ScheduledOn = &v
}

// GetWorkflowTaskGroupId returns the WorkflowTaskGroupId field value
func (o *WorkflowTaskBulkCreate) GetWorkflowTaskGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowTaskGroupId
}

// GetWorkflowTaskGroupIdOk returns a tuple with the WorkflowTaskGroupId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkCreate) GetWorkflowTaskGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WorkflowTaskGroupId, true
}

// SetWorkflowTaskGroupId sets field value
func (o *WorkflowTaskBulkCreate) SetWorkflowTaskGroupId(v string) {
	o.WorkflowTaskGroupId = v
}

func (o WorkflowTaskBulkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssigneeId) {
		toSerialize["assigneeId"] = o.AssigneeId
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.ScheduledOn) {
		toSerialize["scheduledOn"] = o.ScheduledOn
	}
	if true {
		toSerialize["workflowTaskGroupId"] = o.WorkflowTaskGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskBulkCreate struct {
	value *WorkflowTaskBulkCreate
	isSet bool
}

func (v NullableWorkflowTaskBulkCreate) Get() *WorkflowTaskBulkCreate {
	return v.value
}

func (v *NullableWorkflowTaskBulkCreate) Set(val *WorkflowTaskBulkCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskBulkCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskBulkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskBulkCreate(val *WorkflowTaskBulkCreate) *NullableWorkflowTaskBulkCreate {
	return &NullableWorkflowTaskBulkCreate{value: val, isSet: true}
}

func (v NullableWorkflowTaskBulkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskBulkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


