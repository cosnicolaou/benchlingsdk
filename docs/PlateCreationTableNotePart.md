# PlateCreationTableNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 
**PlateSchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewPlateCreationTableNotePart

`func NewPlateCreationTableNotePart() *PlateCreationTableNotePart`

NewPlateCreationTableNotePart instantiates a new PlateCreationTableNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlateCreationTableNotePartWithDefaults

`func NewPlateCreationTableNotePartWithDefaults() *PlateCreationTableNotePart`

NewPlateCreationTableNotePartWithDefaults instantiates a new PlateCreationTableNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *PlateCreationTableNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *PlateCreationTableNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *PlateCreationTableNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *PlateCreationTableNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *PlateCreationTableNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlateCreationTableNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlateCreationTableNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlateCreationTableNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApiId

`func (o *PlateCreationTableNotePart) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *PlateCreationTableNotePart) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *PlateCreationTableNotePart) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *PlateCreationTableNotePart) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *PlateCreationTableNotePart) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *PlateCreationTableNotePart) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *PlateCreationTableNotePart) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *PlateCreationTableNotePart) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetPlateSchemaId

`func (o *PlateCreationTableNotePart) GetPlateSchemaId() string`

GetPlateSchemaId returns the PlateSchemaId field if non-nil, zero value otherwise.

### GetPlateSchemaIdOk

`func (o *PlateCreationTableNotePart) GetPlateSchemaIdOk() (*string, bool)`

GetPlateSchemaIdOk returns a tuple with the PlateSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateSchemaId

`func (o *PlateCreationTableNotePart) SetPlateSchemaId(v string)`

SetPlateSchemaId sets PlateSchemaId field to given value.

### HasPlateSchemaId

`func (o *PlateCreationTableNotePart) HasPlateSchemaId() bool`

HasPlateSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


