# BatchesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchIds** | **[]string** |  | 
**Reason** | **string** | The reason for archiving the provided Batches. Accepted reasons may differ based on tenant configuration.  | 

## Methods

### NewBatchesArchive

`func NewBatchesArchive(batchIds []string, reason string, ) *BatchesArchive`

NewBatchesArchive instantiates a new BatchesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchesArchiveWithDefaults

`func NewBatchesArchiveWithDefaults() *BatchesArchive`

NewBatchesArchiveWithDefaults instantiates a new BatchesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchIds

`func (o *BatchesArchive) GetBatchIds() []string`

GetBatchIds returns the BatchIds field if non-nil, zero value otherwise.

### GetBatchIdsOk

`func (o *BatchesArchive) GetBatchIdsOk() (*[]string, bool)`

GetBatchIdsOk returns a tuple with the BatchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIds

`func (o *BatchesArchive) SetBatchIds(v []string)`

SetBatchIds sets BatchIds field to given value.


### GetReason

`func (o *BatchesArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BatchesArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BatchesArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


