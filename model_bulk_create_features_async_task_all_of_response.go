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

// BulkCreateFeaturesAsyncTaskAllOfResponse struct for BulkCreateFeaturesAsyncTaskAllOfResponse
type BulkCreateFeaturesAsyncTaskAllOfResponse struct {
	Features []Feature `json:"features,omitempty"`
}

// NewBulkCreateFeaturesAsyncTaskAllOfResponse instantiates a new BulkCreateFeaturesAsyncTaskAllOfResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCreateFeaturesAsyncTaskAllOfResponse() *BulkCreateFeaturesAsyncTaskAllOfResponse {
	this := BulkCreateFeaturesAsyncTaskAllOfResponse{}
	return &this
}

// NewBulkCreateFeaturesAsyncTaskAllOfResponseWithDefaults instantiates a new BulkCreateFeaturesAsyncTaskAllOfResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCreateFeaturesAsyncTaskAllOfResponseWithDefaults() *BulkCreateFeaturesAsyncTaskAllOfResponse {
	this := BulkCreateFeaturesAsyncTaskAllOfResponse{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BulkCreateFeaturesAsyncTaskAllOfResponse) GetFeatures() []Feature {
	if o == nil || isNil(o.Features) {
		var ret []Feature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateFeaturesAsyncTaskAllOfResponse) GetFeaturesOk() ([]Feature, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BulkCreateFeaturesAsyncTaskAllOfResponse) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []Feature and assigns it to the Features field.
func (o *BulkCreateFeaturesAsyncTaskAllOfResponse) SetFeatures(v []Feature) {
	o.Features = v
}

func (o BulkCreateFeaturesAsyncTaskAllOfResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableBulkCreateFeaturesAsyncTaskAllOfResponse struct {
	value *BulkCreateFeaturesAsyncTaskAllOfResponse
	isSet bool
}

func (v NullableBulkCreateFeaturesAsyncTaskAllOfResponse) Get() *BulkCreateFeaturesAsyncTaskAllOfResponse {
	return v.value
}

func (v *NullableBulkCreateFeaturesAsyncTaskAllOfResponse) Set(val *BulkCreateFeaturesAsyncTaskAllOfResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCreateFeaturesAsyncTaskAllOfResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCreateFeaturesAsyncTaskAllOfResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCreateFeaturesAsyncTaskAllOfResponse(val *BulkCreateFeaturesAsyncTaskAllOfResponse) *NullableBulkCreateFeaturesAsyncTaskAllOfResponse {
	return &NullableBulkCreateFeaturesAsyncTaskAllOfResponse{value: val, isSet: true}
}

func (v NullableBulkCreateFeaturesAsyncTaskAllOfResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCreateFeaturesAsyncTaskAllOfResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


