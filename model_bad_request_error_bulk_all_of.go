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

// BadRequestErrorBulkAllOf struct for BadRequestErrorBulkAllOf
type BadRequestErrorBulkAllOf struct {
	Error *BadRequestErrorBulkAllOfError `json:"error,omitempty"`
}

// NewBadRequestErrorBulkAllOf instantiates a new BadRequestErrorBulkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestErrorBulkAllOf() *BadRequestErrorBulkAllOf {
	this := BadRequestErrorBulkAllOf{}
	return &this
}

// NewBadRequestErrorBulkAllOfWithDefaults instantiates a new BadRequestErrorBulkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestErrorBulkAllOfWithDefaults() *BadRequestErrorBulkAllOf {
	this := BadRequestErrorBulkAllOf{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BadRequestErrorBulkAllOf) GetError() BadRequestErrorBulkAllOfError {
	if o == nil || isNil(o.Error) {
		var ret BadRequestErrorBulkAllOfError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestErrorBulkAllOf) GetErrorOk() (*BadRequestErrorBulkAllOfError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BadRequestErrorBulkAllOf) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given BadRequestErrorBulkAllOfError and assigns it to the Error field.
func (o *BadRequestErrorBulkAllOf) SetError(v BadRequestErrorBulkAllOfError) {
	o.Error = &v
}

func (o BadRequestErrorBulkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBadRequestErrorBulkAllOf struct {
	value *BadRequestErrorBulkAllOf
	isSet bool
}

func (v NullableBadRequestErrorBulkAllOf) Get() *BadRequestErrorBulkAllOf {
	return v.value
}

func (v *NullableBadRequestErrorBulkAllOf) Set(val *BadRequestErrorBulkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestErrorBulkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestErrorBulkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestErrorBulkAllOf(val *BadRequestErrorBulkAllOf) *NullableBadRequestErrorBulkAllOf {
	return &NullableBadRequestErrorBulkAllOf{value: val, isSet: true}
}

func (v NullableBadRequestErrorBulkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestErrorBulkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


