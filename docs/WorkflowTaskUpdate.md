# WorkflowTaskUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** | The id of the user assigned to the task | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**ScheduledOn** | Pointer to **string** | The date on which the task is scheduled to be executed | [optional] 

## Methods

### NewWorkflowTaskUpdate

`func NewWorkflowTaskUpdate() *WorkflowTaskUpdate`

NewWorkflowTaskUpdate instantiates a new WorkflowTaskUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskUpdateWithDefaults

`func NewWorkflowTaskUpdateWithDefaults() *WorkflowTaskUpdate`

NewWorkflowTaskUpdateWithDefaults instantiates a new WorkflowTaskUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *WorkflowTaskUpdate) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *WorkflowTaskUpdate) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *WorkflowTaskUpdate) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *WorkflowTaskUpdate) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetAssigneeId

`func (o *WorkflowTaskUpdate) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *WorkflowTaskUpdate) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *WorkflowTaskUpdate) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *WorkflowTaskUpdate) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetFields

`func (o *WorkflowTaskUpdate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowTaskUpdate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowTaskUpdate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowTaskUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetScheduledOn

`func (o *WorkflowTaskUpdate) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *WorkflowTaskUpdate) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *WorkflowTaskUpdate) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *WorkflowTaskUpdate) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


