# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**Entity** | Pointer to [**GenericEntity**](GenericEntity.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**Entry** | Pointer to [**StageEntry**](StageEntry.md) |  | [optional] 
**Updates** | Pointer to **[]string** | These properties have been updated, causing this message  | [optional] 
**StageEntry** | Pointer to [**StageEntry**](StageEntry.md) |  | [optional] 
**Request** | Pointer to [**Request**](Request.md) |  | [optional] 
**AssayRun** | Pointer to [**AssayRun**](AssayRun.md) |  | [optional] 
**AutomationInputGenerator** | Pointer to [**AutomationFile**](AutomationFile.md) |  | [optional] 
**AutomationOutputProcessor** | Pointer to [**AutomationFile**](AutomationFile.md) |  | [optional] 
**WorkflowTaskGroup** | Pointer to [**WorkflowTaskGroup**](WorkflowTaskGroup.md) |  | [optional] 
**WorkflowTask** | Pointer to [**WorkflowTask**](WorkflowTask.md) |  | [optional] 
**WorkflowOutput** | Pointer to [**WorkflowOutput**](WorkflowOutput.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Event) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *Event) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Event) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Event) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Event) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *Event) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *Event) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *Event) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *Event) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *Event) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Event) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Event) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Event) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Event) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Event) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetEntity

`func (o *Event) GetEntity() GenericEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Event) GetEntityOk() (*GenericEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Event) SetEntity(v GenericEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Event) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Event) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEntry

`func (o *Event) GetEntry() StageEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *Event) GetEntryOk() (*StageEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *Event) SetEntry(v StageEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *Event) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetUpdates

`func (o *Event) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *Event) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *Event) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *Event) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetStageEntry

`func (o *Event) GetStageEntry() StageEntry`

GetStageEntry returns the StageEntry field if non-nil, zero value otherwise.

### GetStageEntryOk

`func (o *Event) GetStageEntryOk() (*StageEntry, bool)`

GetStageEntryOk returns a tuple with the StageEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageEntry

`func (o *Event) SetStageEntry(v StageEntry)`

SetStageEntry sets StageEntry field to given value.

### HasStageEntry

`func (o *Event) HasStageEntry() bool`

HasStageEntry returns a boolean if a field has been set.

### GetRequest

`func (o *Event) GetRequest() Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Event) GetRequestOk() (*Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Event) SetRequest(v Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *Event) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetAssayRun

`func (o *Event) GetAssayRun() AssayRun`

GetAssayRun returns the AssayRun field if non-nil, zero value otherwise.

### GetAssayRunOk

`func (o *Event) GetAssayRunOk() (*AssayRun, bool)`

GetAssayRunOk returns a tuple with the AssayRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRun

`func (o *Event) SetAssayRun(v AssayRun)`

SetAssayRun sets AssayRun field to given value.

### HasAssayRun

`func (o *Event) HasAssayRun() bool`

HasAssayRun returns a boolean if a field has been set.

### GetAutomationInputGenerator

`func (o *Event) GetAutomationInputGenerator() AutomationFile`

GetAutomationInputGenerator returns the AutomationInputGenerator field if non-nil, zero value otherwise.

### GetAutomationInputGeneratorOk

`func (o *Event) GetAutomationInputGeneratorOk() (*AutomationFile, bool)`

GetAutomationInputGeneratorOk returns a tuple with the AutomationInputGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationInputGenerator

`func (o *Event) SetAutomationInputGenerator(v AutomationFile)`

SetAutomationInputGenerator sets AutomationInputGenerator field to given value.

### HasAutomationInputGenerator

`func (o *Event) HasAutomationInputGenerator() bool`

HasAutomationInputGenerator returns a boolean if a field has been set.

### GetAutomationOutputProcessor

`func (o *Event) GetAutomationOutputProcessor() AutomationFile`

GetAutomationOutputProcessor returns the AutomationOutputProcessor field if non-nil, zero value otherwise.

### GetAutomationOutputProcessorOk

`func (o *Event) GetAutomationOutputProcessorOk() (*AutomationFile, bool)`

GetAutomationOutputProcessorOk returns a tuple with the AutomationOutputProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationOutputProcessor

`func (o *Event) SetAutomationOutputProcessor(v AutomationFile)`

SetAutomationOutputProcessor sets AutomationOutputProcessor field to given value.

### HasAutomationOutputProcessor

`func (o *Event) HasAutomationOutputProcessor() bool`

HasAutomationOutputProcessor returns a boolean if a field has been set.

### GetWorkflowTaskGroup

`func (o *Event) GetWorkflowTaskGroup() WorkflowTaskGroup`

GetWorkflowTaskGroup returns the WorkflowTaskGroup field if non-nil, zero value otherwise.

### GetWorkflowTaskGroupOk

`func (o *Event) GetWorkflowTaskGroupOk() (*WorkflowTaskGroup, bool)`

GetWorkflowTaskGroupOk returns a tuple with the WorkflowTaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTaskGroup

`func (o *Event) SetWorkflowTaskGroup(v WorkflowTaskGroup)`

SetWorkflowTaskGroup sets WorkflowTaskGroup field to given value.

### HasWorkflowTaskGroup

`func (o *Event) HasWorkflowTaskGroup() bool`

HasWorkflowTaskGroup returns a boolean if a field has been set.

### GetWorkflowTask

`func (o *Event) GetWorkflowTask() WorkflowTask`

GetWorkflowTask returns the WorkflowTask field if non-nil, zero value otherwise.

### GetWorkflowTaskOk

`func (o *Event) GetWorkflowTaskOk() (*WorkflowTask, bool)`

GetWorkflowTaskOk returns a tuple with the WorkflowTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTask

`func (o *Event) SetWorkflowTask(v WorkflowTask)`

SetWorkflowTask sets WorkflowTask field to given value.

### HasWorkflowTask

`func (o *Event) HasWorkflowTask() bool`

HasWorkflowTask returns a boolean if a field has been set.

### GetWorkflowOutput

`func (o *Event) GetWorkflowOutput() WorkflowOutput`

GetWorkflowOutput returns the WorkflowOutput field if non-nil, zero value otherwise.

### GetWorkflowOutputOk

`func (o *Event) GetWorkflowOutputOk() (*WorkflowOutput, bool)`

GetWorkflowOutputOk returns a tuple with the WorkflowOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowOutput

`func (o *Event) SetWorkflowOutput(v WorkflowOutput)`

SetWorkflowOutput sets WorkflowOutput field to given value.

### HasWorkflowOutput

`func (o *Event) HasWorkflowOutput() bool`

HasWorkflowOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


