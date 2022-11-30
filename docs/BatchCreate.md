# BatchCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultConcentration** | Pointer to [**DefaultConcentrationSummary**](DefaultConcentrationSummary.md) |  | [optional] 
**EntityId** | Pointer to **string** | API identifier for the entity that the batch will be added to. | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 

## Methods

### NewBatchCreate

`func NewBatchCreate() *BatchCreate`

NewBatchCreate instantiates a new BatchCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCreateWithDefaults

`func NewBatchCreateWithDefaults() *BatchCreate`

NewBatchCreateWithDefaults instantiates a new BatchCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultConcentration

`func (o *BatchCreate) GetDefaultConcentration() DefaultConcentrationSummary`

GetDefaultConcentration returns the DefaultConcentration field if non-nil, zero value otherwise.

### GetDefaultConcentrationOk

`func (o *BatchCreate) GetDefaultConcentrationOk() (*DefaultConcentrationSummary, bool)`

GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcentration

`func (o *BatchCreate) SetDefaultConcentration(v DefaultConcentrationSummary)`

SetDefaultConcentration sets DefaultConcentration field to given value.

### HasDefaultConcentration

`func (o *BatchCreate) HasDefaultConcentration() bool`

HasDefaultConcentration returns a boolean if a field has been set.

### GetEntityId

`func (o *BatchCreate) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *BatchCreate) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *BatchCreate) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *BatchCreate) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetFields

`func (o *BatchCreate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BatchCreate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BatchCreate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BatchCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


