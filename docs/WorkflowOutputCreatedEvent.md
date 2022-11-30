# WorkflowOutputCreatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**WorkflowOutput** | Pointer to [**WorkflowOutput**](WorkflowOutput.md) |  | [optional] 

## Methods

### NewWorkflowOutputCreatedEvent

`func NewWorkflowOutputCreatedEvent() *WorkflowOutputCreatedEvent`

NewWorkflowOutputCreatedEvent instantiates a new WorkflowOutputCreatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputCreatedEventWithDefaults

`func NewWorkflowOutputCreatedEventWithDefaults() *WorkflowOutputCreatedEvent`

NewWorkflowOutputCreatedEventWithDefaults instantiates a new WorkflowOutputCreatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkflowOutputCreatedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowOutputCreatedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowOutputCreatedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowOutputCreatedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *WorkflowOutputCreatedEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkflowOutputCreatedEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkflowOutputCreatedEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *WorkflowOutputCreatedEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *WorkflowOutputCreatedEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *WorkflowOutputCreatedEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *WorkflowOutputCreatedEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *WorkflowOutputCreatedEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *WorkflowOutputCreatedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowOutputCreatedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowOutputCreatedEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowOutputCreatedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *WorkflowOutputCreatedEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkflowOutputCreatedEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkflowOutputCreatedEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkflowOutputCreatedEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *WorkflowOutputCreatedEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *WorkflowOutputCreatedEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetEventType

`func (o *WorkflowOutputCreatedEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WorkflowOutputCreatedEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WorkflowOutputCreatedEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *WorkflowOutputCreatedEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetWorkflowOutput

`func (o *WorkflowOutputCreatedEvent) GetWorkflowOutput() WorkflowOutput`

GetWorkflowOutput returns the WorkflowOutput field if non-nil, zero value otherwise.

### GetWorkflowOutputOk

`func (o *WorkflowOutputCreatedEvent) GetWorkflowOutputOk() (*WorkflowOutput, bool)`

GetWorkflowOutputOk returns a tuple with the WorkflowOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutput

`func (o *WorkflowOutputCreatedEvent) SetWorkflowOutput(v WorkflowOutput)`

SetWorkflowOutput sets WorkflowOutput field to given value.

### HasWorkflowOutput

`func (o *WorkflowOutputCreatedEvent) HasWorkflowOutput() bool`

HasWorkflowOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


