# EntrySchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntrySchemas** | Pointer to [**[]EntrySchemaDetailed**](EntrySchemaDetailed.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewEntrySchemasPaginatedList

`func NewEntrySchemasPaginatedList() *EntrySchemasPaginatedList`

NewEntrySchemasPaginatedList instantiates a new EntrySchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrySchemasPaginatedListWithDefaults

`func NewEntrySchemasPaginatedListWithDefaults() *EntrySchemasPaginatedList`

NewEntrySchemasPaginatedListWithDefaults instantiates a new EntrySchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntrySchemas

`func (o *EntrySchemasPaginatedList) GetEntrySchemas() []EntrySchemaDetailed`

GetEntrySchemas returns the EntrySchemas field if non-nil, zero value otherwise.

### GetEntrySchemasOk

`func (o *EntrySchemasPaginatedList) GetEntrySchemasOk() (*[]EntrySchemaDetailed, bool)`

GetEntrySchemasOk returns a tuple with the EntrySchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySchemas

`func (o *EntrySchemasPaginatedList) SetEntrySchemas(v []EntrySchemaDetailed)`

SetEntrySchemas sets EntrySchemas field to given value.

### HasEntrySchemas

`func (o *EntrySchemasPaginatedList) HasEntrySchemas() bool`

HasEntrySchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *EntrySchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *EntrySchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *EntrySchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *EntrySchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


