# DropdownAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveRecord** | Pointer to [**NullableArchiveRecord**](ArchiveRecord.md) |  | [optional] 
**Options** | Pointer to [**[]DropdownOption**](DropdownOption.md) | Array of dropdown options | [optional] 

## Methods

### NewDropdownAllOf

`func NewDropdownAllOf() *DropdownAllOf`

NewDropdownAllOf instantiates a new DropdownAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropdownAllOfWithDefaults

`func NewDropdownAllOfWithDefaults() *DropdownAllOf`

NewDropdownAllOfWithDefaults instantiates a new DropdownAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveRecord

`func (o *DropdownAllOf) GetArchiveRecord() ArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *DropdownAllOf) GetArchiveRecordOk() (*ArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *DropdownAllOf) SetArchiveRecord(v ArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *DropdownAllOf) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *DropdownAllOf) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *DropdownAllOf) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetOptions

`func (o *DropdownAllOf) GetOptions() []DropdownOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DropdownAllOf) GetOptionsOk() (*[]DropdownOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DropdownAllOf) SetOptions(v []DropdownOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DropdownAllOf) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


