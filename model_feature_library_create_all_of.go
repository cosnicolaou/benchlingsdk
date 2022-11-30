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

// FeatureLibraryCreateAllOf struct for FeatureLibraryCreateAllOf
type FeatureLibraryCreateAllOf struct {
	// The organization that will own the feature library. The requesting user must be an administrator of the organization. If unspecified and the organization allows personal ownables, then the requesting user will own the feature library 
	OrganizationId *string `json:"organizationId,omitempty"`
}

// NewFeatureLibraryCreateAllOf instantiates a new FeatureLibraryCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureLibraryCreateAllOf() *FeatureLibraryCreateAllOf {
	this := FeatureLibraryCreateAllOf{}
	return &this
}

// NewFeatureLibraryCreateAllOfWithDefaults instantiates a new FeatureLibraryCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureLibraryCreateAllOfWithDefaults() *FeatureLibraryCreateAllOf {
	this := FeatureLibraryCreateAllOf{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *FeatureLibraryCreateAllOf) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLibraryCreateAllOf) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *FeatureLibraryCreateAllOf) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *FeatureLibraryCreateAllOf) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

func (o FeatureLibraryCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureLibraryCreateAllOf struct {
	value *FeatureLibraryCreateAllOf
	isSet bool
}

func (v NullableFeatureLibraryCreateAllOf) Get() *FeatureLibraryCreateAllOf {
	return v.value
}

func (v *NullableFeatureLibraryCreateAllOf) Set(val *FeatureLibraryCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureLibraryCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureLibraryCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureLibraryCreateAllOf(val *FeatureLibraryCreateAllOf) *NullableFeatureLibraryCreateAllOf {
	return &NullableFeatureLibraryCreateAllOf{value: val, isSet: true}
}

func (v NullableFeatureLibraryCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureLibraryCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

