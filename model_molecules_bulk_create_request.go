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

// MoleculesBulkCreateRequest struct for MoleculesBulkCreateRequest
type MoleculesBulkCreateRequest struct {
	Molecules []MoleculeCreate `json:"molecules,omitempty"`
}

// NewMoleculesBulkCreateRequest instantiates a new MoleculesBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoleculesBulkCreateRequest() *MoleculesBulkCreateRequest {
	this := MoleculesBulkCreateRequest{}
	return &this
}

// NewMoleculesBulkCreateRequestWithDefaults instantiates a new MoleculesBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoleculesBulkCreateRequestWithDefaults() *MoleculesBulkCreateRequest {
	this := MoleculesBulkCreateRequest{}
	return &this
}

// GetMolecules returns the Molecules field value if set, zero value otherwise.
func (o *MoleculesBulkCreateRequest) GetMolecules() []MoleculeCreate {
	if o == nil || isNil(o.Molecules) {
		var ret []MoleculeCreate
		return ret
	}
	return o.Molecules
}

// GetMoleculesOk returns a tuple with the Molecules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculesBulkCreateRequest) GetMoleculesOk() ([]MoleculeCreate, bool) {
	if o == nil || isNil(o.Molecules) {
    return nil, false
	}
	return o.Molecules, true
}

// HasMolecules returns a boolean if a field has been set.
func (o *MoleculesBulkCreateRequest) HasMolecules() bool {
	if o != nil && !isNil(o.Molecules) {
		return true
	}

	return false
}

// SetMolecules gets a reference to the given []MoleculeCreate and assigns it to the Molecules field.
func (o *MoleculesBulkCreateRequest) SetMolecules(v []MoleculeCreate) {
	o.Molecules = v
}

func (o MoleculesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Molecules) {
		toSerialize["molecules"] = o.Molecules
	}
	return json.Marshal(toSerialize)
}

type NullableMoleculesBulkCreateRequest struct {
	value *MoleculesBulkCreateRequest
	isSet bool
}

func (v NullableMoleculesBulkCreateRequest) Get() *MoleculesBulkCreateRequest {
	return v.value
}

func (v *NullableMoleculesBulkCreateRequest) Set(val *MoleculesBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoleculesBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoleculesBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoleculesBulkCreateRequest(val *MoleculesBulkCreateRequest) *NullableMoleculesBulkCreateRequest {
	return &NullableMoleculesBulkCreateRequest{value: val, isSet: true}
}

func (v NullableMoleculesBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoleculesBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


