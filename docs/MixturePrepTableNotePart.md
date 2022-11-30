# MixturePrepTableNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 
**MixtureSchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewMixturePrepTableNotePart

`func NewMixturePrepTableNotePart() *MixturePrepTableNotePart`

NewMixturePrepTableNotePart instantiates a new MixturePrepTableNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixturePrepTableNotePartWithDefaults

`func NewMixturePrepTableNotePartWithDefaults() *MixturePrepTableNotePart`

NewMixturePrepTableNotePartWithDefaults instantiates a new MixturePrepTableNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *MixturePrepTableNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *MixturePrepTableNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *MixturePrepTableNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *MixturePrepTableNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *MixturePrepTableNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MixturePrepTableNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MixturePrepTableNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MixturePrepTableNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApiId

`func (o *MixturePrepTableNotePart) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *MixturePrepTableNotePart) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *MixturePrepTableNotePart) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *MixturePrepTableNotePart) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *MixturePrepTableNotePart) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MixturePrepTableNotePart) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *MixturePrepTableNotePart) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *MixturePrepTableNotePart) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetMixtureSchemaId

`func (o *MixturePrepTableNotePart) GetMixtureSchemaId() string`

GetMixtureSchemaId returns the MixtureSchemaId field if non-nil, zero value otherwise.

### GetMixtureSchemaIdOk

`func (o *MixturePrepTableNotePart) GetMixtureSchemaIdOk() (*string, bool)`

GetMixtureSchemaIdOk returns a tuple with the MixtureSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixtureSchemaId

`func (o *MixturePrepTableNotePart) SetMixtureSchemaId(v string)`

SetMixtureSchemaId sets MixtureSchemaId field to given value.

### HasMixtureSchemaId

`func (o *MixturePrepTableNotePart) HasMixtureSchemaId() bool`

HasMixtureSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


