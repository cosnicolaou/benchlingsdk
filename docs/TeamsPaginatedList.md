# TeamsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**[]Team**](Team.md) |  | [optional] 

## Methods

### NewTeamsPaginatedList

`func NewTeamsPaginatedList() *TeamsPaginatedList`

NewTeamsPaginatedList instantiates a new TeamsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsPaginatedListWithDefaults

`func NewTeamsPaginatedListWithDefaults() *TeamsPaginatedList`

NewTeamsPaginatedListWithDefaults instantiates a new TeamsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *TeamsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *TeamsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *TeamsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *TeamsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTeams

`func (o *TeamsPaginatedList) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TeamsPaginatedList) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TeamsPaginatedList) SetTeams(v []Team)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *TeamsPaginatedList) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


