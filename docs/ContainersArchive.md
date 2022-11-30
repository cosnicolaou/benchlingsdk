# ContainersArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerIds** | **[]string** | Array of container IDs | 
**Reason** | **string** | Reason that containers are being archived.  | 
**ShouldRemoveBarcodes** | Pointer to **bool** | Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items.  | [optional] 

## Methods

### NewContainersArchive

`func NewContainersArchive(containerIds []string, reason string, ) *ContainersArchive`

NewContainersArchive instantiates a new ContainersArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersArchiveWithDefaults

`func NewContainersArchiveWithDefaults() *ContainersArchive`

NewContainersArchiveWithDefaults instantiates a new ContainersArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerIds

`func (o *ContainersArchive) GetContainerIds() []string`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *ContainersArchive) GetContainerIdsOk() (*[]string, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *ContainersArchive) SetContainerIds(v []string)`

SetContainerIds sets ContainerIds field to given value.


### GetReason

`func (o *ContainersArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ContainersArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ContainersArchive) SetReason(v string)`

SetReason sets Reason field to given value.


### GetShouldRemoveBarcodes

`func (o *ContainersArchive) GetShouldRemoveBarcodes() bool`

GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field if non-nil, zero value otherwise.

### GetShouldRemoveBarcodesOk

`func (o *ContainersArchive) GetShouldRemoveBarcodesOk() (*bool, bool)`

GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRemoveBarcodes

`func (o *ContainersArchive) SetShouldRemoveBarcodes(v bool)`

SetShouldRemoveBarcodes sets ShouldRemoveBarcodes field to given value.

### HasShouldRemoveBarcodes

`func (o *ContainersArchive) HasShouldRemoveBarcodes() bool`

HasShouldRemoveBarcodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


