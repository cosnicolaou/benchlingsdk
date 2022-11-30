# ContainersBulkUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]ContainerBulkUpdateItem**](ContainerBulkUpdateItem.md) |  | 

## Methods

### NewContainersBulkUpdateRequest

`func NewContainersBulkUpdateRequest(containers []ContainerBulkUpdateItem, ) *ContainersBulkUpdateRequest`

NewContainersBulkUpdateRequest instantiates a new ContainersBulkUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersBulkUpdateRequestWithDefaults

`func NewContainersBulkUpdateRequestWithDefaults() *ContainersBulkUpdateRequest`

NewContainersBulkUpdateRequestWithDefaults instantiates a new ContainersBulkUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *ContainersBulkUpdateRequest) GetContainers() []ContainerBulkUpdateItem`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ContainersBulkUpdateRequest) GetContainersOk() (*[]ContainerBulkUpdateItem, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ContainersBulkUpdateRequest) SetContainers(v []ContainerBulkUpdateItem)`

SetContainers sets Containers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


