# WorkflowTaskWriteBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeId** | Pointer to **string** | The id of the user assigned to the task | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**ScheduledOn** | Pointer to **string** | The date on which the task is scheduled to be executed | [optional] 

## Methods

### NewWorkflowTaskWriteBase

`func NewWorkflowTaskWriteBase() *WorkflowTaskWriteBase`

NewWorkflowTaskWriteBase instantiates a new WorkflowTaskWriteBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskWriteBaseWithDefaults

`func NewWorkflowTaskWriteBaseWithDefaults() *WorkflowTaskWriteBase`

NewWorkflowTaskWriteBaseWithDefaults instantiates a new WorkflowTaskWriteBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeId

`func (o *WorkflowTaskWriteBase) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *WorkflowTaskWriteBase) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *WorkflowTaskWriteBase) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *WorkflowTaskWriteBase) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetFields

`func (o *WorkflowTaskWriteBase) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WorkflowTaskWriteBase) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WorkflowTaskWriteBase) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WorkflowTaskWriteBase) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetScheduledOn

`func (o *WorkflowTaskWriteBase) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *WorkflowTaskWriteBase) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *WorkflowTaskWriteBase) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *WorkflowTaskWriteBase) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


