# NucleotideConsensusAlignmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Files** | [**[]NucleotideAlignmentBaseFilesInner**](NucleotideAlignmentBaseFilesInner.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**NewSequence** | Pointer to [**DnaConsensusAlignmentCreateAllOfNewSequence**](DnaConsensusAlignmentCreateAllOfNewSequence.md) |  | [optional] 
**SequenceId** | Pointer to **string** |  | [optional] 

## Methods

### NewNucleotideConsensusAlignmentCreate

`func NewNucleotideConsensusAlignmentCreate(algorithm string, files []NucleotideAlignmentBaseFilesInner, ) *NucleotideConsensusAlignmentCreate`

NewNucleotideConsensusAlignmentCreate instantiates a new NucleotideConsensusAlignmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideConsensusAlignmentCreateWithDefaults

`func NewNucleotideConsensusAlignmentCreateWithDefaults() *NucleotideConsensusAlignmentCreate`

NewNucleotideConsensusAlignmentCreateWithDefaults instantiates a new NucleotideConsensusAlignmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *NucleotideConsensusAlignmentCreate) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *NucleotideConsensusAlignmentCreate) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *NucleotideConsensusAlignmentCreate) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetFiles

`func (o *NucleotideConsensusAlignmentCreate) GetFiles() []NucleotideAlignmentBaseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NucleotideConsensusAlignmentCreate) GetFilesOk() (*[]NucleotideAlignmentBaseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NucleotideConsensusAlignmentCreate) SetFiles(v []NucleotideAlignmentBaseFilesInner)`

SetFiles sets Files field to given value.


### GetName

`func (o *NucleotideConsensusAlignmentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NucleotideConsensusAlignmentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NucleotideConsensusAlignmentCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NucleotideConsensusAlignmentCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewSequence

`func (o *NucleotideConsensusAlignmentCreate) GetNewSequence() DnaConsensusAlignmentCreateAllOfNewSequence`

GetNewSequence returns the NewSequence field if non-nil, zero value otherwise.

### GetNewSequenceOk

`func (o *NucleotideConsensusAlignmentCreate) GetNewSequenceOk() (*DnaConsensusAlignmentCreateAllOfNewSequence, bool)`

GetNewSequenceOk returns a tuple with the NewSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSequence

`func (o *NucleotideConsensusAlignmentCreate) SetNewSequence(v DnaConsensusAlignmentCreateAllOfNewSequence)`

SetNewSequence sets NewSequence field to given value.

### HasNewSequence

`func (o *NucleotideConsensusAlignmentCreate) HasNewSequence() bool`

HasNewSequence returns a boolean if a field has been set.

### GetSequenceId

`func (o *NucleotideConsensusAlignmentCreate) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *NucleotideConsensusAlignmentCreate) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *NucleotideConsensusAlignmentCreate) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *NucleotideConsensusAlignmentCreate) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


