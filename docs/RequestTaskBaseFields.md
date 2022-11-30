# RequestTaskBaseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**ModelMap**](map.md) | Schema fields to set on the request task. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**SampleGroupIds** | Pointer to **[]string** | IDs of all request sample groups now associated with this task. | [optional] 

## Methods

### NewRequestTaskBaseFields

`func NewRequestTaskBaseFields() *RequestTaskBaseFields`

NewRequestTaskBaseFields instantiates a new RequestTaskBaseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTaskBaseFieldsWithDefaults

`func NewRequestTaskBaseFieldsWithDefaults() *RequestTaskBaseFields`

NewRequestTaskBaseFieldsWithDefaults instantiates a new RequestTaskBaseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *RequestTaskBaseFields) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestTaskBaseFields) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestTaskBaseFields) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestTaskBaseFields) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSampleGroupIds

`func (o *RequestTaskBaseFields) GetSampleGroupIds() []string`

GetSampleGroupIds returns the SampleGroupIds field if non-nil, zero value otherwise.

### GetSampleGroupIdsOk

`func (o *RequestTaskBaseFields) GetSampleGroupIdsOk() (*[]string, bool)`

GetSampleGroupIdsOk returns a tuple with the SampleGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroupIds

`func (o *RequestTaskBaseFields) SetSampleGroupIds(v []string)`

SetSampleGroupIds sets SampleGroupIds field to given value.

### HasSampleGroupIds

`func (o *RequestTaskBaseFields) HasSampleGroupIds() bool`

HasSampleGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


