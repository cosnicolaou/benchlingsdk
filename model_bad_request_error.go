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

// BadRequestError struct for BadRequestError
type BadRequestError struct {
	Error *BadRequestErrorError `json:"error,omitempty"`
}

// NewBadRequestError instantiates a new BadRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestError() *BadRequestError {
	this := BadRequestError{}
	return &this
}

// NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestErrorWithDefaults() *BadRequestError {
	this := BadRequestError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BadRequestError) GetError() BadRequestErrorError {
	if o == nil || isNil(o.Error) {
		var ret BadRequestErrorError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestError) GetErrorOk() (*BadRequestErrorError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BadRequestError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given BadRequestErrorError and assigns it to the Error field.
func (o *BadRequestError) SetError(v BadRequestErrorError) {
	o.Error = &v
}

func (o BadRequestError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBadRequestError struct {
	value *BadRequestError
	isSet bool
}

func (v NullableBadRequestError) Get() *BadRequestError {
	return v.value
}

func (v *NullableBadRequestError) Set(val *BadRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestError(val *BadRequestError) *NullableBadRequestError {
	return &NullableBadRequestError{value: val, isSet: true}
}

func (v NullableBadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

