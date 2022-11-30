# BenchlingApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableArchiveRecord**](ArchiveRecord.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the app was created | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime at which the the app was last modified | [optional] [readonly] 
**WebUrl** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBenchlingApp

`func NewBenchlingApp() *BenchlingApp`

NewBenchlingApp instantiates a new BenchlingApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchlingAppWithDefaults

`func NewBenchlingAppWithDefaults() *BenchlingApp`

NewBenchlingAppWithDefaults instantiates a new BenchlingApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BenchlingApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BenchlingApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BenchlingApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BenchlingApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *BenchlingApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BenchlingApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BenchlingApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BenchlingApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiUrl

`func (o *BenchlingApp) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *BenchlingApp) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *BenchlingApp) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *BenchlingApp) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *BenchlingApp) GetArchiveRecord() ArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *BenchlingApp) GetArchiveRecordOk() (*ArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *BenchlingApp) SetArchiveRecord(v ArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *BenchlingApp) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *BenchlingApp) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *BenchlingApp) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *BenchlingApp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BenchlingApp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BenchlingApp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BenchlingApp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *BenchlingApp) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *BenchlingApp) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *BenchlingApp) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *BenchlingApp) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetId

`func (o *BenchlingApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BenchlingApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BenchlingApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BenchlingApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BenchlingApp) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BenchlingApp) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BenchlingApp) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BenchlingApp) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetWebUrl

`func (o *BenchlingApp) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *BenchlingApp) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *BenchlingApp) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *BenchlingApp) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


