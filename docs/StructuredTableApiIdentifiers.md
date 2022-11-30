# StructuredTableApiIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 

## Methods

### NewStructuredTableApiIdentifiers

`func NewStructuredTableApiIdentifiers() *StructuredTableApiIdentifiers`

NewStructuredTableApiIdentifiers instantiates a new StructuredTableApiIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredTableApiIdentifiersWithDefaults

`func NewStructuredTableApiIdentifiersWithDefaults() *StructuredTableApiIdentifiers`

NewStructuredTableApiIdentifiersWithDefaults instantiates a new StructuredTableApiIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiId

`func (o *StructuredTableApiIdentifiers) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *StructuredTableApiIdentifiers) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *StructuredTableApiIdentifiers) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *StructuredTableApiIdentifiers) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *StructuredTableApiIdentifiers) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *StructuredTableApiIdentifiers) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *StructuredTableApiIdentifiers) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *StructuredTableApiIdentifiers) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


