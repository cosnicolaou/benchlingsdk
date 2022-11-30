# RegisterEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityIds** | **[]string** | Array of entity IDs | 
**NamingStrategy** | [**NamingStrategy**](NamingStrategy.md) |  | 

## Methods

### NewRegisterEntities

`func NewRegisterEntities(entityIds []string, namingStrategy NamingStrategy, ) *RegisterEntities`

NewRegisterEntities instantiates a new RegisterEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterEntitiesWithDefaults

`func NewRegisterEntitiesWithDefaults() *RegisterEntities`

NewRegisterEntitiesWithDefaults instantiates a new RegisterEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityIds

`func (o *RegisterEntities) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *RegisterEntities) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *RegisterEntities) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.


### GetNamingStrategy

`func (o *RegisterEntities) GetNamingStrategy() NamingStrategy`

GetNamingStrategy returns the NamingStrategy field if non-nil, zero value otherwise.

### GetNamingStrategyOk

`func (o *RegisterEntities) GetNamingStrategyOk() (*NamingStrategy, bool)`

GetNamingStrategyOk returns a tuple with the NamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingStrategy

`func (o *RegisterEntities) SetNamingStrategy(v NamingStrategy)`

SetNamingStrategy sets NamingStrategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


