# WorkflowTaskGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **string** | ID of the folder that contains the workflow task group | 
**Name** | Pointer to **string** | The name of the workflow task group | [optional] 
**WatcherIds** | Pointer to **[]string** | IDs of the users watching the workflow task group | [optional] 
**SchemaId** | **string** | The workflow task schema of tasks in this task group | 

## Methods

### NewWorkflowTaskGroupCreate

`func NewWorkflowTaskGroupCreate(folderId string, schemaId string, ) *WorkflowTaskGroupCreate`

NewWorkflowTaskGroupCreate instantiates a new WorkflowTaskGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupCreateWithDefaults

`func NewWorkflowTaskGroupCreateWithDefaults() *WorkflowTaskGroupCreate`

NewWorkflowTaskGroupCreateWithDefaults instantiates a new WorkflowTaskGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *WorkflowTaskGroupCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *WorkflowTaskGroupCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *WorkflowTaskGroupCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetName

`func (o *WorkflowTaskGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskGroupCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskGroupCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWatcherIds

`func (o *WorkflowTaskGroupCreate) GetWatcherIds() []string`

GetWatcherIds returns the WatcherIds field if non-nil, zero value otherwise.

### GetWatcherIdsOk

`func (o *WorkflowTaskGroupCreate) GetWatcherIdsOk() (*[]string, bool)`

GetWatcherIdsOk returns a tuple with the WatcherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcherIds

`func (o *WorkflowTaskGroupCreate) SetWatcherIds(v []string)`

SetWatcherIds sets WatcherIds field to given value.

### HasWatcherIds

`func (o *WorkflowTaskGroupCreate) HasWatcherIds() bool`

HasWatcherIds returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowTaskGroupCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowTaskGroupCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowTaskGroupCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


