# RequestResponseSamplesItemBatch

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
**InaccessibleId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of this inaccessible item. Example values: \&quot;custom_entity\&quot;, \&quot;container\&quot;, \&quot;plate\&quot;, \&quot;dna_sequence\&quot;  | [optional] 

## Methods

### NewRequestResponseSamplesItemBatch

`func NewRequestResponseSamplesItemBatch() *RequestResponseSamplesItemBatch`

NewRequestResponseSamplesItemBatch instantiates a new RequestResponseSamplesItemBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestResponseSamplesItemBatchWithDefaults

`func NewRequestResponseSamplesItemBatchWithDefaults() *RequestResponseSamplesItemBatch`

NewRequestResponseSamplesItemBatchWithDefaults instantiates a new RequestResponseSamplesItemBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *RequestResponseSamplesItemBatch) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *RequestResponseSamplesItemBatch) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *RequestResponseSamplesItemBatch) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *RequestResponseSamplesItemBatch) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *RequestResponseSamplesItemBatch) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *RequestResponseSamplesItemBatch) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *RequestResponseSamplesItemBatch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestResponseSamplesItemBatch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestResponseSamplesItemBatch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestResponseSamplesItemBatch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *RequestResponseSamplesItemBatch) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RequestResponseSamplesItemBatch) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RequestResponseSamplesItemBatch) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RequestResponseSamplesItemBatch) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDefaultConcentration

`func (o *RequestResponseSamplesItemBatch) GetDefaultConcentration() Measurement`

GetDefaultConcentration returns the DefaultConcentration field if non-nil, zero value otherwise.

### GetDefaultConcentrationOk

`func (o *RequestResponseSamplesItemBatch) GetDefaultConcentrationOk() (*Measurement, bool)`

GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcentration

`func (o *RequestResponseSamplesItemBatch) SetDefaultConcentration(v Measurement)`

SetDefaultConcentration sets DefaultConcentration field to given value.

### HasDefaultConcentration

`func (o *RequestResponseSamplesItemBatch) HasDefaultConcentration() bool`

HasDefaultConcentration returns a boolean if a field has been set.

### GetEntity

`func (o *RequestResponseSamplesItemBatch) GetEntity() BatchEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RequestResponseSamplesItemBatch) GetEntityOk() (*BatchEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RequestResponseSamplesItemBatch) SetEntity(v BatchEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RequestResponseSamplesItemBatch) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFields

`func (o *RequestResponseSamplesItemBatch) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestResponseSamplesItemBatch) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestResponseSamplesItemBatch) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestResponseSamplesItemBatch) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *RequestResponseSamplesItemBatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestResponseSamplesItemBatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestResponseSamplesItemBatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestResponseSamplesItemBatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RequestResponseSamplesItemBatch) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RequestResponseSamplesItemBatch) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RequestResponseSamplesItemBatch) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RequestResponseSamplesItemBatch) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *RequestResponseSamplesItemBatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestResponseSamplesItemBatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestResponseSamplesItemBatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestResponseSamplesItemBatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchema

`func (o *RequestResponseSamplesItemBatch) GetSchema() SchemaProperty2`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RequestResponseSamplesItemBatch) GetSchemaOk() (*SchemaProperty2, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RequestResponseSamplesItemBatch) SetSchema(v SchemaProperty2)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RequestResponseSamplesItemBatch) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *RequestResponseSamplesItemBatch) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *RequestResponseSamplesItemBatch) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *RequestResponseSamplesItemBatch) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *RequestResponseSamplesItemBatch) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *RequestResponseSamplesItemBatch) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *RequestResponseSamplesItemBatch) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetInaccessibleId

`func (o *RequestResponseSamplesItemBatch) GetInaccessibleId() string`

GetInaccessibleId returns the InaccessibleId field if non-nil, zero value otherwise.

### GetInaccessibleIdOk

`func (o *RequestResponseSamplesItemBatch) GetInaccessibleIdOk() (*string, bool)`

GetInaccessibleIdOk returns a tuple with the InaccessibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInaccessibleId

`func (o *RequestResponseSamplesItemBatch) SetInaccessibleId(v string)`

SetInaccessibleId sets InaccessibleId field to given value.

### HasInaccessibleId

`func (o *RequestResponseSamplesItemBatch) HasInaccessibleId() bool`

HasInaccessibleId returns a boolean if a field has been set.

### GetType

`func (o *RequestResponseSamplesItemBatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestResponseSamplesItemBatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestResponseSamplesItemBatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RequestResponseSamplesItemBatch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


