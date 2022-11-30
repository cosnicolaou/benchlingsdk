# UserValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationComment** | Pointer to **string** | A string explaining the reason for the provided validation status. | [optional] 
**ValidationStatus** | Pointer to **string** | Valid values for this enum depend on whether it is used to set a value (e.g., in a POST request), or is a return value for an existing result. Acceptable values for setting a status are &#39;VALID&#39; or &#39;INVALID&#39;. Possible return values are &#39;VALID&#39;, &#39;INVALID&#39;, or &#39;PARTIALLY_VALID&#39; (a result will be partially valid if it has some valid fields and some invalid fields).  | [optional] 

## Methods

### NewUserValidation

`func NewUserValidation() *UserValidation`

NewUserValidation instantiates a new UserValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserValidationWithDefaults

`func NewUserValidationWithDefaults() *UserValidation`

NewUserValidationWithDefaults instantiates a new UserValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationComment

`func (o *UserValidation) GetValidationComment() string`

GetValidationComment returns the ValidationComment field if non-nil, zero value otherwise.

### GetValidationCommentOk

`func (o *UserValidation) GetValidationCommentOk() (*string, bool)`

GetValidationCommentOk returns a tuple with the ValidationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationComment

`func (o *UserValidation) SetValidationComment(v string)`

SetValidationComment sets ValidationComment field to given value.

### HasValidationComment

`func (o *UserValidation) HasValidationComment() bool`

HasValidationComment returns a boolean if a field has been set.

### GetValidationStatus

`func (o *UserValidation) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *UserValidation) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *UserValidation) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *UserValidation) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


