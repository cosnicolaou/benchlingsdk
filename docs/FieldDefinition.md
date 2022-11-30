# FieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsMulti** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**FieldType**](FieldType.md) |  | [optional] 

## Methods

### NewFieldDefinition

`func NewFieldDefinition() *FieldDefinition`

NewFieldDefinition instantiates a new FieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDefinitionWithDefaults

`func NewFieldDefinitionWithDefaults() *FieldDefinition`

NewFieldDefinitionWithDefaults instantiates a new FieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *FieldDefinition) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *FieldDefinition) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *FieldDefinition) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *FieldDefinition) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *FieldDefinition) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *FieldDefinition) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetId

`func (o *FieldDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FieldDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMulti

`func (o *FieldDefinition) GetIsMulti() bool`

GetIsMulti returns the IsMulti field if non-nil, zero value otherwise.

### GetIsMultiOk

`func (o *FieldDefinition) GetIsMultiOk() (*bool, bool)`

GetIsMultiOk returns a tuple with the IsMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulti

`func (o *FieldDefinition) SetIsMulti(v bool)`

SetIsMulti sets IsMulti field to given value.

### HasIsMulti

`func (o *FieldDefinition) HasIsMulti() bool`

HasIsMulti returns a boolean if a field has been set.

### GetIsRequired

`func (o *FieldDefinition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *FieldDefinition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *FieldDefinition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *FieldDefinition) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetName

`func (o *FieldDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FieldDefinition) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDefinition) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDefinition) SetType(v FieldType)`

SetType sets Type field to given value.

### HasType

`func (o *FieldDefinition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


