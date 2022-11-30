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

// AssayResultSchemasPaginatedList struct for AssayResultSchemasPaginatedList
type AssayResultSchemasPaginatedList struct {
	AssayResultSchemas []AssayResultSchema `json:"assayResultSchemas,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// NewAssayResultSchemasPaginatedList instantiates a new AssayResultSchemasPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResultSchemasPaginatedList() *AssayResultSchemasPaginatedList {
	this := AssayResultSchemasPaginatedList{}
	return &this
}

// NewAssayResultSchemasPaginatedListWithDefaults instantiates a new AssayResultSchemasPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultSchemasPaginatedListWithDefaults() *AssayResultSchemasPaginatedList {
	this := AssayResultSchemasPaginatedList{}
	return &this
}

// GetAssayResultSchemas returns the AssayResultSchemas field value if set, zero value otherwise.
func (o *AssayResultSchemasPaginatedList) GetAssayResultSchemas() []AssayResultSchema {
	if o == nil || isNil(o.AssayResultSchemas) {
		var ret []AssayResultSchema
		return ret
	}
	return o.AssayResultSchemas
}

// GetAssayResultSchemasOk returns a tuple with the AssayResultSchemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchemasPaginatedList) GetAssayResultSchemasOk() ([]AssayResultSchema, bool) {
	if o == nil || isNil(o.AssayResultSchemas) {
    return nil, false
	}
	return o.AssayResultSchemas, true
}

// HasAssayResultSchemas returns a boolean if a field has been set.
func (o *AssayResultSchemasPaginatedList) HasAssayResultSchemas() bool {
	if o != nil && !isNil(o.AssayResultSchemas) {
		return true
	}

	return false
}

// SetAssayResultSchemas gets a reference to the given []AssayResultSchema and assigns it to the AssayResultSchemas field.
func (o *AssayResultSchemasPaginatedList) SetAssayResultSchemas(v []AssayResultSchema) {
	o.AssayResultSchemas = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *AssayResultSchemasPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchemasPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *AssayResultSchemasPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *AssayResultSchemasPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

func (o AssayResultSchemasPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssayResultSchemas) {
		toSerialize["assayResultSchemas"] = o.AssayResultSchemas
	}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResultSchemasPaginatedList struct {
	value *AssayResultSchemasPaginatedList
	isSet bool
}

func (v NullableAssayResultSchemasPaginatedList) Get() *AssayResultSchemasPaginatedList {
	return v.value
}

func (v *NullableAssayResultSchemasPaginatedList) Set(val *AssayResultSchemasPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResultSchemasPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResultSchemasPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResultSchemasPaginatedList(val *AssayResultSchemasPaginatedList) *NullableAssayResultSchemasPaginatedList {
	return &NullableAssayResultSchemasPaginatedList{value: val, isSet: true}
}

func (v NullableAssayResultSchemasPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResultSchemasPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

