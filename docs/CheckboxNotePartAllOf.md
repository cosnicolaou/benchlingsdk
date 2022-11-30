# CheckboxNotePartAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checked** | Pointer to **bool** | Indicates whether the checkbox is checked or not.  | [optional] 
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in text via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Text** | Pointer to **string** | The textual contents of the note. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckboxNotePartAllOf

`func NewCheckboxNotePartAllOf() *CheckboxNotePartAllOf`

NewCheckboxNotePartAllOf instantiates a new CheckboxNotePartAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckboxNotePartAllOfWithDefaults

`func NewCheckboxNotePartAllOfWithDefaults() *CheckboxNotePartAllOf`

NewCheckboxNotePartAllOfWithDefaults instantiates a new CheckboxNotePartAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecked

`func (o *CheckboxNotePartAllOf) GetChecked() bool`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *CheckboxNotePartAllOf) GetCheckedOk() (*bool, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *CheckboxNotePartAllOf) SetChecked(v bool)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *CheckboxNotePartAllOf) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetLinks

`func (o *CheckboxNotePartAllOf) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckboxNotePartAllOf) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckboxNotePartAllOf) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CheckboxNotePartAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *CheckboxNotePartAllOf) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CheckboxNotePartAllOf) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CheckboxNotePartAllOf) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CheckboxNotePartAllOf) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *CheckboxNotePartAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckboxNotePartAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckboxNotePartAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckboxNotePartAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


