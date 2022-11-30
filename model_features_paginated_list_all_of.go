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

// FeaturesPaginatedListAllOf struct for FeaturesPaginatedListAllOf
type FeaturesPaginatedListAllOf struct {
	// List of features for the page
	Features []Feature `json:"features,omitempty"`
}

// NewFeaturesPaginatedListAllOf instantiates a new FeaturesPaginatedListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesPaginatedListAllOf() *FeaturesPaginatedListAllOf {
	this := FeaturesPaginatedListAllOf{}
	return &this
}

// NewFeaturesPaginatedListAllOfWithDefaults instantiates a new FeaturesPaginatedListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesPaginatedListAllOfWithDefaults() *FeaturesPaginatedListAllOf {
	this := FeaturesPaginatedListAllOf{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FeaturesPaginatedListAllOf) GetFeatures() []Feature {
	if o == nil || isNil(o.Features) {
		var ret []Feature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesPaginatedListAllOf) GetFeaturesOk() ([]Feature, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FeaturesPaginatedListAllOf) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []Feature and assigns it to the Features field.
func (o *FeaturesPaginatedListAllOf) SetFeatures(v []Feature) {
	o.Features = v
}

func (o FeaturesPaginatedListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableFeaturesPaginatedListAllOf struct {
	value *FeaturesPaginatedListAllOf
	isSet bool
}

func (v NullableFeaturesPaginatedListAllOf) Get() *FeaturesPaginatedListAllOf {
	return v.value
}

func (v *NullableFeaturesPaginatedListAllOf) Set(val *FeaturesPaginatedListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesPaginatedListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesPaginatedListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesPaginatedListAllOf(val *FeaturesPaginatedListAllOf) *NullableFeaturesPaginatedListAllOf {
	return &NullableFeaturesPaginatedListAllOf{value: val, isSet: true}
}

func (v NullableFeaturesPaginatedListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesPaginatedListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

