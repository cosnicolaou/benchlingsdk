# RequestFulfillmentsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**RequestFulfillments** | Pointer to [**[]RequestFulfillment**](RequestFulfillment.md) |  | [optional] 

## Methods

### NewRequestFulfillmentsPaginatedList

`func NewRequestFulfillmentsPaginatedList() *RequestFulfillmentsPaginatedList`

NewRequestFulfillmentsPaginatedList instantiates a new RequestFulfillmentsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFulfillmentsPaginatedListWithDefaults

`func NewRequestFulfillmentsPaginatedListWithDefaults() *RequestFulfillmentsPaginatedList`

NewRequestFulfillmentsPaginatedListWithDefaults instantiates a new RequestFulfillmentsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *RequestFulfillmentsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RequestFulfillmentsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RequestFulfillmentsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RequestFulfillmentsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRequestFulfillments

`func (o *RequestFulfillmentsPaginatedList) GetRequestFulfillments() []RequestFulfillment`

GetRequestFulfillments returns the RequestFulfillments field if non-nil, zero value otherwise.

### GetRequestFulfillmentsOk

`func (o *RequestFulfillmentsPaginatedList) GetRequestFulfillmentsOk() (*[]RequestFulfillment, bool)`

GetRequestFulfillmentsOk returns a tuple with the RequestFulfillments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFulfillments

`func (o *RequestFulfillmentsPaginatedList) SetRequestFulfillments(v []RequestFulfillment)`

SetRequestFulfillments sets RequestFulfillments field to given value.

### HasRequestFulfillments

`func (o *RequestFulfillmentsPaginatedList) HasRequestFulfillments() bool`

HasRequestFulfillments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


