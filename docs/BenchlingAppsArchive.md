# BenchlingAppsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | **[]string** | Array of app IDs | 
**Reason** | **string** | Reason that apps are being archived. Actual reason enum varies by tenant. | 

## Methods

### NewBenchlingAppsArchive

`func NewBenchlingAppsArchive(appIds []string, reason string, ) *BenchlingAppsArchive`

NewBenchlingAppsArchive instantiates a new BenchlingAppsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchlingAppsArchiveWithDefaults

`func NewBenchlingAppsArchiveWithDefaults() *BenchlingAppsArchive`

NewBenchlingAppsArchiveWithDefaults instantiates a new BenchlingAppsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *BenchlingAppsArchive) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *BenchlingAppsArchive) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *BenchlingAppsArchive) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.


### GetReason

`func (o *BenchlingAppsArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BenchlingAppsArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BenchlingAppsArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


