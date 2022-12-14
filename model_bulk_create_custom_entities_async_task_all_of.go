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

// BulkCreateCustomEntitiesAsyncTaskAllOf struct for BulkCreateCustomEntitiesAsyncTaskAllOf
type BulkCreateCustomEntitiesAsyncTaskAllOf struct {
	Response *BulkCreateCustomEntitiesAsyncTaskAllOfResponse `json:"response,omitempty"`
}

// NewBulkCreateCustomEntitiesAsyncTaskAllOf instantiates a new BulkCreateCustomEntitiesAsyncTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCreateCustomEntitiesAsyncTaskAllOf() *BulkCreateCustomEntitiesAsyncTaskAllOf {
	this := BulkCreateCustomEntitiesAsyncTaskAllOf{}
	return &this
}

// NewBulkCreateCustomEntitiesAsyncTaskAllOfWithDefaults instantiates a new BulkCreateCustomEntitiesAsyncTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCreateCustomEntitiesAsyncTaskAllOfWithDefaults() *BulkCreateCustomEntitiesAsyncTaskAllOf {
	this := BulkCreateCustomEntitiesAsyncTaskAllOf{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *BulkCreateCustomEntitiesAsyncTaskAllOf) GetResponse() BulkCreateCustomEntitiesAsyncTaskAllOfResponse {
	if o == nil || isNil(o.Response) {
		var ret BulkCreateCustomEntitiesAsyncTaskAllOfResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateCustomEntitiesAsyncTaskAllOf) GetResponseOk() (*BulkCreateCustomEntitiesAsyncTaskAllOfResponse, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *BulkCreateCustomEntitiesAsyncTaskAllOf) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given BulkCreateCustomEntitiesAsyncTaskAllOfResponse and assigns it to the Response field.
func (o *BulkCreateCustomEntitiesAsyncTaskAllOf) SetResponse(v BulkCreateCustomEntitiesAsyncTaskAllOfResponse) {
	o.Response = &v
}

func (o BulkCreateCustomEntitiesAsyncTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableBulkCreateCustomEntitiesAsyncTaskAllOf struct {
	value *BulkCreateCustomEntitiesAsyncTaskAllOf
	isSet bool
}

func (v NullableBulkCreateCustomEntitiesAsyncTaskAllOf) Get() *BulkCreateCustomEntitiesAsyncTaskAllOf {
	return v.value
}

func (v *NullableBulkCreateCustomEntitiesAsyncTaskAllOf) Set(val *BulkCreateCustomEntitiesAsyncTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCreateCustomEntitiesAsyncTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCreateCustomEntitiesAsyncTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCreateCustomEntitiesAsyncTaskAllOf(val *BulkCreateCustomEntitiesAsyncTaskAllOf) *NullableBulkCreateCustomEntitiesAsyncTaskAllOf {
	return &NullableBulkCreateCustomEntitiesAsyncTaskAllOf{value: val, isSet: true}
}

func (v NullableBulkCreateCustomEntitiesAsyncTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCreateCustomEntitiesAsyncTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


