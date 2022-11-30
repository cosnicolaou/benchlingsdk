# SampleGroupStatusUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SampleGroupId** | **string** | The string id of the sample group | 
**Status** | [**SampleGroupStatus**](SampleGroupStatus.md) |  | 

## Methods

### NewSampleGroupStatusUpdate

`func NewSampleGroupStatusUpdate(sampleGroupId string, status SampleGroupStatus, ) *SampleGroupStatusUpdate`

NewSampleGroupStatusUpdate instantiates a new SampleGroupStatusUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleGroupStatusUpdateWithDefaults

`func NewSampleGroupStatusUpdateWithDefaults() *SampleGroupStatusUpdate`

NewSampleGroupStatusUpdateWithDefaults instantiates a new SampleGroupStatusUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSampleGroupId

`func (o *SampleGroupStatusUpdate) GetSampleGroupId() string`

GetSampleGroupId returns the SampleGroupId field if non-nil, zero value otherwise.

### GetSampleGroupIdOk

`func (o *SampleGroupStatusUpdate) GetSampleGroupIdOk() (*string, bool)`

GetSampleGroupIdOk returns a tuple with the SampleGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroupId

`func (o *SampleGroupStatusUpdate) SetSampleGroupId(v string)`

SetSampleGroupId sets SampleGroupId field to given value.


### GetStatus

`func (o *SampleGroupStatusUpdate) GetStatus() SampleGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SampleGroupStatusUpdate) GetStatusOk() (*SampleGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SampleGroupStatusUpdate) SetStatus(v SampleGroupStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


