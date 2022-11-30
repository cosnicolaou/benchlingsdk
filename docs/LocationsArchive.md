# LocationsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationIds** | **[]string** | Array of location IDs | 
**Reason** | **string** | Reason that locations are being archived.  | 
**ShouldRemoveBarcodes** | Pointer to **bool** | Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items.  | [optional] 

## Methods

### NewLocationsArchive

`func NewLocationsArchive(locationIds []string, reason string, ) *LocationsArchive`

NewLocationsArchive instantiates a new LocationsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationsArchiveWithDefaults

`func NewLocationsArchiveWithDefaults() *LocationsArchive`

NewLocationsArchiveWithDefaults instantiates a new LocationsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationIds

`func (o *LocationsArchive) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *LocationsArchive) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *LocationsArchive) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.


### GetReason

`func (o *LocationsArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LocationsArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LocationsArchive) SetReason(v string)`

SetReason sets Reason field to given value.


### GetShouldRemoveBarcodes

`func (o *LocationsArchive) GetShouldRemoveBarcodes() bool`

GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field if non-nil, zero value otherwise.

### GetShouldRemoveBarcodesOk

`func (o *LocationsArchive) GetShouldRemoveBarcodesOk() (*bool, bool)`

GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRemoveBarcodes

`func (o *LocationsArchive) SetShouldRemoveBarcodes(v bool)`

SetShouldRemoveBarcodes sets ShouldRemoveBarcodes field to given value.

### HasShouldRemoveBarcodes

`func (o *LocationsArchive) HasShouldRemoveBarcodes() bool`

HasShouldRemoveBarcodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


