# BaseAssaySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**FieldDefinitions** | Pointer to [**[]SchemaFieldDefinitionsInner**](SchemaFieldDefinitionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DerivedFrom** | Pointer to **NullableString** |  | [optional] 
**Organization** | Pointer to [**BaseAssaySchemaAllOfOrganization**](BaseAssaySchemaAllOfOrganization.md) |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseAssaySchema

`func NewBaseAssaySchema() *BaseAssaySchema`

NewBaseAssaySchema instantiates a new BaseAssaySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAssaySchemaWithDefaults

`func NewBaseAssaySchemaWithDefaults() *BaseAssaySchema`

NewBaseAssaySchemaWithDefaults instantiates a new BaseAssaySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *BaseAssaySchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *BaseAssaySchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *BaseAssaySchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *BaseAssaySchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *BaseAssaySchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *BaseAssaySchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *BaseAssaySchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *BaseAssaySchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *BaseAssaySchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *BaseAssaySchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *BaseAssaySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseAssaySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseAssaySchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseAssaySchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseAssaySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseAssaySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseAssaySchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseAssaySchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BaseAssaySchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseAssaySchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseAssaySchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseAssaySchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDerivedFrom

`func (o *BaseAssaySchema) GetDerivedFrom() string`

GetDerivedFrom returns the DerivedFrom field if non-nil, zero value otherwise.

### GetDerivedFromOk

`func (o *BaseAssaySchema) GetDerivedFromOk() (*string, bool)`

GetDerivedFromOk returns a tuple with the DerivedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFrom

`func (o *BaseAssaySchema) SetDerivedFrom(v string)`

SetDerivedFrom sets DerivedFrom field to given value.

### HasDerivedFrom

`func (o *BaseAssaySchema) HasDerivedFrom() bool`

HasDerivedFrom returns a boolean if a field has been set.

### SetDerivedFromNil

`func (o *BaseAssaySchema) SetDerivedFromNil(b bool)`

 SetDerivedFromNil sets the value for DerivedFrom to be an explicit nil

### UnsetDerivedFrom
`func (o *BaseAssaySchema) UnsetDerivedFrom()`

UnsetDerivedFrom ensures that no value is present for DerivedFrom, not even an explicit nil
### GetOrganization

`func (o *BaseAssaySchema) GetOrganization() BaseAssaySchemaAllOfOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BaseAssaySchema) GetOrganizationOk() (*BaseAssaySchemaAllOfOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BaseAssaySchema) SetOrganization(v BaseAssaySchemaAllOfOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BaseAssaySchema) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSystemName

`func (o *BaseAssaySchema) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *BaseAssaySchema) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *BaseAssaySchema) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *BaseAssaySchema) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


