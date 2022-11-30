# Box

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Barcode** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**EmptyContainers** | Pointer to **int32** | The number of containers in the box that have no contents. | [optional] [readonly] 
**EmptyPositions** | Pointer to **int32** | The number of empty positions for adding additional containers in the box. | [optional] [readonly] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FilledPositions** | Pointer to **int32** | The number of containers currently in the box. | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ParentStorageId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty3**](SchemaProperty3.md) |  | [optional] 
**Size** | Pointer to **int32** | The size of the box (i.e. how many containers it can store). | [optional] [readonly] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBox

`func NewBox() *Box`

NewBox instantiates a new Box object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxWithDefaults

`func NewBoxWithDefaults() *Box`

NewBoxWithDefaults instantiates a new Box object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *Box) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Box) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Box) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Box) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Box) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Box) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBarcode

`func (o *Box) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Box) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Box) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Box) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *Box) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *Box) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetCreatedAt

`func (o *Box) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Box) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Box) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Box) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Box) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Box) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Box) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Box) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEmptyContainers

`func (o *Box) GetEmptyContainers() int32`

GetEmptyContainers returns the EmptyContainers field if non-nil, zero value otherwise.

### GetEmptyContainersOk

`func (o *Box) GetEmptyContainersOk() (*int32, bool)`

GetEmptyContainersOk returns a tuple with the EmptyContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyContainers

`func (o *Box) SetEmptyContainers(v int32)`

SetEmptyContainers sets EmptyContainers field to given value.

### HasEmptyContainers

`func (o *Box) HasEmptyContainers() bool`

HasEmptyContainers returns a boolean if a field has been set.

### GetEmptyPositions

`func (o *Box) GetEmptyPositions() int32`

GetEmptyPositions returns the EmptyPositions field if non-nil, zero value otherwise.

### GetEmptyPositionsOk

`func (o *Box) GetEmptyPositionsOk() (*int32, bool)`

GetEmptyPositionsOk returns a tuple with the EmptyPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyPositions

`func (o *Box) SetEmptyPositions(v int32)`

SetEmptyPositions sets EmptyPositions field to given value.

### HasEmptyPositions

`func (o *Box) HasEmptyPositions() bool`

HasEmptyPositions returns a boolean if a field has been set.

### GetFields

`func (o *Box) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Box) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Box) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Box) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilledPositions

`func (o *Box) GetFilledPositions() int32`

GetFilledPositions returns the FilledPositions field if non-nil, zero value otherwise.

### GetFilledPositionsOk

`func (o *Box) GetFilledPositionsOk() (*int32, bool)`

GetFilledPositionsOk returns a tuple with the FilledPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledPositions

`func (o *Box) SetFilledPositions(v int32)`

SetFilledPositions sets FilledPositions field to given value.

### HasFilledPositions

`func (o *Box) HasFilledPositions() bool`

HasFilledPositions returns a boolean if a field has been set.

### GetId

`func (o *Box) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Box) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Box) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Box) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Box) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Box) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Box) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Box) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Box) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Box) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Box) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Box) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *Box) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *Box) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *Box) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *Box) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### SetParentStorageIdNil

`func (o *Box) SetParentStorageIdNil(b bool)`

 SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil

### UnsetParentStorageId
`func (o *Box) UnsetParentStorageId()`

UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
### GetProjectId

`func (o *Box) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Box) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Box) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Box) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Box) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Box) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSchema

`func (o *Box) GetSchema() SchemaProperty3`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Box) GetSchemaOk() (*SchemaProperty3, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Box) SetSchema(v SchemaProperty3)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Box) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Box) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Box) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetSize

`func (o *Box) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Box) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Box) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Box) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetWebURL

`func (o *Box) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Box) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Box) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Box) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


