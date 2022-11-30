# WorkflowOutputsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**WorkflowOutputArchiveReason**](WorkflowOutputArchiveReason.md) |  | 
**WorkflowOutputIds** | **[]string** |  | 

## Methods

### NewWorkflowOutputsArchive

`func NewWorkflowOutputsArchive(reason WorkflowOutputArchiveReason, workflowOutputIds []string, ) *WorkflowOutputsArchive`

NewWorkflowOutputsArchive instantiates a new WorkflowOutputsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputsArchiveWithDefaults

`func NewWorkflowOutputsArchiveWithDefaults() *WorkflowOutputsArchive`

NewWorkflowOutputsArchiveWithDefaults instantiates a new WorkflowOutputsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *WorkflowOutputsArchive) GetReason() WorkflowOutputArchiveReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkflowOutputsArchive) GetReasonOk() (*WorkflowOutputArchiveReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkflowOutputsArchive) SetReason(v WorkflowOutputArchiveReason)`

SetReason sets Reason field to given value.


### GetWorkflowOutputIds

`func (o *WorkflowOutputsArchive) GetWorkflowOutputIds() []string`

GetWorkflowOutputIds returns the WorkflowOutputIds field if non-nil, zero value otherwise.

### GetWorkflowOutputIdsOk

`func (o *WorkflowOutputsArchive) GetWorkflowOutputIdsOk() (*[]string, bool)`

GetWorkflowOutputIdsOk returns a tuple with the WorkflowOutputIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutputIds

`func (o *WorkflowOutputsArchive) SetWorkflowOutputIds(v []string)`

SetWorkflowOutputIds sets WorkflowOutputIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


