# SchemaFieldDefinitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsMulti** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**NumericMax** | Pointer to **NullableFloat32** |  | [optional] 
**NumericMin** | Pointer to **NullableFloat32** |  | [optional] 
**DecimalPrecision** | Pointer to **NullableFloat32** |  | [optional] 
**LegalTextDropdownId** | Pointer to **NullableString** |  | [optional] 
**DropdownId** | Pointer to **NullableString** |  | [optional] 
**SchemaId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSchemaFieldDefinitionsInner

`func NewSchemaFieldDefinitionsInner() *SchemaFieldDefinitionsInner`

NewSchemaFieldDefinitionsInner instantiates a new SchemaFieldDefinitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaFieldDefinitionsInnerWithDefaults

`func NewSchemaFieldDefinitionsInnerWithDefaults() *SchemaFieldDefinitionsInner`

NewSchemaFieldDefinitionsInnerWithDefaults instantiates a new SchemaFieldDefinitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *SchemaFieldDefinitionsInner) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *SchemaFieldDefinitionsInner) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *SchemaFieldDefinitionsInner) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *SchemaFieldDefinitionsInner) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *SchemaFieldDefinitionsInner) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *SchemaFieldDefinitionsInner) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetId

`func (o *SchemaFieldDefinitionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaFieldDefinitionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaFieldDefinitionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaFieldDefinitionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMulti

`func (o *SchemaFieldDefinitionsInner) GetIsMulti() bool`

GetIsMulti returns the IsMulti field if non-nil, zero value otherwise.

### GetIsMultiOk

`func (o *SchemaFieldDefinitionsInner) GetIsMultiOk() (*bool, bool)`

GetIsMultiOk returns a tuple with the IsMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulti

`func (o *SchemaFieldDefinitionsInner) SetIsMulti(v bool)`

SetIsMulti sets IsMulti field to given value.

### HasIsMulti

`func (o *SchemaFieldDefinitionsInner) HasIsMulti() bool`

HasIsMulti returns a boolean if a field has been set.

### GetIsRequired

`func (o *SchemaFieldDefinitionsInner) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *SchemaFieldDefinitionsInner) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *SchemaFieldDefinitionsInner) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *SchemaFieldDefinitionsInner) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetName

`func (o *SchemaFieldDefinitionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaFieldDefinitionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaFieldDefinitionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaFieldDefinitionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SchemaFieldDefinitionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaFieldDefinitionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaFieldDefinitionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SchemaFieldDefinitionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumericMax

`func (o *SchemaFieldDefinitionsInner) GetNumericMax() float32`

GetNumericMax returns the NumericMax field if non-nil, zero value otherwise.

### GetNumericMaxOk

`func (o *SchemaFieldDefinitionsInner) GetNumericMaxOk() (*float32, bool)`

GetNumericMaxOk returns a tuple with the NumericMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericMax

`func (o *SchemaFieldDefinitionsInner) SetNumericMax(v float32)`

SetNumericMax sets NumericMax field to given value.

### HasNumericMax

`func (o *SchemaFieldDefinitionsInner) HasNumericMax() bool`

HasNumericMax returns a boolean if a field has been set.

### SetNumericMaxNil

`func (o *SchemaFieldDefinitionsInner) SetNumericMaxNil(b bool)`

 SetNumericMaxNil sets the value for NumericMax to be an explicit nil

### UnsetNumericMax
`func (o *SchemaFieldDefinitionsInner) UnsetNumericMax()`

UnsetNumericMax ensures that no value is present for NumericMax, not even an explicit nil
### GetNumericMin

`func (o *SchemaFieldDefinitionsInner) GetNumericMin() float32`

GetNumericMin returns the NumericMin field if non-nil, zero value otherwise.

### GetNumericMinOk

`func (o *SchemaFieldDefinitionsInner) GetNumericMinOk() (*float32, bool)`

GetNumericMinOk returns a tuple with the NumericMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericMin

`func (o *SchemaFieldDefinitionsInner) SetNumericMin(v float32)`

SetNumericMin sets NumericMin field to given value.

### HasNumericMin

`func (o *SchemaFieldDefinitionsInner) HasNumericMin() bool`

HasNumericMin returns a boolean if a field has been set.

### SetNumericMinNil

`func (o *SchemaFieldDefinitionsInner) SetNumericMinNil(b bool)`

 SetNumericMinNil sets the value for NumericMin to be an explicit nil

### UnsetNumericMin
`func (o *SchemaFieldDefinitionsInner) UnsetNumericMin()`

UnsetNumericMin ensures that no value is present for NumericMin, not even an explicit nil
### GetDecimalPrecision

`func (o *SchemaFieldDefinitionsInner) GetDecimalPrecision() float32`

GetDecimalPrecision returns the DecimalPrecision field if non-nil, zero value otherwise.

### GetDecimalPrecisionOk

`func (o *SchemaFieldDefinitionsInner) GetDecimalPrecisionOk() (*float32, bool)`

GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPrecision

`func (o *SchemaFieldDefinitionsInner) SetDecimalPrecision(v float32)`

SetDecimalPrecision sets DecimalPrecision field to given value.

### HasDecimalPrecision

`func (o *SchemaFieldDefinitionsInner) HasDecimalPrecision() bool`

HasDecimalPrecision returns a boolean if a field has been set.

### SetDecimalPrecisionNil

`func (o *SchemaFieldDefinitionsInner) SetDecimalPrecisionNil(b bool)`

 SetDecimalPrecisionNil sets the value for DecimalPrecision to be an explicit nil

### UnsetDecimalPrecision
`func (o *SchemaFieldDefinitionsInner) UnsetDecimalPrecision()`

UnsetDecimalPrecision ensures that no value is present for DecimalPrecision, not even an explicit nil
### GetLegalTextDropdownId

`func (o *SchemaFieldDefinitionsInner) GetLegalTextDropdownId() string`

GetLegalTextDropdownId returns the LegalTextDropdownId field if non-nil, zero value otherwise.

### GetLegalTextDropdownIdOk

`func (o *SchemaFieldDefinitionsInner) GetLegalTextDropdownIdOk() (*string, bool)`

GetLegalTextDropdownIdOk returns a tuple with the LegalTextDropdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalTextDropdownId

`func (o *SchemaFieldDefinitionsInner) SetLegalTextDropdownId(v string)`

SetLegalTextDropdownId sets LegalTextDropdownId field to given value.

### HasLegalTextDropdownId

`func (o *SchemaFieldDefinitionsInner) HasLegalTextDropdownId() bool`

HasLegalTextDropdownId returns a boolean if a field has been set.

### SetLegalTextDropdownIdNil

`func (o *SchemaFieldDefinitionsInner) SetLegalTextDropdownIdNil(b bool)`

 SetLegalTextDropdownIdNil sets the value for LegalTextDropdownId to be an explicit nil

### UnsetLegalTextDropdownId
`func (o *SchemaFieldDefinitionsInner) UnsetLegalTextDropdownId()`

UnsetLegalTextDropdownId ensures that no value is present for LegalTextDropdownId, not even an explicit nil
### GetDropdownId

`func (o *SchemaFieldDefinitionsInner) GetDropdownId() string`

GetDropdownId returns the DropdownId field if non-nil, zero value otherwise.

### GetDropdownIdOk

`func (o *SchemaFieldDefinitionsInner) GetDropdownIdOk() (*string, bool)`

GetDropdownIdOk returns a tuple with the DropdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropdownId

`func (o *SchemaFieldDefinitionsInner) SetDropdownId(v string)`

SetDropdownId sets DropdownId field to given value.

### HasDropdownId

`func (o *SchemaFieldDefinitionsInner) HasDropdownId() bool`

HasDropdownId returns a boolean if a field has been set.

### SetDropdownIdNil

`func (o *SchemaFieldDefinitionsInner) SetDropdownIdNil(b bool)`

 SetDropdownIdNil sets the value for DropdownId to be an explicit nil

### UnsetDropdownId
`func (o *SchemaFieldDefinitionsInner) UnsetDropdownId()`

UnsetDropdownId ensures that no value is present for DropdownId, not even an explicit nil
### GetSchemaId

`func (o *SchemaFieldDefinitionsInner) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SchemaFieldDefinitionsInner) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SchemaFieldDefinitionsInner) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *SchemaFieldDefinitionsInner) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### SetSchemaIdNil

`func (o *SchemaFieldDefinitionsInner) SetSchemaIdNil(b bool)`

 SetSchemaIdNil sets the value for SchemaId to be an explicit nil

### UnsetSchemaId
`func (o *SchemaFieldDefinitionsInner) UnsetSchemaId()`

UnsetSchemaId ensures that no value is present for SchemaId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


