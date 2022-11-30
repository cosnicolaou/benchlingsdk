# RequestAllOf

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

### NewRequestAllOf

`func NewRequestAllOf() *RequestAllOf`

NewRequestAllOf instantiates a new RequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAllOfWithDefaults

`func NewRequestAllOfWithDefaults() *RequestAllOf`

NewRequestAllOfWithDefaults instantiates a new RequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *RequestAllOf) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *RequestAllOf) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *RequestAllOf) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *RequestAllOf) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetAssignees

`func (o *RequestAllOf) GetAssignees() []OneOfRequestUserAssigneeRequestTeamAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *RequestAllOf) GetAssigneesOk() (*[]OneOfRequestUserAssigneeRequestTeamAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *RequestAllOf) SetAssignees(v []OneOfRequestUserAssigneeRequestTeamAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *RequestAllOf) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RequestAllOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestAllOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestAllOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *RequestAllOf) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RequestAllOf) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RequestAllOf) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RequestAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisplayId

`func (o *RequestAllOf) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *RequestAllOf) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *RequestAllOf) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *RequestAllOf) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetFields

`func (o *RequestAllOf) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestAllOf) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestAllOf) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *RequestAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *RequestAllOf) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RequestAllOf) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RequestAllOf) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RequestAllOf) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRequestStatus

`func (o *RequestAllOf) GetRequestStatus() RequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *RequestAllOf) GetRequestStatusOk() (*RequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *RequestAllOf) SetRequestStatus(v RequestStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *RequestAllOf) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetRequestor

`func (o *RequestAllOf) GetRequestor() UserSummary`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *RequestAllOf) GetRequestorOk() (*UserSummary, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *RequestAllOf) SetRequestor(v UserSummary)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *RequestAllOf) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetSampleGroups

`func (o *RequestAllOf) GetSampleGroups() []RequestSampleGroup`

GetSampleGroups returns the SampleGroups field if non-nil, zero value otherwise.

### GetSampleGroupsOk

`func (o *RequestAllOf) GetSampleGroupsOk() (*[]RequestSampleGroup, bool)`

GetSampleGroupsOk returns a tuple with the SampleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroups

`func (o *RequestAllOf) SetSampleGroups(v []RequestSampleGroup)`

SetSampleGroups sets SampleGroups field to given value.

### HasSampleGroups

`func (o *RequestAllOf) HasSampleGroups() bool`

HasSampleGroups returns a boolean if a field has been set.

### GetScheduledOn

`func (o *RequestAllOf) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *RequestAllOf) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *RequestAllOf) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *RequestAllOf) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### SetScheduledOnNil

`func (o *RequestAllOf) SetScheduledOnNil(b bool)`

 SetScheduledOnNil sets the value for ScheduledOn to be an explicit nil

### UnsetScheduledOn
`func (o *RequestAllOf) UnsetScheduledOn()`

UnsetScheduledOn ensures that no value is present for ScheduledOn, not even an explicit nil
### GetSchema

`func (o *RequestAllOf) GetSchema() SchemaSummary`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RequestAllOf) GetSchemaOk() (*SchemaSummary, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RequestAllOf) SetSchema(v SchemaSummary)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RequestAllOf) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTasks

`func (o *RequestAllOf) GetTasks() []RequestTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RequestAllOf) GetTasksOk() (*[]RequestTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RequestAllOf) SetTasks(v []RequestTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *RequestAllOf) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetWebURL

`func (o *RequestAllOf) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *RequestAllOf) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *RequestAllOf) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *RequestAllOf) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


