# ContainerContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batch** | Pointer to [**NullableContainerContentBatch**](ContainerContentBatch.md) |  | [optional] 
**Concentration** | Pointer to [**Measurement**](Measurement.md) |  | [optional] 
**Entity** | Pointer to [**NullableContainerContentEntity**](ContainerContentEntity.md) |  | [optional] 

## Methods

### NewContainerContent

`func NewContainerContent() *ContainerContent`

NewContainerContent instantiates a new ContainerContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerContentWithDefaults

`func NewContainerContentWithDefaults() *ContainerContent`

NewContainerContentWithDefaults instantiates a new ContainerContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatch

`func (o *ContainerContent) GetBatch() ContainerContentBatch`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *ContainerContent) GetBatchOk() (*ContainerContentBatch, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *ContainerContent) SetBatch(v ContainerContentBatch)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *ContainerContent) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### SetBatchNil

`func (o *ContainerContent) SetBatchNil(b bool)`

 SetBatchNil sets the value for Batch to be an explicit nil

### UnsetBatch
`func (o *ContainerContent) UnsetBatch()`

UnsetBatch ensures that no value is present for Batch, not even an explicit nil
### GetConcentration

`func (o *ContainerContent) GetConcentration() Measurement`

GetConcentration returns the Concentration field if non-nil, zero value otherwise.

### GetConcentrationOk

`func (o *ContainerContent) GetConcentrationOk() (*Measurement, bool)`

GetConcentrationOk returns a tuple with the Concentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentration

`func (o *ContainerContent) SetConcentration(v Measurement)`

SetConcentration sets Concentration field to given value.

### HasConcentration

`func (o *ContainerContent) HasConcentration() bool`

HasConcentration returns a boolean if a field has been set.

### GetEntity

`func (o *ContainerContent) GetEntity() ContainerContentEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ContainerContent) GetEntityOk() (*ContainerContentEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ContainerContent) SetEntity(v ContainerContentEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ContainerContent) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *ContainerContent) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *ContainerContent) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


