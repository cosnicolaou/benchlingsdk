# RegistrationTableNotePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indentation** | Pointer to **int32** | All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example. | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 
**ApiId** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]StructuredTableColumnInfo**](StructuredTableColumnInfo.md) |  | [optional] 
**EntitySchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistrationTableNotePart

`func NewRegistrationTableNotePart() *RegistrationTableNotePart`

NewRegistrationTableNotePart instantiates a new RegistrationTableNotePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationTableNotePartWithDefaults

`func NewRegistrationTableNotePartWithDefaults() *RegistrationTableNotePart`

NewRegistrationTableNotePartWithDefaults instantiates a new RegistrationTableNotePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndentation

`func (o *RegistrationTableNotePart) GetIndentation() int32`

GetIndentation returns the Indentation field if non-nil, zero value otherwise.

### GetIndentationOk

`func (o *RegistrationTableNotePart) GetIndentationOk() (*int32, bool)`

GetIndentationOk returns a tuple with the Indentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentation

`func (o *RegistrationTableNotePart) SetIndentation(v int32)`

SetIndentation sets Indentation field to given value.

### HasIndentation

`func (o *RegistrationTableNotePart) HasIndentation() bool`

HasIndentation returns a boolean if a field has been set.

### GetType

`func (o *RegistrationTableNotePart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistrationTableNotePart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistrationTableNotePart) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistrationTableNotePart) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApiId

`func (o *RegistrationTableNotePart) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *RegistrationTableNotePart) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *RegistrationTableNotePart) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *RegistrationTableNotePart) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetColumns

`func (o *RegistrationTableNotePart) GetColumns() []StructuredTableColumnInfo`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *RegistrationTableNotePart) GetColumnsOk() (*[]StructuredTableColumnInfo, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *RegistrationTableNotePart) SetColumns(v []StructuredTableColumnInfo)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *RegistrationTableNotePart) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEntitySchemaId

`func (o *RegistrationTableNotePart) GetEntitySchemaId() string`

GetEntitySchemaId returns the EntitySchemaId field if non-nil, zero value otherwise.

### GetEntitySchemaIdOk

`func (o *RegistrationTableNotePart) GetEntitySchemaIdOk() (*string, bool)`

GetEntitySchemaIdOk returns a tuple with the EntitySchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySchemaId

`func (o *RegistrationTableNotePart) SetEntitySchemaId(v string)`

SetEntitySchemaId sets EntitySchemaId field to given value.

### HasEntitySchemaId

`func (o *RegistrationTableNotePart) HasEntitySchemaId() bool`

HasEntitySchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


