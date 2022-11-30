# AutomationInputGeneratorCompletedV2BetaEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ExcludedProperties** | Pointer to **[]string** | These properties have been dropped from the payload due to size.  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableEventBaseSchema**](EventBaseSchema.md) |  | [optional] 
**AutomationInputGenerator** | Pointer to [**AutomationFile**](AutomationFile.md) |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 

## Methods

### NewAutomationInputGeneratorCompletedV2BetaEvent

`func NewAutomationInputGeneratorCompletedV2BetaEvent() *AutomationInputGeneratorCompletedV2BetaEvent`

NewAutomationInputGeneratorCompletedV2BetaEvent instantiates a new AutomationInputGeneratorCompletedV2BetaEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationInputGeneratorCompletedV2BetaEventWithDefaults

`func NewAutomationInputGeneratorCompletedV2BetaEventWithDefaults() *AutomationInputGeneratorCompletedV2BetaEvent`

NewAutomationInputGeneratorCompletedV2BetaEventWithDefaults instantiates a new AutomationInputGeneratorCompletedV2BetaEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeprecated

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetExcludedProperties

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetExcludedProperties() []string`

GetExcludedProperties returns the ExcludedProperties field if non-nil, zero value otherwise.

### GetExcludedPropertiesOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetExcludedPropertiesOk() (*[]string, bool)`

GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedProperties

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetExcludedProperties(v []string)`

SetExcludedProperties sets ExcludedProperties field to given value.

### HasExcludedProperties

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasExcludedProperties() bool`

HasExcludedProperties returns a boolean if a field has been set.

### GetId

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchema

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetSchema() EventBaseSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetSchemaOk() (*EventBaseSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetSchema(v EventBaseSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *AutomationInputGeneratorCompletedV2BetaEvent) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetAutomationInputGenerator

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetAutomationInputGenerator() AutomationFile`

GetAutomationInputGenerator returns the AutomationInputGenerator field if non-nil, zero value otherwise.

### GetAutomationInputGeneratorOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetAutomationInputGeneratorOk() (*AutomationFile, bool)`

GetAutomationInputGeneratorOk returns a tuple with the AutomationInputGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationInputGenerator

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetAutomationInputGenerator(v AutomationFile)`

SetAutomationInputGenerator sets AutomationInputGenerator field to given value.

### HasAutomationInputGenerator

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasAutomationInputGenerator() bool`

HasAutomationInputGenerator returns a boolean if a field has been set.

### GetEventType

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


