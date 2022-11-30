# RequestSampleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Samples** | Pointer to [**map[string]RequestSampleGroupSamplesValue**](RequestSampleGroupSamplesValue.md) | The key for each RequestSample should match one of the samplesSchema[n].name property in the request schema json.  | [optional] 

## Methods

### NewRequestSampleGroup

`func NewRequestSampleGroup() *RequestSampleGroup`

NewRequestSampleGroup instantiates a new RequestSampleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSampleGroupWithDefaults

`func NewRequestSampleGroupWithDefaults() *RequestSampleGroup`

NewRequestSampleGroupWithDefaults instantiates a new RequestSampleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestSampleGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestSampleGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestSampleGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestSampleGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSamples

`func (o *RequestSampleGroup) GetSamples() map[string]RequestSampleGroupSamplesValue`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *RequestSampleGroup) GetSamplesOk() (*map[string]RequestSampleGroupSamplesValue, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *RequestSampleGroup) SetSamples(v map[string]RequestSampleGroupSamplesValue)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *RequestSampleGroup) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


