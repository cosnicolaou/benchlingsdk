# EntryTemplateDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **int32** | 1 indexed day signifier. | [optional] 
**Notes** | Pointer to [**[]EntryTemplateDayNotesInner**](EntryTemplateDayNotesInner.md) |  | [optional] 

## Methods

### NewEntryTemplateDay

`func NewEntryTemplateDay() *EntryTemplateDay`

NewEntryTemplateDay instantiates a new EntryTemplateDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTemplateDayWithDefaults

`func NewEntryTemplateDayWithDefaults() *EntryTemplateDay`

NewEntryTemplateDayWithDefaults instantiates a new EntryTemplateDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *EntryTemplateDay) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *EntryTemplateDay) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *EntryTemplateDay) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *EntryTemplateDay) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetNotes

`func (o *EntryTemplateDay) GetNotes() []EntryTemplateDayNotesInner`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EntryTemplateDay) GetNotesOk() (*[]EntryTemplateDayNotesInner, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EntryTemplateDay) SetNotes(v []EntryTemplateDayNotesInner)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EntryTemplateDay) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


