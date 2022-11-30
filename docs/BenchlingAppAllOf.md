# BenchlingAppAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **string** |  | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableArchiveRecord**](ArchiveRecord.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the app was created | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime at which the the app was last modified | [optional] [readonly] 
**WebUrl** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBenchlingAppAllOf

`func NewBenchlingAppAllOf() *BenchlingAppAllOf`

NewBenchlingAppAllOf instantiates a new BenchlingAppAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchlingAppAllOfWithDefaults

`func NewBenchlingAppAllOfWithDefaults() *BenchlingAppAllOf`

NewBenchlingAppAllOfWithDefaults instantiates a new BenchlingAppAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *BenchlingAppAllOf) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *BenchlingAppAllOf) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *BenchlingAppAllOf) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *BenchlingAppAllOf) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *BenchlingAppAllOf) GetArchiveRecord() ArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *BenchlingAppAllOf) GetArchiveRecordOk() (*ArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *BenchlingAppAllOf) SetArchiveRecord(v ArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *BenchlingAppAllOf) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *BenchlingAppAllOf) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *BenchlingAppAllOf) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *BenchlingAppAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BenchlingAppAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BenchlingAppAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BenchlingAppAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BenchlingAppAllOf) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BenchlingAppAllOf) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BenchlingAppAllOf) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BenchlingAppAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetId

`func (o *BenchlingAppAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BenchlingAppAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BenchlingAppAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BenchlingAppAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BenchlingAppAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BenchlingAppAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BenchlingAppAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BenchlingAppAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetWebUrl

`func (o *BenchlingAppAllOf) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *BenchlingAppAllOf) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *BenchlingAppAllOf) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *BenchlingAppAllOf) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


