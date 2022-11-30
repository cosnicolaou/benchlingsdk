# AutomationInputGeneratorUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **NullableString** | The ID of the file (blob) associated with the input generator. Set to &#x60;null&#x60; to remove an existing file from the generator. | [optional] 

## Methods

### NewAutomationInputGeneratorUpdate

`func NewAutomationInputGeneratorUpdate() *AutomationInputGeneratorUpdate`

NewAutomationInputGeneratorUpdate instantiates a new AutomationInputGeneratorUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationInputGeneratorUpdateWithDefaults

`func NewAutomationInputGeneratorUpdateWithDefaults() *AutomationInputGeneratorUpdate`

NewAutomationInputGeneratorUpdateWithDefaults instantiates a new AutomationInputGeneratorUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *AutomationInputGeneratorUpdate) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AutomationInputGeneratorUpdate) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AutomationInputGeneratorUpdate) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AutomationInputGeneratorUpdate) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *AutomationInputGeneratorUpdate) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *AutomationInputGeneratorUpdate) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


