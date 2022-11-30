# AssayResultsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayResultIds** | **[]string** |  | 
**Reason** | Pointer to **string** | The reason for archiving the provided results. Accepted reasons may differ based on tenant configuration | [optional] 

## Methods

### NewAssayResultsArchive

`func NewAssayResultsArchive(assayResultIds []string, ) *AssayResultsArchive`

NewAssayResultsArchive instantiates a new AssayResultsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayResultsArchiveWithDefaults

`func NewAssayResultsArchiveWithDefaults() *AssayResultsArchive`

NewAssayResultsArchiveWithDefaults instantiates a new AssayResultsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayResultIds

`func (o *AssayResultsArchive) GetAssayResultIds() []string`

GetAssayResultIds returns the AssayResultIds field if non-nil, zero value otherwise.

### GetAssayResultIdsOk

`func (o *AssayResultsArchive) GetAssayResultIdsOk() (*[]string, bool)`

GetAssayResultIdsOk returns a tuple with the AssayResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayResultIds

`func (o *AssayResultsArchive) SetAssayResultIds(v []string)`

SetAssayResultIds sets AssayResultIds field to given value.


### GetReason

`func (o *AssayResultsArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AssayResultsArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AssayResultsArchive) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AssayResultsArchive) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


