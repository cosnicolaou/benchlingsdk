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

// CustomEntitiesBulkCreateRequest struct for CustomEntitiesBulkCreateRequest
type CustomEntitiesBulkCreateRequest struct {
	CustomEntities []CustomEntityBulkCreate `json:"customEntities"`
}

// NewCustomEntitiesBulkCreateRequest instantiates a new CustomEntitiesBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEntitiesBulkCreateRequest(customEntities []CustomEntityBulkCreate) *CustomEntitiesBulkCreateRequest {
	this := CustomEntitiesBulkCreateRequest{}
	this.CustomEntities = customEntities
	return &this
}

// NewCustomEntitiesBulkCreateRequestWithDefaults instantiates a new CustomEntitiesBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEntitiesBulkCreateRequestWithDefaults() *CustomEntitiesBulkCreateRequest {
	this := CustomEntitiesBulkCreateRequest{}
	return &this
}

// GetCustomEntities returns the CustomEntities field value
func (o *CustomEntitiesBulkCreateRequest) GetCustomEntities() []CustomEntityBulkCreate {
	if o == nil {
		var ret []CustomEntityBulkCreate
		return ret
	}

	return o.CustomEntities
}

// GetCustomEntitiesOk returns a tuple with the CustomEntities field value
// and a boolean to check if the value has been set.
func (o *CustomEntitiesBulkCreateRequest) GetCustomEntitiesOk() ([]CustomEntityBulkCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.CustomEntities, true
}

// SetCustomEntities sets field value
func (o *CustomEntitiesBulkCreateRequest) SetCustomEntities(v []CustomEntityBulkCreate) {
	o.CustomEntities = v
}

func (o CustomEntitiesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customEntities"] = o.CustomEntities
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEntitiesBulkCreateRequest struct {
	value *CustomEntitiesBulkCreateRequest
	isSet bool
}

func (v NullableCustomEntitiesBulkCreateRequest) Get() *CustomEntitiesBulkCreateRequest {
	return v.value
}

func (v *NullableCustomEntitiesBulkCreateRequest) Set(val *CustomEntitiesBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEntitiesBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEntitiesBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEntitiesBulkCreateRequest(val *CustomEntitiesBulkCreateRequest) *NullableCustomEntitiesBulkCreateRequest {
	return &NullableCustomEntitiesBulkCreateRequest{value: val, isSet: true}
}

func (v NullableCustomEntitiesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEntitiesBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


