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

// DefaultConcentrationSummary struct for DefaultConcentrationSummary
type DefaultConcentrationSummary struct {
	// Units of measurement.
	Units *string `json:"units,omitempty"`
	// Amount of measurement.
	Value *float32 `json:"value,omitempty"`
}

// NewDefaultConcentrationSummary instantiates a new DefaultConcentrationSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultConcentrationSummary() *DefaultConcentrationSummary {
	this := DefaultConcentrationSummary{}
	return &this
}

// NewDefaultConcentrationSummaryWithDefaults instantiates a new DefaultConcentrationSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultConcentrationSummaryWithDefaults() *DefaultConcentrationSummary {
	this := DefaultConcentrationSummary{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *DefaultConcentrationSummary) GetUnits() string {
	if o == nil || isNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConcentrationSummary) GetUnitsOk() (*string, bool) {
	if o == nil || isNil(o.Units) {
    return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *DefaultConcentrationSummary) HasUnits() bool {
	if o != nil && !isNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *DefaultConcentrationSummary) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DefaultConcentrationSummary) GetValue() float32 {
	if o == nil || isNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultConcentrationSummary) GetValueOk() (*float32, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DefaultConcentrationSummary) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *DefaultConcentrationSummary) SetValue(v float32) {
	o.Value = &v
}

func (o DefaultConcentrationSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultConcentrationSummary struct {
	value *DefaultConcentrationSummary
	isSet bool
}

func (v NullableDefaultConcentrationSummary) Get() *DefaultConcentrationSummary {
	return v.value
}

func (v *NullableDefaultConcentrationSummary) Set(val *DefaultConcentrationSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultConcentrationSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultConcentrationSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultConcentrationSummary(val *DefaultConcentrationSummary) *NullableDefaultConcentrationSummary {
	return &NullableDefaultConcentrationSummary{value: val, isSet: true}
}

func (v NullableDefaultConcentrationSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultConcentrationSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


