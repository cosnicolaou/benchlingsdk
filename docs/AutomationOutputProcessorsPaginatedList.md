# AutomationOutputProcessorsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomationOutputProcessors** | Pointer to [**[]AutomationOutputProcessor**](AutomationOutputProcessor.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAutomationOutputProcessorsPaginatedList

`func NewAutomationOutputProcessorsPaginatedList() *AutomationOutputProcessorsPaginatedList`

NewAutomationOutputProcessorsPaginatedList instantiates a new AutomationOutputProcessorsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationOutputProcessorsPaginatedListWithDefaults

`func NewAutomationOutputProcessorsPaginatedListWithDefaults() *AutomationOutputProcessorsPaginatedList`

NewAutomationOutputProcessorsPaginatedListWithDefaults instantiates a new AutomationOutputProcessorsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomationOutputProcessors

`func (o *AutomationOutputProcessorsPaginatedList) GetAutomationOutputProcessors() []AutomationOutputProcessor`

GetAutomationOutputProcessors returns the AutomationOutputProcessors field if non-nil, zero value otherwise.

### GetAutomationOutputProcessorsOk

`func (o *AutomationOutputProcessorsPaginatedList) GetAutomationOutputProcessorsOk() (*[]AutomationOutputProcessor, bool)`

GetAutomationOutputProcessorsOk returns a tuple with the AutomationOutputProcessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationOutputProcessors

`func (o *AutomationOutputProcessorsPaginatedList) SetAutomationOutputProcessors(v []AutomationOutputProcessor)`

SetAutomationOutputProcessors sets AutomationOutputProcessors field to given value.

### HasAutomationOutputProcessors

`func (o *AutomationOutputProcessorsPaginatedList) HasAutomationOutputProcessors() bool`

HasAutomationOutputProcessors returns a boolean if a field has been set.

### GetNextToken

`func (o *AutomationOutputProcessorsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AutomationOutputProcessorsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AutomationOutputProcessorsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AutomationOutputProcessorsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


