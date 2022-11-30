# Dropdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the dropdown | [optional] 
**Name** | Pointer to **string** | Name of the dropdown | [optional] 
**ArchiveRecord** | Pointer to [**NullableArchiveRecord**](ArchiveRecord.md) |  | [optional] 
**Options** | Pointer to [**[]DropdownOption**](DropdownOption.md) | Array of dropdown options | [optional] 

## Methods

### NewDropdown

`func NewDropdown() *Dropdown`

NewDropdown instantiates a new Dropdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropdownWithDefaults

`func NewDropdownWithDefaults() *Dropdown`

NewDropdownWithDefaults instantiates a new Dropdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dropdown) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dropdown) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dropdown) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dropdown) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Dropdown) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dropdown) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dropdown) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dropdown) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *Dropdown) GetArchiveRecord() ArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Dropdown) GetArchiveRecordOk() (*ArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Dropdown) SetArchiveRecord(v ArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Dropdown) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Dropdown) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Dropdown) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetOptions

`func (o *Dropdown) GetOptions() []DropdownOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Dropdown) GetOptionsOk() (*[]DropdownOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Dropdown) SetOptions(v []DropdownOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Dropdown) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


