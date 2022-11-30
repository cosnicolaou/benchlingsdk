# EntriesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]Entry**](Entry.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewEntriesPaginatedList

`func NewEntriesPaginatedList() *EntriesPaginatedList`

NewEntriesPaginatedList instantiates a new EntriesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntriesPaginatedListWithDefaults

`func NewEntriesPaginatedListWithDefaults() *EntriesPaginatedList`

NewEntriesPaginatedListWithDefaults instantiates a new EntriesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *EntriesPaginatedList) GetEntries() []Entry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *EntriesPaginatedList) GetEntriesOk() (*[]Entry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *EntriesPaginatedList) SetEntries(v []Entry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *EntriesPaginatedList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetNextToken

`func (o *EntriesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *EntriesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *EntriesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *EntriesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


