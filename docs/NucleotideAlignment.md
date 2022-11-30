# NucleotideAlignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Alignment in the API. | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Alignment was created | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Alignment was last modified | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReferenceSequenceId** | Pointer to **string** | The ID of the template or consensus Sequence associated with the Alignment | [optional] 
**WebURL** | Pointer to **string** | The Benchling web UI url to view the Alignment | [optional] 
**AlignedSequences** | Pointer to [**[]AlignedNucleotideSequence**](AlignedNucleotideSequence.md) |  | [optional] 

## Methods

### NewNucleotideAlignment

`func NewNucleotideAlignment() *NucleotideAlignment`

NewNucleotideAlignment instantiates a new NucleotideAlignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideAlignmentWithDefaults

`func NewNucleotideAlignmentWithDefaults() *NucleotideAlignment`

NewNucleotideAlignmentWithDefaults instantiates a new NucleotideAlignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *NucleotideAlignment) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *NucleotideAlignment) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *NucleotideAlignment) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *NucleotideAlignment) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NucleotideAlignment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NucleotideAlignment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NucleotideAlignment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NucleotideAlignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *NucleotideAlignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NucleotideAlignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NucleotideAlignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NucleotideAlignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *NucleotideAlignment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NucleotideAlignment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NucleotideAlignment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *NucleotideAlignment) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *NucleotideAlignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NucleotideAlignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NucleotideAlignment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NucleotideAlignment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceSequenceId

`func (o *NucleotideAlignment) GetReferenceSequenceId() string`

GetReferenceSequenceId returns the ReferenceSequenceId field if non-nil, zero value otherwise.

### GetReferenceSequenceIdOk

`func (o *NucleotideAlignment) GetReferenceSequenceIdOk() (*string, bool)`

GetReferenceSequenceIdOk returns a tuple with the ReferenceSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSequenceId

`func (o *NucleotideAlignment) SetReferenceSequenceId(v string)`

SetReferenceSequenceId sets ReferenceSequenceId field to given value.

### HasReferenceSequenceId

`func (o *NucleotideAlignment) HasReferenceSequenceId() bool`

HasReferenceSequenceId returns a boolean if a field has been set.

### GetWebURL

`func (o *NucleotideAlignment) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *NucleotideAlignment) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *NucleotideAlignment) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *NucleotideAlignment) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetAlignedSequences

`func (o *NucleotideAlignment) GetAlignedSequences() []AlignedNucleotideSequence`

GetAlignedSequences returns the AlignedSequences field if non-nil, zero value otherwise.

### GetAlignedSequencesOk

`func (o *NucleotideAlignment) GetAlignedSequencesOk() (*[]AlignedNucleotideSequence, bool)`

GetAlignedSequencesOk returns a tuple with the AlignedSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignedSequences

`func (o *NucleotideAlignment) SetAlignedSequences(v []AlignedNucleotideSequence)`

SetAlignedSequences sets AlignedSequences field to given value.

### HasAlignedSequences

`func (o *NucleotideAlignment) HasAlignedSequences() bool`

HasAlignedSequences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


