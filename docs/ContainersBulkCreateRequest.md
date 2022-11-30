# ContainersBulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]ContainerCreate**](ContainerCreate.md) |  | 

## Methods

### NewContainersBulkCreateRequest

`func NewContainersBulkCreateRequest(containers []ContainerCreate, ) *ContainersBulkCreateRequest`

NewContainersBulkCreateRequest instantiates a new ContainersBulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersBulkCreateRequestWithDefaults

`func NewContainersBulkCreateRequestWithDefaults() *ContainersBulkCreateRequest`

NewContainersBulkCreateRequestWithDefaults instantiates a new ContainersBulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *ContainersBulkCreateRequest) GetContainers() []ContainerCreate`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ContainersBulkCreateRequest) GetContainersOk() (*[]ContainerCreate, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ContainersBulkCreateRequest) SetContainers(v []ContainerCreate)`

SetContainers sets Containers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


