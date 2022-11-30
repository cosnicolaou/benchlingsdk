# DropdownCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the dropdown | 
**Options** | [**[]DropdownOptionCreate**](DropdownOptionCreate.md) | Options to set for the dropdown | 
**RegistryId** | Pointer to **NullableString** | ID of registry in which to create the dropdown. Required if multiple registries exist. | [optional] 

## Methods

### NewDropdownCreate

`func NewDropdownCreate(name string, options []DropdownOptionCreate, ) *DropdownCreate`

NewDropdownCreate instantiates a new DropdownCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropdownCreateWithDefaults

`func NewDropdownCreateWithDefaults() *DropdownCreate`

NewDropdownCreateWithDefaults instantiates a new DropdownCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DropdownCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DropdownCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DropdownCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *DropdownCreate) GetOptions() []DropdownOptionCreate`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DropdownCreate) GetOptionsOk() (*[]DropdownOptionCreate, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DropdownCreate) SetOptions(v []DropdownOptionCreate)`

SetOptions sets Options field to given value.


### GetRegistryId

`func (o *DropdownCreate) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *DropdownCreate) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *DropdownCreate) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *DropdownCreate) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *DropdownCreate) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *DropdownCreate) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


