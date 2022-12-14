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

// RequestResponseSamplesInnerEntity The sample, if it is an entity resource. Null otherwise.
type RequestResponseSamplesInnerEntity struct {
	RequestResponseSamplesItemEntity
}

// NewRequestResponseSamplesInnerEntity instantiates a new RequestResponseSamplesInnerEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResponseSamplesInnerEntity() *RequestResponseSamplesInnerEntity {
	this := RequestResponseSamplesInnerEntity{}
	return &this
}

// NewRequestResponseSamplesInnerEntityWithDefaults instantiates a new RequestResponseSamplesInnerEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResponseSamplesInnerEntityWithDefaults() *RequestResponseSamplesInnerEntity {
	this := RequestResponseSamplesInnerEntity{}
	return &this
}

func (o RequestResponseSamplesInnerEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRequestResponseSamplesItemEntity, errRequestResponseSamplesItemEntity := json.Marshal(o.RequestResponseSamplesItemEntity)
	if errRequestResponseSamplesItemEntity != nil {
		return []byte{}, errRequestResponseSamplesItemEntity
	}
	errRequestResponseSamplesItemEntity = json.Unmarshal([]byte(serializedRequestResponseSamplesItemEntity), &toSerialize)
	if errRequestResponseSamplesItemEntity != nil {
		return []byte{}, errRequestResponseSamplesItemEntity
	}
	return json.Marshal(toSerialize)
}

type NullableRequestResponseSamplesInnerEntity struct {
	value *RequestResponseSamplesInnerEntity
	isSet bool
}

func (v NullableRequestResponseSamplesInnerEntity) Get() *RequestResponseSamplesInnerEntity {
	return v.value
}

func (v *NullableRequestResponseSamplesInnerEntity) Set(val *RequestResponseSamplesInnerEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResponseSamplesInnerEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResponseSamplesInnerEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResponseSamplesInnerEntity(val *RequestResponseSamplesInnerEntity) *NullableRequestResponseSamplesInnerEntity {
	return &NullableRequestResponseSamplesInnerEntity{value: val, isSet: true}
}

func (v NullableRequestResponseSamplesInnerEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResponseSamplesInnerEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


