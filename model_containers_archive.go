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

// ContainersArchive struct for ContainersArchive
type ContainersArchive struct {
	// Array of container IDs
	ContainerIds []string `json:"containerIds"`
	// Reason that containers are being archived. 
	Reason string `json:"reason"`
	// Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items. 
	ShouldRemoveBarcodes *bool `json:"shouldRemoveBarcodes,omitempty"`
}

// NewContainersArchive instantiates a new ContainersArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainersArchive(containerIds []string, reason string) *ContainersArchive {
	this := ContainersArchive{}
	this.ContainerIds = containerIds
	this.Reason = reason
	return &this
}

// NewContainersArchiveWithDefaults instantiates a new ContainersArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainersArchiveWithDefaults() *ContainersArchive {
	this := ContainersArchive{}
	return &this
}

// GetContainerIds returns the ContainerIds field value
func (o *ContainersArchive) GetContainerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContainerIds
}

// GetContainerIdsOk returns a tuple with the ContainerIds field value
// and a boolean to check if the value has been set.
func (o *ContainersArchive) GetContainerIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContainerIds, true
}

// SetContainerIds sets field value
func (o *ContainersArchive) SetContainerIds(v []string) {
	o.ContainerIds = v
}

// GetReason returns the Reason field value
func (o *ContainersArchive) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ContainersArchive) GetReasonOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ContainersArchive) SetReason(v string) {
	o.Reason = v
}

// GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field value if set, zero value otherwise.
func (o *ContainersArchive) GetShouldRemoveBarcodes() bool {
	if o == nil || isNil(o.ShouldRemoveBarcodes) {
		var ret bool
		return ret
	}
	return *o.ShouldRemoveBarcodes
}

// GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainersArchive) GetShouldRemoveBarcodesOk() (*bool, bool) {
	if o == nil || isNil(o.ShouldRemoveBarcodes) {
    return nil, false
	}
	return o.ShouldRemoveBarcodes, true
}

// HasShouldRemoveBarcodes returns a boolean if a field has been set.
func (o *ContainersArchive) HasShouldRemoveBarcodes() bool {
	if o != nil && !isNil(o.ShouldRemoveBarcodes) {
		return true
	}

	return false
}

// SetShouldRemoveBarcodes gets a reference to the given bool and assigns it to the ShouldRemoveBarcodes field.
func (o *ContainersArchive) SetShouldRemoveBarcodes(v bool) {
	o.ShouldRemoveBarcodes = &v
}

func (o ContainersArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["containerIds"] = o.ContainerIds
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.ShouldRemoveBarcodes) {
		toSerialize["shouldRemoveBarcodes"] = o.ShouldRemoveBarcodes
	}
	return json.Marshal(toSerialize)
}

type NullableContainersArchive struct {
	value *ContainersArchive
	isSet bool
}

func (v NullableContainersArchive) Get() *ContainersArchive {
	return v.value
}

func (v *NullableContainersArchive) Set(val *ContainersArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableContainersArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableContainersArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainersArchive(val *ContainersArchive) *NullableContainersArchive {
	return &NullableContainersArchive{value: val, isSet: true}
}

func (v NullableContainersArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainersArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


