# WorkflowTaskGroupWriteBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | Pointer to **string** | ID of the folder that contains the workflow task group | [optional] 
**Name** | Pointer to **string** | The name of the workflow task group | [optional] 
**WatcherIds** | Pointer to **[]string** | IDs of the users watching the workflow task group | [optional] 

## Methods

### NewWorkflowTaskGroupWriteBase

`func NewWorkflowTaskGroupWriteBase() *WorkflowTaskGroupWriteBase`

NewWorkflowTaskGroupWriteBase instantiates a new WorkflowTaskGroupWriteBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupWriteBaseWithDefaults

`func NewWorkflowTaskGroupWriteBaseWithDefaults() *WorkflowTaskGroupWriteBase`

NewWorkflowTaskGroupWriteBaseWithDefaults instantiates a new WorkflowTaskGroupWriteBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *WorkflowTaskGroupWriteBase) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *WorkflowTaskGroupWriteBase) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *WorkflowTaskGroupWriteBase) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *WorkflowTaskGroupWriteBase) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskGroupWriteBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskGroupWriteBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskGroupWriteBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskGroupWriteBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWatcherIds

`func (o *WorkflowTaskGroupWriteBase) GetWatcherIds() []string`

GetWatcherIds returns the WatcherIds field if non-nil, zero value otherwise.

### GetWatcherIdsOk

`func (o *WorkflowTaskGroupWriteBase) GetWatcherIdsOk() (*[]string, bool)`

GetWatcherIdsOk returns a tuple with the WatcherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcherIds

`func (o *WorkflowTaskGroupWriteBase) SetWatcherIds(v []string)`

SetWatcherIds sets WatcherIds field to given value.

### HasWatcherIds

`func (o *WorkflowTaskGroupWriteBase) HasWatcherIds() bool`

HasWatcherIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


