# ContainerUpdateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**Volume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) |  | [optional] 

## Methods

### NewContainerUpdateAllOf

`func NewContainerUpdateAllOf() *ContainerUpdateAllOf`

NewContainerUpdateAllOf instantiates a new ContainerUpdateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerUpdateAllOfWithDefaults

`func NewContainerUpdateAllOfWithDefaults() *ContainerUpdateAllOf`

NewContainerUpdateAllOfWithDefaults instantiates a new ContainerUpdateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ContainerUpdateAllOf) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ContainerUpdateAllOf) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ContainerUpdateAllOf) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ContainerUpdateAllOf) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ContainerUpdateAllOf) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ContainerUpdateAllOf) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetQuantity

`func (o *ContainerUpdateAllOf) GetQuantity() ContainerQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ContainerUpdateAllOf) GetQuantityOk() (*ContainerQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ContainerUpdateAllOf) SetQuantity(v ContainerQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ContainerUpdateAllOf) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetVolume

`func (o *ContainerUpdateAllOf) GetVolume() DeprecatedContainerVolumeForInput`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ContainerUpdateAllOf) GetVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ContainerUpdateAllOf) SetVolume(v DeprecatedContainerVolumeForInput)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ContainerUpdateAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


