# WorkflowTaskSchemaBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**FieldDefinitions** | Pointer to [**[]SchemaFieldDefinitionsInner**](SchemaFieldDefinitionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CanSetAssigneeOnTaskCreation** | Pointer to **bool** | Whether or not tasks of this schema can be created with a non-null assignee. | [optional] 
**DefaultCreationFolderId** | Pointer to **NullableString** | ID of the default folder for creating workflow task groups | [optional] 
**DefaultEntryExecutionFolderId** | Pointer to **NullableString** | ID of the default folder for workflow task execution entries | [optional] 
**DefaultResponsibleTeam** | Pointer to [**NullableTeamSummary**](TeamSummary.md) |  | [optional] 
**EntryTemplateId** | Pointer to **NullableString** | The ID of the template of the entries tasks of this schema will be executed into. | [optional] 
**Prefix** | Pointer to **string** | The prefix for the displayId of tasks of this schema. | [optional] 
**StatusLifecycle** | Pointer to [**WorkflowTaskStatusLifecycle**](WorkflowTaskStatusLifecycle.md) |  | [optional] 
**TaskGroupPrefix** | Pointer to **string** | The prefix for the displayId of task groups containing tasks of this schema | [optional] 
**WorkflowOutputSchema** | Pointer to [**NullableWorkflowOutputSchema**](WorkflowOutputSchema.md) |  | [optional] 

## Methods

### NewWorkflowTaskSchemaBase

`func NewWorkflowTaskSchemaBase() *WorkflowTaskSchemaBase`

NewWorkflowTaskSchemaBase instantiates a new WorkflowTaskSchemaBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskSchemaBaseWithDefaults

`func NewWorkflowTaskSchemaBaseWithDefaults() *WorkflowTaskSchemaBase`

NewWorkflowTaskSchemaBaseWithDefaults instantiates a new WorkflowTaskSchemaBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *WorkflowTaskSchemaBase) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *WorkflowTaskSchemaBase) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *WorkflowTaskSchemaBase) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *WorkflowTaskSchemaBase) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *WorkflowTaskSchemaBase) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *WorkflowTaskSchemaBase) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *WorkflowTaskSchemaBase) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *WorkflowTaskSchemaBase) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *WorkflowTaskSchemaBase) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *WorkflowTaskSchemaBase) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskSchemaBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskSchemaBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskSchemaBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskSchemaBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskSchemaBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskSchemaBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskSchemaBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskSchemaBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WorkflowTaskSchemaBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTaskSchemaBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTaskSchemaBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTaskSchemaBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchemaBase) GetCanSetAssigneeOnTaskCreation() bool`

GetCanSetAssigneeOnTaskCreation returns the CanSetAssigneeOnTaskCreation field if non-nil, zero value otherwise.

### GetCanSetAssigneeOnTaskCreationOk

`func (o *WorkflowTaskSchemaBase) GetCanSetAssigneeOnTaskCreationOk() (*bool, bool)`

GetCanSetAssigneeOnTaskCreationOk returns a tuple with the CanSetAssigneeOnTaskCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchemaBase) SetCanSetAssigneeOnTaskCreation(v bool)`

SetCanSetAssigneeOnTaskCreation sets CanSetAssigneeOnTaskCreation field to given value.

### HasCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchemaBase) HasCanSetAssigneeOnTaskCreation() bool`

HasCanSetAssigneeOnTaskCreation returns a boolean if a field has been set.

### GetDefaultCreationFolderId

`func (o *WorkflowTaskSchemaBase) GetDefaultCreationFolderId() string`

GetDefaultCreationFolderId returns the DefaultCreationFolderId field if non-nil, zero value otherwise.

### GetDefaultCreationFolderIdOk

`func (o *WorkflowTaskSchemaBase) GetDefaultCreationFolderIdOk() (*string, bool)`

GetDefaultCreationFolderIdOk returns a tuple with the DefaultCreationFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCreationFolderId

`func (o *WorkflowTaskSchemaBase) SetDefaultCreationFolderId(v string)`

SetDefaultCreationFolderId sets DefaultCreationFolderId field to given value.

### HasDefaultCreationFolderId

`func (o *WorkflowTaskSchemaBase) HasDefaultCreationFolderId() bool`

HasDefaultCreationFolderId returns a boolean if a field has been set.

### SetDefaultCreationFolderIdNil

`func (o *WorkflowTaskSchemaBase) SetDefaultCreationFolderIdNil(b bool)`

 SetDefaultCreationFolderIdNil sets the value for DefaultCreationFolderId to be an explicit nil

### UnsetDefaultCreationFolderId
`func (o *WorkflowTaskSchemaBase) UnsetDefaultCreationFolderId()`

UnsetDefaultCreationFolderId ensures that no value is present for DefaultCreationFolderId, not even an explicit nil
### GetDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchemaBase) GetDefaultEntryExecutionFolderId() string`

GetDefaultEntryExecutionFolderId returns the DefaultEntryExecutionFolderId field if non-nil, zero value otherwise.

### GetDefaultEntryExecutionFolderIdOk

`func (o *WorkflowTaskSchemaBase) GetDefaultEntryExecutionFolderIdOk() (*string, bool)`

GetDefaultEntryExecutionFolderIdOk returns a tuple with the DefaultEntryExecutionFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchemaBase) SetDefaultEntryExecutionFolderId(v string)`

SetDefaultEntryExecutionFolderId sets DefaultEntryExecutionFolderId field to given value.

### HasDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchemaBase) HasDefaultEntryExecutionFolderId() bool`

HasDefaultEntryExecutionFolderId returns a boolean if a field has been set.

### SetDefaultEntryExecutionFolderIdNil

`func (o *WorkflowTaskSchemaBase) SetDefaultEntryExecutionFolderIdNil(b bool)`

 SetDefaultEntryExecutionFolderIdNil sets the value for DefaultEntryExecutionFolderId to be an explicit nil

### UnsetDefaultEntryExecutionFolderId
`func (o *WorkflowTaskSchemaBase) UnsetDefaultEntryExecutionFolderId()`

UnsetDefaultEntryExecutionFolderId ensures that no value is present for DefaultEntryExecutionFolderId, not even an explicit nil
### GetDefaultResponsibleTeam

`func (o *WorkflowTaskSchemaBase) GetDefaultResponsibleTeam() TeamSummary`

GetDefaultResponsibleTeam returns the DefaultResponsibleTeam field if non-nil, zero value otherwise.

### GetDefaultResponsibleTeamOk

`func (o *WorkflowTaskSchemaBase) GetDefaultResponsibleTeamOk() (*TeamSummary, bool)`

GetDefaultResponsibleTeamOk returns a tuple with the DefaultResponsibleTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResponsibleTeam

`func (o *WorkflowTaskSchemaBase) SetDefaultResponsibleTeam(v TeamSummary)`

SetDefaultResponsibleTeam sets DefaultResponsibleTeam field to given value.

### HasDefaultResponsibleTeam

`func (o *WorkflowTaskSchemaBase) HasDefaultResponsibleTeam() bool`

HasDefaultResponsibleTeam returns a boolean if a field has been set.

### SetDefaultResponsibleTeamNil

`func (o *WorkflowTaskSchemaBase) SetDefaultResponsibleTeamNil(b bool)`

 SetDefaultResponsibleTeamNil sets the value for DefaultResponsibleTeam to be an explicit nil

### UnsetDefaultResponsibleTeam
`func (o *WorkflowTaskSchemaBase) UnsetDefaultResponsibleTeam()`

UnsetDefaultResponsibleTeam ensures that no value is present for DefaultResponsibleTeam, not even an explicit nil
### GetEntryTemplateId

`func (o *WorkflowTaskSchemaBase) GetEntryTemplateId() string`

GetEntryTemplateId returns the EntryTemplateId field if non-nil, zero value otherwise.

### GetEntryTemplateIdOk

`func (o *WorkflowTaskSchemaBase) GetEntryTemplateIdOk() (*string, bool)`

GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTemplateId

`func (o *WorkflowTaskSchemaBase) SetEntryTemplateId(v string)`

SetEntryTemplateId sets EntryTemplateId field to given value.

### HasEntryTemplateId

`func (o *WorkflowTaskSchemaBase) HasEntryTemplateId() bool`

HasEntryTemplateId returns a boolean if a field has been set.

### SetEntryTemplateIdNil

`func (o *WorkflowTaskSchemaBase) SetEntryTemplateIdNil(b bool)`

 SetEntryTemplateIdNil sets the value for EntryTemplateId to be an explicit nil

### UnsetEntryTemplateId
`func (o *WorkflowTaskSchemaBase) UnsetEntryTemplateId()`

UnsetEntryTemplateId ensures that no value is present for EntryTemplateId, not even an explicit nil
### GetPrefix

`func (o *WorkflowTaskSchemaBase) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WorkflowTaskSchemaBase) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WorkflowTaskSchemaBase) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WorkflowTaskSchemaBase) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetStatusLifecycle

`func (o *WorkflowTaskSchemaBase) GetStatusLifecycle() WorkflowTaskStatusLifecycle`

GetStatusLifecycle returns the StatusLifecycle field if non-nil, zero value otherwise.

### GetStatusLifecycleOk

`func (o *WorkflowTaskSchemaBase) GetStatusLifecycleOk() (*WorkflowTaskStatusLifecycle, bool)`

GetStatusLifecycleOk returns a tuple with the StatusLifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLifecycle

`func (o *WorkflowTaskSchemaBase) SetStatusLifecycle(v WorkflowTaskStatusLifecycle)`

SetStatusLifecycle sets StatusLifecycle field to given value.

### HasStatusLifecycle

`func (o *WorkflowTaskSchemaBase) HasStatusLifecycle() bool`

HasStatusLifecycle returns a boolean if a field has been set.

### GetTaskGroupPrefix

`func (o *WorkflowTaskSchemaBase) GetTaskGroupPrefix() string`

GetTaskGroupPrefix returns the TaskGroupPrefix field if non-nil, zero value otherwise.

### GetTaskGroupPrefixOk

`func (o *WorkflowTaskSchemaBase) GetTaskGroupPrefixOk() (*string, bool)`

GetTaskGroupPrefixOk returns a tuple with the TaskGroupPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroupPrefix

`func (o *WorkflowTaskSchemaBase) SetTaskGroupPrefix(v string)`

SetTaskGroupPrefix sets TaskGroupPrefix field to given value.

### HasTaskGroupPrefix

`func (o *WorkflowTaskSchemaBase) HasTaskGroupPrefix() bool`

HasTaskGroupPrefix returns a boolean if a field has been set.

### GetWorkflowOutputSchema

`func (o *WorkflowTaskSchemaBase) GetWorkflowOutputSchema() WorkflowOutputSchema`

GetWorkflowOutputSchema returns the WorkflowOutputSchema field if non-nil, zero value otherwise.

### GetWorkflowOutputSchemaOk

`func (o *WorkflowTaskSchemaBase) GetWorkflowOutputSchemaOk() (*WorkflowOutputSchema, bool)`

GetWorkflowOutputSchemaOk returns a tuple with the WorkflowOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutputSchema

`func (o *WorkflowTaskSchemaBase) SetWorkflowOutputSchema(v WorkflowOutputSchema)`

SetWorkflowOutputSchema sets WorkflowOutputSchema field to given value.

### HasWorkflowOutputSchema

`func (o *WorkflowTaskSchemaBase) HasWorkflowOutputSchema() bool`

HasWorkflowOutputSchema returns a boolean if a field has been set.

### SetWorkflowOutputSchemaNil

`func (o *WorkflowTaskSchemaBase) SetWorkflowOutputSchemaNil(b bool)`

 SetWorkflowOutputSchemaNil sets the value for WorkflowOutputSchema to be an explicit nil

### UnsetWorkflowOutputSchema
`func (o *WorkflowTaskSchemaBase) UnsetWorkflowOutputSchema()`

UnsetWorkflowOutputSchema ensures that no value is present for WorkflowOutputSchema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


