# RequestFulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time the fulfillment was created | [optional] 
**EntryId** | Pointer to **string** | ID of the entry this fulfillment was executed in, if any | [optional] 
**Id** | Pointer to **string** | ID of the request fulfillment | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Request Fulfillment was last modified | [optional] [readonly] 
**RequestId** | Pointer to **string** | ID of the request this fulfillment fulfills | [optional] 
**SampleGroup** | Pointer to [**NullableRequestFulfillmentSampleGroup**](RequestFulfillmentSampleGroup.md) |  | [optional] 
**Status** | Pointer to [**SampleGroupStatus**](SampleGroupStatus.md) |  | [optional] 

## Methods

### NewRequestFulfillment

`func NewRequestFulfillment() *RequestFulfillment`

NewRequestFulfillment instantiates a new RequestFulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFulfillmentWithDefaults

`func NewRequestFulfillmentWithDefaults() *RequestFulfillment`

NewRequestFulfillmentWithDefaults instantiates a new RequestFulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RequestFulfillment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestFulfillment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestFulfillment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestFulfillment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntryId

`func (o *RequestFulfillment) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *RequestFulfillment) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *RequestFulfillment) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *RequestFulfillment) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetId

`func (o *RequestFulfillment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestFulfillment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestFulfillment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestFulfillment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RequestFulfillment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RequestFulfillment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RequestFulfillment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RequestFulfillment) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRequestId

`func (o *RequestFulfillment) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RequestFulfillment) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RequestFulfillment) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RequestFulfillment) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSampleGroup

`func (o *RequestFulfillment) GetSampleGroup() RequestFulfillmentSampleGroup`

GetSampleGroup returns the SampleGroup field if non-nil, zero value otherwise.

### GetSampleGroupOk

`func (o *RequestFulfillment) GetSampleGroupOk() (*RequestFulfillmentSampleGroup, bool)`

GetSampleGroupOk returns a tuple with the SampleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroup

`func (o *RequestFulfillment) SetSampleGroup(v RequestFulfillmentSampleGroup)`

SetSampleGroup sets SampleGroup field to given value.

### HasSampleGroup

`func (o *RequestFulfillment) HasSampleGroup() bool`

HasSampleGroup returns a boolean if a field has been set.

### SetSampleGroupNil

`func (o *RequestFulfillment) SetSampleGroupNil(b bool)`

 SetSampleGroupNil sets the value for SampleGroup to be an explicit nil

### UnsetSampleGroup
`func (o *RequestFulfillment) UnsetSampleGroup()`

UnsetSampleGroup ensures that no value is present for SampleGroup, not even an explicit nil
### GetStatus

`func (o *RequestFulfillment) GetStatus() SampleGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestFulfillment) GetStatusOk() (*SampleGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestFulfillment) SetStatus(v SampleGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestFulfillment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


