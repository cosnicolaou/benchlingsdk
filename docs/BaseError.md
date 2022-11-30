# BaseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseError

`func NewBaseError() *BaseError`

NewBaseError instantiates a new BaseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseErrorWithDefaults

`func NewBaseErrorWithDefaults() *BaseError`

NewBaseErrorWithDefaults instantiates a new BaseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BaseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BaseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BaseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BaseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *BaseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserMessage

`func (o *BaseError) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *BaseError) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *BaseError) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *BaseError) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


