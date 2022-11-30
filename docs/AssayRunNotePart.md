# AssayRunNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**AssayRunId** | Pointer to **NullableString** |  | [optional] 
**AssayRunSchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewAssayRunNotePart

`func NewAssayRunNotePart() *AssayRunNotePart`

NewAssayRunNotePart instantiates a new AssayRunNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunNotePartWithDefaults

`func NewAssayRunNotePartWithDefaults() *AssayRunNotePart`

NewAssayRunNotePartWithDefaults instantiates a new AssayRunNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *AssayRunNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *AssayRunNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *AssayRunNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *AssayRunNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *AssayRunNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssayRunNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssayRunNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssayRunNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssayRunId

`func (o *AssayRunNotePart) GetAssayRunId() string`

GetAssayRunId returns the AssayRunId field if non-nil, zero value otherwise.

### GetAssayRunIdOk

`func (o *AssayRunNotePart) GetAssayRunIdOk() (*string, bool)`

GetAssayRunIdOk returns a tuple with the AssayRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunId

`func (o *AssayRunNotePart) SetAssayRunId(v string)`

SetAssayRunId sets AssayRunId field to given value.

### HasAssayRunId

`func (o *AssayRunNotePart) HasAssayRunId() bool`

HasAssayRunId returns a boolean if a field has been set.

### SetAssayRunIdNil

`func (o *AssayRunNotePart) SetAssayRunIdNil(b bool)`

 SetAssayRunIdNil sets the value for AssayRunId to be an explicit nil

### UnsetAssayRunId
`func (o *AssayRunNotePart) UnsetAssayRunId()`

UnsetAssayRunId ensures that no value is present for AssayRunId, not even an explicit nil
### GetAssayRunSchemaId

`func (o *AssayRunNotePart) GetAssayRunSchemaId() string`

GetAssayRunSchemaId returns the AssayRunSchemaId field if non-nil, zero value otherwise.

### GetAssayRunSchemaIdOk

`func (o *AssayRunNotePart) GetAssayRunSchemaIdOk() (*string, bool)`

GetAssayRunSchemaIdOk returns a tuple with the AssayRunSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRunSchemaId

`func (o *AssayRunNotePart) SetAssayRunSchemaId(v string)`

SetAssayRunSchemaId sets AssayRunSchemaId field to given value.

### HasAssayRunSchemaId

`func (o *AssayRunNotePart) HasAssayRunSchemaId() bool`

HasAssayRunSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


