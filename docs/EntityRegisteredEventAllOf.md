# EntityRegisteredEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**GenericEntity**](GenericEntity.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 

## Methods

### NewEntityRegisteredEventAllOf

`func NewEntityRegisteredEventAllOf() *EntityRegisteredEventAllOf`

NewEntityRegisteredEventAllOf instantiates a new EntityRegisteredEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRegisteredEventAllOfWithDefaults

`func NewEntityRegisteredEventAllOfWithDefaults() *EntityRegisteredEventAllOf`

NewEntityRegisteredEventAllOfWithDefaults instantiates a new EntityRegisteredEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *EntityRegisteredEventAllOf) GetEntity() GenericEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *EntityRegisteredEventAllOf) GetEntityOk() (*GenericEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *EntityRegisteredEventAllOf) SetEntity(v GenericEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *EntityRegisteredEventAllOf) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEventType

`func (o *EntityRegisteredEventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EntityRegisteredEventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EntityRegisteredEventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EntityRegisteredEventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


