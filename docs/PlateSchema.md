# PlateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**FieldDefinitions** | Pointer to [**[]SchemaFieldDefinitionsInner**](SchemaFieldDefinitionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Prefix** | Pointer to **string** |  | [optional] 
**RegistryId** | Pointer to **string** |  | [optional] 
**ContainerSchema** | Pointer to [**NullableBoxSchemaAllOfContainerSchema**](BoxSchemaAllOfContainerSchema.md) |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**PlateType** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 

## Methods

### NewPlateSchema

`func NewPlateSchema() *PlateSchema`

NewPlateSchema instantiates a new PlateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlateSchemaWithDefaults

`func NewPlateSchemaWithDefaults() *PlateSchema`

NewPlateSchemaWithDefaults instantiates a new PlateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *PlateSchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *PlateSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *PlateSchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *PlateSchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *PlateSchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *PlateSchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *PlateSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *PlateSchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *PlateSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *PlateSchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *PlateSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlateSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlateSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlateSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PlateSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlateSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlateSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlateSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PlateSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlateSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlateSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlateSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrefix

`func (o *PlateSchema) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PlateSchema) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PlateSchema) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PlateSchema) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegistryId

`func (o *PlateSchema) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *PlateSchema) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *PlateSchema) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *PlateSchema) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetContainerSchema

`func (o *PlateSchema) GetContainerSchema() BoxSchemaAllOfContainerSchema`

GetContainerSchema returns the ContainerSchema field if non-nil, zero value otherwise.

### GetContainerSchemaOk

`func (o *PlateSchema) GetContainerSchemaOk() (*BoxSchemaAllOfContainerSchema, bool)`

GetContainerSchemaOk returns a tuple with the ContainerSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSchema

`func (o *PlateSchema) SetContainerSchema(v BoxSchemaAllOfContainerSchema)`

SetContainerSchema sets ContainerSchema field to given value.

### HasContainerSchema

`func (o *PlateSchema) HasContainerSchema() bool`

HasContainerSchema returns a boolean if a field has been set.

### SetContainerSchemaNil

`func (o *PlateSchema) SetContainerSchemaNil(b bool)`

 SetContainerSchemaNil sets the value for ContainerSchema to be an explicit nil

### UnsetContainerSchema
`func (o *PlateSchema) UnsetContainerSchema()`

UnsetContainerSchema ensures that no value is present for ContainerSchema, not even an explicit nil
### GetHeight

`func (o *PlateSchema) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PlateSchema) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PlateSchema) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PlateSchema) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetPlateType

`func (o *PlateSchema) GetPlateType() string`

GetPlateType returns the PlateType field if non-nil, zero value otherwise.

### GetPlateTypeOk

`func (o *PlateSchema) GetPlateTypeOk() (*string, bool)`

GetPlateTypeOk returns a tuple with the PlateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlateType

`func (o *PlateSchema) SetPlateType(v string)`

SetPlateType sets PlateType field to given value.

### HasPlateType

`func (o *PlateSchema) HasPlateType() bool`

HasPlateType returns a boolean if a field has been set.

### GetWidth

`func (o *PlateSchema) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PlateSchema) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PlateSchema) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PlateSchema) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


