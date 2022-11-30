# FloatFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsMulti** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DecimalPrecision** | Pointer to **NullableFloat32** |  | [optional] 
**LegalTextDropdownId** | Pointer to **NullableString** |  | [optional] 
**NumericMax** | Pointer to **NullableFloat32** |  | [optional] 
**NumericMin** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewFloatFieldDefinition

`func NewFloatFieldDefinition() *FloatFieldDefinition`

NewFloatFieldDefinition instantiates a new FloatFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatFieldDefinitionWithDefaults

`func NewFloatFieldDefinitionWithDefaults() *FloatFieldDefinition`

NewFloatFieldDefinitionWithDefaults instantiates a new FloatFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *FloatFieldDefinition) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *FloatFieldDefinition) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *FloatFieldDefinition) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *FloatFieldDefinition) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *FloatFieldDefinition) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *FloatFieldDefinition) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetId

`func (o *FloatFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatFieldDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FloatFieldDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMulti

`func (o *FloatFieldDefinition) GetIsMulti() bool`

GetIsMulti returns the IsMulti field if non-nil, zero value otherwise.

### GetIsMultiOk

`func (o *FloatFieldDefinition) GetIsMultiOk() (*bool, bool)`

GetIsMultiOk returns a tuple with the IsMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulti

`func (o *FloatFieldDefinition) SetIsMulti(v bool)`

SetIsMulti sets IsMulti field to given value.

### HasIsMulti

`func (o *FloatFieldDefinition) HasIsMulti() bool`

HasIsMulti returns a boolean if a field has been set.

### GetIsRequired

`func (o *FloatFieldDefinition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *FloatFieldDefinition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *FloatFieldDefinition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *FloatFieldDefinition) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetName

`func (o *FloatFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloatFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloatFieldDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FloatFieldDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FloatFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FloatFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FloatFieldDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FloatFieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDecimalPrecision

`func (o *FloatFieldDefinition) GetDecimalPrecision() float32`

GetDecimalPrecision returns the DecimalPrecision field if non-nil, zero value otherwise.

### GetDecimalPrecisionOk

`func (o *FloatFieldDefinition) GetDecimalPrecisionOk() (*float32, bool)`

GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPrecision

`func (o *FloatFieldDefinition) SetDecimalPrecision(v float32)`

SetDecimalPrecision sets DecimalPrecision field to given value.

### HasDecimalPrecision

`func (o *FloatFieldDefinition) HasDecimalPrecision() bool`

HasDecimalPrecision returns a boolean if a field has been set.

### SetDecimalPrecisionNil

`func (o *FloatFieldDefinition) SetDecimalPrecisionNil(b bool)`

 SetDecimalPrecisionNil sets the value for DecimalPrecision to be an explicit nil

### UnsetDecimalPrecision
`func (o *FloatFieldDefinition) UnsetDecimalPrecision()`

UnsetDecimalPrecision ensures that no value is present for DecimalPrecision, not even an explicit nil
### GetLegalTextDropdownId

`func (o *FloatFieldDefinition) GetLegalTextDropdownId() string`

GetLegalTextDropdownId returns the LegalTextDropdownId field if non-nil, zero value otherwise.

### GetLegalTextDropdownIdOk

`func (o *FloatFieldDefinition) GetLegalTextDropdownIdOk() (*string, bool)`

GetLegalTextDropdownIdOk returns a tuple with the LegalTextDropdownId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalTextDropdownId

`func (o *FloatFieldDefinition) SetLegalTextDropdownId(v string)`

SetLegalTextDropdownId sets LegalTextDropdownId field to given value.

### HasLegalTextDropdownId

`func (o *FloatFieldDefinition) HasLegalTextDropdownId() bool`

HasLegalTextDropdownId returns a boolean if a field has been set.

### SetLegalTextDropdownIdNil

`func (o *FloatFieldDefinition) SetLegalTextDropdownIdNil(b bool)`

 SetLegalTextDropdownIdNil sets the value for LegalTextDropdownId to be an explicit nil

### UnsetLegalTextDropdownId
`func (o *FloatFieldDefinition) UnsetLegalTextDropdownId()`

UnsetLegalTextDropdownId ensures that no value is present for LegalTextDropdownId, not even an explicit nil
### GetNumericMax

`func (o *FloatFieldDefinition) GetNumericMax() float32`

GetNumericMax returns the NumericMax field if non-nil, zero value otherwise.

### GetNumericMaxOk

`func (o *FloatFieldDefinition) GetNumericMaxOk() (*float32, bool)`

GetNumericMaxOk returns a tuple with the NumericMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericMax

`func (o *FloatFieldDefinition) SetNumericMax(v float32)`

SetNumericMax sets NumericMax field to given value.

### HasNumericMax

`func (o *FloatFieldDefinition) HasNumericMax() bool`

HasNumericMax returns a boolean if a field has been set.

### SetNumericMaxNil

`func (o *FloatFieldDefinition) SetNumericMaxNil(b bool)`

 SetNumericMaxNil sets the value for NumericMax to be an explicit nil

### UnsetNumericMax
`func (o *FloatFieldDefinition) UnsetNumericMax()`

UnsetNumericMax ensures that no value is present for NumericMax, not even an explicit nil
### GetNumericMin

`func (o *FloatFieldDefinition) GetNumericMin() float32`

GetNumericMin returns the NumericMin field if non-nil, zero value otherwise.

### GetNumericMinOk

`func (o *FloatFieldDefinition) GetNumericMinOk() (*float32, bool)`

GetNumericMinOk returns a tuple with the NumericMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericMin

`func (o *FloatFieldDefinition) SetNumericMin(v float32)`

SetNumericMin sets NumericMin field to given value.

### HasNumericMin

`func (o *FloatFieldDefinition) HasNumericMin() bool`

HasNumericMin returns a boolean if a field has been set.

### SetNumericMinNil

`func (o *FloatFieldDefinition) SetNumericMinNil(b bool)`

 SetNumericMinNil sets the value for NumericMin to be an explicit nil

### UnsetNumericMin
`func (o *FloatFieldDefinition) UnsetNumericMin()`

UnsetNumericMin ensures that no value is present for NumericMin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


