# AssayRunSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRunSchemas** | Pointer to [**[]AssayRunSchema**](AssayRunSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayRunSchemasPaginatedList

`func NewAssayRunSchemasPaginatedList() *AssayRunSchemasPaginatedList`

NewAssayRunSchemasPaginatedList instantiates a new AssayRunSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunSchemasPaginatedListWithDefaults

`func NewAssayRunSchemasPaginatedListWithDefaults() *AssayRunSchemasPaginatedList`

NewAssayRunSchemasPaginatedListWithDefaults instantiates a new AssayRunSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRunSchemas

`func (o *AssayRunSchemasPaginatedList) GetAssayRunSchemas() []AssayRunSchema`

GetAssayRunSchemas returns the AssayRunSchemas field if non-nil, zero value otherwise.

### GetAssayRunSchemasOk

`func (o *AssayRunSchemasPaginatedList) GetAssayRunSchemasOk() (*[]AssayRunSchema, bool)`

GetAssayRunSchemasOk returns a tuple with the AssayRunSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunSchemas

`func (o *AssayRunSchemasPaginatedList) SetAssayRunSchemas(v []AssayRunSchema)`

SetAssayRunSchemas sets AssayRunSchemas field to given value.

### HasAssayRunSchemas

`func (o *AssayRunSchemasPaginatedList) HasAssayRunSchemas() bool`

HasAssayRunSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *AssayRunSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AssayRunSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AssayRunSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AssayRunSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


