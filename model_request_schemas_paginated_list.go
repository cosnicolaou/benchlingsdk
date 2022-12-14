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

// RequestSchemasPaginatedList struct for RequestSchemasPaginatedList
type RequestSchemasPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	RequestSchemas []RequestSchema `json:"requestSchemas,omitempty"`
}

// NewRequestSchemasPaginatedList instantiates a new RequestSchemasPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSchemasPaginatedList() *RequestSchemasPaginatedList {
	this := RequestSchemasPaginatedList{}
	return &this
}

// NewRequestSchemasPaginatedListWithDefaults instantiates a new RequestSchemasPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSchemasPaginatedListWithDefaults() *RequestSchemasPaginatedList {
	this := RequestSchemasPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *RequestSchemasPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSchemasPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *RequestSchemasPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *RequestSchemasPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetRequestSchemas returns the RequestSchemas field value if set, zero value otherwise.
func (o *RequestSchemasPaginatedList) GetRequestSchemas() []RequestSchema {
	if o == nil || isNil(o.RequestSchemas) {
		var ret []RequestSchema
		return ret
	}
	return o.RequestSchemas
}

// GetRequestSchemasOk returns a tuple with the RequestSchemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSchemasPaginatedList) GetRequestSchemasOk() ([]RequestSchema, bool) {
	if o == nil || isNil(o.RequestSchemas) {
    return nil, false
	}
	return o.RequestSchemas, true
}

// HasRequestSchemas returns a boolean if a field has been set.
func (o *RequestSchemasPaginatedList) HasRequestSchemas() bool {
	if o != nil && !isNil(o.RequestSchemas) {
		return true
	}

	return false
}

// SetRequestSchemas gets a reference to the given []RequestSchema and assigns it to the RequestSchemas field.
func (o *RequestSchemasPaginatedList) SetRequestSchemas(v []RequestSchema) {
	o.RequestSchemas = v
}

func (o RequestSchemasPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.RequestSchemas) {
		toSerialize["requestSchemas"] = o.RequestSchemas
	}
	return json.Marshal(toSerialize)
}

type NullableRequestSchemasPaginatedList struct {
	value *RequestSchemasPaginatedList
	isSet bool
}

func (v NullableRequestSchemasPaginatedList) Get() *RequestSchemasPaginatedList {
	return v.value
}

func (v *NullableRequestSchemasPaginatedList) Set(val *RequestSchemasPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSchemasPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSchemasPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSchemasPaginatedList(val *RequestSchemasPaginatedList) *NullableRequestSchemasPaginatedList {
	return &NullableRequestSchemasPaginatedList{value: val, isSet: true}
}

func (v NullableRequestSchemasPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSchemasPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


