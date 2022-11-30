# SampleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSampleGroup

`func NewSampleGroup() *SampleGroup`

NewSampleGroup instantiates a new SampleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleGroupWithDefaults

`func NewSampleGroupWithDefaults() *SampleGroup`

NewSampleGroupWithDefaults instantiates a new SampleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SampleGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SampleGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SampleGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SampleGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSamples

`func (o *SampleGroup) GetSamples() map[string]interface{}`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *SampleGroup) GetSamplesOk() (*map[string]interface{}, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *SampleGroup) SetSamples(v map[string]interface{})`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *SampleGroup) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


