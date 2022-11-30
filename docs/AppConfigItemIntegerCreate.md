# AppConfigItemIntegerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | App id to which this config item belongs. | 
**Path** | **[]string** | Array-based representation of config item&#39;s location in the tree in order from top to bottom. | 
**Type** | **string** |  | 
**Value** | **int32** |  | 

## Methods

### NewAppConfigItemIntegerCreate

`func NewAppConfigItemIntegerCreate(appId string, path []string, type_ string, value int32, ) *AppConfigItemIntegerCreate`

NewAppConfigItemIntegerCreate instantiates a new AppConfigItemIntegerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigItemIntegerCreateWithDefaults

`func NewAppConfigItemIntegerCreateWithDefaults() *AppConfigItemIntegerCreate`

NewAppConfigItemIntegerCreateWithDefaults instantiates a new AppConfigItemIntegerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppConfigItemIntegerCreate) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppConfigItemIntegerCreate) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppConfigItemIntegerCreate) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetPath

`func (o *AppConfigItemIntegerCreate) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppConfigItemIntegerCreate) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppConfigItemIntegerCreate) SetPath(v []string)`

SetPath sets Path field to given value.


### GetType

`func (o *AppConfigItemIntegerCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppConfigItemIntegerCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppConfigItemIntegerCreate) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AppConfigItemIntegerCreate) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppConfigItemIntegerCreate) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppConfigItemIntegerCreate) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


