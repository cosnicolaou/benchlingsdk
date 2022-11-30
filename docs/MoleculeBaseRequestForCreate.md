# MoleculeBaseRequestForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the Molecule. | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Molecule&#39;s authors. | [optional] 
**ChemicalStructure** | [**MoleculeBaseRequestChemicalStructure**](MoleculeBaseRequestChemicalStructure.md) |  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Molecule. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Molecule. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the Molecule.  | [optional] 
**Name** | **string** | Name of the Molecule.  | 
**SchemaId** | Pointer to **string** | ID of the Molecule&#39;s schema.  | [optional] 

## Methods

### NewMoleculeBaseRequestForCreate

`func NewMoleculeBaseRequestForCreate(chemicalStructure MoleculeBaseRequestChemicalStructure, name string, ) *MoleculeBaseRequestForCreate`

NewMoleculeBaseRequestForCreate instantiates a new MoleculeBaseRequestForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculeBaseRequestForCreateWithDefaults

`func NewMoleculeBaseRequestForCreateWithDefaults() *MoleculeBaseRequestForCreate`

NewMoleculeBaseRequestForCreateWithDefaults instantiates a new MoleculeBaseRequestForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *MoleculeBaseRequestForCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MoleculeBaseRequestForCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MoleculeBaseRequestForCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MoleculeBaseRequestForCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *MoleculeBaseRequestForCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *MoleculeBaseRequestForCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *MoleculeBaseRequestForCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *MoleculeBaseRequestForCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetChemicalStructure

`func (o *MoleculeBaseRequestForCreate) GetChemicalStructure() MoleculeBaseRequestChemicalStructure`

GetChemicalStructure returns the ChemicalStructure field if non-nil, zero value otherwise.

### GetChemicalStructureOk

`func (o *MoleculeBaseRequestForCreate) GetChemicalStructureOk() (*MoleculeBaseRequestChemicalStructure, bool)`

GetChemicalStructureOk returns a tuple with the ChemicalStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChemicalStructure

`func (o *MoleculeBaseRequestForCreate) SetChemicalStructure(v MoleculeBaseRequestChemicalStructure)`

SetChemicalStructure sets ChemicalStructure field to given value.


### GetCustomFields

`func (o *MoleculeBaseRequestForCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MoleculeBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MoleculeBaseRequestForCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MoleculeBaseRequestForCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *MoleculeBaseRequestForCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MoleculeBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MoleculeBaseRequestForCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MoleculeBaseRequestForCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *MoleculeBaseRequestForCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MoleculeBaseRequestForCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MoleculeBaseRequestForCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MoleculeBaseRequestForCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *MoleculeBaseRequestForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoleculeBaseRequestForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoleculeBaseRequestForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *MoleculeBaseRequestForCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *MoleculeBaseRequestForCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *MoleculeBaseRequestForCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *MoleculeBaseRequestForCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


