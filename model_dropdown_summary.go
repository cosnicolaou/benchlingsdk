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

// DropdownSummary struct for DropdownSummary
type DropdownSummary struct {
	// ID of the dropdown
	Id *string `json:"id,omitempty"`
	// Name of the dropdown
	Name *string `json:"name,omitempty"`
}

// NewDropdownSummary instantiates a new DropdownSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropdownSummary() *DropdownSummary {
	this := DropdownSummary{}
	return &this
}

// NewDropdownSummaryWithDefaults instantiates a new DropdownSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropdownSummaryWithDefaults() *DropdownSummary {
	this := DropdownSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DropdownSummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropdownSummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DropdownSummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DropdownSummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DropdownSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropdownSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DropdownSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DropdownSummary) SetName(v string) {
	o.Name = &v
}

func (o DropdownSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDropdownSummary struct {
	value *DropdownSummary
	isSet bool
}

func (v NullableDropdownSummary) Get() *DropdownSummary {
	return v.value
}

func (v *NullableDropdownSummary) Set(val *DropdownSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDropdownSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDropdownSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropdownSummary(val *DropdownSummary) *NullableDropdownSummary {
	return &NullableDropdownSummary{value: val, isSet: true}
}

func (v NullableDropdownSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropdownSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


