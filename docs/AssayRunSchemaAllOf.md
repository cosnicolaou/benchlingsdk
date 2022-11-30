# AssayRunSchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomationInputFileConfigs** | Pointer to [**[]AssayRunSchemaAllOfAutomationInputFileConfigs**](AssayRunSchemaAllOfAutomationInputFileConfigs.md) |  | [optional] 
**AutomationOutputFileConfigs** | Pointer to [**[]AssayRunSchemaAllOfAutomationInputFileConfigs**](AssayRunSchemaAllOfAutomationInputFileConfigs.md) |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Assay Run Schema was last modified | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAssayRunSchemaAllOf

`func NewAssayRunSchemaAllOf() *AssayRunSchemaAllOf`

NewAssayRunSchemaAllOf instantiates a new AssayRunSchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunSchemaAllOfWithDefaults

`func NewAssayRunSchemaAllOfWithDefaults() *AssayRunSchemaAllOf`

NewAssayRunSchemaAllOfWithDefaults instantiates a new AssayRunSchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomationInputFileConfigs

`func (o *AssayRunSchemaAllOf) GetAutomationInputFileConfigs() []AssayRunSchemaAllOfAutomationInputFileConfigs`

GetAutomationInputFileConfigs returns the AutomationInputFileConfigs field if non-nil, zero value otherwise.

### GetAutomationInputFileConfigsOk

`func (o *AssayRunSchemaAllOf) GetAutomationInputFileConfigsOk() (*[]AssayRunSchemaAllOfAutomationInputFileConfigs, bool)`

GetAutomationInputFileConfigsOk returns a tuple with the AutomationInputFileConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationInputFileConfigs

`func (o *AssayRunSchemaAllOf) SetAutomationInputFileConfigs(v []AssayRunSchemaAllOfAutomationInputFileConfigs)`

SetAutomationInputFileConfigs sets AutomationInputFileConfigs field to given value.

### HasAutomationInputFileConfigs

`func (o *AssayRunSchemaAllOf) HasAutomationInputFileConfigs() bool`

HasAutomationInputFileConfigs returns a boolean if a field has been set.

### GetAutomationOutputFileConfigs

`func (o *AssayRunSchemaAllOf) GetAutomationOutputFileConfigs() []AssayRunSchemaAllOfAutomationInputFileConfigs`

GetAutomationOutputFileConfigs returns the AutomationOutputFileConfigs field if non-nil, zero value otherwise.

### GetAutomationOutputFileConfigsOk

`func (o *AssayRunSchemaAllOf) GetAutomationOutputFileConfigsOk() (*[]AssayRunSchemaAllOfAutomationInputFileConfigs, bool)`

GetAutomationOutputFileConfigsOk returns a tuple with the AutomationOutputFileConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationOutputFileConfigs

`func (o *AssayRunSchemaAllOf) SetAutomationOutputFileConfigs(v []AssayRunSchemaAllOfAutomationInputFileConfigs)`

SetAutomationOutputFileConfigs sets AutomationOutputFileConfigs field to given value.

### HasAutomationOutputFileConfigs

`func (o *AssayRunSchemaAllOf) HasAutomationOutputFileConfigs() bool`

HasAutomationOutputFileConfigs returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AssayRunSchemaAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AssayRunSchemaAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AssayRunSchemaAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AssayRunSchemaAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetType

`func (o *AssayRunSchemaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssayRunSchemaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssayRunSchemaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssayRunSchemaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


