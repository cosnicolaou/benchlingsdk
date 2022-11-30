# RequestResponseSamplesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to [**NullableRequestResponseSamplesInnerBatch**](RequestResponseSamplesInnerBatch.md) |  | [optional] 
**Entity** | Pointer to [**NullableRequestResponseSamplesInnerEntity**](RequestResponseSamplesInnerEntity.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the sample, based on the status of the stage run that generated it. | [optional] 

## Methods

### NewRequestResponseSamplesInner

`func NewRequestResponseSamplesInner() *RequestResponseSamplesInner`

NewRequestResponseSamplesInner instantiates a new RequestResponseSamplesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestResponseSamplesInnerWithDefaults

`func NewRequestResponseSamplesInnerWithDefaults() *RequestResponseSamplesInner`

NewRequestResponseSamplesInnerWithDefaults instantiates a new RequestResponseSamplesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *RequestResponseSamplesInner) GetBatch() RequestResponseSamplesInnerBatch`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *RequestResponseSamplesInner) GetBatchOk() (*RequestResponseSamplesInnerBatch, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *RequestResponseSamplesInner) SetBatch(v RequestResponseSamplesInnerBatch)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *RequestResponseSamplesInner) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### SetBatchNil

`func (o *RequestResponseSamplesInner) SetBatchNil(b bool)`

 SetBatchNil sets the value for Batch to be an explicit nil

### UnsetBatch
`func (o *RequestResponseSamplesInner) UnsetBatch()`

UnsetBatch ensures that no value is present for Batch, not even an explicit nil
### GetEntity

`func (o *RequestResponseSamplesInner) GetEntity() RequestResponseSamplesInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RequestResponseSamplesInner) GetEntityOk() (*RequestResponseSamplesInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RequestResponseSamplesInner) SetEntity(v RequestResponseSamplesInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RequestResponseSamplesInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *RequestResponseSamplesInner) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *RequestResponseSamplesInner) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil
### GetStatus

`func (o *RequestResponseSamplesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestResponseSamplesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestResponseSamplesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestResponseSamplesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


