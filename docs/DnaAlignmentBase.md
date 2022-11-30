# DnaAlignmentBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** |  | 
**Files** | [**[]DnaAlignmentBaseFilesInner**](DnaAlignmentBaseFilesInner.md) |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewDnaAlignmentBase

`func NewDnaAlignmentBase(algorithm string, files []DnaAlignmentBaseFilesInner, ) *DnaAlignmentBase`

NewDnaAlignmentBase instantiates a new DnaAlignmentBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaAlignmentBaseWithDefaults

`func NewDnaAlignmentBaseWithDefaults() *DnaAlignmentBase`

NewDnaAlignmentBaseWithDefaults instantiates a new DnaAlignmentBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *DnaAlignmentBase) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DnaAlignmentBase) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DnaAlignmentBase) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetFiles

`func (o *DnaAlignmentBase) GetFiles() []DnaAlignmentBaseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *DnaAlignmentBase) GetFilesOk() (*[]DnaAlignmentBaseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *DnaAlignmentBase) SetFiles(v []DnaAlignmentBaseFilesInner)`

SetFiles sets Files field to given value.


### GetName

`func (o *DnaAlignmentBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaAlignmentBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaAlignmentBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnaAlignmentBase) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


