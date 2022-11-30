# AppConfigItemApiMixin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** |  | [optional] [readonly] 
**App** | Pointer to [**AppConfigItemApiMixinApp**](AppConfigItemApiMixinApp.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the app config item was last modified | [optional] [readonly] 
**Path** | Pointer to **[]string** | Array-based representation of config item&#39;s location in the tree in order from top to bottom. | [optional] 
**Type** | Pointer to **string** | Type of the app config item | [optional] 

## Methods

### NewAppConfigItemApiMixin

`func NewAppConfigItemApiMixin() *AppConfigItemApiMixin`

NewAppConfigItemApiMixin instantiates a new AppConfigItemApiMixin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigItemApiMixinWithDefaults

`func NewAppConfigItemApiMixinWithDefaults() *AppConfigItemApiMixin`

NewAppConfigItemApiMixinWithDefaults instantiates a new AppConfigItemApiMixin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *AppConfigItemApiMixin) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AppConfigItemApiMixin) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AppConfigItemApiMixin) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *AppConfigItemApiMixin) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetApp

`func (o *AppConfigItemApiMixin) GetApp() AppConfigItemApiMixinApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AppConfigItemApiMixin) GetAppOk() (*AppConfigItemApiMixinApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AppConfigItemApiMixin) SetApp(v AppConfigItemApiMixinApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AppConfigItemApiMixin) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *AppConfigItemApiMixin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppConfigItemApiMixin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppConfigItemApiMixin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppConfigItemApiMixin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AppConfigItemApiMixin) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AppConfigItemApiMixin) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AppConfigItemApiMixin) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AppConfigItemApiMixin) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPath

`func (o *AppConfigItemApiMixin) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppConfigItemApiMixin) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppConfigItemApiMixin) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AppConfigItemApiMixin) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *AppConfigItemApiMixin) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppConfigItemApiMixin) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppConfigItemApiMixin) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppConfigItemApiMixin) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


