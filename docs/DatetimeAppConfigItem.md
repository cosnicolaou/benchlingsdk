# DatetimeAppConfigItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** |  | [optional] [readonly] 
**App** | Pointer to [**AppConfigItemApiMixinApp**](AppConfigItemApiMixinApp.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the app config item was last modified | [optional] [readonly] 
**Path** | Pointer to **[]string** | Array-based representation of config item&#39;s location in the tree in order from top to bottom. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequiredConfig** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDatetimeAppConfigItem

`func NewDatetimeAppConfigItem() *DatetimeAppConfigItem`

NewDatetimeAppConfigItem instantiates a new DatetimeAppConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatetimeAppConfigItemWithDefaults

`func NewDatetimeAppConfigItemWithDefaults() *DatetimeAppConfigItem`

NewDatetimeAppConfigItemWithDefaults instantiates a new DatetimeAppConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *DatetimeAppConfigItem) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *DatetimeAppConfigItem) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *DatetimeAppConfigItem) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *DatetimeAppConfigItem) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetApp

`func (o *DatetimeAppConfigItem) GetApp() AppConfigItemApiMixinApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *DatetimeAppConfigItem) GetAppOk() (*AppConfigItemApiMixinApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *DatetimeAppConfigItem) SetApp(v AppConfigItemApiMixinApp)`

SetApp sets App field to given value.

### HasApp

`func (o *DatetimeAppConfigItem) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *DatetimeAppConfigItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatetimeAppConfigItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatetimeAppConfigItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatetimeAppConfigItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *DatetimeAppConfigItem) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DatetimeAppConfigItem) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DatetimeAppConfigItem) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *DatetimeAppConfigItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPath

`func (o *DatetimeAppConfigItem) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DatetimeAppConfigItem) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DatetimeAppConfigItem) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DatetimeAppConfigItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *DatetimeAppConfigItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatetimeAppConfigItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatetimeAppConfigItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatetimeAppConfigItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *DatetimeAppConfigItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatetimeAppConfigItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatetimeAppConfigItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatetimeAppConfigItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredConfig

`func (o *DatetimeAppConfigItem) GetRequiredConfig() bool`

GetRequiredConfig returns the RequiredConfig field if non-nil, zero value otherwise.

### GetRequiredConfigOk

`func (o *DatetimeAppConfigItem) GetRequiredConfigOk() (*bool, bool)`

GetRequiredConfigOk returns a tuple with the RequiredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfig

`func (o *DatetimeAppConfigItem) SetRequiredConfig(v bool)`

SetRequiredConfig sets RequiredConfig field to given value.

### HasRequiredConfig

`func (o *DatetimeAppConfigItem) HasRequiredConfig() bool`

HasRequiredConfig returns a boolean if a field has been set.

### GetValue

`func (o *DatetimeAppConfigItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DatetimeAppConfigItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DatetimeAppConfigItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DatetimeAppConfigItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DatetimeAppConfigItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DatetimeAppConfigItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


