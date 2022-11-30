# DnaConsensusAlignmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Files** | [**[]DnaAlignmentBaseFilesInner**](DnaAlignmentBaseFilesInner.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**NewSequence** | Pointer to [**DnaConsensusAlignmentCreateAllOfNewSequence**](DnaConsensusAlignmentCreateAllOfNewSequence.md) |  | [optional] 
**SequenceId** | Pointer to **string** |  | [optional] 

## Methods

### NewDnaConsensusAlignmentCreate

`func NewDnaConsensusAlignmentCreate(algorithm string, files []DnaAlignmentBaseFilesInner, ) *DnaConsensusAlignmentCreate`

NewDnaConsensusAlignmentCreate instantiates a new DnaConsensusAlignmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaConsensusAlignmentCreateWithDefaults

`func NewDnaConsensusAlignmentCreateWithDefaults() *DnaConsensusAlignmentCreate`

NewDnaConsensusAlignmentCreateWithDefaults instantiates a new DnaConsensusAlignmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *DnaConsensusAlignmentCreate) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DnaConsensusAlignmentCreate) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DnaConsensusAlignmentCreate) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetFiles

`func (o *DnaConsensusAlignmentCreate) GetFiles() []DnaAlignmentBaseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DnaConsensusAlignmentCreate) GetFilesOk() (*[]DnaAlignmentBaseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DnaConsensusAlignmentCreate) SetFiles(v []DnaAlignmentBaseFilesInner)`

SetFiles sets Files field to given value.


### GetName

`func (o *DnaConsensusAlignmentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaConsensusAlignmentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaConsensusAlignmentCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnaConsensusAlignmentCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewSequence

`func (o *DnaConsensusAlignmentCreate) GetNewSequence() DnaConsensusAlignmentCreateAllOfNewSequence`

GetNewSequence returns the NewSequence field if non-nil, zero value otherwise.

### GetNewSequenceOk

`func (o *DnaConsensusAlignmentCreate) GetNewSequenceOk() (*DnaConsensusAlignmentCreateAllOfNewSequence, bool)`

GetNewSequenceOk returns a tuple with the NewSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSequence

`func (o *DnaConsensusAlignmentCreate) SetNewSequence(v DnaConsensusAlignmentCreateAllOfNewSequence)`

SetNewSequence sets NewSequence field to given value.

### HasNewSequence

`func (o *DnaConsensusAlignmentCreate) HasNewSequence() bool`

HasNewSequence returns a boolean if a field has been set.

### GetSequenceId

`func (o *DnaConsensusAlignmentCreate) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *DnaConsensusAlignmentCreate) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *DnaConsensusAlignmentCreate) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *DnaConsensusAlignmentCreate) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


