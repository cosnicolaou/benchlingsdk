# WorkflowOutputBulkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**WorkflowTaskId** | **string** | The ID of the workflow task this output belogns to | 

## Methods

### NewWorkflowOutputBulkCreate

`func NewWorkflowOutputBulkCreate(workflowTaskId string, ) *WorkflowOutputBulkCreate`

NewWorkflowOutputBulkCreate instantiates a new WorkflowOutputBulkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputBulkCreateWithDefaults

`func NewWorkflowOutputBulkCreateWithDefaults() *WorkflowOutputBulkCreate`

NewWorkflowOutputBulkCreateWithDefaults instantiates a new WorkflowOutputBulkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *WorkflowOutputBulkCreate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowOutputBulkCreate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowOutputBulkCreate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowOutputBulkCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetWorkflowTaskId

`func (o *WorkflowOutputBulkCreate) GetWorkflowTaskId() string`

GetWorkflowTaskId returns the WorkflowTaskId field if non-nil, zero value otherwise.

### GetWorkflowTaskIdOk

`func (o *WorkflowOutputBulkCreate) GetWorkflowTaskIdOk() (*string, bool)`

GetWorkflowTaskIdOk returns a tuple with the WorkflowTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskId

`func (o *WorkflowOutputBulkCreate) SetWorkflowTaskId(v string)`

SetWorkflowTaskId sets WorkflowTaskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


