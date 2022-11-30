# LocationsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationsPaginatedList

`func NewLocationsPaginatedList() *LocationsPaginatedList`

NewLocationsPaginatedList instantiates a new LocationsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationsPaginatedListWithDefaults

`func NewLocationsPaginatedListWithDefaults() *LocationsPaginatedList`

NewLocationsPaginatedListWithDefaults instantiates a new LocationsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *LocationsPaginatedList) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *LocationsPaginatedList) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *LocationsPaginatedList) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *LocationsPaginatedList) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetNextToken

`func (o *LocationsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *LocationsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *LocationsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *LocationsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


