# WorkflowTask

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
**ExecutionType** | Pointer to **string** | The method by which the task of the workflow is executed | [optional] 

## Methods

### NewWorkflowTask

`func NewWorkflowTask() *WorkflowTask`

NewWorkflowTask instantiates a new WorkflowTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskWithDefaults

`func NewWorkflowTaskWithDefaults() *WorkflowTask`

NewWorkflowTaskWithDefaults instantiates a new WorkflowTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *WorkflowTask) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *WorkflowTask) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *WorkflowTask) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *WorkflowTask) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignee

`func (o *WorkflowTask) GetAssignee() UserSummary`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *WorkflowTask) GetAssigneeOk() (*UserSummary, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *WorkflowTask) SetAssignee(v UserSummary)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *WorkflowTask) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *WorkflowTask) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *WorkflowTask) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetClonedFrom

`func (o *WorkflowTask) GetClonedFrom() WorkflowTaskSummary`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *WorkflowTask) GetClonedFromOk() (*WorkflowTaskSummary, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *WorkflowTask) SetClonedFrom(v WorkflowTaskSummary)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *WorkflowTask) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### SetClonedFromNil

`func (o *WorkflowTask) SetClonedFromNil(b bool)`

 SetClonedFromNil sets the value for ClonedFrom to be an explicit nil

### UnsetClonedFrom
`func (o *WorkflowTask) UnsetClonedFrom()`

UnsetClonedFrom ensures that no value is present for ClonedFrom, not even an explicit nil
### GetCreationOrigin

`func (o *WorkflowTask) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowTask) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowTask) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowTask) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowTask) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowTask) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowTask) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowTask) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetExecutionOrigin

`func (o *WorkflowTask) GetExecutionOrigin() WorkflowTaskExecutionOrigin`

GetExecutionOrigin returns the ExecutionOrigin field if non-nil, zero value otherwise.

### GetExecutionOriginOk

`func (o *WorkflowTask) GetExecutionOriginOk() (*WorkflowTaskExecutionOrigin, bool)`

GetExecutionOriginOk returns a tuple with the ExecutionOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrigin

`func (o *WorkflowTask) SetExecutionOrigin(v WorkflowTaskExecutionOrigin)`

SetExecutionOrigin sets ExecutionOrigin field to given value.

### HasExecutionOrigin

`func (o *WorkflowTask) HasExecutionOrigin() bool`

HasExecutionOrigin returns a boolean if a field has been set.

### SetExecutionOriginNil

`func (o *WorkflowTask) SetExecutionOriginNil(b bool)`

 SetExecutionOriginNil sets the value for ExecutionOrigin to be an explicit nil

### UnsetExecutionOrigin
`func (o *WorkflowTask) UnsetExecutionOrigin()`

UnsetExecutionOrigin ensures that no value is present for ExecutionOrigin, not even an explicit nil
### GetFields

`func (o *WorkflowTask) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowTask) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowTask) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowTask) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTask) GetOutputs() []WorkflowOutputSummary`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTask) GetOutputsOk() (*[]WorkflowOutputSummary, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTask) SetOutputs(v []WorkflowOutputSummary)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTask) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetScheduledOn

`func (o *WorkflowTask) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *WorkflowTask) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *WorkflowTask) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *WorkflowTask) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### SetScheduledOnNil

`func (o *WorkflowTask) SetScheduledOnNil(b bool)`

 SetScheduledOnNil sets the value for ScheduledOn to be an explicit nil

### UnsetScheduledOn
`func (o *WorkflowTask) UnsetScheduledOn()`

UnsetScheduledOn ensures that no value is present for ScheduledOn, not even an explicit nil
### GetStatus

`func (o *WorkflowTask) GetStatus() WorkflowTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTask) GetStatusOk() (*WorkflowTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTask) SetStatus(v WorkflowTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowTask) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowTask) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowTask) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowTask) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *WorkflowTask) GetWorkflowTaskGroup() WorkflowTaskGroupSummary`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *WorkflowTask) GetWorkflowTaskGroupOk() (*WorkflowTaskGroupSummary, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *WorkflowTask) SetWorkflowTaskGroup(v WorkflowTaskGroupSummary)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *WorkflowTask) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.

### GetExecutionType

`func (o *WorkflowTask) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *WorkflowTask) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *WorkflowTask) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *WorkflowTask) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


