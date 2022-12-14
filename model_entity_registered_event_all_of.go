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

// EntityRegisteredEventAllOf struct for EntityRegisteredEventAllOf
type EntityRegisteredEventAllOf struct {
	Entity *GenericEntity `json:"entity,omitempty"`
	EventType *string `json:"eventType,omitempty"`
}

// NewEntityRegisteredEventAllOf instantiates a new EntityRegisteredEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRegisteredEventAllOf() *EntityRegisteredEventAllOf {
	this := EntityRegisteredEventAllOf{}
	return &this
}

// NewEntityRegisteredEventAllOfWithDefaults instantiates a new EntityRegisteredEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRegisteredEventAllOfWithDefaults() *EntityRegisteredEventAllOf {
	this := EntityRegisteredEventAllOf{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *EntityRegisteredEventAllOf) GetEntity() GenericEntity {
	if o == nil || isNil(o.Entity) {
		var ret GenericEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRegisteredEventAllOf) GetEntityOk() (*GenericEntity, bool) {
	if o == nil || isNil(o.Entity) {
    return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *EntityRegisteredEventAllOf) HasEntity() bool {
	if o != nil && !isNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given GenericEntity and assigns it to the Entity field.
func (o *EntityRegisteredEventAllOf) SetEntity(v GenericEntity) {
	o.Entity = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EntityRegisteredEventAllOf) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRegisteredEventAllOf) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EntityRegisteredEventAllOf) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *EntityRegisteredEventAllOf) SetEventType(v string) {
	o.EventType = &v
}

func (o EntityRegisteredEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return json.Marshal(toSerialize)
}

type NullableEntityRegisteredEventAllOf struct {
	value *EntityRegisteredEventAllOf
	isSet bool
}

func (v NullableEntityRegisteredEventAllOf) Get() *EntityRegisteredEventAllOf {
	return v.value
}

func (v *NullableEntityRegisteredEventAllOf) Set(val *EntityRegisteredEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRegisteredEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRegisteredEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRegisteredEventAllOf(val *EntityRegisteredEventAllOf) *NullableEntityRegisteredEventAllOf {
	return &NullableEntityRegisteredEventAllOf{value: val, isSet: true}
}

func (v NullableEntityRegisteredEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRegisteredEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


