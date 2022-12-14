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

// DnaAlignmentBase struct for DnaAlignmentBase
type DnaAlignmentBase struct {
	Algorithm string `json:"algorithm"`
	Files []DnaAlignmentBaseFilesInner `json:"files"`
	Name *string `json:"name,omitempty"`
}

// NewDnaAlignmentBase instantiates a new DnaAlignmentBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaAlignmentBase(algorithm string, files []DnaAlignmentBaseFilesInner) *DnaAlignmentBase {
	this := DnaAlignmentBase{}
	this.Algorithm = algorithm
	this.Files = files
	return &this
}

// NewDnaAlignmentBaseWithDefaults instantiates a new DnaAlignmentBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaAlignmentBaseWithDefaults() *DnaAlignmentBase {
	this := DnaAlignmentBase{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *DnaAlignmentBase) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *DnaAlignmentBase) GetAlgorithmOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *DnaAlignmentBase) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetFiles returns the Files field value
func (o *DnaAlignmentBase) GetFiles() []DnaAlignmentBaseFilesInner {
	if o == nil {
		var ret []DnaAlignmentBaseFilesInner
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *DnaAlignmentBase) GetFilesOk() ([]DnaAlignmentBaseFilesInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *DnaAlignmentBase) SetFiles(v []DnaAlignmentBaseFilesInner) {
	o.Files = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnaAlignmentBase) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignmentBase) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnaAlignmentBase) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnaAlignmentBase) SetName(v string) {
	o.Name = &v
}

func (o DnaAlignmentBase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableDnaAlignmentBase struct {
	value *DnaAlignmentBase
	isSet bool
}

func (v NullableDnaAlignmentBase) Get() *DnaAlignmentBase {
	return v.value
}

func (v *NullableDnaAlignmentBase) Set(val *DnaAlignmentBase) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaAlignmentBase) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaAlignmentBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaAlignmentBase(val *DnaAlignmentBase) *NullableDnaAlignmentBase {
	return &NullableDnaAlignmentBase{value: val, isSet: true}
}

func (v NullableDnaAlignmentBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaAlignmentBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


