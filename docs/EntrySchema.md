# EntrySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the entry schema | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Entry Schema was last modified | [optional] 
**Name** | Pointer to **string** | Name of the entry schema | [optional] 

## Methods

### NewEntrySchema

`func NewEntrySchema() *EntrySchema`

NewEntrySchema instantiates a new EntrySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrySchemaWithDefaults

`func NewEntrySchemaWithDefaults() *EntrySchema`

NewEntrySchemaWithDefaults instantiates a new EntrySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntrySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntrySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntrySchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntrySchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntrySchema) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntrySchema) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntrySchema) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntrySchema) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *EntrySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntrySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntrySchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntrySchema) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


