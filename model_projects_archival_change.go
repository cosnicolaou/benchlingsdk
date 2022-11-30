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

// ProjectsArchivalChange IDs of all items that were archived or unarchived, grouped by resource type. This includes the IDs of projects along with any IDs of project contents that were unarchived. 
type ProjectsArchivalChange struct {
	AaSequenceIds []string `json:"aaSequenceIds,omitempty"`
	CustomEntityIds []string `json:"customEntityIds,omitempty"`
	DnaSequenceIds []string `json:"dnaSequenceIds,omitempty"`
	EntryIds []string `json:"entryIds,omitempty"`
	FolderIds []string `json:"folderIds,omitempty"`
	MixtureIds []string `json:"mixtureIds,omitempty"`
	OligoIds []string `json:"oligoIds,omitempty"`
	ProjectIds []string `json:"projectIds,omitempty"`
	ProtocolIds []string `json:"protocolIds,omitempty"`
}

// NewProjectsArchivalChange instantiates a new ProjectsArchivalChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsArchivalChange() *ProjectsArchivalChange {
	this := ProjectsArchivalChange{}
	return &this
}

// NewProjectsArchivalChangeWithDefaults instantiates a new ProjectsArchivalChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsArchivalChangeWithDefaults() *ProjectsArchivalChange {
	this := ProjectsArchivalChange{}
	return &this
}

// GetAaSequenceIds returns the AaSequenceIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetAaSequenceIds() []string {
	if o == nil || isNil(o.AaSequenceIds) {
		var ret []string
		return ret
	}
	return o.AaSequenceIds
}

// GetAaSequenceIdsOk returns a tuple with the AaSequenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetAaSequenceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AaSequenceIds) {
    return nil, false
	}
	return o.AaSequenceIds, true
}

// HasAaSequenceIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasAaSequenceIds() bool {
	if o != nil && !isNil(o.AaSequenceIds) {
		return true
	}

	return false
}

// SetAaSequenceIds gets a reference to the given []string and assigns it to the AaSequenceIds field.
func (o *ProjectsArchivalChange) SetAaSequenceIds(v []string) {
	o.AaSequenceIds = v
}

// GetCustomEntityIds returns the CustomEntityIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetCustomEntityIds() []string {
	if o == nil || isNil(o.CustomEntityIds) {
		var ret []string
		return ret
	}
	return o.CustomEntityIds
}

// GetCustomEntityIdsOk returns a tuple with the CustomEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetCustomEntityIdsOk() ([]string, bool) {
	if o == nil || isNil(o.CustomEntityIds) {
    return nil, false
	}
	return o.CustomEntityIds, true
}

// HasCustomEntityIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasCustomEntityIds() bool {
	if o != nil && !isNil(o.CustomEntityIds) {
		return true
	}

	return false
}

// SetCustomEntityIds gets a reference to the given []string and assigns it to the CustomEntityIds field.
func (o *ProjectsArchivalChange) SetCustomEntityIds(v []string) {
	o.CustomEntityIds = v
}

// GetDnaSequenceIds returns the DnaSequenceIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetDnaSequenceIds() []string {
	if o == nil || isNil(o.DnaSequenceIds) {
		var ret []string
		return ret
	}
	return o.DnaSequenceIds
}

// GetDnaSequenceIdsOk returns a tuple with the DnaSequenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetDnaSequenceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.DnaSequenceIds) {
    return nil, false
	}
	return o.DnaSequenceIds, true
}

// HasDnaSequenceIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasDnaSequenceIds() bool {
	if o != nil && !isNil(o.DnaSequenceIds) {
		return true
	}

	return false
}

// SetDnaSequenceIds gets a reference to the given []string and assigns it to the DnaSequenceIds field.
func (o *ProjectsArchivalChange) SetDnaSequenceIds(v []string) {
	o.DnaSequenceIds = v
}

// GetEntryIds returns the EntryIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetEntryIds() []string {
	if o == nil || isNil(o.EntryIds) {
		var ret []string
		return ret
	}
	return o.EntryIds
}

// GetEntryIdsOk returns a tuple with the EntryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetEntryIdsOk() ([]string, bool) {
	if o == nil || isNil(o.EntryIds) {
    return nil, false
	}
	return o.EntryIds, true
}

// HasEntryIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasEntryIds() bool {
	if o != nil && !isNil(o.EntryIds) {
		return true
	}

	return false
}

// SetEntryIds gets a reference to the given []string and assigns it to the EntryIds field.
func (o *ProjectsArchivalChange) SetEntryIds(v []string) {
	o.EntryIds = v
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetFolderIds() []string {
	if o == nil || isNil(o.FolderIds) {
		var ret []string
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetFolderIdsOk() ([]string, bool) {
	if o == nil || isNil(o.FolderIds) {
    return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasFolderIds() bool {
	if o != nil && !isNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []string and assigns it to the FolderIds field.
func (o *ProjectsArchivalChange) SetFolderIds(v []string) {
	o.FolderIds = v
}

// GetMixtureIds returns the MixtureIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetMixtureIds() []string {
	if o == nil || isNil(o.MixtureIds) {
		var ret []string
		return ret
	}
	return o.MixtureIds
}

// GetMixtureIdsOk returns a tuple with the MixtureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetMixtureIdsOk() ([]string, bool) {
	if o == nil || isNil(o.MixtureIds) {
    return nil, false
	}
	return o.MixtureIds, true
}

// HasMixtureIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasMixtureIds() bool {
	if o != nil && !isNil(o.MixtureIds) {
		return true
	}

	return false
}

// SetMixtureIds gets a reference to the given []string and assigns it to the MixtureIds field.
func (o *ProjectsArchivalChange) SetMixtureIds(v []string) {
	o.MixtureIds = v
}

// GetOligoIds returns the OligoIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetOligoIds() []string {
	if o == nil || isNil(o.OligoIds) {
		var ret []string
		return ret
	}
	return o.OligoIds
}

// GetOligoIdsOk returns a tuple with the OligoIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetOligoIdsOk() ([]string, bool) {
	if o == nil || isNil(o.OligoIds) {
    return nil, false
	}
	return o.OligoIds, true
}

// HasOligoIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasOligoIds() bool {
	if o != nil && !isNil(o.OligoIds) {
		return true
	}

	return false
}

// SetOligoIds gets a reference to the given []string and assigns it to the OligoIds field.
func (o *ProjectsArchivalChange) SetOligoIds(v []string) {
	o.OligoIds = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetProjectIds() []string {
	if o == nil || isNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetProjectIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ProjectIds) {
    return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasProjectIds() bool {
	if o != nil && !isNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ProjectsArchivalChange) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetProtocolIds returns the ProtocolIds field value if set, zero value otherwise.
func (o *ProjectsArchivalChange) GetProtocolIds() []string {
	if o == nil || isNil(o.ProtocolIds) {
		var ret []string
		return ret
	}
	return o.ProtocolIds
}

// GetProtocolIdsOk returns a tuple with the ProtocolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsArchivalChange) GetProtocolIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ProtocolIds) {
    return nil, false
	}
	return o.ProtocolIds, true
}

// HasProtocolIds returns a boolean if a field has been set.
func (o *ProjectsArchivalChange) HasProtocolIds() bool {
	if o != nil && !isNil(o.ProtocolIds) {
		return true
	}

	return false
}

// SetProtocolIds gets a reference to the given []string and assigns it to the ProtocolIds field.
func (o *ProjectsArchivalChange) SetProtocolIds(v []string) {
	o.ProtocolIds = v
}

func (o ProjectsArchivalChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AaSequenceIds) {
		toSerialize["aaSequenceIds"] = o.AaSequenceIds
	}
	if !isNil(o.CustomEntityIds) {
		toSerialize["customEntityIds"] = o.CustomEntityIds
	}
	if !isNil(o.DnaSequenceIds) {
		toSerialize["dnaSequenceIds"] = o.DnaSequenceIds
	}
	if !isNil(o.EntryIds) {
		toSerialize["entryIds"] = o.EntryIds
	}
	if !isNil(o.FolderIds) {
		toSerialize["folderIds"] = o.FolderIds
	}
	if !isNil(o.MixtureIds) {
		toSerialize["mixtureIds"] = o.MixtureIds
	}
	if !isNil(o.OligoIds) {
		toSerialize["oligoIds"] = o.OligoIds
	}
	if !isNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !isNil(o.ProtocolIds) {
		toSerialize["protocolIds"] = o.ProtocolIds
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsArchivalChange struct {
	value *ProjectsArchivalChange
	isSet bool
}

func (v NullableProjectsArchivalChange) Get() *ProjectsArchivalChange {
	return v.value
}

func (v *NullableProjectsArchivalChange) Set(val *ProjectsArchivalChange) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsArchivalChange) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsArchivalChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsArchivalChange(val *ProjectsArchivalChange) *NullableProjectsArchivalChange {
	return &NullableProjectsArchivalChange{value: val, isSet: true}
}

func (v NullableProjectsArchivalChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsArchivalChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


