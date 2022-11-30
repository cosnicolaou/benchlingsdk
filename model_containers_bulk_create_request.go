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

// ContainersBulkCreateRequest struct for ContainersBulkCreateRequest
type ContainersBulkCreateRequest struct {
	Containers []ContainerCreate `json:"containers"`
}

// NewContainersBulkCreateRequest instantiates a new ContainersBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainersBulkCreateRequest(containers []ContainerCreate) *ContainersBulkCreateRequest {
	this := ContainersBulkCreateRequest{}
	this.Containers = containers
	return &this
}

// NewContainersBulkCreateRequestWithDefaults instantiates a new ContainersBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainersBulkCreateRequestWithDefaults() *ContainersBulkCreateRequest {
	this := ContainersBulkCreateRequest{}
	return &this
}

// GetContainers returns the Containers field value
func (o *ContainersBulkCreateRequest) GetContainers() []ContainerCreate {
	if o == nil {
		var ret []ContainerCreate
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
func (o *ContainersBulkCreateRequest) GetContainersOk() ([]ContainerCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ContainersBulkCreateRequest) SetContainers(v []ContainerCreate) {
	o.Containers = v
}

func (o ContainersBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containers"] = o.Containers
	}
	return json.Marshal(toSerialize)
}

type NullableContainersBulkCreateRequest struct {
	value *ContainersBulkCreateRequest
	isSet bool
}

func (v NullableContainersBulkCreateRequest) Get() *ContainersBulkCreateRequest {
	return v.value
}

func (v *NullableContainersBulkCreateRequest) Set(val *ContainersBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainersBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainersBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainersBulkCreateRequest(val *ContainersBulkCreateRequest) *NullableContainersBulkCreateRequest {
	return &NullableContainersBulkCreateRequest{value: val, isSet: true}
}

func (v NullableContainersBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainersBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


