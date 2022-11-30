# RequestSampleGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | [**map[string]RequestSampleGroupSamplesValue**](RequestSampleGroupSamplesValue.md) | The key for each RequestSample should match one of the samplesSchema[n].name property in the request schema json.  | 

## Methods

### NewRequestSampleGroupCreate

`func NewRequestSampleGroupCreate(samples map[string]RequestSampleGroupSamplesValue, ) *RequestSampleGroupCreate`

NewRequestSampleGroupCreate instantiates a new RequestSampleGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSampleGroupCreateWithDefaults

`func NewRequestSampleGroupCreateWithDefaults() *RequestSampleGroupCreate`

NewRequestSampleGroupCreateWithDefaults instantiates a new RequestSampleGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *RequestSampleGroupCreate) GetSamples() map[string]RequestSampleGroupSamplesValue`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *RequestSampleGroupCreate) GetSamplesOk() (*map[string]RequestSampleGroupSamplesValue, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *RequestSampleGroupCreate) SetSamples(v map[string]RequestSampleGroupSamplesValue)`

SetSamples sets Samples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


