# WorkflowTaskSchema

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
**ExecutionType** | Pointer to **string** | The method by which instances of this schema are executed | [optional] 

## Methods

### NewWorkflowTaskSchema

`func NewWorkflowTaskSchema() *WorkflowTaskSchema`

NewWorkflowTaskSchema instantiates a new WorkflowTaskSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskSchemaWithDefaults

`func NewWorkflowTaskSchemaWithDefaults() *WorkflowTaskSchema`

NewWorkflowTaskSchemaWithDefaults instantiates a new WorkflowTaskSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *WorkflowTaskSchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *WorkflowTaskSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *WorkflowTaskSchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *WorkflowTaskSchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *WorkflowTaskSchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *WorkflowTaskSchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *WorkflowTaskSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *WorkflowTaskSchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *WorkflowTaskSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *WorkflowTaskSchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WorkflowTaskSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTaskSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTaskSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTaskSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchema) GetCanSetAssigneeOnTaskCreation() bool`

GetCanSetAssigneeOnTaskCreation returns the CanSetAssigneeOnTaskCreation field if non-nil, zero value otherwise.

### GetCanSetAssigneeOnTaskCreationOk

`func (o *WorkflowTaskSchema) GetCanSetAssigneeOnTaskCreationOk() (*bool, bool)`

GetCanSetAssigneeOnTaskCreationOk returns a tuple with the CanSetAssigneeOnTaskCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchema) SetCanSetAssigneeOnTaskCreation(v bool)`

SetCanSetAssigneeOnTaskCreation sets CanSetAssigneeOnTaskCreation field to given value.

### HasCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchema) HasCanSetAssigneeOnTaskCreation() bool`

HasCanSetAssigneeOnTaskCreation returns a boolean if a field has been set.

### GetDefaultCreationFolderId

`func (o *WorkflowTaskSchema) GetDefaultCreationFolderId() string`

GetDefaultCreationFolderId returns the DefaultCreationFolderId field if non-nil, zero value otherwise.

### GetDefaultCreationFolderIdOk

`func (o *WorkflowTaskSchema) GetDefaultCreationFolderIdOk() (*string, bool)`

GetDefaultCreationFolderIdOk returns a tuple with the DefaultCreationFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCreationFolderId

`func (o *WorkflowTaskSchema) SetDefaultCreationFolderId(v string)`

SetDefaultCreationFolderId sets DefaultCreationFolderId field to given value.

### HasDefaultCreationFolderId

`func (o *WorkflowTaskSchema) HasDefaultCreationFolderId() bool`

HasDefaultCreationFolderId returns a boolean if a field has been set.

### SetDefaultCreationFolderIdNil

`func (o *WorkflowTaskSchema) SetDefaultCreationFolderIdNil(b bool)`

 SetDefaultCreationFolderIdNil sets the value for DefaultCreationFolderId to be an explicit nil

### UnsetDefaultCreationFolderId
`func (o *WorkflowTaskSchema) UnsetDefaultCreationFolderId()`

UnsetDefaultCreationFolderId ensures that no value is present for DefaultCreationFolderId, not even an explicit nil
### GetDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchema) GetDefaultEntryExecutionFolderId() string`

GetDefaultEntryExecutionFolderId returns the DefaultEntryExecutionFolderId field if non-nil, zero value otherwise.

### GetDefaultEntryExecutionFolderIdOk

`func (o *WorkflowTaskSchema) GetDefaultEntryExecutionFolderIdOk() (*string, bool)`

GetDefaultEntryExecutionFolderIdOk returns a tuple with the DefaultEntryExecutionFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchema) SetDefaultEntryExecutionFolderId(v string)`

SetDefaultEntryExecutionFolderId sets DefaultEntryExecutionFolderId field to given value.

### HasDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchema) HasDefaultEntryExecutionFolderId() bool`

HasDefaultEntryExecutionFolderId returns a boolean if a field has been set.

### SetDefaultEntryExecutionFolderIdNil

`func (o *WorkflowTaskSchema) SetDefaultEntryExecutionFolderIdNil(b bool)`

 SetDefaultEntryExecutionFolderIdNil sets the value for DefaultEntryExecutionFolderId to be an explicit nil

### UnsetDefaultEntryExecutionFolderId
`func (o *WorkflowTaskSchema) UnsetDefaultEntryExecutionFolderId()`

UnsetDefaultEntryExecutionFolderId ensures that no value is present for DefaultEntryExecutionFolderId, not even an explicit nil
### GetDefaultResponsibleTeam

`func (o *WorkflowTaskSchema) GetDefaultResponsibleTeam() TeamSummary`

GetDefaultResponsibleTeam returns the DefaultResponsibleTeam field if non-nil, zero value otherwise.

### GetDefaultResponsibleTeamOk

`func (o *WorkflowTaskSchema) GetDefaultResponsibleTeamOk() (*TeamSummary, bool)`

GetDefaultResponsibleTeamOk returns a tuple with the DefaultResponsibleTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResponsibleTeam

`func (o *WorkflowTaskSchema) SetDefaultResponsibleTeam(v TeamSummary)`

SetDefaultResponsibleTeam sets DefaultResponsibleTeam field to given value.

### HasDefaultResponsibleTeam

`func (o *WorkflowTaskSchema) HasDefaultResponsibleTeam() bool`

HasDefaultResponsibleTeam returns a boolean if a field has been set.

### SetDefaultResponsibleTeamNil

`func (o *WorkflowTaskSchema) SetDefaultResponsibleTeamNil(b bool)`

 SetDefaultResponsibleTeamNil sets the value for DefaultResponsibleTeam to be an explicit nil

### UnsetDefaultResponsibleTeam
`func (o *WorkflowTaskSchema) UnsetDefaultResponsibleTeam()`

UnsetDefaultResponsibleTeam ensures that no value is present for DefaultResponsibleTeam, not even an explicit nil
### GetEntryTemplateId

`func (o *WorkflowTaskSchema) GetEntryTemplateId() string`

GetEntryTemplateId returns the EntryTemplateId field if non-nil, zero value otherwise.

### GetEntryTemplateIdOk

`func (o *WorkflowTaskSchema) GetEntryTemplateIdOk() (*string, bool)`

GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTemplateId

`func (o *WorkflowTaskSchema) SetEntryTemplateId(v string)`

SetEntryTemplateId sets EntryTemplateId field to given value.

### HasEntryTemplateId

`func (o *WorkflowTaskSchema) HasEntryTemplateId() bool`

HasEntryTemplateId returns a boolean if a field has been set.

### SetEntryTemplateIdNil

`func (o *WorkflowTaskSchema) SetEntryTemplateIdNil(b bool)`

 SetEntryTemplateIdNil sets the value for EntryTemplateId to be an explicit nil

### UnsetEntryTemplateId
`func (o *WorkflowTaskSchema) UnsetEntryTemplateId()`

UnsetEntryTemplateId ensures that no value is present for EntryTemplateId, not even an explicit nil
### GetPrefix

`func (o *WorkflowTaskSchema) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WorkflowTaskSchema) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WorkflowTaskSchema) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WorkflowTaskSchema) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetStatusLifecycle

`func (o *WorkflowTaskSchema) GetStatusLifecycle() WorkflowTaskStatusLifecycle`

GetStatusLifecycle returns the StatusLifecycle field if non-nil, zero value otherwise.

### GetStatusLifecycleOk

`func (o *WorkflowTaskSchema) GetStatusLifecycleOk() (*WorkflowTaskStatusLifecycle, bool)`

GetStatusLifecycleOk returns a tuple with the StatusLifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLifecycle

`func (o *WorkflowTaskSchema) SetStatusLifecycle(v WorkflowTaskStatusLifecycle)`

SetStatusLifecycle sets StatusLifecycle field to given value.

### HasStatusLifecycle

`func (o *WorkflowTaskSchema) HasStatusLifecycle() bool`

HasStatusLifecycle returns a boolean if a field has been set.

### GetTaskGroupPrefix

`func (o *WorkflowTaskSchema) GetTaskGroupPrefix() string`

GetTaskGroupPrefix returns the TaskGroupPrefix field if non-nil, zero value otherwise.

### GetTaskGroupPrefixOk

`func (o *WorkflowTaskSchema) GetTaskGroupPrefixOk() (*string, bool)`

GetTaskGroupPrefixOk returns a tuple with the TaskGroupPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroupPrefix

`func (o *WorkflowTaskSchema) SetTaskGroupPrefix(v string)`

SetTaskGroupPrefix sets TaskGroupPrefix field to given value.

### HasTaskGroupPrefix

`func (o *WorkflowTaskSchema) HasTaskGroupPrefix() bool`

HasTaskGroupPrefix returns a boolean if a field has been set.

### GetWorkflowOutputSchema

`func (o *WorkflowTaskSchema) GetWorkflowOutputSchema() WorkflowOutputSchema`

GetWorkflowOutputSchema returns the WorkflowOutputSchema field if non-nil, zero value otherwise.

### GetWorkflowOutputSchemaOk

`func (o *WorkflowTaskSchema) GetWorkflowOutputSchemaOk() (*WorkflowOutputSchema, bool)`

GetWorkflowOutputSchemaOk returns a tuple with the WorkflowOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutputSchema

`func (o *WorkflowTaskSchema) SetWorkflowOutputSchema(v WorkflowOutputSchema)`

SetWorkflowOutputSchema sets WorkflowOutputSchema field to given value.

### HasWorkflowOutputSchema

`func (o *WorkflowTaskSchema) HasWorkflowOutputSchema() bool`

HasWorkflowOutputSchema returns a boolean if a field has been set.

### SetWorkflowOutputSchemaNil

`func (o *WorkflowTaskSchema) SetWorkflowOutputSchemaNil(b bool)`

 SetWorkflowOutputSchemaNil sets the value for WorkflowOutputSchema to be an explicit nil

### UnsetWorkflowOutputSchema
`func (o *WorkflowTaskSchema) UnsetWorkflowOutputSchema()`

UnsetWorkflowOutputSchema ensures that no value is present for WorkflowOutputSchema, not even an explicit nil
### GetExecutionType

`func (o *WorkflowTaskSchema) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *WorkflowTaskSchema) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *WorkflowTaskSchema) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *WorkflowTaskSchema) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


