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

// AssayRunCreatedEvent struct for AssayRunCreatedEvent
type AssayRunCreatedEvent struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	// These properties have been dropped from the payload due to size. 
	ExcludedProperties []string `json:"excludedProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Schema NullableEventBaseSchema `json:"schema,omitempty"`
	AssayRun *AssayRun `json:"assayRun,omitempty"`
	EventType *string `json:"eventType,omitempty"`
}

// NewAssayRunCreatedEvent instantiates a new AssayRunCreatedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayRunCreatedEvent() *AssayRunCreatedEvent {
	this := AssayRunCreatedEvent{}
	return &this
}

// NewAssayRunCreatedEventWithDefaults instantiates a new AssayRunCreatedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayRunCreatedEventWithDefaults() *AssayRunCreatedEvent {
	this := AssayRunCreatedEvent{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssayRunCreatedEvent) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunCreatedEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AssayRunCreatedEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AssayRunCreatedEvent) GetDeprecated() bool {
	if o == nil || isNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunCreatedEvent) GetDeprecatedOk() (*bool, bool) {
	if o == nil || isNil(o.Deprecated) {
    return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasDeprecated() bool {
	if o != nil && !isNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *AssayRunCreatedEvent) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetExcludedProperties returns the ExcludedProperties field value if set, zero value otherwise.
func (o *AssayRunCreatedEvent) GetExcludedProperties() []string {
	if o == nil || isNil(o.ExcludedProperties) {
		var ret []string
		return ret
	}
	return o.ExcludedProperties
}

// GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunCreatedEvent) GetExcludedPropertiesOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedProperties) {
    return nil, false
	}
	return o.ExcludedProperties, true
}

// HasExcludedProperties returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasExcludedProperties() bool {
	if o != nil && !isNil(o.ExcludedProperties) {
		return true
	}

	return false
}

// SetExcludedProperties gets a reference to the given []string and assigns it to the ExcludedProperties field.
func (o *AssayRunCreatedEvent) SetExcludedProperties(v []string) {
	o.ExcludedProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssayRunCreatedEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunCreatedEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssayRunCreatedEvent) SetId(v string) {
	o.Id = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayRunCreatedEvent) GetSchema() EventBaseSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret EventBaseSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayRunCreatedEvent) GetSchemaOk() (*EventBaseSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableEventBaseSchema and assigns it to the Schema field.
func (o *AssayRunCreatedEvent) SetSchema(v EventBaseSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *AssayRunCreatedEvent) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *AssayRunCreatedEvent) UnsetSchema() {
	o.Schema.Unset()
}

// GetAssayRun returns the AssayRun field value if set, zero value otherwise.
func (o *AssayRunCreatedEvent) GetAssayRun() AssayRun {
	if o == nil || isNil(o.AssayRun) {
		var ret AssayRun
		return ret
	}
	return *o.AssayRun
}

// GetAssayRunOk returns a tuple with the AssayRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunCreatedEvent) GetAssayRunOk() (*AssayRun, bool) {
	if o == nil || isNil(o.AssayRun) {
    return nil, false
	}
	return o.AssayRun, true
}

// HasAssayRun returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasAssayRun() bool {
	if o != nil && !isNil(o.AssayRun) {
		return true
	}

	return false
}

// SetAssayRun gets a reference to the given AssayRun and assigns it to the AssayRun field.
func (o *AssayRunCreatedEvent) SetAssayRun(v AssayRun) {
	o.AssayRun = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *AssayRunCreatedEvent) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunCreatedEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *AssayRunCreatedEvent) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *AssayRunCreatedEvent) SetEventType(v string) {
	o.EventType = &v
}

func (o AssayRunCreatedEvent) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AssayRun) {
		toSerialize["assayRun"] = o.AssayRun
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return json.Marshal(toSerialize)
}

type NullableAssayRunCreatedEvent struct {
	value *AssayRunCreatedEvent
	isSet bool
}

func (v NullableAssayRunCreatedEvent) Get() *AssayRunCreatedEvent {
	return v.value
}

func (v *NullableAssayRunCreatedEvent) Set(val *AssayRunCreatedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayRunCreatedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayRunCreatedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayRunCreatedEvent(val *AssayRunCreatedEvent) *NullableAssayRunCreatedEvent {
	return &NullableAssayRunCreatedEvent{value: val, isSet: true}
}

func (v NullableAssayRunCreatedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayRunCreatedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


