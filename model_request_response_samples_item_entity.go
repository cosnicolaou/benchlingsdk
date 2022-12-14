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

// RequestResponseSamplesItemEntity struct for RequestResponseSamplesItemEntity
type RequestResponseSamplesItemEntity struct {
	EntityOrInaccessibleResource
}

// NewRequestResponseSamplesItemEntity instantiates a new RequestResponseSamplesItemEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResponseSamplesItemEntity() *RequestResponseSamplesItemEntity {
	this := RequestResponseSamplesItemEntity{}
	return &this
}

// NewRequestResponseSamplesItemEntityWithDefaults instantiates a new RequestResponseSamplesItemEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResponseSamplesItemEntityWithDefaults() *RequestResponseSamplesItemEntity {
	this := RequestResponseSamplesItemEntity{}
	return &this
}

func (o RequestResponseSamplesItemEntity) MarshalJSON() ([]byte, error) {
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

type NullableRequestResponseSamplesItemEntity struct {
	value *RequestResponseSamplesItemEntity
	isSet bool
}

func (v NullableRequestResponseSamplesItemEntity) Get() *RequestResponseSamplesItemEntity {
	return v.value
}

func (v *NullableRequestResponseSamplesItemEntity) Set(val *RequestResponseSamplesItemEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResponseSamplesItemEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResponseSamplesItemEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResponseSamplesItemEntity(val *RequestResponseSamplesItemEntity) *NullableRequestResponseSamplesItemEntity {
	return &NullableRequestResponseSamplesItemEntity{value: val, isSet: true}
}

func (v NullableRequestResponseSamplesItemEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResponseSamplesItemEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


