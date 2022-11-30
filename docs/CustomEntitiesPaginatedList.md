# CustomEntitiesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomEntities** | Pointer to [**[]CustomEntity**](CustomEntity.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomEntitiesPaginatedList

`func NewCustomEntitiesPaginatedList() *CustomEntitiesPaginatedList`

NewCustomEntitiesPaginatedList instantiates a new CustomEntitiesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntitiesPaginatedListWithDefaults

`func NewCustomEntitiesPaginatedListWithDefaults() *CustomEntitiesPaginatedList`

NewCustomEntitiesPaginatedListWithDefaults instantiates a new CustomEntitiesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomEntities

`func (o *CustomEntitiesPaginatedList) GetCustomEntities() []CustomEntity`

GetCustomEntities returns the CustomEntities field if non-nil, zero value otherwise.

### GetCustomEntitiesOk

`func (o *CustomEntitiesPaginatedList) GetCustomEntitiesOk() (*[]CustomEntity, bool)`

GetCustomEntitiesOk returns a tuple with the CustomEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEntities

`func (o *CustomEntitiesPaginatedList) SetCustomEntities(v []CustomEntity)`

SetCustomEntities sets CustomEntities field to given value.

### HasCustomEntities

`func (o *CustomEntitiesPaginatedList) HasCustomEntities() bool`

HasCustomEntities returns a boolean if a field has been set.

### GetNextToken

`func (o *CustomEntitiesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *CustomEntitiesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *CustomEntitiesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *CustomEntitiesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


