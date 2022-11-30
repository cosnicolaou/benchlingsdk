# RequestResponseSamplesInnerBatch

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

### NewRequestResponseSamplesInnerBatch

`func NewRequestResponseSamplesInnerBatch() *RequestResponseSamplesInnerBatch`

NewRequestResponseSamplesInnerBatch instantiates a new RequestResponseSamplesInnerBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestResponseSamplesInnerBatchWithDefaults

`func NewRequestResponseSamplesInnerBatchWithDefaults() *RequestResponseSamplesInnerBatch`

NewRequestResponseSamplesInnerBatchWithDefaults instantiates a new RequestResponseSamplesInnerBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *RequestResponseSamplesInnerBatch) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *RequestResponseSamplesInnerBatch) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *RequestResponseSamplesInnerBatch) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *RequestResponseSamplesInnerBatch) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *RequestResponseSamplesInnerBatch) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *RequestResponseSamplesInnerBatch) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetCreatedAt

`func (o *RequestResponseSamplesInnerBatch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestResponseSamplesInnerBatch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestResponseSamplesInnerBatch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestResponseSamplesInnerBatch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *RequestResponseSamplesInnerBatch) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RequestResponseSamplesInnerBatch) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RequestResponseSamplesInnerBatch) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RequestResponseSamplesInnerBatch) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDefaultConcentration

`func (o *RequestResponseSamplesInnerBatch) GetDefaultConcentration() Measurement`

GetDefaultConcentration returns the DefaultConcentration field if non-nil, zero value otherwise.

### GetDefaultConcentrationOk

`func (o *RequestResponseSamplesInnerBatch) GetDefaultConcentrationOk() (*Measurement, bool)`

GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcentration

`func (o *RequestResponseSamplesInnerBatch) SetDefaultConcentration(v Measurement)`

SetDefaultConcentration sets DefaultConcentration field to given value.

### HasDefaultConcentration

`func (o *RequestResponseSamplesInnerBatch) HasDefaultConcentration() bool`

HasDefaultConcentration returns a boolean if a field has been set.

### GetEntity

`func (o *RequestResponseSamplesInnerBatch) GetEntity() BatchEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RequestResponseSamplesInnerBatch) GetEntityOk() (*BatchEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RequestResponseSamplesInnerBatch) SetEntity(v BatchEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RequestResponseSamplesInnerBatch) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFields

`func (o *RequestResponseSamplesInnerBatch) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestResponseSamplesInnerBatch) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestResponseSamplesInnerBatch) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestResponseSamplesInnerBatch) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *RequestResponseSamplesInnerBatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestResponseSamplesInnerBatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestResponseSamplesInnerBatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestResponseSamplesInnerBatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RequestResponseSamplesInnerBatch) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RequestResponseSamplesInnerBatch) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RequestResponseSamplesInnerBatch) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RequestResponseSamplesInnerBatch) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *RequestResponseSamplesInnerBatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestResponseSamplesInnerBatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestResponseSamplesInnerBatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RequestResponseSamplesInnerBatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchema

`func (o *RequestResponseSamplesInnerBatch) GetSchema() SchemaProperty2`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RequestResponseSamplesInnerBatch) GetSchemaOk() (*SchemaProperty2, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RequestResponseSamplesInnerBatch) SetSchema(v SchemaProperty2)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RequestResponseSamplesInnerBatch) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *RequestResponseSamplesInnerBatch) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *RequestResponseSamplesInnerBatch) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetWebURL

`func (o *RequestResponseSamplesInnerBatch) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *RequestResponseSamplesInnerBatch) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *RequestResponseSamplesInnerBatch) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *RequestResponseSamplesInnerBatch) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetInaccessibleId

`func (o *RequestResponseSamplesInnerBatch) GetInaccessibleId() string`

GetInaccessibleId returns the InaccessibleId field if non-nil, zero value otherwise.

### GetInaccessibleIdOk

`func (o *RequestResponseSamplesInnerBatch) GetInaccessibleIdOk() (*string, bool)`

GetInaccessibleIdOk returns a tuple with the InaccessibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInaccessibleId

`func (o *RequestResponseSamplesInnerBatch) SetInaccessibleId(v string)`

SetInaccessibleId sets InaccessibleId field to given value.

### HasInaccessibleId

`func (o *RequestResponseSamplesInnerBatch) HasInaccessibleId() bool`

HasInaccessibleId returns a boolean if a field has been set.

### GetType

`func (o *RequestResponseSamplesInnerBatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestResponseSamplesInnerBatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestResponseSamplesInnerBatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RequestResponseSamplesInnerBatch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


