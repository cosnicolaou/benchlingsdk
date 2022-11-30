# OrganizationsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Organizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 

## Methods

### NewOrganizationsPaginatedList

`func NewOrganizationsPaginatedList() *OrganizationsPaginatedList`

NewOrganizationsPaginatedList instantiates a new OrganizationsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsPaginatedListWithDefaults

`func NewOrganizationsPaginatedListWithDefaults() *OrganizationsPaginatedList`

NewOrganizationsPaginatedListWithDefaults instantiates a new OrganizationsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *OrganizationsPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *OrganizationsPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *OrganizationsPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *OrganizationsPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetOrganizations

`func (o *OrganizationsPaginatedList) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *OrganizationsPaginatedList) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *OrganizationsPaginatedList) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *OrganizationsPaginatedList) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


