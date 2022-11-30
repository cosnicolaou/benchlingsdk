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

// UserValidation struct for UserValidation
type UserValidation struct {
	// A string explaining the reason for the provided validation status.
	ValidationComment *string `json:"validationComment,omitempty"`
	// Valid values for this enum depend on whether it is used to set a value (e.g., in a POST request), or is a return value for an existing result. Acceptable values for setting a status are 'VALID' or 'INVALID'. Possible return values are 'VALID', 'INVALID', or 'PARTIALLY_VALID' (a result will be partially valid if it has some valid fields and some invalid fields). 
	ValidationStatus *string `json:"validationStatus,omitempty"`
}

// NewUserValidation instantiates a new UserValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserValidation() *UserValidation {
	this := UserValidation{}
	return &this
}

// NewUserValidationWithDefaults instantiates a new UserValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserValidationWithDefaults() *UserValidation {
	this := UserValidation{}
	return &this
}

// GetValidationComment returns the ValidationComment field value if set, zero value otherwise.
func (o *UserValidation) GetValidationComment() string {
	if o == nil || isNil(o.ValidationComment) {
		var ret string
		return ret
	}
	return *o.ValidationComment
}

// GetValidationCommentOk returns a tuple with the ValidationComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserValidation) GetValidationCommentOk() (*string, bool) {
	if o == nil || isNil(o.ValidationComment) {
    return nil, false
	}
	return o.ValidationComment, true
}

// HasValidationComment returns a boolean if a field has been set.
func (o *UserValidation) HasValidationComment() bool {
	if o != nil && !isNil(o.ValidationComment) {
		return true
	}

	return false
}

// SetValidationComment gets a reference to the given string and assigns it to the ValidationComment field.
func (o *UserValidation) SetValidationComment(v string) {
	o.ValidationComment = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *UserValidation) GetValidationStatus() string {
	if o == nil || isNil(o.ValidationStatus) {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserValidation) GetValidationStatusOk() (*string, bool) {
	if o == nil || isNil(o.ValidationStatus) {
    return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *UserValidation) HasValidationStatus() bool {
	if o != nil && !isNil(o.ValidationStatus) {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *UserValidation) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

func (o UserValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ValidationComment) {
		toSerialize["validationComment"] = o.ValidationComment
	}
	if !isNil(o.ValidationStatus) {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableUserValidation struct {
	value *UserValidation
	isSet bool
}

func (v NullableUserValidation) Get() *UserValidation {
	return v.value
}

func (v *NullableUserValidation) Set(val *UserValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserValidation(val *UserValidation) *NullableUserValidation {
	return &NullableUserValidation{value: val, isSet: true}
}

func (v NullableUserValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

