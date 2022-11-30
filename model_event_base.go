/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
	"time"
)

// EventBase struct for EventBase
type EventBase struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	// These properties have been dropped from the payload due to size. 
	ExcludedProperties []string `json:"excludedProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Schema NullableEventBaseSchema `json:"schema,omitempty"`
}

// NewEventBase instantiates a new EventBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBase() *EventBase {
	this := EventBase{}
	return &this
}

// NewEventBaseWithDefaults instantiates a new EventBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBaseWithDefaults() *EventBase {
	this := EventBase{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EventBase) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBase) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventBase) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EventBase) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *EventBase) GetDeprecated() bool {
	if o == nil || isNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBase) GetDeprecatedOk() (*bool, bool) {
	if o == nil || isNil(o.Deprecated) {
    return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *EventBase) HasDeprecated() bool {
	if o != nil && !isNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *EventBase) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetExcludedProperties returns the ExcludedProperties field value if set, zero value otherwise.
func (o *EventBase) GetExcludedProperties() []string {
	if o == nil || isNil(o.ExcludedProperties) {
		var ret []string
		return ret
	}
	return o.ExcludedProperties
}

// GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBase) GetExcludedPropertiesOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedProperties) {
    return nil, false
	}
	return o.ExcludedProperties, true
}

// HasExcludedProperties returns a boolean if a field has been set.
func (o *EventBase) HasExcludedProperties() bool {
	if o != nil && !isNil(o.ExcludedProperties) {
		return true
	}

	return false
}

// SetExcludedProperties gets a reference to the given []string and assigns it to the ExcludedProperties field.
func (o *EventBase) SetExcludedProperties(v []string) {
	o.ExcludedProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventBase) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBase) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventBase) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventBase) SetId(v string) {
	o.Id = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventBase) GetSchema() EventBaseSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret EventBaseSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventBase) GetSchemaOk() (*EventBaseSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *EventBase) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableEventBaseSchema and assigns it to the Schema field.
func (o *EventBase) SetSchema(v EventBaseSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *EventBase) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *EventBase) UnsetSchema() {
	o.Schema.Unset()
}

func (o EventBase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEventBase struct {
	value *EventBase
	isSet bool
}

func (v NullableEventBase) Get() *EventBase {
	return v.value
}

func (v *NullableEventBase) Set(val *EventBase) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBase) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBase(val *EventBase) *NullableEventBase {
	return &NullableEventBase{value: val, isSet: true}
}

func (v NullableEventBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


