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

// DnaTemplateAlignmentCreate struct for DnaTemplateAlignmentCreate
type DnaTemplateAlignmentCreate struct {
	Algorithm string `json:"algorithm"`
	Files []DnaAlignmentBaseFilesInner `json:"files"`
	Name *string `json:"name,omitempty"`
	TemplateSequenceId string `json:"templateSequenceId"`
}

// NewDnaTemplateAlignmentCreate instantiates a new DnaTemplateAlignmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaTemplateAlignmentCreate(algorithm string, files []DnaAlignmentBaseFilesInner, templateSequenceId string) *DnaTemplateAlignmentCreate {
	this := DnaTemplateAlignmentCreate{}
	this.Algorithm = algorithm
	this.Files = files
	this.TemplateSequenceId = templateSequenceId
	return &this
}

// NewDnaTemplateAlignmentCreateWithDefaults instantiates a new DnaTemplateAlignmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaTemplateAlignmentCreateWithDefaults() *DnaTemplateAlignmentCreate {
	this := DnaTemplateAlignmentCreate{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *DnaTemplateAlignmentCreate) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *DnaTemplateAlignmentCreate) GetAlgorithmOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *DnaTemplateAlignmentCreate) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetFiles returns the Files field value
func (o *DnaTemplateAlignmentCreate) GetFiles() []DnaAlignmentBaseFilesInner {
	if o == nil {
		var ret []DnaAlignmentBaseFilesInner
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *DnaTemplateAlignmentCreate) GetFilesOk() ([]DnaAlignmentBaseFilesInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *DnaTemplateAlignmentCreate) SetFiles(v []DnaAlignmentBaseFilesInner) {
	o.Files = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnaTemplateAlignmentCreate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaTemplateAlignmentCreate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnaTemplateAlignmentCreate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnaTemplateAlignmentCreate) SetName(v string) {
	o.Name = &v
}

// GetTemplateSequenceId returns the TemplateSequenceId field value
func (o *DnaTemplateAlignmentCreate) GetTemplateSequenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateSequenceId
}

// GetTemplateSequenceIdOk returns a tuple with the TemplateSequenceId field value
// and a boolean to check if the value has been set.
func (o *DnaTemplateAlignmentCreate) GetTemplateSequenceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TemplateSequenceId, true
}

// SetTemplateSequenceId sets field value
func (o *DnaTemplateAlignmentCreate) SetTemplateSequenceId(v string) {
	o.TemplateSequenceId = v
}

func (o DnaTemplateAlignmentCreate) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["templateSequenceId"] = o.TemplateSequenceId
	}
	return json.Marshal(toSerialize)
}

type NullableDnaTemplateAlignmentCreate struct {
	value *DnaTemplateAlignmentCreate
	isSet bool
}

func (v NullableDnaTemplateAlignmentCreate) Get() *DnaTemplateAlignmentCreate {
	return v.value
}

func (v *NullableDnaTemplateAlignmentCreate) Set(val *DnaTemplateAlignmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaTemplateAlignmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaTemplateAlignmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaTemplateAlignmentCreate(val *DnaTemplateAlignmentCreate) *NullableDnaTemplateAlignmentCreate {
	return &NullableDnaTemplateAlignmentCreate{value: val, isSet: true}
}

func (v NullableDnaTemplateAlignmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaTemplateAlignmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

