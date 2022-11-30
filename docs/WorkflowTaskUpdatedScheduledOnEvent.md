# WorkflowTaskUpdatedScheduledOnEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**WorkflowTask** | Pointer to [**WorkflowTask**](WorkflowTask.md) |  | [optional] 

## Methods

### NewWorkflowTaskUpdatedScheduledOnEvent

`func NewWorkflowTaskUpdatedScheduledOnEvent() *WorkflowTaskUpdatedScheduledOnEvent`

NewWorkflowTaskUpdatedScheduledOnEvent instantiates a new WorkflowTaskUpdatedScheduledOnEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskUpdatedScheduledOnEventWithDefaults

`func NewWorkflowTaskUpdatedScheduledOnEventWithDefaults() *WorkflowTaskUpdatedScheduledOnEvent`

NewWorkflowTaskUpdatedScheduledOnEventWithDefaults instantiates a new WorkflowTaskUpdatedScheduledOnEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *WorkflowTaskUpdatedScheduledOnEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetEventType

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetWorkflowTask

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetWorkflowTask() WorkflowTask`

GetWorkflowTask returns the WorkflowTask field if non-nil, zero value otherwise.

### GetWorkflowTaskOk

`func (o *WorkflowTaskUpdatedScheduledOnEvent) GetWorkflowTaskOk() (*WorkflowTask, bool)`

GetWorkflowTaskOk returns a tuple with the WorkflowTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTask

`func (o *WorkflowTaskUpdatedScheduledOnEvent) SetWorkflowTask(v WorkflowTask)`

SetWorkflowTask sets WorkflowTask field to given value.

### HasWorkflowTask

`func (o *WorkflowTaskUpdatedScheduledOnEvent) HasWorkflowTask() bool`

HasWorkflowTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


