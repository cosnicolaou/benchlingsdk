# WorkflowTaskGroupBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationOrigin** | Pointer to [**CreationOrigin**](CreationOrigin.md) |  | [optional] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**Folder** | Pointer to [**Folder**](Folder.md) |  | [optional] 
**Outputs** | Pointer to [**[]WorkflowOutputSummary**](WorkflowOutputSummary.md) | The outputs of the workflow task group | [optional] 
**ResponsibleTeam** | Pointer to [**NullableTeamSummary**](TeamSummary.md) |  | [optional] 
**Watchers** | Pointer to [**[]UserSummary**](UserSummary.md) | The users watching the workflow task group | [optional] 
**WebURL** | Pointer to **string** | URL of the workflow task group | [optional] [readonly] 
**WorkflowTaskSchema** | Pointer to [**WorkflowTaskSchemaSummary**](WorkflowTaskSchemaSummary.md) |  | [optional] 
**WorkflowTasks** | Pointer to [**[]WorkflowTaskSummary**](WorkflowTaskSummary.md) | The input tasks to the workflow task group | [optional] 

## Methods

### NewWorkflowTaskGroupBaseAllOf

`func NewWorkflowTaskGroupBaseAllOf() *WorkflowTaskGroupBaseAllOf`

NewWorkflowTaskGroupBaseAllOf instantiates a new WorkflowTaskGroupBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupBaseAllOfWithDefaults

`func NewWorkflowTaskGroupBaseAllOfWithDefaults() *WorkflowTaskGroupBaseAllOf`

NewWorkflowTaskGroupBaseAllOfWithDefaults instantiates a new WorkflowTaskGroupBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationOrigin

`func (o *WorkflowTaskGroupBaseAllOf) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowTaskGroupBaseAllOf) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowTaskGroupBaseAllOf) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowTaskGroupBaseAllOf) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowTaskGroupBaseAllOf) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowTaskGroupBaseAllOf) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowTaskGroupBaseAllOf) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowTaskGroupBaseAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFolder

`func (o *WorkflowTaskGroupBaseAllOf) GetFolder() Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *WorkflowTaskGroupBaseAllOf) GetFolderOk() (*Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *WorkflowTaskGroupBaseAllOf) SetFolder(v Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *WorkflowTaskGroupBaseAllOf) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTaskGroupBaseAllOf) GetOutputs() []WorkflowOutputSummary`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTaskGroupBaseAllOf) GetOutputsOk() (*[]WorkflowOutputSummary, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTaskGroupBaseAllOf) SetOutputs(v []WorkflowOutputSummary)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTaskGroupBaseAllOf) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetResponsibleTeam

`func (o *WorkflowTaskGroupBaseAllOf) GetResponsibleTeam() TeamSummary`

GetResponsibleTeam returns the ResponsibleTeam field if non-nil, zero value otherwise.

### GetResponsibleTeamOk

`func (o *WorkflowTaskGroupBaseAllOf) GetResponsibleTeamOk() (*TeamSummary, bool)`

GetResponsibleTeamOk returns a tuple with the ResponsibleTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleTeam

`func (o *WorkflowTaskGroupBaseAllOf) SetResponsibleTeam(v TeamSummary)`

SetResponsibleTeam sets ResponsibleTeam field to given value.

### HasResponsibleTeam

`func (o *WorkflowTaskGroupBaseAllOf) HasResponsibleTeam() bool`

HasResponsibleTeam returns a boolean if a field has been set.

### SetResponsibleTeamNil

`func (o *WorkflowTaskGroupBaseAllOf) SetResponsibleTeamNil(b bool)`

 SetResponsibleTeamNil sets the value for ResponsibleTeam to be an explicit nil

### UnsetResponsibleTeam
`func (o *WorkflowTaskGroupBaseAllOf) UnsetResponsibleTeam()`

UnsetResponsibleTeam ensures that no value is present for ResponsibleTeam, not even an explicit nil
### GetWatchers

`func (o *WorkflowTaskGroupBaseAllOf) GetWatchers() []UserSummary`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *WorkflowTaskGroupBaseAllOf) GetWatchersOk() (*[]UserSummary, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *WorkflowTaskGroupBaseAllOf) SetWatchers(v []UserSummary)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *WorkflowTaskGroupBaseAllOf) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowTaskGroupBaseAllOf) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowTaskGroupBaseAllOf) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowTaskGroupBaseAllOf) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowTaskGroupBaseAllOf) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskSchema

`func (o *WorkflowTaskGroupBaseAllOf) GetWorkflowTaskSchema() WorkflowTaskSchemaSummary`

GetWorkflowTaskSchema returns the WorkflowTaskSchema field if non-nil, zero value otherwise.

### GetWorkflowTaskSchemaOk

`func (o *WorkflowTaskGroupBaseAllOf) GetWorkflowTaskSchemaOk() (*WorkflowTaskSchemaSummary, bool)`

GetWorkflowTaskSchemaOk returns a tuple with the WorkflowTaskSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskSchema

`func (o *WorkflowTaskGroupBaseAllOf) SetWorkflowTaskSchema(v WorkflowTaskSchemaSummary)`

SetWorkflowTaskSchema sets WorkflowTaskSchema field to given value.

### HasWorkflowTaskSchema

`func (o *WorkflowTaskGroupBaseAllOf) HasWorkflowTaskSchema() bool`

HasWorkflowTaskSchema returns a boolean if a field has been set.

### GetWorkflowTasks

`func (o *WorkflowTaskGroupBaseAllOf) GetWorkflowTasks() []WorkflowTaskSummary`

GetWorkflowTasks returns the WorkflowTasks field if non-nil, zero value otherwise.

### GetWorkflowTasksOk

`func (o *WorkflowTaskGroupBaseAllOf) GetWorkflowTasksOk() (*[]WorkflowTaskSummary, bool)`

GetWorkflowTasksOk returns a tuple with the WorkflowTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTasks

`func (o *WorkflowTaskGroupBaseAllOf) SetWorkflowTasks(v []WorkflowTaskSummary)`

SetWorkflowTasks sets WorkflowTasks field to given value.

### HasWorkflowTasks

`func (o *WorkflowTaskGroupBaseAllOf) HasWorkflowTasks() bool`

HasWorkflowTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


