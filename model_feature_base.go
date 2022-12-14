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

// FeatureBase struct for FeatureBase
type FeatureBase struct {
	// The color of the annotations generated by the feature. Must be a valid hex string
	Color *string `json:"color,omitempty"`
	// The id of the feature library the feature belongs to
	FeatureLibraryId *string `json:"featureLibraryId,omitempty"`
	// The type of feature, like gene, promoter, etc. Note: This is an arbitrary string, not an enum 
	FeatureType NullableString `json:"featureType,omitempty"`
	// The name of the feature
	Name *string `json:"name,omitempty"`
	// The pattern used for matching during auto-annotation.
	Pattern *string `json:"pattern,omitempty"`
}

// NewFeatureBase instantiates a new FeatureBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureBase() *FeatureBase {
	this := FeatureBase{}
	return &this
}

// NewFeatureBaseWithDefaults instantiates a new FeatureBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureBaseWithDefaults() *FeatureBase {
	this := FeatureBase{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *FeatureBase) GetColor() string {
	if o == nil || isNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureBase) GetColorOk() (*string, bool) {
	if o == nil || isNil(o.Color) {
    return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *FeatureBase) HasColor() bool {
	if o != nil && !isNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *FeatureBase) SetColor(v string) {
	o.Color = &v
}

// GetFeatureLibraryId returns the FeatureLibraryId field value if set, zero value otherwise.
func (o *FeatureBase) GetFeatureLibraryId() string {
	if o == nil || isNil(o.FeatureLibraryId) {
		var ret string
		return ret
	}
	return *o.FeatureLibraryId
}

// GetFeatureLibraryIdOk returns a tuple with the FeatureLibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureBase) GetFeatureLibraryIdOk() (*string, bool) {
	if o == nil || isNil(o.FeatureLibraryId) {
    return nil, false
	}
	return o.FeatureLibraryId, true
}

// HasFeatureLibraryId returns a boolean if a field has been set.
func (o *FeatureBase) HasFeatureLibraryId() bool {
	if o != nil && !isNil(o.FeatureLibraryId) {
		return true
	}

	return false
}

// SetFeatureLibraryId gets a reference to the given string and assigns it to the FeatureLibraryId field.
func (o *FeatureBase) SetFeatureLibraryId(v string) {
	o.FeatureLibraryId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureBase) GetFeatureType() string {
	if o == nil || isNil(o.FeatureType.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureType.Get()
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureBase) GetFeatureTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FeatureType.Get(), o.FeatureType.IsSet()
}

// HasFeatureType returns a boolean if a field has been set.
func (o *FeatureBase) HasFeatureType() bool {
	if o != nil && o.FeatureType.IsSet() {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given NullableString and assigns it to the FeatureType field.
func (o *FeatureBase) SetFeatureType(v string) {
	o.FeatureType.Set(&v)
}
// SetFeatureTypeNil sets the value for FeatureType to be an explicit nil
func (o *FeatureBase) SetFeatureTypeNil() {
	o.FeatureType.Set(nil)
}

// UnsetFeatureType ensures that no value is present for FeatureType, not even an explicit nil
func (o *FeatureBase) UnsetFeatureType() {
	o.FeatureType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeatureBase) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureBase) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeatureBase) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeatureBase) SetName(v string) {
	o.Name = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *FeatureBase) GetPattern() string {
	if o == nil || isNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureBase) GetPatternOk() (*string, bool) {
	if o == nil || isNil(o.Pattern) {
    return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *FeatureBase) HasPattern() bool {
	if o != nil && !isNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *FeatureBase) SetPattern(v string) {
	o.Pattern = &v
}

func (o FeatureBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !isNil(o.FeatureLibraryId) {
		toSerialize["featureLibraryId"] = o.FeatureLibraryId
	}
	if o.FeatureType.IsSet() {
		toSerialize["featureType"] = o.FeatureType.Get()
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureBase struct {
	value *FeatureBase
	isSet bool
}

func (v NullableFeatureBase) Get() *FeatureBase {
	return v.value
}

func (v *NullableFeatureBase) Set(val *FeatureBase) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureBase) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureBase(val *FeatureBase) *NullableFeatureBase {
	return &NullableFeatureBase{value: val, isSet: true}
}

func (v NullableFeatureBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


