# EntitySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**FieldDefinitions** | Pointer to [**[]SchemaFieldDefinitionsInner**](SchemaFieldDefinitionsInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Prefix** | Pointer to **string** |  | [optional] 
**RegistryId** | Pointer to **string** |  | [optional] 
**Constraint** | Pointer to [**NullableEntitySchemaAllOfConstraint**](EntitySchemaAllOfConstraint.md) |  | [optional] 
**ContainableType** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Entity Schema was last modified | [optional] 

## Methods

### NewEntitySchema

`func NewEntitySchema() *EntitySchema`

NewEntitySchema instantiates a new EntitySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaWithDefaults

`func NewEntitySchemaWithDefaults() *EntitySchema`

NewEntitySchemaWithDefaults instantiates a new EntitySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *EntitySchema) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *EntitySchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *EntitySchema) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *EntitySchema) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *EntitySchema) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *EntitySchema) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetFieldDefinitions

`func (o *EntitySchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *EntitySchema) GetFieldDefinitionsOk() (*[]SchemaFieldDefinitionsInner, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *EntitySchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner)`

SetFieldDefinitions sets FieldDefinitions field to given value.

### HasFieldDefinitions

`func (o *EntitySchema) HasFieldDefinitions() bool`

HasFieldDefinitions returns a boolean if a field has been set.

### GetId

`func (o *EntitySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitySchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitySchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntitySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitySchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitySchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EntitySchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitySchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitySchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntitySchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrefix

`func (o *EntitySchema) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EntitySchema) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EntitySchema) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *EntitySchema) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegistryId

`func (o *EntitySchema) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *EntitySchema) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *EntitySchema) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *EntitySchema) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetConstraint

`func (o *EntitySchema) GetConstraint() EntitySchemaAllOfConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *EntitySchema) GetConstraintOk() (*EntitySchemaAllOfConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *EntitySchema) SetConstraint(v EntitySchemaAllOfConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *EntitySchema) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### SetConstraintNil

`func (o *EntitySchema) SetConstraintNil(b bool)`

 SetConstraintNil sets the value for Constraint to be an explicit nil

### UnsetConstraint
`func (o *EntitySchema) UnsetConstraint()`

UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
### GetContainableType

`func (o *EntitySchema) GetContainableType() string`

GetContainableType returns the ContainableType field if non-nil, zero value otherwise.

### GetContainableTypeOk

`func (o *EntitySchema) GetContainableTypeOk() (*string, bool)`

GetContainableTypeOk returns a tuple with the ContainableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainableType

`func (o *EntitySchema) SetContainableType(v string)`

SetContainableType sets ContainableType field to given value.

### HasContainableType

`func (o *EntitySchema) HasContainableType() bool`

HasContainableType returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntitySchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntitySchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntitySchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntitySchema) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


