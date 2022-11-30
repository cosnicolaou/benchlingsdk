# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**PasswordLastChangedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *User) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *User) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *User) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *User) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsSuspended

`func (o *User) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *User) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *User) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *User) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetPasswordLastChangedAt

`func (o *User) GetPasswordLastChangedAt() string`

GetPasswordLastChangedAt returns the PasswordLastChangedAt field if non-nil, zero value otherwise.

### GetPasswordLastChangedAtOk

`func (o *User) GetPasswordLastChangedAtOk() (*string, bool)`

GetPasswordLastChangedAtOk returns a tuple with the PasswordLastChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChangedAt

`func (o *User) SetPasswordLastChangedAt(v string)`

SetPasswordLastChangedAt sets PasswordLastChangedAt field to given value.

### HasPasswordLastChangedAt

`func (o *User) HasPasswordLastChangedAt() bool`

HasPasswordLastChangedAt returns a boolean if a field has been set.

### SetPasswordLastChangedAtNil

`func (o *User) SetPasswordLastChangedAtNil(b bool)`

 SetPasswordLastChangedAtNil sets the value for PasswordLastChangedAt to be an explicit nil

### UnsetPasswordLastChangedAt
`func (o *User) UnsetPasswordLastChangedAt()`

UnsetPasswordLastChangedAt ensures that no value is present for PasswordLastChangedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


