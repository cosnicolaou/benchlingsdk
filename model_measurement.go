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

// Measurement struct for Measurement
type Measurement struct {
	// Can only be null if value is also null
	Units NullableString `json:"units"`
	// Can only be null if units is also null
	Value NullableFloat32 `json:"value"`
}

// NewMeasurement instantiates a new Measurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurement(units NullableString, value NullableFloat32) *Measurement {
	this := Measurement{}
	this.Units = units
	this.Value = value
	return &this
}

// NewMeasurementWithDefaults instantiates a new Measurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementWithDefaults() *Measurement {
	this := Measurement{}
	return &this
}

// GetUnits returns the Units field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Measurement) GetUnits() string {
	if o == nil || o.Units.Get() == nil {
		var ret string
		return ret
	}

	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Measurement) GetUnitsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// SetUnits sets field value
func (o *Measurement) SetUnits(v string) {
	o.Units.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Measurement) GetValue() float32 {
	if o == nil || o.Value.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Measurement) GetValueOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *Measurement) SetValue(v float32) {
	o.Value.Set(&v)
}

func (o Measurement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["units"] = o.Units.Get()
	}
	if true {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMeasurement struct {
	value *Measurement
	isSet bool
}

func (v NullableMeasurement) Get() *Measurement {
	return v.value
}

func (v *NullableMeasurement) Set(val *Measurement) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurement(val *Measurement) *NullableMeasurement {
	return &NullableMeasurement{value: val, isSet: true}
}

func (v NullableMeasurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

