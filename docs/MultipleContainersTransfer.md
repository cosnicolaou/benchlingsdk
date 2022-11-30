# MultipleContainersTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceBatchId** | Pointer to **string** | ID of the batch that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**SourceContainerId** | Pointer to **string** | ID of the container that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**SourceEntityId** | Pointer to **string** | ID of the entity that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId.  | [optional] 
**TransferQuantity** | Pointer to [**ContainerTransferBaseTransferQuantity**](ContainerTransferBaseTransferQuantity.md) |  | [optional] 
**TransferVolume** | Pointer to [**ContainerTransferBaseTransferVolume**](ContainerTransferBaseTransferVolume.md) |  | [optional] 
**DestinationContainerId** | **string** | ID of container that will be transferred into. | 
**FinalQuantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**FinalVolume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) |  | [optional] 
**SourceConcentration** | Pointer to **map[string]interface{}** | Concentration at which to transfer entities or batches. Not applicable for container-to-container transfers (the concentration of the source container’s contents will be used).  | [optional] 

## Methods

### NewMultipleContainersTransfer

`func NewMultipleContainersTransfer(destinationContainerId string, ) *MultipleContainersTransfer`

NewMultipleContainersTransfer instantiates a new MultipleContainersTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleContainersTransferWithDefaults

`func NewMultipleContainersTransferWithDefaults() *MultipleContainersTransfer`

NewMultipleContainersTransferWithDefaults instantiates a new MultipleContainersTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBatchId

`func (o *MultipleContainersTransfer) GetSourceBatchId() string`

GetSourceBatchId returns the SourceBatchId field if non-nil, zero value otherwise.

### GetSourceBatchIdOk

`func (o *MultipleContainersTransfer) GetSourceBatchIdOk() (*string, bool)`

GetSourceBatchIdOk returns a tuple with the SourceBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBatchId

`func (o *MultipleContainersTransfer) SetSourceBatchId(v string)`

SetSourceBatchId sets SourceBatchId field to given value.

### HasSourceBatchId

`func (o *MultipleContainersTransfer) HasSourceBatchId() bool`

HasSourceBatchId returns a boolean if a field has been set.

### GetSourceContainerId

`func (o *MultipleContainersTransfer) GetSourceContainerId() string`

GetSourceContainerId returns the SourceContainerId field if non-nil, zero value otherwise.

### GetSourceContainerIdOk

`func (o *MultipleContainersTransfer) GetSourceContainerIdOk() (*string, bool)`

GetSourceContainerIdOk returns a tuple with the SourceContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContainerId

`func (o *MultipleContainersTransfer) SetSourceContainerId(v string)`

SetSourceContainerId sets SourceContainerId field to given value.

### HasSourceContainerId

`func (o *MultipleContainersTransfer) HasSourceContainerId() bool`

HasSourceContainerId returns a boolean if a field has been set.

### GetSourceEntityId

`func (o *MultipleContainersTransfer) GetSourceEntityId() string`

GetSourceEntityId returns the SourceEntityId field if non-nil, zero value otherwise.

### GetSourceEntityIdOk

`func (o *MultipleContainersTransfer) GetSourceEntityIdOk() (*string, bool)`

GetSourceEntityIdOk returns a tuple with the SourceEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityId

`func (o *MultipleContainersTransfer) SetSourceEntityId(v string)`

SetSourceEntityId sets SourceEntityId field to given value.

### HasSourceEntityId

`func (o *MultipleContainersTransfer) HasSourceEntityId() bool`

HasSourceEntityId returns a boolean if a field has been set.

### GetTransferQuantity

`func (o *MultipleContainersTransfer) GetTransferQuantity() ContainerTransferBaseTransferQuantity`

GetTransferQuantity returns the TransferQuantity field if non-nil, zero value otherwise.

### GetTransferQuantityOk

`func (o *MultipleContainersTransfer) GetTransferQuantityOk() (*ContainerTransferBaseTransferQuantity, bool)`

GetTransferQuantityOk returns a tuple with the TransferQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferQuantity

`func (o *MultipleContainersTransfer) SetTransferQuantity(v ContainerTransferBaseTransferQuantity)`

SetTransferQuantity sets TransferQuantity field to given value.

### HasTransferQuantity

`func (o *MultipleContainersTransfer) HasTransferQuantity() bool`

HasTransferQuantity returns a boolean if a field has been set.

### GetTransferVolume

`func (o *MultipleContainersTransfer) GetTransferVolume() ContainerTransferBaseTransferVolume`

GetTransferVolume returns the TransferVolume field if non-nil, zero value otherwise.

### GetTransferVolumeOk

`func (o *MultipleContainersTransfer) GetTransferVolumeOk() (*ContainerTransferBaseTransferVolume, bool)`

GetTransferVolumeOk returns a tuple with the TransferVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferVolume

`func (o *MultipleContainersTransfer) SetTransferVolume(v ContainerTransferBaseTransferVolume)`

SetTransferVolume sets TransferVolume field to given value.

### HasTransferVolume

`func (o *MultipleContainersTransfer) HasTransferVolume() bool`

HasTransferVolume returns a boolean if a field has been set.

### GetDestinationContainerId

`func (o *MultipleContainersTransfer) GetDestinationContainerId() string`

GetDestinationContainerId returns the DestinationContainerId field if non-nil, zero value otherwise.

### GetDestinationContainerIdOk

`func (o *MultipleContainersTransfer) GetDestinationContainerIdOk() (*string, bool)`

GetDestinationContainerIdOk returns a tuple with the DestinationContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContainerId

`func (o *MultipleContainersTransfer) SetDestinationContainerId(v string)`

SetDestinationContainerId sets DestinationContainerId field to given value.


### GetFinalQuantity

`func (o *MultipleContainersTransfer) GetFinalQuantity() ContainerQuantity`

GetFinalQuantity returns the FinalQuantity field if non-nil, zero value otherwise.

### GetFinalQuantityOk

`func (o *MultipleContainersTransfer) GetFinalQuantityOk() (*ContainerQuantity, bool)`

GetFinalQuantityOk returns a tuple with the FinalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalQuantity

`func (o *MultipleContainersTransfer) SetFinalQuantity(v ContainerQuantity)`

SetFinalQuantity sets FinalQuantity field to given value.

### HasFinalQuantity

`func (o *MultipleContainersTransfer) HasFinalQuantity() bool`

HasFinalQuantity returns a boolean if a field has been set.

### GetFinalVolume

`func (o *MultipleContainersTransfer) GetFinalVolume() DeprecatedContainerVolumeForInput`

GetFinalVolume returns the FinalVolume field if non-nil, zero value otherwise.

### GetFinalVolumeOk

`func (o *MultipleContainersTransfer) GetFinalVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetFinalVolumeOk returns a tuple with the FinalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalVolume

`func (o *MultipleContainersTransfer) SetFinalVolume(v DeprecatedContainerVolumeForInput)`

SetFinalVolume sets FinalVolume field to given value.

### HasFinalVolume

`func (o *MultipleContainersTransfer) HasFinalVolume() bool`

HasFinalVolume returns a boolean if a field has been set.

### GetSourceConcentration

`func (o *MultipleContainersTransfer) GetSourceConcentration() map[string]interface{}`

GetSourceConcentration returns the SourceConcentration field if non-nil, zero value otherwise.

### GetSourceConcentrationOk

`func (o *MultipleContainersTransfer) GetSourceConcentrationOk() (*map[string]interface{}, bool)`

GetSourceConcentrationOk returns a tuple with the SourceConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConcentration

`func (o *MultipleContainersTransfer) SetSourceConcentration(v map[string]interface{})`

SetSourceConcentration sets SourceConcentration field to given value.

### HasSourceConcentration

`func (o *MultipleContainersTransfer) HasSourceConcentration() bool`

HasSourceConcentration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


