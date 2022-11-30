# MoleculeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the Molecule. | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Molecule&#39;s authors. | [optional] 
**ChemicalStructure** | [**MoleculeBaseRequestChemicalStructure**](MoleculeBaseRequestChemicalStructure.md) |  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Molecule. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Molecule. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the entity. Can be left empty when registryId is present. | [optional] 
**Name** | **string** | Name of the Molecule.  | 
**SchemaId** | Pointer to **string** | ID of the Molecule&#39;s schema.  | [optional] 
**EntityRegistryId** | Pointer to **string** | Entity registry ID to set for the registered entity. Cannot specify both entityRegistryId and namingStrategy at the same time.  | [optional] 
**NamingStrategy** | Pointer to [**NamingStrategy**](NamingStrategy.md) |  | [optional] 
**RegistryId** | Pointer to **string** | Registry ID into which entity should be registered. this is the ID of the registry which was configured for a particular organization To get available registryIds, use the [/registries endpoint](#/Registry/listRegistries)  Required in order for entities to be created directly in the registry.  | [optional] 

## Methods

### NewMoleculeCreate

`func NewMoleculeCreate(chemicalStructure MoleculeBaseRequestChemicalStructure, name string, ) *MoleculeCreate`

NewMoleculeCreate instantiates a new MoleculeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculeCreateWithDefaults

`func NewMoleculeCreateWithDefaults() *MoleculeCreate`

NewMoleculeCreateWithDefaults instantiates a new MoleculeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *MoleculeCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MoleculeCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MoleculeCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MoleculeCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *MoleculeCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *MoleculeCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *MoleculeCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *MoleculeCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetChemicalStructure

`func (o *MoleculeCreate) GetChemicalStructure() MoleculeBaseRequestChemicalStructure`

GetChemicalStructure returns the ChemicalStructure field if non-nil, zero value otherwise.

### GetChemicalStructureOk

`func (o *MoleculeCreate) GetChemicalStructureOk() (*MoleculeBaseRequestChemicalStructure, bool)`

GetChemicalStructureOk returns a tuple with the ChemicalStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChemicalStructure

`func (o *MoleculeCreate) SetChemicalStructure(v MoleculeBaseRequestChemicalStructure)`

SetChemicalStructure sets ChemicalStructure field to given value.


### GetCustomFields

`func (o *MoleculeCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MoleculeCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MoleculeCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MoleculeCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *MoleculeCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MoleculeCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MoleculeCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MoleculeCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *MoleculeCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MoleculeCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MoleculeCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MoleculeCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *MoleculeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoleculeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoleculeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *MoleculeCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *MoleculeCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *MoleculeCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *MoleculeCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *MoleculeCreate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *MoleculeCreate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *MoleculeCreate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *MoleculeCreate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetNamingStrategy

`func (o *MoleculeCreate) GetNamingStrategy() NamingStrategy`

GetNamingStrategy returns the NamingStrategy field if non-nil, zero value otherwise.

### GetNamingStrategyOk

`func (o *MoleculeCreate) GetNamingStrategyOk() (*NamingStrategy, bool)`

GetNamingStrategyOk returns a tuple with the NamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingStrategy

`func (o *MoleculeCreate) SetNamingStrategy(v NamingStrategy)`

SetNamingStrategy sets NamingStrategy field to given value.

### HasNamingStrategy

`func (o *MoleculeCreate) HasNamingStrategy() bool`

HasNamingStrategy returns a boolean if a field has been set.

### GetRegistryId

`func (o *MoleculeCreate) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *MoleculeCreate) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *MoleculeCreate) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *MoleculeCreate) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


