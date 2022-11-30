# DropdownSummariesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dropdowns** | Pointer to [**[]DropdownSummary**](DropdownSummary.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewDropdownSummariesPaginatedList

`func NewDropdownSummariesPaginatedList() *DropdownSummariesPaginatedList`

NewDropdownSummariesPaginatedList instantiates a new DropdownSummariesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropdownSummariesPaginatedListWithDefaults

`func NewDropdownSummariesPaginatedListWithDefaults() *DropdownSummariesPaginatedList`

NewDropdownSummariesPaginatedListWithDefaults instantiates a new DropdownSummariesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropdowns

`func (o *DropdownSummariesPaginatedList) GetDropdowns() []DropdownSummary`

GetDropdowns returns the Dropdowns field if non-nil, zero value otherwise.

### GetDropdownsOk

`func (o *DropdownSummariesPaginatedList) GetDropdownsOk() (*[]DropdownSummary, bool)`

GetDropdownsOk returns a tuple with the Dropdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropdowns

`func (o *DropdownSummariesPaginatedList) SetDropdowns(v []DropdownSummary)`

SetDropdowns sets Dropdowns field to given value.

### HasDropdowns

`func (o *DropdownSummariesPaginatedList) HasDropdowns() bool`

HasDropdowns returns a boolean if a field has been set.

### GetNextToken

`func (o *DropdownSummariesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *DropdownSummariesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *DropdownSummariesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *DropdownSummariesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


