# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Entry in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableEntryArchiveRecord**](EntryArchiveRecord.md) |  | [optional] 
**AssignedReviewers** | Pointer to [**[]UserSummary**](UserSummary.md) | Array of users assigned to review the entry, if any.  | [optional] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) | Array of UserSummary Resources of the authors of the entry. This defaults to the creator but can be manually changed.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the entry was created at | [optional] [readonly] 
**Creator** | Pointer to [**EntryCreator**](EntryCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**Days** | Pointer to [**[]EntryDay**](EntryDay.md) | Array of day objects. Each day object has a date field (string) and notes field (array of notes, expand further for details on note types).  | [optional] 
**DisplayId** | Pointer to **string** | User-friendly ID of the entry | [optional] 
**EntryTemplateId** | Pointer to **NullableString** | ID of the Entry Template this Entry was created from | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder that contains the entry | [optional] 
**Id** | Pointer to **string** | ID of the entry | [optional] 
**ModifiedAt** | Pointer to **string** | DateTime the entry was last modified | [optional] 
**Name** | Pointer to **string** | Title of the entry | [optional] 
**ReviewRecord** | Pointer to **map[string]interface{}** | Review record if set | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty4**](SchemaProperty4.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the entry | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *Entry) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *Entry) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *Entry) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *Entry) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *Entry) GetArchiveRecord() EntryArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Entry) GetArchiveRecordOk() (*EntryArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Entry) SetArchiveRecord(v EntryArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Entry) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Entry) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Entry) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetAssignedReviewers

`func (o *Entry) GetAssignedReviewers() []UserSummary`

GetAssignedReviewers returns the AssignedReviewers field if non-nil, zero value otherwise.

### GetAssignedReviewersOk

`func (o *Entry) GetAssignedReviewersOk() (*[]UserSummary, bool)`

GetAssignedReviewersOk returns a tuple with the AssignedReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedReviewers

`func (o *Entry) SetAssignedReviewers(v []UserSummary)`

SetAssignedReviewers sets AssignedReviewers field to given value.

### HasAssignedReviewers

`func (o *Entry) HasAssignedReviewers() bool`

HasAssignedReviewers returns a boolean if a field has been set.

### GetAuthors

`func (o *Entry) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Entry) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Entry) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *Entry) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Entry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Entry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Entry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Entry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Entry) GetCreator() EntryCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Entry) GetCreatorOk() (*EntryCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Entry) SetCreator(v EntryCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Entry) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *Entry) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Entry) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Entry) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Entry) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDays

`func (o *Entry) GetDays() []EntryDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Entry) GetDaysOk() (*[]EntryDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Entry) SetDays(v []EntryDay)`

SetDays sets Days field to given value.

### HasDays

`func (o *Entry) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDisplayId

`func (o *Entry) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Entry) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Entry) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Entry) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetEntryTemplateId

`func (o *Entry) GetEntryTemplateId() string`

GetEntryTemplateId returns the EntryTemplateId field if non-nil, zero value otherwise.

### GetEntryTemplateIdOk

`func (o *Entry) GetEntryTemplateIdOk() (*string, bool)`

GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTemplateId

`func (o *Entry) SetEntryTemplateId(v string)`

SetEntryTemplateId sets EntryTemplateId field to given value.

### HasEntryTemplateId

`func (o *Entry) HasEntryTemplateId() bool`

HasEntryTemplateId returns a boolean if a field has been set.

### SetEntryTemplateIdNil

`func (o *Entry) SetEntryTemplateIdNil(b bool)`

 SetEntryTemplateIdNil sets the value for EntryTemplateId to be an explicit nil

### UnsetEntryTemplateId
`func (o *Entry) UnsetEntryTemplateId()`

UnsetEntryTemplateId ensures that no value is present for EntryTemplateId, not even an explicit nil
### GetFields

`func (o *Entry) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Entry) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Entry) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Entry) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *Entry) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Entry) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Entry) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Entry) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *Entry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Entry) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Entry) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Entry) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Entry) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Entry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReviewRecord

`func (o *Entry) GetReviewRecord() map[string]interface{}`

GetReviewRecord returns the ReviewRecord field if non-nil, zero value otherwise.

### GetReviewRecordOk

`func (o *Entry) GetReviewRecordOk() (*map[string]interface{}, bool)`

GetReviewRecordOk returns a tuple with the ReviewRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRecord

`func (o *Entry) SetReviewRecord(v map[string]interface{})`

SetReviewRecord sets ReviewRecord field to given value.

### HasReviewRecord

`func (o *Entry) HasReviewRecord() bool`

HasReviewRecord returns a boolean if a field has been set.

### SetReviewRecordNil

`func (o *Entry) SetReviewRecordNil(b bool)`

 SetReviewRecordNil sets the value for ReviewRecord to be an explicit nil

### UnsetReviewRecord
`func (o *Entry) UnsetReviewRecord()`

UnsetReviewRecord ensures that no value is present for ReviewRecord, not even an explicit nil
### GetSchema

`func (o *Entry) GetSchema() SchemaProperty4`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Entry) GetSchemaOk() (*SchemaProperty4, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Entry) SetSchema(v SchemaProperty4)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Entry) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Entry) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Entry) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *Entry) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Entry) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Entry) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Entry) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


