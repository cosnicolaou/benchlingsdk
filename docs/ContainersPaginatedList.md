# ContainersPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]Container**](Container.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewContainersPaginatedList

`func NewContainersPaginatedList() *ContainersPaginatedList`

NewContainersPaginatedList instantiates a new ContainersPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersPaginatedListWithDefaults

`func NewContainersPaginatedListWithDefaults() *ContainersPaginatedList`

NewContainersPaginatedListWithDefaults instantiates a new ContainersPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *ContainersPaginatedList) GetContainers() []Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ContainersPaginatedList) GetContainersOk() (*[]Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ContainersPaginatedList) SetContainers(v []Container)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ContainersPaginatedList) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetNextToken

`func (o *ContainersPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ContainersPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ContainersPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ContainersPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


