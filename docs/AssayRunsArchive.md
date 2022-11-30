# AssayRunsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRunIds** | **[]string** |  | 
**Reason** | **string** | The reason for archiving the provided Assay Runs. Accepted reasons may differ based on tenant configuration.  | 

## Methods

### NewAssayRunsArchive

`func NewAssayRunsArchive(assayRunIds []string, reason string, ) *AssayRunsArchive`

NewAssayRunsArchive instantiates a new AssayRunsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunsArchiveWithDefaults

`func NewAssayRunsArchiveWithDefaults() *AssayRunsArchive`

NewAssayRunsArchiveWithDefaults instantiates a new AssayRunsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRunIds

`func (o *AssayRunsArchive) GetAssayRunIds() []string`

GetAssayRunIds returns the AssayRunIds field if non-nil, zero value otherwise.

### GetAssayRunIdsOk

`func (o *AssayRunsArchive) GetAssayRunIdsOk() (*[]string, bool)`

GetAssayRunIdsOk returns a tuple with the AssayRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunIds

`func (o *AssayRunsArchive) SetAssayRunIds(v []string)`

SetAssayRunIds sets AssayRunIds field to given value.


### GetReason

`func (o *AssayRunsArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AssayRunsArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AssayRunsArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


