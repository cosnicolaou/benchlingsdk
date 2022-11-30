# FeatureLibraryCreateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** | The organization that will own the feature library. The requesting user must be an administrator of the organization. If unspecified and the organization allows personal ownables, then the requesting user will own the feature library  | [optional] 

## Methods

### NewFeatureLibraryCreateAllOf

`func NewFeatureLibraryCreateAllOf() *FeatureLibraryCreateAllOf`

NewFeatureLibraryCreateAllOf instantiates a new FeatureLibraryCreateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureLibraryCreateAllOfWithDefaults

`func NewFeatureLibraryCreateAllOfWithDefaults() *FeatureLibraryCreateAllOf`

NewFeatureLibraryCreateAllOfWithDefaults instantiates a new FeatureLibraryCreateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *FeatureLibraryCreateAllOf) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FeatureLibraryCreateAllOf) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FeatureLibraryCreateAllOf) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FeatureLibraryCreateAllOf) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


