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

// AlignedSequence struct for AlignedSequence
type AlignedSequence struct {
	Bases *string `json:"bases,omitempty"`
	DnaSequenceId NullableString `json:"dnaSequenceId,omitempty"`
	Name *string `json:"name,omitempty"`
	// Fraction of bases between trimStart and trimEnd that match the template bases. Only present for Template Alignments; Will be empty for Consensus Alignments. 
	PairwiseIdentity *float32 `json:"pairwiseIdentity,omitempty"`
	TrimEnd *int32 `json:"trimEnd,omitempty"`
	TrimStart *int32 `json:"trimStart,omitempty"`
}

// NewAlignedSequence instantiates a new AlignedSequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlignedSequence() *AlignedSequence {
	this := AlignedSequence{}
	return &this
}

// NewAlignedSequenceWithDefaults instantiates a new AlignedSequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlignedSequenceWithDefaults() *AlignedSequence {
	this := AlignedSequence{}
	return &this
}

// GetBases returns the Bases field value if set, zero value otherwise.
func (o *AlignedSequence) GetBases() string {
	if o == nil || isNil(o.Bases) {
		var ret string
		return ret
	}
	return *o.Bases
}

// GetBasesOk returns a tuple with the Bases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedSequence) GetBasesOk() (*string, bool) {
	if o == nil || isNil(o.Bases) {
    return nil, false
	}
	return o.Bases, true
}

// HasBases returns a boolean if a field has been set.
func (o *AlignedSequence) HasBases() bool {
	if o != nil && !isNil(o.Bases) {
		return true
	}

	return false
}

// SetBases gets a reference to the given string and assigns it to the Bases field.
func (o *AlignedSequence) SetBases(v string) {
	o.Bases = &v
}

// GetDnaSequenceId returns the DnaSequenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlignedSequence) GetDnaSequenceId() string {
	if o == nil || isNil(o.DnaSequenceId.Get()) {
		var ret string
		return ret
	}
	return *o.DnaSequenceId.Get()
}

// GetDnaSequenceIdOk returns a tuple with the DnaSequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlignedSequence) GetDnaSequenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DnaSequenceId.Get(), o.DnaSequenceId.IsSet()
}

// HasDnaSequenceId returns a boolean if a field has been set.
func (o *AlignedSequence) HasDnaSequenceId() bool {
	if o != nil && o.DnaSequenceId.IsSet() {
		return true
	}

	return false
}

// SetDnaSequenceId gets a reference to the given NullableString and assigns it to the DnaSequenceId field.
func (o *AlignedSequence) SetDnaSequenceId(v string) {
	o.DnaSequenceId.Set(&v)
}
// SetDnaSequenceIdNil sets the value for DnaSequenceId to be an explicit nil
func (o *AlignedSequence) SetDnaSequenceIdNil() {
	o.DnaSequenceId.Set(nil)
}

// UnsetDnaSequenceId ensures that no value is present for DnaSequenceId, not even an explicit nil
func (o *AlignedSequence) UnsetDnaSequenceId() {
	o.DnaSequenceId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlignedSequence) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedSequence) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlignedSequence) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlignedSequence) SetName(v string) {
	o.Name = &v
}

// GetPairwiseIdentity returns the PairwiseIdentity field value if set, zero value otherwise.
func (o *AlignedSequence) GetPairwiseIdentity() float32 {
	if o == nil || isNil(o.PairwiseIdentity) {
		var ret float32
		return ret
	}
	return *o.PairwiseIdentity
}

// GetPairwiseIdentityOk returns a tuple with the PairwiseIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedSequence) GetPairwiseIdentityOk() (*float32, bool) {
	if o == nil || isNil(o.PairwiseIdentity) {
    return nil, false
	}
	return o.PairwiseIdentity, true
}

// HasPairwiseIdentity returns a boolean if a field has been set.
func (o *AlignedSequence) HasPairwiseIdentity() bool {
	if o != nil && !isNil(o.PairwiseIdentity) {
		return true
	}

	return false
}

// SetPairwiseIdentity gets a reference to the given float32 and assigns it to the PairwiseIdentity field.
func (o *AlignedSequence) SetPairwiseIdentity(v float32) {
	o.PairwiseIdentity = &v
}

// GetTrimEnd returns the TrimEnd field value if set, zero value otherwise.
func (o *AlignedSequence) GetTrimEnd() int32 {
	if o == nil || isNil(o.TrimEnd) {
		var ret int32
		return ret
	}
	return *o.TrimEnd
}

// GetTrimEndOk returns a tuple with the TrimEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedSequence) GetTrimEndOk() (*int32, bool) {
	if o == nil || isNil(o.TrimEnd) {
    return nil, false
	}
	return o.TrimEnd, true
}

// HasTrimEnd returns a boolean if a field has been set.
func (o *AlignedSequence) HasTrimEnd() bool {
	if o != nil && !isNil(o.TrimEnd) {
		return true
	}

	return false
}

// SetTrimEnd gets a reference to the given int32 and assigns it to the TrimEnd field.
func (o *AlignedSequence) SetTrimEnd(v int32) {
	o.TrimEnd = &v
}

// GetTrimStart returns the TrimStart field value if set, zero value otherwise.
func (o *AlignedSequence) GetTrimStart() int32 {
	if o == nil || isNil(o.TrimStart) {
		var ret int32
		return ret
	}
	return *o.TrimStart
}

// GetTrimStartOk returns a tuple with the TrimStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlignedSequence) GetTrimStartOk() (*int32, bool) {
	if o == nil || isNil(o.TrimStart) {
    return nil, false
	}
	return o.TrimStart, true
}

// HasTrimStart returns a boolean if a field has been set.
func (o *AlignedSequence) HasTrimStart() bool {
	if o != nil && !isNil(o.TrimStart) {
		return true
	}

	return false
}

// SetTrimStart gets a reference to the given int32 and assigns it to the TrimStart field.
func (o *AlignedSequence) SetTrimStart(v int32) {
	o.TrimStart = &v
}

func (o AlignedSequence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Bases) {
		toSerialize["bases"] = o.Bases
	}
	if o.DnaSequenceId.IsSet() {
		toSerialize["dnaSequenceId"] = o.DnaSequenceId.Get()
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PairwiseIdentity) {
		toSerialize["pairwiseIdentity"] = o.PairwiseIdentity
	}
	if !isNil(o.TrimEnd) {
		toSerialize["trimEnd"] = o.TrimEnd
	}
	if !isNil(o.TrimStart) {
		toSerialize["trimStart"] = o.TrimStart
	}
	return json.Marshal(toSerialize)
}

type NullableAlignedSequence struct {
	value *AlignedSequence
	isSet bool
}

func (v NullableAlignedSequence) Get() *AlignedSequence {
	return v.value
}

func (v *NullableAlignedSequence) Set(val *AlignedSequence) {
	v.value = val
	v.isSet = true
}

func (v NullableAlignedSequence) IsSet() bool {
	return v.isSet
}

func (v *NullableAlignedSequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlignedSequence(val *AlignedSequence) *NullableAlignedSequence {
	return &NullableAlignedSequence{value: val, isSet: true}
}

func (v NullableAlignedSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlignedSequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

