# TableNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in the caption via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Table** | Pointer to [**EntryTable**](EntryTable.md) |  | [optional] 
**Text** | Pointer to **string** | The caption of the table. | [optional] 

## Methods

### NewTableNotePart

`func NewTableNotePart() *TableNotePart`

NewTableNotePart instantiates a new TableNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableNotePartWithDefaults

`func NewTableNotePartWithDefaults() *TableNotePart`

NewTableNotePartWithDefaults instantiates a new TableNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *TableNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *TableNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *TableNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *TableNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *TableNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TableNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TableNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TableNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *TableNotePart) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TableNotePart) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TableNotePart) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TableNotePart) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTable

`func (o *TableNotePart) GetTable() EntryTable`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *TableNotePart) GetTableOk() (*EntryTable, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *TableNotePart) SetTable(v EntryTable)`

SetTable sets Table field to given value.

### HasTable

`func (o *TableNotePart) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetText

`func (o *TableNotePart) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TableNotePart) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TableNotePart) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TableNotePart) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


