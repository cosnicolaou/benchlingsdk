# CustomEntityBaseRequestForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the custom entity | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the custom entity&#39;s authors. | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the custom entity. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Schema fields to set on the custom entity. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder that the entity is moved into | [optional] 
**Name** | **string** |  | 
**SchemaId** | **string** |  | 

## Methods

### NewCustomEntityBaseRequestForCreate

`func NewCustomEntityBaseRequestForCreate(name string, schemaId string, ) *CustomEntityBaseRequestForCreate`

NewCustomEntityBaseRequestForCreate instantiates a new CustomEntityBaseRequestForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityBaseRequestForCreateWithDefaults

`func NewCustomEntityBaseRequestForCreateWithDefaults() *CustomEntityBaseRequestForCreate`

NewCustomEntityBaseRequestForCreateWithDefaults instantiates a new CustomEntityBaseRequestForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *CustomEntityBaseRequestForCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CustomEntityBaseRequestForCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CustomEntityBaseRequestForCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CustomEntityBaseRequestForCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *CustomEntityBaseRequestForCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *CustomEntityBaseRequestForCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *CustomEntityBaseRequestForCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *CustomEntityBaseRequestForCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomEntityBaseRequestForCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomEntityBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomEntityBaseRequestForCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomEntityBaseRequestForCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *CustomEntityBaseRequestForCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomEntityBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomEntityBaseRequestForCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomEntityBaseRequestForCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *CustomEntityBaseRequestForCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *CustomEntityBaseRequestForCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *CustomEntityBaseRequestForCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *CustomEntityBaseRequestForCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *CustomEntityBaseRequestForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEntityBaseRequestForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEntityBaseRequestForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *CustomEntityBaseRequestForCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *CustomEntityBaseRequestForCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *CustomEntityBaseRequestForCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


