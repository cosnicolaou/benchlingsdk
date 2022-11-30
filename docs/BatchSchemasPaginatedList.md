# BatchSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchSchemas** | Pointer to [**[]BatchSchema**](BatchSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchSchemasPaginatedList

`func NewBatchSchemasPaginatedList() *BatchSchemasPaginatedList`

NewBatchSchemasPaginatedList instantiates a new BatchSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSchemasPaginatedListWithDefaults

`func NewBatchSchemasPaginatedListWithDefaults() *BatchSchemasPaginatedList`

NewBatchSchemasPaginatedListWithDefaults instantiates a new BatchSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchSchemas

`func (o *BatchSchemasPaginatedList) GetBatchSchemas() []BatchSchema`

GetBatchSchemas returns the BatchSchemas field if non-nil, zero value otherwise.

### GetBatchSchemasOk

`func (o *BatchSchemasPaginatedList) GetBatchSchemasOk() (*[]BatchSchema, bool)`

GetBatchSchemasOk returns a tuple with the BatchSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSchemas

`func (o *BatchSchemasPaginatedList) SetBatchSchemas(v []BatchSchema)`

SetBatchSchemas sets BatchSchemas field to given value.

### HasBatchSchemas

`func (o *BatchSchemasPaginatedList) HasBatchSchemas() bool`

HasBatchSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *BatchSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BatchSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BatchSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BatchSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


