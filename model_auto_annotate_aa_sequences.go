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

// AutoAnnotateAaSequences struct for AutoAnnotateAaSequences
type AutoAnnotateAaSequences struct {
	// Array of AA sequence IDs.
	AaSequenceIds []string `json:"aaSequenceIds"`
	// Array of feature library IDs.
	FeatureLibraryIds []string `json:"featureLibraryIds"`
}

// NewAutoAnnotateAaSequences instantiates a new AutoAnnotateAaSequences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoAnnotateAaSequences(aaSequenceIds []string, featureLibraryIds []string) *AutoAnnotateAaSequences {
	this := AutoAnnotateAaSequences{}
	this.AaSequenceIds = aaSequenceIds
	this.FeatureLibraryIds = featureLibraryIds
	return &this
}

// NewAutoAnnotateAaSequencesWithDefaults instantiates a new AutoAnnotateAaSequences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoAnnotateAaSequencesWithDefaults() *AutoAnnotateAaSequences {
	this := AutoAnnotateAaSequences{}
	return &this
}

// GetAaSequenceIds returns the AaSequenceIds field value
func (o *AutoAnnotateAaSequences) GetAaSequenceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AaSequenceIds
}

// GetAaSequenceIdsOk returns a tuple with the AaSequenceIds field value
// and a boolean to check if the value has been set.
func (o *AutoAnnotateAaSequences) GetAaSequenceIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AaSequenceIds, true
}

// SetAaSequenceIds sets field value
func (o *AutoAnnotateAaSequences) SetAaSequenceIds(v []string) {
	o.AaSequenceIds = v
}

// GetFeatureLibraryIds returns the FeatureLibraryIds field value
func (o *AutoAnnotateAaSequences) GetFeatureLibraryIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FeatureLibraryIds
}

// GetFeatureLibraryIdsOk returns a tuple with the FeatureLibraryIds field value
// and a boolean to check if the value has been set.
func (o *AutoAnnotateAaSequences) GetFeatureLibraryIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FeatureLibraryIds, true
}

// SetFeatureLibraryIds sets field value
func (o *AutoAnnotateAaSequences) SetFeatureLibraryIds(v []string) {
	o.FeatureLibraryIds = v
}

func (o AutoAnnotateAaSequences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aaSequenceIds"] = o.AaSequenceIds
	}
	if true {
		toSerialize["featureLibraryIds"] = o.FeatureLibraryIds
	}
	return json.Marshal(toSerialize)
}

type NullableAutoAnnotateAaSequences struct {
	value *AutoAnnotateAaSequences
	isSet bool
}

func (v NullableAutoAnnotateAaSequences) Get() *AutoAnnotateAaSequences {
	return v.value
}

func (v *NullableAutoAnnotateAaSequences) Set(val *AutoAnnotateAaSequences) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoAnnotateAaSequences) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoAnnotateAaSequences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoAnnotateAaSequences(val *AutoAnnotateAaSequences) *NullableAutoAnnotateAaSequences {
	return &NullableAutoAnnotateAaSequences{value: val, isSet: true}
}

func (v NullableAutoAnnotateAaSequences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoAnnotateAaSequences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

