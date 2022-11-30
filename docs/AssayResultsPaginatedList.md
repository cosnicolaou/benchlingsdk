# AssayResultsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayResults** | Pointer to [**[]AssayResult**](AssayResult.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayResultsPaginatedList

`func NewAssayResultsPaginatedList() *AssayResultsPaginatedList`

NewAssayResultsPaginatedList instantiates a new AssayResultsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayResultsPaginatedListWithDefaults

`func NewAssayResultsPaginatedListWithDefaults() *AssayResultsPaginatedList`

NewAssayResultsPaginatedListWithDefaults instantiates a new AssayResultsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayResults

`func (o *AssayResultsPaginatedList) GetAssayResults() []AssayResult`

GetAssayResults returns the AssayResults field if non-nil, zero value otherwise.

### GetAssayResultsOk

`func (o *AssayResultsPaginatedList) GetAssayResultsOk() (*[]AssayResult, bool)`

GetAssayResultsOk returns a tuple with the AssayResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayResults

`func (o *AssayResultsPaginatedList) SetAssayResults(v []AssayResult)`

SetAssayResults sets AssayResults field to given value.

### HasAssayResults

`func (o *AssayResultsPaginatedList) HasAssayResults() bool`

HasAssayResults returns a boolean if a field has been set.

### GetNextToken

`func (o *AssayResultsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AssayResultsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AssayResultsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AssayResultsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


