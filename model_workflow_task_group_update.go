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

// WorkflowTaskGroupUpdate struct for WorkflowTaskGroupUpdate
type WorkflowTaskGroupUpdate struct {
	// ID of the folder that contains the workflow task group
	FolderId *string `json:"folderId,omitempty"`
	// The name of the workflow task group
	Name *string `json:"name,omitempty"`
	// IDs of the users watching the workflow task group
	WatcherIds []string `json:"watcherIds,omitempty"`
}

// NewWorkflowTaskGroupUpdate instantiates a new WorkflowTaskGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskGroupUpdate() *WorkflowTaskGroupUpdate {
	this := WorkflowTaskGroupUpdate{}
	return &this
}

// NewWorkflowTaskGroupUpdateWithDefaults instantiates a new WorkflowTaskGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskGroupUpdateWithDefaults() *WorkflowTaskGroupUpdate {
	this := WorkflowTaskGroupUpdate{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *WorkflowTaskGroupUpdate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupUpdate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *WorkflowTaskGroupUpdate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *WorkflowTaskGroupUpdate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskGroupUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskGroupUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskGroupUpdate) SetName(v string) {
	o.Name = &v
}

// GetWatcherIds returns the WatcherIds field value if set, zero value otherwise.
func (o *WorkflowTaskGroupUpdate) GetWatcherIds() []string {
	if o == nil || isNil(o.WatcherIds) {
		var ret []string
		return ret
	}
	return o.WatcherIds
}

// GetWatcherIdsOk returns a tuple with the WatcherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupUpdate) GetWatcherIdsOk() ([]string, bool) {
	if o == nil || isNil(o.WatcherIds) {
    return nil, false
	}
	return o.WatcherIds, true
}

// HasWatcherIds returns a boolean if a field has been set.
func (o *WorkflowTaskGroupUpdate) HasWatcherIds() bool {
	if o != nil && !isNil(o.WatcherIds) {
		return true
	}

	return false
}

// SetWatcherIds gets a reference to the given []string and assigns it to the WatcherIds field.
func (o *WorkflowTaskGroupUpdate) SetWatcherIds(v []string) {
	o.WatcherIds = v
}

func (o WorkflowTaskGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.WatcherIds) {
		toSerialize["watcherIds"] = o.WatcherIds
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskGroupUpdate struct {
	value *WorkflowTaskGroupUpdate
	isSet bool
}

func (v NullableWorkflowTaskGroupUpdate) Get() *WorkflowTaskGroupUpdate {
	return v.value
}

func (v *NullableWorkflowTaskGroupUpdate) Set(val *WorkflowTaskGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskGroupUpdate(val *WorkflowTaskGroupUpdate) *NullableWorkflowTaskGroupUpdate {
	return &NullableWorkflowTaskGroupUpdate{value: val, isSet: true}
}

func (v NullableWorkflowTaskGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


