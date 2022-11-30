# AutomationOutputProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRunId** | Pointer to **string** |  | [optional] 
**AutomationFileConfig** | Pointer to [**AutomationFileAutomationFileConfig**](AutomationFileAutomationFileConfig.md) |  | [optional] 
**File** | Pointer to [**NullableAutomationFileFile**](AutomationFileFile.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ApiURL** | Pointer to **string** | The canonical url of the Automation Output Processor in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**ArchiveRecord**](ArchiveRecord.md) |  | [optional] 
**CompleteWithErrors** | Pointer to **bool** | Specifies whether file processing should complete with errors. False means any error in output file processing will result in no actions being committed. True means that if row-level errors occur, then failing rows and their errors will be saved to errorFile, and actions from successful rows will be committed. | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Automation Output Processor was created | [optional] 
**ErrorFile** | Pointer to [**NullableBlob**](Blob.md) |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Automation Output Processor was last modified | [optional] 
**ProgressStats** | Pointer to [**AutomationProgressStats**](AutomationProgressStats.md) |  | [optional] 
**Transforms** | Pointer to [**[]LabAutomationTransform**](LabAutomationTransform.md) |  | [optional] 

## Methods

### NewAutomationOutputProcessor

`func NewAutomationOutputProcessor() *AutomationOutputProcessor`

NewAutomationOutputProcessor instantiates a new AutomationOutputProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationOutputProcessorWithDefaults

`func NewAutomationOutputProcessorWithDefaults() *AutomationOutputProcessor`

NewAutomationOutputProcessorWithDefaults instantiates a new AutomationOutputProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRunId

`func (o *AutomationOutputProcessor) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *AutomationOutputProcessor) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *AutomationOutputProcessor) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.

### HasAssayRunId

`func (o *AutomationOutputProcessor) HasAssayRunId() bool`

HasAssayRunId returns a boolean if a field has been set.

### GetAutomationFileConfig

`func (o *AutomationOutputProcessor) GetAutomationFileConfig() AutomationFileAutomationFileConfig`

GetAutomationFileConfig returns the AutomationFileConfig field if non-nil, zero value otherwise.

### GetAutomationFileConfigOk

`func (o *AutomationOutputProcessor) GetAutomationFileConfigOk() (*AutomationFileAutomationFileConfig, bool)`

GetAutomationFileConfigOk returns a tuple with the AutomationFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationFileConfig

`func (o *AutomationOutputProcessor) SetAutomationFileConfig(v AutomationFileAutomationFileConfig)`

SetAutomationFileConfig sets AutomationFileConfig field to given value.

### HasAutomationFileConfig

`func (o *AutomationOutputProcessor) HasAutomationFileConfig() bool`

HasAutomationFileConfig returns a boolean if a field has been set.

### GetFile

`func (o *AutomationOutputProcessor) GetFile() AutomationFileFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AutomationOutputProcessor) GetFileOk() (*AutomationFileFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AutomationOutputProcessor) SetFile(v AutomationFileFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AutomationOutputProcessor) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *AutomationOutputProcessor) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *AutomationOutputProcessor) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetId

`func (o *AutomationOutputProcessor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationOutputProcessor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationOutputProcessor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationOutputProcessor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AutomationOutputProcessor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationOutputProcessor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationOutputProcessor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomationOutputProcessor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiURL

`func (o *AutomationOutputProcessor) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AutomationOutputProcessor) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AutomationOutputProcessor) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *AutomationOutputProcessor) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *AutomationOutputProcessor) GetArchiveRecord() ArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *AutomationOutputProcessor) GetArchiveRecordOk() (*ArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *AutomationOutputProcessor) SetArchiveRecord(v ArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *AutomationOutputProcessor) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### GetCompleteWithErrors

`func (o *AutomationOutputProcessor) GetCompleteWithErrors() bool`

GetCompleteWithErrors returns the CompleteWithErrors field if non-nil, zero value otherwise.

### GetCompleteWithErrorsOk

`func (o *AutomationOutputProcessor) GetCompleteWithErrorsOk() (*bool, bool)`

GetCompleteWithErrorsOk returns a tuple with the CompleteWithErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteWithErrors

`func (o *AutomationOutputProcessor) SetCompleteWithErrors(v bool)`

SetCompleteWithErrors sets CompleteWithErrors field to given value.

### HasCompleteWithErrors

`func (o *AutomationOutputProcessor) HasCompleteWithErrors() bool`

HasCompleteWithErrors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationOutputProcessor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationOutputProcessor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationOutputProcessor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationOutputProcessor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorFile

`func (o *AutomationOutputProcessor) GetErrorFile() Blob`

GetErrorFile returns the ErrorFile field if non-nil, zero value otherwise.

### GetErrorFileOk

`func (o *AutomationOutputProcessor) GetErrorFileOk() (*Blob, bool)`

GetErrorFileOk returns a tuple with the ErrorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFile

`func (o *AutomationOutputProcessor) SetErrorFile(v Blob)`

SetErrorFile sets ErrorFile field to given value.

### HasErrorFile

`func (o *AutomationOutputProcessor) HasErrorFile() bool`

HasErrorFile returns a boolean if a field has been set.

### SetErrorFileNil

`func (o *AutomationOutputProcessor) SetErrorFileNil(b bool)`

 SetErrorFileNil sets the value for ErrorFile to be an explicit nil

### UnsetErrorFile
`func (o *AutomationOutputProcessor) UnsetErrorFile()`

UnsetErrorFile ensures that no value is present for ErrorFile, not even an explicit nil
### GetModifiedAt

`func (o *AutomationOutputProcessor) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutomationOutputProcessor) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutomationOutputProcessor) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AutomationOutputProcessor) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetProgressStats

`func (o *AutomationOutputProcessor) GetProgressStats() AutomationProgressStats`

GetProgressStats returns the ProgressStats field if non-nil, zero value otherwise.

### GetProgressStatsOk

`func (o *AutomationOutputProcessor) GetProgressStatsOk() (*AutomationProgressStats, bool)`

GetProgressStatsOk returns a tuple with the ProgressStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressStats

`func (o *AutomationOutputProcessor) SetProgressStats(v AutomationProgressStats)`

SetProgressStats sets ProgressStats field to given value.

### HasProgressStats

`func (o *AutomationOutputProcessor) HasProgressStats() bool`

HasProgressStats returns a boolean if a field has been set.

### GetTransforms

`func (o *AutomationOutputProcessor) GetTransforms() []*LabAutomationTransform`

GetTransforms returns the Transforms field if non-nil, zero value otherwise.

### GetTransformsOk

`func (o *AutomationOutputProcessor) GetTransformsOk() (*[]*LabAutomationTransform, bool)`

GetTransformsOk returns a tuple with the Transforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransforms

`func (o *AutomationOutputProcessor) SetTransforms(v []*LabAutomationTransform)`

SetTransforms sets Transforms field to given value.

### HasTransforms

`func (o *AutomationOutputProcessor) HasTransforms() bool`

HasTransforms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


