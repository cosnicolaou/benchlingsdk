# WorkflowTaskGroupsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**WorkflowTaskGroups** | Pointer to [**[]WorkflowTaskGroup**](WorkflowTaskGroup.md) |  | [optional] 

## Methods

### NewWorkflowTaskGroupsPaginatedList

`func NewWorkflowTaskGroupsPaginatedList() *WorkflowTaskGroupsPaginatedList`

NewWorkflowTaskGroupsPaginatedList instantiates a new WorkflowTaskGroupsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupsPaginatedListWithDefaults

`func NewWorkflowTaskGroupsPaginatedListWithDefaults() *WorkflowTaskGroupsPaginatedList`

NewWorkflowTaskGroupsPaginatedListWithDefaults instantiates a new WorkflowTaskGroupsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *WorkflowTaskGroupsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *WorkflowTaskGroupsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *WorkflowTaskGroupsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *WorkflowTaskGroupsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetWorkflowTaskGroups

`func (o *WorkflowTaskGroupsPaginatedList) GetWorkflowTaskGroups() []WorkflowTaskGroup`

GetWorkflowTaskGroups returns the WorkflowTaskGroups field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupsOk

`func (o *WorkflowTaskGroupsPaginatedList) GetWorkflowTaskGroupsOk() (*[]WorkflowTaskGroup, bool)`

GetWorkflowTaskGroupsOk returns a tuple with the WorkflowTaskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroups

`func (o *WorkflowTaskGroupsPaginatedList) SetWorkflowTaskGroups(v []WorkflowTaskGroup)`

SetWorkflowTaskGroups sets WorkflowTaskGroups field to given value.

### HasWorkflowTaskGroups

`func (o *WorkflowTaskGroupsPaginatedList) HasWorkflowTaskGroups() bool`

HasWorkflowTaskGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


