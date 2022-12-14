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

// RequestResponseSamplesInner struct for RequestResponseSamplesInner
type RequestResponseSamplesInner struct {
	Batch NullableRequestResponseSamplesInnerBatch `json:"batch,omitempty"`
	Entity NullableRequestResponseSamplesInnerEntity `json:"entity,omitempty"`
	// The status of the sample, based on the status of the stage run that generated it.
	Status *string `json:"status,omitempty"`
}

// NewRequestResponseSamplesInner instantiates a new RequestResponseSamplesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResponseSamplesInner() *RequestResponseSamplesInner {
	this := RequestResponseSamplesInner{}
	return &this
}

// NewRequestResponseSamplesInnerWithDefaults instantiates a new RequestResponseSamplesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResponseSamplesInnerWithDefaults() *RequestResponseSamplesInner {
	this := RequestResponseSamplesInner{}
	return &this
}

// GetBatch returns the Batch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestResponseSamplesInner) GetBatch() RequestResponseSamplesInnerBatch {
	if o == nil || isNil(o.Batch.Get()) {
		var ret RequestResponseSamplesInnerBatch
		return ret
	}
	return *o.Batch.Get()
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestResponseSamplesInner) GetBatchOk() (*RequestResponseSamplesInnerBatch, bool) {
	if o == nil {
    return nil, false
	}
	return o.Batch.Get(), o.Batch.IsSet()
}

// HasBatch returns a boolean if a field has been set.
func (o *RequestResponseSamplesInner) HasBatch() bool {
	if o != nil && o.Batch.IsSet() {
		return true
	}

	return false
}

// SetBatch gets a reference to the given NullableRequestResponseSamplesInnerBatch and assigns it to the Batch field.
func (o *RequestResponseSamplesInner) SetBatch(v RequestResponseSamplesInnerBatch) {
	o.Batch.Set(&v)
}
// SetBatchNil sets the value for Batch to be an explicit nil
func (o *RequestResponseSamplesInner) SetBatchNil() {
	o.Batch.Set(nil)
}

// UnsetBatch ensures that no value is present for Batch, not even an explicit nil
func (o *RequestResponseSamplesInner) UnsetBatch() {
	o.Batch.Unset()
}

// GetEntity returns the Entity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestResponseSamplesInner) GetEntity() RequestResponseSamplesInnerEntity {
	if o == nil || isNil(o.Entity.Get()) {
		var ret RequestResponseSamplesInnerEntity
		return ret
	}
	return *o.Entity.Get()
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestResponseSamplesInner) GetEntityOk() (*RequestResponseSamplesInnerEntity, bool) {
	if o == nil {
    return nil, false
	}
	return o.Entity.Get(), o.Entity.IsSet()
}

// HasEntity returns a boolean if a field has been set.
func (o *RequestResponseSamplesInner) HasEntity() bool {
	if o != nil && o.Entity.IsSet() {
		return true
	}

	return false
}

// SetEntity gets a reference to the given NullableRequestResponseSamplesInnerEntity and assigns it to the Entity field.
func (o *RequestResponseSamplesInner) SetEntity(v RequestResponseSamplesInnerEntity) {
	o.Entity.Set(&v)
}
// SetEntityNil sets the value for Entity to be an explicit nil
func (o *RequestResponseSamplesInner) SetEntityNil() {
	o.Entity.Set(nil)
}

// UnsetEntity ensures that no value is present for Entity, not even an explicit nil
func (o *RequestResponseSamplesInner) UnsetEntity() {
	o.Entity.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestResponseSamplesInner) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesInner) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestResponseSamplesInner) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RequestResponseSamplesInner) SetStatus(v string) {
	o.Status = &v
}

func (o RequestResponseSamplesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Batch.IsSet() {
		toSerialize["batch"] = o.Batch.Get()
	}
	if o.Entity.IsSet() {
		toSerialize["entity"] = o.Entity.Get()
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRequestResponseSamplesInner struct {
	value *RequestResponseSamplesInner
	isSet bool
}

func (v NullableRequestResponseSamplesInner) Get() *RequestResponseSamplesInner {
	return v.value
}

func (v *NullableRequestResponseSamplesInner) Set(val *RequestResponseSamplesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResponseSamplesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResponseSamplesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResponseSamplesInner(val *RequestResponseSamplesInner) *NullableRequestResponseSamplesInner {
	return &NullableRequestResponseSamplesInner{value: val, isSet: true}
}

func (v NullableRequestResponseSamplesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResponseSamplesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


