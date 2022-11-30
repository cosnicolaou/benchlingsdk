# RequestsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]Request**](Request.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestsPaginatedList

`func NewRequestsPaginatedList() *RequestsPaginatedList`

NewRequestsPaginatedList instantiates a new RequestsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsPaginatedListWithDefaults

`func NewRequestsPaginatedListWithDefaults() *RequestsPaginatedList`

NewRequestsPaginatedListWithDefaults instantiates a new RequestsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *RequestsPaginatedList) GetRequests() []Request`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *RequestsPaginatedList) GetRequestsOk() (*[]Request, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *RequestsPaginatedList) SetRequests(v []Request)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *RequestsPaginatedList) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetNextToken

`func (o *RequestsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RequestsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RequestsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RequestsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


