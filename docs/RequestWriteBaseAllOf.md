# RequestWriteBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee**](OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee.md) | Array of assignees | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | The request&#39;s fields | [optional] 
**ProjectId** | Pointer to **string** | The ID of the project to which the request belongs. | [optional] 
**RequestorId** | Pointer to **NullableString** | ID of the user making the request. If unspecified, the requestor is the request creator.  | [optional] 
**SampleGroups** | Pointer to [**[]RequestSampleGroupCreate**](RequestSampleGroupCreate.md) |  | [optional] 
**ScheduledOn** | Pointer to **string** | Date the request is scheduled to be executed on, in YYYY-MM-DD format. | [optional] 

## Methods

### NewRequestWriteBaseAllOf

`func NewRequestWriteBaseAllOf() *RequestWriteBaseAllOf`

NewRequestWriteBaseAllOf instantiates a new RequestWriteBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWriteBaseAllOfWithDefaults

`func NewRequestWriteBaseAllOfWithDefaults() *RequestWriteBaseAllOf`

NewRequestWriteBaseAllOfWithDefaults instantiates a new RequestWriteBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *RequestWriteBaseAllOf) GetAssignees() []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *RequestWriteBaseAllOf) GetAssigneesOk() (*[]OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *RequestWriteBaseAllOf) SetAssignees(v []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *RequestWriteBaseAllOf) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetFields

`func (o *RequestWriteBaseAllOf) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestWriteBaseAllOf) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestWriteBaseAllOf) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RequestWriteBaseAllOf) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetProjectId

`func (o *RequestWriteBaseAllOf) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RequestWriteBaseAllOf) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RequestWriteBaseAllOf) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RequestWriteBaseAllOf) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRequestorId

`func (o *RequestWriteBaseAllOf) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *RequestWriteBaseAllOf) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *RequestWriteBaseAllOf) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *RequestWriteBaseAllOf) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### SetRequestorIdNil

`func (o *RequestWriteBaseAllOf) SetRequestorIdNil(b bool)`

 SetRequestorIdNil sets the value for RequestorId to be an explicit nil

### UnsetRequestorId
`func (o *RequestWriteBaseAllOf) UnsetRequestorId()`

UnsetRequestorId ensures that no value is present for RequestorId, not even an explicit nil
### GetSampleGroups

`func (o *RequestWriteBaseAllOf) GetSampleGroups() []RequestSampleGroupCreate`

GetSampleGroups returns the SampleGroups field if non-nil, zero value otherwise.

### GetSampleGroupsOk

`func (o *RequestWriteBaseAllOf) GetSampleGroupsOk() (*[]RequestSampleGroupCreate, bool)`

GetSampleGroupsOk returns a tuple with the SampleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroups

`func (o *RequestWriteBaseAllOf) SetSampleGroups(v []RequestSampleGroupCreate)`

SetSampleGroups sets SampleGroups field to given value.

### HasSampleGroups

`func (o *RequestWriteBaseAllOf) HasSampleGroups() bool`

HasSampleGroups returns a boolean if a field has been set.

### GetScheduledOn

`func (o *RequestWriteBaseAllOf) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *RequestWriteBaseAllOf) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *RequestWriteBaseAllOf) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *RequestWriteBaseAllOf) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


