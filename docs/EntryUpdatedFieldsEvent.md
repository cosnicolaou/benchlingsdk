# EntryUpdatedFieldsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**Updates** | Pointer to **[]string** | These properties have been updated, causing this message  | [optional] 
**Entry** | Pointer to [**Entry**](Entry.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 

## Methods

### NewEntryUpdatedFieldsEvent

`func NewEntryUpdatedFieldsEvent() *EntryUpdatedFieldsEvent`

NewEntryUpdatedFieldsEvent instantiates a new EntryUpdatedFieldsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryUpdatedFieldsEventWithDefaults

`func NewEntryUpdatedFieldsEventWithDefaults() *EntryUpdatedFieldsEvent`

NewEntryUpdatedFieldsEventWithDefaults instantiates a new EntryUpdatedFieldsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EntryUpdatedFieldsEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntryUpdatedFieldsEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntryUpdatedFieldsEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntryUpdatedFieldsEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *EntryUpdatedFieldsEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *EntryUpdatedFieldsEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *EntryUpdatedFieldsEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *EntryUpdatedFieldsEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *EntryUpdatedFieldsEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *EntryUpdatedFieldsEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *EntryUpdatedFieldsEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *EntryUpdatedFieldsEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *EntryUpdatedFieldsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryUpdatedFieldsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryUpdatedFieldsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntryUpdatedFieldsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *EntryUpdatedFieldsEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *EntryUpdatedFieldsEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *EntryUpdatedFieldsEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *EntryUpdatedFieldsEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *EntryUpdatedFieldsEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *EntryUpdatedFieldsEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetUpdates

`func (o *EntryUpdatedFieldsEvent) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *EntryUpdatedFieldsEvent) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *EntryUpdatedFieldsEvent) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *EntryUpdatedFieldsEvent) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetEntry

`func (o *EntryUpdatedFieldsEvent) GetEntry() Entry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *EntryUpdatedFieldsEvent) GetEntryOk() (*Entry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *EntryUpdatedFieldsEvent) SetEntry(v Entry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *EntryUpdatedFieldsEvent) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetEventType

`func (o *EntryUpdatedFieldsEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EntryUpdatedFieldsEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EntryUpdatedFieldsEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EntryUpdatedFieldsEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


