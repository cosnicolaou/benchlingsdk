# OligosPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Oligos** | Pointer to [**[]Oligo**](Oligo.md) |  | [optional] 

## Methods

### NewOligosPaginatedList

`func NewOligosPaginatedList() *OligosPaginatedList`

NewOligosPaginatedList instantiates a new OligosPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOligosPaginatedListWithDefaults

`func NewOligosPaginatedListWithDefaults() *OligosPaginatedList`

NewOligosPaginatedListWithDefaults instantiates a new OligosPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *OligosPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *OligosPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *OligosPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *OligosPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetOligos

`func (o *OligosPaginatedList) GetOligos() []Oligo`

GetOligos returns the Oligos field if non-nil, zero value otherwise.

### GetOligosOk

`func (o *OligosPaginatedList) GetOligosOk() (*[]Oligo, bool)`

GetOligosOk returns a tuple with the Oligos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOligos

`func (o *OligosPaginatedList) SetOligos(v []Oligo)`

SetOligos sets Oligos field to given value.

### HasOligos

`func (o *OligosPaginatedList) HasOligos() bool`

HasOligos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


