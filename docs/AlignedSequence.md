# AlignedSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bases** | Pointer to **string** |  | [optional] 
**DnaSequenceId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PairwiseIdentity** | Pointer to **float32** | Fraction of bases between trimStart and trimEnd that match the template bases. Only present for Template Alignments; Will be empty for Consensus Alignments.  | [optional] 
**TrimEnd** | Pointer to **int32** |  | [optional] 
**TrimStart** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlignedSequence

`func NewAlignedSequence() *AlignedSequence`

NewAlignedSequence instantiates a new AlignedSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlignedSequenceWithDefaults

`func NewAlignedSequenceWithDefaults() *AlignedSequence`

NewAlignedSequenceWithDefaults instantiates a new AlignedSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBases

`func (o *AlignedSequence) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *AlignedSequence) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *AlignedSequence) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *AlignedSequence) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetDnaSequenceId

`func (o *AlignedSequence) GetDnaSequenceId() string`

GetDnaSequenceId returns the DnaSequenceId field if non-nil, zero value otherwise.

### GetDnaSequenceIdOk

`func (o *AlignedSequence) GetDnaSequenceIdOk() (*string, bool)`

GetDnaSequenceIdOk returns a tuple with the DnaSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaSequenceId

`func (o *AlignedSequence) SetDnaSequenceId(v string)`

SetDnaSequenceId sets DnaSequenceId field to given value.

### HasDnaSequenceId

`func (o *AlignedSequence) HasDnaSequenceId() bool`

HasDnaSequenceId returns a boolean if a field has been set.

### SetDnaSequenceIdNil

`func (o *AlignedSequence) SetDnaSequenceIdNil(b bool)`

 SetDnaSequenceIdNil sets the value for DnaSequenceId to be an explicit nil

### UnsetDnaSequenceId
`func (o *AlignedSequence) UnsetDnaSequenceId()`

UnsetDnaSequenceId ensures that no value is present for DnaSequenceId, not even an explicit nil
### GetName

`func (o *AlignedSequence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlignedSequence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlignedSequence) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlignedSequence) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPairwiseIdentity

`func (o *AlignedSequence) GetPairwiseIdentity() float32`

GetPairwiseIdentity returns the PairwiseIdentity field if non-nil, zero value otherwise.

### GetPairwiseIdentityOk

`func (o *AlignedSequence) GetPairwiseIdentityOk() (*float32, bool)`

GetPairwiseIdentityOk returns a tuple with the PairwiseIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairwiseIdentity

`func (o *AlignedSequence) SetPairwiseIdentity(v float32)`

SetPairwiseIdentity sets PairwiseIdentity field to given value.

### HasPairwiseIdentity

`func (o *AlignedSequence) HasPairwiseIdentity() bool`

HasPairwiseIdentity returns a boolean if a field has been set.

### GetTrimEnd

`func (o *AlignedSequence) GetTrimEnd() int32`

GetTrimEnd returns the TrimEnd field if non-nil, zero value otherwise.

### GetTrimEndOk

`func (o *AlignedSequence) GetTrimEndOk() (*int32, bool)`

GetTrimEndOk returns a tuple with the TrimEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimEnd

`func (o *AlignedSequence) SetTrimEnd(v int32)`

SetTrimEnd sets TrimEnd field to given value.

### HasTrimEnd

`func (o *AlignedSequence) HasTrimEnd() bool`

HasTrimEnd returns a boolean if a field has been set.

### GetTrimStart

`func (o *AlignedSequence) GetTrimStart() int32`

GetTrimStart returns the TrimStart field if non-nil, zero value otherwise.

### GetTrimStartOk

`func (o *AlignedSequence) GetTrimStartOk() (*int32, bool)`

GetTrimStartOk returns a tuple with the TrimStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimStart

`func (o *AlignedSequence) SetTrimStart(v int32)`

SetTrimStart sets TrimStart field to given value.

### HasTrimStart

`func (o *AlignedSequence) HasTrimStart() bool`

HasTrimStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


