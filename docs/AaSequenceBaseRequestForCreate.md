# AaSequenceBaseRequestForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the AA sequence | [optional] 
**AminoAcids** | **string** | Amino acids for the AA sequence.  | 
**Annotations** | Pointer to [**[]AaAnnotation**](AaAnnotation.md) | Annotations to create on the AA sequence.  | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the AA sequence&#39;s authors. | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the AA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the AA sequence. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the AA sequence.  | [optional] 
**Name** | **string** | Name of the AA sequence.  | 
**SchemaId** | Pointer to **string** | ID of the AA sequence&#39;s schema.  | [optional] 

## Methods

### NewAaSequenceBaseRequestForCreate

`func NewAaSequenceBaseRequestForCreate(aminoAcids string, name string, ) *AaSequenceBaseRequestForCreate`

NewAaSequenceBaseRequestForCreate instantiates a new AaSequenceBaseRequestForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaSequenceBaseRequestForCreateWithDefaults

`func NewAaSequenceBaseRequestForCreateWithDefaults() *AaSequenceBaseRequestForCreate`

NewAaSequenceBaseRequestForCreateWithDefaults instantiates a new AaSequenceBaseRequestForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *AaSequenceBaseRequestForCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AaSequenceBaseRequestForCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AaSequenceBaseRequestForCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AaSequenceBaseRequestForCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAminoAcids

`func (o *AaSequenceBaseRequestForCreate) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *AaSequenceBaseRequestForCreate) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *AaSequenceBaseRequestForCreate) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.


### GetAnnotations

`func (o *AaSequenceBaseRequestForCreate) GetAnnotations() []AaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AaSequenceBaseRequestForCreate) GetAnnotationsOk() (*[]AaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AaSequenceBaseRequestForCreate) SetAnnotations(v []AaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AaSequenceBaseRequestForCreate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *AaSequenceBaseRequestForCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *AaSequenceBaseRequestForCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *AaSequenceBaseRequestForCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *AaSequenceBaseRequestForCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *AaSequenceBaseRequestForCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AaSequenceBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AaSequenceBaseRequestForCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AaSequenceBaseRequestForCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *AaSequenceBaseRequestForCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AaSequenceBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AaSequenceBaseRequestForCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AaSequenceBaseRequestForCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *AaSequenceBaseRequestForCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AaSequenceBaseRequestForCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AaSequenceBaseRequestForCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AaSequenceBaseRequestForCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *AaSequenceBaseRequestForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AaSequenceBaseRequestForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AaSequenceBaseRequestForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *AaSequenceBaseRequestForCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *AaSequenceBaseRequestForCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *AaSequenceBaseRequestForCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *AaSequenceBaseRequestForCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


