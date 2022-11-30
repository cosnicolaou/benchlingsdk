# AaSequenceBaseRequest

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

## Methods

### NewAaSequenceBaseRequest

`func NewAaSequenceBaseRequest() *AaSequenceBaseRequest`

NewAaSequenceBaseRequest instantiates a new AaSequenceBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaSequenceBaseRequestWithDefaults

`func NewAaSequenceBaseRequestWithDefaults() *AaSequenceBaseRequest`

NewAaSequenceBaseRequestWithDefaults instantiates a new AaSequenceBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *AaSequenceBaseRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AaSequenceBaseRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AaSequenceBaseRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AaSequenceBaseRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAminoAcids

`func (o *AaSequenceBaseRequest) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *AaSequenceBaseRequest) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *AaSequenceBaseRequest) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *AaSequenceBaseRequest) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetAnnotations

`func (o *AaSequenceBaseRequest) GetAnnotations() []AaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AaSequenceBaseRequest) GetAnnotationsOk() (*[]AaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AaSequenceBaseRequest) SetAnnotations(v []AaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AaSequenceBaseRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *AaSequenceBaseRequest) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *AaSequenceBaseRequest) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *AaSequenceBaseRequest) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *AaSequenceBaseRequest) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *AaSequenceBaseRequest) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AaSequenceBaseRequest) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AaSequenceBaseRequest) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AaSequenceBaseRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *AaSequenceBaseRequest) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AaSequenceBaseRequest) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AaSequenceBaseRequest) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AaSequenceBaseRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *AaSequenceBaseRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AaSequenceBaseRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AaSequenceBaseRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AaSequenceBaseRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *AaSequenceBaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AaSequenceBaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AaSequenceBaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AaSequenceBaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *AaSequenceBaseRequest) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *AaSequenceBaseRequest) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *AaSequenceBaseRequest) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *AaSequenceBaseRequest) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


