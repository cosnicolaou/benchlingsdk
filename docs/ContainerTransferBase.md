# ContainerTransferBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceBatchId** | Pointer to **string** | ID of the batch that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**SourceContainerId** | Pointer to **string** | ID of the container that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**SourceEntityId** | Pointer to **string** | ID of the entity that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**TransferQuantity** | Pointer to [**ContainerTransferBaseTransferQuantity**](ContainerTransferBaseTransferQuantity.md) |  | [optional] 
**TransferVolume** | Pointer to [**ContainerTransferBaseTransferVolume**](ContainerTransferBaseTransferVolume.md) |  | [optional] 

## Methods

### NewContainerTransferBase

`func NewContainerTransferBase() *ContainerTransferBase`

NewContainerTransferBase instantiates a new ContainerTransferBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTransferBaseWithDefaults

`func NewContainerTransferBaseWithDefaults() *ContainerTransferBase`

NewContainerTransferBaseWithDefaults instantiates a new ContainerTransferBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBatchId

`func (o *ContainerTransferBase) GetSourceBatchId() string`

GetSourceBatchId returns the SourceBatchId field if non-nil, zero value otherwise.

### GetSourceBatchIdOk

`func (o *ContainerTransferBase) GetSourceBatchIdOk() (*string, bool)`

GetSourceBatchIdOk returns a tuple with the SourceBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBatchId

`func (o *ContainerTransferBase) SetSourceBatchId(v string)`

SetSourceBatchId sets SourceBatchId field to given value.

### HasSourceBatchId

`func (o *ContainerTransferBase) HasSourceBatchId() bool`

HasSourceBatchId returns a boolean if a field has been set.

### GetSourceContainerId

`func (o *ContainerTransferBase) GetSourceContainerId() string`

GetSourceContainerId returns the SourceContainerId field if non-nil, zero value otherwise.

### GetSourceContainerIdOk

`func (o *ContainerTransferBase) GetSourceContainerIdOk() (*string, bool)`

GetSourceContainerIdOk returns a tuple with the SourceContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContainerId

`func (o *ContainerTransferBase) SetSourceContainerId(v string)`

SetSourceContainerId sets SourceContainerId field to given value.

### HasSourceContainerId

`func (o *ContainerTransferBase) HasSourceContainerId() bool`

HasSourceContainerId returns a boolean if a field has been set.

### GetSourceEntityId

`func (o *ContainerTransferBase) GetSourceEntityId() string`

GetSourceEntityId returns the SourceEntityId field if non-nil, zero value otherwise.

### GetSourceEntityIdOk

`func (o *ContainerTransferBase) GetSourceEntityIdOk() (*string, bool)`

GetSourceEntityIdOk returns a tuple with the SourceEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityId

`func (o *ContainerTransferBase) SetSourceEntityId(v string)`

SetSourceEntityId sets SourceEntityId field to given value.

### HasSourceEntityId

`func (o *ContainerTransferBase) HasSourceEntityId() bool`

HasSourceEntityId returns a boolean if a field has been set.

### GetTransferQuantity

`func (o *ContainerTransferBase) GetTransferQuantity() ContainerTransferBaseTransferQuantity`

GetTransferQuantity returns the TransferQuantity field if non-nil, zero value otherwise.

### GetTransferQuantityOk

`func (o *ContainerTransferBase) GetTransferQuantityOk() (*ContainerTransferBaseTransferQuantity, bool)`

GetTransferQuantityOk returns a tuple with the TransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferQuantity

`func (o *ContainerTransferBase) SetTransferQuantity(v ContainerTransferBaseTransferQuantity)`

SetTransferQuantity sets TransferQuantity field to given value.

### HasTransferQuantity

`func (o *ContainerTransferBase) HasTransferQuantity() bool`

HasTransferQuantity returns a boolean if a field has been set.

### GetTransferVolume

`func (o *ContainerTransferBase) GetTransferVolume() ContainerTransferBaseTransferVolume`

GetTransferVolume returns the TransferVolume field if non-nil, zero value otherwise.

### GetTransferVolumeOk

`func (o *ContainerTransferBase) GetTransferVolumeOk() (*ContainerTransferBaseTransferVolume, bool)`

GetTransferVolumeOk returns a tuple with the TransferVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferVolume

`func (o *ContainerTransferBase) SetTransferVolume(v ContainerTransferBaseTransferVolume)`

SetTransferVolume sets TransferVolume field to given value.

### HasTransferVolume

`func (o *ContainerTransferBase) HasTransferVolume() bool`

HasTransferVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


