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

// FolderCreate struct for FolderCreate
type FolderCreate struct {
	// The name of the new folder.
	Name string `json:"name"`
	// The ID of the parent folder.
	ParentFolderId string `json:"parentFolderId"`
}

// NewFolderCreate instantiates a new FolderCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderCreate(name string, parentFolderId string) *FolderCreate {
	this := FolderCreate{}
	this.Name = name
	this.ParentFolderId = parentFolderId
	return &this
}

// NewFolderCreateWithDefaults instantiates a new FolderCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderCreateWithDefaults() *FolderCreate {
	this := FolderCreate{}
	return &this
}

// GetName returns the Name field value
func (o *FolderCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FolderCreate) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FolderCreate) SetName(v string) {
	o.Name = v
}

// GetParentFolderId returns the ParentFolderId field value
func (o *FolderCreate) GetParentFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value
// and a boolean to check if the value has been set.
func (o *FolderCreate) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ParentFolderId, true
}

// SetParentFolderId sets field value
func (o *FolderCreate) SetParentFolderId(v string) {
	o.ParentFolderId = v
}

func (o FolderCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["parentFolderId"] = o.ParentFolderId
	}
	return json.Marshal(toSerialize)
}

type NullableFolderCreate struct {
	value *FolderCreate
	isSet bool
}

func (v NullableFolderCreate) Get() *FolderCreate {
	return v.value
}

func (v *NullableFolderCreate) Set(val *FolderCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderCreate(val *FolderCreate) *NullableFolderCreate {
	return &NullableFolderCreate{value: val, isSet: true}
}

func (v NullableFolderCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

