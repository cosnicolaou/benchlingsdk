# PlateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** |  | [optional] 
**ContainerSchemaId** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentStorageId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**SchemaId** | **string** |  | 
**Wells** | Pointer to [**map[string]PlateCreateWellsValue**](PlateCreateWellsValue.md) |  | [optional] 

## Methods

### NewPlateCreate

`func NewPlateCreate(schemaId string, ) *PlateCreate`

NewPlateCreate instantiates a new PlateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlateCreateWithDefaults

`func NewPlateCreateWithDefaults() *PlateCreate`

NewPlateCreateWithDefaults instantiates a new PlateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *PlateCreate) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *PlateCreate) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *PlateCreate) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *PlateCreate) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetContainerSchemaId

`func (o *PlateCreate) GetContainerSchemaId() string`

GetContainerSchemaId returns the ContainerSchemaId field if non-nil, zero value otherwise.

### GetContainerSchemaIdOk

`func (o *PlateCreate) GetContainerSchemaIdOk() (*string, bool)`

GetContainerSchemaIdOk returns a tuple with the ContainerSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSchemaId

`func (o *PlateCreate) SetContainerSchemaId(v string)`

SetContainerSchemaId sets ContainerSchemaId field to given value.

### HasContainerSchemaId

`func (o *PlateCreate) HasContainerSchemaId() bool`

HasContainerSchemaId returns a boolean if a field has been set.

### GetFields

`func (o *PlateCreate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PlateCreate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PlateCreate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PlateCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *PlateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlateCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlateCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *PlateCreate) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *PlateCreate) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *PlateCreate) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *PlateCreate) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### GetProjectId

`func (o *PlateCreate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PlateCreate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PlateCreate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PlateCreate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSchemaId

`func (o *PlateCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *PlateCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *PlateCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.


### GetWells

`func (o *PlateCreate) GetWells() map[string]PlateCreateWellsValue`

GetWells returns the Wells field if non-nil, zero value otherwise.

### GetWellsOk

`func (o *PlateCreate) GetWellsOk() (*map[string]PlateCreateWellsValue, bool)`

GetWellsOk returns a tuple with the Wells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWells

`func (o *PlateCreate) SetWells(v map[string]PlateCreateWellsValue)`

SetWells sets Wells field to given value.

### HasWells

`func (o *PlateCreate) HasWells() bool`

HasWells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


