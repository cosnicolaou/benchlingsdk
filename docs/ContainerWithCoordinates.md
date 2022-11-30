# ContainerWithCoordinates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableContainerArchiveRecord**](ContainerArchiveRecord.md) |  | [optional] 
**Barcode** | Pointer to **NullableString** |  | [optional] 
**CheckoutRecord** | Pointer to [**CheckoutRecord**](CheckoutRecord.md) |  | [optional] [readonly] 
**Contents** | Pointer to [**[]ContainerContent**](ContainerContent.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] [readonly] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ParentStorageId** | Pointer to **NullableString** |  | [optional] 
**ParentStorageSchema** | Pointer to [**NullableContainerParentStorageSchema**](ContainerParentStorageSchema.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**Schema** | Pointer to [**NullableContainerParentStorageSchema**](ContainerParentStorageSchema.md) |  | [optional] 
**Volume** | Pointer to [**DeprecatedContainerVolumeForResponse**](DeprecatedContainerVolumeForResponse.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 
**GridNumber** | Pointer to **float32** |  | [optional] 
**GridPosition** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerWithCoordinates

`func NewContainerWithCoordinates() *ContainerWithCoordinates`

NewContainerWithCoordinates instantiates a new ContainerWithCoordinates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithCoordinatesWithDefaults

`func NewContainerWithCoordinatesWithDefaults() *ContainerWithCoordinates`

NewContainerWithCoordinatesWithDefaults instantiates a new ContainerWithCoordinates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *ContainerWithCoordinates) GetArchiveRecord() ContainerArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *ContainerWithCoordinates) GetArchiveRecordOk() (*ContainerArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *ContainerWithCoordinates) SetArchiveRecord(v ContainerArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *ContainerWithCoordinates) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *ContainerWithCoordinates) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *ContainerWithCoordinates) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBarcode

`func (o *ContainerWithCoordinates) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ContainerWithCoordinates) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ContainerWithCoordinates) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ContainerWithCoordinates) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *ContainerWithCoordinates) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *ContainerWithCoordinates) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetCheckoutRecord

`func (o *ContainerWithCoordinates) GetCheckoutRecord() CheckoutRecord`

GetCheckoutRecord returns the CheckoutRecord field if non-nil, zero value otherwise.

### GetCheckoutRecordOk

`func (o *ContainerWithCoordinates) GetCheckoutRecordOk() (*CheckoutRecord, bool)`

GetCheckoutRecordOk returns a tuple with the CheckoutRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutRecord

`func (o *ContainerWithCoordinates) SetCheckoutRecord(v CheckoutRecord)`

SetCheckoutRecord sets CheckoutRecord field to given value.

### HasCheckoutRecord

`func (o *ContainerWithCoordinates) HasCheckoutRecord() bool`

HasCheckoutRecord returns a boolean if a field has been set.

### GetContents

`func (o *ContainerWithCoordinates) GetContents() []ContainerContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ContainerWithCoordinates) GetContentsOk() (*[]ContainerContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ContainerWithCoordinates) SetContents(v []ContainerContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ContainerWithCoordinates) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContainerWithCoordinates) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContainerWithCoordinates) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContainerWithCoordinates) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContainerWithCoordinates) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *ContainerWithCoordinates) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ContainerWithCoordinates) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ContainerWithCoordinates) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ContainerWithCoordinates) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFields

`func (o *ContainerWithCoordinates) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContainerWithCoordinates) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContainerWithCoordinates) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ContainerWithCoordinates) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *ContainerWithCoordinates) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerWithCoordinates) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerWithCoordinates) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerWithCoordinates) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ContainerWithCoordinates) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ContainerWithCoordinates) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ContainerWithCoordinates) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ContainerWithCoordinates) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *ContainerWithCoordinates) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerWithCoordinates) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerWithCoordinates) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerWithCoordinates) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *ContainerWithCoordinates) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *ContainerWithCoordinates) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *ContainerWithCoordinates) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *ContainerWithCoordinates) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### SetParentStorageIdNil

`func (o *ContainerWithCoordinates) SetParentStorageIdNil(b bool)`

 SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil

### UnsetParentStorageId
`func (o *ContainerWithCoordinates) UnsetParentStorageId()`

UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
### GetParentStorageSchema

`func (o *ContainerWithCoordinates) GetParentStorageSchema() ContainerParentStorageSchema`

GetParentStorageSchema returns the ParentStorageSchema field if non-nil, zero value otherwise.

### GetParentStorageSchemaOk

`func (o *ContainerWithCoordinates) GetParentStorageSchemaOk() (*ContainerParentStorageSchema, bool)`

GetParentStorageSchemaOk returns a tuple with the ParentStorageSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageSchema

`func (o *ContainerWithCoordinates) SetParentStorageSchema(v ContainerParentStorageSchema)`

SetParentStorageSchema sets ParentStorageSchema field to given value.

### HasParentStorageSchema

`func (o *ContainerWithCoordinates) HasParentStorageSchema() bool`

HasParentStorageSchema returns a boolean if a field has been set.

### SetParentStorageSchemaNil

`func (o *ContainerWithCoordinates) SetParentStorageSchemaNil(b bool)`

 SetParentStorageSchemaNil sets the value for ParentStorageSchema to be an explicit nil

### UnsetParentStorageSchema
`func (o *ContainerWithCoordinates) UnsetParentStorageSchema()`

UnsetParentStorageSchema ensures that no value is present for ParentStorageSchema, not even an explicit nil
### GetProjectId

`func (o *ContainerWithCoordinates) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ContainerWithCoordinates) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ContainerWithCoordinates) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ContainerWithCoordinates) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ContainerWithCoordinates) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ContainerWithCoordinates) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetQuantity

`func (o *ContainerWithCoordinates) GetQuantity() ContainerQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ContainerWithCoordinates) GetQuantityOk() (*ContainerQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ContainerWithCoordinates) SetQuantity(v ContainerQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ContainerWithCoordinates) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSchema

`func (o *ContainerWithCoordinates) GetSchema() ContainerParentStorageSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ContainerWithCoordinates) GetSchemaOk() (*ContainerParentStorageSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ContainerWithCoordinates) SetSchema(v ContainerParentStorageSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ContainerWithCoordinates) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ContainerWithCoordinates) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ContainerWithCoordinates) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetVolume

`func (o *ContainerWithCoordinates) GetVolume() DeprecatedContainerVolumeForResponse`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ContainerWithCoordinates) GetVolumeOk() (*DeprecatedContainerVolumeForResponse, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ContainerWithCoordinates) SetVolume(v DeprecatedContainerVolumeForResponse)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ContainerWithCoordinates) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWebURL

`func (o *ContainerWithCoordinates) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *ContainerWithCoordinates) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *ContainerWithCoordinates) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *ContainerWithCoordinates) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetGridNumber

`func (o *ContainerWithCoordinates) GetGridNumber() float32`

GetGridNumber returns the GridNumber field if non-nil, zero value otherwise.

### GetGridNumberOk

`func (o *ContainerWithCoordinates) GetGridNumberOk() (*float32, bool)`

GetGridNumberOk returns a tuple with the GridNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNumber

`func (o *ContainerWithCoordinates) SetGridNumber(v float32)`

SetGridNumber sets GridNumber field to given value.

### HasGridNumber

`func (o *ContainerWithCoordinates) HasGridNumber() bool`

HasGridNumber returns a boolean if a field has been set.

### GetGridPosition

`func (o *ContainerWithCoordinates) GetGridPosition() string`

GetGridPosition returns the GridPosition field if non-nil, zero value otherwise.

### GetGridPositionOk

`func (o *ContainerWithCoordinates) GetGridPositionOk() (*string, bool)`

GetGridPositionOk returns a tuple with the GridPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPosition

`func (o *ContainerWithCoordinates) SetGridPosition(v string)`

SetGridPosition sets GridPosition field to given value.

### HasGridPosition

`func (o *ContainerWithCoordinates) HasGridPosition() bool`

HasGridPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


