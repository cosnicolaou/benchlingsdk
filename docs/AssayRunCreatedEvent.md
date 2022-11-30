# AssayRunCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**AssayRun** | Pointer to [**AssayRun**](AssayRun.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayRunCreatedEvent

`func NewAssayRunCreatedEvent() *AssayRunCreatedEvent`

NewAssayRunCreatedEvent instantiates a new AssayRunCreatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunCreatedEventWithDefaults

`func NewAssayRunCreatedEventWithDefaults() *AssayRunCreatedEvent`

NewAssayRunCreatedEventWithDefaults instantiates a new AssayRunCreatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AssayRunCreatedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssayRunCreatedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssayRunCreatedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssayRunCreatedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *AssayRunCreatedEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AssayRunCreatedEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AssayRunCreatedEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AssayRunCreatedEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *AssayRunCreatedEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *AssayRunCreatedEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *AssayRunCreatedEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *AssayRunCreatedEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *AssayRunCreatedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayRunCreatedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayRunCreatedEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayRunCreatedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *AssayRunCreatedEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AssayRunCreatedEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AssayRunCreatedEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AssayRunCreatedEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *AssayRunCreatedEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *AssayRunCreatedEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetAssayRun

`func (o *AssayRunCreatedEvent) GetAssayRun() AssayRun`

GetAssayRun returns the AssayRun field if non-nil, zero value otherwise.

### GetAssayRunOk

`func (o *AssayRunCreatedEvent) GetAssayRunOk() (*AssayRun, bool)`

GetAssayRunOk returns a tuple with the AssayRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRun

`func (o *AssayRunCreatedEvent) SetAssayRun(v AssayRun)`

SetAssayRun sets AssayRun field to given value.

### HasAssayRun

`func (o *AssayRunCreatedEvent) HasAssayRun() bool`

HasAssayRun returns a boolean if a field has been set.

### GetEventType

`func (o *AssayRunCreatedEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AssayRunCreatedEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AssayRunCreatedEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AssayRunCreatedEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


