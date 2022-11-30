# CreateEntityIntoRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityRegistryId** | Pointer to **string** | Entity registry ID to set for the registered entity. Cannot specify both entityRegistryId and namingStrategy at the same time.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the entity. Can be left empty when registryId is present. | [optional] 
**NamingStrategy** | Pointer to [**NamingStrategy**](NamingStrategy.md) |  | [optional] 
**RegistryId** | Pointer to **string** | Registry ID into which entity should be registered. this is the ID of the registry which was configured for a particular organization To get available registryIds, use the [/registries endpoint](#/Registry/listRegistries)  Required in order for entities to be created directly in the registry.  | [optional] 

## Methods

### NewCreateEntityIntoRegistry

`func NewCreateEntityIntoRegistry() *CreateEntityIntoRegistry`

NewCreateEntityIntoRegistry instantiates a new CreateEntityIntoRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityIntoRegistryWithDefaults

`func NewCreateEntityIntoRegistryWithDefaults() *CreateEntityIntoRegistry`

NewCreateEntityIntoRegistryWithDefaults instantiates a new CreateEntityIntoRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityRegistryId

`func (o *CreateEntityIntoRegistry) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *CreateEntityIntoRegistry) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *CreateEntityIntoRegistry) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *CreateEntityIntoRegistry) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetFolderId

`func (o *CreateEntityIntoRegistry) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CreateEntityIntoRegistry) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CreateEntityIntoRegistry) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CreateEntityIntoRegistry) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetNamingStrategy

`func (o *CreateEntityIntoRegistry) GetNamingStrategy() NamingStrategy`

GetNamingStrategy returns the NamingStrategy field if non-nil, zero value otherwise.

### GetNamingStrategyOk

`func (o *CreateEntityIntoRegistry) GetNamingStrategyOk() (*NamingStrategy, bool)`

GetNamingStrategyOk returns a tuple with the NamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingStrategy

`func (o *CreateEntityIntoRegistry) SetNamingStrategy(v NamingStrategy)`

SetNamingStrategy sets NamingStrategy field to given value.

### HasNamingStrategy

`func (o *CreateEntityIntoRegistry) HasNamingStrategy() bool`

HasNamingStrategy returns a boolean if a field has been set.

### GetRegistryId

`func (o *CreateEntityIntoRegistry) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *CreateEntityIntoRegistry) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *CreateEntityIntoRegistry) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *CreateEntityIntoRegistry) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


