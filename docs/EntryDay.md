# EntryDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | A Date string | [optional] 
**Notes** | Pointer to [**[]EntryDayNotesInner**](EntryDayNotesInner.md) |  | [optional] 

## Methods

### NewEntryDay

`func NewEntryDay() *EntryDay`

NewEntryDay instantiates a new EntryDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryDayWithDefaults

`func NewEntryDayWithDefaults() *EntryDay`

NewEntryDayWithDefaults instantiates a new EntryDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EntryDay) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EntryDay) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EntryDay) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EntryDay) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNotes

`func (o *EntryDay) GetNotes() []EntryDayNotesInner`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EntryDay) GetNotesOk() (*[]EntryDayNotesInner, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EntryDay) SetNotes(v []EntryDayNotesInner)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EntryDay) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


