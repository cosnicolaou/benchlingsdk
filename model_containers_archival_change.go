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

// ContainersArchivalChange IDs of all items that were unarchived, grouped by resource type. This includes the IDs of containers that were unarchived. 
type ContainersArchivalChange struct {
	ContainerIds []string `json:"containerIds,omitempty"`
}

// NewContainersArchivalChange instantiates a new ContainersArchivalChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainersArchivalChange() *ContainersArchivalChange {
	this := ContainersArchivalChange{}
	return &this
}

// NewContainersArchivalChangeWithDefaults instantiates a new ContainersArchivalChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainersArchivalChangeWithDefaults() *ContainersArchivalChange {
	this := ContainersArchivalChange{}
	return &this
}

// GetContainerIds returns the ContainerIds field value if set, zero value otherwise.
func (o *ContainersArchivalChange) GetContainerIds() []string {
	if o == nil || isNil(o.ContainerIds) {
		var ret []string
		return ret
	}
	return o.ContainerIds
}

// GetContainerIdsOk returns a tuple with the ContainerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersArchivalChange) GetContainerIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ContainerIds) {
    return nil, false
	}
	return o.ContainerIds, true
}

// HasContainerIds returns a boolean if a field has been set.
func (o *ContainersArchivalChange) HasContainerIds() bool {
	if o != nil && !isNil(o.ContainerIds) {
		return true
	}

	return false
}

// SetContainerIds gets a reference to the given []string and assigns it to the ContainerIds field.
func (o *ContainersArchivalChange) SetContainerIds(v []string) {
	o.ContainerIds = v
}

func (o ContainersArchivalChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContainerIds) {
		toSerialize["containerIds"] = o.ContainerIds
	}
	return json.Marshal(toSerialize)
}

type NullableContainersArchivalChange struct {
	value *ContainersArchivalChange
	isSet bool
}

func (v NullableContainersArchivalChange) Get() *ContainersArchivalChange {
	return v.value
}

func (v *NullableContainersArchivalChange) Set(val *ContainersArchivalChange) {
	v.value = val
	v.isSet = true
}

func (v NullableContainersArchivalChange) IsSet() bool {
	return v.isSet
}

func (v *NullableContainersArchivalChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainersArchivalChange(val *ContainersArchivalChange) *NullableContainersArchivalChange {
	return &NullableContainersArchivalChange{value: val, isSet: true}
}

func (v NullableContainersArchivalChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainersArchivalChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

