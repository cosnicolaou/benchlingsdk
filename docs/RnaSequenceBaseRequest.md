# RnaSequenceBaseRequest

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

## Methods

### NewRnaSequenceBaseRequest

`func NewRnaSequenceBaseRequest() *RnaSequenceBaseRequest`

NewRnaSequenceBaseRequest instantiates a new RnaSequenceBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaSequenceBaseRequestWithDefaults

`func NewRnaSequenceBaseRequestWithDefaults() *RnaSequenceBaseRequest`

NewRnaSequenceBaseRequestWithDefaults instantiates a new RnaSequenceBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *RnaSequenceBaseRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *RnaSequenceBaseRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *RnaSequenceBaseRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *RnaSequenceBaseRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *RnaSequenceBaseRequest) GetAnnotations() []RnaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RnaSequenceBaseRequest) GetAnnotationsOk() (*[]RnaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RnaSequenceBaseRequest) SetAnnotations(v []RnaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *RnaSequenceBaseRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *RnaSequenceBaseRequest) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *RnaSequenceBaseRequest) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *RnaSequenceBaseRequest) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *RnaSequenceBaseRequest) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *RnaSequenceBaseRequest) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *RnaSequenceBaseRequest) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *RnaSequenceBaseRequest) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *RnaSequenceBaseRequest) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetCustomFields

`func (o *RnaSequenceBaseRequest) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RnaSequenceBaseRequest) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RnaSequenceBaseRequest) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RnaSequenceBaseRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *RnaSequenceBaseRequest) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RnaSequenceBaseRequest) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RnaSequenceBaseRequest) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RnaSequenceBaseRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *RnaSequenceBaseRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RnaSequenceBaseRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RnaSequenceBaseRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RnaSequenceBaseRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIsCircular

`func (o *RnaSequenceBaseRequest) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *RnaSequenceBaseRequest) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *RnaSequenceBaseRequest) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.

### HasIsCircular

`func (o *RnaSequenceBaseRequest) HasIsCircular() bool`

HasIsCircular returns a boolean if a field has been set.

### GetName

`func (o *RnaSequenceBaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RnaSequenceBaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RnaSequenceBaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RnaSequenceBaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimers

`func (o *RnaSequenceBaseRequest) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *RnaSequenceBaseRequest) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *RnaSequenceBaseRequest) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *RnaSequenceBaseRequest) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetSchemaId

`func (o *RnaSequenceBaseRequest) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *RnaSequenceBaseRequest) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *RnaSequenceBaseRequest) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *RnaSequenceBaseRequest) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTranslations

`func (o *RnaSequenceBaseRequest) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *RnaSequenceBaseRequest) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *RnaSequenceBaseRequest) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *RnaSequenceBaseRequest) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


