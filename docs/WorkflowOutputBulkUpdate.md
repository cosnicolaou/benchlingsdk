# WorkflowOutputBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**WorkflowOutputId** | Pointer to **string** | The ID of the workflow output | [optional] 

## Methods

### NewWorkflowOutputBulkUpdate

`func NewWorkflowOutputBulkUpdate() *WorkflowOutputBulkUpdate`

NewWorkflowOutputBulkUpdate instantiates a new WorkflowOutputBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputBulkUpdateWithDefaults

`func NewWorkflowOutputBulkUpdateWithDefaults() *WorkflowOutputBulkUpdate`

NewWorkflowOutputBulkUpdateWithDefaults instantiates a new WorkflowOutputBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *WorkflowOutputBulkUpdate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowOutputBulkUpdate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowOutputBulkUpdate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowOutputBulkUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetWorkflowOutputId

`func (o *WorkflowOutputBulkUpdate) GetWorkflowOutputId() string`

GetWorkflowOutputId returns the WorkflowOutputId field if non-nil, zero value otherwise.

### GetWorkflowOutputIdOk

`func (o *WorkflowOutputBulkUpdate) GetWorkflowOutputIdOk() (*string, bool)`

GetWorkflowOutputIdOk returns a tuple with the WorkflowOutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutputId

`func (o *WorkflowOutputBulkUpdate) SetWorkflowOutputId(v string)`

SetWorkflowOutputId sets WorkflowOutputId field to given value.

### HasWorkflowOutputId

`func (o *WorkflowOutputBulkUpdate) HasWorkflowOutputId() bool`

HasWorkflowOutputId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


