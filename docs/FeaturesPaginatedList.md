# FeaturesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]Feature**](Feature.md) | List of features for the page | [optional] 

## Methods

### NewFeaturesPaginatedList

`func NewFeaturesPaginatedList() *FeaturesPaginatedList`

NewFeaturesPaginatedList instantiates a new FeaturesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesPaginatedListWithDefaults

`func NewFeaturesPaginatedListWithDefaults() *FeaturesPaginatedList`

NewFeaturesPaginatedListWithDefaults instantiates a new FeaturesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *FeaturesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *FeaturesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *FeaturesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *FeaturesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetFeatures

`func (o *FeaturesPaginatedList) GetFeatures() []Feature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeaturesPaginatedList) GetFeaturesOk() (*[]Feature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeaturesPaginatedList) SetFeatures(v []Feature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FeaturesPaginatedList) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


