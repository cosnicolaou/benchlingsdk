# AppConfigItemCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | App id to which this config item belongs. | 
**Path** | **[]string** | Array-based representation of config item&#39;s location in the tree in order from top to bottom. | 
**Type** | **string** |  | 
**Value** | **string** | The value of a json create object should be json parseable. | 

## Methods

### NewAppConfigItemCreate

`func NewAppConfigItemCreate(appId string, path []string, type_ string, value string, ) *AppConfigItemCreate`

NewAppConfigItemCreate instantiates a new AppConfigItemCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigItemCreateWithDefaults

`func NewAppConfigItemCreateWithDefaults() *AppConfigItemCreate`

NewAppConfigItemCreateWithDefaults instantiates a new AppConfigItemCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppConfigItemCreate) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppConfigItemCreate) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppConfigItemCreate) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetPath

`func (o *AppConfigItemCreate) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AppConfigItemCreate) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AppConfigItemCreate) SetPath(v []string)`

SetPath sets Path field to given value.


### GetType

`func (o *AppConfigItemCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppConfigItemCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppConfigItemCreate) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AppConfigItemCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppConfigItemCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppConfigItemCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


