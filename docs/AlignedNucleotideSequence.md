# AlignedNucleotideSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bases** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PairwiseIdentity** | Pointer to **float32** | Fraction of bases between trimStart and trimEnd that match the template bases. Only present for Template Alignments; Will be empty for Consensus Alignments.  | [optional] 
**SequenceId** | Pointer to **NullableString** |  | [optional] 
**TrimEnd** | Pointer to **int32** |  | [optional] 
**TrimStart** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlignedNucleotideSequence

`func NewAlignedNucleotideSequence() *AlignedNucleotideSequence`

NewAlignedNucleotideSequence instantiates a new AlignedNucleotideSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlignedNucleotideSequenceWithDefaults

`func NewAlignedNucleotideSequenceWithDefaults() *AlignedNucleotideSequence`

NewAlignedNucleotideSequenceWithDefaults instantiates a new AlignedNucleotideSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBases

`func (o *AlignedNucleotideSequence) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *AlignedNucleotideSequence) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *AlignedNucleotideSequence) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *AlignedNucleotideSequence) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetName

`func (o *AlignedNucleotideSequence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlignedNucleotideSequence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlignedNucleotideSequence) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlignedNucleotideSequence) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPairwiseIdentity

`func (o *AlignedNucleotideSequence) GetPairwiseIdentity() float32`

GetPairwiseIdentity returns the PairwiseIdentity field if non-nil, zero value otherwise.

### GetPairwiseIdentityOk

`func (o *AlignedNucleotideSequence) GetPairwiseIdentityOk() (*float32, bool)`

GetPairwiseIdentityOk returns a tuple with the PairwiseIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairwiseIdentity

`func (o *AlignedNucleotideSequence) SetPairwiseIdentity(v float32)`

SetPairwiseIdentity sets PairwiseIdentity field to given value.

### HasPairwiseIdentity

`func (o *AlignedNucleotideSequence) HasPairwiseIdentity() bool`

HasPairwiseIdentity returns a boolean if a field has been set.

### GetSequenceId

`func (o *AlignedNucleotideSequence) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *AlignedNucleotideSequence) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *AlignedNucleotideSequence) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *AlignedNucleotideSequence) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### SetSequenceIdNil

`func (o *AlignedNucleotideSequence) SetSequenceIdNil(b bool)`

 SetSequenceIdNil sets the value for SequenceId to be an explicit nil

### UnsetSequenceId
`func (o *AlignedNucleotideSequence) UnsetSequenceId()`

UnsetSequenceId ensures that no value is present for SequenceId, not even an explicit nil
### GetTrimEnd

`func (o *AlignedNucleotideSequence) GetTrimEnd() int32`

GetTrimEnd returns the TrimEnd field if non-nil, zero value otherwise.

### GetTrimEndOk

`func (o *AlignedNucleotideSequence) GetTrimEndOk() (*int32, bool)`

GetTrimEndOk returns a tuple with the TrimEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimEnd

`func (o *AlignedNucleotideSequence) SetTrimEnd(v int32)`

SetTrimEnd sets TrimEnd field to given value.

### HasTrimEnd

`func (o *AlignedNucleotideSequence) HasTrimEnd() bool`

HasTrimEnd returns a boolean if a field has been set.

### GetTrimStart

`func (o *AlignedNucleotideSequence) GetTrimStart() int32`

GetTrimStart returns the TrimStart field if non-nil, zero value otherwise.

### GetTrimStartOk

`func (o *AlignedNucleotideSequence) GetTrimStartOk() (*int32, bool)`

GetTrimStartOk returns a tuple with the TrimStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimStart

`func (o *AlignedNucleotideSequence) SetTrimStart(v int32)`

SetTrimStart sets TrimStart field to given value.

### HasTrimStart

`func (o *AlignedNucleotideSequence) HasTrimStart() bool`

HasTrimStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


