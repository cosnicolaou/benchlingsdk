# AssayRunSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**FieldDefinitions** | Pointer to [**[]SchemaFieldDefinitionsInner**](SchemaFieldDefinitionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**DerivedFrom** | Pointer to **NullableString** |  | [optional] 
**Organization** | Pointer to [**BaseAssaySchemaAllOfOrganization**](BaseAssaySchemaAllOfOrganization.md) |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**AutomationInputFileConfigs** | Pointer to [**[]AssayRunSchemaAllOfAutomationInputFileConfigs**](AssayRunSchemaAllOfAutomationInputFileConfigs.md) |  | [optional] 
**AutomationOutputFileConfigs** | Pointer to [**[]AssayRunSchemaAllOfAutomationInputFileConfigs**](AssayRunSchemaAllOfAutomationInputFileConfigs.md) |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Assay Run Schema was last modified | [optional] 

## Methods

### NewAssayRunSchema

`func NewAssayRunSchema() *AssayRunSchema`

NewAssayRunSchema instantiates a new AssayRunSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunSchemaWithDefaults

`func NewAssayRunSchemaWithDefaults() *AssayRunSchema`

NewAssayRunSchemaWithDefaults instantiates a new AssayRunSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *AssayRunSchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *AssayRunSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *AssayRunSchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *AssayRunSchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *AssayRunSchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *AssayRunSchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *AssayRunSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *AssayRunSchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *AssayRunSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *AssayRunSchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *AssayRunSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssayRunSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssayRunSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssayRunSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AssayRunSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssayRunSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssayRunSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssayRunSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AssayRunSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssayRunSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssayRunSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssayRunSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDerivedFrom

`func (o *AssayRunSchema) GetDerivedFrom() string`

GetDerivedFrom returns the DerivedFrom field if non-nil, zero value otherwise.

### GetDerivedFromOk

`func (o *AssayRunSchema) GetDerivedFromOk() (*string, bool)`

GetDerivedFromOk returns a tuple with the DerivedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFrom

`func (o *AssayRunSchema) SetDerivedFrom(v string)`

SetDerivedFrom sets DerivedFrom field to given value.

### HasDerivedFrom

`func (o *AssayRunSchema) HasDerivedFrom() bool`

HasDerivedFrom returns a boolean if a field has been set.

### SetDerivedFromNil

`func (o *AssayRunSchema) SetDerivedFromNil(b bool)`

 SetDerivedFromNil sets the value for DerivedFrom to be an explicit nil

### UnsetDerivedFrom
`func (o *AssayRunSchema) UnsetDerivedFrom()`

UnsetDerivedFrom ensures that no value is present for DerivedFrom, not even an explicit nil
### GetOrganization

`func (o *AssayRunSchema) GetOrganization() BaseAssaySchemaAllOfOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AssayRunSchema) GetOrganizationOk() (*BaseAssaySchemaAllOfOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AssayRunSchema) SetOrganization(v BaseAssaySchemaAllOfOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AssayRunSchema) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSystemName

`func (o *AssayRunSchema) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *AssayRunSchema) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *AssayRunSchema) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *AssayRunSchema) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetAutomationInputFileConfigs

`func (o *AssayRunSchema) GetAutomationInputFileConfigs() []AssayRunSchemaAllOfAutomationInputFileConfigs`

GetAutomationInputFileConfigs returns the AutomationInputFileConfigs field if non-nil, zero value otherwise.

### GetAutomationInputFileConfigsOk

`func (o *AssayRunSchema) GetAutomationInputFileConfigsOk() (*[]AssayRunSchemaAllOfAutomationInputFileConfigs, bool)`

GetAutomationInputFileConfigsOk returns a tuple with the AutomationInputFileConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationInputFileConfigs

`func (o *AssayRunSchema) SetAutomationInputFileConfigs(v []AssayRunSchemaAllOfAutomationInputFileConfigs)`

SetAutomationInputFileConfigs sets AutomationInputFileConfigs field to given value.

### HasAutomationInputFileConfigs

`func (o *AssayRunSchema) HasAutomationInputFileConfigs() bool`

HasAutomationInputFileConfigs returns a boolean if a field has been set.

### GetAutomationOutputFileConfigs

`func (o *AssayRunSchema) GetAutomationOutputFileConfigs() []AssayRunSchemaAllOfAutomationInputFileConfigs`

GetAutomationOutputFileConfigs returns the AutomationOutputFileConfigs field if non-nil, zero value otherwise.

### GetAutomationOutputFileConfigsOk

`func (o *AssayRunSchema) GetAutomationOutputFileConfigsOk() (*[]AssayRunSchemaAllOfAutomationInputFileConfigs, bool)`

GetAutomationOutputFileConfigsOk returns a tuple with the AutomationOutputFileConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationOutputFileConfigs

`func (o *AssayRunSchema) SetAutomationOutputFileConfigs(v []AssayRunSchemaAllOfAutomationInputFileConfigs)`

SetAutomationOutputFileConfigs sets AutomationOutputFileConfigs field to given value.

### HasAutomationOutputFileConfigs

`func (o *AssayRunSchema) HasAutomationOutputFileConfigs() bool`

HasAutomationOutputFileConfigs returns a boolean if a field has been set.

### GetModifiedAt

`func (o *AssayRunSchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AssayRunSchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AssayRunSchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AssayRunSchema) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


