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

// ConflictErrorErrorAllOf struct for ConflictErrorErrorAllOf
type ConflictErrorErrorAllOf struct {
	Conflicts []map[string]interface{} `json:"conflicts,omitempty"`
}

// NewConflictErrorErrorAllOf instantiates a new ConflictErrorErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictErrorErrorAllOf() *ConflictErrorErrorAllOf {
	this := ConflictErrorErrorAllOf{}
	return &this
}

// NewConflictErrorErrorAllOfWithDefaults instantiates a new ConflictErrorErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictErrorErrorAllOfWithDefaults() *ConflictErrorErrorAllOf {
	this := ConflictErrorErrorAllOf{}
	return &this
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *ConflictErrorErrorAllOf) GetConflicts() []map[string]interface{} {
	if o == nil || isNil(o.Conflicts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictErrorErrorAllOf) GetConflictsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Conflicts) {
    return nil, false
	}
	return o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *ConflictErrorErrorAllOf) HasConflicts() bool {
	if o != nil && !isNil(o.Conflicts) {
		return true
	}

	return false
}

// SetConflicts gets a reference to the given []map[string]interface{} and assigns it to the Conflicts field.
func (o *ConflictErrorErrorAllOf) SetConflicts(v []map[string]interface{}) {
	o.Conflicts = v
}

func (o ConflictErrorErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Conflicts) {
		toSerialize["conflicts"] = o.Conflicts
	}
	return json.Marshal(toSerialize)
}

type NullableConflictErrorErrorAllOf struct {
	value *ConflictErrorErrorAllOf
	isSet bool
}

func (v NullableConflictErrorErrorAllOf) Get() *ConflictErrorErrorAllOf {
	return v.value
}

func (v *NullableConflictErrorErrorAllOf) Set(val *ConflictErrorErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictErrorErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictErrorErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictErrorErrorAllOf(val *ConflictErrorErrorAllOf) *NullableConflictErrorErrorAllOf {
	return &NullableConflictErrorErrorAllOf{value: val, isSet: true}
}

func (v NullableConflictErrorErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictErrorErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


