# FeatureLibraryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | DateTime the Feature Library was created | [optional] 
**Id** | Pointer to **string** | The id of the feature library | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Feature Library was last modified | [optional] 
**Owner** | Pointer to [**NullableAnyOfOrganizationUserSummary**](anyOf&lt;Organization,UserSummary&gt;.md) |  | [optional] 
**WebURL** | Pointer to **string** | The Benchling web UI url to view the Feature Library | [optional] 

## Methods

### NewFeatureLibraryAllOf

`func NewFeatureLibraryAllOf() *FeatureLibraryAllOf`

NewFeatureLibraryAllOf instantiates a new FeatureLibraryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureLibraryAllOfWithDefaults

`func NewFeatureLibraryAllOfWithDefaults() *FeatureLibraryAllOf`

NewFeatureLibraryAllOfWithDefaults instantiates a new FeatureLibraryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FeatureLibraryAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeatureLibraryAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeatureLibraryAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeatureLibraryAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *FeatureLibraryAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureLibraryAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureLibraryAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureLibraryAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *FeatureLibraryAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FeatureLibraryAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FeatureLibraryAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *FeatureLibraryAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetOwner

`func (o *FeatureLibraryAllOf) GetOwner() AnyOfOrganizationUserSummary`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FeatureLibraryAllOf) GetOwnerOk() (*AnyOfOrganizationUserSummary, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FeatureLibraryAllOf) SetOwner(v AnyOfOrganizationUserSummary)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FeatureLibraryAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *FeatureLibraryAllOf) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *FeatureLibraryAllOf) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetWebURL

`func (o *FeatureLibraryAllOf) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *FeatureLibraryAllOf) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *FeatureLibraryAllOf) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *FeatureLibraryAllOf) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


