# FeatureLibrariesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**FeatureLibraries** | Pointer to [**[]FeatureLibrary**](FeatureLibrary.md) |  | [optional] 

## Methods

### NewFeatureLibrariesPaginatedList

`func NewFeatureLibrariesPaginatedList() *FeatureLibrariesPaginatedList`

NewFeatureLibrariesPaginatedList instantiates a new FeatureLibrariesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureLibrariesPaginatedListWithDefaults

`func NewFeatureLibrariesPaginatedListWithDefaults() *FeatureLibrariesPaginatedList`

NewFeatureLibrariesPaginatedListWithDefaults instantiates a new FeatureLibrariesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *FeatureLibrariesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *FeatureLibrariesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *FeatureLibrariesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *FeatureLibrariesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetFeatureLibraries

`func (o *FeatureLibrariesPaginatedList) GetFeatureLibraries() []FeatureLibrary`

GetFeatureLibraries returns the FeatureLibraries field if non-nil, zero value otherwise.

### GetFeatureLibrariesOk

`func (o *FeatureLibrariesPaginatedList) GetFeatureLibrariesOk() (*[]FeatureLibrary, bool)`

GetFeatureLibrariesOk returns a tuple with the FeatureLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLibraries

`func (o *FeatureLibrariesPaginatedList) SetFeatureLibraries(v []FeatureLibrary)`

SetFeatureLibraries sets FeatureLibraries field to given value.

### HasFeatureLibraries

`func (o *FeatureLibrariesPaginatedList) HasFeatureLibraries() bool`

HasFeatureLibraries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


