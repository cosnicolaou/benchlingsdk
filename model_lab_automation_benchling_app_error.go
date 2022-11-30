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

// LabAutomationBenchlingAppError struct for LabAutomationBenchlingAppError
type LabAutomationBenchlingAppError struct {
	Message *string `json:"message,omitempty"`
}

// NewLabAutomationBenchlingAppError instantiates a new LabAutomationBenchlingAppError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabAutomationBenchlingAppError() *LabAutomationBenchlingAppError {
	this := LabAutomationBenchlingAppError{}
	return &this
}

// NewLabAutomationBenchlingAppErrorWithDefaults instantiates a new LabAutomationBenchlingAppError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabAutomationBenchlingAppErrorWithDefaults() *LabAutomationBenchlingAppError {
	this := LabAutomationBenchlingAppError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LabAutomationBenchlingAppError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationBenchlingAppError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LabAutomationBenchlingAppError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LabAutomationBenchlingAppError) SetMessage(v string) {
	o.Message = &v
}

func (o LabAutomationBenchlingAppError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableLabAutomationBenchlingAppError struct {
	value *LabAutomationBenchlingAppError
	isSet bool
}

func (v NullableLabAutomationBenchlingAppError) Get() *LabAutomationBenchlingAppError {
	return v.value
}

func (v *NullableLabAutomationBenchlingAppError) Set(val *LabAutomationBenchlingAppError) {
	v.value = val
	v.isSet = true
}

func (v NullableLabAutomationBenchlingAppError) IsSet() bool {
	return v.isSet
}

func (v *NullableLabAutomationBenchlingAppError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabAutomationBenchlingAppError(val *LabAutomationBenchlingAppError) *NullableLabAutomationBenchlingAppError {
	return &NullableLabAutomationBenchlingAppError{value: val, isSet: true}
}

func (v NullableLabAutomationBenchlingAppError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabAutomationBenchlingAppError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


