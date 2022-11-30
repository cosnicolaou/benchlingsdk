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

// ContainerContentsList struct for ContainerContentsList
type ContainerContentsList struct {
	Contents []ContainerContent `json:"contents,omitempty"`
}

// NewContainerContentsList instantiates a new ContainerContentsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContentsList() *ContainerContentsList {
	this := ContainerContentsList{}
	return &this
}

// NewContainerContentsListWithDefaults instantiates a new ContainerContentsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContentsListWithDefaults() *ContainerContentsList {
	this := ContainerContentsList{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *ContainerContentsList) GetContents() []ContainerContent {
	if o == nil || isNil(o.Contents) {
		var ret []ContainerContent
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentsList) GetContentsOk() ([]ContainerContent, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *ContainerContentsList) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []ContainerContent and assigns it to the Contents field.
func (o *ContainerContentsList) SetContents(v []ContainerContent) {
	o.Contents = v
}

func (o ContainerContentsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableContainerContentsList struct {
	value *ContainerContentsList
	isSet bool
}

func (v NullableContainerContentsList) Get() *ContainerContentsList {
	return v.value
}

func (v *NullableContainerContentsList) Set(val *ContainerContentsList) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContentsList) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContentsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContentsList(val *ContainerContentsList) *NullableContainerContentsList {
	return &NullableContainerContentsList{value: val, isSet: true}
}

func (v NullableContainerContentsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContentsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


