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

// FeatureLibraryUpdate Inputs for updating a feature library
type FeatureLibraryUpdate struct {
	// The description for the feature library
	Description *string `json:"description,omitempty"`
	// The name of the feature library
	Name *string `json:"name,omitempty"`
}

// NewFeatureLibraryUpdate instantiates a new FeatureLibraryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureLibraryUpdate() *FeatureLibraryUpdate {
	this := FeatureLibraryUpdate{}
	return &this
}

// NewFeatureLibraryUpdateWithDefaults instantiates a new FeatureLibraryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureLibraryUpdateWithDefaults() *FeatureLibraryUpdate {
	this := FeatureLibraryUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureLibraryUpdate) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLibraryUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureLibraryUpdate) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureLibraryUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeatureLibraryUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLibraryUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeatureLibraryUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeatureLibraryUpdate) SetName(v string) {
	o.Name = &v
}

func (o FeatureLibraryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureLibraryUpdate struct {
	value *FeatureLibraryUpdate
	isSet bool
}

func (v NullableFeatureLibraryUpdate) Get() *FeatureLibraryUpdate {
	return v.value
}

func (v *NullableFeatureLibraryUpdate) Set(val *FeatureLibraryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureLibraryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureLibraryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureLibraryUpdate(val *FeatureLibraryUpdate) *NullableFeatureLibraryUpdate {
	return &NullableFeatureLibraryUpdate{value: val, isSet: true}
}

func (v NullableFeatureLibraryUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureLibraryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


