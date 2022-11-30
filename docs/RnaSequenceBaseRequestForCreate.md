# RnaSequenceBaseRequestForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the RNA sequence | [optional] 
**Annotations** | Pointer to [**[]RnaAnnotation**](RnaAnnotation.md) | Annotations to create on the RNA sequence.  | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the RNA sequence&#39;s authors. | [optional] 
**Bases** | **string** | Base pairs for the RNA sequence.  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the RNA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the RNA sequence. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the RNA sequence.  | [optional] 
**IsCircular** | Pointer to **bool** | Whether the RNA sequence is circular or linear. RNA sequences can only be linear  | [optional] 
**Name** | **string** | Name of the RNA sequence.  | 
**Primers** | Pointer to [**[]Primer**](Primer.md) |  | [optional] 
**SchemaId** | Pointer to **string** | ID of the RNA sequence&#39;s schema.  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) | Translations to create on the RNA sequence. Translations are specified by either a combination of &#39;start&#39; and &#39;end&#39; fields, or a list of regions. Both cannot be provided.  | [optional] 

## Methods

### NewRnaSequenceBaseRequestForCreate

`func NewRnaSequenceBaseRequestForCreate(bases string, name string, ) *RnaSequenceBaseRequestForCreate`

NewRnaSequenceBaseRequestForCreate instantiates a new RnaSequenceBaseRequestForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaSequenceBaseRequestForCreateWithDefaults

`func NewRnaSequenceBaseRequestForCreateWithDefaults() *RnaSequenceBaseRequestForCreate`

NewRnaSequenceBaseRequestForCreateWithDefaults instantiates a new RnaSequenceBaseRequestForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *RnaSequenceBaseRequestForCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *RnaSequenceBaseRequestForCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *RnaSequenceBaseRequestForCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *RnaSequenceBaseRequestForCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *RnaSequenceBaseRequestForCreate) GetAnnotations() []RnaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RnaSequenceBaseRequestForCreate) GetAnnotationsOk() (*[]RnaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RnaSequenceBaseRequestForCreate) SetAnnotations(v []RnaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *RnaSequenceBaseRequestForCreate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *RnaSequenceBaseRequestForCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *RnaSequenceBaseRequestForCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *RnaSequenceBaseRequestForCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *RnaSequenceBaseRequestForCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *RnaSequenceBaseRequestForCreate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *RnaSequenceBaseRequestForCreate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *RnaSequenceBaseRequestForCreate) SetBases(v string)`

SetBases sets Bases field to given value.


### GetCustomFields

`func (o *RnaSequenceBaseRequestForCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RnaSequenceBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RnaSequenceBaseRequestForCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RnaSequenceBaseRequestForCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *RnaSequenceBaseRequestForCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RnaSequenceBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RnaSequenceBaseRequestForCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RnaSequenceBaseRequestForCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *RnaSequenceBaseRequestForCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RnaSequenceBaseRequestForCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RnaSequenceBaseRequestForCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RnaSequenceBaseRequestForCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIsCircular

`func (o *RnaSequenceBaseRequestForCreate) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *RnaSequenceBaseRequestForCreate) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *RnaSequenceBaseRequestForCreate) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.

### HasIsCircular

`func (o *RnaSequenceBaseRequestForCreate) HasIsCircular() bool`

HasIsCircular returns a boolean if a field has been set.

### GetName

`func (o *RnaSequenceBaseRequestForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RnaSequenceBaseRequestForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RnaSequenceBaseRequestForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrimers

`func (o *RnaSequenceBaseRequestForCreate) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *RnaSequenceBaseRequestForCreate) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *RnaSequenceBaseRequestForCreate) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *RnaSequenceBaseRequestForCreate) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetSchemaId

`func (o *RnaSequenceBaseRequestForCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *RnaSequenceBaseRequestForCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *RnaSequenceBaseRequestForCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *RnaSequenceBaseRequestForCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTranslations

`func (o *RnaSequenceBaseRequestForCreate) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *RnaSequenceBaseRequestForCreate) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *RnaSequenceBaseRequestForCreate) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *RnaSequenceBaseRequestForCreate) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


