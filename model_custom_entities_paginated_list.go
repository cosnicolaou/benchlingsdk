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

// CustomEntitiesPaginatedList struct for CustomEntitiesPaginatedList
type CustomEntitiesPaginatedList struct {
	CustomEntities []CustomEntity `json:"customEntities,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// NewCustomEntitiesPaginatedList instantiates a new CustomEntitiesPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEntitiesPaginatedList() *CustomEntitiesPaginatedList {
	this := CustomEntitiesPaginatedList{}
	return &this
}

// NewCustomEntitiesPaginatedListWithDefaults instantiates a new CustomEntitiesPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEntitiesPaginatedListWithDefaults() *CustomEntitiesPaginatedList {
	this := CustomEntitiesPaginatedList{}
	return &this
}

// GetCustomEntities returns the CustomEntities field value if set, zero value otherwise.
func (o *CustomEntitiesPaginatedList) GetCustomEntities() []CustomEntity {
	if o == nil || isNil(o.CustomEntities) {
		var ret []CustomEntity
		return ret
	}
	return o.CustomEntities
}

// GetCustomEntitiesOk returns a tuple with the CustomEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntitiesPaginatedList) GetCustomEntitiesOk() ([]CustomEntity, bool) {
	if o == nil || isNil(o.CustomEntities) {
    return nil, false
	}
	return o.CustomEntities, true
}

// HasCustomEntities returns a boolean if a field has been set.
func (o *CustomEntitiesPaginatedList) HasCustomEntities() bool {
	if o != nil && !isNil(o.CustomEntities) {
		return true
	}

	return false
}

// SetCustomEntities gets a reference to the given []CustomEntity and assigns it to the CustomEntities field.
func (o *CustomEntitiesPaginatedList) SetCustomEntities(v []CustomEntity) {
	o.CustomEntities = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *CustomEntitiesPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntitiesPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *CustomEntitiesPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *CustomEntitiesPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

func (o CustomEntitiesPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CustomEntities) {
		toSerialize["customEntities"] = o.CustomEntities
	}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEntitiesPaginatedList struct {
	value *CustomEntitiesPaginatedList
	isSet bool
}

func (v NullableCustomEntitiesPaginatedList) Get() *CustomEntitiesPaginatedList {
	return v.value
}

func (v *NullableCustomEntitiesPaginatedList) Set(val *CustomEntitiesPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEntitiesPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEntitiesPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEntitiesPaginatedList(val *CustomEntitiesPaginatedList) *NullableCustomEntitiesPaginatedList {
	return &NullableCustomEntitiesPaginatedList{value: val, isSet: true}
}

func (v NullableCustomEntitiesPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEntitiesPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

