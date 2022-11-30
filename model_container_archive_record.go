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

// ContainerArchiveRecord struct for ContainerArchiveRecord
type ContainerArchiveRecord struct {
	Reason *string `json:"reason,omitempty"`
}

// NewContainerArchiveRecord instantiates a new ContainerArchiveRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerArchiveRecord() *ContainerArchiveRecord {
	this := ContainerArchiveRecord{}
	return &this
}

// NewContainerArchiveRecordWithDefaults instantiates a new ContainerArchiveRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerArchiveRecordWithDefaults() *ContainerArchiveRecord {
	this := ContainerArchiveRecord{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ContainerArchiveRecord) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerArchiveRecord) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ContainerArchiveRecord) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ContainerArchiveRecord) SetReason(v string) {
	o.Reason = &v
}

func (o ContainerArchiveRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableContainerArchiveRecord struct {
	value *ContainerArchiveRecord
	isSet bool
}

func (v NullableContainerArchiveRecord) Get() *ContainerArchiveRecord {
	return v.value
}

func (v *NullableContainerArchiveRecord) Set(val *ContainerArchiveRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerArchiveRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerArchiveRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerArchiveRecord(val *ContainerArchiveRecord) *NullableContainerArchiveRecord {
	return &NullableContainerArchiveRecord{value: val, isSet: true}
}

func (v NullableContainerArchiveRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerArchiveRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


