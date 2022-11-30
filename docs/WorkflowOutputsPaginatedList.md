# WorkflowOutputsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**WorkflowOutputs** | Pointer to [**[]WorkflowOutput**](WorkflowOutput.md) |  | [optional] 

## Methods

### NewWorkflowOutputsPaginatedList

`func NewWorkflowOutputsPaginatedList() *WorkflowOutputsPaginatedList`

NewWorkflowOutputsPaginatedList instantiates a new WorkflowOutputsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputsPaginatedListWithDefaults

`func NewWorkflowOutputsPaginatedListWithDefaults() *WorkflowOutputsPaginatedList`

NewWorkflowOutputsPaginatedListWithDefaults instantiates a new WorkflowOutputsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *WorkflowOutputsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *WorkflowOutputsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *WorkflowOutputsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *WorkflowOutputsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetWorkflowOutputs

`func (o *WorkflowOutputsPaginatedList) GetWorkflowOutputs() []WorkflowOutput`

GetWorkflowOutputs returns the WorkflowOutputs field if non-nil, zero value otherwise.

### GetWorkflowOutputsOk

`func (o *WorkflowOutputsPaginatedList) GetWorkflowOutputsOk() (*[]WorkflowOutput, bool)`

GetWorkflowOutputsOk returns a tuple with the WorkflowOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutputs

`func (o *WorkflowOutputsPaginatedList) SetWorkflowOutputs(v []WorkflowOutput)`

SetWorkflowOutputs sets WorkflowOutputs field to given value.

### HasWorkflowOutputs

`func (o *WorkflowOutputsPaginatedList) HasWorkflowOutputs() bool`

HasWorkflowOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


