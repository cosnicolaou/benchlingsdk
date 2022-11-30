# RequestTaskBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the request task | 
**Fields** | Pointer to [**ModelMap**](map.md) | Schema fields to set on the request task. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**SampleGroupIds** | Pointer to **[]string** | IDs of all request sample groups now associated with this task. | [optional] 

## Methods

### NewRequestTaskBase

`func NewRequestTaskBase(id string, ) *RequestTaskBase`

NewRequestTaskBase instantiates a new RequestTaskBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTaskBaseWithDefaults

`func NewRequestTaskBaseWithDefaults() *RequestTaskBase`

NewRequestTaskBaseWithDefaults instantiates a new RequestTaskBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestTaskBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestTaskBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestTaskBase) SetId(v string)`

SetId sets Id field to given value.


### GetFields

`func (o *RequestTaskBase) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestTaskBase) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestTaskBase) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestTaskBase) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSampleGroupIds

`func (o *RequestTaskBase) GetSampleGroupIds() []string`

GetSampleGroupIds returns the SampleGroupIds field if non-nil, zero value otherwise.

### GetSampleGroupIdsOk

`func (o *RequestTaskBase) GetSampleGroupIdsOk() (*[]string, bool)`

GetSampleGroupIdsOk returns a tuple with the SampleGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroupIds

`func (o *RequestTaskBase) SetSampleGroupIds(v []string)`

SetSampleGroupIds sets SampleGroupIds field to given value.

### HasSampleGroupIds

`func (o *RequestTaskBase) HasSampleGroupIds() bool`

HasSampleGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


