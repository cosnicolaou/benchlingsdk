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

// ContainerQuantity Quantity of a container, well, or transfer. Supports mass, volume, and other quantities.
type ContainerQuantity struct {
	Units NullableString `json:"units,omitempty"`
	Value NullableFloat32 `json:"value,omitempty"`
}

// NewContainerQuantity instantiates a new ContainerQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerQuantity() *ContainerQuantity {
	this := ContainerQuantity{}
	return &this
}

// NewContainerQuantityWithDefaults instantiates a new ContainerQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerQuantityWithDefaults() *ContainerQuantity {
	this := ContainerQuantity{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerQuantity) GetUnits() string {
	if o == nil || isNil(o.Units.Get()) {
		var ret string
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerQuantity) GetUnitsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *ContainerQuantity) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableString and assigns it to the Units field.
func (o *ContainerQuantity) SetUnits(v string) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *ContainerQuantity) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *ContainerQuantity) UnsetUnits() {
	o.Units.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerQuantity) GetValue() float32 {
	if o == nil || isNil(o.Value.Get()) {
		var ret float32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerQuantity) GetValueOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ContainerQuantity) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat32 and assigns it to the Value field.
func (o *ContainerQuantity) SetValue(v float32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ContainerQuantity) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ContainerQuantity) UnsetValue() {
	o.Value.Unset()
}

func (o ContainerQuantity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Units.IsSet() {
		toSerialize["units"] = o.Units.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContainerQuantity struct {
	value *ContainerQuantity
	isSet bool
}

func (v NullableContainerQuantity) Get() *ContainerQuantity {
	return v.value
}

func (v *NullableContainerQuantity) Set(val *ContainerQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerQuantity(val *ContainerQuantity) *NullableContainerQuantity {
	return &NullableContainerQuantity{value: val, isSet: true}
}

func (v NullableContainerQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


