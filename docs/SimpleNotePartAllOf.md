# SimpleNotePartAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in text via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Text** | Pointer to **string** | The textual contents of the note. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSimpleNotePartAllOf

`func NewSimpleNotePartAllOf() *SimpleNotePartAllOf`

NewSimpleNotePartAllOf instantiates a new SimpleNotePartAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleNotePartAllOfWithDefaults

`func NewSimpleNotePartAllOfWithDefaults() *SimpleNotePartAllOf`

NewSimpleNotePartAllOfWithDefaults instantiates a new SimpleNotePartAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SimpleNotePartAllOf) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SimpleNotePartAllOf) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SimpleNotePartAllOf) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SimpleNotePartAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *SimpleNotePartAllOf) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SimpleNotePartAllOf) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SimpleNotePartAllOf) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SimpleNotePartAllOf) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *SimpleNotePartAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleNotePartAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleNotePartAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleNotePartAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


