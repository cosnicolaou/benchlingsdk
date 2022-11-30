# EntryTableCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to [**EntryTableCellLink**](EntryTableCellLink.md) |  | [optional] 
**Text** | Pointer to **string** | The textual content of the cell. If the cell was originally a formula, this will be the evaluated version of the formula.  | [optional] 

## Methods

### NewEntryTableCell

`func NewEntryTableCell() *EntryTableCell`

NewEntryTableCell instantiates a new EntryTableCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTableCellWithDefaults

`func NewEntryTableCellWithDefaults() *EntryTableCell`

NewEntryTableCellWithDefaults instantiates a new EntryTableCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *EntryTableCell) GetLink() EntryTableCellLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *EntryTableCell) GetLinkOk() (*EntryTableCellLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *EntryTableCell) SetLink(v EntryTableCellLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *EntryTableCell) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetText

`func (o *EntryTableCell) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EntryTableCell) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EntryTableCell) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EntryTableCell) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


