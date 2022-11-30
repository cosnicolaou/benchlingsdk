# BatchesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batches** | Pointer to [**[]Batch**](Batch.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchesPaginatedList

`func NewBatchesPaginatedList() *BatchesPaginatedList`

NewBatchesPaginatedList instantiates a new BatchesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchesPaginatedListWithDefaults

`func NewBatchesPaginatedListWithDefaults() *BatchesPaginatedList`

NewBatchesPaginatedListWithDefaults instantiates a new BatchesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatches

`func (o *BatchesPaginatedList) GetBatches() []Batch`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *BatchesPaginatedList) GetBatchesOk() (*[]Batch, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *BatchesPaginatedList) SetBatches(v []Batch)`

SetBatches sets Batches field to given value.

### HasBatches

`func (o *BatchesPaginatedList) HasBatches() bool`

HasBatches returns a boolean if a field has been set.

### GetNextToken

`func (o *BatchesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BatchesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BatchesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BatchesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


