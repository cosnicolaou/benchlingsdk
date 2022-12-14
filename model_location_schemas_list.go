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

// LocationSchemasList struct for LocationSchemasList
type LocationSchemasList struct {
	LocationSchemas []LocationSchema `json:"locationSchemas,omitempty"`
}

// NewLocationSchemasList instantiates a new LocationSchemasList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationSchemasList() *LocationSchemasList {
	this := LocationSchemasList{}
	return &this
}

// NewLocationSchemasListWithDefaults instantiates a new LocationSchemasList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationSchemasListWithDefaults() *LocationSchemasList {
	this := LocationSchemasList{}
	return &this
}

// GetLocationSchemas returns the LocationSchemas field value if set, zero value otherwise.
func (o *LocationSchemasList) GetLocationSchemas() []LocationSchema {
	if o == nil || isNil(o.LocationSchemas) {
		var ret []LocationSchema
		return ret
	}
	return o.LocationSchemas
}

// GetLocationSchemasOk returns a tuple with the LocationSchemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationSchemasList) GetLocationSchemasOk() ([]LocationSchema, bool) {
	if o == nil || isNil(o.LocationSchemas) {
    return nil, false
	}
	return o.LocationSchemas, true
}

// HasLocationSchemas returns a boolean if a field has been set.
func (o *LocationSchemasList) HasLocationSchemas() bool {
	if o != nil && !isNil(o.LocationSchemas) {
		return true
	}

	return false
}

// SetLocationSchemas gets a reference to the given []LocationSchema and assigns it to the LocationSchemas field.
func (o *LocationSchemasList) SetLocationSchemas(v []LocationSchema) {
	o.LocationSchemas = v
}

func (o LocationSchemasList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocationSchemas) {
		toSerialize["locationSchemas"] = o.LocationSchemas
	}
	return json.Marshal(toSerialize)
}

type NullableLocationSchemasList struct {
	value *LocationSchemasList
	isSet bool
}

func (v NullableLocationSchemasList) Get() *LocationSchemasList {
	return v.value
}

func (v *NullableLocationSchemasList) Set(val *LocationSchemasList) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationSchemasList) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationSchemasList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationSchemasList(val *LocationSchemasList) *NullableLocationSchemasList {
	return &NullableLocationSchemasList{value: val, isSet: true}
}

func (v NullableLocationSchemasList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationSchemasList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


