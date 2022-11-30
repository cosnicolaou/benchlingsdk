# AutomationOutputProcessorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Automation Output Processor in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**ArchiveRecord**](ArchiveRecord.md) |  | [optional] 
**CompleteWithErrors** | Pointer to **bool** | Specifies whether file processing should complete with errors. False means any error in output file processing will result in no actions being committed. True means that if row-level errors occur, then failing rows and their errors will be saved to errorFile, and actions from successful rows will be committed. | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Automation Output Processor was created | [optional] 
**ErrorFile** | Pointer to [**NullableBlob**](Blob.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Automation Output Processor was last modified | [optional] 
**ProgressStats** | Pointer to [**AutomationProgressStats**](AutomationProgressStats.md) |  | [optional] 
**Transforms** | Pointer to [**[]LabAutomationTransform**](LabAutomationTransform.md) |  | [optional] 

## Methods

### NewAutomationOutputProcessorAllOf

`func NewAutomationOutputProcessorAllOf() *AutomationOutputProcessorAllOf`

NewAutomationOutputProcessorAllOf instantiates a new AutomationOutputProcessorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationOutputProcessorAllOfWithDefaults

`func NewAutomationOutputProcessorAllOfWithDefaults() *AutomationOutputProcessorAllOf`

NewAutomationOutputProcessorAllOfWithDefaults instantiates a new AutomationOutputProcessorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *AutomationOutputProcessorAllOf) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AutomationOutputProcessorAllOf) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AutomationOutputProcessorAllOf) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *AutomationOutputProcessorAllOf) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *AutomationOutputProcessorAllOf) GetArchiveRecord() ArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *AutomationOutputProcessorAllOf) GetArchiveRecordOk() (*ArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *AutomationOutputProcessorAllOf) SetArchiveRecord(v ArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *AutomationOutputProcessorAllOf) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### GetCompleteWithErrors

`func (o *AutomationOutputProcessorAllOf) GetCompleteWithErrors() bool`

GetCompleteWithErrors returns the CompleteWithErrors field if non-nil, zero value otherwise.

### GetCompleteWithErrorsOk

`func (o *AutomationOutputProcessorAllOf) GetCompleteWithErrorsOk() (*bool, bool)`

GetCompleteWithErrorsOk returns a tuple with the CompleteWithErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteWithErrors

`func (o *AutomationOutputProcessorAllOf) SetCompleteWithErrors(v bool)`

SetCompleteWithErrors sets CompleteWithErrors field to given value.

### HasCompleteWithErrors

`func (o *AutomationOutputProcessorAllOf) HasCompleteWithErrors() bool`

HasCompleteWithErrors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationOutputProcessorAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationOutputProcessorAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationOutputProcessorAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationOutputProcessorAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorFile

`func (o *AutomationOutputProcessorAllOf) GetErrorFile() Blob`

GetErrorFile returns the ErrorFile field if non-nil, zero value otherwise.

### GetErrorFileOk

`func (o *AutomationOutputProcessorAllOf) GetErrorFileOk() (*Blob, bool)`

GetErrorFileOk returns a tuple with the ErrorFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFile

`func (o *AutomationOutputProcessorAllOf) SetErrorFile(v Blob)`

SetErrorFile sets ErrorFile field to given value.

### HasErrorFile

`func (o *AutomationOutputProcessorAllOf) HasErrorFile() bool`

HasErrorFile returns a boolean if a field has been set.

### SetErrorFileNil

`func (o *AutomationOutputProcessorAllOf) SetErrorFileNil(b bool)`

 SetErrorFileNil sets the value for ErrorFile to be an explicit nil

### UnsetErrorFile
`func (o *AutomationOutputProcessorAllOf) UnsetErrorFile()`

UnsetErrorFile ensures that no value is present for ErrorFile, not even an explicit nil
### GetId

`func (o *AutomationOutputProcessorAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationOutputProcessorAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationOutputProcessorAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationOutputProcessorAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AutomationOutputProcessorAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutomationOutputProcessorAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutomationOutputProcessorAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AutomationOutputProcessorAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetProgressStats

`func (o *AutomationOutputProcessorAllOf) GetProgressStats() AutomationProgressStats`

GetProgressStats returns the ProgressStats field if non-nil, zero value otherwise.

### GetProgressStatsOk

`func (o *AutomationOutputProcessorAllOf) GetProgressStatsOk() (*AutomationProgressStats, bool)`

GetProgressStatsOk returns a tuple with the ProgressStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressStats

`func (o *AutomationOutputProcessorAllOf) SetProgressStats(v AutomationProgressStats)`

SetProgressStats sets ProgressStats field to given value.

### HasProgressStats

`func (o *AutomationOutputProcessorAllOf) HasProgressStats() bool`

HasProgressStats returns a boolean if a field has been set.

### GetTransforms

`func (o *AutomationOutputProcessorAllOf) GetTransforms() []*LabAutomationTransform`

GetTransforms returns the Transforms field if non-nil, zero value otherwise.

### GetTransformsOk

`func (o *AutomationOutputProcessorAllOf) GetTransformsOk() (*[]*LabAutomationTransform, bool)`

GetTransformsOk returns a tuple with the Transforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransforms

`func (o *AutomationOutputProcessorAllOf) SetTransforms(v []*LabAutomationTransform)`

SetTransforms sets Transforms field to given value.

### HasTransforms

`func (o *AutomationOutputProcessorAllOf) HasTransforms() bool`

HasTransforms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


