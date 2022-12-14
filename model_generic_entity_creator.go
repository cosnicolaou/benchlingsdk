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

// GenericEntityCreator struct for GenericEntityCreator
type GenericEntityCreator struct {
	Handle *string `json:"handle,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGenericEntityCreator instantiates a new GenericEntityCreator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericEntityCreator() *GenericEntityCreator {
	this := GenericEntityCreator{}
	return &this
}

// NewGenericEntityCreatorWithDefaults instantiates a new GenericEntityCreator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericEntityCreatorWithDefaults() *GenericEntityCreator {
	this := GenericEntityCreator{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *GenericEntityCreator) GetHandle() string {
	if o == nil || isNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEntityCreator) GetHandleOk() (*string, bool) {
	if o == nil || isNil(o.Handle) {
    return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *GenericEntityCreator) HasHandle() bool {
	if o != nil && !isNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *GenericEntityCreator) SetHandle(v string) {
	o.Handle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GenericEntityCreator) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEntityCreator) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GenericEntityCreator) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GenericEntityCreator) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GenericEntityCreator) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEntityCreator) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GenericEntityCreator) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GenericEntityCreator) SetName(v string) {
	o.Name = &v
}

func (o GenericEntityCreator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGenericEntityCreator struct {
	value *GenericEntityCreator
	isSet bool
}

func (v NullableGenericEntityCreator) Get() *GenericEntityCreator {
	return v.value
}

func (v *NullableGenericEntityCreator) Set(val *GenericEntityCreator) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericEntityCreator) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericEntityCreator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericEntityCreator(val *GenericEntityCreator) *NullableGenericEntityCreator {
	return &NullableGenericEntityCreator{value: val, isSet: true}
}

func (v NullableGenericEntityCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericEntityCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


