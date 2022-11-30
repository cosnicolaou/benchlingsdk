# BenchlingAppsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Apps** | Pointer to [**[]BenchlingApp**](BenchlingApp.md) |  | [optional] 

## Methods

### NewBenchlingAppsPaginatedList

`func NewBenchlingAppsPaginatedList() *BenchlingAppsPaginatedList`

NewBenchlingAppsPaginatedList instantiates a new BenchlingAppsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchlingAppsPaginatedListWithDefaults

`func NewBenchlingAppsPaginatedListWithDefaults() *BenchlingAppsPaginatedList`

NewBenchlingAppsPaginatedListWithDefaults instantiates a new BenchlingAppsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *BenchlingAppsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BenchlingAppsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BenchlingAppsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BenchlingAppsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetApps

`func (o *BenchlingAppsPaginatedList) GetApps() []BenchlingApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BenchlingAppsPaginatedList) GetAppsOk() (*[]BenchlingApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BenchlingAppsPaginatedList) SetApps(v []BenchlingApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *BenchlingAppsPaginatedList) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


