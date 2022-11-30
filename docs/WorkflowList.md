# WorkflowList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflows** | Pointer to [**[]LegacyWorkflow**](LegacyWorkflow.md) |  | [optional] 

## Methods

### NewWorkflowList

`func NewWorkflowList() *WorkflowList`

NewWorkflowList instantiates a new WorkflowList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowListWithDefaults

`func NewWorkflowListWithDefaults() *WorkflowList`

NewWorkflowListWithDefaults instantiates a new WorkflowList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflows

`func (o *WorkflowList) GetWorkflows() []LegacyWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowList) GetWorkflowsOk() (*[]LegacyWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowList) SetWorkflows(v []LegacyWorkflow)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowList) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


