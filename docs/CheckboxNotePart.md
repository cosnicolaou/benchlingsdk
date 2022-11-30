# CheckboxNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**Checked** | Pointer to **bool** | Indicates whether the checkbox is checked or not.  | [optional] 
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in text via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Text** | Pointer to **string** | The textual contents of the note. | [optional] 

## Methods

### NewCheckboxNotePart

`func NewCheckboxNotePart() *CheckboxNotePart`

NewCheckboxNotePart instantiates a new CheckboxNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckboxNotePartWithDefaults

`func NewCheckboxNotePartWithDefaults() *CheckboxNotePart`

NewCheckboxNotePartWithDefaults instantiates a new CheckboxNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *CheckboxNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *CheckboxNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *CheckboxNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *CheckboxNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *CheckboxNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckboxNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckboxNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckboxNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChecked

`func (o *CheckboxNotePart) GetChecked() bool`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *CheckboxNotePart) GetCheckedOk() (*bool, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *CheckboxNotePart) SetChecked(v bool)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *CheckboxNotePart) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetLinks

`func (o *CheckboxNotePart) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckboxNotePart) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckboxNotePart) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CheckboxNotePart) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *CheckboxNotePart) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CheckboxNotePart) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CheckboxNotePart) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CheckboxNotePart) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


