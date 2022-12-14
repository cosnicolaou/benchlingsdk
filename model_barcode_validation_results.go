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

// BarcodeValidationResults struct for BarcodeValidationResults
type BarcodeValidationResults struct {
	ValidationResults []BarcodeValidationResult `json:"validationResults,omitempty"`
}

// NewBarcodeValidationResults instantiates a new BarcodeValidationResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBarcodeValidationResults() *BarcodeValidationResults {
	this := BarcodeValidationResults{}
	return &this
}

// NewBarcodeValidationResultsWithDefaults instantiates a new BarcodeValidationResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBarcodeValidationResultsWithDefaults() *BarcodeValidationResults {
	this := BarcodeValidationResults{}
	return &this
}

// GetValidationResults returns the ValidationResults field value if set, zero value otherwise.
func (o *BarcodeValidationResults) GetValidationResults() []BarcodeValidationResult {
	if o == nil || isNil(o.ValidationResults) {
		var ret []BarcodeValidationResult
		return ret
	}
	return o.ValidationResults
}

// GetValidationResultsOk returns a tuple with the ValidationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarcodeValidationResults) GetValidationResultsOk() ([]BarcodeValidationResult, bool) {
	if o == nil || isNil(o.ValidationResults) {
    return nil, false
	}
	return o.ValidationResults, true
}

// HasValidationResults returns a boolean if a field has been set.
func (o *BarcodeValidationResults) HasValidationResults() bool {
	if o != nil && !isNil(o.ValidationResults) {
		return true
	}

	return false
}

// SetValidationResults gets a reference to the given []BarcodeValidationResult and assigns it to the ValidationResults field.
func (o *BarcodeValidationResults) SetValidationResults(v []BarcodeValidationResult) {
	o.ValidationResults = v
}

func (o BarcodeValidationResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ValidationResults) {
		toSerialize["validationResults"] = o.ValidationResults
	}
	return json.Marshal(toSerialize)
}

type NullableBarcodeValidationResults struct {
	value *BarcodeValidationResults
	isSet bool
}

func (v NullableBarcodeValidationResults) Get() *BarcodeValidationResults {
	return v.value
}

func (v *NullableBarcodeValidationResults) Set(val *BarcodeValidationResults) {
	v.value = val
	v.isSet = true
}

func (v NullableBarcodeValidationResults) IsSet() bool {
	return v.isSet
}

func (v *NullableBarcodeValidationResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBarcodeValidationResults(val *BarcodeValidationResults) *NullableBarcodeValidationResults {
	return &NullableBarcodeValidationResults{value: val, isSet: true}
}

func (v NullableBarcodeValidationResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBarcodeValidationResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


