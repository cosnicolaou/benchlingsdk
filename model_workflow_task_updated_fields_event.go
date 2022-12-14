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

// WorkflowTaskUpdatedFieldsEvent struct for WorkflowTaskUpdatedFieldsEvent
type WorkflowTaskUpdatedFieldsEvent struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	// These properties have been dropped from the payload due to size. 
	ExcludedProperties []string `json:"excludedProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Schema NullableEventBaseSchema `json:"schema,omitempty"`
	EventType *string `json:"eventType,omitempty"`
	WorkflowTask *WorkflowTask `json:"workflowTask,omitempty"`
}

// NewWorkflowTaskUpdatedFieldsEvent instantiates a new WorkflowTaskUpdatedFieldsEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskUpdatedFieldsEvent() *WorkflowTaskUpdatedFieldsEvent {
	this := WorkflowTaskUpdatedFieldsEvent{}
	return &this
}

// NewWorkflowTaskUpdatedFieldsEventWithDefaults instantiates a new WorkflowTaskUpdatedFieldsEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskUpdatedFieldsEventWithDefaults() *WorkflowTaskUpdatedFieldsEvent {
	this := WorkflowTaskUpdatedFieldsEvent{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkflowTaskUpdatedFieldsEvent) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *WorkflowTaskUpdatedFieldsEvent) GetDeprecated() bool {
	if o == nil || isNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) GetDeprecatedOk() (*bool, bool) {
	if o == nil || isNil(o.Deprecated) {
    return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasDeprecated() bool {
	if o != nil && !isNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetExcludedProperties returns the ExcludedProperties field value if set, zero value otherwise.
func (o *WorkflowTaskUpdatedFieldsEvent) GetExcludedProperties() []string {
	if o == nil || isNil(o.ExcludedProperties) {
		var ret []string
		return ret
	}
	return o.ExcludedProperties
}

// GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) GetExcludedPropertiesOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedProperties) {
    return nil, false
	}
	return o.ExcludedProperties, true
}

// HasExcludedProperties returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasExcludedProperties() bool {
	if o != nil && !isNil(o.ExcludedProperties) {
		return true
	}

	return false
}

// SetExcludedProperties gets a reference to the given []string and assigns it to the ExcludedProperties field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetExcludedProperties(v []string) {
	o.ExcludedProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowTaskUpdatedFieldsEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetId(v string) {
	o.Id = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskUpdatedFieldsEvent) GetSchema() EventBaseSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret EventBaseSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskUpdatedFieldsEvent) GetSchemaOk() (*EventBaseSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableEventBaseSchema and assigns it to the Schema field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetSchema(v EventBaseSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *WorkflowTaskUpdatedFieldsEvent) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *WorkflowTaskUpdatedFieldsEvent) UnsetSchema() {
	o.Schema.Unset()
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *WorkflowTaskUpdatedFieldsEvent) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetWorkflowTask returns the WorkflowTask field value if set, zero value otherwise.
func (o *WorkflowTaskUpdatedFieldsEvent) GetWorkflowTask() WorkflowTask {
	if o == nil || isNil(o.WorkflowTask) {
		var ret WorkflowTask
		return ret
	}
	return *o.WorkflowTask
}

// GetWorkflowTaskOk returns a tuple with the WorkflowTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) GetWorkflowTaskOk() (*WorkflowTask, bool) {
	if o == nil || isNil(o.WorkflowTask) {
    return nil, false
	}
	return o.WorkflowTask, true
}

// HasWorkflowTask returns a boolean if a field has been set.
func (o *WorkflowTaskUpdatedFieldsEvent) HasWorkflowTask() bool {
	if o != nil && !isNil(o.WorkflowTask) {
		return true
	}

	return false
}

// SetWorkflowTask gets a reference to the given WorkflowTask and assigns it to the WorkflowTask field.
func (o *WorkflowTaskUpdatedFieldsEvent) SetWorkflowTask(v WorkflowTask) {
	o.WorkflowTask = &v
}

func (o WorkflowTaskUpdatedFieldsEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !isNil(o.ExcludedProperties) {
		toSerialize["excludedProperties"] = o.ExcludedProperties
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.WorkflowTask) {
		toSerialize["workflowTask"] = o.WorkflowTask
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskUpdatedFieldsEvent struct {
	value *WorkflowTaskUpdatedFieldsEvent
	isSet bool
}

func (v NullableWorkflowTaskUpdatedFieldsEvent) Get() *WorkflowTaskUpdatedFieldsEvent {
	return v.value
}

func (v *NullableWorkflowTaskUpdatedFieldsEvent) Set(val *WorkflowTaskUpdatedFieldsEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskUpdatedFieldsEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskUpdatedFieldsEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskUpdatedFieldsEvent(val *WorkflowTaskUpdatedFieldsEvent) *NullableWorkflowTaskUpdatedFieldsEvent {
	return &NullableWorkflowTaskUpdatedFieldsEvent{value: val, isSet: true}
}

func (v NullableWorkflowTaskUpdatedFieldsEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskUpdatedFieldsEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


