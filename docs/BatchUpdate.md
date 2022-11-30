# BatchUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultConcentration** | Pointer to [**DefaultConcentrationSummary**](DefaultConcentrationSummary.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 

## Methods

### NewBatchUpdate

`func NewBatchUpdate() *BatchUpdate`

NewBatchUpdate instantiates a new BatchUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateWithDefaults

`func NewBatchUpdateWithDefaults() *BatchUpdate`

NewBatchUpdateWithDefaults instantiates a new BatchUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultConcentration

`func (o *BatchUpdate) GetDefaultConcentration() DefaultConcentrationSummary`

GetDefaultConcentration returns the DefaultConcentration field if non-nil, zero value otherwise.

### GetDefaultConcentrationOk

`func (o *BatchUpdate) GetDefaultConcentrationOk() (*DefaultConcentrationSummary, bool)`

GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcentration

`func (o *BatchUpdate) SetDefaultConcentration(v DefaultConcentrationSummary)`

SetDefaultConcentration sets DefaultConcentration field to given value.

### HasDefaultConcentration

`func (o *BatchUpdate) HasDefaultConcentration() bool`

HasDefaultConcentration returns a boolean if a field has been set.

### GetFields

`func (o *BatchUpdate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BatchUpdate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BatchUpdate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BatchUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


