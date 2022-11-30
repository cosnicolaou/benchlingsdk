# RequestTasksBulkCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]RequestTask**](RequestTask.md) | The created tasks | [optional] 

## Methods

### NewRequestTasksBulkCreateResponse

`func NewRequestTasksBulkCreateResponse() *RequestTasksBulkCreateResponse`

NewRequestTasksBulkCreateResponse instantiates a new RequestTasksBulkCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTasksBulkCreateResponseWithDefaults

`func NewRequestTasksBulkCreateResponseWithDefaults() *RequestTasksBulkCreateResponse`

NewRequestTasksBulkCreateResponseWithDefaults instantiates a new RequestTasksBulkCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *RequestTasksBulkCreateResponse) GetTasks() []RequestTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RequestTasksBulkCreateResponse) GetTasksOk() (*[]RequestTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RequestTasksBulkCreateResponse) SetTasks(v []RequestTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *RequestTasksBulkCreateResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


