# AppConfigItemsBulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConfigurationItems** | [**[]AppConfigItemCreate**](AppConfigItemCreate.md) |  | 

## Methods

### NewAppConfigItemsBulkCreateRequest

`func NewAppConfigItemsBulkCreateRequest(appConfigurationItems []AppConfigItemCreate, ) *AppConfigItemsBulkCreateRequest`

NewAppConfigItemsBulkCreateRequest instantiates a new AppConfigItemsBulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigItemsBulkCreateRequestWithDefaults

`func NewAppConfigItemsBulkCreateRequestWithDefaults() *AppConfigItemsBulkCreateRequest`

NewAppConfigItemsBulkCreateRequestWithDefaults instantiates a new AppConfigItemsBulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConfigurationItems

`func (o *AppConfigItemsBulkCreateRequest) GetAppConfigurationItems() []AppConfigItemCreate`

GetAppConfigurationItems returns the AppConfigurationItems field if non-nil, zero value otherwise.

### GetAppConfigurationItemsOk

`func (o *AppConfigItemsBulkCreateRequest) GetAppConfigurationItemsOk() (*[]AppConfigItemCreate, bool)`

GetAppConfigurationItemsOk returns a tuple with the AppConfigurationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConfigurationItems

`func (o *AppConfigItemsBulkCreateRequest) SetAppConfigurationItems(v []AppConfigItemCreate)`

SetAppConfigurationItems sets AppConfigurationItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


