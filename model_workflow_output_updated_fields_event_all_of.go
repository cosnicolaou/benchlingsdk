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

// WorkflowOutputUpdatedFieldsEventAllOf struct for WorkflowOutputUpdatedFieldsEventAllOf
type WorkflowOutputUpdatedFieldsEventAllOf struct {
	EventType *string `json:"eventType,omitempty"`
	WorkflowOutput *WorkflowOutput `json:"workflowOutput,omitempty"`
}

// NewWorkflowOutputUpdatedFieldsEventAllOf instantiates a new WorkflowOutputUpdatedFieldsEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOutputUpdatedFieldsEventAllOf() *WorkflowOutputUpdatedFieldsEventAllOf {
	this := WorkflowOutputUpdatedFieldsEventAllOf{}
	return &this
}

// NewWorkflowOutputUpdatedFieldsEventAllOfWithDefaults instantiates a new WorkflowOutputUpdatedFieldsEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOutputUpdatedFieldsEventAllOfWithDefaults() *WorkflowOutputUpdatedFieldsEventAllOf {
	this := WorkflowOutputUpdatedFieldsEventAllOf{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) SetEventType(v string) {
	o.EventType = &v
}

// GetWorkflowOutput returns the WorkflowOutput field value if set, zero value otherwise.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) GetWorkflowOutput() WorkflowOutput {
	if o == nil || isNil(o.WorkflowOutput) {
		var ret WorkflowOutput
		return ret
	}
	return *o.WorkflowOutput
}

// GetWorkflowOutputOk returns a tuple with the WorkflowOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) GetWorkflowOutputOk() (*WorkflowOutput, bool) {
	if o == nil || isNil(o.WorkflowOutput) {
    return nil, false
	}
	return o.WorkflowOutput, true
}

// HasWorkflowOutput returns a boolean if a field has been set.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) HasWorkflowOutput() bool {
	if o != nil && !isNil(o.WorkflowOutput) {
		return true
	}

	return false
}

// SetWorkflowOutput gets a reference to the given WorkflowOutput and assigns it to the WorkflowOutput field.
func (o *WorkflowOutputUpdatedFieldsEventAllOf) SetWorkflowOutput(v WorkflowOutput) {
	o.WorkflowOutput = &v
}

func (o WorkflowOutputUpdatedFieldsEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.WorkflowOutput) {
		toSerialize["workflowOutput"] = o.WorkflowOutput
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowOutputUpdatedFieldsEventAllOf struct {
	value *WorkflowOutputUpdatedFieldsEventAllOf
	isSet bool
}

func (v NullableWorkflowOutputUpdatedFieldsEventAllOf) Get() *WorkflowOutputUpdatedFieldsEventAllOf {
	return v.value
}

func (v *NullableWorkflowOutputUpdatedFieldsEventAllOf) Set(val *WorkflowOutputUpdatedFieldsEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOutputUpdatedFieldsEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOutputUpdatedFieldsEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOutputUpdatedFieldsEventAllOf(val *WorkflowOutputUpdatedFieldsEventAllOf) *NullableWorkflowOutputUpdatedFieldsEventAllOf {
	return &NullableWorkflowOutputUpdatedFieldsEventAllOf{value: val, isSet: true}
}

func (v NullableWorkflowOutputUpdatedFieldsEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOutputUpdatedFieldsEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


