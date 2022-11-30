# WorkflowOutputAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationOrigin** | Pointer to [**CreationOrigin**](CreationOrigin.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Task** | Pointer to [**WorkflowTaskSummary**](WorkflowTaskSummary.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the workflow output | [optional] [readonly] 
**WorkflowTaskGroup** | Pointer to [**WorkflowTaskGroupSummary**](WorkflowTaskGroupSummary.md) |  | [optional] 

## Methods

### NewWorkflowOutputAllOf

`func NewWorkflowOutputAllOf() *WorkflowOutputAllOf`

NewWorkflowOutputAllOf instantiates a new WorkflowOutputAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputAllOfWithDefaults

`func NewWorkflowOutputAllOfWithDefaults() *WorkflowOutputAllOf`

NewWorkflowOutputAllOfWithDefaults instantiates a new WorkflowOutputAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationOrigin

`func (o *WorkflowOutputAllOf) GetCreationOrigin() CreationOrigin`

GetCreationOrigin returns the CreationOrigin field if non-nil, zero value otherwise.

### GetCreationOriginOk

`func (o *WorkflowOutputAllOf) GetCreationOriginOk() (*CreationOrigin, bool)`

GetCreationOriginOk returns a tuple with the CreationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationOrigin

`func (o *WorkflowOutputAllOf) SetCreationOrigin(v CreationOrigin)`

SetCreationOrigin sets CreationOrigin field to given value.

### HasCreationOrigin

`func (o *WorkflowOutputAllOf) HasCreationOrigin() bool`

HasCreationOrigin returns a boolean if a field has been set.

### GetFields

`func (o *WorkflowOutputAllOf) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowOutputAllOf) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowOutputAllOf) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowOutputAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetTask

`func (o *WorkflowOutputAllOf) GetTask() WorkflowTaskSummary`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *WorkflowOutputAllOf) GetTaskOk() (*WorkflowTaskSummary, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *WorkflowOutputAllOf) SetTask(v WorkflowTaskSummary)`

SetTask sets Task field to given value.

### HasTask

`func (o *WorkflowOutputAllOf) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetWebURL

`func (o *WorkflowOutputAllOf) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *WorkflowOutputAllOf) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *WorkflowOutputAllOf) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *WorkflowOutputAllOf) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *WorkflowOutputAllOf) GetWorkflowTaskGroup() WorkflowTaskGroupSummary`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *WorkflowOutputAllOf) GetWorkflowTaskGroupOk() (*WorkflowTaskGroupSummary, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *WorkflowOutputAllOf) SetWorkflowTaskGroup(v WorkflowTaskGroupSummary)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *WorkflowOutputAllOf) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


