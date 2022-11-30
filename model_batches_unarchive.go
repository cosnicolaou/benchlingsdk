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

// BatchesUnarchive The request body for unarchiving Batches. 
type BatchesUnarchive struct {
	BatchIds []string `json:"batchIds"`
}

// NewBatchesUnarchive instantiates a new BatchesUnarchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchesUnarchive(batchIds []string) *BatchesUnarchive {
	this := BatchesUnarchive{}
	this.BatchIds = batchIds
	return &this
}

// NewBatchesUnarchiveWithDefaults instantiates a new BatchesUnarchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchesUnarchiveWithDefaults() *BatchesUnarchive {
	this := BatchesUnarchive{}
	return &this
}

// GetBatchIds returns the BatchIds field value
func (o *BatchesUnarchive) GetBatchIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BatchIds
}

// GetBatchIdsOk returns a tuple with the BatchIds field value
// and a boolean to check if the value has been set.
func (o *BatchesUnarchive) GetBatchIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BatchIds, true
}

// SetBatchIds sets field value
func (o *BatchesUnarchive) SetBatchIds(v []string) {
	o.BatchIds = v
}

func (o BatchesUnarchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["batchIds"] = o.BatchIds
	}
	return json.Marshal(toSerialize)
}

type NullableBatchesUnarchive struct {
	value *BatchesUnarchive
	isSet bool
}

func (v NullableBatchesUnarchive) Get() *BatchesUnarchive {
	return v.value
}

func (v *NullableBatchesUnarchive) Set(val *BatchesUnarchive) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchesUnarchive) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchesUnarchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchesUnarchive(val *BatchesUnarchive) *NullableBatchesUnarchive {
	return &NullableBatchesUnarchive{value: val, isSet: true}
}

func (v NullableBatchesUnarchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchesUnarchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


