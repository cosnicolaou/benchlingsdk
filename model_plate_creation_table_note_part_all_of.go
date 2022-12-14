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

// PlateCreationTableNotePartAllOf struct for PlateCreationTableNotePartAllOf
type PlateCreationTableNotePartAllOf struct {
	PlateSchemaId *string `json:"plateSchemaId,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPlateCreationTableNotePartAllOf instantiates a new PlateCreationTableNotePartAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlateCreationTableNotePartAllOf() *PlateCreationTableNotePartAllOf {
	this := PlateCreationTableNotePartAllOf{}
	return &this
}

// NewPlateCreationTableNotePartAllOfWithDefaults instantiates a new PlateCreationTableNotePartAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlateCreationTableNotePartAllOfWithDefaults() *PlateCreationTableNotePartAllOf {
	this := PlateCreationTableNotePartAllOf{}
	return &this
}

// GetPlateSchemaId returns the PlateSchemaId field value if set, zero value otherwise.
func (o *PlateCreationTableNotePartAllOf) GetPlateSchemaId() string {
	if o == nil || isNil(o.PlateSchemaId) {
		var ret string
		return ret
	}
	return *o.PlateSchemaId
}

// GetPlateSchemaIdOk returns a tuple with the PlateSchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateCreationTableNotePartAllOf) GetPlateSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.PlateSchemaId) {
    return nil, false
	}
	return o.PlateSchemaId, true
}

// HasPlateSchemaId returns a boolean if a field has been set.
func (o *PlateCreationTableNotePartAllOf) HasPlateSchemaId() bool {
	if o != nil && !isNil(o.PlateSchemaId) {
		return true
	}

	return false
}

// SetPlateSchemaId gets a reference to the given string and assigns it to the PlateSchemaId field.
func (o *PlateCreationTableNotePartAllOf) SetPlateSchemaId(v string) {
	o.PlateSchemaId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlateCreationTableNotePartAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateCreationTableNotePartAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlateCreationTableNotePartAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlateCreationTableNotePartAllOf) SetType(v string) {
	o.Type = &v
}

func (o PlateCreationTableNotePartAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PlateSchemaId) {
		toSerialize["plateSchemaId"] = o.PlateSchemaId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePlateCreationTableNotePartAllOf struct {
	value *PlateCreationTableNotePartAllOf
	isSet bool
}

func (v NullablePlateCreationTableNotePartAllOf) Get() *PlateCreationTableNotePartAllOf {
	return v.value
}

func (v *NullablePlateCreationTableNotePartAllOf) Set(val *PlateCreationTableNotePartAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlateCreationTableNotePartAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlateCreationTableNotePartAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlateCreationTableNotePartAllOf(val *PlateCreationTableNotePartAllOf) *NullablePlateCreationTableNotePartAllOf {
	return &NullablePlateCreationTableNotePartAllOf{value: val, isSet: true}
}

func (v NullablePlateCreationTableNotePartAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlateCreationTableNotePartAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


