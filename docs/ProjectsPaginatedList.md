# ProjectsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]Project**](Project.md) |  | [optional] 

## Methods

### NewProjectsPaginatedList

`func NewProjectsPaginatedList() *ProjectsPaginatedList`

NewProjectsPaginatedList instantiates a new ProjectsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsPaginatedListWithDefaults

`func NewProjectsPaginatedListWithDefaults() *ProjectsPaginatedList`

NewProjectsPaginatedListWithDefaults instantiates a new ProjectsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *ProjectsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ProjectsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ProjectsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ProjectsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetProjects

`func (o *ProjectsPaginatedList) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectsPaginatedList) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectsPaginatedList) SetProjects(v []Project)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectsPaginatedList) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


