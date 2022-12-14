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

// LocationSchemasPaginatedList struct for LocationSchemasPaginatedList
type LocationSchemasPaginatedList struct {
	LocationSchemas []LocationSchema `json:"locationSchemas,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// NewLocationSchemasPaginatedList instantiates a new LocationSchemasPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationSchemasPaginatedList() *LocationSchemasPaginatedList {
	this := LocationSchemasPaginatedList{}
	return &this
}

// NewLocationSchemasPaginatedListWithDefaults instantiates a new LocationSchemasPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationSchemasPaginatedListWithDefaults() *LocationSchemasPaginatedList {
	this := LocationSchemasPaginatedList{}
	return &this
}

// GetLocationSchemas returns the LocationSchemas field value if set, zero value otherwise.
func (o *LocationSchemasPaginatedList) GetLocationSchemas() []LocationSchema {
	if o == nil || isNil(o.LocationSchemas) {
		var ret []LocationSchema
		return ret
	}
	return o.LocationSchemas
}

// GetLocationSchemasOk returns a tuple with the LocationSchemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationSchemasPaginatedList) GetLocationSchemasOk() ([]LocationSchema, bool) {
	if o == nil || isNil(o.LocationSchemas) {
    return nil, false
	}
	return o.LocationSchemas, true
}

// HasLocationSchemas returns a boolean if a field has been set.
func (o *LocationSchemasPaginatedList) HasLocationSchemas() bool {
	if o != nil && !isNil(o.LocationSchemas) {
		return true
	}

	return false
}

// SetLocationSchemas gets a reference to the given []LocationSchema and assigns it to the LocationSchemas field.
func (o *LocationSchemasPaginatedList) SetLocationSchemas(v []LocationSchema) {
	o.LocationSchemas = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *LocationSchemasPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationSchemasPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *LocationSchemasPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *LocationSchemasPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

func (o LocationSchemasPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocationSchemas) {
		toSerialize["locationSchemas"] = o.LocationSchemas
	}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableLocationSchemasPaginatedList struct {
	value *LocationSchemasPaginatedList
	isSet bool
}

func (v NullableLocationSchemasPaginatedList) Get() *LocationSchemasPaginatedList {
	return v.value
}

func (v *NullableLocationSchemasPaginatedList) Set(val *LocationSchemasPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationSchemasPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationSchemasPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationSchemasPaginatedList(val *LocationSchemasPaginatedList) *NullableLocationSchemasPaginatedList {
	return &NullableLocationSchemasPaginatedList{value: val, isSet: true}
}

func (v NullableLocationSchemasPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationSchemasPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


