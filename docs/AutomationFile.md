# AutomationFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRunId** | Pointer to **string** |  | [optional] 
**AutomationFileConfig** | Pointer to [**AutomationFileAutomationFileConfig**](AutomationFileAutomationFileConfig.md) |  | [optional] 
**File** | Pointer to [**NullableAutomationFileFile**](AutomationFileFile.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAutomationFile

`func NewAutomationFile() *AutomationFile`

NewAutomationFile instantiates a new AutomationFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationFileWithDefaults

`func NewAutomationFileWithDefaults() *AutomationFile`

NewAutomationFileWithDefaults instantiates a new AutomationFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRunId

`func (o *AutomationFile) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *AutomationFile) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *AutomationFile) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.

### HasAssayRunId

`func (o *AutomationFile) HasAssayRunId() bool`

HasAssayRunId returns a boolean if a field has been set.

### GetAutomationFileConfig

`func (o *AutomationFile) GetAutomationFileConfig() AutomationFileAutomationFileConfig`

GetAutomationFileConfig returns the AutomationFileConfig field if non-nil, zero value otherwise.

### GetAutomationFileConfigOk

`func (o *AutomationFile) GetAutomationFileConfigOk() (*AutomationFileAutomationFileConfig, bool)`

GetAutomationFileConfigOk returns a tuple with the AutomationFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationFileConfig

`func (o *AutomationFile) SetAutomationFileConfig(v AutomationFileAutomationFileConfig)`

SetAutomationFileConfig sets AutomationFileConfig field to given value.

### HasAutomationFileConfig

`func (o *AutomationFile) HasAutomationFileConfig() bool`

HasAutomationFileConfig returns a boolean if a field has been set.

### GetFile

`func (o *AutomationFile) GetFile() AutomationFileFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AutomationFile) GetFileOk() (*AutomationFileFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AutomationFile) SetFile(v AutomationFileFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AutomationFile) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *AutomationFile) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *AutomationFile) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetId

`func (o *AutomationFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AutomationFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomationFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


