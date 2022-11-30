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

// WorkflowTaskGroupCreate struct for WorkflowTaskGroupCreate
type WorkflowTaskGroupCreate struct {
	// ID of the folder that contains the workflow task group
	FolderId string `json:"folderId"`
	// The name of the workflow task group
	Name *string `json:"name,omitempty"`
	// IDs of the users watching the workflow task group
	WatcherIds []string `json:"watcherIds,omitempty"`
	// The workflow task schema of tasks in this task group
	SchemaId string `json:"schemaId"`
}

// NewWorkflowTaskGroupCreate instantiates a new WorkflowTaskGroupCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskGroupCreate(folderId string, schemaId string) *WorkflowTaskGroupCreate {
	this := WorkflowTaskGroupCreate{}
	this.FolderId = folderId
	this.SchemaId = schemaId
	return &this
}

// NewWorkflowTaskGroupCreateWithDefaults instantiates a new WorkflowTaskGroupCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskGroupCreateWithDefaults() *WorkflowTaskGroupCreate {
	this := WorkflowTaskGroupCreate{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *WorkflowTaskGroupCreate) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupCreate) GetFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *WorkflowTaskGroupCreate) SetFolderId(v string) {
	o.FolderId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskGroupCreate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupCreate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskGroupCreate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskGroupCreate) SetName(v string) {
	o.Name = &v
}

// GetWatcherIds returns the WatcherIds field value if set, zero value otherwise.
func (o *WorkflowTaskGroupCreate) GetWatcherIds() []string {
	if o == nil || isNil(o.WatcherIds) {
		var ret []string
		return ret
	}
	return o.WatcherIds
}

// GetWatcherIdsOk returns a tuple with the WatcherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupCreate) GetWatcherIdsOk() ([]string, bool) {
	if o == nil || isNil(o.WatcherIds) {
    return nil, false
	}
	return o.WatcherIds, true
}

// HasWatcherIds returns a boolean if a field has been set.
func (o *WorkflowTaskGroupCreate) HasWatcherIds() bool {
	if o != nil && !isNil(o.WatcherIds) {
		return true
	}

	return false
}

// SetWatcherIds gets a reference to the given []string and assigns it to the WatcherIds field.
func (o *WorkflowTaskGroupCreate) SetWatcherIds(v []string) {
	o.WatcherIds = v
}

// GetSchemaId returns the SchemaId field value
func (o *WorkflowTaskGroupCreate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskGroupCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *WorkflowTaskGroupCreate) SetSchemaId(v string) {
	o.SchemaId = v
}

func (o WorkflowTaskGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["folderId"] = o.FolderId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.WatcherIds) {
		toSerialize["watcherIds"] = o.WatcherIds
	}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskGroupCreate struct {
	value *WorkflowTaskGroupCreate
	isSet bool
}

func (v NullableWorkflowTaskGroupCreate) Get() *WorkflowTaskGroupCreate {
	return v.value
}

func (v *NullableWorkflowTaskGroupCreate) Set(val *WorkflowTaskGroupCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskGroupCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskGroupCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskGroupCreate(val *WorkflowTaskGroupCreate) *NullableWorkflowTaskGroupCreate {
	return &NullableWorkflowTaskGroupCreate{value: val, isSet: true}
}

func (v NullableWorkflowTaskGroupCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskGroupCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


