# BlobCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data64** | **string** | base64 encoded file contents | 
**Md5** | **string** | The MD5 hash of the blob part. Note: this should be the hash of the raw data of the blob part, not the hash of the base64 encoding.  | 
**MimeType** | Pointer to **string** | eg. application/jpeg | [optional] [default to "application/octet-stream"]
**Name** | **string** | Name of the blob | 
**Type** | **string** | One of RAW_FILE or VISUALIZATION. If VISUALIZATION, the blob may be displayed as an image preview.  | 

## Methods

### NewBlobCreate

`func NewBlobCreate(data64 string, md5 string, name string, type_ string, ) *BlobCreate`

NewBlobCreate instantiates a new BlobCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobCreateWithDefaults

`func NewBlobCreateWithDefaults() *BlobCreate`

NewBlobCreateWithDefaults instantiates a new BlobCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData64

`func (o *BlobCreate) GetData64() string`

GetData64 returns the Data64 field if non-nil, zero value otherwise.

### GetData64Ok

`func (o *BlobCreate) GetData64Ok() (*string, bool)`

GetData64Ok returns a tuple with the Data64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData64

`func (o *BlobCreate) SetData64(v string)`

SetData64 sets Data64 field to given value.


### GetMd5

`func (o *BlobCreate) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *BlobCreate) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *BlobCreate) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetMimeType

`func (o *BlobCreate) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BlobCreate) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BlobCreate) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BlobCreate) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *BlobCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlobCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlobCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlobCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlobCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlobCreate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


