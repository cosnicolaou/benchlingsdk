# WorkflowTasksArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**WorkflowTaskArchiveReason**](WorkflowTaskArchiveReason.md) |  | 
**WorkflowTaskIds** | **[]string** |  | 

## Methods

### NewWorkflowTasksArchive

`func NewWorkflowTasksArchive(reason WorkflowTaskArchiveReason, workflowTaskIds []string, ) *WorkflowTasksArchive`

NewWorkflowTasksArchive instantiates a new WorkflowTasksArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTasksArchiveWithDefaults

`func NewWorkflowTasksArchiveWithDefaults() *WorkflowTasksArchive`

NewWorkflowTasksArchiveWithDefaults instantiates a new WorkflowTasksArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *WorkflowTasksArchive) GetReason() WorkflowTaskArchiveReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkflowTasksArchive) GetReasonOk() (*WorkflowTaskArchiveReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkflowTasksArchive) SetReason(v WorkflowTaskArchiveReason)`

SetReason sets Reason field to given value.


### GetWorkflowTaskIds

`func (o *WorkflowTasksArchive) GetWorkflowTaskIds() []string`

GetWorkflowTaskIds returns the WorkflowTaskIds field if non-nil, zero value otherwise.

### GetWorkflowTaskIdsOk

`func (o *WorkflowTasksArchive) GetWorkflowTaskIdsOk() (*[]string, bool)`

GetWorkflowTaskIdsOk returns a tuple with the WorkflowTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskIds

`func (o *WorkflowTasksArchive) SetWorkflowTaskIds(v []string)`

SetWorkflowTaskIds sets WorkflowTaskIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


