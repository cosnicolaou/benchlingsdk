# OligoBaseRequestForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the Oligo | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the Oligo&#39;s authors. | [optional] 
**Bases** | **string** | Base pairs of the oligo.  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the Oligo. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the Oligo. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the Oligo.  | [optional] 
**Name** | **string** | Name of the Oligo.  | 
**SchemaId** | Pointer to **string** | ID of the oligo&#39;s schema.  | [optional] 

## Methods

### NewOligoBaseRequestForCreate

`func NewOligoBaseRequestForCreate(bases string, name string, ) *OligoBaseRequestForCreate`

NewOligoBaseRequestForCreate instantiates a new OligoBaseRequestForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOligoBaseRequestForCreateWithDefaults

`func NewOligoBaseRequestForCreateWithDefaults() *OligoBaseRequestForCreate`

NewOligoBaseRequestForCreateWithDefaults instantiates a new OligoBaseRequestForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *OligoBaseRequestForCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *OligoBaseRequestForCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *OligoBaseRequestForCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *OligoBaseRequestForCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAuthorIds

`func (o *OligoBaseRequestForCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *OligoBaseRequestForCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *OligoBaseRequestForCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *OligoBaseRequestForCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *OligoBaseRequestForCreate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *OligoBaseRequestForCreate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *OligoBaseRequestForCreate) SetBases(v string)`

SetBases sets Bases field to given value.


### GetCustomFields

`func (o *OligoBaseRequestForCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OligoBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OligoBaseRequestForCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OligoBaseRequestForCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *OligoBaseRequestForCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *OligoBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *OligoBaseRequestForCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *OligoBaseRequestForCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *OligoBaseRequestForCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *OligoBaseRequestForCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *OligoBaseRequestForCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *OligoBaseRequestForCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetName

`func (o *OligoBaseRequestForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OligoBaseRequestForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OligoBaseRequestForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSchemaId

`func (o *OligoBaseRequestForCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *OligoBaseRequestForCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *OligoBaseRequestForCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *OligoBaseRequestForCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


