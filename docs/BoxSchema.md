# BoxSchema

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
**Width** | Pointer to **float32** |  | [optional] 

## Methods

### NewBoxSchema

`func NewBoxSchema() *BoxSchema`

NewBoxSchema instantiates a new BoxSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoxSchemaWithDefaults

`func NewBoxSchemaWithDefaults() *BoxSchema`

NewBoxSchemaWithDefaults instantiates a new BoxSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *BoxSchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *BoxSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *BoxSchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *BoxSchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *BoxSchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *BoxSchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *BoxSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *BoxSchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *BoxSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *BoxSchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *BoxSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoxSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoxSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BoxSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoxSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoxSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoxSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoxSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BoxSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoxSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoxSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BoxSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrefix

`func (o *BoxSchema) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *BoxSchema) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *BoxSchema) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *BoxSchema) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegistryId

`func (o *BoxSchema) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *BoxSchema) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *BoxSchema) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *BoxSchema) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetContainerSchema

`func (o *BoxSchema) GetContainerSchema() BoxSchemaAllOfContainerSchema`

GetContainerSchema returns the ContainerSchema field if non-nil, zero value otherwise.

### GetContainerSchemaOk

`func (o *BoxSchema) GetContainerSchemaOk() (*BoxSchemaAllOfContainerSchema, bool)`

GetContainerSchemaOk returns a tuple with the ContainerSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSchema

`func (o *BoxSchema) SetContainerSchema(v BoxSchemaAllOfContainerSchema)`

SetContainerSchema sets ContainerSchema field to given value.

### HasContainerSchema

`func (o *BoxSchema) HasContainerSchema() bool`

HasContainerSchema returns a boolean if a field has been set.

### SetContainerSchemaNil

`func (o *BoxSchema) SetContainerSchemaNil(b bool)`

 SetContainerSchemaNil sets the value for ContainerSchema to be an explicit nil

### UnsetContainerSchema
`func (o *BoxSchema) UnsetContainerSchema()`

UnsetContainerSchema ensures that no value is present for ContainerSchema, not even an explicit nil
### GetHeight

`func (o *BoxSchema) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BoxSchema) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BoxSchema) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BoxSchema) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *BoxSchema) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BoxSchema) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BoxSchema) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *BoxSchema) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


