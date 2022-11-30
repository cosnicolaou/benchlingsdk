# AssayRunUpdatedFieldsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**Updates** | Pointer to **[]string** | These properties have been updated, causing this message  | [optional] 
**AssayRun** | Pointer to [**AssayRun**](AssayRun.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayRunUpdatedFieldsEvent

`func NewAssayRunUpdatedFieldsEvent() *AssayRunUpdatedFieldsEvent`

NewAssayRunUpdatedFieldsEvent instantiates a new AssayRunUpdatedFieldsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunUpdatedFieldsEventWithDefaults

`func NewAssayRunUpdatedFieldsEventWithDefaults() *AssayRunUpdatedFieldsEvent`

NewAssayRunUpdatedFieldsEventWithDefaults instantiates a new AssayRunUpdatedFieldsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AssayRunUpdatedFieldsEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssayRunUpdatedFieldsEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssayRunUpdatedFieldsEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssayRunUpdatedFieldsEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *AssayRunUpdatedFieldsEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AssayRunUpdatedFieldsEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AssayRunUpdatedFieldsEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AssayRunUpdatedFieldsEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *AssayRunUpdatedFieldsEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *AssayRunUpdatedFieldsEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *AssayRunUpdatedFieldsEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *AssayRunUpdatedFieldsEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *AssayRunUpdatedFieldsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayRunUpdatedFieldsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayRunUpdatedFieldsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayRunUpdatedFieldsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *AssayRunUpdatedFieldsEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AssayRunUpdatedFieldsEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AssayRunUpdatedFieldsEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AssayRunUpdatedFieldsEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *AssayRunUpdatedFieldsEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *AssayRunUpdatedFieldsEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetUpdates

`func (o *AssayRunUpdatedFieldsEvent) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *AssayRunUpdatedFieldsEvent) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *AssayRunUpdatedFieldsEvent) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *AssayRunUpdatedFieldsEvent) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetAssayRun

`func (o *AssayRunUpdatedFieldsEvent) GetAssayRun() AssayRun`

GetAssayRun returns the AssayRun field if non-nil, zero value otherwise.

### GetAssayRunOk

`func (o *AssayRunUpdatedFieldsEvent) GetAssayRunOk() (*AssayRun, bool)`

GetAssayRunOk returns a tuple with the AssayRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRun

`func (o *AssayRunUpdatedFieldsEvent) SetAssayRun(v AssayRun)`

SetAssayRun sets AssayRun field to given value.

### HasAssayRun

`func (o *AssayRunUpdatedFieldsEvent) HasAssayRun() bool`

HasAssayRun returns a boolean if a field has been set.

### GetEventType

`func (o *AssayRunUpdatedFieldsEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AssayRunUpdatedFieldsEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AssayRunUpdatedFieldsEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AssayRunUpdatedFieldsEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


