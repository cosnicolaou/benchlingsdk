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

// AaSequenceBulkUpdateAllOf struct for AaSequenceBulkUpdateAllOf
type AaSequenceBulkUpdateAllOf struct {
	Id *string `json:"id,omitempty"`
}

// NewAaSequenceBulkUpdateAllOf instantiates a new AaSequenceBulkUpdateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaSequenceBulkUpdateAllOf() *AaSequenceBulkUpdateAllOf {
	this := AaSequenceBulkUpdateAllOf{}
	return &this
}

// NewAaSequenceBulkUpdateAllOfWithDefaults instantiates a new AaSequenceBulkUpdateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaSequenceBulkUpdateAllOfWithDefaults() *AaSequenceBulkUpdateAllOf {
	this := AaSequenceBulkUpdateAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AaSequenceBulkUpdateAllOf) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceBulkUpdateAllOf) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AaSequenceBulkUpdateAllOf) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AaSequenceBulkUpdateAllOf) SetId(v string) {
	o.Id = &v
}

func (o AaSequenceBulkUpdateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAaSequenceBulkUpdateAllOf struct {
	value *AaSequenceBulkUpdateAllOf
	isSet bool
}

func (v NullableAaSequenceBulkUpdateAllOf) Get() *AaSequenceBulkUpdateAllOf {
	return v.value
}

func (v *NullableAaSequenceBulkUpdateAllOf) Set(val *AaSequenceBulkUpdateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAaSequenceBulkUpdateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAaSequenceBulkUpdateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaSequenceBulkUpdateAllOf(val *AaSequenceBulkUpdateAllOf) *NullableAaSequenceBulkUpdateAllOf {
	return &NullableAaSequenceBulkUpdateAllOf{value: val, isSet: true}
}

func (v NullableAaSequenceBulkUpdateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaSequenceBulkUpdateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


