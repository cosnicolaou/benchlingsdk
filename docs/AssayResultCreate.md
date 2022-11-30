# AssayResultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldValidation** | Pointer to [**map[string]UserValidation**](UserValidation.md) | Dictionary mapping field names to UserValidation Resources.  | [optional] 
**Fields** | [**AssayResultCreateFields**](AssayResultCreateFields.md) |  | 
**Id** | Pointer to **string** | UUID | [optional] 
**ProjectId** | Pointer to **NullableString** | The project that the assay result should be uploaded to. Only users with read access to the project will be able to read the assay result. Leaving this empty will result in only the creator having read access.  | [optional] 
**SchemaId** | **string** | ID of result schema under which to upload this result | 

## Methods

### NewAssayResultCreate

`func NewAssayResultCreate(fields AssayResultCreateFields, schemaId string, ) *AssayResultCreate`

NewAssayResultCreate instantiates a new AssayResultCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayResultCreateWithDefaults

`func NewAssayResultCreateWithDefaults() *AssayResultCreate`

NewAssayResultCreateWithDefaults instantiates a new AssayResultCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValidation

`func (o *AssayResultCreate) GetFieldValidation() map[string]UserValidation`

GetFieldValidation returns the FieldValidation field if non-nil, zero value otherwise.

### GetFieldValidationOk

`func (o *AssayResultCreate) GetFieldValidationOk() (*map[string]UserValidation, bool)`

GetFieldValidationOk returns a tuple with the FieldValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValidation

`func (o *AssayResultCreate) SetFieldValidation(v map[string]UserValidation)`

SetFieldValidation sets FieldValidation field to given value.

### HasFieldValidation

`func (o *AssayResultCreate) HasFieldValidation() bool`

HasFieldValidation returns a boolean if a field has been set.

### GetFields

`func (o *AssayResultCreate) GetFields() AssayResultCreateFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AssayResultCreate) GetFieldsOk() (*AssayResultCreateFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AssayResultCreate) SetFields(v AssayResultCreateFields)`

SetFields sets Fields field to given value.


### GetId

`func (o *AssayResultCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayResultCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayResultCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayResultCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *AssayResultCreate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssayResultCreate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssayResultCreate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AssayResultCreate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *AssayResultCreate) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *AssayResultCreate) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSchemaId

`func (o *AssayResultCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *AssayResultCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *AssayResultCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


