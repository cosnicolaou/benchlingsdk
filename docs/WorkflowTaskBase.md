# WorkflowTaskBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | User-friendly ID of the workflow task | [optional] 
**Id** | Pointer to **string** | The ID of the workflow task | [optional] [readonly] 
**Assignee** | Pointer to [**NullableUserSummary**](UserSummary.md) |  | [optional] 
**ClonedFrom** | Pointer to [**NullableWorkflowTaskSummary**](WorkflowTaskSummary.md) |  | [optional] 
**CreationOrigin** | Pointer to [**CreationOrigin**](CreationOrigin.md) |  | [optional] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**ExecutionOrigin** | Pointer to [**NullableWorkflowTaskExecutionOrigin**](WorkflowTaskExecutionOrigin.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Outputs** | Pointer to [**[]WorkflowOutputSummary**](WorkflowOutputSummary.md) |  | [optional] 
**ScheduledOn** | Pointer to **NullableString** | The date on which the task is scheduled to be executed | [optional] 
**Status** | Pointer to [**WorkflowTaskStatus**](WorkflowTaskStatus.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the workflow task | [optional] [readonly] 
**WorkflowTaskGroup** | Pointer to [**WorkflowTaskGroupSummary**](WorkflowTaskGroupSummary.md) |  | [optional] 

## Methods

### NewWorkflowTaskBase

`func NewWorkflowTaskBase() *WorkflowTaskBase`

NewWorkflowTaskBase instantiates a new WorkflowTaskBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskBaseWithDefaults

`func NewWorkflowTaskBaseWithDefaults() *WorkflowTaskBase`

NewWorkflowTaskBaseWithDefaults instantiates a new WorkflowTaskBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *WorkflowTaskBase) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *WorkflowTaskBase) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *WorkflowTaskBase) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *WorkflowTaskBase) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignee

`func (o *WorkflowTaskBase) GetAssignee() UserSummary`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *WorkflowTaskBase) GetAssigneeOk() (*UserSummary, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *WorkflowTaskBase) SetAssignee(v UserSummary)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *WorkflowTaskBase) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *WorkflowTaskBase) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *WorkflowTaskBase) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetClonedFrom

`func (o *WorkflowTaskBase) GetClonedFrom() WorkflowTaskSummary`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *WorkflowTaskBase) GetClonedFromOk() (*WorkflowTaskSummary, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *WorkflowTaskBase) SetClonedFrom(v WorkflowTaskSummary)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *WorkflowTaskBase) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### SetClonedFromNil

`func (o *WorkflowTaskBase) SetClonedFromNil(b bool)`

 SetClonedFromNil sets the value for ClonedFrom to be an explicit nil

### UnsetClonedFrom
`func (o *WorkflowTaskBase) UnsetClonedFrom()`

UnsetClonedFrom ensures that no value is present for ClonedFrom, not even an explicit nil
### GetCreationOrigin

`func (o *WorkflowTaskBase) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowTaskBase) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowTaskBase) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowTaskBase) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowTaskBase) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowTaskBase) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowTaskBase) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowTaskBase) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetExecutionOrigin

`func (o *WorkflowTaskBase) GetExecutionOrigin() WorkflowTaskExecutionOrigin`

GetExecutionOrigin returns the ExecutionOrigin field if non-nil, zero value otherwise.

### GetExecutionOriginOk

`func (o *WorkflowTaskBase) GetExecutionOriginOk() (*WorkflowTaskExecutionOrigin, bool)`

GetExecutionOriginOk returns a tuple with the ExecutionOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrigin

`func (o *WorkflowTaskBase) SetExecutionOrigin(v WorkflowTaskExecutionOrigin)`

SetExecutionOrigin sets ExecutionOrigin field to given value.

### HasExecutionOrigin

`func (o *WorkflowTaskBase) HasExecutionOrigin() bool`

HasExecutionOrigin returns a boolean if a field has been set.

### SetExecutionOriginNil

`func (o *WorkflowTaskBase) SetExecutionOriginNil(b bool)`

 SetExecutionOriginNil sets the value for ExecutionOrigin to be an explicit nil

### UnsetExecutionOrigin
`func (o *WorkflowTaskBase) UnsetExecutionOrigin()`

UnsetExecutionOrigin ensures that no value is present for ExecutionOrigin, not even an explicit nil
### GetFields

`func (o *WorkflowTaskBase) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowTaskBase) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowTaskBase) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowTaskBase) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTaskBase) GetOutputs() []WorkflowOutputSummary`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTaskBase) GetOutputsOk() (*[]WorkflowOutputSummary, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTaskBase) SetOutputs(v []WorkflowOutputSummary)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTaskBase) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetScheduledOn

`func (o *WorkflowTaskBase) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *WorkflowTaskBase) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *WorkflowTaskBase) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *WorkflowTaskBase) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### SetScheduledOnNil

`func (o *WorkflowTaskBase) SetScheduledOnNil(b bool)`

 SetScheduledOnNil sets the value for ScheduledOn to be an explicit nil

### UnsetScheduledOn
`func (o *WorkflowTaskBase) UnsetScheduledOn()`

UnsetScheduledOn ensures that no value is present for ScheduledOn, not even an explicit nil
### GetStatus

`func (o *WorkflowTaskBase) GetStatus() WorkflowTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTaskBase) GetStatusOk() (*WorkflowTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTaskBase) SetStatus(v WorkflowTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowTaskBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowTaskBase) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowTaskBase) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowTaskBase) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowTaskBase) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *WorkflowTaskBase) GetWorkflowTaskGroup() WorkflowTaskGroupSummary`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *WorkflowTaskBase) GetWorkflowTaskGroupOk() (*WorkflowTaskGroupSummary, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *WorkflowTaskBase) SetWorkflowTaskGroup(v WorkflowTaskGroupSummary)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *WorkflowTaskBase) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


