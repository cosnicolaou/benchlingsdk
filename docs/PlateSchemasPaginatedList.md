# PlateSchemasPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlateSchemas** | Pointer to [**[]PlateSchema**](PlateSchema.md) |  | [optional] [readonly] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewPlateSchemasPaginatedList

`func NewPlateSchemasPaginatedList() *PlateSchemasPaginatedList`

NewPlateSchemasPaginatedList instantiates a new PlateSchemasPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlateSchemasPaginatedListWithDefaults

`func NewPlateSchemasPaginatedListWithDefaults() *PlateSchemasPaginatedList`

NewPlateSchemasPaginatedListWithDefaults instantiates a new PlateSchemasPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlateSchemas

`func (o *PlateSchemasPaginatedList) GetPlateSchemas() []PlateSchema`

GetPlateSchemas returns the PlateSchemas field if non-nil, zero value otherwise.

### GetPlateSchemasOk

`func (o *PlateSchemasPaginatedList) GetPlateSchemasOk() (*[]PlateSchema, bool)`

GetPlateSchemasOk returns a tuple with the PlateSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateSchemas

`func (o *PlateSchemasPaginatedList) SetPlateSchemas(v []PlateSchema)`

SetPlateSchemas sets PlateSchemas field to given value.

### HasPlateSchemas

`func (o *PlateSchemasPaginatedList) HasPlateSchemas() bool`

HasPlateSchemas returns a boolean if a field has been set.

### GetNextToken

`func (o *PlateSchemasPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *PlateSchemasPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *PlateSchemasPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *PlateSchemasPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


