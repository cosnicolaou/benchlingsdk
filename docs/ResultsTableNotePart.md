# ResultsTableNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 
**AssayResultSchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewResultsTableNotePart

`func NewResultsTableNotePart() *ResultsTableNotePart`

NewResultsTableNotePart instantiates a new ResultsTableNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsTableNotePartWithDefaults

`func NewResultsTableNotePartWithDefaults() *ResultsTableNotePart`

NewResultsTableNotePartWithDefaults instantiates a new ResultsTableNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *ResultsTableNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *ResultsTableNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *ResultsTableNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *ResultsTableNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *ResultsTableNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResultsTableNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResultsTableNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResultsTableNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApiId

`func (o *ResultsTableNotePart) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *ResultsTableNotePart) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *ResultsTableNotePart) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *ResultsTableNotePart) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *ResultsTableNotePart) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ResultsTableNotePart) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ResultsTableNotePart) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ResultsTableNotePart) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetAssayResultSchemaId

`func (o *ResultsTableNotePart) GetAssayResultSchemaId() string`

GetAssayResultSchemaId returns the AssayResultSchemaId field if non-nil, zero value otherwise.

### GetAssayResultSchemaIdOk

`func (o *ResultsTableNotePart) GetAssayResultSchemaIdOk() (*string, bool)`

GetAssayResultSchemaIdOk returns a tuple with the AssayResultSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayResultSchemaId

`func (o *ResultsTableNotePart) SetAssayResultSchemaId(v string)`

SetAssayResultSchemaId sets AssayResultSchemaId field to given value.

### HasAssayResultSchemaId

`func (o *ResultsTableNotePart) HasAssayResultSchemaId() bool`

HasAssayResultSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


