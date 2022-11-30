# EntitySchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitySchemas** | Pointer to [**[]EntitySchema**](EntitySchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewEntitySchemasPaginatedList

`func NewEntitySchemasPaginatedList() *EntitySchemasPaginatedList`

NewEntitySchemasPaginatedList instantiates a new EntitySchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemasPaginatedListWithDefaults

`func NewEntitySchemasPaginatedListWithDefaults() *EntitySchemasPaginatedList`

NewEntitySchemasPaginatedListWithDefaults instantiates a new EntitySchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitySchemas

`func (o *EntitySchemasPaginatedList) GetEntitySchemas() []EntitySchema`

GetEntitySchemas returns the EntitySchemas field if non-nil, zero value otherwise.

### GetEntitySchemasOk

`func (o *EntitySchemasPaginatedList) GetEntitySchemasOk() (*[]EntitySchema, bool)`

GetEntitySchemasOk returns a tuple with the EntitySchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySchemas

`func (o *EntitySchemasPaginatedList) SetEntitySchemas(v []EntitySchema)`

SetEntitySchemas sets EntitySchemas field to given value.

### HasEntitySchemas

`func (o *EntitySchemasPaginatedList) HasEntitySchemas() bool`

HasEntitySchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *EntitySchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *EntitySchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *EntitySchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *EntitySchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


