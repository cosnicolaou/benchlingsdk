# BaseNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** | The type of the note.  Type determines what other fields are present. | [optional] 

## Methods

### NewBaseNotePart

`func NewBaseNotePart() *BaseNotePart`

NewBaseNotePart instantiates a new BaseNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNotePartWithDefaults

`func NewBaseNotePartWithDefaults() *BaseNotePart`

NewBaseNotePartWithDefaults instantiates a new BaseNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *BaseNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *BaseNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *BaseNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *BaseNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *BaseNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseNotePart) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


