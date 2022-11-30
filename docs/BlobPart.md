# BlobPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ETag** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewBlobPart

`func NewBlobPart() *BlobPart`

NewBlobPart instantiates a new BlobPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobPartWithDefaults

`func NewBlobPartWithDefaults() *BlobPart`

NewBlobPartWithDefaults instantiates a new BlobPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetETag

`func (o *BlobPart) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *BlobPart) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *BlobPart) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *BlobPart) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetPartNumber

`func (o *BlobPart) GetPartNumber() int32`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BlobPart) GetPartNumberOk() (*int32, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BlobPart) SetPartNumber(v int32)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BlobPart) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


