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

// WorkflowTaskSchemaBase struct for WorkflowTaskSchemaBase
type WorkflowTaskSchemaBase struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	FieldDefinitions []SchemaFieldDefinitionsInner `json:"fieldDefinitions,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	// Whether or not tasks of this schema can be created with a non-null assignee.
	CanSetAssigneeOnTaskCreation *bool `json:"canSetAssigneeOnTaskCreation,omitempty"`
	// ID of the default folder for creating workflow task groups
	DefaultCreationFolderId NullableString `json:"defaultCreationFolderId,omitempty"`
	// ID of the default folder for workflow task execution entries
	DefaultEntryExecutionFolderId NullableString `json:"defaultEntryExecutionFolderId,omitempty"`
	DefaultResponsibleTeam NullableTeamSummary `json:"defaultResponsibleTeam,omitempty"`
	// The ID of the template of the entries tasks of this schema will be executed into.
	EntryTemplateId NullableString `json:"entryTemplateId,omitempty"`
	// The prefix for the displayId of tasks of this schema.
	Prefix *string `json:"prefix,omitempty"`
	StatusLifecycle *WorkflowTaskStatusLifecycle `json:"statusLifecycle,omitempty"`
	// The prefix for the displayId of task groups containing tasks of this schema
	TaskGroupPrefix *string `json:"taskGroupPrefix,omitempty"`
	WorkflowOutputSchema NullableWorkflowOutputSchema `json:"workflowOutputSchema,omitempty"`
}

// NewWorkflowTaskSchemaBase instantiates a new WorkflowTaskSchemaBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskSchemaBase() *WorkflowTaskSchemaBase {
	this := WorkflowTaskSchemaBase{}
	return &this
}

// NewWorkflowTaskSchemaBaseWithDefaults instantiates a new WorkflowTaskSchemaBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskSchemaBaseWithDefaults() *WorkflowTaskSchemaBase {
	this := WorkflowTaskSchemaBase{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchemaBase) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchemaBase) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *WorkflowTaskSchemaBase) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *WorkflowTaskSchemaBase) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *WorkflowTaskSchemaBase) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetFieldDefinitions returns the FieldDefinitions field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetFieldDefinitions() []SchemaFieldDefinitionsInner {
	if o == nil || isNil(o.FieldDefinitions) {
		var ret []SchemaFieldDefinitionsInner
		return ret
	}
	return o.FieldDefinitions
}

// GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetFieldDefinitionsOk() ([]SchemaFieldDefinitionsInner, bool) {
	if o == nil || isNil(o.FieldDefinitions) {
    return nil, false
	}
	return o.FieldDefinitions, true
}

// HasFieldDefinitions returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasFieldDefinitions() bool {
	if o != nil && !isNil(o.FieldDefinitions) {
		return true
	}

	return false
}

// SetFieldDefinitions gets a reference to the given []SchemaFieldDefinitionsInner and assigns it to the FieldDefinitions field.
func (o *WorkflowTaskSchemaBase) SetFieldDefinitions(v []SchemaFieldDefinitionsInner) {
	o.FieldDefinitions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowTaskSchemaBase) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskSchemaBase) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowTaskSchemaBase) SetType(v string) {
	o.Type = &v
}

// GetCanSetAssigneeOnTaskCreation returns the CanSetAssigneeOnTaskCreation field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetCanSetAssigneeOnTaskCreation() bool {
	if o == nil || isNil(o.CanSetAssigneeOnTaskCreation) {
		var ret bool
		return ret
	}
	return *o.CanSetAssigneeOnTaskCreation
}

// GetCanSetAssigneeOnTaskCreationOk returns a tuple with the CanSetAssigneeOnTaskCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetCanSetAssigneeOnTaskCreationOk() (*bool, bool) {
	if o == nil || isNil(o.CanSetAssigneeOnTaskCreation) {
    return nil, false
	}
	return o.CanSetAssigneeOnTaskCreation, true
}

// HasCanSetAssigneeOnTaskCreation returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasCanSetAssigneeOnTaskCreation() bool {
	if o != nil && !isNil(o.CanSetAssigneeOnTaskCreation) {
		return true
	}

	return false
}

// SetCanSetAssigneeOnTaskCreation gets a reference to the given bool and assigns it to the CanSetAssigneeOnTaskCreation field.
func (o *WorkflowTaskSchemaBase) SetCanSetAssigneeOnTaskCreation(v bool) {
	o.CanSetAssigneeOnTaskCreation = &v
}

// GetDefaultCreationFolderId returns the DefaultCreationFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchemaBase) GetDefaultCreationFolderId() string {
	if o == nil || isNil(o.DefaultCreationFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultCreationFolderId.Get()
}

// GetDefaultCreationFolderIdOk returns a tuple with the DefaultCreationFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchemaBase) GetDefaultCreationFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultCreationFolderId.Get(), o.DefaultCreationFolderId.IsSet()
}

// HasDefaultCreationFolderId returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasDefaultCreationFolderId() bool {
	if o != nil && o.DefaultCreationFolderId.IsSet() {
		return true
	}

	return false
}

// SetDefaultCreationFolderId gets a reference to the given NullableString and assigns it to the DefaultCreationFolderId field.
func (o *WorkflowTaskSchemaBase) SetDefaultCreationFolderId(v string) {
	o.DefaultCreationFolderId.Set(&v)
}
// SetDefaultCreationFolderIdNil sets the value for DefaultCreationFolderId to be an explicit nil
func (o *WorkflowTaskSchemaBase) SetDefaultCreationFolderIdNil() {
	o.DefaultCreationFolderId.Set(nil)
}

// UnsetDefaultCreationFolderId ensures that no value is present for DefaultCreationFolderId, not even an explicit nil
func (o *WorkflowTaskSchemaBase) UnsetDefaultCreationFolderId() {
	o.DefaultCreationFolderId.Unset()
}

// GetDefaultEntryExecutionFolderId returns the DefaultEntryExecutionFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchemaBase) GetDefaultEntryExecutionFolderId() string {
	if o == nil || isNil(o.DefaultEntryExecutionFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultEntryExecutionFolderId.Get()
}

// GetDefaultEntryExecutionFolderIdOk returns a tuple with the DefaultEntryExecutionFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchemaBase) GetDefaultEntryExecutionFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultEntryExecutionFolderId.Get(), o.DefaultEntryExecutionFolderId.IsSet()
}

// HasDefaultEntryExecutionFolderId returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasDefaultEntryExecutionFolderId() bool {
	if o != nil && o.DefaultEntryExecutionFolderId.IsSet() {
		return true
	}

	return false
}

// SetDefaultEntryExecutionFolderId gets a reference to the given NullableString and assigns it to the DefaultEntryExecutionFolderId field.
func (o *WorkflowTaskSchemaBase) SetDefaultEntryExecutionFolderId(v string) {
	o.DefaultEntryExecutionFolderId.Set(&v)
}
// SetDefaultEntryExecutionFolderIdNil sets the value for DefaultEntryExecutionFolderId to be an explicit nil
func (o *WorkflowTaskSchemaBase) SetDefaultEntryExecutionFolderIdNil() {
	o.DefaultEntryExecutionFolderId.Set(nil)
}

// UnsetDefaultEntryExecutionFolderId ensures that no value is present for DefaultEntryExecutionFolderId, not even an explicit nil
func (o *WorkflowTaskSchemaBase) UnsetDefaultEntryExecutionFolderId() {
	o.DefaultEntryExecutionFolderId.Unset()
}

// GetDefaultResponsibleTeam returns the DefaultResponsibleTeam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchemaBase) GetDefaultResponsibleTeam() TeamSummary {
	if o == nil || isNil(o.DefaultResponsibleTeam.Get()) {
		var ret TeamSummary
		return ret
	}
	return *o.DefaultResponsibleTeam.Get()
}

// GetDefaultResponsibleTeamOk returns a tuple with the DefaultResponsibleTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchemaBase) GetDefaultResponsibleTeamOk() (*TeamSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultResponsibleTeam.Get(), o.DefaultResponsibleTeam.IsSet()
}

// HasDefaultResponsibleTeam returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasDefaultResponsibleTeam() bool {
	if o != nil && o.DefaultResponsibleTeam.IsSet() {
		return true
	}

	return false
}

// SetDefaultResponsibleTeam gets a reference to the given NullableTeamSummary and assigns it to the DefaultResponsibleTeam field.
func (o *WorkflowTaskSchemaBase) SetDefaultResponsibleTeam(v TeamSummary) {
	o.DefaultResponsibleTeam.Set(&v)
}
// SetDefaultResponsibleTeamNil sets the value for DefaultResponsibleTeam to be an explicit nil
func (o *WorkflowTaskSchemaBase) SetDefaultResponsibleTeamNil() {
	o.DefaultResponsibleTeam.Set(nil)
}

// UnsetDefaultResponsibleTeam ensures that no value is present for DefaultResponsibleTeam, not even an explicit nil
func (o *WorkflowTaskSchemaBase) UnsetDefaultResponsibleTeam() {
	o.DefaultResponsibleTeam.Unset()
}

// GetEntryTemplateId returns the EntryTemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchemaBase) GetEntryTemplateId() string {
	if o == nil || isNil(o.EntryTemplateId.Get()) {
		var ret string
		return ret
	}
	return *o.EntryTemplateId.Get()
}

// GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchemaBase) GetEntryTemplateIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntryTemplateId.Get(), o.EntryTemplateId.IsSet()
}

// HasEntryTemplateId returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasEntryTemplateId() bool {
	if o != nil && o.EntryTemplateId.IsSet() {
		return true
	}

	return false
}

// SetEntryTemplateId gets a reference to the given NullableString and assigns it to the EntryTemplateId field.
func (o *WorkflowTaskSchemaBase) SetEntryTemplateId(v string) {
	o.EntryTemplateId.Set(&v)
}
// SetEntryTemplateIdNil sets the value for EntryTemplateId to be an explicit nil
func (o *WorkflowTaskSchemaBase) SetEntryTemplateIdNil() {
	o.EntryTemplateId.Set(nil)
}

// UnsetEntryTemplateId ensures that no value is present for EntryTemplateId, not even an explicit nil
func (o *WorkflowTaskSchemaBase) UnsetEntryTemplateId() {
	o.EntryTemplateId.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *WorkflowTaskSchemaBase) SetPrefix(v string) {
	o.Prefix = &v
}

// GetStatusLifecycle returns the StatusLifecycle field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetStatusLifecycle() WorkflowTaskStatusLifecycle {
	if o == nil || isNil(o.StatusLifecycle) {
		var ret WorkflowTaskStatusLifecycle
		return ret
	}
	return *o.StatusLifecycle
}

// GetStatusLifecycleOk returns a tuple with the StatusLifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetStatusLifecycleOk() (*WorkflowTaskStatusLifecycle, bool) {
	if o == nil || isNil(o.StatusLifecycle) {
    return nil, false
	}
	return o.StatusLifecycle, true
}

// HasStatusLifecycle returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasStatusLifecycle() bool {
	if o != nil && !isNil(o.StatusLifecycle) {
		return true
	}

	return false
}

// SetStatusLifecycle gets a reference to the given WorkflowTaskStatusLifecycle and assigns it to the StatusLifecycle field.
func (o *WorkflowTaskSchemaBase) SetStatusLifecycle(v WorkflowTaskStatusLifecycle) {
	o.StatusLifecycle = &v
}

// GetTaskGroupPrefix returns the TaskGroupPrefix field value if set, zero value otherwise.
func (o *WorkflowTaskSchemaBase) GetTaskGroupPrefix() string {
	if o == nil || isNil(o.TaskGroupPrefix) {
		var ret string
		return ret
	}
	return *o.TaskGroupPrefix
}

// GetTaskGroupPrefixOk returns a tuple with the TaskGroupPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchemaBase) GetTaskGroupPrefixOk() (*string, bool) {
	if o == nil || isNil(o.TaskGroupPrefix) {
    return nil, false
	}
	return o.TaskGroupPrefix, true
}

// HasTaskGroupPrefix returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasTaskGroupPrefix() bool {
	if o != nil && !isNil(o.TaskGroupPrefix) {
		return true
	}

	return false
}

// SetTaskGroupPrefix gets a reference to the given string and assigns it to the TaskGroupPrefix field.
func (o *WorkflowTaskSchemaBase) SetTaskGroupPrefix(v string) {
	o.TaskGroupPrefix = &v
}

// GetWorkflowOutputSchema returns the WorkflowOutputSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchemaBase) GetWorkflowOutputSchema() WorkflowOutputSchema {
	if o == nil || isNil(o.WorkflowOutputSchema.Get()) {
		var ret WorkflowOutputSchema
		return ret
	}
	return *o.WorkflowOutputSchema.Get()
}

// GetWorkflowOutputSchemaOk returns a tuple with the WorkflowOutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchemaBase) GetWorkflowOutputSchemaOk() (*WorkflowOutputSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.WorkflowOutputSchema.Get(), o.WorkflowOutputSchema.IsSet()
}

// HasWorkflowOutputSchema returns a boolean if a field has been set.
func (o *WorkflowTaskSchemaBase) HasWorkflowOutputSchema() bool {
	if o != nil && o.WorkflowOutputSchema.IsSet() {
		return true
	}

	return false
}

// SetWorkflowOutputSchema gets a reference to the given NullableWorkflowOutputSchema and assigns it to the WorkflowOutputSchema field.
func (o *WorkflowTaskSchemaBase) SetWorkflowOutputSchema(v WorkflowOutputSchema) {
	o.WorkflowOutputSchema.Set(&v)
}
// SetWorkflowOutputSchemaNil sets the value for WorkflowOutputSchema to be an explicit nil
func (o *WorkflowTaskSchemaBase) SetWorkflowOutputSchemaNil() {
	o.WorkflowOutputSchema.Set(nil)
}

// UnsetWorkflowOutputSchema ensures that no value is present for WorkflowOutputSchema, not even an explicit nil
func (o *WorkflowTaskSchemaBase) UnsetWorkflowOutputSchema() {
	o.WorkflowOutputSchema.Unset()
}

func (o WorkflowTaskSchemaBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.FieldDefinitions) {
		toSerialize["fieldDefinitions"] = o.FieldDefinitions
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.CanSetAssigneeOnTaskCreation) {
		toSerialize["canSetAssigneeOnTaskCreation"] = o.CanSetAssigneeOnTaskCreation
	}
	if o.DefaultCreationFolderId.IsSet() {
		toSerialize["defaultCreationFolderId"] = o.DefaultCreationFolderId.Get()
	}
	if o.DefaultEntryExecutionFolderId.IsSet() {
		toSerialize["defaultEntryExecutionFolderId"] = o.DefaultEntryExecutionFolderId.Get()
	}
	if o.DefaultResponsibleTeam.IsSet() {
		toSerialize["defaultResponsibleTeam"] = o.DefaultResponsibleTeam.Get()
	}
	if o.EntryTemplateId.IsSet() {
		toSerialize["entryTemplateId"] = o.EntryTemplateId.Get()
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.StatusLifecycle) {
		toSerialize["statusLifecycle"] = o.StatusLifecycle
	}
	if !isNil(o.TaskGroupPrefix) {
		toSerialize["taskGroupPrefix"] = o.TaskGroupPrefix
	}
	if o.WorkflowOutputSchema.IsSet() {
		toSerialize["workflowOutputSchema"] = o.WorkflowOutputSchema.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskSchemaBase struct {
	value *WorkflowTaskSchemaBase
	isSet bool
}

func (v NullableWorkflowTaskSchemaBase) Get() *WorkflowTaskSchemaBase {
	return v.value
}

func (v *NullableWorkflowTaskSchemaBase) Set(val *WorkflowTaskSchemaBase) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskSchemaBase) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskSchemaBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskSchemaBase(val *WorkflowTaskSchemaBase) *NullableWorkflowTaskSchemaBase {
	return &NullableWorkflowTaskSchemaBase{value: val, isSet: true}
}

func (v NullableWorkflowTaskSchemaBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskSchemaBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

