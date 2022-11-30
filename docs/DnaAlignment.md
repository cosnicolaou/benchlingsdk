# DnaAlignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the DNA Alignment in the API. | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the DNA Alignment was created | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the DNA Alignment was last modified | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReferenceSequenceId** | Pointer to **string** | The ID of the template or consensus DNA Sequence associated with the DNA Alignment | [optional] 
**WebURL** | Pointer to **string** | The Benchling web UI url to view the DNA Alignment | [optional] 
**AlignedSequences** | Pointer to [**[]AlignedSequence**](AlignedSequence.md) |  | [optional] 

## Methods

### NewDnaAlignment

`func NewDnaAlignment() *DnaAlignment`

NewDnaAlignment instantiates a new DnaAlignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaAlignmentWithDefaults

`func NewDnaAlignmentWithDefaults() *DnaAlignment`

NewDnaAlignmentWithDefaults instantiates a new DnaAlignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *DnaAlignment) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *DnaAlignment) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *DnaAlignment) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *DnaAlignment) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DnaAlignment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DnaAlignment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DnaAlignment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DnaAlignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DnaAlignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnaAlignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnaAlignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnaAlignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *DnaAlignment) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DnaAlignment) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DnaAlignment) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *DnaAlignment) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *DnaAlignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaAlignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaAlignment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnaAlignment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceSequenceId

`func (o *DnaAlignment) GetReferenceSequenceId() string`

GetReferenceSequenceId returns the ReferenceSequenceId field if non-nil, zero value otherwise.

### GetReferenceSequenceIdOk

`func (o *DnaAlignment) GetReferenceSequenceIdOk() (*string, bool)`

GetReferenceSequenceIdOk returns a tuple with the ReferenceSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSequenceId

`func (o *DnaAlignment) SetReferenceSequenceId(v string)`

SetReferenceSequenceId sets ReferenceSequenceId field to given value.

### HasReferenceSequenceId

`func (o *DnaAlignment) HasReferenceSequenceId() bool`

HasReferenceSequenceId returns a boolean if a field has been set.

### GetWebURL

`func (o *DnaAlignment) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *DnaAlignment) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *DnaAlignment) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *DnaAlignment) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetAlignedSequences

`func (o *DnaAlignment) GetAlignedSequences() []AlignedSequence`

GetAlignedSequences returns the AlignedSequences field if non-nil, zero value otherwise.

### GetAlignedSequencesOk

`func (o *DnaAlignment) GetAlignedSequencesOk() (*[]AlignedSequence, bool)`

GetAlignedSequencesOk returns a tuple with the AlignedSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignedSequences

`func (o *DnaAlignment) SetAlignedSequences(v []AlignedSequence)`

SetAlignedSequences sets AlignedSequences field to given value.

### HasAlignedSequences

`func (o *DnaAlignment) HasAlignedSequences() bool`

HasAlignedSequences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


