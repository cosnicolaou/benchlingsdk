# EventsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsPaginatedList

`func NewEventsPaginatedList() *EventsPaginatedList`

NewEventsPaginatedList instantiates a new EventsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsPaginatedListWithDefaults

`func NewEventsPaginatedListWithDefaults() *EventsPaginatedList`

NewEventsPaginatedListWithDefaults instantiates a new EventsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsPaginatedList) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsPaginatedList) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsPaginatedList) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EventsPaginatedList) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetNextToken

`func (o *EventsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *EventsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *EventsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *EventsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


