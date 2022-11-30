# WorkflowOutputUpdatedFieldsEvent

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

### NewWorkflowOutputUpdatedFieldsEvent

`func NewWorkflowOutputUpdatedFieldsEvent() *WorkflowOutputUpdatedFieldsEvent`

NewWorkflowOutputUpdatedFieldsEvent instantiates a new WorkflowOutputUpdatedFieldsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOutputUpdatedFieldsEventWithDefaults

`func NewWorkflowOutputUpdatedFieldsEventWithDefaults() *WorkflowOutputUpdatedFieldsEvent`

NewWorkflowOutputUpdatedFieldsEventWithDefaults instantiates a new WorkflowOutputUpdatedFieldsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkflowOutputUpdatedFieldsEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowOutputUpdatedFieldsEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowOutputUpdatedFieldsEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *WorkflowOutputUpdatedFieldsEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkflowOutputUpdatedFieldsEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *WorkflowOutputUpdatedFieldsEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *WorkflowOutputUpdatedFieldsEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *WorkflowOutputUpdatedFieldsEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *WorkflowOutputUpdatedFieldsEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *WorkflowOutputUpdatedFieldsEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowOutputUpdatedFieldsEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowOutputUpdatedFieldsEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *WorkflowOutputUpdatedFieldsEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkflowOutputUpdatedFieldsEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkflowOutputUpdatedFieldsEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *WorkflowOutputUpdatedFieldsEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *WorkflowOutputUpdatedFieldsEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetEventType

`func (o *WorkflowOutputUpdatedFieldsEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WorkflowOutputUpdatedFieldsEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *WorkflowOutputUpdatedFieldsEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetWorkflowOutput

`func (o *WorkflowOutputUpdatedFieldsEvent) GetWorkflowOutput() WorkflowOutput`

GetWorkflowOutput returns the WorkflowOutput field if non-nil, zero value otherwise.

### GetWorkflowOutputOk

`func (o *WorkflowOutputUpdatedFieldsEvent) GetWorkflowOutputOk() (*WorkflowOutput, bool)`

GetWorkflowOutputOk returns a tuple with the WorkflowOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutput

`func (o *WorkflowOutputUpdatedFieldsEvent) SetWorkflowOutput(v WorkflowOutput)`

SetWorkflowOutput sets WorkflowOutput field to given value.

### HasWorkflowOutput

`func (o *WorkflowOutputUpdatedFieldsEvent) HasWorkflowOutput() bool`

HasWorkflowOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


