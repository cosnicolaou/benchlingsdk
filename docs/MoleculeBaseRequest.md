# MoleculeBaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the Molecule. | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Molecule&#39;s authors. | [optional] 
**ChemicalStructure** | Pointer to [**MoleculeBaseRequestChemicalStructure**](MoleculeBaseRequestChemicalStructure.md) |  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Molecule. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Molecule. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the Molecule.  | [optional] 
**Name** | Pointer to **string** | Name of the Molecule.  | [optional] 
**SchemaId** | Pointer to **string** | ID of the Molecule&#39;s schema.  | [optional] 

## Methods

### NewMoleculeBaseRequest

`func NewMoleculeBaseRequest() *MoleculeBaseRequest`

NewMoleculeBaseRequest instantiates a new MoleculeBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculeBaseRequestWithDefaults

`func NewMoleculeBaseRequestWithDefaults() *MoleculeBaseRequest`

NewMoleculeBaseRequestWithDefaults instantiates a new MoleculeBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *MoleculeBaseRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MoleculeBaseRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MoleculeBaseRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MoleculeBaseRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *MoleculeBaseRequest) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *MoleculeBaseRequest) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *MoleculeBaseRequest) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *MoleculeBaseRequest) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetChemicalStructure

`func (o *MoleculeBaseRequest) GetChemicalStructure() MoleculeBaseRequestChemicalStructure`

GetChemicalStructure returns the ChemicalStructure field if non-nil, zero value otherwise.

### GetChemicalStructureOk

`func (o *MoleculeBaseRequest) GetChemicalStructureOk() (*MoleculeBaseRequestChemicalStructure, bool)`

GetChemicalStructureOk returns a tuple with the ChemicalStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChemicalStructure

`func (o *MoleculeBaseRequest) SetChemicalStructure(v MoleculeBaseRequestChemicalStructure)`

SetChemicalStructure sets ChemicalStructure field to given value.

### HasChemicalStructure

`func (o *MoleculeBaseRequest) HasChemicalStructure() bool`

HasChemicalStructure returns a boolean if a field has been set.

### GetCustomFields

`func (o *MoleculeBaseRequest) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MoleculeBaseRequest) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MoleculeBaseRequest) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MoleculeBaseRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *MoleculeBaseRequest) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MoleculeBaseRequest) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MoleculeBaseRequest) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MoleculeBaseRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *MoleculeBaseRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MoleculeBaseRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MoleculeBaseRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MoleculeBaseRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *MoleculeBaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoleculeBaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoleculeBaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MoleculeBaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *MoleculeBaseRequest) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *MoleculeBaseRequest) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *MoleculeBaseRequest) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *MoleculeBaseRequest) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


