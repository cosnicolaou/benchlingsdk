# ContainerTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceBatchId** | Pointer to **string** | ID of the batch that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**SourceContainerId** | Pointer to **string** | ID of the container that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**SourceEntityId** | Pointer to **string** | ID of the entity that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**TransferQuantity** | Pointer to [**ContainerTransferBaseTransferQuantity**](ContainerTransferBaseTransferQuantity.md) |  | [optional] 
**TransferVolume** | Pointer to [**ContainerTransferBaseTransferVolume**](ContainerTransferBaseTransferVolume.md) |  | [optional] 
**DestinationContents** | [**[]ContainerTransferDestinationContentsItem**](ContainerTransferDestinationContentsItem.md) | This represents what the contents of the destination container should look like post-transfer.  | 
**DestinationQuantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) | This represents the desired final quantity of the destination container, post-dilution. If you don&#39;t want to dilute the destination container, you can omit this parameter. Supports mass, volume, and other quantities.  | [optional] 
**DestinationVolume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) | Deprecated - use destinationQuantity instead.  | [optional] 

## Methods

### NewContainerTransfer

`func NewContainerTransfer(destinationContents []ContainerTransferDestinationContentsItem, ) *ContainerTransfer`

NewContainerTransfer instantiates a new ContainerTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTransferWithDefaults

`func NewContainerTransferWithDefaults() *ContainerTransfer`

NewContainerTransferWithDefaults instantiates a new ContainerTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBatchId

`func (o *ContainerTransfer) GetSourceBatchId() string`

GetSourceBatchId returns the SourceBatchId field if non-nil, zero value otherwise.

### GetSourceBatchIdOk

`func (o *ContainerTransfer) GetSourceBatchIdOk() (*string, bool)`

GetSourceBatchIdOk returns a tuple with the SourceBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBatchId

`func (o *ContainerTransfer) SetSourceBatchId(v string)`

SetSourceBatchId sets SourceBatchId field to given value.

### HasSourceBatchId

`func (o *ContainerTransfer) HasSourceBatchId() bool`

HasSourceBatchId returns a boolean if a field has been set.

### GetSourceContainerId

`func (o *ContainerTransfer) GetSourceContainerId() string`

GetSourceContainerId returns the SourceContainerId field if non-nil, zero value otherwise.

### GetSourceContainerIdOk

`func (o *ContainerTransfer) GetSourceContainerIdOk() (*string, bool)`

GetSourceContainerIdOk returns a tuple with the SourceContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContainerId

`func (o *ContainerTransfer) SetSourceContainerId(v string)`

SetSourceContainerId sets SourceContainerId field to given value.

### HasSourceContainerId

`func (o *ContainerTransfer) HasSourceContainerId() bool`

HasSourceContainerId returns a boolean if a field has been set.

### GetSourceEntityId

`func (o *ContainerTransfer) GetSourceEntityId() string`

GetSourceEntityId returns the SourceEntityId field if non-nil, zero value otherwise.

### GetSourceEntityIdOk

`func (o *ContainerTransfer) GetSourceEntityIdOk() (*string, bool)`

GetSourceEntityIdOk returns a tuple with the SourceEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityId

`func (o *ContainerTransfer) SetSourceEntityId(v string)`

SetSourceEntityId sets SourceEntityId field to given value.

### HasSourceEntityId

`func (o *ContainerTransfer) HasSourceEntityId() bool`

HasSourceEntityId returns a boolean if a field has been set.

### GetTransferQuantity

`func (o *ContainerTransfer) GetTransferQuantity() ContainerTransferBaseTransferQuantity`

GetTransferQuantity returns the TransferQuantity field if non-nil, zero value otherwise.

### GetTransferQuantityOk

`func (o *ContainerTransfer) GetTransferQuantityOk() (*ContainerTransferBaseTransferQuantity, bool)`

GetTransferQuantityOk returns a tuple with the TransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferQuantity

`func (o *ContainerTransfer) SetTransferQuantity(v ContainerTransferBaseTransferQuantity)`

SetTransferQuantity sets TransferQuantity field to given value.

### HasTransferQuantity

`func (o *ContainerTransfer) HasTransferQuantity() bool`

HasTransferQuantity returns a boolean if a field has been set.

### GetTransferVolume

`func (o *ContainerTransfer) GetTransferVolume() ContainerTransferBaseTransferVolume`

GetTransferVolume returns the TransferVolume field if non-nil, zero value otherwise.

### GetTransferVolumeOk

`func (o *ContainerTransfer) GetTransferVolumeOk() (*ContainerTransferBaseTransferVolume, bool)`

GetTransferVolumeOk returns a tuple with the TransferVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferVolume

`func (o *ContainerTransfer) SetTransferVolume(v ContainerTransferBaseTransferVolume)`

SetTransferVolume sets TransferVolume field to given value.

### HasTransferVolume

`func (o *ContainerTransfer) HasTransferVolume() bool`

HasTransferVolume returns a boolean if a field has been set.

### GetDestinationContents

`func (o *ContainerTransfer) GetDestinationContents() []ContainerTransferDestinationContentsItem`

GetDestinationContents returns the DestinationContents field if non-nil, zero value otherwise.

### GetDestinationContentsOk

`func (o *ContainerTransfer) GetDestinationContentsOk() (*[]ContainerTransferDestinationContentsItem, bool)`

GetDestinationContentsOk returns a tuple with the DestinationContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContents

`func (o *ContainerTransfer) SetDestinationContents(v []ContainerTransferDestinationContentsItem)`

SetDestinationContents sets DestinationContents field to given value.


### GetDestinationQuantity

`func (o *ContainerTransfer) GetDestinationQuantity() ContainerQuantity`

GetDestinationQuantity returns the DestinationQuantity field if non-nil, zero value otherwise.

### GetDestinationQuantityOk

`func (o *ContainerTransfer) GetDestinationQuantityOk() (*ContainerQuantity, bool)`

GetDestinationQuantityOk returns a tuple with the DestinationQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationQuantity

`func (o *ContainerTransfer) SetDestinationQuantity(v ContainerQuantity)`

SetDestinationQuantity sets DestinationQuantity field to given value.

### HasDestinationQuantity

`func (o *ContainerTransfer) HasDestinationQuantity() bool`

HasDestinationQuantity returns a boolean if a field has been set.

### GetDestinationVolume

`func (o *ContainerTransfer) GetDestinationVolume() DeprecatedContainerVolumeForInput`

GetDestinationVolume returns the DestinationVolume field if non-nil, zero value otherwise.

### GetDestinationVolumeOk

`func (o *ContainerTransfer) GetDestinationVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetDestinationVolumeOk returns a tuple with the DestinationVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVolume

`func (o *ContainerTransfer) SetDestinationVolume(v DeprecatedContainerVolumeForInput)`

SetDestinationVolume sets DestinationVolume field to given value.

### HasDestinationVolume

`func (o *ContainerTransfer) HasDestinationVolume() bool`

HasDestinationVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


