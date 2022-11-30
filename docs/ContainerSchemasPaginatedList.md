# ContainerSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerSchemas** | Pointer to [**[]ContainerSchema**](ContainerSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerSchemasPaginatedList

`func NewContainerSchemasPaginatedList() *ContainerSchemasPaginatedList`

NewContainerSchemasPaginatedList instantiates a new ContainerSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerSchemasPaginatedListWithDefaults

`func NewContainerSchemasPaginatedListWithDefaults() *ContainerSchemasPaginatedList`

NewContainerSchemasPaginatedListWithDefaults instantiates a new ContainerSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerSchemas

`func (o *ContainerSchemasPaginatedList) GetContainerSchemas() []ContainerSchema`

GetContainerSchemas returns the ContainerSchemas field if non-nil, zero value otherwise.

### GetContainerSchemasOk

`func (o *ContainerSchemasPaginatedList) GetContainerSchemasOk() (*[]ContainerSchema, bool)`

GetContainerSchemasOk returns a tuple with the ContainerSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSchemas

`func (o *ContainerSchemasPaginatedList) SetContainerSchemas(v []ContainerSchema)`

SetContainerSchemas sets ContainerSchemas field to given value.

### HasContainerSchemas

`func (o *ContainerSchemasPaginatedList) HasContainerSchemas() bool`

HasContainerSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *ContainerSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ContainerSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ContainerSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ContainerSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


