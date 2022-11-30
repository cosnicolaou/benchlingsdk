# ConflictErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 
**Conflicts** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewConflictErrorError

`func NewConflictErrorError() *ConflictErrorError`

NewConflictErrorError instantiates a new ConflictErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictErrorErrorWithDefaults

`func NewConflictErrorErrorWithDefaults() *ConflictErrorError`

NewConflictErrorErrorWithDefaults instantiates a new ConflictErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConflictErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConflictErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConflictErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConflictErrorError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ConflictErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConflictErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConflictErrorError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConflictErrorError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserMessage

`func (o *ConflictErrorError) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *ConflictErrorError) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *ConflictErrorError) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *ConflictErrorError) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.

### GetConflicts

`func (o *ConflictErrorError) GetConflicts() []map[string]interface{}`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *ConflictErrorError) GetConflictsOk() (*[]map[string]interface{}, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *ConflictErrorError) SetConflicts(v []map[string]interface{})`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *ConflictErrorError) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


