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

// CustomEntitySummary struct for CustomEntitySummary
type CustomEntitySummary struct {
	EntityType *string `json:"entityType,omitempty"`
	Id *string `json:"id,omitempty"`
	// Deprecated
	Type *string `json:"type,omitempty"`
}

// NewCustomEntitySummary instantiates a new CustomEntitySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEntitySummary() *CustomEntitySummary {
	this := CustomEntitySummary{}
	return &this
}

// NewCustomEntitySummaryWithDefaults instantiates a new CustomEntitySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEntitySummaryWithDefaults() *CustomEntitySummary {
	this := CustomEntitySummary{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CustomEntitySummary) GetEntityType() string {
	if o == nil || isNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntitySummary) GetEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.EntityType) {
    return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CustomEntitySummary) HasEntityType() bool {
	if o != nil && !isNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CustomEntitySummary) SetEntityType(v string) {
	o.EntityType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomEntitySummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntitySummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomEntitySummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomEntitySummary) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
// Deprecated
func (o *CustomEntitySummary) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomEntitySummary) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomEntitySummary) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
// Deprecated
func (o *CustomEntitySummary) SetType(v string) {
	o.Type = &v
}

func (o CustomEntitySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEntitySummary struct {
	value *CustomEntitySummary
	isSet bool
}

func (v NullableCustomEntitySummary) Get() *CustomEntitySummary {
	return v.value
}

func (v *NullableCustomEntitySummary) Set(val *CustomEntitySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEntitySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEntitySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEntitySummary(val *CustomEntitySummary) *NullableCustomEntitySummary {
	return &NullableCustomEntitySummary{value: val, isSet: true}
}

func (v NullableCustomEntitySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEntitySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

