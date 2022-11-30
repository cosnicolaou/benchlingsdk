# DnaOligosPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**DnaOligos** | Pointer to [**[]DnaOligo**](DnaOligo.md) |  | [optional] 

## Methods

### NewDnaOligosPaginatedList

`func NewDnaOligosPaginatedList() *DnaOligosPaginatedList`

NewDnaOligosPaginatedList instantiates a new DnaOligosPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaOligosPaginatedListWithDefaults

`func NewDnaOligosPaginatedListWithDefaults() *DnaOligosPaginatedList`

NewDnaOligosPaginatedListWithDefaults instantiates a new DnaOligosPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *DnaOligosPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *DnaOligosPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *DnaOligosPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *DnaOligosPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetDnaOligos

`func (o *DnaOligosPaginatedList) GetDnaOligos() []DnaOligo`

GetDnaOligos returns the DnaOligos field if non-nil, zero value otherwise.

### GetDnaOligosOk

`func (o *DnaOligosPaginatedList) GetDnaOligosOk() (*[]DnaOligo, bool)`

GetDnaOligosOk returns a tuple with the DnaOligos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaOligos

`func (o *DnaOligosPaginatedList) SetDnaOligos(v []DnaOligo)`

SetDnaOligos sets DnaOligos field to given value.

### HasDnaOligos

`func (o *DnaOligosPaginatedList) HasDnaOligos() bool`

HasDnaOligos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


