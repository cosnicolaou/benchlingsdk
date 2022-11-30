# OligoCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the Oligo | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Oligo&#39;s authors. | [optional] 
**Bases** | **string** | Base pairs of the oligo.  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Oligo. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Oligo. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the entity. Can be left empty when registryId is present. | [optional] 
**Name** | **string** | Name of the Oligo.  | 
**SchemaId** | Pointer to **string** | ID of the oligo&#39;s schema.  | [optional] 
**EntityRegistryId** | Pointer to **string** | Entity registry ID to set for the registered entity. Cannot specify both entityRegistryId and namingStrategy at the same time.  | [optional] 
**NamingStrategy** | Pointer to [**NamingStrategy**](NamingStrategy.md) |  | [optional] 
**RegistryId** | Pointer to **string** | Registry ID into which entity should be registered. this is the ID of the registry which was configured for a particular organization To get available registryIds, use the [/registries endpoint](#/Registry/listRegistries)  Required in order for entities to be created directly in the registry.  | [optional] 

## Methods

### NewOligoCreate

`func NewOligoCreate(bases string, name string, ) *OligoCreate`

NewOligoCreate instantiates a new OligoCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOligoCreateWithDefaults

`func NewOligoCreateWithDefaults() *OligoCreate`

NewOligoCreateWithDefaults instantiates a new OligoCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *OligoCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *OligoCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *OligoCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *OligoCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *OligoCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *OligoCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *OligoCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *OligoCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *OligoCreate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *OligoCreate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *OligoCreate) SetBases(v string)`

SetBases sets Bases field to given value.


### GetCustomFields

`func (o *OligoCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OligoCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OligoCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OligoCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *OligoCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *OligoCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *OligoCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *OligoCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *OligoCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *OligoCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *OligoCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *OligoCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *OligoCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OligoCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OligoCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *OligoCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *OligoCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *OligoCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *OligoCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *OligoCreate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *OligoCreate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *OligoCreate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *OligoCreate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetNamingStrategy

`func (o *OligoCreate) GetNamingStrategy() NamingStrategy`

GetNamingStrategy returns the NamingStrategy field if non-nil, zero value otherwise.

### GetNamingStrategyOk

`func (o *OligoCreate) GetNamingStrategyOk() (*NamingStrategy, bool)`

GetNamingStrategyOk returns a tuple with the NamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingStrategy

`func (o *OligoCreate) SetNamingStrategy(v NamingStrategy)`

SetNamingStrategy sets NamingStrategy field to given value.

### HasNamingStrategy

`func (o *OligoCreate) HasNamingStrategy() bool`

HasNamingStrategy returns a boolean if a field has been set.

### GetRegistryId

`func (o *OligoCreate) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *OligoCreate) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *OligoCreate) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *OligoCreate) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


