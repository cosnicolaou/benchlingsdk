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

// MoleculeBaseRequestChemicalStructure Chemical structure of the Molecule. 
type MoleculeBaseRequestChemicalStructure struct {
	StructureFormat *string `json:"structureFormat,omitempty"`
	// Chemical structure in SMILES or molfile format.
	Value *string `json:"value,omitempty"`
}

// NewMoleculeBaseRequestChemicalStructure instantiates a new MoleculeBaseRequestChemicalStructure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoleculeBaseRequestChemicalStructure() *MoleculeBaseRequestChemicalStructure {
	this := MoleculeBaseRequestChemicalStructure{}
	return &this
}

// NewMoleculeBaseRequestChemicalStructureWithDefaults instantiates a new MoleculeBaseRequestChemicalStructure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoleculeBaseRequestChemicalStructureWithDefaults() *MoleculeBaseRequestChemicalStructure {
	this := MoleculeBaseRequestChemicalStructure{}
	return &this
}

// GetStructureFormat returns the StructureFormat field value if set, zero value otherwise.
func (o *MoleculeBaseRequestChemicalStructure) GetStructureFormat() string {
	if o == nil || isNil(o.StructureFormat) {
		var ret string
		return ret
	}
	return *o.StructureFormat
}

// GetStructureFormatOk returns a tuple with the StructureFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeBaseRequestChemicalStructure) GetStructureFormatOk() (*string, bool) {
	if o == nil || isNil(o.StructureFormat) {
    return nil, false
	}
	return o.StructureFormat, true
}

// HasStructureFormat returns a boolean if a field has been set.
func (o *MoleculeBaseRequestChemicalStructure) HasStructureFormat() bool {
	if o != nil && !isNil(o.StructureFormat) {
		return true
	}

	return false
}

// SetStructureFormat gets a reference to the given string and assigns it to the StructureFormat field.
func (o *MoleculeBaseRequestChemicalStructure) SetStructureFormat(v string) {
	o.StructureFormat = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MoleculeBaseRequestChemicalStructure) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeBaseRequestChemicalStructure) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MoleculeBaseRequestChemicalStructure) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MoleculeBaseRequestChemicalStructure) SetValue(v string) {
	o.Value = &v
}

func (o MoleculeBaseRequestChemicalStructure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StructureFormat) {
		toSerialize["structureFormat"] = o.StructureFormat
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMoleculeBaseRequestChemicalStructure struct {
	value *MoleculeBaseRequestChemicalStructure
	isSet bool
}

func (v NullableMoleculeBaseRequestChemicalStructure) Get() *MoleculeBaseRequestChemicalStructure {
	return v.value
}

func (v *NullableMoleculeBaseRequestChemicalStructure) Set(val *MoleculeBaseRequestChemicalStructure) {
	v.value = val
	v.isSet = true
}

func (v NullableMoleculeBaseRequestChemicalStructure) IsSet() bool {
	return v.isSet
}

func (v *NullableMoleculeBaseRequestChemicalStructure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoleculeBaseRequestChemicalStructure(val *MoleculeBaseRequestChemicalStructure) *NullableMoleculeBaseRequestChemicalStructure {
	return &NullableMoleculeBaseRequestChemicalStructure{value: val, isSet: true}
}

func (v NullableMoleculeBaseRequestChemicalStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoleculeBaseRequestChemicalStructure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


