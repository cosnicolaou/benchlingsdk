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

// AssayResultsBulkGet struct for AssayResultsBulkGet
type AssayResultsBulkGet struct {
	AssayResults []AssayResult `json:"assayResults,omitempty"`
}

// NewAssayResultsBulkGet instantiates a new AssayResultsBulkGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResultsBulkGet() *AssayResultsBulkGet {
	this := AssayResultsBulkGet{}
	return &this
}

// NewAssayResultsBulkGetWithDefaults instantiates a new AssayResultsBulkGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultsBulkGetWithDefaults() *AssayResultsBulkGet {
	this := AssayResultsBulkGet{}
	return &this
}

// GetAssayResults returns the AssayResults field value if set, zero value otherwise.
func (o *AssayResultsBulkGet) GetAssayResults() []AssayResult {
	if o == nil || isNil(o.AssayResults) {
		var ret []AssayResult
		return ret
	}
	return o.AssayResults
}

// GetAssayResultsOk returns a tuple with the AssayResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultsBulkGet) GetAssayResultsOk() ([]AssayResult, bool) {
	if o == nil || isNil(o.AssayResults) {
    return nil, false
	}
	return o.AssayResults, true
}

// HasAssayResults returns a boolean if a field has been set.
func (o *AssayResultsBulkGet) HasAssayResults() bool {
	if o != nil && !isNil(o.AssayResults) {
		return true
	}

	return false
}

// SetAssayResults gets a reference to the given []AssayResult and assigns it to the AssayResults field.
func (o *AssayResultsBulkGet) SetAssayResults(v []AssayResult) {
	o.AssayResults = v
}

func (o AssayResultsBulkGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssayResults) {
		toSerialize["assayResults"] = o.AssayResults
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResultsBulkGet struct {
	value *AssayResultsBulkGet
	isSet bool
}

func (v NullableAssayResultsBulkGet) Get() *AssayResultsBulkGet {
	return v.value
}

func (v *NullableAssayResultsBulkGet) Set(val *AssayResultsBulkGet) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResultsBulkGet) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResultsBulkGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResultsBulkGet(val *AssayResultsBulkGet) *NullableAssayResultsBulkGet {
	return &NullableAssayResultsBulkGet{value: val, isSet: true}
}

func (v NullableAssayResultsBulkGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResultsBulkGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


