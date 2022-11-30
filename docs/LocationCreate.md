# LocationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Name** | **string** |  | 
**ParentStorageId** | Pointer to **string** |  | [optional] 
**SchemaId** | **string** |  | 

## Methods

### NewLocationCreate

`func NewLocationCreate(name string, schemaId string, ) *LocationCreate`

NewLocationCreate instantiates a new LocationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationCreateWithDefaults

`func NewLocationCreateWithDefaults() *LocationCreate`

NewLocationCreateWithDefaults instantiates a new LocationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *LocationCreate) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *LocationCreate) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *LocationCreate) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *LocationCreate) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetFields

`func (o *LocationCreate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LocationCreate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LocationCreate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *LocationCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *LocationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetParentStorageId

`func (o *LocationCreate) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *LocationCreate) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *LocationCreate) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *LocationCreate) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### GetSchemaId

`func (o *LocationCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *LocationCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *LocationCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


