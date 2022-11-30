# WorkflowTaskBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** | The id of the user assigned to the task | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**ScheduledOn** | Pointer to **string** | The date on which the task is scheduled to be executed | [optional] 
**WorkflowTaskId** | Pointer to **string** | The workflow task ID | [optional] 

## Methods

### NewWorkflowTaskBulkUpdate

`func NewWorkflowTaskBulkUpdate() *WorkflowTaskBulkUpdate`

NewWorkflowTaskBulkUpdate instantiates a new WorkflowTaskBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskBulkUpdateWithDefaults

`func NewWorkflowTaskBulkUpdateWithDefaults() *WorkflowTaskBulkUpdate`

NewWorkflowTaskBulkUpdateWithDefaults instantiates a new WorkflowTaskBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *WorkflowTaskBulkUpdate) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *WorkflowTaskBulkUpdate) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *WorkflowTaskBulkUpdate) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *WorkflowTaskBulkUpdate) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetAssigneeId

`func (o *WorkflowTaskBulkUpdate) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *WorkflowTaskBulkUpdate) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *WorkflowTaskBulkUpdate) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *WorkflowTaskBulkUpdate) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetFields

`func (o *WorkflowTaskBulkUpdate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowTaskBulkUpdate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowTaskBulkUpdate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowTaskBulkUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetScheduledOn

`func (o *WorkflowTaskBulkUpdate) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *WorkflowTaskBulkUpdate) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *WorkflowTaskBulkUpdate) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *WorkflowTaskBulkUpdate) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### GetWorkflowTaskId

`func (o *WorkflowTaskBulkUpdate) GetWorkflowTaskId() string`

GetWorkflowTaskId returns the WorkflowTaskId field if non-nil, zero value otherwise.

### GetWorkflowTaskIdOk

`func (o *WorkflowTaskBulkUpdate) GetWorkflowTaskIdOk() (*string, bool)`

GetWorkflowTaskIdOk returns a tuple with the WorkflowTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskId

`func (o *WorkflowTaskBulkUpdate) SetWorkflowTaskId(v string)`

SetWorkflowTaskId sets WorkflowTaskId field to given value.

### HasWorkflowTaskId

`func (o *WorkflowTaskBulkUpdate) HasWorkflowTaskId() bool`

HasWorkflowTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


