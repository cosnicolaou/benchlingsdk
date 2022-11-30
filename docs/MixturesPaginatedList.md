# MixturesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mixtures** | Pointer to [**[]Mixture**](Mixture.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewMixturesPaginatedList

`func NewMixturesPaginatedList() *MixturesPaginatedList`

NewMixturesPaginatedList instantiates a new MixturesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixturesPaginatedListWithDefaults

`func NewMixturesPaginatedListWithDefaults() *MixturesPaginatedList`

NewMixturesPaginatedListWithDefaults instantiates a new MixturesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMixtures

`func (o *MixturesPaginatedList) GetMixtures() []Mixture`

GetMixtures returns the Mixtures field if non-nil, zero value otherwise.

### GetMixturesOk

`func (o *MixturesPaginatedList) GetMixturesOk() (*[]Mixture, bool)`

GetMixturesOk returns a tuple with the Mixtures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixtures

`func (o *MixturesPaginatedList) SetMixtures(v []Mixture)`

SetMixtures sets Mixtures field to given value.

### HasMixtures

`func (o *MixturesPaginatedList) HasMixtures() bool`

HasMixtures returns a boolean if a field has been set.

### GetNextToken

`func (o *MixturesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MixturesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MixturesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *MixturesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


