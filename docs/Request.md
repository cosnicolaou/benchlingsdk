# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | The canonical url of the Request in the API. | [optional] [readonly] 
**Assignees** | Pointer to [**[]OneOfRequestUserAssigneeRequestTeamAssignee**](OneOfRequestUserAssigneeRequestTeamAssignee.md) | Array of assignees | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Date and time the request was created | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**DisplayId** | Pointer to **string** | User-friendly ID of the request | [optional] [readonly] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for the request | [optional] [readonly] 
**ProjectId** | Pointer to **string** | The ID of the project to which the request belongs. | [optional] 
**RequestStatus** | Pointer to [**RequestStatus**](RequestStatus.md) |  | [optional] 
**Requestor** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**SampleGroups** | Pointer to [**[]RequestSampleGroup**](RequestSampleGroup.md) |  | [optional] 
**ScheduledOn** | Pointer to **NullableString** | Date the request is scheduled to be executed on, in YYYY-MM-DD format. | [optional] 
**Schema** | Pointer to [**SchemaSummary**](SchemaSummary.md) |  | [optional] 
**Tasks** | Pointer to [**[]RequestTask**](RequestTask.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the request | [optional] [readonly] 

## Methods

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *Request) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *Request) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *Request) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *Request) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetAssignees

`func (o *Request) GetAssignees() []OneOfRequestUserAssigneeRequestTeamAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *Request) GetAssigneesOk() (*[]OneOfRequestUserAssigneeRequestTeamAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *Request) SetAssignees(v []OneOfRequestUserAssigneeRequestTeamAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *Request) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Request) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Request) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Request) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Request) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Request) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Request) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Request) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Request) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisplayId

`func (o *Request) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Request) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Request) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Request) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetFields

`func (o *Request) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Request) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Request) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Request) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *Request) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Request) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Request) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Request) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *Request) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Request) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Request) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Request) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRequestStatus

`func (o *Request) GetRequestStatus() RequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *Request) GetRequestStatusOk() (*RequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *Request) SetRequestStatus(v RequestStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *Request) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetRequestor

`func (o *Request) GetRequestor() UserSummary`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *Request) GetRequestorOk() (*UserSummary, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *Request) SetRequestor(v UserSummary)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *Request) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetSampleGroups

`func (o *Request) GetSampleGroups() []RequestSampleGroup`

GetSampleGroups returns the SampleGroups field if non-nil, zero value otherwise.

### GetSampleGroupsOk

`func (o *Request) GetSampleGroupsOk() (*[]RequestSampleGroup, bool)`

GetSampleGroupsOk returns a tuple with the SampleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroups

`func (o *Request) SetSampleGroups(v []RequestSampleGroup)`

SetSampleGroups sets SampleGroups field to given value.

### HasSampleGroups

`func (o *Request) HasSampleGroups() bool`

HasSampleGroups returns a boolean if a field has been set.

### GetScheduledOn

`func (o *Request) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *Request) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *Request) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *Request) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### SetScheduledOnNil

`func (o *Request) SetScheduledOnNil(b bool)`

 SetScheduledOnNil sets the value for ScheduledOn to be an explicit nil

### UnsetScheduledOn
`func (o *Request) UnsetScheduledOn()`

UnsetScheduledOn ensures that no value is present for ScheduledOn, not even an explicit nil
### GetSchema

`func (o *Request) GetSchema() SchemaSummary`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Request) GetSchemaOk() (*SchemaSummary, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Request) SetSchema(v SchemaSummary)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Request) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTasks

`func (o *Request) GetTasks() []RequestTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Request) GetTasksOk() (*[]RequestTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Request) SetTasks(v []RequestTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *Request) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetWebURL

`func (o *Request) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Request) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Request) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Request) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


