# RequestSchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModifiedAt** | Pointer to **time.Time** | DateTime the Request Schema was last modified | [optional] 
**Organization** | Pointer to [**RequestSchemaAllOfOrganization**](RequestSchemaAllOfOrganization.md) |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRequestSchemaAllOf

`func NewRequestSchemaAllOf() *RequestSchemaAllOf`

NewRequestSchemaAllOf instantiates a new RequestSchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSchemaAllOfWithDefaults

`func NewRequestSchemaAllOfWithDefaults() *RequestSchemaAllOf`

NewRequestSchemaAllOfWithDefaults instantiates a new RequestSchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifiedAt

`func (o *RequestSchemaAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RequestSchemaAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RequestSchemaAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RequestSchemaAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetOrganization

`func (o *RequestSchemaAllOf) GetOrganization() RequestSchemaAllOfOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RequestSchemaAllOf) GetOrganizationOk() (*RequestSchemaAllOfOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RequestSchemaAllOf) SetOrganization(v RequestSchemaAllOfOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RequestSchemaAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSystemName

`func (o *RequestSchemaAllOf) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *RequestSchemaAllOf) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *RequestSchemaAllOf) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *RequestSchemaAllOf) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetType

`func (o *RequestSchemaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestSchemaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestSchemaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RequestSchemaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


