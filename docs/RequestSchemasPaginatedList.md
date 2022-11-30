# RequestSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**RequestSchemas** | Pointer to [**[]RequestSchema**](RequestSchema.md) |  | [optional] [readonly] 

## Methods

### NewRequestSchemasPaginatedList

`func NewRequestSchemasPaginatedList() *RequestSchemasPaginatedList`

NewRequestSchemasPaginatedList instantiates a new RequestSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSchemasPaginatedListWithDefaults

`func NewRequestSchemasPaginatedListWithDefaults() *RequestSchemasPaginatedList`

NewRequestSchemasPaginatedListWithDefaults instantiates a new RequestSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *RequestSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RequestSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RequestSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RequestSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRequestSchemas

`func (o *RequestSchemasPaginatedList) GetRequestSchemas() []RequestSchema`

GetRequestSchemas returns the RequestSchemas field if non-nil, zero value otherwise.

### GetRequestSchemasOk

`func (o *RequestSchemasPaginatedList) GetRequestSchemasOk() (*[]RequestSchema, bool)`

GetRequestSchemasOk returns a tuple with the RequestSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSchemas

`func (o *RequestSchemasPaginatedList) SetRequestSchemas(v []RequestSchema)`

SetRequestSchemas sets RequestSchemas field to given value.

### HasRequestSchemas

`func (o *RequestSchemasPaginatedList) HasRequestSchemas() bool`

HasRequestSchemas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


