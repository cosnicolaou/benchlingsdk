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

// MixturesBulkCreateRequest struct for MixturesBulkCreateRequest
type MixturesBulkCreateRequest struct {
	Mixtures []MixtureCreate `json:"mixtures"`
}

// NewMixturesBulkCreateRequest instantiates a new MixturesBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixturesBulkCreateRequest(mixtures []MixtureCreate) *MixturesBulkCreateRequest {
	this := MixturesBulkCreateRequest{}
	this.Mixtures = mixtures
	return &this
}

// NewMixturesBulkCreateRequestWithDefaults instantiates a new MixturesBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixturesBulkCreateRequestWithDefaults() *MixturesBulkCreateRequest {
	this := MixturesBulkCreateRequest{}
	return &this
}

// GetMixtures returns the Mixtures field value
func (o *MixturesBulkCreateRequest) GetMixtures() []MixtureCreate {
	if o == nil {
		var ret []MixtureCreate
		return ret
	}

	return o.Mixtures
}

// GetMixturesOk returns a tuple with the Mixtures field value
// and a boolean to check if the value has been set.
func (o *MixturesBulkCreateRequest) GetMixturesOk() ([]MixtureCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.Mixtures, true
}

// SetMixtures sets field value
func (o *MixturesBulkCreateRequest) SetMixtures(v []MixtureCreate) {
	o.Mixtures = v
}

func (o MixturesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mixtures"] = o.Mixtures
	}
	return json.Marshal(toSerialize)
}

type NullableMixturesBulkCreateRequest struct {
	value *MixturesBulkCreateRequest
	isSet bool
}

func (v NullableMixturesBulkCreateRequest) Get() *MixturesBulkCreateRequest {
	return v.value
}

func (v *NullableMixturesBulkCreateRequest) Set(val *MixturesBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMixturesBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMixturesBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixturesBulkCreateRequest(val *MixturesBulkCreateRequest) *NullableMixturesBulkCreateRequest {
	return &NullableMixturesBulkCreateRequest{value: val, isSet: true}
}

func (v NullableMixturesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixturesBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


