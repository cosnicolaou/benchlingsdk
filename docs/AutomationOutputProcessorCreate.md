# AutomationOutputProcessorCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRunId** | **string** |  | 
**AutomationFileConfigName** | **string** |  | 
**CompleteWithErrors** | Pointer to **bool** | Specifies whether file processing should complete with errors. False means any error in output file processing will result in no actions being committed. True means that if row-level errors occur, then failing rows and their errors will be saved to errorFile, and actions from successful rows will be committed. | [optional] 
**FileId** | **string** | The ID of a blob link to process. | 

## Methods

### NewAutomationOutputProcessorCreate

`func NewAutomationOutputProcessorCreate(assayRunId string, automationFileConfigName string, fileId string, ) *AutomationOutputProcessorCreate`

NewAutomationOutputProcessorCreate instantiates a new AutomationOutputProcessorCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationOutputProcessorCreateWithDefaults

`func NewAutomationOutputProcessorCreateWithDefaults() *AutomationOutputProcessorCreate`

NewAutomationOutputProcessorCreateWithDefaults instantiates a new AutomationOutputProcessorCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRunId

`func (o *AutomationOutputProcessorCreate) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *AutomationOutputProcessorCreate) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *AutomationOutputProcessorCreate) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.


### GetAutomationFileConfigName

`func (o *AutomationOutputProcessorCreate) GetAutomationFileConfigName() string`

GetAutomationFileConfigName returns the AutomationFileConfigName field if non-nil, zero value otherwise.

### GetAutomationFileConfigNameOk

`func (o *AutomationOutputProcessorCreate) GetAutomationFileConfigNameOk() (*string, bool)`

GetAutomationFileConfigNameOk returns a tuple with the AutomationFileConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationFileConfigName

`func (o *AutomationOutputProcessorCreate) SetAutomationFileConfigName(v string)`

SetAutomationFileConfigName sets AutomationFileConfigName field to given value.


### GetCompleteWithErrors

`func (o *AutomationOutputProcessorCreate) GetCompleteWithErrors() bool`

GetCompleteWithErrors returns the CompleteWithErrors field if non-nil, zero value otherwise.

### GetCompleteWithErrorsOk

`func (o *AutomationOutputProcessorCreate) GetCompleteWithErrorsOk() (*bool, bool)`

GetCompleteWithErrorsOk returns a tuple with the CompleteWithErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteWithErrors

`func (o *AutomationOutputProcessorCreate) SetCompleteWithErrors(v bool)`

SetCompleteWithErrors sets CompleteWithErrors field to given value.

### HasCompleteWithErrors

`func (o *AutomationOutputProcessorCreate) HasCompleteWithErrors() bool`

HasCompleteWithErrors returns a boolean if a field has been set.

### GetFileId

`func (o *AutomationOutputProcessorCreate) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AutomationOutputProcessorCreate) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AutomationOutputProcessorCreate) SetFileId(v string)`

SetFileId sets FileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


