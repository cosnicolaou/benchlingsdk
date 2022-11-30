# AaSequenceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the AA sequence | [optional] 
**AminoAcids** | Pointer to **string** | Amino acids for the AA sequence.  | [optional] 
**Annotations** | Pointer to [**[]AaAnnotation**](AaAnnotation.md) | Annotations to create on the AA sequence.  | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the AA sequence&#39;s authors. | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the AA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the AA sequence. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the AA sequence.  | [optional] 
**Name** | Pointer to **string** | Name of the AA sequence.  | [optional] 
**SchemaId** | Pointer to **string** | ID of the AA sequence&#39;s schema.  | [optional] 
**EntityRegistryId** | Pointer to **string** |  | [optional] 

## Methods

### NewAaSequenceUpdate

`func NewAaSequenceUpdate() *AaSequenceUpdate`

NewAaSequenceUpdate instantiates a new AaSequenceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaSequenceUpdateWithDefaults

`func NewAaSequenceUpdateWithDefaults() *AaSequenceUpdate`

NewAaSequenceUpdateWithDefaults instantiates a new AaSequenceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *AaSequenceUpdate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AaSequenceUpdate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AaSequenceUpdate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AaSequenceUpdate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAminoAcids

`func (o *AaSequenceUpdate) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *AaSequenceUpdate) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *AaSequenceUpdate) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *AaSequenceUpdate) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetAnnotations

`func (o *AaSequenceUpdate) GetAnnotations() []AaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AaSequenceUpdate) GetAnnotationsOk() (*[]AaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AaSequenceUpdate) SetAnnotations(v []AaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AaSequenceUpdate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *AaSequenceUpdate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *AaSequenceUpdate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *AaSequenceUpdate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *AaSequenceUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *AaSequenceUpdate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AaSequenceUpdate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AaSequenceUpdate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AaSequenceUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *AaSequenceUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AaSequenceUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AaSequenceUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AaSequenceUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *AaSequenceUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AaSequenceUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AaSequenceUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AaSequenceUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *AaSequenceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AaSequenceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AaSequenceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AaSequenceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *AaSequenceUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *AaSequenceUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *AaSequenceUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *AaSequenceUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *AaSequenceUpdate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *AaSequenceUpdate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *AaSequenceUpdate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *AaSequenceUpdate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


