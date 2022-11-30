# Batch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the result was created | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] [readonly] 
**DefaultConcentration** | Pointer to [**Measurement**](Measurement.md) |  | [optional] 
**Entity** | Pointer to [**BatchEntity**](BatchEntity.md) |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to [**NullableSchemaProperty2**](SchemaProperty2.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBatch

`func NewBatch() *Batch`

NewBatch instantiates a new Batch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWithDefaults

`func NewBatchWithDefaults() *Batch`

NewBatchWithDefaults instantiates a new Batch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *Batch) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Batch) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Batch) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Batch) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Batch) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Batch) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *Batch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Batch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Batch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Batch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Batch) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Batch) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Batch) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Batch) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDefaultConcentration

`func (o *Batch) GetDefaultConcentration() Measurement`

GetDefaultConcentration returns the DefaultConcentration field if non-nil, zero value otherwise.

### GetDefaultConcentrationOk

`func (o *Batch) GetDefaultConcentrationOk() (*Measurement, bool)`

GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcentration

`func (o *Batch) SetDefaultConcentration(v Measurement)`

SetDefaultConcentration sets DefaultConcentration field to given value.

### HasDefaultConcentration

`func (o *Batch) HasDefaultConcentration() bool`

HasDefaultConcentration returns a boolean if a field has been set.

### GetEntity

`func (o *Batch) GetEntity() BatchEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Batch) GetEntityOk() (*BatchEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Batch) SetEntity(v BatchEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Batch) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFields

`func (o *Batch) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Batch) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Batch) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Batch) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *Batch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Batch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Batch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Batch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Batch) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Batch) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Batch) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Batch) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Batch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Batch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Batch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Batch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchema

`func (o *Batch) GetSchema() SchemaProperty2`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Batch) GetSchemaOk() (*SchemaProperty2, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Batch) SetSchema(v SchemaProperty2)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Batch) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Batch) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Batch) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *Batch) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Batch) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Batch) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Batch) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


