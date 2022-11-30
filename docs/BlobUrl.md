# BlobUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadURL** | Pointer to **string** | a pre-signed download url. | [optional] 
**ExpiresAt** | Pointer to **int32** | The unix time that the download URL expires at. | [optional] 

## Methods

### NewBlobUrl

`func NewBlobUrl() *BlobUrl`

NewBlobUrl instantiates a new BlobUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobUrlWithDefaults

`func NewBlobUrlWithDefaults() *BlobUrl`

NewBlobUrlWithDefaults instantiates a new BlobUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadURL

`func (o *BlobUrl) GetDownloadURL() string`

GetDownloadURL returns the DownloadURL field if non-nil, zero value otherwise.

### GetDownloadURLOk

`func (o *BlobUrl) GetDownloadURLOk() (*string, bool)`

GetDownloadURLOk returns a tuple with the DownloadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadURL

`func (o *BlobUrl) SetDownloadURL(v string)`

SetDownloadURL sets DownloadURL field to given value.

### HasDownloadURL

`func (o *BlobUrl) HasDownloadURL() bool`

HasDownloadURL returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BlobUrl) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BlobUrl) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BlobUrl) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BlobUrl) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


