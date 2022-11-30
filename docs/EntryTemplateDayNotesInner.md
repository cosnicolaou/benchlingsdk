# EntryTemplateDayNotesInner

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

### NewEntryTemplateDayNotesInner

`func NewEntryTemplateDayNotesInner() *EntryTemplateDayNotesInner`

NewEntryTemplateDayNotesInner instantiates a new EntryTemplateDayNotesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTemplateDayNotesInnerWithDefaults

`func NewEntryTemplateDayNotesInnerWithDefaults() *EntryTemplateDayNotesInner`

NewEntryTemplateDayNotesInnerWithDefaults instantiates a new EntryTemplateDayNotesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *EntryTemplateDayNotesInner) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *EntryTemplateDayNotesInner) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *EntryTemplateDayNotesInner) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *EntryTemplateDayNotesInner) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *EntryTemplateDayNotesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntryTemplateDayNotesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntryTemplateDayNotesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntryTemplateDayNotesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *EntryTemplateDayNotesInner) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntryTemplateDayNotesInner) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntryTemplateDayNotesInner) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntryTemplateDayNotesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *EntryTemplateDayNotesInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EntryTemplateDayNotesInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EntryTemplateDayNotesInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EntryTemplateDayNotesInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTable

`func (o *EntryTemplateDayNotesInner) GetTable() EntryTable`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *EntryTemplateDayNotesInner) GetTableOk() (*EntryTable, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *EntryTemplateDayNotesInner) SetTable(v EntryTable)`

SetTable sets Table field to given value.

### HasTable

`func (o *EntryTemplateDayNotesInner) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetChecked

`func (o *EntryTemplateDayNotesInner) GetChecked() bool`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *EntryTemplateDayNotesInner) GetCheckedOk() (*bool, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *EntryTemplateDayNotesInner) SetChecked(v bool)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *EntryTemplateDayNotesInner) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetExternalFileId

`func (o *EntryTemplateDayNotesInner) GetExternalFileId() string`

GetExternalFileId returns the ExternalFileId field if non-nil, zero value otherwise.

### GetExternalFileIdOk

`func (o *EntryTemplateDayNotesInner) GetExternalFileIdOk() (*string, bool)`

GetExternalFileIdOk returns a tuple with the ExternalFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFileId

`func (o *EntryTemplateDayNotesInner) SetExternalFileId(v string)`

SetExternalFileId sets ExternalFileId field to given value.

### HasExternalFileId

`func (o *EntryTemplateDayNotesInner) HasExternalFileId() bool`

HasExternalFileId returns a boolean if a field has been set.

### GetAssayRunId

`func (o *EntryTemplateDayNotesInner) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *EntryTemplateDayNotesInner) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *EntryTemplateDayNotesInner) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.

### HasAssayRunId

`func (o *EntryTemplateDayNotesInner) HasAssayRunId() bool`

HasAssayRunId returns a boolean if a field has been set.

### SetAssayRunIdNil

`func (o *EntryTemplateDayNotesInner) SetAssayRunIdNil(b bool)`

 SetAssayRunIdNil sets the value for AssayRunId to be an explicit nil

### UnsetAssayRunId
`func (o *EntryTemplateDayNotesInner) UnsetAssayRunId()`

UnsetAssayRunId ensures that no value is present for AssayRunId, not even an explicit nil
### GetAssayRunSchemaId

`func (o *EntryTemplateDayNotesInner) GetAssayRunSchemaId() string`

GetAssayRunSchemaId returns the AssayRunSchemaId field if non-nil, zero value otherwise.

### GetAssayRunSchemaIdOk

`func (o *EntryTemplateDayNotesInner) GetAssayRunSchemaIdOk() (*string, bool)`

GetAssayRunSchemaIdOk returns a tuple with the AssayRunSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunSchemaId

`func (o *EntryTemplateDayNotesInner) SetAssayRunSchemaId(v string)`

SetAssayRunSchemaId sets AssayRunSchemaId field to given value.

### HasAssayRunSchemaId

`func (o *EntryTemplateDayNotesInner) HasAssayRunSchemaId() bool`

HasAssayRunSchemaId returns a boolean if a field has been set.

### GetApiId

`func (o *EntryTemplateDayNotesInner) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *EntryTemplateDayNotesInner) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *EntryTemplateDayNotesInner) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *EntryTemplateDayNotesInner) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *EntryTemplateDayNotesInner) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *EntryTemplateDayNotesInner) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *EntryTemplateDayNotesInner) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *EntryTemplateDayNotesInner) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetAssayResultSchemaId

`func (o *EntryTemplateDayNotesInner) GetAssayResultSchemaId() string`

GetAssayResultSchemaId returns the AssayResultSchemaId field if non-nil, zero value otherwise.

### GetAssayResultSchemaIdOk

`func (o *EntryTemplateDayNotesInner) GetAssayResultSchemaIdOk() (*string, bool)`

GetAssayResultSchemaIdOk returns a tuple with the AssayResultSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayResultSchemaId

`func (o *EntryTemplateDayNotesInner) SetAssayResultSchemaId(v string)`

SetAssayResultSchemaId sets AssayResultSchemaId field to given value.

### HasAssayResultSchemaId

`func (o *EntryTemplateDayNotesInner) HasAssayResultSchemaId() bool`

HasAssayResultSchemaId returns a boolean if a field has been set.

### GetEntitySchemaId

`func (o *EntryTemplateDayNotesInner) GetEntitySchemaId() string`

GetEntitySchemaId returns the EntitySchemaId field if non-nil, zero value otherwise.

### GetEntitySchemaIdOk

`func (o *EntryTemplateDayNotesInner) GetEntitySchemaIdOk() (*string, bool)`

GetEntitySchemaIdOk returns a tuple with the EntitySchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySchemaId

`func (o *EntryTemplateDayNotesInner) SetEntitySchemaId(v string)`

SetEntitySchemaId sets EntitySchemaId field to given value.

### HasEntitySchemaId

`func (o *EntryTemplateDayNotesInner) HasEntitySchemaId() bool`

HasEntitySchemaId returns a boolean if a field has been set.

### GetPlateSchemaId

`func (o *EntryTemplateDayNotesInner) GetPlateSchemaId() string`

GetPlateSchemaId returns the PlateSchemaId field if non-nil, zero value otherwise.

### GetPlateSchemaIdOk

`func (o *EntryTemplateDayNotesInner) GetPlateSchemaIdOk() (*string, bool)`

GetPlateSchemaIdOk returns a tuple with the PlateSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateSchemaId

`func (o *EntryTemplateDayNotesInner) SetPlateSchemaId(v string)`

SetPlateSchemaId sets PlateSchemaId field to given value.

### HasPlateSchemaId

`func (o *EntryTemplateDayNotesInner) HasPlateSchemaId() bool`

HasPlateSchemaId returns a boolean if a field has been set.

### GetBoxSchemaId

`func (o *EntryTemplateDayNotesInner) GetBoxSchemaId() string`

GetBoxSchemaId returns the BoxSchemaId field if non-nil, zero value otherwise.

### GetBoxSchemaIdOk

`func (o *EntryTemplateDayNotesInner) GetBoxSchemaIdOk() (*string, bool)`

GetBoxSchemaIdOk returns a tuple with the BoxSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxSchemaId

`func (o *EntryTemplateDayNotesInner) SetBoxSchemaId(v string)`

SetBoxSchemaId sets BoxSchemaId field to given value.

### HasBoxSchemaId

`func (o *EntryTemplateDayNotesInner) HasBoxSchemaId() bool`

HasBoxSchemaId returns a boolean if a field has been set.

### GetMixtureSchemaId

`func (o *EntryTemplateDayNotesInner) GetMixtureSchemaId() string`

GetMixtureSchemaId returns the MixtureSchemaId field if non-nil, zero value otherwise.

### GetMixtureSchemaIdOk

`func (o *EntryTemplateDayNotesInner) GetMixtureSchemaIdOk() (*string, bool)`

GetMixtureSchemaIdOk returns a tuple with the MixtureSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixtureSchemaId

`func (o *EntryTemplateDayNotesInner) SetMixtureSchemaId(v string)`

SetMixtureSchemaId sets MixtureSchemaId field to given value.

### HasMixtureSchemaId

`func (o *EntryTemplateDayNotesInner) HasMixtureSchemaId() bool`

HasMixtureSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


