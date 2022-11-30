# NucleotideAlignmentsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Alignments** | Pointer to [**[]NucleotideAlignmentSummary**](NucleotideAlignmentSummary.md) |  | [optional] 

## Methods

### NewNucleotideAlignmentsPaginatedList

`func NewNucleotideAlignmentsPaginatedList() *NucleotideAlignmentsPaginatedList`

NewNucleotideAlignmentsPaginatedList instantiates a new NucleotideAlignmentsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideAlignmentsPaginatedListWithDefaults

`func NewNucleotideAlignmentsPaginatedListWithDefaults() *NucleotideAlignmentsPaginatedList`

NewNucleotideAlignmentsPaginatedListWithDefaults instantiates a new NucleotideAlignmentsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *NucleotideAlignmentsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *NucleotideAlignmentsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *NucleotideAlignmentsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *NucleotideAlignmentsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetAlignments

`func (o *NucleotideAlignmentsPaginatedList) GetAlignments() []NucleotideAlignmentSummary`

GetAlignments returns the Alignments field if non-nil, zero value otherwise.

### GetAlignmentsOk

`func (o *NucleotideAlignmentsPaginatedList) GetAlignmentsOk() (*[]NucleotideAlignmentSummary, bool)`

GetAlignmentsOk returns a tuple with the Alignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignments

`func (o *NucleotideAlignmentsPaginatedList) SetAlignments(v []NucleotideAlignmentSummary)`

SetAlignments sets Alignments field to given value.

### HasAlignments

`func (o *NucleotideAlignmentsPaginatedList) HasAlignments() bool`

HasAlignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


