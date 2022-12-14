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

// WorkflowTaskGroupCreatedEventAllOf struct for WorkflowTaskGroupCreatedEventAllOf
type WorkflowTaskGroupCreatedEventAllOf struct {
	EventType *string `json:"eventType,omitempty"`
	WorkflowTaskGroup *WorkflowTaskGroup `json:"workflowTaskGroup,omitempty"`
}

// NewWorkflowTaskGroupCreatedEventAllOf instantiates a new WorkflowTaskGroupCreatedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskGroupCreatedEventAllOf() *WorkflowTaskGroupCreatedEventAllOf {
	this := WorkflowTaskGroupCreatedEventAllOf{}
	return &this
}

// NewWorkflowTaskGroupCreatedEventAllOfWithDefaults instantiates a new WorkflowTaskGroupCreatedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskGroupCreatedEventAllOfWithDefaults() *WorkflowTaskGroupCreatedEventAllOf {
	this := WorkflowTaskGroupCreatedEventAllOf{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *WorkflowTaskGroupCreatedEventAllOf) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupCreatedEventAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *WorkflowTaskGroupCreatedEventAllOf) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *WorkflowTaskGroupCreatedEventAllOf) SetEventType(v string) {
	o.EventType = &v
}

// GetWorkflowTaskGroup returns the WorkflowTaskGroup field value if set, zero value otherwise.
func (o *WorkflowTaskGroupCreatedEventAllOf) GetWorkflowTaskGroup() WorkflowTaskGroup {
	if o == nil || isNil(o.WorkflowTaskGroup) {
		var ret WorkflowTaskGroup
		return ret
	}
	return *o.WorkflowTaskGroup
}

// GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupCreatedEventAllOf) GetWorkflowTaskGroupOk() (*WorkflowTaskGroup, bool) {
	if o == nil || isNil(o.WorkflowTaskGroup) {
    return nil, false
	}
	return o.WorkflowTaskGroup, true
}

// HasWorkflowTaskGroup returns a boolean if a field has been set.
func (o *WorkflowTaskGroupCreatedEventAllOf) HasWorkflowTaskGroup() bool {
	if o != nil && !isNil(o.WorkflowTaskGroup) {
		return true
	}

	return false
}

// SetWorkflowTaskGroup gets a reference to the given WorkflowTaskGroup and assigns it to the WorkflowTaskGroup field.
func (o *WorkflowTaskGroupCreatedEventAllOf) SetWorkflowTaskGroup(v WorkflowTaskGroup) {
	o.WorkflowTaskGroup = &v
}

func (o WorkflowTaskGroupCreatedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.WorkflowTaskGroup) {
		toSerialize["workflowTaskGroup"] = o.WorkflowTaskGroup
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskGroupCreatedEventAllOf struct {
	value *WorkflowTaskGroupCreatedEventAllOf
	isSet bool
}

func (v NullableWorkflowTaskGroupCreatedEventAllOf) Get() *WorkflowTaskGroupCreatedEventAllOf {
	return v.value
}

func (v *NullableWorkflowTaskGroupCreatedEventAllOf) Set(val *WorkflowTaskGroupCreatedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskGroupCreatedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskGroupCreatedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskGroupCreatedEventAllOf(val *WorkflowTaskGroupCreatedEventAllOf) *NullableWorkflowTaskGroupCreatedEventAllOf {
	return &NullableWorkflowTaskGroupCreatedEventAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskGroupCreatedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskGroupCreatedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


