# EntryExternalFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadURL** | Pointer to **string** | A short-lived URL that can be used to download the original file.  | [optional] 
**ExpiresAt** | Pointer to **int32** | UNIX timestamp when downloadURL expires. | [optional] 
**Id** | Pointer to **string** | ID of the external file | [optional] 
**Size** | Pointer to **int32** | Size, in bytes, of the external file | [optional] 

## Methods

### NewEntryExternalFile

`func NewEntryExternalFile() *EntryExternalFile`

NewEntryExternalFile instantiates a new EntryExternalFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryExternalFileWithDefaults

`func NewEntryExternalFileWithDefaults() *EntryExternalFile`

NewEntryExternalFileWithDefaults instantiates a new EntryExternalFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadURL

`func (o *EntryExternalFile) GetDownloadURL() string`

GetDownloadURL returns the DownloadURL field if non-nil, zero value otherwise.

### GetDownloadURLOk

`func (o *EntryExternalFile) GetDownloadURLOk() (*string, bool)`

GetDownloadURLOk returns a tuple with the DownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadURL

`func (o *EntryExternalFile) SetDownloadURL(v string)`

SetDownloadURL sets DownloadURL field to given value.

### HasDownloadURL

`func (o *EntryExternalFile) HasDownloadURL() bool`

HasDownloadURL returns a boolean if a field has been set.

### GetExpiresAt

`func (o *EntryExternalFile) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EntryExternalFile) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EntryExternalFile) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *EntryExternalFile) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *EntryExternalFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryExternalFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryExternalFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntryExternalFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *EntryExternalFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EntryExternalFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EntryExternalFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *EntryExternalFile) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


