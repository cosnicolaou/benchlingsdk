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

// BlobsBulkGet struct for BlobsBulkGet
type BlobsBulkGet struct {
	Blobs []Blob `json:"blobs,omitempty"`
}

// NewBlobsBulkGet instantiates a new BlobsBulkGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobsBulkGet() *BlobsBulkGet {
	this := BlobsBulkGet{}
	return &this
}

// NewBlobsBulkGetWithDefaults instantiates a new BlobsBulkGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobsBulkGetWithDefaults() *BlobsBulkGet {
	this := BlobsBulkGet{}
	return &this
}

// GetBlobs returns the Blobs field value if set, zero value otherwise.
func (o *BlobsBulkGet) GetBlobs() []Blob {
	if o == nil || isNil(o.Blobs) {
		var ret []Blob
		return ret
	}
	return o.Blobs
}

// GetBlobsOk returns a tuple with the Blobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobsBulkGet) GetBlobsOk() ([]Blob, bool) {
	if o == nil || isNil(o.Blobs) {
    return nil, false
	}
	return o.Blobs, true
}

// HasBlobs returns a boolean if a field has been set.
func (o *BlobsBulkGet) HasBlobs() bool {
	if o != nil && !isNil(o.Blobs) {
		return true
	}

	return false
}

// SetBlobs gets a reference to the given []Blob and assigns it to the Blobs field.
func (o *BlobsBulkGet) SetBlobs(v []Blob) {
	o.Blobs = v
}

func (o BlobsBulkGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Blobs) {
		toSerialize["blobs"] = o.Blobs
	}
	return json.Marshal(toSerialize)
}

type NullableBlobsBulkGet struct {
	value *BlobsBulkGet
	isSet bool
}

func (v NullableBlobsBulkGet) Get() *BlobsBulkGet {
	return v.value
}

func (v *NullableBlobsBulkGet) Set(val *BlobsBulkGet) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobsBulkGet) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobsBulkGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobsBulkGet(val *BlobsBulkGet) *NullableBlobsBulkGet {
	return &NullableBlobsBulkGet{value: val, isSet: true}
}

func (v NullableBlobsBulkGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobsBulkGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


