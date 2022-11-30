# Plate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Barcode** | Pointer to **NullableString** | Barcode of the plate | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the container was created | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the plate | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the plate was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the plate, defaults to barcode if name is not provided. | [optional] 
**ParentStorageId** | Pointer to **NullableString** | ID of containing parent storage (e.g. loc_k2lNspzS). | [optional] 
**ProjectId** | Pointer to **NullableString** | ID of the project if set | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty3**](SchemaProperty3.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 
**Wells** | Pointer to [**map[string]PlateWellsValue**](PlateWellsValue.md) | Well contents of the plate, keyed by position string (eg. \&quot;A1\&quot;). | [optional] 

## Methods

### NewPlate

`func NewPlate() *Plate`

NewPlate instantiates a new Plate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlateWithDefaults

`func NewPlateWithDefaults() *Plate`

NewPlateWithDefaults instantiates a new Plate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *Plate) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Plate) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Plate) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Plate) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Plate) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Plate) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBarcode

`func (o *Plate) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Plate) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Plate) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Plate) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *Plate) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *Plate) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetCreatedAt

`func (o *Plate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Plate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Plate) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Plate) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Plate) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Plate) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFields

`func (o *Plate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Plate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Plate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Plate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *Plate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Plate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Plate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Plate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Plate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Plate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Plate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *Plate) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *Plate) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *Plate) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *Plate) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### SetParentStorageIdNil

`func (o *Plate) SetParentStorageIdNil(b bool)`

 SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil

### UnsetParentStorageId
`func (o *Plate) UnsetParentStorageId()`

UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
### GetProjectId

`func (o *Plate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Plate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Plate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Plate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Plate) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Plate) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSchema

`func (o *Plate) GetSchema() SchemaProperty3`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Plate) GetSchemaOk() (*SchemaProperty3, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Plate) SetSchema(v SchemaProperty3)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Plate) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Plate) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Plate) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetType

`func (o *Plate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Plate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Plate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Plate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebURL

`func (o *Plate) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Plate) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Plate) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Plate) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetWells

`func (o *Plate) GetWells() map[string]PlateWellsValue`

GetWells returns the Wells field if non-nil, zero value otherwise.

### GetWellsOk

`func (o *Plate) GetWellsOk() (*map[string]PlateWellsValue, bool)`

GetWellsOk returns a tuple with the Wells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWells

`func (o *Plate) SetWells(v map[string]PlateWellsValue)`

SetWells sets Wells field to given value.

### HasWells

`func (o *Plate) HasWells() bool`

HasWells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


