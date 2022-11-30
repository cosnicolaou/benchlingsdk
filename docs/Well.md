# Well

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableContainerArchiveRecord**](ContainerArchiveRecord.md) |  | [optional] 
**Barcode** | Pointer to **string** | Barcode of the well | [optional] 
**CheckoutRecord** | Pointer to [**CheckoutRecord**](CheckoutRecord.md) |  | [optional] [readonly] 
**Contents** | Pointer to [**[]ContainerContent**](ContainerContent.md) | Array of well contents, each with a batch and concentration | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the well was created | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the well | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the well was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the well, defaults to barcode if name is not provided. | [optional] 
**ParentStorageId** | Pointer to **NullableString** | ID of containing parent storage, a plate well with a coordinate (e.g. plt_2bAks9dx:a2). | [optional] 
**ParentStorageSchema** | Pointer to [**SchemaSummary**](SchemaSummary.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** | ID of the project if set | [optional] 
**Quantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**Schema** | Pointer to [**NullableAaSequenceSchema**](AaSequenceSchema.md) |  | [optional] 
**Volume** | Pointer to [**DeprecatedContainerVolumeForResponse**](DeprecatedContainerVolumeForResponse.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewWell

`func NewWell() *Well`

NewWell instantiates a new Well object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellWithDefaults

`func NewWellWithDefaults() *Well`

NewWellWithDefaults instantiates a new Well object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *Well) GetArchiveRecord() ContainerArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Well) GetArchiveRecordOk() (*ContainerArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Well) SetArchiveRecord(v ContainerArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Well) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Well) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Well) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBarcode

`func (o *Well) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Well) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Well) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Well) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetCheckoutRecord

`func (o *Well) GetCheckoutRecord() CheckoutRecord`

GetCheckoutRecord returns the CheckoutRecord field if non-nil, zero value otherwise.

### GetCheckoutRecordOk

`func (o *Well) GetCheckoutRecordOk() (*CheckoutRecord, bool)`

GetCheckoutRecordOk returns a tuple with the CheckoutRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutRecord

`func (o *Well) SetCheckoutRecord(v CheckoutRecord)`

SetCheckoutRecord sets CheckoutRecord field to given value.

### HasCheckoutRecord

`func (o *Well) HasCheckoutRecord() bool`

HasCheckoutRecord returns a boolean if a field has been set.

### GetContents

`func (o *Well) GetContents() []ContainerContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Well) GetContentsOk() (*[]ContainerContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Well) SetContents(v []ContainerContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *Well) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Well) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Well) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Well) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Well) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Well) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Well) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Well) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Well) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFields

`func (o *Well) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Well) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Well) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Well) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *Well) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Well) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Well) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Well) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Well) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Well) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Well) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Well) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Well) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Well) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Well) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Well) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *Well) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *Well) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *Well) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *Well) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.

### SetParentStorageIdNil

`func (o *Well) SetParentStorageIdNil(b bool)`

 SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil

### UnsetParentStorageId
`func (o *Well) UnsetParentStorageId()`

UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
### GetParentStorageSchema

`func (o *Well) GetParentStorageSchema() SchemaSummary`

GetParentStorageSchema returns the ParentStorageSchema field if non-nil, zero value otherwise.

### GetParentStorageSchemaOk

`func (o *Well) GetParentStorageSchemaOk() (*SchemaSummary, bool)`

GetParentStorageSchemaOk returns a tuple with the ParentStorageSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageSchema

`func (o *Well) SetParentStorageSchema(v SchemaSummary)`

SetParentStorageSchema sets ParentStorageSchema field to given value.

### HasParentStorageSchema

`func (o *Well) HasParentStorageSchema() bool`

HasParentStorageSchema returns a boolean if a field has been set.

### GetProjectId

`func (o *Well) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Well) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Well) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Well) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Well) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Well) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetQuantity

`func (o *Well) GetQuantity() ContainerQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Well) GetQuantityOk() (*ContainerQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Well) SetQuantity(v ContainerQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Well) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSchema

`func (o *Well) GetSchema() AaSequenceSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Well) GetSchemaOk() (*AaSequenceSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Well) SetSchema(v AaSequenceSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Well) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Well) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Well) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetVolume

`func (o *Well) GetVolume() DeprecatedContainerVolumeForResponse`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Well) GetVolumeOk() (*DeprecatedContainerVolumeForResponse, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Well) SetVolume(v DeprecatedContainerVolumeForResponse)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Well) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWebURL

`func (o *Well) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Well) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Well) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Well) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


