# RnaOligoBulkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Aliases** | Pointer to **[]string** | Aliases to add to the Oligo | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Oligo&#39;s authors. | [optional] 
**Bases** | Pointer to **string** | Base pairs of the oligo.  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Oligo. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Oligo. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the Oligo.  | [optional] 
**Name** | Pointer to **string** | Name of the Oligo.  | [optional] 
**SchemaId** | Pointer to **string** | ID of the oligo&#39;s schema.  | [optional] 

## Methods

### NewRnaOligoBulkUpdate

`func NewRnaOligoBulkUpdate() *RnaOligoBulkUpdate`

NewRnaOligoBulkUpdate instantiates a new RnaOligoBulkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaOligoBulkUpdateWithDefaults

`func NewRnaOligoBulkUpdateWithDefaults() *RnaOligoBulkUpdate`

NewRnaOligoBulkUpdateWithDefaults instantiates a new RnaOligoBulkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RnaOligoBulkUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RnaOligoBulkUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RnaOligoBulkUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RnaOligoBulkUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAliases

`func (o *RnaOligoBulkUpdate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *RnaOligoBulkUpdate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *RnaOligoBulkUpdate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *RnaOligoBulkUpdate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *RnaOligoBulkUpdate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *RnaOligoBulkUpdate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *RnaOligoBulkUpdate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *RnaOligoBulkUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *RnaOligoBulkUpdate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *RnaOligoBulkUpdate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *RnaOligoBulkUpdate) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *RnaOligoBulkUpdate) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetCustomFields

`func (o *RnaOligoBulkUpdate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RnaOligoBulkUpdate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RnaOligoBulkUpdate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RnaOligoBulkUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *RnaOligoBulkUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RnaOligoBulkUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RnaOligoBulkUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RnaOligoBulkUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *RnaOligoBulkUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RnaOligoBulkUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RnaOligoBulkUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RnaOligoBulkUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *RnaOligoBulkUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RnaOligoBulkUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RnaOligoBulkUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RnaOligoBulkUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *RnaOligoBulkUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *RnaOligoBulkUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *RnaOligoBulkUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *RnaOligoBulkUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


