# BoxesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boxes** | Pointer to [**[]Box**](Box.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewBoxesPaginatedList

`func NewBoxesPaginatedList() *BoxesPaginatedList`

NewBoxesPaginatedList instantiates a new BoxesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxesPaginatedListWithDefaults

`func NewBoxesPaginatedListWithDefaults() *BoxesPaginatedList`

NewBoxesPaginatedListWithDefaults instantiates a new BoxesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxes

`func (o *BoxesPaginatedList) GetBoxes() []Box`

GetBoxes returns the Boxes field if non-nil, zero value otherwise.

### GetBoxesOk

`func (o *BoxesPaginatedList) GetBoxesOk() (*[]Box, bool)`

GetBoxesOk returns a tuple with the Boxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxes

`func (o *BoxesPaginatedList) SetBoxes(v []Box)`

SetBoxes sets Boxes field to given value.

### HasBoxes

`func (o *BoxesPaginatedList) HasBoxes() bool`

HasBoxes returns a boolean if a field has been set.

### GetNextToken

`func (o *BoxesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BoxesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BoxesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BoxesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


