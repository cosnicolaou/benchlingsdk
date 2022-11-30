# LocationSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationSchemas** | Pointer to [**[]LocationSchema**](LocationSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationSchemasPaginatedList

`func NewLocationSchemasPaginatedList() *LocationSchemasPaginatedList`

NewLocationSchemasPaginatedList instantiates a new LocationSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationSchemasPaginatedListWithDefaults

`func NewLocationSchemasPaginatedListWithDefaults() *LocationSchemasPaginatedList`

NewLocationSchemasPaginatedListWithDefaults instantiates a new LocationSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationSchemas

`func (o *LocationSchemasPaginatedList) GetLocationSchemas() []LocationSchema`

GetLocationSchemas returns the LocationSchemas field if non-nil, zero value otherwise.

### GetLocationSchemasOk

`func (o *LocationSchemasPaginatedList) GetLocationSchemasOk() (*[]LocationSchema, bool)`

GetLocationSchemasOk returns a tuple with the LocationSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationSchemas

`func (o *LocationSchemasPaginatedList) SetLocationSchemas(v []LocationSchema)`

SetLocationSchemas sets LocationSchemas field to given value.

### HasLocationSchemas

`func (o *LocationSchemasPaginatedList) HasLocationSchemas() bool`

HasLocationSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *LocationSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *LocationSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *LocationSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *LocationSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


