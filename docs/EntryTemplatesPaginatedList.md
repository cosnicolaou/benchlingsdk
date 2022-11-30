# EntryTemplatesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryTemplates** | Pointer to [**[]EntryTemplate**](EntryTemplate.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewEntryTemplatesPaginatedList

`func NewEntryTemplatesPaginatedList() *EntryTemplatesPaginatedList`

NewEntryTemplatesPaginatedList instantiates a new EntryTemplatesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTemplatesPaginatedListWithDefaults

`func NewEntryTemplatesPaginatedListWithDefaults() *EntryTemplatesPaginatedList`

NewEntryTemplatesPaginatedListWithDefaults instantiates a new EntryTemplatesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryTemplates

`func (o *EntryTemplatesPaginatedList) GetEntryTemplates() []EntryTemplate`

GetEntryTemplates returns the EntryTemplates field if non-nil, zero value otherwise.

### GetEntryTemplatesOk

`func (o *EntryTemplatesPaginatedList) GetEntryTemplatesOk() (*[]EntryTemplate, bool)`

GetEntryTemplatesOk returns a tuple with the EntryTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTemplates

`func (o *EntryTemplatesPaginatedList) SetEntryTemplates(v []EntryTemplate)`

SetEntryTemplates sets EntryTemplates field to given value.

### HasEntryTemplates

`func (o *EntryTemplatesPaginatedList) HasEntryTemplates() bool`

HasEntryTemplates returns a boolean if a field has been set.

### GetNextToken

`func (o *EntryTemplatesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *EntryTemplatesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *EntryTemplatesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *EntryTemplatesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


