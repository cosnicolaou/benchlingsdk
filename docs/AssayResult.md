# AssayResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAssayResultArchiveRecord**](AssayResultArchiveRecord.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the result was created | [optional] [readonly] 
**Creator** | Pointer to [**AssayResultCreator**](AssayResultCreator.md) |  | [optional] 
**EntryId** | Pointer to **NullableString** | ID of the entry that this result is attached to | [optional] 
**FieldValidation** | Pointer to [**map[string]UserValidation**](UserValidation.md) | Object mapping field names to a UserValidation Resource object for that field. To **set** validation for a result, you *must* use this object.  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the result | [optional] 
**IsReviewed** | Pointer to **bool** | Whether or not this result is attached to an accepted entry | [optional] 
**ProjectId** | Pointer to **NullableString** | ID of the project to insert the result into | [optional] 
**Schema** | Pointer to [**SchemaProperty**](SchemaProperty.md) |  | [optional] 
**ValidationComment** | Pointer to **string** |  | [optional] [readonly] 
**ValidationStatus** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAssayResult

`func NewAssayResult() *AssayResult`

NewAssayResult instantiates a new AssayResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayResultWithDefaults

`func NewAssayResultWithDefaults() *AssayResult`

NewAssayResultWithDefaults instantiates a new AssayResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *AssayResult) GetArchiveRecord() AssayResultArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *AssayResult) GetArchiveRecordOk() (*AssayResultArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *AssayResult) SetArchiveRecord(v AssayResultArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *AssayResult) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *AssayResult) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *AssayResult) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *AssayResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssayResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssayResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssayResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *AssayResult) GetCreator() AssayResultCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *AssayResult) GetCreatorOk() (*AssayResultCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *AssayResult) SetCreator(v AssayResultCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *AssayResult) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEntryId

`func (o *AssayResult) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *AssayResult) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *AssayResult) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *AssayResult) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### SetEntryIdNil

`func (o *AssayResult) SetEntryIdNil(b bool)`

 SetEntryIdNil sets the value for EntryId to be an explicit nil

### UnsetEntryId
`func (o *AssayResult) UnsetEntryId()`

UnsetEntryId ensures that no value is present for EntryId, not even an explicit nil
### GetFieldValidation

`func (o *AssayResult) GetFieldValidation() map[string]UserValidation`

GetFieldValidation returns the FieldValidation field if non-nil, zero value otherwise.

### GetFieldValidationOk

`func (o *AssayResult) GetFieldValidationOk() (*map[string]UserValidation, bool)`

GetFieldValidationOk returns a tuple with the FieldValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValidation

`func (o *AssayResult) SetFieldValidation(v map[string]UserValidation)`

SetFieldValidation sets FieldValidation field to given value.

### HasFieldValidation

`func (o *AssayResult) HasFieldValidation() bool`

HasFieldValidation returns a boolean if a field has been set.

### GetFields

`func (o *AssayResult) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AssayResult) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AssayResult) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AssayResult) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *AssayResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsReviewed

`func (o *AssayResult) GetIsReviewed() bool`

GetIsReviewed returns the IsReviewed field if non-nil, zero value otherwise.

### GetIsReviewedOk

`func (o *AssayResult) GetIsReviewedOk() (*bool, bool)`

GetIsReviewedOk returns a tuple with the IsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReviewed

`func (o *AssayResult) SetIsReviewed(v bool)`

SetIsReviewed sets IsReviewed field to given value.

### HasIsReviewed

`func (o *AssayResult) HasIsReviewed() bool`

HasIsReviewed returns a boolean if a field has been set.

### GetProjectId

`func (o *AssayResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssayResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssayResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AssayResult) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *AssayResult) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AssayResult) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSchema

`func (o *AssayResult) GetSchema() SchemaProperty`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AssayResult) GetSchemaOk() (*SchemaProperty, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AssayResult) SetSchema(v SchemaProperty)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AssayResult) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetValidationComment

`func (o *AssayResult) GetValidationComment() string`

GetValidationComment returns the ValidationComment field if non-nil, zero value otherwise.

### GetValidationCommentOk

`func (o *AssayResult) GetValidationCommentOk() (*string, bool)`

GetValidationCommentOk returns a tuple with the ValidationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationComment

`func (o *AssayResult) SetValidationComment(v string)`

SetValidationComment sets ValidationComment field to given value.

### HasValidationComment

`func (o *AssayResult) HasValidationComment() bool`

HasValidationComment returns a boolean if a field has been set.

### GetValidationStatus

`func (o *AssayResult) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *AssayResult) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *AssayResult) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *AssayResult) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


