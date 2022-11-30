# RegistrySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**FieldDefinitions** | Pointer to [**[]SchemaFieldDefinitionsInner**](SchemaFieldDefinitionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**RegistryId** | Pointer to **string** |  | [optional] 

## Methods

### NewRegistrySchema

`func NewRegistrySchema() *RegistrySchema`

NewRegistrySchema instantiates a new RegistrySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrySchemaWithDefaults

`func NewRegistrySchemaWithDefaults() *RegistrySchema`

NewRegistrySchemaWithDefaults instantiates a new RegistrySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *RegistrySchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *RegistrySchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *RegistrySchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *RegistrySchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *RegistrySchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *RegistrySchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *RegistrySchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *RegistrySchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *RegistrySchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *RegistrySchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *RegistrySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrySchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistrySchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RegistrySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistrySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistrySchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistrySchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RegistrySchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistrySchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistrySchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistrySchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrefix

`func (o *RegistrySchema) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RegistrySchema) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RegistrySchema) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RegistrySchema) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegistryId

`func (o *RegistrySchema) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *RegistrySchema) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *RegistrySchema) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *RegistrySchema) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


