# ContainerBulkUpdateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentStorageId** | Pointer to **string** | ID of containing parent storage, can also specify a coordinate for plates and boxes (e.g. plt_2bAks9dx:a2). | [optional] 
**ContainerId** | **string** |  | 
**Quantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**Volume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) |  | [optional] 

## Methods

### NewContainerBulkUpdateItem

`func NewContainerBulkUpdateItem(containerId string, ) *ContainerBulkUpdateItem`

NewContainerBulkUpdateItem instantiates a new ContainerBulkUpdateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerBulkUpdateItemWithDefaults

`func NewContainerBulkUpdateItemWithDefaults() *ContainerBulkUpdateItem`

NewContainerBulkUpdateItemWithDefaults instantiates a new ContainerBulkUpdateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ContainerBulkUpdateItem) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContainerBulkUpdateItem) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContainerBulkUpdateItem) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ContainerBulkUpdateItem) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *ContainerBulkUpdateItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerBulkUpdateItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerBulkUpdateItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerBulkUpdateItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *ContainerBulkUpdateItem) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *ContainerBulkUpdateItem) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *ContainerBulkUpdateItem) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *ContainerBulkUpdateItem) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### GetContainerId

`func (o *ContainerBulkUpdateItem) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ContainerBulkUpdateItem) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ContainerBulkUpdateItem) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetQuantity

`func (o *ContainerBulkUpdateItem) GetQuantity() ContainerQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ContainerBulkUpdateItem) GetQuantityOk() (*ContainerQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ContainerBulkUpdateItem) SetQuantity(v ContainerQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ContainerBulkUpdateItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetVolume

`func (o *ContainerBulkUpdateItem) GetVolume() DeprecatedContainerVolumeForInput`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ContainerBulkUpdateItem) GetVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ContainerBulkUpdateItem) SetVolume(v DeprecatedContainerVolumeForInput)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ContainerBulkUpdateItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


