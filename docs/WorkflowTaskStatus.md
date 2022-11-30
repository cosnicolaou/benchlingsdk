# WorkflowTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The status label | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the workflow task status | [optional] [readonly] 
**StatusType** | Pointer to **string** | The status type | [optional] [readonly] 

## Methods

### NewWorkflowTaskStatus

`func NewWorkflowTaskStatus() *WorkflowTaskStatus`

NewWorkflowTaskStatus instantiates a new WorkflowTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskStatusWithDefaults

`func NewWorkflowTaskStatusWithDefaults() *WorkflowTaskStatus`

NewWorkflowTaskStatusWithDefaults instantiates a new WorkflowTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowTaskStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowTaskStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowTaskStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowTaskStatus) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusType

`func (o *WorkflowTaskStatus) GetStatusType() string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *WorkflowTaskStatus) GetStatusTypeOk() (*string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *WorkflowTaskStatus) SetStatusType(v string)`

SetStatusType sets StatusType field to given value.

### HasStatusType

`func (o *WorkflowTaskStatus) HasStatusType() bool`

HasStatusType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


