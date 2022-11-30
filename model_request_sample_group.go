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

// RequestSampleGroup struct for RequestSampleGroup
type RequestSampleGroup struct {
	Id *string `json:"id,omitempty"`
	// The key for each RequestSample should match one of the samplesSchema[n].name property in the request schema json. 
	Samples *map[string]RequestSampleGroupSamplesValue `json:"samples,omitempty"`
}

// NewRequestSampleGroup instantiates a new RequestSampleGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSampleGroup() *RequestSampleGroup {
	this := RequestSampleGroup{}
	return &this
}

// NewRequestSampleGroupWithDefaults instantiates a new RequestSampleGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSampleGroupWithDefaults() *RequestSampleGroup {
	this := RequestSampleGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequestSampleGroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSampleGroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequestSampleGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RequestSampleGroup) SetId(v string) {
	o.Id = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *RequestSampleGroup) GetSamples() map[string]RequestSampleGroupSamplesValue {
	if o == nil || isNil(o.Samples) {
		var ret map[string]RequestSampleGroupSamplesValue
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSampleGroup) GetSamplesOk() (*map[string]RequestSampleGroupSamplesValue, bool) {
	if o == nil || isNil(o.Samples) {
    return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *RequestSampleGroup) HasSamples() bool {
	if o != nil && !isNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given map[string]RequestSampleGroupSamplesValue and assigns it to the Samples field.
func (o *RequestSampleGroup) SetSamples(v map[string]RequestSampleGroupSamplesValue) {
	o.Samples = &v
}

func (o RequestSampleGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	return json.Marshal(toSerialize)
}

type NullableRequestSampleGroup struct {
	value *RequestSampleGroup
	isSet bool
}

func (v NullableRequestSampleGroup) Get() *RequestSampleGroup {
	return v.value
}

func (v *NullableRequestSampleGroup) Set(val *RequestSampleGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSampleGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSampleGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSampleGroup(val *RequestSampleGroup) *NullableRequestSampleGroup {
	return &NullableRequestSampleGroup{value: val, isSet: true}
}

func (v NullableRequestSampleGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSampleGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


