# AssayResultSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayResultSchemas** | Pointer to [**[]AssayResultSchema**](AssayResultSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayResultSchemasPaginatedList

`func NewAssayResultSchemasPaginatedList() *AssayResultSchemasPaginatedList`

NewAssayResultSchemasPaginatedList instantiates a new AssayResultSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayResultSchemasPaginatedListWithDefaults

`func NewAssayResultSchemasPaginatedListWithDefaults() *AssayResultSchemasPaginatedList`

NewAssayResultSchemasPaginatedListWithDefaults instantiates a new AssayResultSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayResultSchemas

`func (o *AssayResultSchemasPaginatedList) GetAssayResultSchemas() []AssayResultSchema`

GetAssayResultSchemas returns the AssayResultSchemas field if non-nil, zero value otherwise.

### GetAssayResultSchemasOk

`func (o *AssayResultSchemasPaginatedList) GetAssayResultSchemasOk() (*[]AssayResultSchema, bool)`

GetAssayResultSchemasOk returns a tuple with the AssayResultSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayResultSchemas

`func (o *AssayResultSchemasPaginatedList) SetAssayResultSchemas(v []AssayResultSchema)`

SetAssayResultSchemas sets AssayResultSchemas field to given value.

### HasAssayResultSchemas

`func (o *AssayResultSchemasPaginatedList) HasAssayResultSchemas() bool`

HasAssayResultSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *AssayResultSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AssayResultSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AssayResultSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AssayResultSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


