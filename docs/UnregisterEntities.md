# UnregisterEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityIds** | **[]string** | Array of entity IDs | 
**FolderId** | **string** | ID of the folder that the entities should be moved to | 

## Methods

### NewUnregisterEntities

`func NewUnregisterEntities(entityIds []string, folderId string, ) *UnregisterEntities`

NewUnregisterEntities instantiates a new UnregisterEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisterEntitiesWithDefaults

`func NewUnregisterEntitiesWithDefaults() *UnregisterEntities`

NewUnregisterEntitiesWithDefaults instantiates a new UnregisterEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityIds

`func (o *UnregisterEntities) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *UnregisterEntities) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *UnregisterEntities) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.


### GetFolderId

`func (o *UnregisterEntities) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UnregisterEntities) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UnregisterEntities) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


