# MoleculeBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**EntityRegistryId** | Pointer to **string** |  | [optional] 
**Aliases** | Pointer to **[]string** | Aliases to add to the Molecule. | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Molecule&#39;s authors. | [optional] 
**ChemicalStructure** | Pointer to [**MoleculeBaseRequestChemicalStructure**](MoleculeBaseRequestChemicalStructure.md) |  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Molecule. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Molecule. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the Molecule.  | [optional] 
**Name** | Pointer to **string** | Name of the Molecule.  | [optional] 
**SchemaId** | Pointer to **string** | ID of the Molecule&#39;s schema.  | [optional] 

## Methods

### NewMoleculeBulkUpdate

`func NewMoleculeBulkUpdate() *MoleculeBulkUpdate`

NewMoleculeBulkUpdate instantiates a new MoleculeBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculeBulkUpdateWithDefaults

`func NewMoleculeBulkUpdateWithDefaults() *MoleculeBulkUpdate`

NewMoleculeBulkUpdateWithDefaults instantiates a new MoleculeBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoleculeBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoleculeBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoleculeBulkUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoleculeBulkUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *MoleculeBulkUpdate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *MoleculeBulkUpdate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *MoleculeBulkUpdate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *MoleculeBulkUpdate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetAliases

`func (o *MoleculeBulkUpdate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MoleculeBulkUpdate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MoleculeBulkUpdate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MoleculeBulkUpdate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *MoleculeBulkUpdate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *MoleculeBulkUpdate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *MoleculeBulkUpdate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *MoleculeBulkUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetChemicalStructure

`func (o *MoleculeBulkUpdate) GetChemicalStructure() MoleculeBaseRequestChemicalStructure`

GetChemicalStructure returns the ChemicalStructure field if non-nil, zero value otherwise.

### GetChemicalStructureOk

`func (o *MoleculeBulkUpdate) GetChemicalStructureOk() (*MoleculeBaseRequestChemicalStructure, bool)`

GetChemicalStructureOk returns a tuple with the ChemicalStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChemicalStructure

`func (o *MoleculeBulkUpdate) SetChemicalStructure(v MoleculeBaseRequestChemicalStructure)`

SetChemicalStructure sets ChemicalStructure field to given value.

### HasChemicalStructure

`func (o *MoleculeBulkUpdate) HasChemicalStructure() bool`

HasChemicalStructure returns a boolean if a field has been set.

### GetCustomFields

`func (o *MoleculeBulkUpdate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MoleculeBulkUpdate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MoleculeBulkUpdate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MoleculeBulkUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *MoleculeBulkUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MoleculeBulkUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MoleculeBulkUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MoleculeBulkUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *MoleculeBulkUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MoleculeBulkUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MoleculeBulkUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MoleculeBulkUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *MoleculeBulkUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoleculeBulkUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoleculeBulkUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MoleculeBulkUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *MoleculeBulkUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *MoleculeBulkUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *MoleculeBulkUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *MoleculeBulkUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


