# GenericEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** |  | [optional] 
**ApiURL** | Pointer to **string** |  | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) | Array of UserSummary Resources of the authors of the entry. This defaults to the creator but can be manually changed.  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to [**GenericEntityCreator**](GenericEntityCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**EntityRegistryId** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty3**](SchemaProperty3.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewGenericEntity

`func NewGenericEntity() *GenericEntity`

NewGenericEntity instantiates a new GenericEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericEntityWithDefaults

`func NewGenericEntityWithDefaults() *GenericEntity`

NewGenericEntityWithDefaults instantiates a new GenericEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *GenericEntity) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GenericEntity) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GenericEntity) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GenericEntity) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetApiURL

`func (o *GenericEntity) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *GenericEntity) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *GenericEntity) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *GenericEntity) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *GenericEntity) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *GenericEntity) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *GenericEntity) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *GenericEntity) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *GenericEntity) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *GenericEntity) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetAuthors

`func (o *GenericEntity) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *GenericEntity) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *GenericEntity) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *GenericEntity) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GenericEntity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GenericEntity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GenericEntity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GenericEntity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *GenericEntity) GetCreator() GenericEntityCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GenericEntity) GetCreatorOk() (*GenericEntityCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GenericEntity) SetCreator(v GenericEntityCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GenericEntity) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *GenericEntity) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GenericEntity) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GenericEntity) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GenericEntity) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *GenericEntity) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *GenericEntity) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *GenericEntity) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *GenericEntity) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *GenericEntity) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *GenericEntity) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *GenericEntity) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GenericEntity) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GenericEntity) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GenericEntity) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *GenericEntity) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GenericEntity) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GenericEntity) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GenericEntity) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *GenericEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GenericEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *GenericEntity) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GenericEntity) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GenericEntity) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *GenericEntity) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *GenericEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *GenericEntity) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *GenericEntity) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *GenericEntity) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *GenericEntity) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *GenericEntity) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *GenericEntity) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *GenericEntity) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *GenericEntity) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *GenericEntity) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *GenericEntity) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *GenericEntity) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *GenericEntity) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *GenericEntity) GetSchema() SchemaProperty3`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GenericEntity) GetSchemaOk() (*SchemaProperty3, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GenericEntity) SetSchema(v SchemaProperty3)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *GenericEntity) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *GenericEntity) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *GenericEntity) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *GenericEntity) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *GenericEntity) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *GenericEntity) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *GenericEntity) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


