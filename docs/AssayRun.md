# AssayRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Run in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**EntryId** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsReviewed** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty1**](SchemaProperty1.md) |  | [optional] 
**ValidationComment** | Pointer to **NullableString** |  | [optional] 
**ValidationStatus** | Pointer to [**AssayRunValidationStatus**](AssayRunValidationStatus.md) |  | [optional] 

## Methods

### NewAssayRun

`func NewAssayRun() *AssayRun`

NewAssayRun instantiates a new AssayRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunWithDefaults

`func NewAssayRunWithDefaults() *AssayRun`

NewAssayRunWithDefaults instantiates a new AssayRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *AssayRun) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AssayRun) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AssayRun) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *AssayRun) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *AssayRun) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *AssayRun) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *AssayRun) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *AssayRun) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *AssayRun) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *AssayRun) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *AssayRun) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssayRun) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssayRun) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssayRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *AssayRun) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *AssayRun) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *AssayRun) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *AssayRun) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEntryId

`func (o *AssayRun) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *AssayRun) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *AssayRun) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *AssayRun) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### SetEntryIdNil

`func (o *AssayRun) SetEntryIdNil(b bool)`

 SetEntryIdNil sets the value for EntryId to be an explicit nil

### UnsetEntryId
`func (o *AssayRun) UnsetEntryId()`

UnsetEntryId ensures that no value is present for EntryId, not even an explicit nil
### GetFields

`func (o *AssayRun) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AssayRun) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AssayRun) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AssayRun) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *AssayRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsReviewed

`func (o *AssayRun) GetIsReviewed() bool`

GetIsReviewed returns the IsReviewed field if non-nil, zero value otherwise.

### GetIsReviewedOk

`func (o *AssayRun) GetIsReviewedOk() (*bool, bool)`

GetIsReviewedOk returns a tuple with the IsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReviewed

`func (o *AssayRun) SetIsReviewed(v bool)`

SetIsReviewed sets IsReviewed field to given value.

### HasIsReviewed

`func (o *AssayRun) HasIsReviewed() bool`

HasIsReviewed returns a boolean if a field has been set.

### GetProjectId

`func (o *AssayRun) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssayRun) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssayRun) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AssayRun) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *AssayRun) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AssayRun) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSchema

`func (o *AssayRun) GetSchema() SchemaProperty1`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AssayRun) GetSchemaOk() (*SchemaProperty1, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AssayRun) SetSchema(v SchemaProperty1)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AssayRun) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *AssayRun) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *AssayRun) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetValidationComment

`func (o *AssayRun) GetValidationComment() string`

GetValidationComment returns the ValidationComment field if non-nil, zero value otherwise.

### GetValidationCommentOk

`func (o *AssayRun) GetValidationCommentOk() (*string, bool)`

GetValidationCommentOk returns a tuple with the ValidationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationComment

`func (o *AssayRun) SetValidationComment(v string)`

SetValidationComment sets ValidationComment field to given value.

### HasValidationComment

`func (o *AssayRun) HasValidationComment() bool`

HasValidationComment returns a boolean if a field has been set.

### SetValidationCommentNil

`func (o *AssayRun) SetValidationCommentNil(b bool)`

 SetValidationCommentNil sets the value for ValidationComment to be an explicit nil

### UnsetValidationComment
`func (o *AssayRun) UnsetValidationComment()`

UnsetValidationComment ensures that no value is present for ValidationComment, not even an explicit nil
### GetValidationStatus

`func (o *AssayRun) GetValidationStatus() AssayRunValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *AssayRun) GetValidationStatusOk() (*AssayRunValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *AssayRun) SetValidationStatus(v AssayRunValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *AssayRun) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


