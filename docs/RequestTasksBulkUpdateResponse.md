# RequestTasksBulkUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]RequestTask**](RequestTask.md) | The tasks to update | [optional] 

## Methods

### NewRequestTasksBulkUpdateResponse

`func NewRequestTasksBulkUpdateResponse() *RequestTasksBulkUpdateResponse`

NewRequestTasksBulkUpdateResponse instantiates a new RequestTasksBulkUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTasksBulkUpdateResponseWithDefaults

`func NewRequestTasksBulkUpdateResponseWithDefaults() *RequestTasksBulkUpdateResponse`

NewRequestTasksBulkUpdateResponseWithDefaults instantiates a new RequestTasksBulkUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *RequestTasksBulkUpdateResponse) GetTasks() []RequestTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RequestTasksBulkUpdateResponse) GetTasksOk() (*[]RequestTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RequestTasksBulkUpdateResponse) SetTasks(v []RequestTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *RequestTasksBulkUpdateResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


