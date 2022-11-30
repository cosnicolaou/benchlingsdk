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

// WorkflowTaskBulkUpdate struct for WorkflowTaskBulkUpdate
type WorkflowTaskBulkUpdate struct {
	StatusId *string `json:"statusId,omitempty"`
	// The id of the user assigned to the task
	AssigneeId *string `json:"assigneeId,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	// The date on which the task is scheduled to be executed
	ScheduledOn *string `json:"scheduledOn,omitempty"`
	// The workflow task ID
	WorkflowTaskId *string `json:"workflowTaskId,omitempty"`
}

// NewWorkflowTaskBulkUpdate instantiates a new WorkflowTaskBulkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskBulkUpdate() *WorkflowTaskBulkUpdate {
	this := WorkflowTaskBulkUpdate{}
	return &this
}

// NewWorkflowTaskBulkUpdateWithDefaults instantiates a new WorkflowTaskBulkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskBulkUpdateWithDefaults() *WorkflowTaskBulkUpdate {
	this := WorkflowTaskBulkUpdate{}
	return &this
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *WorkflowTaskBulkUpdate) GetStatusId() string {
	if o == nil || isNil(o.StatusId) {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkUpdate) GetStatusIdOk() (*string, bool) {
	if o == nil || isNil(o.StatusId) {
    return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *WorkflowTaskBulkUpdate) HasStatusId() bool {
	if o != nil && !isNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *WorkflowTaskBulkUpdate) SetStatusId(v string) {
	o.StatusId = &v
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *WorkflowTaskBulkUpdate) GetAssigneeId() string {
	if o == nil || isNil(o.AssigneeId) {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkUpdate) GetAssigneeIdOk() (*string, bool) {
	if o == nil || isNil(o.AssigneeId) {
    return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *WorkflowTaskBulkUpdate) HasAssigneeId() bool {
	if o != nil && !isNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *WorkflowTaskBulkUpdate) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *WorkflowTaskBulkUpdate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkUpdate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *WorkflowTaskBulkUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *WorkflowTaskBulkUpdate) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetScheduledOn returns the ScheduledOn field value if set, zero value otherwise.
func (o *WorkflowTaskBulkUpdate) GetScheduledOn() string {
	if o == nil || isNil(o.ScheduledOn) {
		var ret string
		return ret
	}
	return *o.ScheduledOn
}

// GetScheduledOnOk returns a tuple with the ScheduledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkUpdate) GetScheduledOnOk() (*string, bool) {
	if o == nil || isNil(o.ScheduledOn) {
    return nil, false
	}
	return o.ScheduledOn, true
}

// HasScheduledOn returns a boolean if a field has been set.
func (o *WorkflowTaskBulkUpdate) HasScheduledOn() bool {
	if o != nil && !isNil(o.ScheduledOn) {
		return true
	}

	return false
}

// SetScheduledOn gets a reference to the given string and assigns it to the ScheduledOn field.
func (o *WorkflowTaskBulkUpdate) SetScheduledOn(v string) {
	o.ScheduledOn = &v
}

// GetWorkflowTaskId returns the WorkflowTaskId field value if set, zero value otherwise.
func (o *WorkflowTaskBulkUpdate) GetWorkflowTaskId() string {
	if o == nil || isNil(o.WorkflowTaskId) {
		var ret string
		return ret
	}
	return *o.WorkflowTaskId
}

// GetWorkflowTaskIdOk returns a tuple with the WorkflowTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskBulkUpdate) GetWorkflowTaskIdOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowTaskId) {
    return nil, false
	}
	return o.WorkflowTaskId, true
}

// HasWorkflowTaskId returns a boolean if a field has been set.
func (o *WorkflowTaskBulkUpdate) HasWorkflowTaskId() bool {
	if o != nil && !isNil(o.WorkflowTaskId) {
		return true
	}

	return false
}

// SetWorkflowTaskId gets a reference to the given string and assigns it to the WorkflowTaskId field.
func (o *WorkflowTaskBulkUpdate) SetWorkflowTaskId(v string) {
	o.WorkflowTaskId = &v
}

func (o WorkflowTaskBulkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StatusId) {
		toSerialize["statusId"] = o.StatusId
	}
	if !isNil(o.AssigneeId) {
		toSerialize["assigneeId"] = o.AssigneeId
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.ScheduledOn) {
		toSerialize["scheduledOn"] = o.ScheduledOn
	}
	if !isNil(o.WorkflowTaskId) {
		toSerialize["workflowTaskId"] = o.WorkflowTaskId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskBulkUpdate struct {
	value *WorkflowTaskBulkUpdate
	isSet bool
}

func (v NullableWorkflowTaskBulkUpdate) Get() *WorkflowTaskBulkUpdate {
	return v.value
}

func (v *NullableWorkflowTaskBulkUpdate) Set(val *WorkflowTaskBulkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskBulkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskBulkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskBulkUpdate(val *WorkflowTaskBulkUpdate) *NullableWorkflowTaskBulkUpdate {
	return &NullableWorkflowTaskBulkUpdate{value: val, isSet: true}
}

func (v NullableWorkflowTaskBulkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskBulkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

