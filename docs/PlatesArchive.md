# PlatesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlateIds** | **[]string** | Array of plate IDs | 
**Reason** | **string** | Reason that plates are being archived.  | 
**ShouldRemoveBarcodes** | **bool** | Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items.  | 

## Methods

### NewPlatesArchive

`func NewPlatesArchive(plateIds []string, reason string, shouldRemoveBarcodes bool, ) *PlatesArchive`

NewPlatesArchive instantiates a new PlatesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatesArchiveWithDefaults

`func NewPlatesArchiveWithDefaults() *PlatesArchive`

NewPlatesArchiveWithDefaults instantiates a new PlatesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlateIds

`func (o *PlatesArchive) GetPlateIds() []string`

GetPlateIds returns the PlateIds field if non-nil, zero value otherwise.

### GetPlateIdsOk

`func (o *PlatesArchive) GetPlateIdsOk() (*[]string, bool)`

GetPlateIdsOk returns a tuple with the PlateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateIds

`func (o *PlatesArchive) SetPlateIds(v []string)`

SetPlateIds sets PlateIds field to given value.


### GetReason

`func (o *PlatesArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlatesArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlatesArchive) SetReason(v string)`

SetReason sets Reason field to given value.


### GetShouldRemoveBarcodes

`func (o *PlatesArchive) GetShouldRemoveBarcodes() bool`

GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field if non-nil, zero value otherwise.

### GetShouldRemoveBarcodesOk

`func (o *PlatesArchive) GetShouldRemoveBarcodesOk() (*bool, bool)`

GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRemoveBarcodes

`func (o *PlatesArchive) SetShouldRemoveBarcodes(v bool)`

SetShouldRemoveBarcodes sets ShouldRemoveBarcodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


