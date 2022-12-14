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

// MoleculesUnarchive The request body for unarchiving Molecules. 
type MoleculesUnarchive struct {
	MoleculeIds []string `json:"moleculeIds"`
}

// NewMoleculesUnarchive instantiates a new MoleculesUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoleculesUnarchive(moleculeIds []string) *MoleculesUnarchive {
	this := MoleculesUnarchive{}
	this.MoleculeIds = moleculeIds
	return &this
}

// NewMoleculesUnarchiveWithDefaults instantiates a new MoleculesUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoleculesUnarchiveWithDefaults() *MoleculesUnarchive {
	this := MoleculesUnarchive{}
	return &this
}

// GetMoleculeIds returns the MoleculeIds field value
func (o *MoleculesUnarchive) GetMoleculeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MoleculeIds
}

// GetMoleculeIdsOk returns a tuple with the MoleculeIds field value
// and a boolean to check if the value has been set.
func (o *MoleculesUnarchive) GetMoleculeIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MoleculeIds, true
}

// SetMoleculeIds sets field value
func (o *MoleculesUnarchive) SetMoleculeIds(v []string) {
	o.MoleculeIds = v
}

func (o MoleculesUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["moleculeIds"] = o.MoleculeIds
	}
	return json.Marshal(toSerialize)
}

type NullableMoleculesUnarchive struct {
	value *MoleculesUnarchive
	isSet bool
}

func (v NullableMoleculesUnarchive) Get() *MoleculesUnarchive {
	return v.value
}

func (v *NullableMoleculesUnarchive) Set(val *MoleculesUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableMoleculesUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableMoleculesUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoleculesUnarchive(val *MoleculesUnarchive) *NullableMoleculesUnarchive {
	return &NullableMoleculesUnarchive{value: val, isSet: true}
}

func (v NullableMoleculesUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoleculesUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


