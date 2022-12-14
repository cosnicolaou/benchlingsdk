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

// RequestTasksBulkCreate struct for RequestTasksBulkCreate
type RequestTasksBulkCreate struct {
	// The schema id of the task to create
	SchemaId string `json:"schemaId"`
	// Schema fields to set on the request task. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	Fields *ModelMap `json:"fields,omitempty"`
	// IDs of all request sample groups now associated with this task.
	SampleGroupIds []string `json:"sampleGroupIds,omitempty"`
}

// NewRequestTasksBulkCreate instantiates a new RequestTasksBulkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTasksBulkCreate(schemaId string) *RequestTasksBulkCreate {
	this := RequestTasksBulkCreate{}
	return &this
}

// NewRequestTasksBulkCreateWithDefaults instantiates a new RequestTasksBulkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTasksBulkCreateWithDefaults() *RequestTasksBulkCreate {
	this := RequestTasksBulkCreate{}
	return &this
}

// GetSchemaId returns the SchemaId field value
func (o *RequestTasksBulkCreate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *RequestTasksBulkCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *RequestTasksBulkCreate) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *RequestTasksBulkCreate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTasksBulkCreate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *RequestTasksBulkCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *RequestTasksBulkCreate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetSampleGroupIds returns the SampleGroupIds field value if set, zero value otherwise.
func (o *RequestTasksBulkCreate) GetSampleGroupIds() []string {
	if o == nil || isNil(o.SampleGroupIds) {
		var ret []string
		return ret
	}
	return o.SampleGroupIds
}

// GetSampleGroupIdsOk returns a tuple with the SampleGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestTasksBulkCreate) GetSampleGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SampleGroupIds) {
    return nil, false
	}
	return o.SampleGroupIds, true
}

// HasSampleGroupIds returns a boolean if a field has been set.
func (o *RequestTasksBulkCreate) HasSampleGroupIds() bool {
	if o != nil && !isNil(o.SampleGroupIds) {
		return true
	}

	return false
}

// SetSampleGroupIds gets a reference to the given []string and assigns it to the SampleGroupIds field.
func (o *RequestTasksBulkCreate) SetSampleGroupIds(v []string) {
	o.SampleGroupIds = v
}

func (o RequestTasksBulkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.SampleGroupIds) {
		toSerialize["sampleGroupIds"] = o.SampleGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableRequestTasksBulkCreate struct {
	value *RequestTasksBulkCreate
	isSet bool
}

func (v NullableRequestTasksBulkCreate) Get() *RequestTasksBulkCreate {
	return v.value
}

func (v *NullableRequestTasksBulkCreate) Set(val *RequestTasksBulkCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTasksBulkCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTasksBulkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTasksBulkCreate(val *RequestTasksBulkCreate) *NullableRequestTasksBulkCreate {
	return &NullableRequestTasksBulkCreate{value: val, isSet: true}
}

func (v NullableRequestTasksBulkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTasksBulkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


