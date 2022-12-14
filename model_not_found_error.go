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

// NotFoundError struct for NotFoundError
type NotFoundError struct {
	Error *NotFoundErrorError `json:"error,omitempty"`
}

// NewNotFoundError instantiates a new NotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundError() *NotFoundError {
	this := NotFoundError{}
	return &this
}

// NewNotFoundErrorWithDefaults instantiates a new NotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundErrorWithDefaults() *NotFoundError {
	this := NotFoundError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NotFoundError) GetError() NotFoundErrorError {
	if o == nil || isNil(o.Error) {
		var ret NotFoundErrorError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundError) GetErrorOk() (*NotFoundErrorError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NotFoundError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given NotFoundErrorError and assigns it to the Error field.
func (o *NotFoundError) SetError(v NotFoundErrorError) {
	o.Error = &v
}

func (o NotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableNotFoundError struct {
	value *NotFoundError
	isSet bool
}

func (v NullableNotFoundError) Get() *NotFoundError {
	return v.value
}

func (v *NullableNotFoundError) Set(val *NotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundError(val *NotFoundError) *NullableNotFoundError {
	return &NullableNotFoundError{value: val, isSet: true}
}

func (v NullableNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


