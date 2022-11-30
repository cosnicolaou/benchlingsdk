# FeatureLibrary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description for the feature library | [optional] 
**Name** | Pointer to **string** | The name of the feature library | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Feature Library was created | [optional] 
**Id** | Pointer to **string** | The id of the feature library | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Feature Library was last modified | [optional] 
**Owner** | Pointer to [**NullableAnyOfOrganizationUserSummary**](anyOf&lt;Organization,UserSummary&gt;.md) |  | [optional] 
**WebURL** | Pointer to **string** | The Benchling web UI url to view the Feature Library | [optional] 

## Methods

### NewFeatureLibrary

`func NewFeatureLibrary() *FeatureLibrary`

NewFeatureLibrary instantiates a new FeatureLibrary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureLibraryWithDefaults

`func NewFeatureLibraryWithDefaults() *FeatureLibrary`

NewFeatureLibraryWithDefaults instantiates a new FeatureLibrary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FeatureLibrary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureLibrary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureLibrary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureLibrary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FeatureLibrary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureLibrary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureLibrary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureLibrary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeatureLibrary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeatureLibrary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeatureLibrary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeatureLibrary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *FeatureLibrary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureLibrary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureLibrary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureLibrary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *FeatureLibrary) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FeatureLibrary) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FeatureLibrary) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *FeatureLibrary) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetOwner

`func (o *FeatureLibrary) GetOwner() AnyOfOrganizationUserSummary`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FeatureLibrary) GetOwnerOk() (*AnyOfOrganizationUserSummary, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FeatureLibrary) SetOwner(v AnyOfOrganizationUserSummary)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FeatureLibrary) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *FeatureLibrary) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *FeatureLibrary) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetWebURL

`func (o *FeatureLibrary) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *FeatureLibrary) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *FeatureLibrary) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *FeatureLibrary) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


