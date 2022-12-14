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

// ConflictError struct for ConflictError
type ConflictError struct {
	Error *ConflictErrorError `json:"error,omitempty"`
}

// NewConflictError instantiates a new ConflictError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictError() *ConflictError {
	this := ConflictError{}
	return &this
}

// NewConflictErrorWithDefaults instantiates a new ConflictError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictErrorWithDefaults() *ConflictError {
	this := ConflictError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConflictError) GetError() ConflictErrorError {
	if o == nil || isNil(o.Error) {
		var ret ConflictErrorError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictError) GetErrorOk() (*ConflictErrorError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConflictError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ConflictErrorError and assigns it to the Error field.
func (o *ConflictError) SetError(v ConflictErrorError) {
	o.Error = &v
}

func (o ConflictError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableConflictError struct {
	value *ConflictError
	isSet bool
}

func (v NullableConflictError) Get() *ConflictError {
	return v.value
}

func (v *NullableConflictError) Set(val *ConflictError) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictError) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictError(val *ConflictError) *NullableConflictError {
	return &NullableConflictError{value: val, isSet: true}
}

func (v NullableConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


