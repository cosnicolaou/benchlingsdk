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

// AssayResultIdsResponse struct for AssayResultIdsResponse
type AssayResultIdsResponse struct {
	AssayResultIds []string `json:"assayResultIds,omitempty"`
}

// NewAssayResultIdsResponse instantiates a new AssayResultIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResultIdsResponse() *AssayResultIdsResponse {
	this := AssayResultIdsResponse{}
	return &this
}

// NewAssayResultIdsResponseWithDefaults instantiates a new AssayResultIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultIdsResponseWithDefaults() *AssayResultIdsResponse {
	this := AssayResultIdsResponse{}
	return &this
}

// GetAssayResultIds returns the AssayResultIds field value if set, zero value otherwise.
func (o *AssayResultIdsResponse) GetAssayResultIds() []string {
	if o == nil || isNil(o.AssayResultIds) {
		var ret []string
		return ret
	}
	return o.AssayResultIds
}

// GetAssayResultIdsOk returns a tuple with the AssayResultIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultIdsResponse) GetAssayResultIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AssayResultIds) {
    return nil, false
	}
	return o.AssayResultIds, true
}

// HasAssayResultIds returns a boolean if a field has been set.
func (o *AssayResultIdsResponse) HasAssayResultIds() bool {
	if o != nil && !isNil(o.AssayResultIds) {
		return true
	}

	return false
}

// SetAssayResultIds gets a reference to the given []string and assigns it to the AssayResultIds field.
func (o *AssayResultIdsResponse) SetAssayResultIds(v []string) {
	o.AssayResultIds = v
}

func (o AssayResultIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssayResultIds) {
		toSerialize["assayResultIds"] = o.AssayResultIds
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResultIdsResponse struct {
	value *AssayResultIdsResponse
	isSet bool
}

func (v NullableAssayResultIdsResponse) Get() *AssayResultIdsResponse {
	return v.value
}

func (v *NullableAssayResultIdsResponse) Set(val *AssayResultIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResultIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResultIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResultIdsResponse(val *AssayResultIdsResponse) *NullableAssayResultIdsResponse {
	return &NullableAssayResultIdsResponse{value: val, isSet: true}
}

func (v NullableAssayResultIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResultIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


