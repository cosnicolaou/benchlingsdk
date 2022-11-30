# WorkflowTaskGroupsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**WorkflowTaskGroupArchiveReason**](WorkflowTaskGroupArchiveReason.md) |  | 
**WorkflowTaskGroupIds** | **[]string** |  | 

## Methods

### NewWorkflowTaskGroupsArchive

`func NewWorkflowTaskGroupsArchive(reason WorkflowTaskGroupArchiveReason, workflowTaskGroupIds []string, ) *WorkflowTaskGroupsArchive`

NewWorkflowTaskGroupsArchive instantiates a new WorkflowTaskGroupsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupsArchiveWithDefaults

`func NewWorkflowTaskGroupsArchiveWithDefaults() *WorkflowTaskGroupsArchive`

NewWorkflowTaskGroupsArchiveWithDefaults instantiates a new WorkflowTaskGroupsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *WorkflowTaskGroupsArchive) GetReason() WorkflowTaskGroupArchiveReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkflowTaskGroupsArchive) GetReasonOk() (*WorkflowTaskGroupArchiveReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkflowTaskGroupsArchive) SetReason(v WorkflowTaskGroupArchiveReason)`

SetReason sets Reason field to given value.


### GetWorkflowTaskGroupIds

`func (o *WorkflowTaskGroupsArchive) GetWorkflowTaskGroupIds() []string`

GetWorkflowTaskGroupIds returns the WorkflowTaskGroupIds field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupIdsOk

`func (o *WorkflowTaskGroupsArchive) GetWorkflowTaskGroupIdsOk() (*[]string, bool)`

GetWorkflowTaskGroupIdsOk returns a tuple with the WorkflowTaskGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroupIds

`func (o *WorkflowTaskGroupsArchive) SetWorkflowTaskGroupIds(v []string)`

SetWorkflowTaskGroupIds sets WorkflowTaskGroupIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


