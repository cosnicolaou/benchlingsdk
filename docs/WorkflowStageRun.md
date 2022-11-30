# WorkflowStageRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the stage run was created | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the stage run | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the stage run | [optional] 
**Status** | Pointer to **string** | Status of the stage run | [optional] 

## Methods

### NewWorkflowStageRun

`func NewWorkflowStageRun() *WorkflowStageRun`

NewWorkflowStageRun instantiates a new WorkflowStageRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStageRunWithDefaults

`func NewWorkflowStageRunWithDefaults() *WorkflowStageRun`

NewWorkflowStageRunWithDefaults instantiates a new WorkflowStageRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkflowStageRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowStageRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowStageRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowStageRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *WorkflowStageRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStageRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStageRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowStageRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowStageRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStageRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStageRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowStageRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowStageRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowStageRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowStageRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowStageRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


