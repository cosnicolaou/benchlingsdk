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

// AssayRunsBulkCreateRequest struct for AssayRunsBulkCreateRequest
type AssayRunsBulkCreateRequest struct {
	AssayRuns []AssayRunCreate `json:"assayRuns"`
}

// NewAssayRunsBulkCreateRequest instantiates a new AssayRunsBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayRunsBulkCreateRequest(assayRuns []AssayRunCreate) *AssayRunsBulkCreateRequest {
	this := AssayRunsBulkCreateRequest{}
	this.AssayRuns = assayRuns
	return &this
}

// NewAssayRunsBulkCreateRequestWithDefaults instantiates a new AssayRunsBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayRunsBulkCreateRequestWithDefaults() *AssayRunsBulkCreateRequest {
	this := AssayRunsBulkCreateRequest{}
	return &this
}

// GetAssayRuns returns the AssayRuns field value
func (o *AssayRunsBulkCreateRequest) GetAssayRuns() []AssayRunCreate {
	if o == nil {
		var ret []AssayRunCreate
		return ret
	}

	return o.AssayRuns
}

// GetAssayRunsOk returns a tuple with the AssayRuns field value
// and a boolean to check if the value has been set.
func (o *AssayRunsBulkCreateRequest) GetAssayRunsOk() ([]AssayRunCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.AssayRuns, true
}

// SetAssayRuns sets field value
func (o *AssayRunsBulkCreateRequest) SetAssayRuns(v []AssayRunCreate) {
	o.AssayRuns = v
}

func (o AssayRunsBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assayRuns"] = o.AssayRuns
	}
	return json.Marshal(toSerialize)
}

type NullableAssayRunsBulkCreateRequest struct {
	value *AssayRunsBulkCreateRequest
	isSet bool
}

func (v NullableAssayRunsBulkCreateRequest) Get() *AssayRunsBulkCreateRequest {
	return v.value
}

func (v *NullableAssayRunsBulkCreateRequest) Set(val *AssayRunsBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayRunsBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayRunsBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayRunsBulkCreateRequest(val *AssayRunsBulkCreateRequest) *NullableAssayRunsBulkCreateRequest {
	return &NullableAssayRunsBulkCreateRequest{value: val, isSet: true}
}

func (v NullableAssayRunsBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayRunsBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


