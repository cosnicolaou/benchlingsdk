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

// MoleculesPaginatedList struct for MoleculesPaginatedList
type MoleculesPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	Molecules []Molecule `json:"molecules,omitempty"`
}

// NewMoleculesPaginatedList instantiates a new MoleculesPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoleculesPaginatedList() *MoleculesPaginatedList {
	this := MoleculesPaginatedList{}
	return &this
}

// NewMoleculesPaginatedListWithDefaults instantiates a new MoleculesPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoleculesPaginatedListWithDefaults() *MoleculesPaginatedList {
	this := MoleculesPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *MoleculesPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculesPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *MoleculesPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *MoleculesPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetMolecules returns the Molecules field value if set, zero value otherwise.
func (o *MoleculesPaginatedList) GetMolecules() []Molecule {
	if o == nil || isNil(o.Molecules) {
		var ret []Molecule
		return ret
	}
	return o.Molecules
}

// GetMoleculesOk returns a tuple with the Molecules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculesPaginatedList) GetMoleculesOk() ([]Molecule, bool) {
	if o == nil || isNil(o.Molecules) {
    return nil, false
	}
	return o.Molecules, true
}

// HasMolecules returns a boolean if a field has been set.
func (o *MoleculesPaginatedList) HasMolecules() bool {
	if o != nil && !isNil(o.Molecules) {
		return true
	}

	return false
}

// SetMolecules gets a reference to the given []Molecule and assigns it to the Molecules field.
func (o *MoleculesPaginatedList) SetMolecules(v []Molecule) {
	o.Molecules = v
}

func (o MoleculesPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.Molecules) {
		toSerialize["molecules"] = o.Molecules
	}
	return json.Marshal(toSerialize)
}

type NullableMoleculesPaginatedList struct {
	value *MoleculesPaginatedList
	isSet bool
}

func (v NullableMoleculesPaginatedList) Get() *MoleculesPaginatedList {
	return v.value
}

func (v *NullableMoleculesPaginatedList) Set(val *MoleculesPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableMoleculesPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableMoleculesPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoleculesPaginatedList(val *MoleculesPaginatedList) *NullableMoleculesPaginatedList {
	return &NullableMoleculesPaginatedList{value: val, isSet: true}
}

func (v NullableMoleculesPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoleculesPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


