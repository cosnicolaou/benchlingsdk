# WorkflowTaskGroupBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | User-friendly ID of the workflow task group | [optional] 
**Id** | Pointer to **string** | The ID of the workflow task group | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the workflow task group | [optional] 
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

### NewWorkflowTaskGroupBase

`func NewWorkflowTaskGroupBase() *WorkflowTaskGroupBase`

NewWorkflowTaskGroupBase instantiates a new WorkflowTaskGroupBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupBaseWithDefaults

`func NewWorkflowTaskGroupBaseWithDefaults() *WorkflowTaskGroupBase`

NewWorkflowTaskGroupBaseWithDefaults instantiates a new WorkflowTaskGroupBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *WorkflowTaskGroupBase) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *WorkflowTaskGroupBase) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *WorkflowTaskGroupBase) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *WorkflowTaskGroupBase) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskGroupBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskGroupBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskGroupBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskGroupBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskGroupBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskGroupBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskGroupBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskGroupBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationOrigin

`func (o *WorkflowTaskGroupBase) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowTaskGroupBase) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowTaskGroupBase) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowTaskGroupBase) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowTaskGroupBase) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowTaskGroupBase) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowTaskGroupBase) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowTaskGroupBase) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFolder

`func (o *WorkflowTaskGroupBase) GetFolder() Folder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *WorkflowTaskGroupBase) GetFolderOk() (*Folder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *WorkflowTaskGroupBase) SetFolder(v Folder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *WorkflowTaskGroupBase) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTaskGroupBase) GetOutputs() []WorkflowOutputSummary`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTaskGroupBase) GetOutputsOk() (*[]WorkflowOutputSummary, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTaskGroupBase) SetOutputs(v []WorkflowOutputSummary)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTaskGroupBase) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetResponsibleTeam

`func (o *WorkflowTaskGroupBase) GetResponsibleTeam() TeamSummary`

GetResponsibleTeam returns the ResponsibleTeam field if non-nil, zero value otherwise.

### GetResponsibleTeamOk

`func (o *WorkflowTaskGroupBase) GetResponsibleTeamOk() (*TeamSummary, bool)`

GetResponsibleTeamOk returns a tuple with the ResponsibleTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleTeam

`func (o *WorkflowTaskGroupBase) SetResponsibleTeam(v TeamSummary)`

SetResponsibleTeam sets ResponsibleTeam field to given value.

### HasResponsibleTeam

`func (o *WorkflowTaskGroupBase) HasResponsibleTeam() bool`

HasResponsibleTeam returns a boolean if a field has been set.

### SetResponsibleTeamNil

`func (o *WorkflowTaskGroupBase) SetResponsibleTeamNil(b bool)`

 SetResponsibleTeamNil sets the value for ResponsibleTeam to be an explicit nil

### UnsetResponsibleTeam
`func (o *WorkflowTaskGroupBase) UnsetResponsibleTeam()`

UnsetResponsibleTeam ensures that no value is present for ResponsibleTeam, not even an explicit nil
### GetWatchers

`func (o *WorkflowTaskGroupBase) GetWatchers() []UserSummary`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *WorkflowTaskGroupBase) GetWatchersOk() (*[]UserSummary, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *WorkflowTaskGroupBase) SetWatchers(v []UserSummary)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *WorkflowTaskGroupBase) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowTaskGroupBase) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowTaskGroupBase) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowTaskGroupBase) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowTaskGroupBase) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskSchema

`func (o *WorkflowTaskGroupBase) GetWorkflowTaskSchema() WorkflowTaskSchemaSummary`

GetWorkflowTaskSchema returns the WorkflowTaskSchema field if non-nil, zero value otherwise.

### GetWorkflowTaskSchemaOk

`func (o *WorkflowTaskGroupBase) GetWorkflowTaskSchemaOk() (*WorkflowTaskSchemaSummary, bool)`

GetWorkflowTaskSchemaOk returns a tuple with the WorkflowTaskSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskSchema

`func (o *WorkflowTaskGroupBase) SetWorkflowTaskSchema(v WorkflowTaskSchemaSummary)`

SetWorkflowTaskSchema sets WorkflowTaskSchema field to given value.

### HasWorkflowTaskSchema

`func (o *WorkflowTaskGroupBase) HasWorkflowTaskSchema() bool`

HasWorkflowTaskSchema returns a boolean if a field has been set.

### GetWorkflowTasks

`func (o *WorkflowTaskGroupBase) GetWorkflowTasks() []WorkflowTaskSummary`

GetWorkflowTasks returns the WorkflowTasks field if non-nil, zero value otherwise.

### GetWorkflowTasksOk

`func (o *WorkflowTaskGroupBase) GetWorkflowTasksOk() (*[]WorkflowTaskSummary, bool)`

GetWorkflowTasksOk returns a tuple with the WorkflowTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTasks

`func (o *WorkflowTaskGroupBase) SetWorkflowTasks(v []WorkflowTaskSummary)`

SetWorkflowTasks sets WorkflowTasks field to given value.

### HasWorkflowTasks

`func (o *WorkflowTaskGroupBase) HasWorkflowTasks() bool`

HasWorkflowTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


