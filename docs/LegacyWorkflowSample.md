# LegacyWorkflowSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | ID of the batch | [optional] 
**ContainerIds** | Pointer to **[]string** | Array of IDs of containers | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the sample was created | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the sample | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the sample | [optional] 

## Methods

### NewLegacyWorkflowSample

`func NewLegacyWorkflowSample() *LegacyWorkflowSample`

NewLegacyWorkflowSample instantiates a new LegacyWorkflowSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWorkflowSampleWithDefaults

`func NewLegacyWorkflowSampleWithDefaults() *LegacyWorkflowSample`

NewLegacyWorkflowSampleWithDefaults instantiates a new LegacyWorkflowSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *LegacyWorkflowSample) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *LegacyWorkflowSample) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *LegacyWorkflowSample) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *LegacyWorkflowSample) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetContainerIds

`func (o *LegacyWorkflowSample) GetContainerIds() []string`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *LegacyWorkflowSample) GetContainerIdsOk() (*[]string, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *LegacyWorkflowSample) SetContainerIds(v []string)`

SetContainerIds sets ContainerIds field to given value.

### HasContainerIds

`func (o *LegacyWorkflowSample) HasContainerIds() bool`

HasContainerIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LegacyWorkflowSample) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacyWorkflowSample) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacyWorkflowSample) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacyWorkflowSample) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LegacyWorkflowSample) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyWorkflowSample) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyWorkflowSample) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyWorkflowSample) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LegacyWorkflowSample) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyWorkflowSample) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyWorkflowSample) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyWorkflowSample) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


