# DnaTemplateAlignmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Files** | [**[]DnaAlignmentBaseFilesInner**](DnaAlignmentBaseFilesInner.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**TemplateSequenceId** | **string** |  | 

## Methods

### NewDnaTemplateAlignmentCreate

`func NewDnaTemplateAlignmentCreate(algorithm string, files []DnaAlignmentBaseFilesInner, templateSequenceId string, ) *DnaTemplateAlignmentCreate`

NewDnaTemplateAlignmentCreate instantiates a new DnaTemplateAlignmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaTemplateAlignmentCreateWithDefaults

`func NewDnaTemplateAlignmentCreateWithDefaults() *DnaTemplateAlignmentCreate`

NewDnaTemplateAlignmentCreateWithDefaults instantiates a new DnaTemplateAlignmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *DnaTemplateAlignmentCreate) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DnaTemplateAlignmentCreate) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DnaTemplateAlignmentCreate) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetFiles

`func (o *DnaTemplateAlignmentCreate) GetFiles() []DnaAlignmentBaseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DnaTemplateAlignmentCreate) GetFilesOk() (*[]DnaAlignmentBaseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DnaTemplateAlignmentCreate) SetFiles(v []DnaAlignmentBaseFilesInner)`

SetFiles sets Files field to given value.


### GetName

`func (o *DnaTemplateAlignmentCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaTemplateAlignmentCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaTemplateAlignmentCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnaTemplateAlignmentCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateSequenceId

`func (o *DnaTemplateAlignmentCreate) GetTemplateSequenceId() string`

GetTemplateSequenceId returns the TemplateSequenceId field if non-nil, zero value otherwise.

### GetTemplateSequenceIdOk

`func (o *DnaTemplateAlignmentCreate) GetTemplateSequenceIdOk() (*string, bool)`

GetTemplateSequenceIdOk returns a tuple with the TemplateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSequenceId

`func (o *DnaTemplateAlignmentCreate) SetTemplateSequenceId(v string)`

SetTemplateSequenceId sets TemplateSequenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


