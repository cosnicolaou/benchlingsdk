# ContainerTransferAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationContents** | [**[]ContainerTransferDestinationContentsItem**](ContainerTransferDestinationContentsItem.md) | This represents what the contents of the destination container should look like post-transfer.  | 
**DestinationQuantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) | This represents the desired final quantity of the destination container, post-dilution. If you don&#39;t want to dilute the destination container, you can omit this parameter. Supports mass, volume, and other quantities.  | [optional] 
**DestinationVolume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) | Deprecated - use destinationQuantity instead.  | [optional] 

## Methods

### NewContainerTransferAllOf

`func NewContainerTransferAllOf(destinationContents []ContainerTransferDestinationContentsItem, ) *ContainerTransferAllOf`

NewContainerTransferAllOf instantiates a new ContainerTransferAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTransferAllOfWithDefaults

`func NewContainerTransferAllOfWithDefaults() *ContainerTransferAllOf`

NewContainerTransferAllOfWithDefaults instantiates a new ContainerTransferAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationContents

`func (o *ContainerTransferAllOf) GetDestinationContents() []ContainerTransferDestinationContentsItem`

GetDestinationContents returns the DestinationContents field if non-nil, zero value otherwise.

### GetDestinationContentsOk

`func (o *ContainerTransferAllOf) GetDestinationContentsOk() (*[]ContainerTransferDestinationContentsItem, bool)`

GetDestinationContentsOk returns a tuple with the DestinationContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContents

`func (o *ContainerTransferAllOf) SetDestinationContents(v []ContainerTransferDestinationContentsItem)`

SetDestinationContents sets DestinationContents field to given value.


### GetDestinationQuantity

`func (o *ContainerTransferAllOf) GetDestinationQuantity() ContainerQuantity`

GetDestinationQuantity returns the DestinationQuantity field if non-nil, zero value otherwise.

### GetDestinationQuantityOk

`func (o *ContainerTransferAllOf) GetDestinationQuantityOk() (*ContainerQuantity, bool)`

GetDestinationQuantityOk returns a tuple with the DestinationQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationQuantity

`func (o *ContainerTransferAllOf) SetDestinationQuantity(v ContainerQuantity)`

SetDestinationQuantity sets DestinationQuantity field to given value.

### HasDestinationQuantity

`func (o *ContainerTransferAllOf) HasDestinationQuantity() bool`

HasDestinationQuantity returns a boolean if a field has been set.

### GetDestinationVolume

`func (o *ContainerTransferAllOf) GetDestinationVolume() DeprecatedContainerVolumeForInput`

GetDestinationVolume returns the DestinationVolume field if non-nil, zero value otherwise.

### GetDestinationVolumeOk

`func (o *ContainerTransferAllOf) GetDestinationVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetDestinationVolumeOk returns a tuple with the DestinationVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVolume

`func (o *ContainerTransferAllOf) SetDestinationVolume(v DeprecatedContainerVolumeForInput)`

SetDestinationVolume sets DestinationVolume field to given value.

### HasDestinationVolume

`func (o *ContainerTransferAllOf) HasDestinationVolume() bool`

HasDestinationVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


