# RnaSequencesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**RnaSequences** | Pointer to [**[]RnaSequence**](RnaSequence.md) |  | [optional] 

## Methods

### NewRnaSequencesPaginatedList

`func NewRnaSequencesPaginatedList() *RnaSequencesPaginatedList`

NewRnaSequencesPaginatedList instantiates a new RnaSequencesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaSequencesPaginatedListWithDefaults

`func NewRnaSequencesPaginatedListWithDefaults() *RnaSequencesPaginatedList`

NewRnaSequencesPaginatedListWithDefaults instantiates a new RnaSequencesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *RnaSequencesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RnaSequencesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RnaSequencesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RnaSequencesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRnaSequences

`func (o *RnaSequencesPaginatedList) GetRnaSequences() []RnaSequence`

GetRnaSequences returns the RnaSequences field if non-nil, zero value otherwise.

### GetRnaSequencesOk

`func (o *RnaSequencesPaginatedList) GetRnaSequencesOk() (*[]RnaSequence, bool)`

GetRnaSequencesOk returns a tuple with the RnaSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaSequences

`func (o *RnaSequencesPaginatedList) SetRnaSequences(v []RnaSequence)`

SetRnaSequences sets RnaSequences field to given value.

### HasRnaSequences

`func (o *RnaSequencesPaginatedList) HasRnaSequences() bool`

HasRnaSequences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


