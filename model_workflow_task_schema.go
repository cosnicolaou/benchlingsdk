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

// WorkflowTaskSchema struct for WorkflowTaskSchema
type WorkflowTaskSchema struct {
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
	// The method by which instances of this schema are executed
	ExecutionType *string `json:"executionType,omitempty"`
}

// NewWorkflowTaskSchema instantiates a new WorkflowTaskSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskSchema() *WorkflowTaskSchema {
	this := WorkflowTaskSchema{}
	return &this
}

// NewWorkflowTaskSchemaWithDefaults instantiates a new WorkflowTaskSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskSchemaWithDefaults() *WorkflowTaskSchema {
	this := WorkflowTaskSchema{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchema) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *WorkflowTaskSchema) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *WorkflowTaskSchema) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *WorkflowTaskSchema) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetFieldDefinitions returns the FieldDefinitions field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner {
	if o == nil || isNil(o.FieldDefinitions) {
		var ret []SchemaFieldDefinitionsInner
		return ret
	}
	return o.FieldDefinitions
}

// GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetFieldDefinitionsOk() ([]SchemaFieldDefinitionsInner, bool) {
	if o == nil || isNil(o.FieldDefinitions) {
    return nil, false
	}
	return o.FieldDefinitions, true
}

// HasFieldDefinitions returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasFieldDefinitions() bool {
	if o != nil && !isNil(o.FieldDefinitions) {
		return true
	}

	return false
}

// SetFieldDefinitions gets a reference to the given []SchemaFieldDefinitionsInner and assigns it to the FieldDefinitions field.
func (o *WorkflowTaskSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner) {
	o.FieldDefinitions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowTaskSchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskSchema) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowTaskSchema) SetType(v string) {
	o.Type = &v
}

// GetCanSetAssigneeOnTaskCreation returns the CanSetAssigneeOnTaskCreation field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetCanSetAssigneeOnTaskCreation() bool {
	if o == nil || isNil(o.CanSetAssigneeOnTaskCreation) {
		var ret bool
		return ret
	}
	return *o.CanSetAssigneeOnTaskCreation
}

// GetCanSetAssigneeOnTaskCreationOk returns a tuple with the CanSetAssigneeOnTaskCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetCanSetAssigneeOnTaskCreationOk() (*bool, bool) {
	if o == nil || isNil(o.CanSetAssigneeOnTaskCreation) {
    return nil, false
	}
	return o.CanSetAssigneeOnTaskCreation, true
}

// HasCanSetAssigneeOnTaskCreation returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasCanSetAssigneeOnTaskCreation() bool {
	if o != nil && !isNil(o.CanSetAssigneeOnTaskCreation) {
		return true
	}

	return false
}

// SetCanSetAssigneeOnTaskCreation gets a reference to the given bool and assigns it to the CanSetAssigneeOnTaskCreation field.
func (o *WorkflowTaskSchema) SetCanSetAssigneeOnTaskCreation(v bool) {
	o.CanSetAssigneeOnTaskCreation = &v
}

// GetDefaultCreationFolderId returns the DefaultCreationFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchema) GetDefaultCreationFolderId() string {
	if o == nil || isNil(o.DefaultCreationFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultCreationFolderId.Get()
}

// GetDefaultCreationFolderIdOk returns a tuple with the DefaultCreationFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchema) GetDefaultCreationFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultCreationFolderId.Get(), o.DefaultCreationFolderId.IsSet()
}

// HasDefaultCreationFolderId returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasDefaultCreationFolderId() bool {
	if o != nil && o.DefaultCreationFolderId.IsSet() {
		return true
	}

	return false
}

// SetDefaultCreationFolderId gets a reference to the given NullableString and assigns it to the DefaultCreationFolderId field.
func (o *WorkflowTaskSchema) SetDefaultCreationFolderId(v string) {
	o.DefaultCreationFolderId.Set(&v)
}
// SetDefaultCreationFolderIdNil sets the value for DefaultCreationFolderId to be an explicit nil
func (o *WorkflowTaskSchema) SetDefaultCreationFolderIdNil() {
	o.DefaultCreationFolderId.Set(nil)
}

// UnsetDefaultCreationFolderId ensures that no value is present for DefaultCreationFolderId, not even an explicit nil
func (o *WorkflowTaskSchema) UnsetDefaultCreationFolderId() {
	o.DefaultCreationFolderId.Unset()
}

// GetDefaultEntryExecutionFolderId returns the DefaultEntryExecutionFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchema) GetDefaultEntryExecutionFolderId() string {
	if o == nil || isNil(o.DefaultEntryExecutionFolderId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultEntryExecutionFolderId.Get()
}

// GetDefaultEntryExecutionFolderIdOk returns a tuple with the DefaultEntryExecutionFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchema) GetDefaultEntryExecutionFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultEntryExecutionFolderId.Get(), o.DefaultEntryExecutionFolderId.IsSet()
}

// HasDefaultEntryExecutionFolderId returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasDefaultEntryExecutionFolderId() bool {
	if o != nil && o.DefaultEntryExecutionFolderId.IsSet() {
		return true
	}

	return false
}

// SetDefaultEntryExecutionFolderId gets a reference to the given NullableString and assigns it to the DefaultEntryExecutionFolderId field.
func (o *WorkflowTaskSchema) SetDefaultEntryExecutionFolderId(v string) {
	o.DefaultEntryExecutionFolderId.Set(&v)
}
// SetDefaultEntryExecutionFolderIdNil sets the value for DefaultEntryExecutionFolderId to be an explicit nil
func (o *WorkflowTaskSchema) SetDefaultEntryExecutionFolderIdNil() {
	o.DefaultEntryExecutionFolderId.Set(nil)
}

// UnsetDefaultEntryExecutionFolderId ensures that no value is present for DefaultEntryExecutionFolderId, not even an explicit nil
func (o *WorkflowTaskSchema) UnsetDefaultEntryExecutionFolderId() {
	o.DefaultEntryExecutionFolderId.Unset()
}

// GetDefaultResponsibleTeam returns the DefaultResponsibleTeam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchema) GetDefaultResponsibleTeam() TeamSummary {
	if o == nil || isNil(o.DefaultResponsibleTeam.Get()) {
		var ret TeamSummary
		return ret
	}
	return *o.DefaultResponsibleTeam.Get()
}

// GetDefaultResponsibleTeamOk returns a tuple with the DefaultResponsibleTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchema) GetDefaultResponsibleTeamOk() (*TeamSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.DefaultResponsibleTeam.Get(), o.DefaultResponsibleTeam.IsSet()
}

// HasDefaultResponsibleTeam returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasDefaultResponsibleTeam() bool {
	if o != nil && o.DefaultResponsibleTeam.IsSet() {
		return true
	}

	return false
}

// SetDefaultResponsibleTeam gets a reference to the given NullableTeamSummary and assigns it to the DefaultResponsibleTeam field.
func (o *WorkflowTaskSchema) SetDefaultResponsibleTeam(v TeamSummary) {
	o.DefaultResponsibleTeam.Set(&v)
}
// SetDefaultResponsibleTeamNil sets the value for DefaultResponsibleTeam to be an explicit nil
func (o *WorkflowTaskSchema) SetDefaultResponsibleTeamNil() {
	o.DefaultResponsibleTeam.Set(nil)
}

// UnsetDefaultResponsibleTeam ensures that no value is present for DefaultResponsibleTeam, not even an explicit nil
func (o *WorkflowTaskSchema) UnsetDefaultResponsibleTeam() {
	o.DefaultResponsibleTeam.Unset()
}

// GetEntryTemplateId returns the EntryTemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchema) GetEntryTemplateId() string {
	if o == nil || isNil(o.EntryTemplateId.Get()) {
		var ret string
		return ret
	}
	return *o.EntryTemplateId.Get()
}

// GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchema) GetEntryTemplateIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntryTemplateId.Get(), o.EntryTemplateId.IsSet()
}

// HasEntryTemplateId returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasEntryTemplateId() bool {
	if o != nil && o.EntryTemplateId.IsSet() {
		return true
	}

	return false
}

// SetEntryTemplateId gets a reference to the given NullableString and assigns it to the EntryTemplateId field.
func (o *WorkflowTaskSchema) SetEntryTemplateId(v string) {
	o.EntryTemplateId.Set(&v)
}
// SetEntryTemplateIdNil sets the value for EntryTemplateId to be an explicit nil
func (o *WorkflowTaskSchema) SetEntryTemplateIdNil() {
	o.EntryTemplateId.Set(nil)
}

// UnsetEntryTemplateId ensures that no value is present for EntryTemplateId, not even an explicit nil
func (o *WorkflowTaskSchema) UnsetEntryTemplateId() {
	o.EntryTemplateId.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *WorkflowTaskSchema) SetPrefix(v string) {
	o.Prefix = &v
}

// GetStatusLifecycle returns the StatusLifecycle field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetStatusLifecycle() WorkflowTaskStatusLifecycle {
	if o == nil || isNil(o.StatusLifecycle) {
		var ret WorkflowTaskStatusLifecycle
		return ret
	}
	return *o.StatusLifecycle
}

// GetStatusLifecycleOk returns a tuple with the StatusLifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetStatusLifecycleOk() (*WorkflowTaskStatusLifecycle, bool) {
	if o == nil || isNil(o.StatusLifecycle) {
    return nil, false
	}
	return o.StatusLifecycle, true
}

// HasStatusLifecycle returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasStatusLifecycle() bool {
	if o != nil && !isNil(o.StatusLifecycle) {
		return true
	}

	return false
}

// SetStatusLifecycle gets a reference to the given WorkflowTaskStatusLifecycle and assigns it to the StatusLifecycle field.
func (o *WorkflowTaskSchema) SetStatusLifecycle(v WorkflowTaskStatusLifecycle) {
	o.StatusLifecycle = &v
}

// GetTaskGroupPrefix returns the TaskGroupPrefix field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetTaskGroupPrefix() string {
	if o == nil || isNil(o.TaskGroupPrefix) {
		var ret string
		return ret
	}
	return *o.TaskGroupPrefix
}

// GetTaskGroupPrefixOk returns a tuple with the TaskGroupPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetTaskGroupPrefixOk() (*string, bool) {
	if o == nil || isNil(o.TaskGroupPrefix) {
    return nil, false
	}
	return o.TaskGroupPrefix, true
}

// HasTaskGroupPrefix returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasTaskGroupPrefix() bool {
	if o != nil && !isNil(o.TaskGroupPrefix) {
		return true
	}

	return false
}

// SetTaskGroupPrefix gets a reference to the given string and assigns it to the TaskGroupPrefix field.
func (o *WorkflowTaskSchema) SetTaskGroupPrefix(v string) {
	o.TaskGroupPrefix = &v
}

// GetWorkflowOutputSchema returns the WorkflowOutputSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskSchema) GetWorkflowOutputSchema() WorkflowOutputSchema {
	if o == nil || isNil(o.WorkflowOutputSchema.Get()) {
		var ret WorkflowOutputSchema
		return ret
	}
	return *o.WorkflowOutputSchema.Get()
}

// GetWorkflowOutputSchemaOk returns a tuple with the WorkflowOutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskSchema) GetWorkflowOutputSchemaOk() (*WorkflowOutputSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.WorkflowOutputSchema.Get(), o.WorkflowOutputSchema.IsSet()
}

// HasWorkflowOutputSchema returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasWorkflowOutputSchema() bool {
	if o != nil && o.WorkflowOutputSchema.IsSet() {
		return true
	}

	return false
}

// SetWorkflowOutputSchema gets a reference to the given NullableWorkflowOutputSchema and assigns it to the WorkflowOutputSchema field.
func (o *WorkflowTaskSchema) SetWorkflowOutputSchema(v WorkflowOutputSchema) {
	o.WorkflowOutputSchema.Set(&v)
}
// SetWorkflowOutputSchemaNil sets the value for WorkflowOutputSchema to be an explicit nil
func (o *WorkflowTaskSchema) SetWorkflowOutputSchemaNil() {
	o.WorkflowOutputSchema.Set(nil)
}

// UnsetWorkflowOutputSchema ensures that no value is present for WorkflowOutputSchema, not even an explicit nil
func (o *WorkflowTaskSchema) UnsetWorkflowOutputSchema() {
	o.WorkflowOutputSchema.Unset()
}

// GetExecutionType returns the ExecutionType field value if set, zero value otherwise.
func (o *WorkflowTaskSchema) GetExecutionType() string {
	if o == nil || isNil(o.ExecutionType) {
		var ret string
		return ret
	}
	return *o.ExecutionType
}

// GetExecutionTypeOk returns a tuple with the ExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSchema) GetExecutionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ExecutionType) {
    return nil, false
	}
	return o.ExecutionType, true
}

// HasExecutionType returns a boolean if a field has been set.
func (o *WorkflowTaskSchema) HasExecutionType() bool {
	if o != nil && !isNil(o.ExecutionType) {
		return true
	}

	return false
}

// SetExecutionType gets a reference to the given string and assigns it to the ExecutionType field.
func (o *WorkflowTaskSchema) SetExecutionType(v string) {
	o.ExecutionType = &v
}

func (o WorkflowTaskSchema) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ExecutionType) {
		toSerialize["executionType"] = o.ExecutionType
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskSchema struct {
	value *WorkflowTaskSchema
	isSet bool
}

func (v NullableWorkflowTaskSchema) Get() *WorkflowTaskSchema {
	return v.value
}

func (v *NullableWorkflowTaskSchema) Set(val *WorkflowTaskSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskSchema(val *WorkflowTaskSchema) *NullableWorkflowTaskSchema {
	return &NullableWorkflowTaskSchema{value: val, isSet: true}
}

func (v NullableWorkflowTaskSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

