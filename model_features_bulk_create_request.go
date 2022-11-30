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

// FeaturesBulkCreateRequest Inputs for bulk creating a new feature
type FeaturesBulkCreateRequest struct {
	Features []FeatureBulkCreate `json:"features,omitempty"`
}

// NewFeaturesBulkCreateRequest instantiates a new FeaturesBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesBulkCreateRequest() *FeaturesBulkCreateRequest {
	this := FeaturesBulkCreateRequest{}
	return &this
}

// NewFeaturesBulkCreateRequestWithDefaults instantiates a new FeaturesBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesBulkCreateRequestWithDefaults() *FeaturesBulkCreateRequest {
	this := FeaturesBulkCreateRequest{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FeaturesBulkCreateRequest) GetFeatures() []FeatureBulkCreate {
	if o == nil || isNil(o.Features) {
		var ret []FeatureBulkCreate
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesBulkCreateRequest) GetFeaturesOk() ([]FeatureBulkCreate, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FeaturesBulkCreateRequest) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []FeatureBulkCreate and assigns it to the Features field.
func (o *FeaturesBulkCreateRequest) SetFeatures(v []FeatureBulkCreate) {
	o.Features = v
}

func (o FeaturesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableFeaturesBulkCreateRequest struct {
	value *FeaturesBulkCreateRequest
	isSet bool
}

func (v NullableFeaturesBulkCreateRequest) Get() *FeaturesBulkCreateRequest {
	return v.value
}

func (v *NullableFeaturesBulkCreateRequest) Set(val *FeaturesBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesBulkCreateRequest(val *FeaturesBulkCreateRequest) *NullableFeaturesBulkCreateRequest {
	return &NullableFeaturesBulkCreateRequest{value: val, isSet: true}
}

func (v NullableFeaturesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

