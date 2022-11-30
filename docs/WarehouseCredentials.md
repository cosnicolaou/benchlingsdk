# WarehouseCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time after which new connections using the username/password will not be permitted. Upon expiration, currently open connections are not terminated.  | [optional] [readonly] 
**Password** | Pointer to **string** | The password to connect to the warehouse. | [optional] [readonly] 
**Username** | Pointer to **string** | The username to connect to the warehouse. | [optional] [readonly] 

## Methods

### NewWarehouseCredentials

`func NewWarehouseCredentials() *WarehouseCredentials`

NewWarehouseCredentials instantiates a new WarehouseCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseCredentialsWithDefaults

`func NewWarehouseCredentialsWithDefaults() *WarehouseCredentials`

NewWarehouseCredentialsWithDefaults instantiates a new WarehouseCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *WarehouseCredentials) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *WarehouseCredentials) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *WarehouseCredentials) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *WarehouseCredentials) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetPassword

`func (o *WarehouseCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WarehouseCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WarehouseCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WarehouseCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *WarehouseCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WarehouseCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WarehouseCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WarehouseCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


