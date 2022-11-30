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

// RnaOligosBulkCreateRequest struct for RnaOligosBulkCreateRequest
type RnaOligosBulkCreateRequest struct {
	RnaOligos []RnaOligoCreate `json:"rnaOligos,omitempty"`
}

// NewRnaOligosBulkCreateRequest instantiates a new RnaOligosBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaOligosBulkCreateRequest() *RnaOligosBulkCreateRequest {
	this := RnaOligosBulkCreateRequest{}
	return &this
}

// NewRnaOligosBulkCreateRequestWithDefaults instantiates a new RnaOligosBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaOligosBulkCreateRequestWithDefaults() *RnaOligosBulkCreateRequest {
	this := RnaOligosBulkCreateRequest{}
	return &this
}

// GetRnaOligos returns the RnaOligos field value if set, zero value otherwise.
func (o *RnaOligosBulkCreateRequest) GetRnaOligos() []RnaOligoCreate {
	if o == nil || isNil(o.RnaOligos) {
		var ret []RnaOligoCreate
		return ret
	}
	return o.RnaOligos
}

// GetRnaOligosOk returns a tuple with the RnaOligos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligosBulkCreateRequest) GetRnaOligosOk() ([]RnaOligoCreate, bool) {
	if o == nil || isNil(o.RnaOligos) {
    return nil, false
	}
	return o.RnaOligos, true
}

// HasRnaOligos returns a boolean if a field has been set.
func (o *RnaOligosBulkCreateRequest) HasRnaOligos() bool {
	if o != nil && !isNil(o.RnaOligos) {
		return true
	}

	return false
}

// SetRnaOligos gets a reference to the given []RnaOligoCreate and assigns it to the RnaOligos field.
func (o *RnaOligosBulkCreateRequest) SetRnaOligos(v []RnaOligoCreate) {
	o.RnaOligos = v
}

func (o RnaOligosBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RnaOligos) {
		toSerialize["rnaOligos"] = o.RnaOligos
	}
	return json.Marshal(toSerialize)
}

type NullableRnaOligosBulkCreateRequest struct {
	value *RnaOligosBulkCreateRequest
	isSet bool
}

func (v NullableRnaOligosBulkCreateRequest) Get() *RnaOligosBulkCreateRequest {
	return v.value
}

func (v *NullableRnaOligosBulkCreateRequest) Set(val *RnaOligosBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaOligosBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaOligosBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaOligosBulkCreateRequest(val *RnaOligosBulkCreateRequest) *NullableRnaOligosBulkCreateRequest {
	return &NullableRnaOligosBulkCreateRequest{value: val, isSet: true}
}

func (v NullableRnaOligosBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaOligosBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

