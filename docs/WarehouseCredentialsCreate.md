# WarehouseCredentialsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | **int32** | Duration, in seconds, that credentials should be active for. Must be greater than 0 and less than 3600.  | 

## Methods

### NewWarehouseCredentialsCreate

`func NewWarehouseCredentialsCreate(expiresIn int32, ) *WarehouseCredentialsCreate`

NewWarehouseCredentialsCreate instantiates a new WarehouseCredentialsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseCredentialsCreateWithDefaults

`func NewWarehouseCredentialsCreateWithDefaults() *WarehouseCredentialsCreate`

NewWarehouseCredentialsCreateWithDefaults instantiates a new WarehouseCredentialsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *WarehouseCredentialsCreate) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *WarehouseCredentialsCreate) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *WarehouseCredentialsCreate) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


