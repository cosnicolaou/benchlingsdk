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

// BulkCreateContainersAsyncTaskAllOfResponse struct for BulkCreateContainersAsyncTaskAllOfResponse
type BulkCreateContainersAsyncTaskAllOfResponse struct {
	Containers []Container `json:"containers,omitempty"`
}

// NewBulkCreateContainersAsyncTaskAllOfResponse instantiates a new BulkCreateContainersAsyncTaskAllOfResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCreateContainersAsyncTaskAllOfResponse() *BulkCreateContainersAsyncTaskAllOfResponse {
	this := BulkCreateContainersAsyncTaskAllOfResponse{}
	return &this
}

// NewBulkCreateContainersAsyncTaskAllOfResponseWithDefaults instantiates a new BulkCreateContainersAsyncTaskAllOfResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCreateContainersAsyncTaskAllOfResponseWithDefaults() *BulkCreateContainersAsyncTaskAllOfResponse {
	this := BulkCreateContainersAsyncTaskAllOfResponse{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *BulkCreateContainersAsyncTaskAllOfResponse) GetContainers() []Container {
	if o == nil || isNil(o.Containers) {
		var ret []Container
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateContainersAsyncTaskAllOfResponse) GetContainersOk() ([]Container, bool) {
	if o == nil || isNil(o.Containers) {
    return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *BulkCreateContainersAsyncTaskAllOfResponse) HasContainers() bool {
	if o != nil && !isNil(o.Containers) {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []Container and assigns it to the Containers field.
func (o *BulkCreateContainersAsyncTaskAllOfResponse) SetContainers(v []Container) {
	o.Containers = v
}

func (o BulkCreateContainersAsyncTaskAllOfResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Containers) {
		toSerialize["containers"] = o.Containers
	}
	return json.Marshal(toSerialize)
}

type NullableBulkCreateContainersAsyncTaskAllOfResponse struct {
	value *BulkCreateContainersAsyncTaskAllOfResponse
	isSet bool
}

func (v NullableBulkCreateContainersAsyncTaskAllOfResponse) Get() *BulkCreateContainersAsyncTaskAllOfResponse {
	return v.value
}

func (v *NullableBulkCreateContainersAsyncTaskAllOfResponse) Set(val *BulkCreateContainersAsyncTaskAllOfResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCreateContainersAsyncTaskAllOfResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCreateContainersAsyncTaskAllOfResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCreateContainersAsyncTaskAllOfResponse(val *BulkCreateContainersAsyncTaskAllOfResponse) *NullableBulkCreateContainersAsyncTaskAllOfResponse {
	return &NullableBulkCreateContainersAsyncTaskAllOfResponse{value: val, isSet: true}
}

func (v NullableBulkCreateContainersAsyncTaskAllOfResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCreateContainersAsyncTaskAllOfResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

