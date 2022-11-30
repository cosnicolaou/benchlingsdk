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

// LabAutomationBenchlingAppErrors struct for LabAutomationBenchlingAppErrors
type LabAutomationBenchlingAppErrors struct {
	TopLevelErrors []LabAutomationBenchlingAppErrorsTopLevelErrorsInner `json:"topLevelErrors,omitempty"`
}

// NewLabAutomationBenchlingAppErrors instantiates a new LabAutomationBenchlingAppErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabAutomationBenchlingAppErrors() *LabAutomationBenchlingAppErrors {
	this := LabAutomationBenchlingAppErrors{}
	return &this
}

// NewLabAutomationBenchlingAppErrorsWithDefaults instantiates a new LabAutomationBenchlingAppErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabAutomationBenchlingAppErrorsWithDefaults() *LabAutomationBenchlingAppErrors {
	this := LabAutomationBenchlingAppErrors{}
	return &this
}

// GetTopLevelErrors returns the TopLevelErrors field value if set, zero value otherwise.
func (o *LabAutomationBenchlingAppErrors) GetTopLevelErrors() []LabAutomationBenchlingAppErrorsTopLevelErrorsInner {
	if o == nil || isNil(o.TopLevelErrors) {
		var ret []LabAutomationBenchlingAppErrorsTopLevelErrorsInner
		return ret
	}
	return o.TopLevelErrors
}

// GetTopLevelErrorsOk returns a tuple with the TopLevelErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationBenchlingAppErrors) GetTopLevelErrorsOk() ([]LabAutomationBenchlingAppErrorsTopLevelErrorsInner, bool) {
	if o == nil || isNil(o.TopLevelErrors) {
    return nil, false
	}
	return o.TopLevelErrors, true
}

// HasTopLevelErrors returns a boolean if a field has been set.
func (o *LabAutomationBenchlingAppErrors) HasTopLevelErrors() bool {
	if o != nil && !isNil(o.TopLevelErrors) {
		return true
	}

	return false
}

// SetTopLevelErrors gets a reference to the given []LabAutomationBenchlingAppErrorsTopLevelErrorsInner and assigns it to the TopLevelErrors field.
func (o *LabAutomationBenchlingAppErrors) SetTopLevelErrors(v []LabAutomationBenchlingAppErrorsTopLevelErrorsInner) {
	o.TopLevelErrors = v
}

func (o LabAutomationBenchlingAppErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TopLevelErrors) {
		toSerialize["topLevelErrors"] = o.TopLevelErrors
	}
	return json.Marshal(toSerialize)
}

type NullableLabAutomationBenchlingAppErrors struct {
	value *LabAutomationBenchlingAppErrors
	isSet bool
}

func (v NullableLabAutomationBenchlingAppErrors) Get() *LabAutomationBenchlingAppErrors {
	return v.value
}

func (v *NullableLabAutomationBenchlingAppErrors) Set(val *LabAutomationBenchlingAppErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableLabAutomationBenchlingAppErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableLabAutomationBenchlingAppErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabAutomationBenchlingAppErrors(val *LabAutomationBenchlingAppErrors) *NullableLabAutomationBenchlingAppErrors {
	return &NullableLabAutomationBenchlingAppErrors{value: val, isSet: true}
}

func (v NullableLabAutomationBenchlingAppErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabAutomationBenchlingAppErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


