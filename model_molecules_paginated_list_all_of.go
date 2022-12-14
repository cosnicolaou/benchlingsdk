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

// MoleculesPaginatedListAllOf struct for MoleculesPaginatedListAllOf
type MoleculesPaginatedListAllOf struct {
	Molecules []Molecule `json:"molecules,omitempty"`
}

// NewMoleculesPaginatedListAllOf instantiates a new MoleculesPaginatedListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoleculesPaginatedListAllOf() *MoleculesPaginatedListAllOf {
	this := MoleculesPaginatedListAllOf{}
	return &this
}

// NewMoleculesPaginatedListAllOfWithDefaults instantiates a new MoleculesPaginatedListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoleculesPaginatedListAllOfWithDefaults() *MoleculesPaginatedListAllOf {
	this := MoleculesPaginatedListAllOf{}
	return &this
}

// GetMolecules returns the Molecules field value if set, zero value otherwise.
func (o *MoleculesPaginatedListAllOf) GetMolecules() []Molecule {
	if o == nil || isNil(o.Molecules) {
		var ret []Molecule
		return ret
	}
	return o.Molecules
}

// GetMoleculesOk returns a tuple with the Molecules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculesPaginatedListAllOf) GetMoleculesOk() ([]Molecule, bool) {
	if o == nil || isNil(o.Molecules) {
    return nil, false
	}
	return o.Molecules, true
}

// HasMolecules returns a boolean if a field has been set.
func (o *MoleculesPaginatedListAllOf) HasMolecules() bool {
	if o != nil && !isNil(o.Molecules) {
		return true
	}

	return false
}

// SetMolecules gets a reference to the given []Molecule and assigns it to the Molecules field.
func (o *MoleculesPaginatedListAllOf) SetMolecules(v []Molecule) {
	o.Molecules = v
}

func (o MoleculesPaginatedListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Molecules) {
		toSerialize["molecules"] = o.Molecules
	}
	return json.Marshal(toSerialize)
}

type NullableMoleculesPaginatedListAllOf struct {
	value *MoleculesPaginatedListAllOf
	isSet bool
}

func (v NullableMoleculesPaginatedListAllOf) Get() *MoleculesPaginatedListAllOf {
	return v.value
}

func (v *NullableMoleculesPaginatedListAllOf) Set(val *MoleculesPaginatedListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoleculesPaginatedListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoleculesPaginatedListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoleculesPaginatedListAllOf(val *MoleculesPaginatedListAllOf) *NullableMoleculesPaginatedListAllOf {
	return &NullableMoleculesPaginatedListAllOf{value: val, isSet: true}
}

func (v NullableMoleculesPaginatedListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoleculesPaginatedListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


