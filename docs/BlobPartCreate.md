# BlobPartCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data64** | **string** |  | 
**Md5** | **string** |  | 
**PartNumber** | **int32** | An integer between 1 to 10,000, inclusive. The part number must be unique per part and indicates the ordering of the part inside the final blob. The part numbers do not need to be consecutive.  | 

## Methods

### NewBlobPartCreate

`func NewBlobPartCreate(data64 string, md5 string, partNumber int32, ) *BlobPartCreate`

NewBlobPartCreate instantiates a new BlobPartCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobPartCreateWithDefaults

`func NewBlobPartCreateWithDefaults() *BlobPartCreate`

NewBlobPartCreateWithDefaults instantiates a new BlobPartCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData64

`func (o *BlobPartCreate) GetData64() string`

GetData64 returns the Data64 field if non-nil, zero value otherwise.

### GetData64Ok

`func (o *BlobPartCreate) GetData64Ok() (*string, bool)`

GetData64Ok returns a tuple with the Data64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData64

`func (o *BlobPartCreate) SetData64(v string)`

SetData64 sets Data64 field to given value.


### GetMd5

`func (o *BlobPartCreate) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *BlobPartCreate) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *BlobPartCreate) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetPartNumber

`func (o *BlobPartCreate) GetPartNumber() int32`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BlobPartCreate) GetPartNumberOk() (*int32, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BlobPartCreate) SetPartNumber(v int32)`

SetPartNumber sets PartNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


