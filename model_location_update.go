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

// LocationUpdate struct for LocationUpdate
type LocationUpdate struct {
	Fields *map[string]Field `json:"fields,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentStorageId *string `json:"parentStorageId,omitempty"`
}

// NewLocationUpdate instantiates a new LocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationUpdate() *LocationUpdate {
	this := LocationUpdate{}
	return &this
}

// NewLocationUpdateWithDefaults instantiates a new LocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationUpdateWithDefaults() *LocationUpdate {
	this := LocationUpdate{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *LocationUpdate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *LocationUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *LocationUpdate) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LocationUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LocationUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LocationUpdate) SetName(v string) {
	o.Name = &v
}

// GetParentStorageId returns the ParentStorageId field value if set, zero value otherwise.
func (o *LocationUpdate) GetParentStorageId() string {
	if o == nil || isNil(o.ParentStorageId) {
		var ret string
		return ret
	}
	return *o.ParentStorageId
}

// GetParentStorageIdOk returns a tuple with the ParentStorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationUpdate) GetParentStorageIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentStorageId) {
    return nil, false
	}
	return o.ParentStorageId, true
}

// HasParentStorageId returns a boolean if a field has been set.
func (o *LocationUpdate) HasParentStorageId() bool {
	if o != nil && !isNil(o.ParentStorageId) {
		return true
	}

	return false
}

// SetParentStorageId gets a reference to the given string and assigns it to the ParentStorageId field.
func (o *LocationUpdate) SetParentStorageId(v string) {
	o.ParentStorageId = &v
}

func (o LocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ParentStorageId) {
		toSerialize["parentStorageId"] = o.ParentStorageId
	}
	return json.Marshal(toSerialize)
}

type NullableLocationUpdate struct {
	value *LocationUpdate
	isSet bool
}

func (v NullableLocationUpdate) Get() *LocationUpdate {
	return v.value
}

func (v *NullableLocationUpdate) Set(val *LocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationUpdate(val *LocationUpdate) *NullableLocationUpdate {
	return &NullableLocationUpdate{value: val, isSet: true}
}

func (v NullableLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


