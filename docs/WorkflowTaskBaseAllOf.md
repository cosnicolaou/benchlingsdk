# WorkflowTaskBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkflowTaskBaseAllOf

`func NewWorkflowTaskBaseAllOf() *WorkflowTaskBaseAllOf`

NewWorkflowTaskBaseAllOf instantiates a new WorkflowTaskBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskBaseAllOfWithDefaults

`func NewWorkflowTaskBaseAllOfWithDefaults() *WorkflowTaskBaseAllOf`

NewWorkflowTaskBaseAllOfWithDefaults instantiates a new WorkflowTaskBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *WorkflowTaskBaseAllOf) GetAssignee() UserSummary`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *WorkflowTaskBaseAllOf) GetAssigneeOk() (*UserSummary, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *WorkflowTaskBaseAllOf) SetAssignee(v UserSummary)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *WorkflowTaskBaseAllOf) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *WorkflowTaskBaseAllOf) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *WorkflowTaskBaseAllOf) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetClonedFrom

`func (o *WorkflowTaskBaseAllOf) GetClonedFrom() WorkflowTaskSummary`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *WorkflowTaskBaseAllOf) GetClonedFromOk() (*WorkflowTaskSummary, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *WorkflowTaskBaseAllOf) SetClonedFrom(v WorkflowTaskSummary)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *WorkflowTaskBaseAllOf) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### SetClonedFromNil

`func (o *WorkflowTaskBaseAllOf) SetClonedFromNil(b bool)`

 SetClonedFromNil sets the value for ClonedFrom to be an explicit nil

### UnsetClonedFrom
`func (o *WorkflowTaskBaseAllOf) UnsetClonedFrom()`

UnsetClonedFrom ensures that no value is present for ClonedFrom, not even an explicit nil
### GetCreationOrigin

`func (o *WorkflowTaskBaseAllOf) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowTaskBaseAllOf) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowTaskBaseAllOf) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowTaskBaseAllOf) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetCreator

`func (o *WorkflowTaskBaseAllOf) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WorkflowTaskBaseAllOf) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WorkflowTaskBaseAllOf) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WorkflowTaskBaseAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetExecutionOrigin

`func (o *WorkflowTaskBaseAllOf) GetExecutionOrigin() WorkflowTaskExecutionOrigin`

GetExecutionOrigin returns the ExecutionOrigin field if non-nil, zero value otherwise.

### GetExecutionOriginOk

`func (o *WorkflowTaskBaseAllOf) GetExecutionOriginOk() (*WorkflowTaskExecutionOrigin, bool)`

GetExecutionOriginOk returns a tuple with the ExecutionOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrigin

`func (o *WorkflowTaskBaseAllOf) SetExecutionOrigin(v WorkflowTaskExecutionOrigin)`

SetExecutionOrigin sets ExecutionOrigin field to given value.

### HasExecutionOrigin

`func (o *WorkflowTaskBaseAllOf) HasExecutionOrigin() bool`

HasExecutionOrigin returns a boolean if a field has been set.

### SetExecutionOriginNil

`func (o *WorkflowTaskBaseAllOf) SetExecutionOriginNil(b bool)`

 SetExecutionOriginNil sets the value for ExecutionOrigin to be an explicit nil

### UnsetExecutionOrigin
`func (o *WorkflowTaskBaseAllOf) UnsetExecutionOrigin()`

UnsetExecutionOrigin ensures that no value is present for ExecutionOrigin, not even an explicit nil
### GetFields

`func (o *WorkflowTaskBaseAllOf) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowTaskBaseAllOf) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowTaskBaseAllOf) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowTaskBaseAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetOutputs

`func (o *WorkflowTaskBaseAllOf) GetOutputs() []WorkflowOutputSummary`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *WorkflowTaskBaseAllOf) GetOutputsOk() (*[]WorkflowOutputSummary, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *WorkflowTaskBaseAllOf) SetOutputs(v []WorkflowOutputSummary)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *WorkflowTaskBaseAllOf) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetScheduledOn

`func (o *WorkflowTaskBaseAllOf) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *WorkflowTaskBaseAllOf) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *WorkflowTaskBaseAllOf) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *WorkflowTaskBaseAllOf) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### SetScheduledOnNil

`func (o *WorkflowTaskBaseAllOf) SetScheduledOnNil(b bool)`

 SetScheduledOnNil sets the value for ScheduledOn to be an explicit nil

### UnsetScheduledOn
`func (o *WorkflowTaskBaseAllOf) UnsetScheduledOn()`

UnsetScheduledOn ensures that no value is present for ScheduledOn, not even an explicit nil
### GetStatus

`func (o *WorkflowTaskBaseAllOf) GetStatus() WorkflowTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowTaskBaseAllOf) GetStatusOk() (*WorkflowTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowTaskBaseAllOf) SetStatus(v WorkflowTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowTaskBaseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowTaskBaseAllOf) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowTaskBaseAllOf) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowTaskBaseAllOf) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowTaskBaseAllOf) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *WorkflowTaskBaseAllOf) GetWorkflowTaskGroup() WorkflowTaskGroupSummary`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *WorkflowTaskBaseAllOf) GetWorkflowTaskGroupOk() (*WorkflowTaskGroupSummary, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *WorkflowTaskBaseAllOf) SetWorkflowTaskGroup(v WorkflowTaskGroupSummary)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *WorkflowTaskBaseAllOf) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


