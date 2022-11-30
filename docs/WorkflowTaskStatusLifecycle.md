# WorkflowTaskStatusLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**InitialStatus** | Pointer to [**WorkflowTaskStatus**](WorkflowTaskStatus.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to [**[]WorkflowTaskStatus**](WorkflowTaskStatus.md) |  | [optional] 
**Transitions** | Pointer to [**[]WorkflowTaskStatusLifecycleTransition**](WorkflowTaskStatusLifecycleTransition.md) |  | [optional] 

## Methods

### NewWorkflowTaskStatusLifecycle

`func NewWorkflowTaskStatusLifecycle() *WorkflowTaskStatusLifecycle`

NewWorkflowTaskStatusLifecycle instantiates a new WorkflowTaskStatusLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskStatusLifecycleWithDefaults

`func NewWorkflowTaskStatusLifecycleWithDefaults() *WorkflowTaskStatusLifecycle`

NewWorkflowTaskStatusLifecycleWithDefaults instantiates a new WorkflowTaskStatusLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTaskStatusLifecycle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskStatusLifecycle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskStatusLifecycle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskStatusLifecycle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialStatus

`func (o *WorkflowTaskStatusLifecycle) GetInitialStatus() WorkflowTaskStatus`

GetInitialStatus returns the InitialStatus field if non-nil, zero value otherwise.

### GetInitialStatusOk

`func (o *WorkflowTaskStatusLifecycle) GetInitialStatusOk() (*WorkflowTaskStatus, bool)`

GetInitialStatusOk returns a tuple with the InitialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStatus

`func (o *WorkflowTaskStatusLifecycle) SetInitialStatus(v WorkflowTaskStatus)`

SetInitialStatus sets InitialStatus field to given value.

### HasInitialStatus

`func (o *WorkflowTaskStatusLifecycle) HasInitialStatus() bool`

HasInitialStatus returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskStatusLifecycle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskStatusLifecycle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskStatusLifecycle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskStatusLifecycle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowTaskStatusLifecycle) GetStatuses() []WorkflowTaskStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowTaskStatusLifecycle) GetStatusesOk() (*[]WorkflowTaskStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowTaskStatusLifecycle) SetStatuses(v []WorkflowTaskStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowTaskStatusLifecycle) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTransitions

`func (o *WorkflowTaskStatusLifecycle) GetTransitions() []WorkflowTaskStatusLifecycleTransition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowTaskStatusLifecycle) GetTransitionsOk() (*[]WorkflowTaskStatusLifecycleTransition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowTaskStatusLifecycle) SetTransitions(v []WorkflowTaskStatusLifecycleTransition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *WorkflowTaskStatusLifecycle) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


