# Blob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The universally unique identifier (UUID) for the blob. | [optional] 
**MimeType** | Pointer to **string** | eg. application/jpeg | [optional] 
**Name** | Pointer to **string** | Name of the blob | [optional] 
**Type** | Pointer to **string** | One of RAW_FILE or VISUALIZATION. If VISUALIZATION, the blob may be displayed as an image preview.  | [optional] 
**UploadStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewBlob

`func NewBlob() *Blob`

NewBlob instantiates a new Blob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobWithDefaults

`func NewBlobWithDefaults() *Blob`

NewBlobWithDefaults instantiates a new Blob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Blob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Blob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Blob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Blob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimeType

`func (o *Blob) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Blob) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Blob) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Blob) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *Blob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Blob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Blob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Blob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Blob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Blob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Blob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Blob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUploadStatus

`func (o *Blob) GetUploadStatus() string`

GetUploadStatus returns the UploadStatus field if non-nil, zero value otherwise.

### GetUploadStatusOk

`func (o *Blob) GetUploadStatusOk() (*string, bool)`

GetUploadStatusOk returns a tuple with the UploadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadStatus

`func (o *Blob) SetUploadStatus(v string)`

SetUploadStatus sets UploadStatus field to given value.

### HasUploadStatus

`func (o *Blob) HasUploadStatus() bool`

HasUploadStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


