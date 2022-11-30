# UserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**PasswordLastChangedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserAllOf

`func NewUserAllOf() *UserAllOf`

NewUserAllOf instantiates a new UserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAllOfWithDefaults

`func NewUserAllOfWithDefaults() *UserAllOf`

NewUserAllOfWithDefaults instantiates a new UserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserAllOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAllOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAllOf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsSuspended

`func (o *UserAllOf) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *UserAllOf) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *UserAllOf) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *UserAllOf) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetPasswordLastChangedAt

`func (o *UserAllOf) GetPasswordLastChangedAt() string`

GetPasswordLastChangedAt returns the PasswordLastChangedAt field if non-nil, zero value otherwise.

### GetPasswordLastChangedAtOk

`func (o *UserAllOf) GetPasswordLastChangedAtOk() (*string, bool)`

GetPasswordLastChangedAtOk returns a tuple with the PasswordLastChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChangedAt

`func (o *UserAllOf) SetPasswordLastChangedAt(v string)`

SetPasswordLastChangedAt sets PasswordLastChangedAt field to given value.

### HasPasswordLastChangedAt

`func (o *UserAllOf) HasPasswordLastChangedAt() bool`

HasPasswordLastChangedAt returns a boolean if a field has been set.

### SetPasswordLastChangedAtNil

`func (o *UserAllOf) SetPasswordLastChangedAtNil(b bool)`

 SetPasswordLastChangedAtNil sets the value for PasswordLastChangedAt to be an explicit nil

### UnsetPasswordLastChangedAt
`func (o *UserAllOf) UnsetPasswordLastChangedAt()`

UnsetPasswordLastChangedAt ensures that no value is present for PasswordLastChangedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


