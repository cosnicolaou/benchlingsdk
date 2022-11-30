# PlatesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Plates** | Pointer to [**[]Plate**](Plate.md) |  | [optional] 

## Methods

### NewPlatesPaginatedList

`func NewPlatesPaginatedList() *PlatesPaginatedList`

NewPlatesPaginatedList instantiates a new PlatesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatesPaginatedListWithDefaults

`func NewPlatesPaginatedListWithDefaults() *PlatesPaginatedList`

NewPlatesPaginatedListWithDefaults instantiates a new PlatesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *PlatesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *PlatesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *PlatesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *PlatesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetPlates

`func (o *PlatesPaginatedList) GetPlates() []Plate`

GetPlates returns the Plates field if non-nil, zero value otherwise.

### GetPlatesOk

`func (o *PlatesPaginatedList) GetPlatesOk() (*[]Plate, bool)`

GetPlatesOk returns a tuple with the Plates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlates

`func (o *PlatesPaginatedList) SetPlates(v []Plate)`

SetPlates sets Plates field to given value.

### HasPlates

`func (o *PlatesPaginatedList) HasPlates() bool`

HasPlates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


