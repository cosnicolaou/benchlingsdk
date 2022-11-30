# ContainerBulkUpdateItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** |  | 
**Quantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**Volume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) |  | [optional] 

## Methods

### NewContainerBulkUpdateItemAllOf

`func NewContainerBulkUpdateItemAllOf(containerId string, ) *ContainerBulkUpdateItemAllOf`

NewContainerBulkUpdateItemAllOf instantiates a new ContainerBulkUpdateItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerBulkUpdateItemAllOfWithDefaults

`func NewContainerBulkUpdateItemAllOfWithDefaults() *ContainerBulkUpdateItemAllOf`

NewContainerBulkUpdateItemAllOfWithDefaults instantiates a new ContainerBulkUpdateItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ContainerBulkUpdateItemAllOf) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ContainerBulkUpdateItemAllOf) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ContainerBulkUpdateItemAllOf) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetQuantity

`func (o *ContainerBulkUpdateItemAllOf) GetQuantity() ContainerQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ContainerBulkUpdateItemAllOf) GetQuantityOk() (*ContainerQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ContainerBulkUpdateItemAllOf) SetQuantity(v ContainerQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ContainerBulkUpdateItemAllOf) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetVolume

`func (o *ContainerBulkUpdateItemAllOf) GetVolume() DeprecatedContainerVolumeForInput`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ContainerBulkUpdateItemAllOf) GetVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ContainerBulkUpdateItemAllOf) SetVolume(v DeprecatedContainerVolumeForInput)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ContainerBulkUpdateItemAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


