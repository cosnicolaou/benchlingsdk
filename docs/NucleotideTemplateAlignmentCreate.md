# NucleotideTemplateAlignmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Files** | [**[]NucleotideAlignmentBaseFilesInner**](NucleotideAlignmentBaseFilesInner.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**TemplateSequenceId** | **string** |  | 

## Methods

### NewNucleotideTemplateAlignmentCreate

`func NewNucleotideTemplateAlignmentCreate(algorithm string, files []NucleotideAlignmentBaseFilesInner, templateSequenceId string, ) *NucleotideTemplateAlignmentCreate`

NewNucleotideTemplateAlignmentCreate instantiates a new NucleotideTemplateAlignmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideTemplateAlignmentCreateWithDefaults

`func NewNucleotideTemplateAlignmentCreateWithDefaults() *NucleotideTemplateAlignmentCreate`

NewNucleotideTemplateAlignmentCreateWithDefaults instantiates a new NucleotideTemplateAlignmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *NucleotideTemplateAlignmentCreate) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *NucleotideTemplateAlignmentCreate) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *NucleotideTemplateAlignmentCreate) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetFiles

`func (o *NucleotideTemplateAlignmentCreate) GetFiles() []NucleotideAlignmentBaseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NucleotideTemplateAlignmentCreate) GetFilesOk() (*[]NucleotideAlignmentBaseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NucleotideTemplateAlignmentCreate) SetFiles(v []NucleotideAlignmentBaseFilesInner)`

SetFiles sets Files field to given value.


### GetName

`func (o *NucleotideTemplateAlignmentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NucleotideTemplateAlignmentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NucleotideTemplateAlignmentCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NucleotideTemplateAlignmentCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateSequenceId

`func (o *NucleotideTemplateAlignmentCreate) GetTemplateSequenceId() string`

GetTemplateSequenceId returns the TemplateSequenceId field if non-nil, zero value otherwise.

### GetTemplateSequenceIdOk

`func (o *NucleotideTemplateAlignmentCreate) GetTemplateSequenceIdOk() (*string, bool)`

GetTemplateSequenceIdOk returns a tuple with the TemplateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSequenceId

`func (o *NucleotideTemplateAlignmentCreate) SetTemplateSequenceId(v string)`

SetTemplateSequenceId sets TemplateSequenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


