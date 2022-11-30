# WorkflowTaskSchemaBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkflowTaskSchemaBaseAllOf

`func NewWorkflowTaskSchemaBaseAllOf() *WorkflowTaskSchemaBaseAllOf`

NewWorkflowTaskSchemaBaseAllOf instantiates a new WorkflowTaskSchemaBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskSchemaBaseAllOfWithDefaults

`func NewWorkflowTaskSchemaBaseAllOfWithDefaults() *WorkflowTaskSchemaBaseAllOf`

NewWorkflowTaskSchemaBaseAllOfWithDefaults instantiates a new WorkflowTaskSchemaBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchemaBaseAllOf) GetCanSetAssigneeOnTaskCreation() bool`

GetCanSetAssigneeOnTaskCreation returns the CanSetAssigneeOnTaskCreation field if non-nil, zero value otherwise.

### GetCanSetAssigneeOnTaskCreationOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetCanSetAssigneeOnTaskCreationOk() (*bool, bool)`

GetCanSetAssigneeOnTaskCreationOk returns a tuple with the CanSetAssigneeOnTaskCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchemaBaseAllOf) SetCanSetAssigneeOnTaskCreation(v bool)`

SetCanSetAssigneeOnTaskCreation sets CanSetAssigneeOnTaskCreation field to given value.

### HasCanSetAssigneeOnTaskCreation

`func (o *WorkflowTaskSchemaBaseAllOf) HasCanSetAssigneeOnTaskCreation() bool`

HasCanSetAssigneeOnTaskCreation returns a boolean if a field has been set.

### GetDefaultCreationFolderId

`func (o *WorkflowTaskSchemaBaseAllOf) GetDefaultCreationFolderId() string`

GetDefaultCreationFolderId returns the DefaultCreationFolderId field if non-nil, zero value otherwise.

### GetDefaultCreationFolderIdOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetDefaultCreationFolderIdOk() (*string, bool)`

GetDefaultCreationFolderIdOk returns a tuple with the DefaultCreationFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCreationFolderId

`func (o *WorkflowTaskSchemaBaseAllOf) SetDefaultCreationFolderId(v string)`

SetDefaultCreationFolderId sets DefaultCreationFolderId field to given value.

### HasDefaultCreationFolderId

`func (o *WorkflowTaskSchemaBaseAllOf) HasDefaultCreationFolderId() bool`

HasDefaultCreationFolderId returns a boolean if a field has been set.

### SetDefaultCreationFolderIdNil

`func (o *WorkflowTaskSchemaBaseAllOf) SetDefaultCreationFolderIdNil(b bool)`

 SetDefaultCreationFolderIdNil sets the value for DefaultCreationFolderId to be an explicit nil

### UnsetDefaultCreationFolderId
`func (o *WorkflowTaskSchemaBaseAllOf) UnsetDefaultCreationFolderId()`

UnsetDefaultCreationFolderId ensures that no value is present for DefaultCreationFolderId, not even an explicit nil
### GetDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchemaBaseAllOf) GetDefaultEntryExecutionFolderId() string`

GetDefaultEntryExecutionFolderId returns the DefaultEntryExecutionFolderId field if non-nil, zero value otherwise.

### GetDefaultEntryExecutionFolderIdOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetDefaultEntryExecutionFolderIdOk() (*string, bool)`

GetDefaultEntryExecutionFolderIdOk returns a tuple with the DefaultEntryExecutionFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchemaBaseAllOf) SetDefaultEntryExecutionFolderId(v string)`

SetDefaultEntryExecutionFolderId sets DefaultEntryExecutionFolderId field to given value.

### HasDefaultEntryExecutionFolderId

`func (o *WorkflowTaskSchemaBaseAllOf) HasDefaultEntryExecutionFolderId() bool`

HasDefaultEntryExecutionFolderId returns a boolean if a field has been set.

### SetDefaultEntryExecutionFolderIdNil

`func (o *WorkflowTaskSchemaBaseAllOf) SetDefaultEntryExecutionFolderIdNil(b bool)`

 SetDefaultEntryExecutionFolderIdNil sets the value for DefaultEntryExecutionFolderId to be an explicit nil

### UnsetDefaultEntryExecutionFolderId
`func (o *WorkflowTaskSchemaBaseAllOf) UnsetDefaultEntryExecutionFolderId()`

UnsetDefaultEntryExecutionFolderId ensures that no value is present for DefaultEntryExecutionFolderId, not even an explicit nil
### GetDefaultResponsibleTeam

`func (o *WorkflowTaskSchemaBaseAllOf) GetDefaultResponsibleTeam() TeamSummary`

GetDefaultResponsibleTeam returns the DefaultResponsibleTeam field if non-nil, zero value otherwise.

### GetDefaultResponsibleTeamOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetDefaultResponsibleTeamOk() (*TeamSummary, bool)`

GetDefaultResponsibleTeamOk returns a tuple with the DefaultResponsibleTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResponsibleTeam

`func (o *WorkflowTaskSchemaBaseAllOf) SetDefaultResponsibleTeam(v TeamSummary)`

SetDefaultResponsibleTeam sets DefaultResponsibleTeam field to given value.

### HasDefaultResponsibleTeam

`func (o *WorkflowTaskSchemaBaseAllOf) HasDefaultResponsibleTeam() bool`

HasDefaultResponsibleTeam returns a boolean if a field has been set.

### SetDefaultResponsibleTeamNil

`func (o *WorkflowTaskSchemaBaseAllOf) SetDefaultResponsibleTeamNil(b bool)`

 SetDefaultResponsibleTeamNil sets the value for DefaultResponsibleTeam to be an explicit nil

### UnsetDefaultResponsibleTeam
`func (o *WorkflowTaskSchemaBaseAllOf) UnsetDefaultResponsibleTeam()`

UnsetDefaultResponsibleTeam ensures that no value is present for DefaultResponsibleTeam, not even an explicit nil
### GetEntryTemplateId

`func (o *WorkflowTaskSchemaBaseAllOf) GetEntryTemplateId() string`

GetEntryTemplateId returns the EntryTemplateId field if non-nil, zero value otherwise.

### GetEntryTemplateIdOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetEntryTemplateIdOk() (*string, bool)`

GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTemplateId

`func (o *WorkflowTaskSchemaBaseAllOf) SetEntryTemplateId(v string)`

SetEntryTemplateId sets EntryTemplateId field to given value.

### HasEntryTemplateId

`func (o *WorkflowTaskSchemaBaseAllOf) HasEntryTemplateId() bool`

HasEntryTemplateId returns a boolean if a field has been set.

### SetEntryTemplateIdNil

`func (o *WorkflowTaskSchemaBaseAllOf) SetEntryTemplateIdNil(b bool)`

 SetEntryTemplateIdNil sets the value for EntryTemplateId to be an explicit nil

### UnsetEntryTemplateId
`func (o *WorkflowTaskSchemaBaseAllOf) UnsetEntryTemplateId()`

UnsetEntryTemplateId ensures that no value is present for EntryTemplateId, not even an explicit nil
### GetPrefix

`func (o *WorkflowTaskSchemaBaseAllOf) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WorkflowTaskSchemaBaseAllOf) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WorkflowTaskSchemaBaseAllOf) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetStatusLifecycle

`func (o *WorkflowTaskSchemaBaseAllOf) GetStatusLifecycle() WorkflowTaskStatusLifecycle`

GetStatusLifecycle returns the StatusLifecycle field if non-nil, zero value otherwise.

### GetStatusLifecycleOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetStatusLifecycleOk() (*WorkflowTaskStatusLifecycle, bool)`

GetStatusLifecycleOk returns a tuple with the StatusLifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLifecycle

`func (o *WorkflowTaskSchemaBaseAllOf) SetStatusLifecycle(v WorkflowTaskStatusLifecycle)`

SetStatusLifecycle sets StatusLifecycle field to given value.

### HasStatusLifecycle

`func (o *WorkflowTaskSchemaBaseAllOf) HasStatusLifecycle() bool`

HasStatusLifecycle returns a boolean if a field has been set.

### GetTaskGroupPrefix

`func (o *WorkflowTaskSchemaBaseAllOf) GetTaskGroupPrefix() string`

GetTaskGroupPrefix returns the TaskGroupPrefix field if non-nil, zero value otherwise.

### GetTaskGroupPrefixOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetTaskGroupPrefixOk() (*string, bool)`

GetTaskGroupPrefixOk returns a tuple with the TaskGroupPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroupPrefix

`func (o *WorkflowTaskSchemaBaseAllOf) SetTaskGroupPrefix(v string)`

SetTaskGroupPrefix sets TaskGroupPrefix field to given value.

### HasTaskGroupPrefix

`func (o *WorkflowTaskSchemaBaseAllOf) HasTaskGroupPrefix() bool`

HasTaskGroupPrefix returns a boolean if a field has been set.

### GetWorkflowOutputSchema

`func (o *WorkflowTaskSchemaBaseAllOf) GetWorkflowOutputSchema() WorkflowOutputSchema`

GetWorkflowOutputSchema returns the WorkflowOutputSchema field if non-nil, zero value otherwise.

### GetWorkflowOutputSchemaOk

`func (o *WorkflowTaskSchemaBaseAllOf) GetWorkflowOutputSchemaOk() (*WorkflowOutputSchema, bool)`

GetWorkflowOutputSchemaOk returns a tuple with the WorkflowOutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutputSchema

`func (o *WorkflowTaskSchemaBaseAllOf) SetWorkflowOutputSchema(v WorkflowOutputSchema)`

SetWorkflowOutputSchema sets WorkflowOutputSchema field to given value.

### HasWorkflowOutputSchema

`func (o *WorkflowTaskSchemaBaseAllOf) HasWorkflowOutputSchema() bool`

HasWorkflowOutputSchema returns a boolean if a field has been set.

### SetWorkflowOutputSchemaNil

`func (o *WorkflowTaskSchemaBaseAllOf) SetWorkflowOutputSchemaNil(b bool)`

 SetWorkflowOutputSchemaNil sets the value for WorkflowOutputSchema to be an explicit nil

### UnsetWorkflowOutputSchema
`func (o *WorkflowTaskSchemaBaseAllOf) UnsetWorkflowOutputSchema()`

UnsetWorkflowOutputSchema ensures that no value is present for WorkflowOutputSchema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


