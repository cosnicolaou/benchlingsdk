# Registry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Registry was last modified | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewRegistry

`func NewRegistry() *Registry`

NewRegistry instantiates a new Registry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWithDefaults

`func NewRegistryWithDefaults() *Registry`

NewRegistryWithDefaults instantiates a new Registry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Registry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Registry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Registry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Registry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Registry) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Registry) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Registry) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Registry) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Registry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Registry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Registry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Registry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *Registry) GetOwner() Organization`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Registry) GetOwnerOk() (*Organization, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Registry) SetOwner(v Organization)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Registry) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


