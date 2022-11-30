# AssayRunCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**AssayRunCreateFields**](AssayRunCreateFields.md) |  | 
**Id** | Pointer to **string** | ID of assay run | [optional] 
**ProjectId** | Pointer to **string** | The project that the assay run should be uploaded to. Only users with read access to the project will be able to read the assay run. Leaving this empty will result in only the creator having read access.  | [optional] 
**SchemaId** | **string** | ID of assay schema that assay run conforms to | 
**ValidationComment** | Pointer to **string** | Additional information about the validation status | [optional] 
**ValidationStatus** | Pointer to [**AssayRunValidationStatus**](AssayRunValidationStatus.md) |  | [optional] 

## Methods

### NewAssayRunCreate

`func NewAssayRunCreate(fields AssayRunCreateFields, schemaId string, ) *AssayRunCreate`

NewAssayRunCreate instantiates a new AssayRunCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunCreateWithDefaults

`func NewAssayRunCreateWithDefaults() *AssayRunCreate`

NewAssayRunCreateWithDefaults instantiates a new AssayRunCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AssayRunCreate) GetFields() AssayRunCreateFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AssayRunCreate) GetFieldsOk() (*AssayRunCreateFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AssayRunCreate) SetFields(v AssayRunCreateFields)`

SetFields sets Fields field to given value.


### GetId

`func (o *AssayRunCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayRunCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayRunCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayRunCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *AssayRunCreate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssayRunCreate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssayRunCreate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *AssayRunCreate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSchemaId

`func (o *AssayRunCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *AssayRunCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *AssayRunCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.


### GetValidationComment

`func (o *AssayRunCreate) GetValidationComment() string`

GetValidationComment returns the ValidationComment field if non-nil, zero value otherwise.

### GetValidationCommentOk

`func (o *AssayRunCreate) GetValidationCommentOk() (*string, bool)`

GetValidationCommentOk returns a tuple with the ValidationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationComment

`func (o *AssayRunCreate) SetValidationComment(v string)`

SetValidationComment sets ValidationComment field to given value.

### HasValidationComment

`func (o *AssayRunCreate) HasValidationComment() bool`

HasValidationComment returns a boolean if a field has been set.

### GetValidationStatus

`func (o *AssayRunCreate) GetValidationStatus() AssayRunValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *AssayRunCreate) GetValidationStatusOk() (*AssayRunValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *AssayRunCreate) SetValidationStatus(v AssayRunValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *AssayRunCreate) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


