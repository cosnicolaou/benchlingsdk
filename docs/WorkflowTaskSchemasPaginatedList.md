# WorkflowTaskSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**WorkflowTaskSchemas** | Pointer to [**[]WorkflowTaskSchema**](WorkflowTaskSchema.md) |  | [optional] 

## Methods

### NewWorkflowTaskSchemasPaginatedList

`func NewWorkflowTaskSchemasPaginatedList() *WorkflowTaskSchemasPaginatedList`

NewWorkflowTaskSchemasPaginatedList instantiates a new WorkflowTaskSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskSchemasPaginatedListWithDefaults

`func NewWorkflowTaskSchemasPaginatedListWithDefaults() *WorkflowTaskSchemasPaginatedList`

NewWorkflowTaskSchemasPaginatedListWithDefaults instantiates a new WorkflowTaskSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *WorkflowTaskSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *WorkflowTaskSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *WorkflowTaskSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *WorkflowTaskSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetWorkflowTaskSchemas

`func (o *WorkflowTaskSchemasPaginatedList) GetWorkflowTaskSchemas() []WorkflowTaskSchema`

GetWorkflowTaskSchemas returns the WorkflowTaskSchemas field if non-nil, zero value otherwise.

### GetWorkflowTaskSchemasOk

`func (o *WorkflowTaskSchemasPaginatedList) GetWorkflowTaskSchemasOk() (*[]WorkflowTaskSchema, bool)`

GetWorkflowTaskSchemasOk returns a tuple with the WorkflowTaskSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskSchemas

`func (o *WorkflowTaskSchemasPaginatedList) SetWorkflowTaskSchemas(v []WorkflowTaskSchema)`

SetWorkflowTaskSchemas sets WorkflowTaskSchemas field to given value.

### HasWorkflowTaskSchemas

`func (o *WorkflowTaskSchemasPaginatedList) HasWorkflowTaskSchemas() bool`

HasWorkflowTaskSchemas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


