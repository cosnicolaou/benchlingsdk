# WorkflowSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | ID of the batch | [optional] 
**ContainerIds** | Pointer to **[]string** | Array of IDs of containers | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the sample was created | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the sample | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the sample | [optional] 

## Methods

### NewWorkflowSample

`func NewWorkflowSample() *WorkflowSample`

NewWorkflowSample instantiates a new WorkflowSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSampleWithDefaults

`func NewWorkflowSampleWithDefaults() *WorkflowSample`

NewWorkflowSampleWithDefaults instantiates a new WorkflowSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *WorkflowSample) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *WorkflowSample) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *WorkflowSample) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *WorkflowSample) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetContainerIds

`func (o *WorkflowSample) GetContainerIds() []string`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *WorkflowSample) GetContainerIdsOk() (*[]string, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *WorkflowSample) SetContainerIds(v []string)`

SetContainerIds sets ContainerIds field to given value.

### HasContainerIds

`func (o *WorkflowSample) HasContainerIds() bool`

HasContainerIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkflowSample) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowSample) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowSample) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowSample) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *WorkflowSample) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSample) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSample) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowSample) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSample) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSample) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSample) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


