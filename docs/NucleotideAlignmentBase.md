# NucleotideAlignmentBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Files** | [**[]NucleotideAlignmentBaseFilesInner**](NucleotideAlignmentBaseFilesInner.md) |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewNucleotideAlignmentBase

`func NewNucleotideAlignmentBase(algorithm string, files []NucleotideAlignmentBaseFilesInner, ) *NucleotideAlignmentBase`

NewNucleotideAlignmentBase instantiates a new NucleotideAlignmentBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideAlignmentBaseWithDefaults

`func NewNucleotideAlignmentBaseWithDefaults() *NucleotideAlignmentBase`

NewNucleotideAlignmentBaseWithDefaults instantiates a new NucleotideAlignmentBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *NucleotideAlignmentBase) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *NucleotideAlignmentBase) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *NucleotideAlignmentBase) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetFiles

`func (o *NucleotideAlignmentBase) GetFiles() []NucleotideAlignmentBaseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NucleotideAlignmentBase) GetFilesOk() (*[]NucleotideAlignmentBaseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NucleotideAlignmentBase) SetFiles(v []NucleotideAlignmentBaseFilesInner)`

SetFiles sets Files field to given value.


### GetName

`func (o *NucleotideAlignmentBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NucleotideAlignmentBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NucleotideAlignmentBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NucleotideAlignmentBase) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


