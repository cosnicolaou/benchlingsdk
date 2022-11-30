# CustomEntityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the custom entity | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the custom entity&#39;s authors. | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the custom entity. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Schema fields to set on the custom entity. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder that the entity is moved into | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**EntityRegistryId** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomEntityUpdate

`func NewCustomEntityUpdate() *CustomEntityUpdate`

NewCustomEntityUpdate instantiates a new CustomEntityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityUpdateWithDefaults

`func NewCustomEntityUpdateWithDefaults() *CustomEntityUpdate`

NewCustomEntityUpdateWithDefaults instantiates a new CustomEntityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *CustomEntityUpdate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CustomEntityUpdate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CustomEntityUpdate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CustomEntityUpdate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *CustomEntityUpdate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *CustomEntityUpdate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *CustomEntityUpdate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *CustomEntityUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomEntityUpdate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomEntityUpdate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomEntityUpdate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomEntityUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *CustomEntityUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomEntityUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomEntityUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomEntityUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *CustomEntityUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CustomEntityUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CustomEntityUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CustomEntityUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *CustomEntityUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEntityUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEntityUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomEntityUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *CustomEntityUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *CustomEntityUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *CustomEntityUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *CustomEntityUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *CustomEntityUpdate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *CustomEntityUpdate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *CustomEntityUpdate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *CustomEntityUpdate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


