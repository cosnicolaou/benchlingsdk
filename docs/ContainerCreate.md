# ContainerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentStorageId** | Pointer to **string** | ID of containing parent storage, can also specify a coordinate for plates and boxes (e.g. plt_2bAks9dx:a2). | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**SchemaId** | **string** |  | 

## Methods

### NewContainerCreate

`func NewContainerCreate(schemaId string, ) *ContainerCreate`

NewContainerCreate instantiates a new ContainerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerCreateWithDefaults

`func NewContainerCreateWithDefaults() *ContainerCreate`

NewContainerCreateWithDefaults instantiates a new ContainerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ContainerCreate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContainerCreate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContainerCreate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ContainerCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *ContainerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *ContainerCreate) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *ContainerCreate) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *ContainerCreate) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *ContainerCreate) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### GetBarcode

`func (o *ContainerCreate) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ContainerCreate) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ContainerCreate) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ContainerCreate) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetProjectId

`func (o *ContainerCreate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ContainerCreate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ContainerCreate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ContainerCreate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ContainerCreate) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ContainerCreate) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSchemaId

`func (o *ContainerCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ContainerCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ContainerCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


