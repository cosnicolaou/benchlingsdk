# DnaSequencesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnaSequences** | Pointer to [**[]DnaSequence**](DnaSequence.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewDnaSequencesPaginatedList

`func NewDnaSequencesPaginatedList() *DnaSequencesPaginatedList`

NewDnaSequencesPaginatedList instantiates a new DnaSequencesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaSequencesPaginatedListWithDefaults

`func NewDnaSequencesPaginatedListWithDefaults() *DnaSequencesPaginatedList`

NewDnaSequencesPaginatedListWithDefaults instantiates a new DnaSequencesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnaSequences

`func (o *DnaSequencesPaginatedList) GetDnaSequences() []DnaSequence`

GetDnaSequences returns the DnaSequences field if non-nil, zero value otherwise.

### GetDnaSequencesOk

`func (o *DnaSequencesPaginatedList) GetDnaSequencesOk() (*[]DnaSequence, bool)`

GetDnaSequencesOk returns a tuple with the DnaSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaSequences

`func (o *DnaSequencesPaginatedList) SetDnaSequences(v []DnaSequence)`

SetDnaSequences sets DnaSequences field to given value.

### HasDnaSequences

`func (o *DnaSequencesPaginatedList) HasDnaSequences() bool`

HasDnaSequences returns a boolean if a field has been set.

### GetNextToken

`func (o *DnaSequencesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *DnaSequencesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *DnaSequencesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *DnaSequencesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


