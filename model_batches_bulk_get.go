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

// BatchesBulkGet struct for BatchesBulkGet
type BatchesBulkGet struct {
	Batches []Batch `json:"batches,omitempty"`
}

// NewBatchesBulkGet instantiates a new BatchesBulkGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchesBulkGet() *BatchesBulkGet {
	this := BatchesBulkGet{}
	return &this
}

// NewBatchesBulkGetWithDefaults instantiates a new BatchesBulkGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchesBulkGetWithDefaults() *BatchesBulkGet {
	this := BatchesBulkGet{}
	return &this
}

// GetBatches returns the Batches field value if set, zero value otherwise.
func (o *BatchesBulkGet) GetBatches() []Batch {
	if o == nil || isNil(o.Batches) {
		var ret []Batch
		return ret
	}
	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchesBulkGet) GetBatchesOk() ([]Batch, bool) {
	if o == nil || isNil(o.Batches) {
    return nil, false
	}
	return o.Batches, true
}

// HasBatches returns a boolean if a field has been set.
func (o *BatchesBulkGet) HasBatches() bool {
	if o != nil && !isNil(o.Batches) {
		return true
	}

	return false
}

// SetBatches gets a reference to the given []Batch and assigns it to the Batches field.
func (o *BatchesBulkGet) SetBatches(v []Batch) {
	o.Batches = v
}

func (o BatchesBulkGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Batches) {
		toSerialize["batches"] = o.Batches
	}
	return json.Marshal(toSerialize)
}

type NullableBatchesBulkGet struct {
	value *BatchesBulkGet
	isSet bool
}

func (v NullableBatchesBulkGet) Get() *BatchesBulkGet {
	return v.value
}

func (v *NullableBatchesBulkGet) Set(val *BatchesBulkGet) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchesBulkGet) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchesBulkGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchesBulkGet(val *BatchesBulkGet) *NullableBatchesBulkGet {
	return &NullableBatchesBulkGet{value: val, isSet: true}
}

func (v NullableBatchesBulkGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchesBulkGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


