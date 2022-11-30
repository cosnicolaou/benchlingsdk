# WorkflowTasksPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**WorkflowTasks** | Pointer to [**[]WorkflowTask**](WorkflowTask.md) |  | [optional] 

## Methods

### NewWorkflowTasksPaginatedList

`func NewWorkflowTasksPaginatedList() *WorkflowTasksPaginatedList`

NewWorkflowTasksPaginatedList instantiates a new WorkflowTasksPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTasksPaginatedListWithDefaults

`func NewWorkflowTasksPaginatedListWithDefaults() *WorkflowTasksPaginatedList`

NewWorkflowTasksPaginatedListWithDefaults instantiates a new WorkflowTasksPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *WorkflowTasksPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *WorkflowTasksPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *WorkflowTasksPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *WorkflowTasksPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetWorkflowTasks

`func (o *WorkflowTasksPaginatedList) GetWorkflowTasks() []WorkflowTask`

GetWorkflowTasks returns the WorkflowTasks field if non-nil, zero value otherwise.

### GetWorkflowTasksOk

`func (o *WorkflowTasksPaginatedList) GetWorkflowTasksOk() (*[]WorkflowTask, bool)`

GetWorkflowTasksOk returns a tuple with the WorkflowTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTasks

`func (o *WorkflowTasksPaginatedList) SetWorkflowTasks(v []WorkflowTask)`

SetWorkflowTasks sets WorkflowTasks field to given value.

### HasWorkflowTasks

`func (o *WorkflowTasksPaginatedList) HasWorkflowTasks() bool`

HasWorkflowTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


