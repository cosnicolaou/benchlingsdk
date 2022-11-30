# FeatureLibraryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description for the feature library | 
**Name** | **string** | The name of the feature library | 
**OrganizationId** | Pointer to **string** | The organization that will own the feature library. The requesting user must be an administrator of the organization. If unspecified and the organization allows personal ownables, then the requesting user will own the feature library  | [optional] 

## Methods

### NewFeatureLibraryCreate

`func NewFeatureLibraryCreate(description string, name string, ) *FeatureLibraryCreate`

NewFeatureLibraryCreate instantiates a new FeatureLibraryCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureLibraryCreateWithDefaults

`func NewFeatureLibraryCreateWithDefaults() *FeatureLibraryCreate`

NewFeatureLibraryCreateWithDefaults instantiates a new FeatureLibraryCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FeatureLibraryCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureLibraryCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureLibraryCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *FeatureLibraryCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureLibraryCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureLibraryCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *FeatureLibraryCreate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FeatureLibraryCreate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FeatureLibraryCreate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FeatureLibraryCreate) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


