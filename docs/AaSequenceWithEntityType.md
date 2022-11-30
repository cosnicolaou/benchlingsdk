# AaSequenceWithEntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Array of aliases | [optional] 
**AminoAcids** | Pointer to **string** | Amino acids of the AA sequence. | [optional] 
**Annotations** | Pointer to [**[]AaAnnotation**](AaAnnotation.md) | Array of annotation objects on the AA sequence. | [optional] 
**ApiURL** | Pointer to **string** | The canonical url of the AA Sequence in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the AA sequence was created. | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields set on the AA sequence. | [optional] 
**EntityRegistryId** | Pointer to **NullableString** | Registry ID of the AA sequence if registered. | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **NullableString** | ID of the folder that contains the AA sequence. | [optional] 
**Id** | Pointer to **string** | ID of the AA sequence. | [optional] 
**Length** | Pointer to **int32** | Number of amino acids in the AA sequence. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the AA sequence was last modified. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the AA sequence. | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** | Registry the AA sequence is registered in. | [optional] 
**Schema** | Pointer to [**NullableAaSequenceSchema**](AaSequenceSchema.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the protein. | [optional] [readonly] 
**EntityType** | Pointer to **string** |  | [optional] 

## Methods

### NewAaSequenceWithEntityType

`func NewAaSequenceWithEntityType() *AaSequenceWithEntityType`

NewAaSequenceWithEntityType instantiates a new AaSequenceWithEntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaSequenceWithEntityTypeWithDefaults

`func NewAaSequenceWithEntityTypeWithDefaults() *AaSequenceWithEntityType`

NewAaSequenceWithEntityTypeWithDefaults instantiates a new AaSequenceWithEntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *AaSequenceWithEntityType) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AaSequenceWithEntityType) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AaSequenceWithEntityType) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AaSequenceWithEntityType) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAminoAcids

`func (o *AaSequenceWithEntityType) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *AaSequenceWithEntityType) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *AaSequenceWithEntityType) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *AaSequenceWithEntityType) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetAnnotations

`func (o *AaSequenceWithEntityType) GetAnnotations() []AaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AaSequenceWithEntityType) GetAnnotationsOk() (*[]AaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AaSequenceWithEntityType) SetAnnotations(v []AaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AaSequenceWithEntityType) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetApiURL

`func (o *AaSequenceWithEntityType) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AaSequenceWithEntityType) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AaSequenceWithEntityType) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *AaSequenceWithEntityType) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *AaSequenceWithEntityType) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *AaSequenceWithEntityType) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *AaSequenceWithEntityType) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *AaSequenceWithEntityType) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *AaSequenceWithEntityType) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *AaSequenceWithEntityType) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *AaSequenceWithEntityType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AaSequenceWithEntityType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AaSequenceWithEntityType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AaSequenceWithEntityType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *AaSequenceWithEntityType) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *AaSequenceWithEntityType) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *AaSequenceWithEntityType) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *AaSequenceWithEntityType) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *AaSequenceWithEntityType) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AaSequenceWithEntityType) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AaSequenceWithEntityType) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AaSequenceWithEntityType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *AaSequenceWithEntityType) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *AaSequenceWithEntityType) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *AaSequenceWithEntityType) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *AaSequenceWithEntityType) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *AaSequenceWithEntityType) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *AaSequenceWithEntityType) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *AaSequenceWithEntityType) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AaSequenceWithEntityType) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AaSequenceWithEntityType) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AaSequenceWithEntityType) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *AaSequenceWithEntityType) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AaSequenceWithEntityType) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AaSequenceWithEntityType) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AaSequenceWithEntityType) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *AaSequenceWithEntityType) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *AaSequenceWithEntityType) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *AaSequenceWithEntityType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AaSequenceWithEntityType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AaSequenceWithEntityType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AaSequenceWithEntityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLength

`func (o *AaSequenceWithEntityType) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AaSequenceWithEntityType) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AaSequenceWithEntityType) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *AaSequenceWithEntityType) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AaSequenceWithEntityType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AaSequenceWithEntityType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AaSequenceWithEntityType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AaSequenceWithEntityType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *AaSequenceWithEntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AaSequenceWithEntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AaSequenceWithEntityType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AaSequenceWithEntityType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *AaSequenceWithEntityType) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *AaSequenceWithEntityType) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *AaSequenceWithEntityType) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *AaSequenceWithEntityType) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *AaSequenceWithEntityType) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *AaSequenceWithEntityType) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *AaSequenceWithEntityType) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *AaSequenceWithEntityType) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *AaSequenceWithEntityType) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *AaSequenceWithEntityType) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *AaSequenceWithEntityType) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *AaSequenceWithEntityType) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *AaSequenceWithEntityType) GetSchema() AaSequenceSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AaSequenceWithEntityType) GetSchemaOk() (*AaSequenceSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AaSequenceWithEntityType) SetSchema(v AaSequenceSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AaSequenceWithEntityType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *AaSequenceWithEntityType) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *AaSequenceWithEntityType) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *AaSequenceWithEntityType) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *AaSequenceWithEntityType) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *AaSequenceWithEntityType) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *AaSequenceWithEntityType) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetEntityType

`func (o *AaSequenceWithEntityType) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AaSequenceWithEntityType) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AaSequenceWithEntityType) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AaSequenceWithEntityType) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


