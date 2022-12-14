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

// WorkflowTaskWriteBase struct for WorkflowTaskWriteBase
type WorkflowTaskWriteBase struct {
	// The id of the user assigned to the task
	AssigneeId *string `json:"assigneeId,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	// The date on which the task is scheduled to be executed
	ScheduledOn *string `json:"scheduledOn,omitempty"`
}

// NewWorkflowTaskWriteBase instantiates a new WorkflowTaskWriteBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskWriteBase() *WorkflowTaskWriteBase {
	this := WorkflowTaskWriteBase{}
	return &this
}

// NewWorkflowTaskWriteBaseWithDefaults instantiates a new WorkflowTaskWriteBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskWriteBaseWithDefaults() *WorkflowTaskWriteBase {
	this := WorkflowTaskWriteBase{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *WorkflowTaskWriteBase) GetAssigneeId() string {
	if o == nil || isNil(o.AssigneeId) {
		var ret string
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskWriteBase) GetAssigneeIdOk() (*string, bool) {
	if o == nil || isNil(o.AssigneeId) {
    return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *WorkflowTaskWriteBase) HasAssigneeId() bool {
	if o != nil && !isNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given string and assigns it to the AssigneeId field.
func (o *WorkflowTaskWriteBase) SetAssigneeId(v string) {
	o.AssigneeId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *WorkflowTaskWriteBase) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskWriteBase) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *WorkflowTaskWriteBase) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *WorkflowTaskWriteBase) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetScheduledOn returns the ScheduledOn field value if set, zero value otherwise.
func (o *WorkflowTaskWriteBase) GetScheduledOn() string {
	if o == nil || isNil(o.ScheduledOn) {
		var ret string
		return ret
	}
	return *o.ScheduledOn
}

// GetScheduledOnOk returns a tuple with the ScheduledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskWriteBase) GetScheduledOnOk() (*string, bool) {
	if o == nil || isNil(o.ScheduledOn) {
    return nil, false
	}
	return o.ScheduledOn, true
}

// HasScheduledOn returns a boolean if a field has been set.
func (o *WorkflowTaskWriteBase) HasScheduledOn() bool {
	if o != nil && !isNil(o.ScheduledOn) {
		return true
	}

	return false
}

// SetScheduledOn gets a reference to the given string and assigns it to the ScheduledOn field.
func (o *WorkflowTaskWriteBase) SetScheduledOn(v string) {
	o.ScheduledOn = &v
}

func (o WorkflowTaskWriteBase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskWriteBase struct {
	value *WorkflowTaskWriteBase
	isSet bool
}

func (v NullableWorkflowTaskWriteBase) Get() *WorkflowTaskWriteBase {
	return v.value
}

func (v *NullableWorkflowTaskWriteBase) Set(val *WorkflowTaskWriteBase) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskWriteBase) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskWriteBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskWriteBase(val *WorkflowTaskWriteBase) *NullableWorkflowTaskWriteBase {
	return &NullableWorkflowTaskWriteBase{value: val, isSet: true}
}

func (v NullableWorkflowTaskWriteBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskWriteBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


