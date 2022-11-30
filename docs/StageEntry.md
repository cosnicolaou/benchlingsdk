# StageEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Stage Entry in the API. | [optional] [readonly] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) | Array of UserSummary Resources of the authors of the stage entry. This defaults to the creator but can be manually changed.  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the stage entry was created at | [optional] [readonly] 
**Creator** | Pointer to [**StageEntryCreator**](StageEntryCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**DisplayId** | Pointer to **string** | User-friendly ID of the stage entry | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder that contains the stage entry | [optional] 
**Id** | Pointer to **string** | ID of the stage entry | [optional] 
**ModifiedAt** | Pointer to **string** | DateTime the stage entry was last modified | [optional] 
**Name** | Pointer to **string** | Title of the stage entry | [optional] 
**ReviewRecord** | Pointer to **map[string]interface{}** | Review record if set | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty4**](SchemaProperty4.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the stage entry | [optional] 
**WorkflowId** | Pointer to **string** | ID of the parent workflow | [optional] 
**WorkflowStageId** | Pointer to **string** | ID of the associated workflow stage | [optional] 

## Methods

### NewStageEntry

`func NewStageEntry() *StageEntry`

NewStageEntry instantiates a new StageEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageEntryWithDefaults

`func NewStageEntryWithDefaults() *StageEntry`

NewStageEntryWithDefaults instantiates a new StageEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *StageEntry) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *StageEntry) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *StageEntry) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *StageEntry) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetAuthors

`func (o *StageEntry) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *StageEntry) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *StageEntry) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *StageEntry) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StageEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StageEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StageEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StageEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *StageEntry) GetCreator() StageEntryCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StageEntry) GetCreatorOk() (*StageEntryCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StageEntry) SetCreator(v StageEntryCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StageEntry) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *StageEntry) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StageEntry) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StageEntry) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StageEntry) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDisplayId

`func (o *StageEntry) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *StageEntry) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *StageEntry) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *StageEntry) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetFields

`func (o *StageEntry) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *StageEntry) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *StageEntry) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *StageEntry) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *StageEntry) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *StageEntry) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *StageEntry) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *StageEntry) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetId

`func (o *StageEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StageEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StageEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StageEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *StageEntry) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *StageEntry) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *StageEntry) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *StageEntry) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *StageEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StageEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StageEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StageEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReviewRecord

`func (o *StageEntry) GetReviewRecord() map[string]interface{}`

GetReviewRecord returns the ReviewRecord field if non-nil, zero value otherwise.

### GetReviewRecordOk

`func (o *StageEntry) GetReviewRecordOk() (*map[string]interface{}, bool)`

GetReviewRecordOk returns a tuple with the ReviewRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRecord

`func (o *StageEntry) SetReviewRecord(v map[string]interface{})`

SetReviewRecord sets ReviewRecord field to given value.

### HasReviewRecord

`func (o *StageEntry) HasReviewRecord() bool`

HasReviewRecord returns a boolean if a field has been set.

### SetReviewRecordNil

`func (o *StageEntry) SetReviewRecordNil(b bool)`

 SetReviewRecordNil sets the value for ReviewRecord to be an explicit nil

### UnsetReviewRecord
`func (o *StageEntry) UnsetReviewRecord()`

UnsetReviewRecord ensures that no value is present for ReviewRecord, not even an explicit nil
### GetSchema

`func (o *StageEntry) GetSchema() SchemaProperty4`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *StageEntry) GetSchemaOk() (*SchemaProperty4, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *StageEntry) SetSchema(v SchemaProperty4)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *StageEntry) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *StageEntry) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *StageEntry) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *StageEntry) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *StageEntry) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *StageEntry) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *StageEntry) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWorkflowId

`func (o *StageEntry) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *StageEntry) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *StageEntry) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *StageEntry) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowStageId

`func (o *StageEntry) GetWorkflowStageId() string`

GetWorkflowStageId returns the WorkflowStageId field if non-nil, zero value otherwise.

### GetWorkflowStageIdOk

`func (o *StageEntry) GetWorkflowStageIdOk() (*string, bool)`

GetWorkflowStageIdOk returns a tuple with the WorkflowStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStageId

`func (o *StageEntry) SetWorkflowStageId(v string)`

SetWorkflowStageId sets WorkflowStageId field to given value.

### HasWorkflowStageId

`func (o *StageEntry) HasWorkflowStageId() bool`

HasWorkflowStageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


