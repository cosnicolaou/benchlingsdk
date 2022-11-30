# CustomEntityWithEntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** |  | [optional] 
**ApiURL** | Pointer to **string** | The canonical url of the Custom Entity in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to [**CustomEntityCreator**](CustomEntityCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**EntityRegistryId** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to [**NullableCustomEntitySchema**](CustomEntitySchema.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 
**EntityType** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomEntityWithEntityType

`func NewCustomEntityWithEntityType() *CustomEntityWithEntityType`

NewCustomEntityWithEntityType instantiates a new CustomEntityWithEntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityWithEntityTypeWithDefaults

`func NewCustomEntityWithEntityTypeWithDefaults() *CustomEntityWithEntityType`

NewCustomEntityWithEntityTypeWithDefaults instantiates a new CustomEntityWithEntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *CustomEntityWithEntityType) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CustomEntityWithEntityType) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CustomEntityWithEntityType) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CustomEntityWithEntityType) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetApiURL

`func (o *CustomEntityWithEntityType) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *CustomEntityWithEntityType) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *CustomEntityWithEntityType) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *CustomEntityWithEntityType) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *CustomEntityWithEntityType) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *CustomEntityWithEntityType) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *CustomEntityWithEntityType) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *CustomEntityWithEntityType) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *CustomEntityWithEntityType) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *CustomEntityWithEntityType) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetAuthors

`func (o *CustomEntityWithEntityType) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CustomEntityWithEntityType) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CustomEntityWithEntityType) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *CustomEntityWithEntityType) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomEntityWithEntityType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomEntityWithEntityType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomEntityWithEntityType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomEntityWithEntityType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *CustomEntityWithEntityType) GetCreator() CustomEntityCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CustomEntityWithEntityType) GetCreatorOk() (*CustomEntityCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CustomEntityWithEntityType) SetCreator(v CustomEntityCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CustomEntityWithEntityType) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomEntityWithEntityType) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomEntityWithEntityType) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomEntityWithEntityType) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomEntityWithEntityType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *CustomEntityWithEntityType) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *CustomEntityWithEntityType) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *CustomEntityWithEntityType) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *CustomEntityWithEntityType) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *CustomEntityWithEntityType) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *CustomEntityWithEntityType) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *CustomEntityWithEntityType) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomEntityWithEntityType) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomEntityWithEntityType) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomEntityWithEntityType) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *CustomEntityWithEntityType) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CustomEntityWithEntityType) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CustomEntityWithEntityType) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CustomEntityWithEntityType) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *CustomEntityWithEntityType) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *CustomEntityWithEntityType) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *CustomEntityWithEntityType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEntityWithEntityType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEntityWithEntityType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomEntityWithEntityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *CustomEntityWithEntityType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CustomEntityWithEntityType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CustomEntityWithEntityType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CustomEntityWithEntityType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *CustomEntityWithEntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEntityWithEntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEntityWithEntityType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomEntityWithEntityType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *CustomEntityWithEntityType) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *CustomEntityWithEntityType) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *CustomEntityWithEntityType) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *CustomEntityWithEntityType) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *CustomEntityWithEntityType) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *CustomEntityWithEntityType) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *CustomEntityWithEntityType) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *CustomEntityWithEntityType) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *CustomEntityWithEntityType) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *CustomEntityWithEntityType) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *CustomEntityWithEntityType) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *CustomEntityWithEntityType) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *CustomEntityWithEntityType) GetSchema() CustomEntitySchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CustomEntityWithEntityType) GetSchemaOk() (*CustomEntitySchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CustomEntityWithEntityType) SetSchema(v CustomEntitySchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CustomEntityWithEntityType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *CustomEntityWithEntityType) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *CustomEntityWithEntityType) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *CustomEntityWithEntityType) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *CustomEntityWithEntityType) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *CustomEntityWithEntityType) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *CustomEntityWithEntityType) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetEntityType

`func (o *CustomEntityWithEntityType) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomEntityWithEntityType) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomEntityWithEntityType) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CustomEntityWithEntityType) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


