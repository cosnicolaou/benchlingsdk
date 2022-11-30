# DnaSequenceBaseRequestForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the DNA sequence | [optional] 
**Annotations** | Pointer to [**[]DnaAnnotation**](DnaAnnotation.md) | Annotations to create on the DNA sequence.  | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the DNA sequence&#39;s authors. | [optional] 
**Bases** | **string** | Base pairs for the DNA sequence.  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the DNA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the DNA sequence. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the DNA sequence.  | [optional] 
**IsCircular** | **bool** | Whether the DNA sequence is circular or linear.  | 
**Name** | **string** | Name of the DNA sequence.  | 
**Primers** | Pointer to [**[]Primer**](Primer.md) |  | [optional] 
**SchemaId** | Pointer to **string** | ID of the DNA sequence&#39;s schema.  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) | Translations to create on the DNA sequence. Translations are specified by either a combination of &#39;start&#39; and &#39;end&#39; fields, or a list of regions. Both cannot be provided.  | [optional] 

## Methods

### NewDnaSequenceBaseRequestForCreate

`func NewDnaSequenceBaseRequestForCreate(bases string, isCircular bool, name string, ) *DnaSequenceBaseRequestForCreate`

NewDnaSequenceBaseRequestForCreate instantiates a new DnaSequenceBaseRequestForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaSequenceBaseRequestForCreateWithDefaults

`func NewDnaSequenceBaseRequestForCreateWithDefaults() *DnaSequenceBaseRequestForCreate`

NewDnaSequenceBaseRequestForCreateWithDefaults instantiates a new DnaSequenceBaseRequestForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *DnaSequenceBaseRequestForCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *DnaSequenceBaseRequestForCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *DnaSequenceBaseRequestForCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *DnaSequenceBaseRequestForCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *DnaSequenceBaseRequestForCreate) GetAnnotations() []DnaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DnaSequenceBaseRequestForCreate) GetAnnotationsOk() (*[]DnaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DnaSequenceBaseRequestForCreate) SetAnnotations(v []DnaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DnaSequenceBaseRequestForCreate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *DnaSequenceBaseRequestForCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *DnaSequenceBaseRequestForCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *DnaSequenceBaseRequestForCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *DnaSequenceBaseRequestForCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *DnaSequenceBaseRequestForCreate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *DnaSequenceBaseRequestForCreate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *DnaSequenceBaseRequestForCreate) SetBases(v string)`

SetBases sets Bases field to given value.


### GetCustomFields

`func (o *DnaSequenceBaseRequestForCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DnaSequenceBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DnaSequenceBaseRequestForCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DnaSequenceBaseRequestForCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *DnaSequenceBaseRequestForCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DnaSequenceBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DnaSequenceBaseRequestForCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DnaSequenceBaseRequestForCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *DnaSequenceBaseRequestForCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DnaSequenceBaseRequestForCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DnaSequenceBaseRequestForCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DnaSequenceBaseRequestForCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIsCircular

`func (o *DnaSequenceBaseRequestForCreate) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *DnaSequenceBaseRequestForCreate) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *DnaSequenceBaseRequestForCreate) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.


### GetName

`func (o *DnaSequenceBaseRequestForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaSequenceBaseRequestForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaSequenceBaseRequestForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrimers

`func (o *DnaSequenceBaseRequestForCreate) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *DnaSequenceBaseRequestForCreate) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *DnaSequenceBaseRequestForCreate) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *DnaSequenceBaseRequestForCreate) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetSchemaId

`func (o *DnaSequenceBaseRequestForCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *DnaSequenceBaseRequestForCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *DnaSequenceBaseRequestForCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *DnaSequenceBaseRequestForCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTranslations

`func (o *DnaSequenceBaseRequestForCreate) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *DnaSequenceBaseRequestForCreate) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *DnaSequenceBaseRequestForCreate) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *DnaSequenceBaseRequestForCreate) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


