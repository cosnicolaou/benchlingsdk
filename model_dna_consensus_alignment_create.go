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

// DnaConsensusAlignmentCreate struct for DnaConsensusAlignmentCreate
type DnaConsensusAlignmentCreate struct {
	Algorithm string `json:"algorithm"`
	Files []DnaAlignmentBaseFilesInner `json:"files"`
	Name *string `json:"name,omitempty"`
	NewSequence *DnaConsensusAlignmentCreateAllOfNewSequence `json:"newSequence,omitempty"`
	SequenceId *string `json:"sequenceId,omitempty"`
}

// NewDnaConsensusAlignmentCreate instantiates a new DnaConsensusAlignmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaConsensusAlignmentCreate(algorithm string, files []DnaAlignmentBaseFilesInner) *DnaConsensusAlignmentCreate {
	this := DnaConsensusAlignmentCreate{}
	this.Algorithm = algorithm
	this.Files = files
	return &this
}

// NewDnaConsensusAlignmentCreateWithDefaults instantiates a new DnaConsensusAlignmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaConsensusAlignmentCreateWithDefaults() *DnaConsensusAlignmentCreate {
	this := DnaConsensusAlignmentCreate{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *DnaConsensusAlignmentCreate) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreate) GetAlgorithmOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *DnaConsensusAlignmentCreate) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetFiles returns the Files field value
func (o *DnaConsensusAlignmentCreate) GetFiles() []DnaAlignmentBaseFilesInner {
	if o == nil {
		var ret []DnaAlignmentBaseFilesInner
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreate) GetFilesOk() ([]DnaAlignmentBaseFilesInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *DnaConsensusAlignmentCreate) SetFiles(v []DnaAlignmentBaseFilesInner) {
	o.Files = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnaConsensusAlignmentCreate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnaConsensusAlignmentCreate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnaConsensusAlignmentCreate) SetName(v string) {
	o.Name = &v
}

// GetNewSequence returns the NewSequence field value if set, zero value otherwise.
func (o *DnaConsensusAlignmentCreate) GetNewSequence() DnaConsensusAlignmentCreateAllOfNewSequence {
	if o == nil || isNil(o.NewSequence) {
		var ret DnaConsensusAlignmentCreateAllOfNewSequence
		return ret
	}
	return *o.NewSequence
}

// GetNewSequenceOk returns a tuple with the NewSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreate) GetNewSequenceOk() (*DnaConsensusAlignmentCreateAllOfNewSequence, bool) {
	if o == nil || isNil(o.NewSequence) {
    return nil, false
	}
	return o.NewSequence, true
}

// HasNewSequence returns a boolean if a field has been set.
func (o *DnaConsensusAlignmentCreate) HasNewSequence() bool {
	if o != nil && !isNil(o.NewSequence) {
		return true
	}

	return false
}

// SetNewSequence gets a reference to the given DnaConsensusAlignmentCreateAllOfNewSequence and assigns it to the NewSequence field.
func (o *DnaConsensusAlignmentCreate) SetNewSequence(v DnaConsensusAlignmentCreateAllOfNewSequence) {
	o.NewSequence = &v
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise.
func (o *DnaConsensusAlignmentCreate) GetSequenceId() string {
	if o == nil || isNil(o.SequenceId) {
		var ret string
		return ret
	}
	return *o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaConsensusAlignmentCreate) GetSequenceIdOk() (*string, bool) {
	if o == nil || isNil(o.SequenceId) {
    return nil, false
	}
	return o.SequenceId, true
}

// HasSequenceId returns a boolean if a field has been set.
func (o *DnaConsensusAlignmentCreate) HasSequenceId() bool {
	if o != nil && !isNil(o.SequenceId) {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given string and assigns it to the SequenceId field.
func (o *DnaConsensusAlignmentCreate) SetSequenceId(v string) {
	o.SequenceId = &v
}

func (o DnaConsensusAlignmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["algorithm"] = o.Algorithm
	}
	if true {
		toSerialize["files"] = o.Files
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NewSequence) {
		toSerialize["newSequence"] = o.NewSequence
	}
	if !isNil(o.SequenceId) {
		toSerialize["sequenceId"] = o.SequenceId
	}
	return json.Marshal(toSerialize)
}

type NullableDnaConsensusAlignmentCreate struct {
	value *DnaConsensusAlignmentCreate
	isSet bool
}

func (v NullableDnaConsensusAlignmentCreate) Get() *DnaConsensusAlignmentCreate {
	return v.value
}

func (v *NullableDnaConsensusAlignmentCreate) Set(val *DnaConsensusAlignmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaConsensusAlignmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaConsensusAlignmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaConsensusAlignmentCreate(val *DnaConsensusAlignmentCreate) *NullableDnaConsensusAlignmentCreate {
	return &NullableDnaConsensusAlignmentCreate{value: val, isSet: true}
}

func (v NullableDnaConsensusAlignmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaConsensusAlignmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


