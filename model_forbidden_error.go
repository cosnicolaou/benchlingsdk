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

// ForbiddenError struct for ForbiddenError
type ForbiddenError struct {
	Error *ForbiddenErrorError `json:"error,omitempty"`
}

// NewForbiddenError instantiates a new ForbiddenError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForbiddenError() *ForbiddenError {
	this := ForbiddenError{}
	return &this
}

// NewForbiddenErrorWithDefaults instantiates a new ForbiddenError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForbiddenErrorWithDefaults() *ForbiddenError {
	this := ForbiddenError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ForbiddenError) GetError() ForbiddenErrorError {
	if o == nil || isNil(o.Error) {
		var ret ForbiddenErrorError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForbiddenError) GetErrorOk() (*ForbiddenErrorError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ForbiddenError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ForbiddenErrorError and assigns it to the Error field.
func (o *ForbiddenError) SetError(v ForbiddenErrorError) {
	o.Error = &v
}

func (o ForbiddenError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableForbiddenError struct {
	value *ForbiddenError
	isSet bool
}

func (v NullableForbiddenError) Get() *ForbiddenError {
	return v.value
}

func (v *NullableForbiddenError) Set(val *ForbiddenError) {
	v.value = val
	v.isSet = true
}

func (v NullableForbiddenError) IsSet() bool {
	return v.isSet
}

func (v *NullableForbiddenError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForbiddenError(val *ForbiddenError) *NullableForbiddenError {
	return &NullableForbiddenError{value: val, isSet: true}
}

func (v NullableForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForbiddenError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


