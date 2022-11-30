# RnaOligosPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**RnaOligos** | Pointer to [**[]RnaOligo**](RnaOligo.md) |  | [optional] 

## Methods

### NewRnaOligosPaginatedList

`func NewRnaOligosPaginatedList() *RnaOligosPaginatedList`

NewRnaOligosPaginatedList instantiates a new RnaOligosPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaOligosPaginatedListWithDefaults

`func NewRnaOligosPaginatedListWithDefaults() *RnaOligosPaginatedList`

NewRnaOligosPaginatedListWithDefaults instantiates a new RnaOligosPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *RnaOligosPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RnaOligosPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RnaOligosPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RnaOligosPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRnaOligos

`func (o *RnaOligosPaginatedList) GetRnaOligos() []RnaOligo`

GetRnaOligos returns the RnaOligos field if non-nil, zero value otherwise.

### GetRnaOligosOk

`func (o *RnaOligosPaginatedList) GetRnaOligosOk() (*[]RnaOligo, bool)`

GetRnaOligosOk returns a tuple with the RnaOligos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaOligos

`func (o *RnaOligosPaginatedList) SetRnaOligos(v []RnaOligo)`

SetRnaOligos sets RnaOligos field to given value.

### HasRnaOligos

`func (o *RnaOligosPaginatedList) HasRnaOligos() bool`

HasRnaOligos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


