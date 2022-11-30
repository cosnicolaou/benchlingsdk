# RequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]AssayResult**](AssayResult.md) |  | [optional] 
**Samples** | Pointer to [**[]RequestResponseSamplesInner**](RequestResponseSamplesInner.md) | Array of samples produced by the request. | [optional] 

## Methods

### NewRequestResponse

`func NewRequestResponse() *RequestResponse`

NewRequestResponse instantiates a new RequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestResponseWithDefaults

`func NewRequestResponseWithDefaults() *RequestResponse`

NewRequestResponseWithDefaults instantiates a new RequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *RequestResponse) GetResults() []AssayResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RequestResponse) GetResultsOk() (*[]AssayResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RequestResponse) SetResults(v []AssayResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *RequestResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSamples

`func (o *RequestResponse) GetSamples() []RequestResponseSamplesInner`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *RequestResponse) GetSamplesOk() (*[]RequestResponseSamplesInner, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *RequestResponse) SetSamples(v []RequestResponseSamplesInner)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *RequestResponse) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


