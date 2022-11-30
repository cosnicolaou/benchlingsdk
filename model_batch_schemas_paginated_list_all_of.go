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

// BatchSchemasPaginatedListAllOf struct for BatchSchemasPaginatedListAllOf
type BatchSchemasPaginatedListAllOf struct {
	NextToken *string `json:"nextToken,omitempty"`
}

// NewBatchSchemasPaginatedListAllOf instantiates a new BatchSchemasPaginatedListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchSchemasPaginatedListAllOf() *BatchSchemasPaginatedListAllOf {
	this := BatchSchemasPaginatedListAllOf{}
	return &this
}

// NewBatchSchemasPaginatedListAllOfWithDefaults instantiates a new BatchSchemasPaginatedListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchSchemasPaginatedListAllOfWithDefaults() *BatchSchemasPaginatedListAllOf {
	this := BatchSchemasPaginatedListAllOf{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *BatchSchemasPaginatedListAllOf) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSchemasPaginatedListAllOf) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *BatchSchemasPaginatedListAllOf) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *BatchSchemasPaginatedListAllOf) SetNextToken(v string) {
	o.NextToken = &v
}

func (o BatchSchemasPaginatedListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableBatchSchemasPaginatedListAllOf struct {
	value *BatchSchemasPaginatedListAllOf
	isSet bool
}

func (v NullableBatchSchemasPaginatedListAllOf) Get() *BatchSchemasPaginatedListAllOf {
	return v.value
}

func (v *NullableBatchSchemasPaginatedListAllOf) Set(val *BatchSchemasPaginatedListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchSchemasPaginatedListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchSchemasPaginatedListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchSchemasPaginatedListAllOf(val *BatchSchemasPaginatedListAllOf) *NullableBatchSchemasPaginatedListAllOf {
	return &NullableBatchSchemasPaginatedListAllOf{value: val, isSet: true}
}

func (v NullableBatchSchemasPaginatedListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchSchemasPaginatedListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

