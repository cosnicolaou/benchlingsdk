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

// TransfersAsyncTaskAllOf struct for TransfersAsyncTaskAllOf
type TransfersAsyncTaskAllOf struct {
	Response *TransfersAsyncTaskAllOfResponse `json:"response,omitempty"`
}

// NewTransfersAsyncTaskAllOf instantiates a new TransfersAsyncTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfersAsyncTaskAllOf() *TransfersAsyncTaskAllOf {
	this := TransfersAsyncTaskAllOf{}
	return &this
}

// NewTransfersAsyncTaskAllOfWithDefaults instantiates a new TransfersAsyncTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransfersAsyncTaskAllOfWithDefaults() *TransfersAsyncTaskAllOf {
	this := TransfersAsyncTaskAllOf{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TransfersAsyncTaskAllOf) GetResponse() TransfersAsyncTaskAllOfResponse {
	if o == nil || isNil(o.Response) {
		var ret TransfersAsyncTaskAllOfResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransfersAsyncTaskAllOf) GetResponseOk() (*TransfersAsyncTaskAllOfResponse, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TransfersAsyncTaskAllOf) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given TransfersAsyncTaskAllOfResponse and assigns it to the Response field.
func (o *TransfersAsyncTaskAllOf) SetResponse(v TransfersAsyncTaskAllOfResponse) {
	o.Response = &v
}

func (o TransfersAsyncTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableTransfersAsyncTaskAllOf struct {
	value *TransfersAsyncTaskAllOf
	isSet bool
}

func (v NullableTransfersAsyncTaskAllOf) Get() *TransfersAsyncTaskAllOf {
	return v.value
}

func (v *NullableTransfersAsyncTaskAllOf) Set(val *TransfersAsyncTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfersAsyncTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfersAsyncTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfersAsyncTaskAllOf(val *TransfersAsyncTaskAllOf) *NullableTransfersAsyncTaskAllOf {
	return &NullableTransfersAsyncTaskAllOf{value: val, isSet: true}
}

func (v NullableTransfersAsyncTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfersAsyncTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


