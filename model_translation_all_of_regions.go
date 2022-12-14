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

// TranslationAllOfRegions struct for TranslationAllOfRegions
type TranslationAllOfRegions struct {
	End *int32 `json:"end,omitempty"`
	Start *int32 `json:"start,omitempty"`
}

// NewTranslationAllOfRegions instantiates a new TranslationAllOfRegions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationAllOfRegions() *TranslationAllOfRegions {
	this := TranslationAllOfRegions{}
	return &this
}

// NewTranslationAllOfRegionsWithDefaults instantiates a new TranslationAllOfRegions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationAllOfRegionsWithDefaults() *TranslationAllOfRegions {
	this := TranslationAllOfRegions{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *TranslationAllOfRegions) GetEnd() int32 {
	if o == nil || isNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationAllOfRegions) GetEndOk() (*int32, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *TranslationAllOfRegions) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *TranslationAllOfRegions) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TranslationAllOfRegions) GetStart() int32 {
	if o == nil || isNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationAllOfRegions) GetStartOk() (*int32, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TranslationAllOfRegions) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *TranslationAllOfRegions) SetStart(v int32) {
	o.Start = &v
}

func (o TranslationAllOfRegions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullableTranslationAllOfRegions struct {
	value *TranslationAllOfRegions
	isSet bool
}

func (v NullableTranslationAllOfRegions) Get() *TranslationAllOfRegions {
	return v.value
}

func (v *NullableTranslationAllOfRegions) Set(val *TranslationAllOfRegions) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationAllOfRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationAllOfRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationAllOfRegions(val *TranslationAllOfRegions) *NullableTranslationAllOfRegions {
	return &NullableTranslationAllOfRegions{value: val, isSet: true}
}

func (v NullableTranslationAllOfRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationAllOfRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


