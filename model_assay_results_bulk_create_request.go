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

// AssayResultsBulkCreateRequest struct for AssayResultsBulkCreateRequest
type AssayResultsBulkCreateRequest struct {
	AssayResults []AssayResultCreate `json:"assayResults"`
}

// NewAssayResultsBulkCreateRequest instantiates a new AssayResultsBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResultsBulkCreateRequest(assayResults []AssayResultCreate) *AssayResultsBulkCreateRequest {
	this := AssayResultsBulkCreateRequest{}
	this.AssayResults = assayResults
	return &this
}

// NewAssayResultsBulkCreateRequestWithDefaults instantiates a new AssayResultsBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultsBulkCreateRequestWithDefaults() *AssayResultsBulkCreateRequest {
	this := AssayResultsBulkCreateRequest{}
	return &this
}

// GetAssayResults returns the AssayResults field value
func (o *AssayResultsBulkCreateRequest) GetAssayResults() []AssayResultCreate {
	if o == nil {
		var ret []AssayResultCreate
		return ret
	}

	return o.AssayResults
}

// GetAssayResultsOk returns a tuple with the AssayResults field value
// and a boolean to check if the value has been set.
func (o *AssayResultsBulkCreateRequest) GetAssayResultsOk() ([]AssayResultCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.AssayResults, true
}

// SetAssayResults sets field value
func (o *AssayResultsBulkCreateRequest) SetAssayResults(v []AssayResultCreate) {
	o.AssayResults = v
}

func (o AssayResultsBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assayResults"] = o.AssayResults
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResultsBulkCreateRequest struct {
	value *AssayResultsBulkCreateRequest
	isSet bool
}

func (v NullableAssayResultsBulkCreateRequest) Get() *AssayResultsBulkCreateRequest {
	return v.value
}

func (v *NullableAssayResultsBulkCreateRequest) Set(val *AssayResultsBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResultsBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResultsBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResultsBulkCreateRequest(val *AssayResultsBulkCreateRequest) *NullableAssayResultsBulkCreateRequest {
	return &NullableAssayResultsBulkCreateRequest{value: val, isSet: true}
}

func (v NullableAssayResultsBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResultsBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


