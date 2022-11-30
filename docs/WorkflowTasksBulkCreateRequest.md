# WorkflowTasksBulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowTasks** | Pointer to [**[]WorkflowTaskBulkCreate**](WorkflowTaskBulkCreate.md) |  | [optional] 

## Methods

### NewWorkflowTasksBulkCreateRequest

`func NewWorkflowTasksBulkCreateRequest() *WorkflowTasksBulkCreateRequest`

NewWorkflowTasksBulkCreateRequest instantiates a new WorkflowTasksBulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTasksBulkCreateRequestWithDefaults

`func NewWorkflowTasksBulkCreateRequestWithDefaults() *WorkflowTasksBulkCreateRequest`

NewWorkflowTasksBulkCreateRequestWithDefaults instantiates a new WorkflowTasksBulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowTasks

`func (o *WorkflowTasksBulkCreateRequest) GetWorkflowTasks() []WorkflowTaskBulkCreate`

GetWorkflowTasks returns the WorkflowTasks field if non-nil, zero value otherwise.

### GetWorkflowTasksOk

`func (o *WorkflowTasksBulkCreateRequest) GetWorkflowTasksOk() (*[]WorkflowTaskBulkCreate, bool)`

GetWorkflowTasksOk returns a tuple with the WorkflowTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTasks

`func (o *WorkflowTasksBulkCreateRequest) SetWorkflowTasks(v []WorkflowTaskBulkCreate)`

SetWorkflowTasks sets WorkflowTasks field to given value.

### HasWorkflowTasks

`func (o *WorkflowTasksBulkCreateRequest) HasWorkflowTasks() bool`

HasWorkflowTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


