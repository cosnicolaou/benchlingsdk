# FloatAppConfigItem

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
**Value** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewFloatAppConfigItem

`func NewFloatAppConfigItem() *FloatAppConfigItem`

NewFloatAppConfigItem instantiates a new FloatAppConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatAppConfigItemWithDefaults

`func NewFloatAppConfigItemWithDefaults() *FloatAppConfigItem`

NewFloatAppConfigItemWithDefaults instantiates a new FloatAppConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *FloatAppConfigItem) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *FloatAppConfigItem) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *FloatAppConfigItem) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *FloatAppConfigItem) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetApp

`func (o *FloatAppConfigItem) GetApp() AppConfigItemApiMixinApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *FloatAppConfigItem) GetAppOk() (*AppConfigItemApiMixinApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *FloatAppConfigItem) SetApp(v AppConfigItemApiMixinApp)`

SetApp sets App field to given value.

### HasApp

`func (o *FloatAppConfigItem) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *FloatAppConfigItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatAppConfigItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatAppConfigItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FloatAppConfigItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *FloatAppConfigItem) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FloatAppConfigItem) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FloatAppConfigItem) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *FloatAppConfigItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPath

`func (o *FloatAppConfigItem) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FloatAppConfigItem) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FloatAppConfigItem) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FloatAppConfigItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *FloatAppConfigItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FloatAppConfigItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FloatAppConfigItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FloatAppConfigItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *FloatAppConfigItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FloatAppConfigItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FloatAppConfigItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FloatAppConfigItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredConfig

`func (o *FloatAppConfigItem) GetRequiredConfig() bool`

GetRequiredConfig returns the RequiredConfig field if non-nil, zero value otherwise.

### GetRequiredConfigOk

`func (o *FloatAppConfigItem) GetRequiredConfigOk() (*bool, bool)`

GetRequiredConfigOk returns a tuple with the RequiredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfig

`func (o *FloatAppConfigItem) SetRequiredConfig(v bool)`

SetRequiredConfig sets RequiredConfig field to given value.

### HasRequiredConfig

`func (o *FloatAppConfigItem) HasRequiredConfig() bool`

HasRequiredConfig returns a boolean if a field has been set.

### GetValue

`func (o *FloatAppConfigItem) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FloatAppConfigItem) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FloatAppConfigItem) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *FloatAppConfigItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FloatAppConfigItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FloatAppConfigItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


