# EntryDayNotesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in the caption via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Text** | Pointer to **string** | The caption of the file attachment. | [optional] 
**Table** | Pointer to [**EntryTable**](EntryTable.md) |  | [optional] 
**Checked** | Pointer to **bool** | Indicates whether the checkbox is checked or not.  | [optional] 
**ExternalFileId** | Pointer to **string** | The ID of the external file. Use the &#39;Get an external file&#39; endpoint to retrieve metadata about it.  | [optional] 
**AssayRunId** | Pointer to **NullableString** |  | [optional] 
**AssayRunSchemaId** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 
**AssayResultSchemaId** | Pointer to **string** |  | [optional] 
**EntitySchemaId** | Pointer to **string** |  | [optional] 
**PlateSchemaId** | Pointer to **string** |  | [optional] 
**BoxSchemaId** | Pointer to **string** |  | [optional] 
**MixtureSchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewEntryDayNotesInner

`func NewEntryDayNotesInner() *EntryDayNotesInner`

NewEntryDayNotesInner instantiates a new EntryDayNotesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryDayNotesInnerWithDefaults

`func NewEntryDayNotesInnerWithDefaults() *EntryDayNotesInner`

NewEntryDayNotesInnerWithDefaults instantiates a new EntryDayNotesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *EntryDayNotesInner) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *EntryDayNotesInner) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *EntryDayNotesInner) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *EntryDayNotesInner) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *EntryDayNotesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntryDayNotesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntryDayNotesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntryDayNotesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *EntryDayNotesInner) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntryDayNotesInner) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntryDayNotesInner) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntryDayNotesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *EntryDayNotesInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EntryDayNotesInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EntryDayNotesInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EntryDayNotesInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTable

`func (o *EntryDayNotesInner) GetTable() EntryTable`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *EntryDayNotesInner) GetTableOk() (*EntryTable, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *EntryDayNotesInner) SetTable(v EntryTable)`

SetTable sets Table field to given value.

### HasTable

`func (o *EntryDayNotesInner) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetChecked

`func (o *EntryDayNotesInner) GetChecked() bool`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *EntryDayNotesInner) GetCheckedOk() (*bool, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *EntryDayNotesInner) SetChecked(v bool)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *EntryDayNotesInner) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetExternalFileId

`func (o *EntryDayNotesInner) GetExternalFileId() string`

GetExternalFileId returns the ExternalFileId field if non-nil, zero value otherwise.

### GetExternalFileIdOk

`func (o *EntryDayNotesInner) GetExternalFileIdOk() (*string, bool)`

GetExternalFileIdOk returns a tuple with the ExternalFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFileId

`func (o *EntryDayNotesInner) SetExternalFileId(v string)`

SetExternalFileId sets ExternalFileId field to given value.

### HasExternalFileId

`func (o *EntryDayNotesInner) HasExternalFileId() bool`

HasExternalFileId returns a boolean if a field has been set.

### GetAssayRunId

`func (o *EntryDayNotesInner) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *EntryDayNotesInner) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *EntryDayNotesInner) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.

### HasAssayRunId

`func (o *EntryDayNotesInner) HasAssayRunId() bool`

HasAssayRunId returns a boolean if a field has been set.

### SetAssayRunIdNil

`func (o *EntryDayNotesInner) SetAssayRunIdNil(b bool)`

 SetAssayRunIdNil sets the value for AssayRunId to be an explicit nil

### UnsetAssayRunId
`func (o *EntryDayNotesInner) UnsetAssayRunId()`

UnsetAssayRunId ensures that no value is present for AssayRunId, not even an explicit nil
### GetAssayRunSchemaId

`func (o *EntryDayNotesInner) GetAssayRunSchemaId() string`

GetAssayRunSchemaId returns the AssayRunSchemaId field if non-nil, zero value otherwise.

### GetAssayRunSchemaIdOk

`func (o *EntryDayNotesInner) GetAssayRunSchemaIdOk() (*string, bool)`

GetAssayRunSchemaIdOk returns a tuple with the AssayRunSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunSchemaId

`func (o *EntryDayNotesInner) SetAssayRunSchemaId(v string)`

SetAssayRunSchemaId sets AssayRunSchemaId field to given value.

### HasAssayRunSchemaId

`func (o *EntryDayNotesInner) HasAssayRunSchemaId() bool`

HasAssayRunSchemaId returns a boolean if a field has been set.

### GetApiId

`func (o *EntryDayNotesInner) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *EntryDayNotesInner) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *EntryDayNotesInner) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *EntryDayNotesInner) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *EntryDayNotesInner) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *EntryDayNotesInner) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *EntryDayNotesInner) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *EntryDayNotesInner) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetAssayResultSchemaId

`func (o *EntryDayNotesInner) GetAssayResultSchemaId() string`

GetAssayResultSchemaId returns the AssayResultSchemaId field if non-nil, zero value otherwise.

### GetAssayResultSchemaIdOk

`func (o *EntryDayNotesInner) GetAssayResultSchemaIdOk() (*string, bool)`

GetAssayResultSchemaIdOk returns a tuple with the AssayResultSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayResultSchemaId

`func (o *EntryDayNotesInner) SetAssayResultSchemaId(v string)`

SetAssayResultSchemaId sets AssayResultSchemaId field to given value.

### HasAssayResultSchemaId

`func (o *EntryDayNotesInner) HasAssayResultSchemaId() bool`

HasAssayResultSchemaId returns a boolean if a field has been set.

### GetEntitySchemaId

`func (o *EntryDayNotesInner) GetEntitySchemaId() string`

GetEntitySchemaId returns the EntitySchemaId field if non-nil, zero value otherwise.

### GetEntitySchemaIdOk

`func (o *EntryDayNotesInner) GetEntitySchemaIdOk() (*string, bool)`

GetEntitySchemaIdOk returns a tuple with the EntitySchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySchemaId

`func (o *EntryDayNotesInner) SetEntitySchemaId(v string)`

SetEntitySchemaId sets EntitySchemaId field to given value.

### HasEntitySchemaId

`func (o *EntryDayNotesInner) HasEntitySchemaId() bool`

HasEntitySchemaId returns a boolean if a field has been set.

### GetPlateSchemaId

`func (o *EntryDayNotesInner) GetPlateSchemaId() string`

GetPlateSchemaId returns the PlateSchemaId field if non-nil, zero value otherwise.

### GetPlateSchemaIdOk

`func (o *EntryDayNotesInner) GetPlateSchemaIdOk() (*string, bool)`

GetPlateSchemaIdOk returns a tuple with the PlateSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateSchemaId

`func (o *EntryDayNotesInner) SetPlateSchemaId(v string)`

SetPlateSchemaId sets PlateSchemaId field to given value.

### HasPlateSchemaId

`func (o *EntryDayNotesInner) HasPlateSchemaId() bool`

HasPlateSchemaId returns a boolean if a field has been set.

### GetBoxSchemaId

`func (o *EntryDayNotesInner) GetBoxSchemaId() string`

GetBoxSchemaId returns the BoxSchemaId field if non-nil, zero value otherwise.

### GetBoxSchemaIdOk

`func (o *EntryDayNotesInner) GetBoxSchemaIdOk() (*string, bool)`

GetBoxSchemaIdOk returns a tuple with the BoxSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxSchemaId

`func (o *EntryDayNotesInner) SetBoxSchemaId(v string)`

SetBoxSchemaId sets BoxSchemaId field to given value.

### HasBoxSchemaId

`func (o *EntryDayNotesInner) HasBoxSchemaId() bool`

HasBoxSchemaId returns a boolean if a field has been set.

### GetMixtureSchemaId

`func (o *EntryDayNotesInner) GetMixtureSchemaId() string`

GetMixtureSchemaId returns the MixtureSchemaId field if non-nil, zero value otherwise.

### GetMixtureSchemaIdOk

`func (o *EntryDayNotesInner) GetMixtureSchemaIdOk() (*string, bool)`

GetMixtureSchemaIdOk returns a tuple with the MixtureSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixtureSchemaId

`func (o *EntryDayNotesInner) SetMixtureSchemaId(v string)`

SetMixtureSchemaId sets MixtureSchemaId field to given value.

### HasMixtureSchemaId

`func (o *EntryDayNotesInner) HasMixtureSchemaId() bool`

HasMixtureSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


