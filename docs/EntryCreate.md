# EntryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorIds** | Pointer to [**EntryCreateAuthorIds**](EntryCreateAuthorIds.md) |  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the entry | [optional] 
**EntryTemplateId** | Pointer to **string** | ID of the template to clone the entry from | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the entry. Must correspond with the schema&#39;s field definitions.  | [optional] 
**FolderId** | **string** | ID of the folder that will contain the entry | 
**InitialTables** | Pointer to [**[]InitialTable**](InitialTable.md) | An array of table API IDs and blob id pairs to seed tables from the template while creating the entry. The entryTemplateId parameter must be set to use this parameter. The table API IDs should be the API Identifiers of the tables in the given template. - If a template table has one row, the values in that row act as default values for cloned entries. - If a template table has multiple rows, there is no default value and those rows are added to the cloned entry along with the provided csv data. - If a table has default values, they will be populated in any respective undefined columns in the csv data. - If a table has no default values, undefined columns from csv data will be empty. - If no csv data is provided for a table, the table in the entry will be populated with whatever values are in the respective template table.  | [optional] 
**Name** | **string** | Name of the entry | 
**SchemaId** | Pointer to **string** | ID of the entry&#39;s schema | [optional] 

## Methods

### NewEntryCreate

`func NewEntryCreate(folderId string, name string, ) *EntryCreate`

NewEntryCreate instantiates a new EntryCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryCreateWithDefaults

`func NewEntryCreateWithDefaults() *EntryCreate`

NewEntryCreateWithDefaults instantiates a new EntryCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorIds

`func (o *EntryCreate) GetAuthorIds() EntryCreateAuthorIds`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *EntryCreate) GetAuthorIdsOk() (*EntryCreateAuthorIds, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *EntryCreate) SetAuthorIds(v EntryCreateAuthorIds)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *EntryCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *EntryCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EntryCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EntryCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EntryCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntryTemplateId

`func (o *EntryCreate) GetEntryTemplateId() string`

GetEntryTemplateId returns the EntryTemplateId field if non-nil, zero value otherwise.

### GetEntryTemplateIdOk

`func (o *EntryCreate) GetEntryTemplateIdOk() (*string, bool)`

GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTemplateId

`func (o *EntryCreate) SetEntryTemplateId(v string)`

SetEntryTemplateId sets EntryTemplateId field to given value.

### HasEntryTemplateId

`func (o *EntryCreate) HasEntryTemplateId() bool`

HasEntryTemplateId returns a boolean if a field has been set.

### GetFields

`func (o *EntryCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EntryCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EntryCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EntryCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *EntryCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *EntryCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *EntryCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetInitialTables

`func (o *EntryCreate) GetInitialTables() []InitialTable`

GetInitialTables returns the InitialTables field if non-nil, zero value otherwise.

### GetInitialTablesOk

`func (o *EntryCreate) GetInitialTablesOk() (*[]InitialTable, bool)`

GetInitialTablesOk returns a tuple with the InitialTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialTables

`func (o *EntryCreate) SetInitialTables(v []InitialTable)`

SetInitialTables sets InitialTables field to given value.

### HasInitialTables

`func (o *EntryCreate) HasInitialTables() bool`

HasInitialTables returns a boolean if a field has been set.

### GetName

`func (o *EntryCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *EntryCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *EntryCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *EntryCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *EntryCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


