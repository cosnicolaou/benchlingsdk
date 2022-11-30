# BoxSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxSchemas** | Pointer to [**[]BoxSchema**](BoxSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewBoxSchemasPaginatedList

`func NewBoxSchemasPaginatedList() *BoxSchemasPaginatedList`

NewBoxSchemasPaginatedList instantiates a new BoxSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxSchemasPaginatedListWithDefaults

`func NewBoxSchemasPaginatedListWithDefaults() *BoxSchemasPaginatedList`

NewBoxSchemasPaginatedListWithDefaults instantiates a new BoxSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxSchemas

`func (o *BoxSchemasPaginatedList) GetBoxSchemas() []BoxSchema`

GetBoxSchemas returns the BoxSchemas field if non-nil, zero value otherwise.

### GetBoxSchemasOk

`func (o *BoxSchemasPaginatedList) GetBoxSchemasOk() (*[]BoxSchema, bool)`

GetBoxSchemasOk returns a tuple with the BoxSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxSchemas

`func (o *BoxSchemasPaginatedList) SetBoxSchemas(v []BoxSchema)`

SetBoxSchemas sets BoxSchemas field to given value.

### HasBoxSchemas

`func (o *BoxSchemasPaginatedList) HasBoxSchemas() bool`

HasBoxSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *BoxSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BoxSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BoxSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BoxSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


