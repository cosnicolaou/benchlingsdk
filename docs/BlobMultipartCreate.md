# BlobMultipartCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MimeType** | Pointer to **string** | eg. application/jpeg | [optional] [default to "application/octet-stream"]
**Name** | **string** | Name of the blob | 
**Type** | **string** | One of RAW_FILE or VISUALIZATION. If VISUALIZATION, the blob may be displayed as an image preview.  | 

## Methods

### NewBlobMultipartCreate

`func NewBlobMultipartCreate(name string, type_ string, ) *BlobMultipartCreate`

NewBlobMultipartCreate instantiates a new BlobMultipartCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobMultipartCreateWithDefaults

`func NewBlobMultipartCreateWithDefaults() *BlobMultipartCreate`

NewBlobMultipartCreateWithDefaults instantiates a new BlobMultipartCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMimeType

`func (o *BlobMultipartCreate) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BlobMultipartCreate) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BlobMultipartCreate) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BlobMultipartCreate) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *BlobMultipartCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlobMultipartCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlobMultipartCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlobMultipartCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlobMultipartCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlobMultipartCreate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


