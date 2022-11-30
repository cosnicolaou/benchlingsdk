# IntegerFieldDefinition

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

## Methods

### NewIntegerFieldDefinition

`func NewIntegerFieldDefinition() *IntegerFieldDefinition`

NewIntegerFieldDefinition instantiates a new IntegerFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegerFieldDefinitionWithDefaults

`func NewIntegerFieldDefinitionWithDefaults() *IntegerFieldDefinition`

NewIntegerFieldDefinitionWithDefaults instantiates a new IntegerFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *IntegerFieldDefinition) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *IntegerFieldDefinition) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *IntegerFieldDefinition) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *IntegerFieldDefinition) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *IntegerFieldDefinition) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *IntegerFieldDefinition) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetId

`func (o *IntegerFieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegerFieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegerFieldDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegerFieldDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMulti

`func (o *IntegerFieldDefinition) GetIsMulti() bool`

GetIsMulti returns the IsMulti field if non-nil, zero value otherwise.

### GetIsMultiOk

`func (o *IntegerFieldDefinition) GetIsMultiOk() (*bool, bool)`

GetIsMultiOk returns a tuple with the IsMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulti

`func (o *IntegerFieldDefinition) SetIsMulti(v bool)`

SetIsMulti sets IsMulti field to given value.

### HasIsMulti

`func (o *IntegerFieldDefinition) HasIsMulti() bool`

HasIsMulti returns a boolean if a field has been set.

### GetIsRequired

`func (o *IntegerFieldDefinition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *IntegerFieldDefinition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *IntegerFieldDefinition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *IntegerFieldDefinition) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetName

`func (o *IntegerFieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegerFieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegerFieldDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegerFieldDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IntegerFieldDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegerFieldDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegerFieldDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegerFieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumericMax

`func (o *IntegerFieldDefinition) GetNumericMax() float32`

GetNumericMax returns the NumericMax field if non-nil, zero value otherwise.

### GetNumericMaxOk

`func (o *IntegerFieldDefinition) GetNumericMaxOk() (*float32, bool)`

GetNumericMaxOk returns a tuple with the NumericMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericMax

`func (o *IntegerFieldDefinition) SetNumericMax(v float32)`

SetNumericMax sets NumericMax field to given value.

### HasNumericMax

`func (o *IntegerFieldDefinition) HasNumericMax() bool`

HasNumericMax returns a boolean if a field has been set.

### SetNumericMaxNil

`func (o *IntegerFieldDefinition) SetNumericMaxNil(b bool)`

 SetNumericMaxNil sets the value for NumericMax to be an explicit nil

### UnsetNumericMax
`func (o *IntegerFieldDefinition) UnsetNumericMax()`

UnsetNumericMax ensures that no value is present for NumericMax, not even an explicit nil
### GetNumericMin

`func (o *IntegerFieldDefinition) GetNumericMin() float32`

GetNumericMin returns the NumericMin field if non-nil, zero value otherwise.

### GetNumericMinOk

`func (o *IntegerFieldDefinition) GetNumericMinOk() (*float32, bool)`

GetNumericMinOk returns a tuple with the NumericMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericMin

`func (o *IntegerFieldDefinition) SetNumericMin(v float32)`

SetNumericMin sets NumericMin field to given value.

### HasNumericMin

`func (o *IntegerFieldDefinition) HasNumericMin() bool`

HasNumericMin returns a boolean if a field has been set.

### SetNumericMinNil

`func (o *IntegerFieldDefinition) SetNumericMinNil(b bool)`

 SetNumericMinNil sets the value for NumericMin to be an explicit nil

### UnsetNumericMin
`func (o *IntegerFieldDefinition) UnsetNumericMin()`

UnsetNumericMin ensures that no value is present for NumericMin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


