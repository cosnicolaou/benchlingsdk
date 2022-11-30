# EntryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorIds** | Pointer to **string** | IDs of users to set as the entry&#39;s authors. | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Schema fields to set on the entry | [optional] 
**FolderId** | Pointer to **string** | ID of the folder that will contain the entry | [optional] 
**Name** | Pointer to **string** | New name of the entry | [optional] 
**SchemaId** | Pointer to **string** | ID of the schema for the entry | [optional] 

## Methods

### NewEntryUpdate

`func NewEntryUpdate() *EntryUpdate`

NewEntryUpdate instantiates a new EntryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryUpdateWithDefaults

`func NewEntryUpdateWithDefaults() *EntryUpdate`

NewEntryUpdateWithDefaults instantiates a new EntryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorIds

`func (o *EntryUpdate) GetAuthorIds() string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *EntryUpdate) GetAuthorIdsOk() (*string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *EntryUpdate) SetAuthorIds(v string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *EntryUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetFields

`func (o *EntryUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EntryUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EntryUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EntryUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *EntryUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *EntryUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *EntryUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *EntryUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *EntryUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntryUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *EntryUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *EntryUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *EntryUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *EntryUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


