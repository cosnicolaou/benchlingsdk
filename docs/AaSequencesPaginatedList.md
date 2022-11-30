# AaSequencesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AaSequences** | Pointer to [**[]AaSequence**](AaSequence.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAaSequencesPaginatedList

`func NewAaSequencesPaginatedList() *AaSequencesPaginatedList`

NewAaSequencesPaginatedList instantiates a new AaSequencesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaSequencesPaginatedListWithDefaults

`func NewAaSequencesPaginatedListWithDefaults() *AaSequencesPaginatedList`

NewAaSequencesPaginatedListWithDefaults instantiates a new AaSequencesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaSequences

`func (o *AaSequencesPaginatedList) GetAaSequences() []AaSequence`

GetAaSequences returns the AaSequences field if non-nil, zero value otherwise.

### GetAaSequencesOk

`func (o *AaSequencesPaginatedList) GetAaSequencesOk() (*[]AaSequence, bool)`

GetAaSequencesOk returns a tuple with the AaSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaSequences

`func (o *AaSequencesPaginatedList) SetAaSequences(v []AaSequence)`

SetAaSequences sets AaSequences field to given value.

### HasAaSequences

`func (o *AaSequencesPaginatedList) HasAaSequences() bool`

HasAaSequences returns a boolean if a field has been set.

### GetNextToken

`func (o *AaSequencesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *AaSequencesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *AaSequencesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *AaSequencesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


