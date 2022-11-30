# RequestTaskSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**RequestTaskSchemas** | Pointer to [**[]RequestTaskSchema**](RequestTaskSchema.md) |  | [optional] [readonly] 

## Methods

### NewRequestTaskSchemasPaginatedList

`func NewRequestTaskSchemasPaginatedList() *RequestTaskSchemasPaginatedList`

NewRequestTaskSchemasPaginatedList instantiates a new RequestTaskSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTaskSchemasPaginatedListWithDefaults

`func NewRequestTaskSchemasPaginatedListWithDefaults() *RequestTaskSchemasPaginatedList`

NewRequestTaskSchemasPaginatedListWithDefaults instantiates a new RequestTaskSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *RequestTaskSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RequestTaskSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RequestTaskSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RequestTaskSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRequestTaskSchemas

`func (o *RequestTaskSchemasPaginatedList) GetRequestTaskSchemas() []RequestTaskSchema`

GetRequestTaskSchemas returns the RequestTaskSchemas field if non-nil, zero value otherwise.

### GetRequestTaskSchemasOk

`func (o *RequestTaskSchemasPaginatedList) GetRequestTaskSchemasOk() (*[]RequestTaskSchema, bool)`

GetRequestTaskSchemasOk returns a tuple with the RequestTaskSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTaskSchemas

`func (o *RequestTaskSchemasPaginatedList) SetRequestTaskSchemas(v []RequestTaskSchema)`

SetRequestTaskSchemas sets RequestTaskSchemas field to given value.

### HasRequestTaskSchemas

`func (o *RequestTaskSchemasPaginatedList) HasRequestTaskSchemas() bool`

HasRequestTaskSchemas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


