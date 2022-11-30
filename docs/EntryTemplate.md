# EntryTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Entry Template in the API. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | DateTime the template was created at | [optional] [readonly] 
**Creator** | Pointer to [**EntryTemplateCreator**](EntryTemplateCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**Days** | Pointer to [**[]EntryTemplateDay**](EntryTemplateDay.md) | Array of day objects. Each day object has a day index (integer) and notes field (array of notes, expand further for details on note types).  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the entry template | [optional] 
**ModifiedAt** | Pointer to **string** | DateTime the template was last modified | [optional] 
**Name** | Pointer to **string** | Title of the template | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty4**](SchemaProperty4.md) |  | [optional] 
**TemplateCollectionId** | Pointer to **string** | ID of the collection that contains the template | [optional] 
**WebURL** | Pointer to **string** | URL of the template | [optional] 

## Methods

### NewEntryTemplate

`func NewEntryTemplate() *EntryTemplate`

NewEntryTemplate instantiates a new EntryTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTemplateWithDefaults

`func NewEntryTemplateWithDefaults() *EntryTemplate`

NewEntryTemplateWithDefaults instantiates a new EntryTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *EntryTemplate) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *EntryTemplate) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *EntryTemplate) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *EntryTemplate) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntryTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntryTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntryTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntryTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *EntryTemplate) GetCreator() EntryTemplateCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *EntryTemplate) GetCreatorOk() (*EntryTemplateCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *EntryTemplate) SetCreator(v EntryTemplateCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *EntryTemplate) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *EntryTemplate) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EntryTemplate) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EntryTemplate) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EntryTemplate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDays

`func (o *EntryTemplate) GetDays() []EntryTemplateDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *EntryTemplate) GetDaysOk() (*[]EntryTemplateDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *EntryTemplate) SetDays(v []EntryTemplateDay)`

SetDays sets Days field to given value.

### HasDays

`func (o *EntryTemplate) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetFields

`func (o *EntryTemplate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EntryTemplate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EntryTemplate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EntryTemplate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *EntryTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntryTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntryTemplate) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntryTemplate) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntryTemplate) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntryTemplate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *EntryTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntryTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchema

`func (o *EntryTemplate) GetSchema() SchemaProperty4`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *EntryTemplate) GetSchemaOk() (*SchemaProperty4, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *EntryTemplate) SetSchema(v SchemaProperty4)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *EntryTemplate) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *EntryTemplate) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *EntryTemplate) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetTemplateCollectionId

`func (o *EntryTemplate) GetTemplateCollectionId() string`

GetTemplateCollectionId returns the TemplateCollectionId field if non-nil, zero value otherwise.

### GetTemplateCollectionIdOk

`func (o *EntryTemplate) GetTemplateCollectionIdOk() (*string, bool)`

GetTemplateCollectionIdOk returns a tuple with the TemplateCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCollectionId

`func (o *EntryTemplate) SetTemplateCollectionId(v string)`

SetTemplateCollectionId sets TemplateCollectionId field to given value.

### HasTemplateCollectionId

`func (o *EntryTemplate) HasTemplateCollectionId() bool`

HasTemplateCollectionId returns a boolean if a field has been set.

### GetWebURL

`func (o *EntryTemplate) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *EntryTemplate) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *EntryTemplate) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *EntryTemplate) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


