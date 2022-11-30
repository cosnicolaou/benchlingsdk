# DnaAlignmentSummary

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

## Methods

### NewDnaAlignmentSummary

`func NewDnaAlignmentSummary() *DnaAlignmentSummary`

NewDnaAlignmentSummary instantiates a new DnaAlignmentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaAlignmentSummaryWithDefaults

`func NewDnaAlignmentSummaryWithDefaults() *DnaAlignmentSummary`

NewDnaAlignmentSummaryWithDefaults instantiates a new DnaAlignmentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *DnaAlignmentSummary) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *DnaAlignmentSummary) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *DnaAlignmentSummary) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *DnaAlignmentSummary) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DnaAlignmentSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DnaAlignmentSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DnaAlignmentSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DnaAlignmentSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DnaAlignmentSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnaAlignmentSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnaAlignmentSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnaAlignmentSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *DnaAlignmentSummary) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DnaAlignmentSummary) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DnaAlignmentSummary) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *DnaAlignmentSummary) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *DnaAlignmentSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaAlignmentSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaAlignmentSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnaAlignmentSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferenceSequenceId

`func (o *DnaAlignmentSummary) GetReferenceSequenceId() string`

GetReferenceSequenceId returns the ReferenceSequenceId field if non-nil, zero value otherwise.

### GetReferenceSequenceIdOk

`func (o *DnaAlignmentSummary) GetReferenceSequenceIdOk() (*string, bool)`

GetReferenceSequenceIdOk returns a tuple with the ReferenceSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSequenceId

`func (o *DnaAlignmentSummary) SetReferenceSequenceId(v string)`

SetReferenceSequenceId sets ReferenceSequenceId field to given value.

### HasReferenceSequenceId

`func (o *DnaAlignmentSummary) HasReferenceSequenceId() bool`

HasReferenceSequenceId returns a boolean if a field has been set.

### GetWebURL

`func (o *DnaAlignmentSummary) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *DnaAlignmentSummary) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *DnaAlignmentSummary) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *DnaAlignmentSummary) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


