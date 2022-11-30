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

// BatchesArchive The request body for archiving Batches. 
type BatchesArchive struct {
	BatchIds []string `json:"batchIds"`
	// The reason for archiving the provided Batches. Accepted reasons may differ based on tenant configuration. 
	Reason string `json:"reason"`
}

// NewBatchesArchive instantiates a new BatchesArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchesArchive(batchIds []string, reason string) *BatchesArchive {
	this := BatchesArchive{}
	this.BatchIds = batchIds
	this.Reason = reason
	return &this
}

// NewBatchesArchiveWithDefaults instantiates a new BatchesArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchesArchiveWithDefaults() *BatchesArchive {
	this := BatchesArchive{}
	return &this
}

// GetBatchIds returns the BatchIds field value
func (o *BatchesArchive) GetBatchIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BatchIds
}

// GetBatchIdsOk returns a tuple with the BatchIds field value
// and a boolean to check if the value has been set.
func (o *BatchesArchive) GetBatchIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BatchIds, true
}

// SetBatchIds sets field value
func (o *BatchesArchive) SetBatchIds(v []string) {
	o.BatchIds = v
}

// GetReason returns the Reason field value
func (o *BatchesArchive) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *BatchesArchive) GetReasonOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *BatchesArchive) SetReason(v string) {
	o.Reason = v
}

func (o BatchesArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["batchIds"] = o.BatchIds
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableBatchesArchive struct {
	value *BatchesArchive
	isSet bool
}

func (v NullableBatchesArchive) Get() *BatchesArchive {
	return v.value
}

func (v *NullableBatchesArchive) Set(val *BatchesArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchesArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchesArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchesArchive(val *BatchesArchive) *NullableBatchesArchive {
	return &NullableBatchesArchive{value: val, isSet: true}
}

func (v NullableBatchesArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchesArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

