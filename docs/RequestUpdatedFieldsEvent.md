# RequestUpdatedFieldsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**Updates** | Pointer to **[]string** | These properties have been updated, causing this message  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**Request** | Pointer to [**Request**](Request.md) |  | [optional] 

## Methods

### NewRequestUpdatedFieldsEvent

`func NewRequestUpdatedFieldsEvent() *RequestUpdatedFieldsEvent`

NewRequestUpdatedFieldsEvent instantiates a new RequestUpdatedFieldsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpdatedFieldsEventWithDefaults

`func NewRequestUpdatedFieldsEventWithDefaults() *RequestUpdatedFieldsEvent`

NewRequestUpdatedFieldsEventWithDefaults instantiates a new RequestUpdatedFieldsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RequestUpdatedFieldsEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestUpdatedFieldsEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestUpdatedFieldsEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestUpdatedFieldsEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *RequestUpdatedFieldsEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *RequestUpdatedFieldsEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *RequestUpdatedFieldsEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *RequestUpdatedFieldsEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *RequestUpdatedFieldsEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *RequestUpdatedFieldsEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *RequestUpdatedFieldsEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *RequestUpdatedFieldsEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *RequestUpdatedFieldsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestUpdatedFieldsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestUpdatedFieldsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestUpdatedFieldsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *RequestUpdatedFieldsEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RequestUpdatedFieldsEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RequestUpdatedFieldsEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RequestUpdatedFieldsEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *RequestUpdatedFieldsEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *RequestUpdatedFieldsEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetUpdates

`func (o *RequestUpdatedFieldsEvent) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *RequestUpdatedFieldsEvent) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *RequestUpdatedFieldsEvent) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *RequestUpdatedFieldsEvent) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetEventType

`func (o *RequestUpdatedFieldsEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *RequestUpdatedFieldsEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *RequestUpdatedFieldsEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *RequestUpdatedFieldsEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetRequest

`func (o *RequestUpdatedFieldsEvent) GetRequest() Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *RequestUpdatedFieldsEvent) GetRequestOk() (*Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *RequestUpdatedFieldsEvent) SetRequest(v Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *RequestUpdatedFieldsEvent) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


