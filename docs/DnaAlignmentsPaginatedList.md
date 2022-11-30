# DnaAlignmentsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**DnaAlignments** | Pointer to [**[]DnaAlignmentSummary**](DnaAlignmentSummary.md) |  | [optional] 

## Methods

### NewDnaAlignmentsPaginatedList

`func NewDnaAlignmentsPaginatedList() *DnaAlignmentsPaginatedList`

NewDnaAlignmentsPaginatedList instantiates a new DnaAlignmentsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaAlignmentsPaginatedListWithDefaults

`func NewDnaAlignmentsPaginatedListWithDefaults() *DnaAlignmentsPaginatedList`

NewDnaAlignmentsPaginatedListWithDefaults instantiates a new DnaAlignmentsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *DnaAlignmentsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *DnaAlignmentsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *DnaAlignmentsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *DnaAlignmentsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetDnaAlignments

`func (o *DnaAlignmentsPaginatedList) GetDnaAlignments() []DnaAlignmentSummary`

GetDnaAlignments returns the DnaAlignments field if non-nil, zero value otherwise.

### GetDnaAlignmentsOk

`func (o *DnaAlignmentsPaginatedList) GetDnaAlignmentsOk() (*[]DnaAlignmentSummary, bool)`

GetDnaAlignmentsOk returns a tuple with the DnaAlignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaAlignments

`func (o *DnaAlignmentsPaginatedList) SetDnaAlignments(v []DnaAlignmentSummary)`

SetDnaAlignments sets DnaAlignments field to given value.

### HasDnaAlignments

`func (o *DnaAlignmentsPaginatedList) HasDnaAlignments() bool`

HasDnaAlignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


