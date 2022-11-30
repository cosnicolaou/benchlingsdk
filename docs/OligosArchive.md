# OligosArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OligoIds** | **[]string** |  | 
**Reason** | [**EntityArchiveReason**](EntityArchiveReason.md) |  | 

## Methods

### NewOligosArchive

`func NewOligosArchive(oligoIds []string, reason EntityArchiveReason, ) *OligosArchive`

NewOligosArchive instantiates a new OligosArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOligosArchiveWithDefaults

`func NewOligosArchiveWithDefaults() *OligosArchive`

NewOligosArchiveWithDefaults instantiates a new OligosArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOligoIds

`func (o *OligosArchive) GetOligoIds() []string`

GetOligoIds returns the OligoIds field if non-nil, zero value otherwise.

### GetOligoIdsOk

`func (o *OligosArchive) GetOligoIdsOk() (*[]string, bool)`

GetOligoIdsOk returns a tuple with the OligoIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOligoIds

`func (o *OligosArchive) SetOligoIds(v []string)`

SetOligoIds sets OligoIds field to given value.


### GetReason

`func (o *OligosArchive) GetReason() EntityArchiveReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OligosArchive) GetReasonOk() (*EntityArchiveReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OligosArchive) SetReason(v EntityArchiveReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


