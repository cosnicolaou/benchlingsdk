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

// PlateSchemasList struct for PlateSchemasList
type PlateSchemasList struct {
	PlateSchemas []PlateSchema `json:"plateSchemas,omitempty"`
}

// NewPlateSchemasList instantiates a new PlateSchemasList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlateSchemasList() *PlateSchemasList {
	this := PlateSchemasList{}
	return &this
}

// NewPlateSchemasListWithDefaults instantiates a new PlateSchemasList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlateSchemasListWithDefaults() *PlateSchemasList {
	this := PlateSchemasList{}
	return &this
}

// GetPlateSchemas returns the PlateSchemas field value if set, zero value otherwise.
func (o *PlateSchemasList) GetPlateSchemas() []PlateSchema {
	if o == nil || isNil(o.PlateSchemas) {
		var ret []PlateSchema
		return ret
	}
	return o.PlateSchemas
}

// GetPlateSchemasOk returns a tuple with the PlateSchemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchemasList) GetPlateSchemasOk() ([]PlateSchema, bool) {
	if o == nil || isNil(o.PlateSchemas) {
    return nil, false
	}
	return o.PlateSchemas, true
}

// HasPlateSchemas returns a boolean if a field has been set.
func (o *PlateSchemasList) HasPlateSchemas() bool {
	if o != nil && !isNil(o.PlateSchemas) {
		return true
	}

	return false
}

// SetPlateSchemas gets a reference to the given []PlateSchema and assigns it to the PlateSchemas field.
func (o *PlateSchemasList) SetPlateSchemas(v []PlateSchema) {
	o.PlateSchemas = v
}

func (o PlateSchemasList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PlateSchemas) {
		toSerialize["plateSchemas"] = o.PlateSchemas
	}
	return json.Marshal(toSerialize)
}

type NullablePlateSchemasList struct {
	value *PlateSchemasList
	isSet bool
}

func (v NullablePlateSchemasList) Get() *PlateSchemasList {
	return v.value
}

func (v *NullablePlateSchemasList) Set(val *PlateSchemasList) {
	v.value = val
	v.isSet = true
}

func (v NullablePlateSchemasList) IsSet() bool {
	return v.isSet
}

func (v *NullablePlateSchemasList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlateSchemasList(val *PlateSchemasList) *NullablePlateSchemasList {
	return &NullablePlateSchemasList{value: val, isSet: true}
}

func (v NullablePlateSchemasList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlateSchemasList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


