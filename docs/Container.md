# Container

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

## Methods

### NewContainer

`func NewContainer() *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *Container) GetArchiveRecord() ContainerArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Container) GetArchiveRecordOk() (*ContainerArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Container) SetArchiveRecord(v ContainerArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Container) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Container) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Container) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBarcode

`func (o *Container) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Container) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Container) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Container) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *Container) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *Container) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetCheckoutRecord

`func (o *Container) GetCheckoutRecord() CheckoutRecord`

GetCheckoutRecord returns the CheckoutRecord field if non-nil, zero value otherwise.

### GetCheckoutRecordOk

`func (o *Container) GetCheckoutRecordOk() (*CheckoutRecord, bool)`

GetCheckoutRecordOk returns a tuple with the CheckoutRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutRecord

`func (o *Container) SetCheckoutRecord(v CheckoutRecord)`

SetCheckoutRecord sets CheckoutRecord field to given value.

### HasCheckoutRecord

`func (o *Container) HasCheckoutRecord() bool`

HasCheckoutRecord returns a boolean if a field has been set.

### GetContents

`func (o *Container) GetContents() []ContainerContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Container) GetContentsOk() (*[]ContainerContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Container) SetContents(v []ContainerContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *Container) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Container) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Container) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Container) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Container) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Container) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Container) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Container) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Container) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFields

`func (o *Container) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Container) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Container) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Container) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *Container) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Container) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Container) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Container) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Container) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Container) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Container) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Container) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Container) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Container) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *Container) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *Container) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *Container) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *Container) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### SetParentStorageIdNil

`func (o *Container) SetParentStorageIdNil(b bool)`

 SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil

### UnsetParentStorageId
`func (o *Container) UnsetParentStorageId()`

UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
### GetParentStorageSchema

`func (o *Container) GetParentStorageSchema() ContainerParentStorageSchema`

GetParentStorageSchema returns the ParentStorageSchema field if non-nil, zero value otherwise.

### GetParentStorageSchemaOk

`func (o *Container) GetParentStorageSchemaOk() (*ContainerParentStorageSchema, bool)`

GetParentStorageSchemaOk returns a tuple with the ParentStorageSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageSchema

`func (o *Container) SetParentStorageSchema(v ContainerParentStorageSchema)`

SetParentStorageSchema sets ParentStorageSchema field to given value.

### HasParentStorageSchema

`func (o *Container) HasParentStorageSchema() bool`

HasParentStorageSchema returns a boolean if a field has been set.

### SetParentStorageSchemaNil

`func (o *Container) SetParentStorageSchemaNil(b bool)`

 SetParentStorageSchemaNil sets the value for ParentStorageSchema to be an explicit nil

### UnsetParentStorageSchema
`func (o *Container) UnsetParentStorageSchema()`

UnsetParentStorageSchema ensures that no value is present for ParentStorageSchema, not even an explicit nil
### GetProjectId

`func (o *Container) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Container) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Container) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Container) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Container) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Container) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetQuantity

`func (o *Container) GetQuantity() ContainerQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Container) GetQuantityOk() (*ContainerQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Container) SetQuantity(v ContainerQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Container) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSchema

`func (o *Container) GetSchema() ContainerParentStorageSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Container) GetSchemaOk() (*ContainerParentStorageSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Container) SetSchema(v ContainerParentStorageSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Container) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Container) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Container) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetVolume

`func (o *Container) GetVolume() DeprecatedContainerVolumeForResponse`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Container) GetVolumeOk() (*DeprecatedContainerVolumeForResponse, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Container) SetVolume(v DeprecatedContainerVolumeForResponse)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Container) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWebURL

`func (o *Container) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Container) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Container) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Container) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


