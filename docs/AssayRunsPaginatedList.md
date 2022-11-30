# AssayRunsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRuns** | Pointer to [**[]AssayRun**](AssayRun.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayRunsPaginatedList

`func NewAssayRunsPaginatedList() *AssayRunsPaginatedList`

NewAssayRunsPaginatedList instantiates a new AssayRunsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunsPaginatedListWithDefaults

`func NewAssayRunsPaginatedListWithDefaults() *AssayRunsPaginatedList`

NewAssayRunsPaginatedListWithDefaults instantiates a new AssayRunsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRuns

`func (o *AssayRunsPaginatedList) GetAssayRuns() []AssayRun`

GetAssayRuns returns the AssayRuns field if non-nil, zero value otherwise.

### GetAssayRunsOk

`func (o *AssayRunsPaginatedList) GetAssayRunsOk() (*[]AssayRun, bool)`

GetAssayRunsOk returns a tuple with the AssayRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRuns

`func (o *AssayRunsPaginatedList) SetAssayRuns(v []AssayRun)`

SetAssayRuns sets AssayRuns field to given value.

### HasAssayRuns

`func (o *AssayRunsPaginatedList) HasAssayRuns() bool`

HasAssayRuns returns a boolean if a field has been set.

### GetNextToken

`func (o *AssayRunsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AssayRunsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AssayRunsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AssayRunsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


