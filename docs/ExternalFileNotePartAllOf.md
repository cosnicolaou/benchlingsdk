# ExternalFileNotePartAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalFileId** | Pointer to **string** | The ID of the external file. Use the &#39;Get an external file&#39; endpoint to retrieve metadata about it.  | [optional] 
**Links** | Pointer to [**[]EntryLink**](EntryLink.md) | Array of links referenced in the caption via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note.  | [optional] 
**Text** | Pointer to **string** | The caption of the file attachment. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalFileNotePartAllOf

`func NewExternalFileNotePartAllOf() *ExternalFileNotePartAllOf`

NewExternalFileNotePartAllOf instantiates a new ExternalFileNotePartAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalFileNotePartAllOfWithDefaults

`func NewExternalFileNotePartAllOfWithDefaults() *ExternalFileNotePartAllOf`

NewExternalFileNotePartAllOfWithDefaults instantiates a new ExternalFileNotePartAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalFileId

`func (o *ExternalFileNotePartAllOf) GetExternalFileId() string`

GetExternalFileId returns the ExternalFileId field if non-nil, zero value otherwise.

### GetExternalFileIdOk

`func (o *ExternalFileNotePartAllOf) GetExternalFileIdOk() (*string, bool)`

GetExternalFileIdOk returns a tuple with the ExternalFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFileId

`func (o *ExternalFileNotePartAllOf) SetExternalFileId(v string)`

SetExternalFileId sets ExternalFileId field to given value.

### HasExternalFileId

`func (o *ExternalFileNotePartAllOf) HasExternalFileId() bool`

HasExternalFileId returns a boolean if a field has been set.

### GetLinks

`func (o *ExternalFileNotePartAllOf) GetLinks() []EntryLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalFileNotePartAllOf) GetLinksOk() (*[]EntryLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalFileNotePartAllOf) SetLinks(v []EntryLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalFileNotePartAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetText

`func (o *ExternalFileNotePartAllOf) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ExternalFileNotePartAllOf) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ExternalFileNotePartAllOf) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ExternalFileNotePartAllOf) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *ExternalFileNotePartAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalFileNotePartAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalFileNotePartAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalFileNotePartAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


