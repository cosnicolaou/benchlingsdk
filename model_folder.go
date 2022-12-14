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

// Folder struct for Folder
type Folder struct {
	ArchiveRecord NullableContainerArchiveRecord `json:"archiveRecord,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// ID of the parent folder, if it exists
	ParentFolderId NullableString `json:"parentFolderId,omitempty"`
	// ID of the containing project
	ProjectId *string `json:"projectId,omitempty"`
}

// NewFolder instantiates a new Folder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolder() *Folder {
	this := Folder{}
	return &this
}

// NewFolderWithDefaults instantiates a new Folder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderWithDefaults() *Folder {
	this := Folder{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Folder) GetArchiveRecord() ContainerArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret ContainerArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Folder) GetArchiveRecordOk() (*ContainerArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *Folder) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableContainerArchiveRecord and assigns it to the ArchiveRecord field.
func (o *Folder) SetArchiveRecord(v ContainerArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *Folder) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *Folder) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Folder) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Folder) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Folder) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Folder) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Folder) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Folder) SetName(v string) {
	o.Name = &v
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Folder) GetParentFolderId() string {
	if o == nil || isNil(o.ParentFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentFolderId.Get()
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Folder) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentFolderId.Get(), o.ParentFolderId.IsSet()
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *Folder) HasParentFolderId() bool {
	if o != nil && o.ParentFolderId.IsSet() {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given NullableString and assigns it to the ParentFolderId field.
func (o *Folder) SetParentFolderId(v string) {
	o.ParentFolderId.Set(&v)
}
// SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil
func (o *Folder) SetParentFolderIdNil() {
	o.ParentFolderId.Set(nil)
}

// UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
func (o *Folder) UnsetParentFolderId() {
	o.ParentFolderId.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Folder) GetProjectId() string {
	if o == nil || isNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folder) GetProjectIdOk() (*string, bool) {
	if o == nil || isNil(o.ProjectId) {
    return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Folder) HasProjectId() bool {
	if o != nil && !isNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Folder) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o Folder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ParentFolderId.IsSet() {
		toSerialize["parentFolderId"] = o.ParentFolderId.Get()
	}
	if !isNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableFolder struct {
	value *Folder
	isSet bool
}

func (v NullableFolder) Get() *Folder {
	return v.value
}

func (v *NullableFolder) Set(val *Folder) {
	v.value = val
	v.isSet = true
}

func (v NullableFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolder(val *Folder) *NullableFolder {
	return &NullableFolder{value: val, isSet: true}
}

func (v NullableFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


