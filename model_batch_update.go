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

// BatchUpdate struct for BatchUpdate
type BatchUpdate struct {
	DefaultConcentration *DefaultConcentrationSummary `json:"defaultConcentration,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
}

// NewBatchUpdate instantiates a new BatchUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchUpdate() *BatchUpdate {
	this := BatchUpdate{}
	return &this
}

// NewBatchUpdateWithDefaults instantiates a new BatchUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchUpdateWithDefaults() *BatchUpdate {
	this := BatchUpdate{}
	return &this
}

// GetDefaultConcentration returns the DefaultConcentration field value if set, zero value otherwise.
func (o *BatchUpdate) GetDefaultConcentration() DefaultConcentrationSummary {
	if o == nil || isNil(o.DefaultConcentration) {
		var ret DefaultConcentrationSummary
		return ret
	}
	return *o.DefaultConcentration
}

// GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdate) GetDefaultConcentrationOk() (*DefaultConcentrationSummary, bool) {
	if o == nil || isNil(o.DefaultConcentration) {
    return nil, false
	}
	return o.DefaultConcentration, true
}

// HasDefaultConcentration returns a boolean if a field has been set.
func (o *BatchUpdate) HasDefaultConcentration() bool {
	if o != nil && !isNil(o.DefaultConcentration) {
		return true
	}

	return false
}

// SetDefaultConcentration gets a reference to the given DefaultConcentrationSummary and assigns it to the DefaultConcentration field.
func (o *BatchUpdate) SetDefaultConcentration(v DefaultConcentrationSummary) {
	o.DefaultConcentration = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *BatchUpdate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *BatchUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *BatchUpdate) SetFields(v map[string]Field) {
	o.Fields = &v
}

func (o BatchUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultConcentration) {
		toSerialize["defaultConcentration"] = o.DefaultConcentration
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableBatchUpdate struct {
	value *BatchUpdate
	isSet bool
}

func (v NullableBatchUpdate) Get() *BatchUpdate {
	return v.value
}

func (v *NullableBatchUpdate) Set(val *BatchUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchUpdate(val *BatchUpdate) *NullableBatchUpdate {
	return &NullableBatchUpdate{value: val, isSet: true}
}

func (v NullableBatchUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


