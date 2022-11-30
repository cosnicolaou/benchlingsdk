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

// ContainersCheckout struct for ContainersCheckout
type ContainersCheckout struct {
	// User or Team API ID.
	AssigneeId string `json:"assigneeId"`
	Comment *string `json:"comment,omitempty"`
	// Array of container IDs.
	ContainerIds []string `json:"containerIds"`
}

// NewContainersCheckout instantiates a new ContainersCheckout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainersCheckout(assigneeId string, containerIds []string) *ContainersCheckout {
	this := ContainersCheckout{}
	this.AssigneeId = assigneeId
	this.ContainerIds = containerIds
	return &this
}

// NewContainersCheckoutWithDefaults instantiates a new ContainersCheckout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainersCheckoutWithDefaults() *ContainersCheckout {
	this := ContainersCheckout{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value
func (o *ContainersCheckout) GetAssigneeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value
// and a boolean to check if the value has been set.
func (o *ContainersCheckout) GetAssigneeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AssigneeId, true
}

// SetAssigneeId sets field value
func (o *ContainersCheckout) SetAssigneeId(v string) {
	o.AssigneeId = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ContainersCheckout) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersCheckout) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ContainersCheckout) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ContainersCheckout) SetComment(v string) {
	o.Comment = &v
}

// GetContainerIds returns the ContainerIds field value
func (o *ContainersCheckout) GetContainerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContainerIds
}

// GetContainerIdsOk returns a tuple with the ContainerIds field value
// and a boolean to check if the value has been set.
func (o *ContainersCheckout) GetContainerIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContainerIds, true
}

// SetContainerIds sets field value
func (o *ContainersCheckout) SetContainerIds(v []string) {
	o.ContainerIds = v
}

func (o ContainersCheckout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assigneeId"] = o.AssigneeId
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["containerIds"] = o.ContainerIds
	}
	return json.Marshal(toSerialize)
}

type NullableContainersCheckout struct {
	value *ContainersCheckout
	isSet bool
}

func (v NullableContainersCheckout) Get() *ContainersCheckout {
	return v.value
}

func (v *NullableContainersCheckout) Set(val *ContainersCheckout) {
	v.value = val
	v.isSet = true
}

func (v NullableContainersCheckout) IsSet() bool {
	return v.isSet
}

func (v *NullableContainersCheckout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainersCheckout(val *ContainersCheckout) *NullableContainersCheckout {
	return &NullableContainersCheckout{value: val, isSet: true}
}

func (v NullableContainersCheckout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainersCheckout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

