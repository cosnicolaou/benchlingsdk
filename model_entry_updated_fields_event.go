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

// EntryUpdatedFieldsEvent struct for EntryUpdatedFieldsEvent
type EntryUpdatedFieldsEvent struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	// These properties have been dropped from the payload due to size. 
	ExcludedProperties []string `json:"excludedProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Schema NullableEventBaseSchema `json:"schema,omitempty"`
	// These properties have been updated, causing this message 
	Updates []string `json:"updates,omitempty"`
	Entry *Entry `json:"entry,omitempty"`
	EventType *string `json:"eventType,omitempty"`
}

// NewEntryUpdatedFieldsEvent instantiates a new EntryUpdatedFieldsEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryUpdatedFieldsEvent() *EntryUpdatedFieldsEvent {
	this := EntryUpdatedFieldsEvent{}
	return &this
}

// NewEntryUpdatedFieldsEventWithDefaults instantiates a new EntryUpdatedFieldsEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryUpdatedFieldsEventWithDefaults() *EntryUpdatedFieldsEvent {
	this := EntryUpdatedFieldsEvent{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EntryUpdatedFieldsEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetDeprecated() bool {
	if o == nil || isNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetDeprecatedOk() (*bool, bool) {
	if o == nil || isNil(o.Deprecated) {
    return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasDeprecated() bool {
	if o != nil && !isNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *EntryUpdatedFieldsEvent) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetExcludedProperties returns the ExcludedProperties field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetExcludedProperties() []string {
	if o == nil || isNil(o.ExcludedProperties) {
		var ret []string
		return ret
	}
	return o.ExcludedProperties
}

// GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetExcludedPropertiesOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedProperties) {
    return nil, false
	}
	return o.ExcludedProperties, true
}

// HasExcludedProperties returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasExcludedProperties() bool {
	if o != nil && !isNil(o.ExcludedProperties) {
		return true
	}

	return false
}

// SetExcludedProperties gets a reference to the given []string and assigns it to the ExcludedProperties field.
func (o *EntryUpdatedFieldsEvent) SetExcludedProperties(v []string) {
	o.ExcludedProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntryUpdatedFieldsEvent) SetId(v string) {
	o.Id = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryUpdatedFieldsEvent) GetSchema() EventBaseSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret EventBaseSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryUpdatedFieldsEvent) GetSchemaOk() (*EventBaseSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableEventBaseSchema and assigns it to the Schema field.
func (o *EntryUpdatedFieldsEvent) SetSchema(v EventBaseSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *EntryUpdatedFieldsEvent) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *EntryUpdatedFieldsEvent) UnsetSchema() {
	o.Schema.Unset()
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetUpdates() []string {
	if o == nil || isNil(o.Updates) {
		var ret []string
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetUpdatesOk() ([]string, bool) {
	if o == nil || isNil(o.Updates) {
    return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasUpdates() bool {
	if o != nil && !isNil(o.Updates) {
		return true
	}

	return false
}

// SetUpdates gets a reference to the given []string and assigns it to the Updates field.
func (o *EntryUpdatedFieldsEvent) SetUpdates(v []string) {
	o.Updates = v
}

// GetEntry returns the Entry field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetEntry() Entry {
	if o == nil || isNil(o.Entry) {
		var ret Entry
		return ret
	}
	return *o.Entry
}

// GetEntryOk returns a tuple with the Entry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetEntryOk() (*Entry, bool) {
	if o == nil || isNil(o.Entry) {
    return nil, false
	}
	return o.Entry, true
}

// HasEntry returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasEntry() bool {
	if o != nil && !isNil(o.Entry) {
		return true
	}

	return false
}

// SetEntry gets a reference to the given Entry and assigns it to the Entry field.
func (o *EntryUpdatedFieldsEvent) SetEntry(v Entry) {
	o.Entry = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EntryUpdatedFieldsEvent) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryUpdatedFieldsEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EntryUpdatedFieldsEvent) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *EntryUpdatedFieldsEvent) SetEventType(v string) {
	o.EventType = &v
}

func (o EntryUpdatedFieldsEvent) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Updates) {
		toSerialize["updates"] = o.Updates
	}
	if !isNil(o.Entry) {
		toSerialize["entry"] = o.Entry
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return json.Marshal(toSerialize)
}

type NullableEntryUpdatedFieldsEvent struct {
	value *EntryUpdatedFieldsEvent
	isSet bool
}

func (v NullableEntryUpdatedFieldsEvent) Get() *EntryUpdatedFieldsEvent {
	return v.value
}

func (v *NullableEntryUpdatedFieldsEvent) Set(val *EntryUpdatedFieldsEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryUpdatedFieldsEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryUpdatedFieldsEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryUpdatedFieldsEvent(val *EntryUpdatedFieldsEvent) *NullableEntryUpdatedFieldsEvent {
	return &NullableEntryUpdatedFieldsEvent{value: val, isSet: true}
}

func (v NullableEntryUpdatedFieldsEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryUpdatedFieldsEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


