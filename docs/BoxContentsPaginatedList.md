# BoxContentsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]ContainerWithCoordinates**](ContainerWithCoordinates.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewBoxContentsPaginatedList

`func NewBoxContentsPaginatedList() *BoxContentsPaginatedList`

NewBoxContentsPaginatedList instantiates a new BoxContentsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxContentsPaginatedListWithDefaults

`func NewBoxContentsPaginatedListWithDefaults() *BoxContentsPaginatedList`

NewBoxContentsPaginatedListWithDefaults instantiates a new BoxContentsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *BoxContentsPaginatedList) GetContainers() []ContainerWithCoordinates`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *BoxContentsPaginatedList) GetContainersOk() (*[]ContainerWithCoordinates, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *BoxContentsPaginatedList) SetContainers(v []ContainerWithCoordinates)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *BoxContentsPaginatedList) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetNextToken

`func (o *BoxContentsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BoxContentsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BoxContentsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BoxContentsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


