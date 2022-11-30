# WorkflowTaskGroupUpdatedWatchersEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**WorkflowTaskGroup** | Pointer to [**WorkflowTaskGroup**](WorkflowTaskGroup.md) |  | [optional] 

## Methods

### NewWorkflowTaskGroupUpdatedWatchersEvent

`func NewWorkflowTaskGroupUpdatedWatchersEvent() *WorkflowTaskGroupUpdatedWatchersEvent`

NewWorkflowTaskGroupUpdatedWatchersEvent instantiates a new WorkflowTaskGroupUpdatedWatchersEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskGroupUpdatedWatchersEventWithDefaults

`func NewWorkflowTaskGroupUpdatedWatchersEventWithDefaults() *WorkflowTaskGroupUpdatedWatchersEvent`

NewWorkflowTaskGroupUpdatedWatchersEventWithDefaults instantiates a new WorkflowTaskGroupUpdatedWatchersEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *WorkflowTaskGroupUpdatedWatchersEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetEventType

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetWorkflowTaskGroup() WorkflowTaskGroup`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) GetWorkflowTaskGroupOk() (*WorkflowTaskGroup, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) SetWorkflowTaskGroup(v WorkflowTaskGroup)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *WorkflowTaskGroupUpdatedWatchersEvent) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


