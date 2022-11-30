# RequestTasksBulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | [**[]RequestTasksBulkCreate**](RequestTasksBulkCreate.md) | The tasks to create | 

## Methods

### NewRequestTasksBulkCreateRequest

`func NewRequestTasksBulkCreateRequest(tasks []RequestTasksBulkCreate, ) *RequestTasksBulkCreateRequest`

NewRequestTasksBulkCreateRequest instantiates a new RequestTasksBulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTasksBulkCreateRequestWithDefaults

`func NewRequestTasksBulkCreateRequestWithDefaults() *RequestTasksBulkCreateRequest`

NewRequestTasksBulkCreateRequestWithDefaults instantiates a new RequestTasksBulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *RequestTasksBulkCreateRequest) GetTasks() []RequestTasksBulkCreate`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RequestTasksBulkCreateRequest) GetTasksOk() (*[]RequestTasksBulkCreate, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RequestTasksBulkCreateRequest) SetTasks(v []RequestTasksBulkCreate)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


