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

// FeatureAllOf struct for FeatureAllOf
type FeatureAllOf struct {
	// The id of the feature
	Id *string `json:"id,omitempty"`
	// The match type of the feature. Used to determine how auto-annotate matches are made.
	MatchType *string `json:"matchType,omitempty"`
}

// NewFeatureAllOf instantiates a new FeatureAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureAllOf() *FeatureAllOf {
	this := FeatureAllOf{}
	return &this
}

// NewFeatureAllOfWithDefaults instantiates a new FeatureAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureAllOfWithDefaults() *FeatureAllOf {
	this := FeatureAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeatureAllOf) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAllOf) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeatureAllOf) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FeatureAllOf) SetId(v string) {
	o.Id = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *FeatureAllOf) GetMatchType() string {
	if o == nil || isNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureAllOf) GetMatchTypeOk() (*string, bool) {
	if o == nil || isNil(o.MatchType) {
    return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *FeatureAllOf) HasMatchType() bool {
	if o != nil && !isNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *FeatureAllOf) SetMatchType(v string) {
	o.MatchType = &v
}

func (o FeatureAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	return json.Marshal(toSerialize)
}

type NullableFeatureAllOf struct {
	value *FeatureAllOf
	isSet bool
}

func (v NullableFeatureAllOf) Get() *FeatureAllOf {
	return v.value
}

func (v *NullableFeatureAllOf) Set(val *FeatureAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureAllOf(val *FeatureAllOf) *NullableFeatureAllOf {
	return &NullableFeatureAllOf{value: val, isSet: true}
}

func (v NullableFeatureAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


