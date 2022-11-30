# AutomationInputGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRunId** | Pointer to **string** |  | [optional] 
**AutomationFileConfig** | Pointer to [**AutomationFileAutomationFileConfig**](AutomationFileAutomationFileConfig.md) |  | [optional] 
**File** | Pointer to [**NullableAutomationFileFile**](AutomationFileFile.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ApiURL** | Pointer to **string** | The canonical url of the Automation Input Generator in the API. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Automation Input Generator was last modified | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Automation Input Generator was last modified | [optional] [readonly] 

## Methods

### NewAutomationInputGenerator

`func NewAutomationInputGenerator() *AutomationInputGenerator`

NewAutomationInputGenerator instantiates a new AutomationInputGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationInputGeneratorWithDefaults

`func NewAutomationInputGeneratorWithDefaults() *AutomationInputGenerator`

NewAutomationInputGeneratorWithDefaults instantiates a new AutomationInputGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRunId

`func (o *AutomationInputGenerator) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *AutomationInputGenerator) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *AutomationInputGenerator) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.

### HasAssayRunId

`func (o *AutomationInputGenerator) HasAssayRunId() bool`

HasAssayRunId returns a boolean if a field has been set.

### GetAutomationFileConfig

`func (o *AutomationInputGenerator) GetAutomationFileConfig() AutomationFileAutomationFileConfig`

GetAutomationFileConfig returns the AutomationFileConfig field if non-nil, zero value otherwise.

### GetAutomationFileConfigOk

`func (o *AutomationInputGenerator) GetAutomationFileConfigOk() (*AutomationFileAutomationFileConfig, bool)`

GetAutomationFileConfigOk returns a tuple with the AutomationFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationFileConfig

`func (o *AutomationInputGenerator) SetAutomationFileConfig(v AutomationFileAutomationFileConfig)`

SetAutomationFileConfig sets AutomationFileConfig field to given value.

### HasAutomationFileConfig

`func (o *AutomationInputGenerator) HasAutomationFileConfig() bool`

HasAutomationFileConfig returns a boolean if a field has been set.

### GetFile

`func (o *AutomationInputGenerator) GetFile() AutomationFileFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AutomationInputGenerator) GetFileOk() (*AutomationFileFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AutomationInputGenerator) SetFile(v AutomationFileFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AutomationInputGenerator) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *AutomationInputGenerator) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *AutomationInputGenerator) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetId

`func (o *AutomationInputGenerator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationInputGenerator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationInputGenerator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationInputGenerator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AutomationInputGenerator) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationInputGenerator) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationInputGenerator) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomationInputGenerator) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiURL

`func (o *AutomationInputGenerator) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AutomationInputGenerator) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AutomationInputGenerator) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *AutomationInputGenerator) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationInputGenerator) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationInputGenerator) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationInputGenerator) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationInputGenerator) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AutomationInputGenerator) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutomationInputGenerator) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutomationInputGenerator) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AutomationInputGenerator) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


