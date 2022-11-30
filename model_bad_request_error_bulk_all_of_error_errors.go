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

// BadRequestErrorBulkAllOfErrorErrors struct for BadRequestErrorBulkAllOfErrorErrors
type BadRequestErrorBulkAllOfErrorErrors struct {
	Index *float32 `json:"index,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewBadRequestErrorBulkAllOfErrorErrors instantiates a new BadRequestErrorBulkAllOfErrorErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestErrorBulkAllOfErrorErrors() *BadRequestErrorBulkAllOfErrorErrors {
	this := BadRequestErrorBulkAllOfErrorErrors{}
	return &this
}

// NewBadRequestErrorBulkAllOfErrorErrorsWithDefaults instantiates a new BadRequestErrorBulkAllOfErrorErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestErrorBulkAllOfErrorErrorsWithDefaults() *BadRequestErrorBulkAllOfErrorErrors {
	this := BadRequestErrorBulkAllOfErrorErrors{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BadRequestErrorBulkAllOfErrorErrors) GetIndex() float32 {
	if o == nil || isNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestErrorBulkAllOfErrorErrors) GetIndexOk() (*float32, bool) {
	if o == nil || isNil(o.Index) {
    return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BadRequestErrorBulkAllOfErrorErrors) HasIndex() bool {
	if o != nil && !isNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *BadRequestErrorBulkAllOfErrorErrors) SetIndex(v float32) {
	o.Index = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BadRequestErrorBulkAllOfErrorErrors) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestErrorBulkAllOfErrorErrors) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BadRequestErrorBulkAllOfErrorErrors) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BadRequestErrorBulkAllOfErrorErrors) SetMessage(v string) {
	o.Message = &v
}

func (o BadRequestErrorBulkAllOfErrorErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableBadRequestErrorBulkAllOfErrorErrors struct {
	value *BadRequestErrorBulkAllOfErrorErrors
	isSet bool
}

func (v NullableBadRequestErrorBulkAllOfErrorErrors) Get() *BadRequestErrorBulkAllOfErrorErrors {
	return v.value
}

func (v *NullableBadRequestErrorBulkAllOfErrorErrors) Set(val *BadRequestErrorBulkAllOfErrorErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestErrorBulkAllOfErrorErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestErrorBulkAllOfErrorErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestErrorBulkAllOfErrorErrors(val *BadRequestErrorBulkAllOfErrorErrors) *NullableBadRequestErrorBulkAllOfErrorErrors {
	return &NullableBadRequestErrorBulkAllOfErrorErrors{value: val, isSet: true}
}

func (v NullableBadRequestErrorBulkAllOfErrorErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestErrorBulkAllOfErrorErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

