# BoxCreationTableNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 
**BoxSchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewBoxCreationTableNotePart

`func NewBoxCreationTableNotePart() *BoxCreationTableNotePart`

NewBoxCreationTableNotePart instantiates a new BoxCreationTableNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxCreationTableNotePartWithDefaults

`func NewBoxCreationTableNotePartWithDefaults() *BoxCreationTableNotePart`

NewBoxCreationTableNotePartWithDefaults instantiates a new BoxCreationTableNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *BoxCreationTableNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *BoxCreationTableNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *BoxCreationTableNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *BoxCreationTableNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *BoxCreationTableNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoxCreationTableNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoxCreationTableNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BoxCreationTableNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApiId

`func (o *BoxCreationTableNotePart) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *BoxCreationTableNotePart) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *BoxCreationTableNotePart) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *BoxCreationTableNotePart) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *BoxCreationTableNotePart) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *BoxCreationTableNotePart) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *BoxCreationTableNotePart) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *BoxCreationTableNotePart) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetBoxSchemaId

`func (o *BoxCreationTableNotePart) GetBoxSchemaId() string`

GetBoxSchemaId returns the BoxSchemaId field if non-nil, zero value otherwise.

### GetBoxSchemaIdOk

`func (o *BoxCreationTableNotePart) GetBoxSchemaIdOk() (*string, bool)`

GetBoxSchemaIdOk returns a tuple with the BoxSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxSchemaId

`func (o *BoxCreationTableNotePart) SetBoxSchemaId(v string)`

SetBoxSchemaId sets BoxSchemaId field to given value.

### HasBoxSchemaId

`func (o *BoxCreationTableNotePart) HasBoxSchemaId() bool`

HasBoxSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


