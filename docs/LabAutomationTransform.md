# LabAutomationTransform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the transform in the API. | [optional] [readonly] 
**BlobId** | Pointer to **NullableString** |  | [optional] 
**CustomTransformId** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to [**LabAutomationBenchlingAppErrors**](LabAutomationBenchlingAppErrors.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the transform was last modified. | [optional] [readonly] 
**OutputProcessorId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewLabAutomationTransform

`func NewLabAutomationTransform() *LabAutomationTransform`

NewLabAutomationTransform instantiates a new LabAutomationTransform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabAutomationTransformWithDefaults

`func NewLabAutomationTransformWithDefaults() *LabAutomationTransform`

NewLabAutomationTransformWithDefaults instantiates a new LabAutomationTransform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *LabAutomationTransform) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *LabAutomationTransform) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *LabAutomationTransform) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *LabAutomationTransform) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetBlobId

`func (o *LabAutomationTransform) GetBlobId() string`

GetBlobId returns the BlobId field if non-nil, zero value otherwise.

### GetBlobIdOk

`func (o *LabAutomationTransform) GetBlobIdOk() (*string, bool)`

GetBlobIdOk returns a tuple with the BlobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobId

`func (o *LabAutomationTransform) SetBlobId(v string)`

SetBlobId sets BlobId field to given value.

### HasBlobId

`func (o *LabAutomationTransform) HasBlobId() bool`

HasBlobId returns a boolean if a field has been set.

### SetBlobIdNil

`func (o *LabAutomationTransform) SetBlobIdNil(b bool)`

 SetBlobIdNil sets the value for BlobId to be an explicit nil

### UnsetBlobId
`func (o *LabAutomationTransform) UnsetBlobId()`

UnsetBlobId ensures that no value is present for BlobId, not even an explicit nil
### GetCustomTransformId

`func (o *LabAutomationTransform) GetCustomTransformId() string`

GetCustomTransformId returns the CustomTransformId field if non-nil, zero value otherwise.

### GetCustomTransformIdOk

`func (o *LabAutomationTransform) GetCustomTransformIdOk() (*string, bool)`

GetCustomTransformIdOk returns a tuple with the CustomTransformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTransformId

`func (o *LabAutomationTransform) SetCustomTransformId(v string)`

SetCustomTransformId sets CustomTransformId field to given value.

### HasCustomTransformId

`func (o *LabAutomationTransform) HasCustomTransformId() bool`

HasCustomTransformId returns a boolean if a field has been set.

### SetCustomTransformIdNil

`func (o *LabAutomationTransform) SetCustomTransformIdNil(b bool)`

 SetCustomTransformIdNil sets the value for CustomTransformId to be an explicit nil

### UnsetCustomTransformId
`func (o *LabAutomationTransform) UnsetCustomTransformId()`

UnsetCustomTransformId ensures that no value is present for CustomTransformId, not even an explicit nil
### GetErrors

`func (o *LabAutomationTransform) GetErrors() LabAutomationBenchlingAppErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LabAutomationTransform) GetErrorsOk() (*LabAutomationBenchlingAppErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LabAutomationTransform) SetErrors(v LabAutomationBenchlingAppErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *LabAutomationTransform) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetId

`func (o *LabAutomationTransform) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabAutomationTransform) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabAutomationTransform) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabAutomationTransform) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *LabAutomationTransform) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LabAutomationTransform) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LabAutomationTransform) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *LabAutomationTransform) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetOutputProcessorId

`func (o *LabAutomationTransform) GetOutputProcessorId() string`

GetOutputProcessorId returns the OutputProcessorId field if non-nil, zero value otherwise.

### GetOutputProcessorIdOk

`func (o *LabAutomationTransform) GetOutputProcessorIdOk() (*string, bool)`

GetOutputProcessorIdOk returns a tuple with the OutputProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputProcessorId

`func (o *LabAutomationTransform) SetOutputProcessorId(v string)`

SetOutputProcessorId sets OutputProcessorId field to given value.

### HasOutputProcessorId

`func (o *LabAutomationTransform) HasOutputProcessorId() bool`

HasOutputProcessorId returns a boolean if a field has been set.

### GetStatus

`func (o *LabAutomationTransform) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LabAutomationTransform) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LabAutomationTransform) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LabAutomationTransform) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


