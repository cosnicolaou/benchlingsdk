# AppConfigurationPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**AppConfigurationItems** | Pointer to [**[]AppConfigItem**](AppConfigItem.md) |  | [optional] 

## Methods

### NewAppConfigurationPaginatedList

`func NewAppConfigurationPaginatedList() *AppConfigurationPaginatedList`

NewAppConfigurationPaginatedList instantiates a new AppConfigurationPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConfigurationPaginatedListWithDefaults

`func NewAppConfigurationPaginatedListWithDefaults() *AppConfigurationPaginatedList`

NewAppConfigurationPaginatedListWithDefaults instantiates a new AppConfigurationPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *AppConfigurationPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AppConfigurationPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AppConfigurationPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AppConfigurationPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetAppConfigurationItems

`func (o *AppConfigurationPaginatedList) GetAppConfigurationItems() []AppConfigItem`

GetAppConfigurationItems returns the AppConfigurationItems field if non-nil, zero value otherwise.

### GetAppConfigurationItemsOk

`func (o *AppConfigurationPaginatedList) GetAppConfigurationItemsOk() (*[]AppConfigItem, bool)`

GetAppConfigurationItemsOk returns a tuple with the AppConfigurationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConfigurationItems

`func (o *AppConfigurationPaginatedList) SetAppConfigurationItems(v []AppConfigItem)`

SetAppConfigurationItems sets AppConfigurationItems field to given value.

### HasAppConfigurationItems

`func (o *AppConfigurationPaginatedList) HasAppConfigurationItems() bool`

HasAppConfigurationItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


