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

// WarehouseCredentialsCreate struct for WarehouseCredentialsCreate
type WarehouseCredentialsCreate struct {
	// Duration, in seconds, that credentials should be active for. Must be greater than 0 and less than 3600. 
	ExpiresIn int32 `json:"expiresIn"`
}

// NewWarehouseCredentialsCreate instantiates a new WarehouseCredentialsCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseCredentialsCreate(expiresIn int32) *WarehouseCredentialsCreate {
	this := WarehouseCredentialsCreate{}
	this.ExpiresIn = expiresIn
	return &this
}

// NewWarehouseCredentialsCreateWithDefaults instantiates a new WarehouseCredentialsCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseCredentialsCreateWithDefaults() *WarehouseCredentialsCreate {
	this := WarehouseCredentialsCreate{}
	return &this
}

// GetExpiresIn returns the ExpiresIn field value
func (o *WarehouseCredentialsCreate) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *WarehouseCredentialsCreate) GetExpiresInOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *WarehouseCredentialsCreate) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

func (o WarehouseCredentialsCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	return json.Marshal(toSerialize)
}

type NullableWarehouseCredentialsCreate struct {
	value *WarehouseCredentialsCreate
	isSet bool
}

func (v NullableWarehouseCredentialsCreate) Get() *WarehouseCredentialsCreate {
	return v.value
}

func (v *NullableWarehouseCredentialsCreate) Set(val *WarehouseCredentialsCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseCredentialsCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseCredentialsCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseCredentialsCreate(val *WarehouseCredentialsCreate) *NullableWarehouseCredentialsCreate {
	return &NullableWarehouseCredentialsCreate{value: val, isSet: true}
}

func (v NullableWarehouseCredentialsCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseCredentialsCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


