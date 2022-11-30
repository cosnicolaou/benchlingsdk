# WorkflowOutputCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**WorkflowTaskId** | **string** | The ID of the workflow task this output belogns to | 

## Methods

### NewWorkflowOutputCreate

`func NewWorkflowOutputCreate(workflowTaskId string, ) *WorkflowOutputCreate`

NewWorkflowOutputCreate instantiates a new WorkflowOutputCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputCreateWithDefaults

`func NewWorkflowOutputCreateWithDefaults() *WorkflowOutputCreate`

NewWorkflowOutputCreateWithDefaults instantiates a new WorkflowOutputCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *WorkflowOutputCreate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowOutputCreate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowOutputCreate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowOutputCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetWorkflowTaskId

`func (o *WorkflowOutputCreate) GetWorkflowTaskId() string`

GetWorkflowTaskId returns the WorkflowTaskId field if non-nil, zero value otherwise.

### GetWorkflowTaskIdOk

`func (o *WorkflowOutputCreate) GetWorkflowTaskIdOk() (*string, bool)`

GetWorkflowTaskIdOk returns a tuple with the WorkflowTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskId

`func (o *WorkflowOutputCreate) SetWorkflowTaskId(v string)`

SetWorkflowTaskId sets WorkflowTaskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


