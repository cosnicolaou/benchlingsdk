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

// DropdownFieldDefinitionAllOf struct for DropdownFieldDefinitionAllOf
type DropdownFieldDefinitionAllOf struct {
	DropdownId NullableString `json:"dropdownId,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDropdownFieldDefinitionAllOf instantiates a new DropdownFieldDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropdownFieldDefinitionAllOf() *DropdownFieldDefinitionAllOf {
	this := DropdownFieldDefinitionAllOf{}
	return &this
}

// NewDropdownFieldDefinitionAllOfWithDefaults instantiates a new DropdownFieldDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropdownFieldDefinitionAllOfWithDefaults() *DropdownFieldDefinitionAllOf {
	this := DropdownFieldDefinitionAllOf{}
	return &this
}

// GetDropdownId returns the DropdownId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DropdownFieldDefinitionAllOf) GetDropdownId() string {
	if o == nil || isNil(o.DropdownId.Get()) {
		var ret string
		return ret
	}
	return *o.DropdownId.Get()
}

// GetDropdownIdOk returns a tuple with the DropdownId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DropdownFieldDefinitionAllOf) GetDropdownIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DropdownId.Get(), o.DropdownId.IsSet()
}

// HasDropdownId returns a boolean if a field has been set.
func (o *DropdownFieldDefinitionAllOf) HasDropdownId() bool {
	if o != nil && o.DropdownId.IsSet() {
		return true
	}

	return false
}

// SetDropdownId gets a reference to the given NullableString and assigns it to the DropdownId field.
func (o *DropdownFieldDefinitionAllOf) SetDropdownId(v string) {
	o.DropdownId.Set(&v)
}
// SetDropdownIdNil sets the value for DropdownId to be an explicit nil
func (o *DropdownFieldDefinitionAllOf) SetDropdownIdNil() {
	o.DropdownId.Set(nil)
}

// UnsetDropdownId ensures that no value is present for DropdownId, not even an explicit nil
func (o *DropdownFieldDefinitionAllOf) UnsetDropdownId() {
	o.DropdownId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DropdownFieldDefinitionAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropdownFieldDefinitionAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DropdownFieldDefinitionAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DropdownFieldDefinitionAllOf) SetType(v string) {
	o.Type = &v
}

func (o DropdownFieldDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DropdownId.IsSet() {
		toSerialize["dropdownId"] = o.DropdownId.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDropdownFieldDefinitionAllOf struct {
	value *DropdownFieldDefinitionAllOf
	isSet bool
}

func (v NullableDropdownFieldDefinitionAllOf) Get() *DropdownFieldDefinitionAllOf {
	return v.value
}

func (v *NullableDropdownFieldDefinitionAllOf) Set(val *DropdownFieldDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDropdownFieldDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDropdownFieldDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropdownFieldDefinitionAllOf(val *DropdownFieldDefinitionAllOf) *NullableDropdownFieldDefinitionAllOf {
	return &NullableDropdownFieldDefinitionAllOf{value: val, isSet: true}
}

func (v NullableDropdownFieldDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropdownFieldDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

