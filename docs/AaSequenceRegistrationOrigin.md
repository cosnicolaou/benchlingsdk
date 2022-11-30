# AaSequenceRegistrationOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginEntryId** | Pointer to **NullableString** |  | [optional] [readonly] 
**RegisteredAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewAaSequenceRegistrationOrigin

`func NewAaSequenceRegistrationOrigin() *AaSequenceRegistrationOrigin`

NewAaSequenceRegistrationOrigin instantiates a new AaSequenceRegistrationOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAaSequenceRegistrationOriginWithDefaults

`func NewAaSequenceRegistrationOriginWithDefaults() *AaSequenceRegistrationOrigin`

NewAaSequenceRegistrationOriginWithDefaults instantiates a new AaSequenceRegistrationOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginEntryId

`func (o *AaSequenceRegistrationOrigin) GetOriginEntryId() string`

GetOriginEntryId returns the OriginEntryId field if non-nil, zero value otherwise.

### GetOriginEntryIdOk

`func (o *AaSequenceRegistrationOrigin) GetOriginEntryIdOk() (*string, bool)`

GetOriginEntryIdOk returns a tuple with the OriginEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginEntryId

`func (o *AaSequenceRegistrationOrigin) SetOriginEntryId(v string)`

SetOriginEntryId sets OriginEntryId field to given value.

### HasOriginEntryId

`func (o *AaSequenceRegistrationOrigin) HasOriginEntryId() bool`

HasOriginEntryId returns a boolean if a field has been set.

### SetOriginEntryIdNil

`func (o *AaSequenceRegistrationOrigin) SetOriginEntryIdNil(b bool)`

 SetOriginEntryIdNil sets the value for OriginEntryId to be an explicit nil

### UnsetOriginEntryId
`func (o *AaSequenceRegistrationOrigin) UnsetOriginEntryId()`

UnsetOriginEntryId ensures that no value is present for OriginEntryId, not even an explicit nil
### GetRegisteredAt

`func (o *AaSequenceRegistrationOrigin) GetRegisteredAt() time.Time`

GetRegisteredAt returns the RegisteredAt field if non-nil, zero value otherwise.

### GetRegisteredAtOk

`func (o *AaSequenceRegistrationOrigin) GetRegisteredAtOk() (*time.Time, bool)`

GetRegisteredAtOk returns a tuple with the RegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAt

`func (o *AaSequenceRegistrationOrigin) SetRegisteredAt(v time.Time)`

SetRegisteredAt sets RegisteredAt field to given value.

### HasRegisteredAt

`func (o *AaSequenceRegistrationOrigin) HasRegisteredAt() bool`

HasRegisteredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


