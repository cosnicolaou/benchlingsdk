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

// AlignedNucleotideSequence struct for AlignedNucleotideSequence
type AlignedNucleotideSequence struct {
	Bases *string `json:"bases,omitempty"`
	Name *string `json:"name,omitempty"`
	// Fraction of bases between trimStart and trimEnd that match the template bases. Only present for Template Alignments; Will be empty for Consensus Alignments. 
	PairwiseIdentity *float32 `json:"pairwiseIdentity,omitempty"`
	SequenceId NullableString `json:"sequenceId,omitempty"`
	TrimEnd *int32 `json:"trimEnd,omitempty"`
	TrimStart *int32 `json:"trimStart,omitempty"`
}

// NewAlignedNucleotideSequence instantiates a new AlignedNucleotideSequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlignedNucleotideSequence() *AlignedNucleotideSequence {
	this := AlignedNucleotideSequence{}
	return &this
}

// NewAlignedNucleotideSequenceWithDefaults instantiates a new AlignedNucleotideSequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlignedNucleotideSequenceWithDefaults() *AlignedNucleotideSequence {
	this := AlignedNucleotideSequence{}
	return &this
}

// GetBases returns the Bases field value if set, zero value otherwise.
func (o *AlignedNucleotideSequence) GetBases() string {
	if o == nil || isNil(o.Bases) {
		var ret string
		return ret
	}
	return *o.Bases
}

// GetBasesOk returns a tuple with the Bases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedNucleotideSequence) GetBasesOk() (*string, bool) {
	if o == nil || isNil(o.Bases) {
    return nil, false
	}
	return o.Bases, true
}

// HasBases returns a boolean if a field has been set.
func (o *AlignedNucleotideSequence) HasBases() bool {
	if o != nil && !isNil(o.Bases) {
		return true
	}

	return false
}

// SetBases gets a reference to the given string and assigns it to the Bases field.
func (o *AlignedNucleotideSequence) SetBases(v string) {
	o.Bases = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlignedNucleotideSequence) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedNucleotideSequence) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlignedNucleotideSequence) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlignedNucleotideSequence) SetName(v string) {
	o.Name = &v
}

// GetPairwiseIdentity returns the PairwiseIdentity field value if set, zero value otherwise.
func (o *AlignedNucleotideSequence) GetPairwiseIdentity() float32 {
	if o == nil || isNil(o.PairwiseIdentity) {
		var ret float32
		return ret
	}
	return *o.PairwiseIdentity
}

// GetPairwiseIdentityOk returns a tuple with the PairwiseIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedNucleotideSequence) GetPairwiseIdentityOk() (*float32, bool) {
	if o == nil || isNil(o.PairwiseIdentity) {
    return nil, false
	}
	return o.PairwiseIdentity, true
}

// HasPairwiseIdentity returns a boolean if a field has been set.
func (o *AlignedNucleotideSequence) HasPairwiseIdentity() bool {
	if o != nil && !isNil(o.PairwiseIdentity) {
		return true
	}

	return false
}

// SetPairwiseIdentity gets a reference to the given float32 and assigns it to the PairwiseIdentity field.
func (o *AlignedNucleotideSequence) SetPairwiseIdentity(v float32) {
	o.PairwiseIdentity = &v
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlignedNucleotideSequence) GetSequenceId() string {
	if o == nil || isNil(o.SequenceId.Get()) {
		var ret string
		return ret
	}
	return *o.SequenceId.Get()
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlignedNucleotideSequence) GetSequenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SequenceId.Get(), o.SequenceId.IsSet()
}

// HasSequenceId returns a boolean if a field has been set.
func (o *AlignedNucleotideSequence) HasSequenceId() bool {
	if o != nil && o.SequenceId.IsSet() {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given NullableString and assigns it to the SequenceId field.
func (o *AlignedNucleotideSequence) SetSequenceId(v string) {
	o.SequenceId.Set(&v)
}
// SetSequenceIdNil sets the value for SequenceId to be an explicit nil
func (o *AlignedNucleotideSequence) SetSequenceIdNil() {
	o.SequenceId.Set(nil)
}

// UnsetSequenceId ensures that no value is present for SequenceId, not even an explicit nil
func (o *AlignedNucleotideSequence) UnsetSequenceId() {
	o.SequenceId.Unset()
}

// GetTrimEnd returns the TrimEnd field value if set, zero value otherwise.
func (o *AlignedNucleotideSequence) GetTrimEnd() int32 {
	if o == nil || isNil(o.TrimEnd) {
		var ret int32
		return ret
	}
	return *o.TrimEnd
}

// GetTrimEndOk returns a tuple with the TrimEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedNucleotideSequence) GetTrimEndOk() (*int32, bool) {
	if o == nil || isNil(o.TrimEnd) {
    return nil, false
	}
	return o.TrimEnd, true
}

// HasTrimEnd returns a boolean if a field has been set.
func (o *AlignedNucleotideSequence) HasTrimEnd() bool {
	if o != nil && !isNil(o.TrimEnd) {
		return true
	}

	return false
}

// SetTrimEnd gets a reference to the given int32 and assigns it to the TrimEnd field.
func (o *AlignedNucleotideSequence) SetTrimEnd(v int32) {
	o.TrimEnd = &v
}

// GetTrimStart returns the TrimStart field value if set, zero value otherwise.
func (o *AlignedNucleotideSequence) GetTrimStart() int32 {
	if o == nil || isNil(o.TrimStart) {
		var ret int32
		return ret
	}
	return *o.TrimStart
}

// GetTrimStartOk returns a tuple with the TrimStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedNucleotideSequence) GetTrimStartOk() (*int32, bool) {
	if o == nil || isNil(o.TrimStart) {
    return nil, false
	}
	return o.TrimStart, true
}

// HasTrimStart returns a boolean if a field has been set.
func (o *AlignedNucleotideSequence) HasTrimStart() bool {
	if o != nil && !isNil(o.TrimStart) {
		return true
	}

	return false
}

// SetTrimStart gets a reference to the given int32 and assigns it to the TrimStart field.
func (o *AlignedNucleotideSequence) SetTrimStart(v int32) {
	o.TrimStart = &v
}

func (o AlignedNucleotideSequence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Bases) {
		toSerialize["bases"] = o.Bases
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PairwiseIdentity) {
		toSerialize["pairwiseIdentity"] = o.PairwiseIdentity
	}
	if o.SequenceId.IsSet() {
		toSerialize["sequenceId"] = o.SequenceId.Get()
	}
	if !isNil(o.TrimEnd) {
		toSerialize["trimEnd"] = o.TrimEnd
	}
	if !isNil(o.TrimStart) {
		toSerialize["trimStart"] = o.TrimStart
	}
	return json.Marshal(toSerialize)
}

type NullableAlignedNucleotideSequence struct {
	value *AlignedNucleotideSequence
	isSet bool
}

func (v NullableAlignedNucleotideSequence) Get() *AlignedNucleotideSequence {
	return v.value
}

func (v *NullableAlignedNucleotideSequence) Set(val *AlignedNucleotideSequence) {
	v.value = val
	v.isSet = true
}

func (v NullableAlignedNucleotideSequence) IsSet() bool {
	return v.isSet
}

func (v *NullableAlignedNucleotideSequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlignedNucleotideSequence(val *AlignedNucleotideSequence) *NullableAlignedNucleotideSequence {
	return &NullableAlignedNucleotideSequence{value: val, isSet: true}
}

func (v NullableAlignedNucleotideSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlignedNucleotideSequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


