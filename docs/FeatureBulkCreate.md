# FeatureBulkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | The color of the annotations generated by the feature. Must be a valid hex string | [optional] 
**FeatureLibraryId** | **string** | The id of the feature library the feature belongs to | 
**FeatureType** | **NullableString** | The type of feature, like gene, promoter, etc. Note: This is an arbitrary string, not an enum  | 
**Name** | **string** | The name of the feature | 
**Pattern** | **string** | The pattern used for matching during auto-annotation. | 
**MatchType** | **string** | The match type of the feature. Used to determine how auto-annotate matches are made. | 

## Methods

### NewFeatureBulkCreate

`func NewFeatureBulkCreate(featureLibraryId string, featureType NullableString, name string, pattern string, matchType string, ) *FeatureBulkCreate`

NewFeatureBulkCreate instantiates a new FeatureBulkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureBulkCreateWithDefaults

`func NewFeatureBulkCreateWithDefaults() *FeatureBulkCreate`

NewFeatureBulkCreateWithDefaults instantiates a new FeatureBulkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *FeatureBulkCreate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FeatureBulkCreate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FeatureBulkCreate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *FeatureBulkCreate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetFeatureLibraryId

`func (o *FeatureBulkCreate) GetFeatureLibraryId() string`

GetFeatureLibraryId returns the FeatureLibraryId field if non-nil, zero value otherwise.

### GetFeatureLibraryIdOk

`func (o *FeatureBulkCreate) GetFeatureLibraryIdOk() (*string, bool)`

GetFeatureLibraryIdOk returns a tuple with the FeatureLibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLibraryId

`func (o *FeatureBulkCreate) SetFeatureLibraryId(v string)`

SetFeatureLibraryId sets FeatureLibraryId field to given value.


### GetFeatureType

`func (o *FeatureBulkCreate) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *FeatureBulkCreate) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *FeatureBulkCreate) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.


### SetFeatureTypeNil

`func (o *FeatureBulkCreate) SetFeatureTypeNil(b bool)`

 SetFeatureTypeNil sets the value for FeatureType to be an explicit nil

### UnsetFeatureType
`func (o *FeatureBulkCreate) UnsetFeatureType()`

UnsetFeatureType ensures that no value is present for FeatureType, not even an explicit nil
### GetName

`func (o *FeatureBulkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureBulkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureBulkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPattern

`func (o *FeatureBulkCreate) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *FeatureBulkCreate) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *FeatureBulkCreate) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetMatchType

`func (o *FeatureBulkCreate) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *FeatureBulkCreate) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *FeatureBulkCreate) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


