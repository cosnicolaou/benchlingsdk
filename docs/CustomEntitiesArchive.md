# CustomEntitiesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomEntityIds** | **[]string** |  | 
**Reason** | [**EntityArchiveReason**](EntityArchiveReason.md) |  | 

## Methods

### NewCustomEntitiesArchive

`func NewCustomEntitiesArchive(customEntityIds []string, reason EntityArchiveReason, ) *CustomEntitiesArchive`

NewCustomEntitiesArchive instantiates a new CustomEntitiesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntitiesArchiveWithDefaults

`func NewCustomEntitiesArchiveWithDefaults() *CustomEntitiesArchive`

NewCustomEntitiesArchiveWithDefaults instantiates a new CustomEntitiesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomEntityIds

`func (o *CustomEntitiesArchive) GetCustomEntityIds() []string`

GetCustomEntityIds returns the CustomEntityIds field if non-nil, zero value otherwise.

### GetCustomEntityIdsOk

`func (o *CustomEntitiesArchive) GetCustomEntityIdsOk() (*[]string, bool)`

GetCustomEntityIdsOk returns a tuple with the CustomEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEntityIds

`func (o *CustomEntitiesArchive) SetCustomEntityIds(v []string)`

SetCustomEntityIds sets CustomEntityIds field to given value.


### GetReason

`func (o *CustomEntitiesArchive) GetReason() EntityArchiveReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CustomEntitiesArchive) GetReasonOk() (*EntityArchiveReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CustomEntitiesArchive) SetReason(v EntityArchiveReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


