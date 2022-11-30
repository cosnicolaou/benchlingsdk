# SimpleNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in text via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Text** | Pointer to **string** | The textual contents of the note. | [optional] 

## Methods

### NewSimpleNotePart

`func NewSimpleNotePart() *SimpleNotePart`

NewSimpleNotePart instantiates a new SimpleNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleNotePartWithDefaults

`func NewSimpleNotePartWithDefaults() *SimpleNotePart`

NewSimpleNotePartWithDefaults instantiates a new SimpleNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *SimpleNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *SimpleNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *SimpleNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *SimpleNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *SimpleNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *SimpleNotePart) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SimpleNotePart) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SimpleNotePart) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SimpleNotePart) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *SimpleNotePart) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SimpleNotePart) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SimpleNotePart) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SimpleNotePart) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


