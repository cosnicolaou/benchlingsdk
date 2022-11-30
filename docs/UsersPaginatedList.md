# UsersPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewUsersPaginatedList

`func NewUsersPaginatedList() *UsersPaginatedList`

NewUsersPaginatedList instantiates a new UsersPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPaginatedListWithDefaults

`func NewUsersPaginatedListWithDefaults() *UsersPaginatedList`

NewUsersPaginatedListWithDefaults instantiates a new UsersPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *UsersPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *UsersPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *UsersPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *UsersPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetUsers

`func (o *UsersPaginatedList) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersPaginatedList) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersPaginatedList) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersPaginatedList) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


