# EntriesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryIds** | **[]string** | Array of entry IDs | 
**Reason** | **string** | Reason that entries are being archived. One of [\&quot;Made in error\&quot;, \&quot;Retired\&quot;, \&quot;Other\&quot;].  | 

## Methods

### NewEntriesArchive

`func NewEntriesArchive(entryIds []string, reason string, ) *EntriesArchive`

NewEntriesArchive instantiates a new EntriesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntriesArchiveWithDefaults

`func NewEntriesArchiveWithDefaults() *EntriesArchive`

NewEntriesArchiveWithDefaults instantiates a new EntriesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryIds

`func (o *EntriesArchive) GetEntryIds() []string`

GetEntryIds returns the EntryIds field if non-nil, zero value otherwise.

### GetEntryIdsOk

`func (o *EntriesArchive) GetEntryIdsOk() (*[]string, bool)`

GetEntryIdsOk returns a tuple with the EntryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryIds

`func (o *EntriesArchive) SetEntryIds(v []string)`

SetEntryIds sets EntryIds field to given value.


### GetReason

`func (o *EntriesArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EntriesArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EntriesArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


