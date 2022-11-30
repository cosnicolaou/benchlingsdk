# RnaSequenceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the RNA sequence | [optional] 
**Annotations** | Pointer to [**[]RnaAnnotation**](RnaAnnotation.md) | Annotations to create on the RNA sequence.  | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the RNA sequence&#39;s authors. | [optional] 
**Bases** | Pointer to **string** | Base pairs for the RNA sequence.  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the RNA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the RNA sequence. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the RNA sequence.  | [optional] 
**IsCircular** | Pointer to **bool** | Whether the RNA sequence is circular or linear. RNA sequences can only be linear  | [optional] 
**Name** | Pointer to **string** | Name of the RNA sequence.  | [optional] 
**Primers** | Pointer to [**[]Primer**](Primer.md) |  | [optional] 
**SchemaId** | Pointer to **string** | ID of the RNA sequence&#39;s schema.  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) | Translations to create on the RNA sequence. Translations are specified by either a combination of &#39;start&#39; and &#39;end&#39; fields, or a list of regions. Both cannot be provided.  | [optional] 
**EntityRegistryId** | Pointer to **string** |  | [optional] 

## Methods

### NewRnaSequenceUpdate

`func NewRnaSequenceUpdate() *RnaSequenceUpdate`

NewRnaSequenceUpdate instantiates a new RnaSequenceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaSequenceUpdateWithDefaults

`func NewRnaSequenceUpdateWithDefaults() *RnaSequenceUpdate`

NewRnaSequenceUpdateWithDefaults instantiates a new RnaSequenceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *RnaSequenceUpdate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *RnaSequenceUpdate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *RnaSequenceUpdate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *RnaSequenceUpdate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *RnaSequenceUpdate) GetAnnotations() []RnaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RnaSequenceUpdate) GetAnnotationsOk() (*[]RnaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RnaSequenceUpdate) SetAnnotations(v []RnaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *RnaSequenceUpdate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *RnaSequenceUpdate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *RnaSequenceUpdate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *RnaSequenceUpdate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *RnaSequenceUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *RnaSequenceUpdate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *RnaSequenceUpdate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *RnaSequenceUpdate) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *RnaSequenceUpdate) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetCustomFields

`func (o *RnaSequenceUpdate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RnaSequenceUpdate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RnaSequenceUpdate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RnaSequenceUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *RnaSequenceUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RnaSequenceUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RnaSequenceUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RnaSequenceUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *RnaSequenceUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RnaSequenceUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RnaSequenceUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RnaSequenceUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIsCircular

`func (o *RnaSequenceUpdate) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *RnaSequenceUpdate) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *RnaSequenceUpdate) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.

### HasIsCircular

`func (o *RnaSequenceUpdate) HasIsCircular() bool`

HasIsCircular returns a boolean if a field has been set.

### GetName

`func (o *RnaSequenceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RnaSequenceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RnaSequenceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RnaSequenceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimers

`func (o *RnaSequenceUpdate) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *RnaSequenceUpdate) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *RnaSequenceUpdate) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *RnaSequenceUpdate) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetSchemaId

`func (o *RnaSequenceUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *RnaSequenceUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *RnaSequenceUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *RnaSequenceUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTranslations

`func (o *RnaSequenceUpdate) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *RnaSequenceUpdate) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *RnaSequenceUpdate) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *RnaSequenceUpdate) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *RnaSequenceUpdate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *RnaSequenceUpdate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *RnaSequenceUpdate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *RnaSequenceUpdate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


