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

// RequestCreateAllOf struct for RequestCreateAllOf
type RequestCreateAllOf struct {
	// ID of the request's schema.
	SchemaId string `json:"schemaId"`
}

// NewRequestCreateAllOf instantiates a new RequestCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreateAllOf(schemaId string) *RequestCreateAllOf {
	this := RequestCreateAllOf{}
	this.SchemaId = schemaId
	return &this
}

// NewRequestCreateAllOfWithDefaults instantiates a new RequestCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreateAllOfWithDefaults() *RequestCreateAllOf {
	this := RequestCreateAllOf{}
	return &this
}

// GetSchemaId returns the SchemaId field value
func (o *RequestCreateAllOf) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *RequestCreateAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *RequestCreateAllOf) SetSchemaId(v string) {
	o.SchemaId = v
}

func (o RequestCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableRequestCreateAllOf struct {
	value *RequestCreateAllOf
	isSet bool
}

func (v NullableRequestCreateAllOf) Get() *RequestCreateAllOf {
	return v.value
}

func (v *NullableRequestCreateAllOf) Set(val *RequestCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreateAllOf(val *RequestCreateAllOf) *NullableRequestCreateAllOf {
	return &NullableRequestCreateAllOf{value: val, isSet: true}
}

func (v NullableRequestCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

