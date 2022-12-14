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

// ContainerContentEntity struct for ContainerContentEntity
type ContainerContentEntity struct {
	EntityOrInaccessibleResource
}

// NewContainerContentEntity instantiates a new ContainerContentEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContentEntity() *ContainerContentEntity {
	this := ContainerContentEntity{}
	return &this
}

// NewContainerContentEntityWithDefaults instantiates a new ContainerContentEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContentEntityWithDefaults() *ContainerContentEntity {
	this := ContainerContentEntity{}
	return &this
}

func (o ContainerContentEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEntityOrInaccessibleResource, errEntityOrInaccessibleResource := json.Marshal(o.EntityOrInaccessibleResource)
	if errEntityOrInaccessibleResource != nil {
		return []byte{}, errEntityOrInaccessibleResource
	}
	errEntityOrInaccessibleResource = json.Unmarshal([]byte(serializedEntityOrInaccessibleResource), &toSerialize)
	if errEntityOrInaccessibleResource != nil {
		return []byte{}, errEntityOrInaccessibleResource
	}
	return json.Marshal(toSerialize)
}

type NullableContainerContentEntity struct {
	value *ContainerContentEntity
	isSet bool
}

func (v NullableContainerContentEntity) Get() *ContainerContentEntity {
	return v.value
}

func (v *NullableContainerContentEntity) Set(val *ContainerContentEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContentEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContentEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContentEntity(val *ContainerContentEntity) *NullableContainerContentEntity {
	return &NullableContainerContentEntity{value: val, isSet: true}
}

func (v NullableContainerContentEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContentEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


