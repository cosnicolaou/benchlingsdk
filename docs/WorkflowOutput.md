# WorkflowOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayId** | Pointer to **string** | User-friendly ID of the workflow task group | [optional] 
**Id** | Pointer to **string** | The ID of the workflow output | [optional] [readonly] 
**CreationOrigin** | Pointer to [**CreationOrigin**](CreationOrigin.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Task** | Pointer to [**WorkflowTaskSummary**](WorkflowTaskSummary.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the workflow output | [optional] [readonly] 
**WorkflowTaskGroup** | Pointer to [**WorkflowTaskGroupSummary**](WorkflowTaskGroupSummary.md) |  | [optional] 

## Methods

### NewWorkflowOutput

`func NewWorkflowOutput() *WorkflowOutput`

NewWorkflowOutput instantiates a new WorkflowOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputWithDefaults

`func NewWorkflowOutputWithDefaults() *WorkflowOutput`

NewWorkflowOutputWithDefaults instantiates a new WorkflowOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayId

`func (o *WorkflowOutput) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *WorkflowOutput) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *WorkflowOutput) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *WorkflowOutput) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetId

`func (o *WorkflowOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationOrigin

`func (o *WorkflowOutput) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowOutput) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowOutput) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowOutput) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetFields

`func (o *WorkflowOutput) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowOutput) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowOutput) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowOutput) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetTask

`func (o *WorkflowOutput) GetTask() WorkflowTaskSummary`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *WorkflowOutput) GetTaskOk() (*WorkflowTaskSummary, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *WorkflowOutput) SetTask(v WorkflowTaskSummary)`

SetTask sets Task field to given value.

### HasTask

`func (o *WorkflowOutput) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowOutput) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowOutput) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowOutput) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowOutput) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *WorkflowOutput) GetWorkflowTaskGroup() WorkflowTaskGroupSummary`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *WorkflowOutput) GetWorkflowTaskGroupOk() (*WorkflowTaskGroupSummary, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *WorkflowOutput) SetWorkflowTaskGroup(v WorkflowTaskGroupSummary)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *WorkflowOutput) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


