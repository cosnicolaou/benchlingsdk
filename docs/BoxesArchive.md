# BoxesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxIds** | **[]string** | Array of box IDs | 
**Reason** | **string** | Reason that boxes are being archived.  | 
**ShouldRemoveBarcodes** | Pointer to **bool** | Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items.  | [optional] 

## Methods

### NewBoxesArchive

`func NewBoxesArchive(boxIds []string, reason string, ) *BoxesArchive`

NewBoxesArchive instantiates a new BoxesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxesArchiveWithDefaults

`func NewBoxesArchiveWithDefaults() *BoxesArchive`

NewBoxesArchiveWithDefaults instantiates a new BoxesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxIds

`func (o *BoxesArchive) GetBoxIds() []string`

GetBoxIds returns the BoxIds field if non-nil, zero value otherwise.

### GetBoxIdsOk

`func (o *BoxesArchive) GetBoxIdsOk() (*[]string, bool)`

GetBoxIdsOk returns a tuple with the BoxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxIds

`func (o *BoxesArchive) SetBoxIds(v []string)`

SetBoxIds sets BoxIds field to given value.


### GetReason

`func (o *BoxesArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BoxesArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BoxesArchive) SetReason(v string)`

SetReason sets Reason field to given value.


### GetShouldRemoveBarcodes

`func (o *BoxesArchive) GetShouldRemoveBarcodes() bool`

GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field if non-nil, zero value otherwise.

### GetShouldRemoveBarcodesOk

`func (o *BoxesArchive) GetShouldRemoveBarcodesOk() (*bool, bool)`

GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRemoveBarcodes

`func (o *BoxesArchive) SetShouldRemoveBarcodes(v bool)`

SetShouldRemoveBarcodes sets ShouldRemoveBarcodes field to given value.

### HasShouldRemoveBarcodes

`func (o *BoxesArchive) HasShouldRemoveBarcodes() bool`

HasShouldRemoveBarcodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


