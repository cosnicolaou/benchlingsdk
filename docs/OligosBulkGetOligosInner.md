# OligosBulkGetOligosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Array of aliases | [optional] 
**ApiURL** | Pointer to **string** |  | [optional] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Bases** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Oligo was created. | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields set on the Oligo. | [optional] 
**EntityRegistryId** | Pointer to **NullableString** | Registry ID of the Oligo if registered. | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **NullableString** | ID of the folder that contains the Oligo. | [optional] 
**Id** | Pointer to **string** | ID of the Oligo. | [optional] 
**Length** | Pointer to **int32** | Number of bases in the Oligo. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Oligo was last modified. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Oligo. | [optional] 
**NucleotideType** | Pointer to **string** |  | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** | Registry the Oligo is registered in. | [optional] 
**Schema** | Pointer to [**NullableAaSequenceSchema**](AaSequenceSchema.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the Oligo. | [optional] [readonly] 

## Methods

### NewOligosBulkGetOligosInner

`func NewOligosBulkGetOligosInner() *OligosBulkGetOligosInner`

NewOligosBulkGetOligosInner instantiates a new OligosBulkGetOligosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOligosBulkGetOligosInnerWithDefaults

`func NewOligosBulkGetOligosInnerWithDefaults() *OligosBulkGetOligosInner`

NewOligosBulkGetOligosInnerWithDefaults instantiates a new OligosBulkGetOligosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *OligosBulkGetOligosInner) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *OligosBulkGetOligosInner) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *OligosBulkGetOligosInner) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *OligosBulkGetOligosInner) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetApiURL

`func (o *OligosBulkGetOligosInner) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *OligosBulkGetOligosInner) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *OligosBulkGetOligosInner) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *OligosBulkGetOligosInner) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *OligosBulkGetOligosInner) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *OligosBulkGetOligosInner) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *OligosBulkGetOligosInner) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *OligosBulkGetOligosInner) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *OligosBulkGetOligosInner) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *OligosBulkGetOligosInner) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBases

`func (o *OligosBulkGetOligosInner) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *OligosBulkGetOligosInner) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *OligosBulkGetOligosInner) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *OligosBulkGetOligosInner) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OligosBulkGetOligosInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OligosBulkGetOligosInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OligosBulkGetOligosInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OligosBulkGetOligosInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *OligosBulkGetOligosInner) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *OligosBulkGetOligosInner) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *OligosBulkGetOligosInner) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *OligosBulkGetOligosInner) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *OligosBulkGetOligosInner) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OligosBulkGetOligosInner) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OligosBulkGetOligosInner) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OligosBulkGetOligosInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *OligosBulkGetOligosInner) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *OligosBulkGetOligosInner) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *OligosBulkGetOligosInner) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *OligosBulkGetOligosInner) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *OligosBulkGetOligosInner) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *OligosBulkGetOligosInner) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *OligosBulkGetOligosInner) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *OligosBulkGetOligosInner) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *OligosBulkGetOligosInner) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *OligosBulkGetOligosInner) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *OligosBulkGetOligosInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *OligosBulkGetOligosInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *OligosBulkGetOligosInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *OligosBulkGetOligosInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *OligosBulkGetOligosInner) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *OligosBulkGetOligosInner) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *OligosBulkGetOligosInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OligosBulkGetOligosInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OligosBulkGetOligosInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OligosBulkGetOligosInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLength

`func (o *OligosBulkGetOligosInner) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *OligosBulkGetOligosInner) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *OligosBulkGetOligosInner) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *OligosBulkGetOligosInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetModifiedAt

`func (o *OligosBulkGetOligosInner) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *OligosBulkGetOligosInner) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *OligosBulkGetOligosInner) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *OligosBulkGetOligosInner) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *OligosBulkGetOligosInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OligosBulkGetOligosInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OligosBulkGetOligosInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OligosBulkGetOligosInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNucleotideType

`func (o *OligosBulkGetOligosInner) GetNucleotideType() string`

GetNucleotideType returns the NucleotideType field if non-nil, zero value otherwise.

### GetNucleotideTypeOk

`func (o *OligosBulkGetOligosInner) GetNucleotideTypeOk() (*string, bool)`

GetNucleotideTypeOk returns a tuple with the NucleotideType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotideType

`func (o *OligosBulkGetOligosInner) SetNucleotideType(v string)`

SetNucleotideType sets NucleotideType field to given value.

### HasNucleotideType

`func (o *OligosBulkGetOligosInner) HasNucleotideType() bool`

HasNucleotideType returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *OligosBulkGetOligosInner) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *OligosBulkGetOligosInner) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *OligosBulkGetOligosInner) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *OligosBulkGetOligosInner) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *OligosBulkGetOligosInner) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *OligosBulkGetOligosInner) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *OligosBulkGetOligosInner) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *OligosBulkGetOligosInner) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *OligosBulkGetOligosInner) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *OligosBulkGetOligosInner) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *OligosBulkGetOligosInner) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *OligosBulkGetOligosInner) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *OligosBulkGetOligosInner) GetSchema() AaSequenceSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OligosBulkGetOligosInner) GetSchemaOk() (*AaSequenceSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OligosBulkGetOligosInner) SetSchema(v AaSequenceSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OligosBulkGetOligosInner) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *OligosBulkGetOligosInner) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *OligosBulkGetOligosInner) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *OligosBulkGetOligosInner) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *OligosBulkGetOligosInner) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *OligosBulkGetOligosInner) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *OligosBulkGetOligosInner) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


