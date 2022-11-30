# RequestCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to [**[]OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee**](OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee.md) | Array of assignees | [optional] 
**Fields** | [**ModelMap**](map.md) | The request&#39;s fields | 
**ProjectId** | **string** | The ID of the project to which the request belongs. | 
**RequestorId** | Pointer to **NullableString** | ID of the user making the request. If unspecified, the requestor is the request creator.  | [optional] 
**SampleGroups** | [**[]RequestSampleGroupCreate**](RequestSampleGroupCreate.md) |  | 
**ScheduledOn** | Pointer to **string** | Date the request is scheduled to be executed on, in YYYY-MM-DD format. | [optional] 
**SchemaId** | **string** | ID of the request&#39;s schema. | 

## Methods

### NewRequestCreate

`func NewRequestCreate(fields ModelMap, projectId string, sampleGroups []RequestSampleGroupCreate, schemaId string, ) *RequestCreate`

NewRequestCreate instantiates a new RequestCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateWithDefaults

`func NewRequestCreateWithDefaults() *RequestCreate`

NewRequestCreateWithDefaults instantiates a new RequestCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *RequestCreate) GetAssignees() []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *RequestCreate) GetAssigneesOk() (*[]OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *RequestCreate) SetAssignees(v []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *RequestCreate) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetFields

`func (o *RequestCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RequestCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RequestCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.


### GetProjectId

`func (o *RequestCreate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RequestCreate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RequestCreate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetRequestorId

`func (o *RequestCreate) GetRequestorId() string`

GetRequestorId returns the RequestorId field if non-nil, zero value otherwise.

### GetRequestorIdOk

`func (o *RequestCreate) GetRequestorIdOk() (*string, bool)`

GetRequestorIdOk returns a tuple with the RequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorId

`func (o *RequestCreate) SetRequestorId(v string)`

SetRequestorId sets RequestorId field to given value.

### HasRequestorId

`func (o *RequestCreate) HasRequestorId() bool`

HasRequestorId returns a boolean if a field has been set.

### SetRequestorIdNil

`func (o *RequestCreate) SetRequestorIdNil(b bool)`

 SetRequestorIdNil sets the value for RequestorId to be an explicit nil

### UnsetRequestorId
`func (o *RequestCreate) UnsetRequestorId()`

UnsetRequestorId ensures that no value is present for RequestorId, not even an explicit nil
### GetSampleGroups

`func (o *RequestCreate) GetSampleGroups() []RequestSampleGroupCreate`

GetSampleGroups returns the SampleGroups field if non-nil, zero value otherwise.

### GetSampleGroupsOk

`func (o *RequestCreate) GetSampleGroupsOk() (*[]RequestSampleGroupCreate, bool)`

GetSampleGroupsOk returns a tuple with the SampleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleGroups

`func (o *RequestCreate) SetSampleGroups(v []RequestSampleGroupCreate)`

SetSampleGroups sets SampleGroups field to given value.


### GetScheduledOn

`func (o *RequestCreate) GetScheduledOn() string`

GetScheduledOn returns the ScheduledOn field if non-nil, zero value otherwise.

### GetScheduledOnOk

`func (o *RequestCreate) GetScheduledOnOk() (*string, bool)`

GetScheduledOnOk returns a tuple with the ScheduledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledOn

`func (o *RequestCreate) SetScheduledOn(v string)`

SetScheduledOn sets ScheduledOn field to given value.

### HasScheduledOn

`func (o *RequestCreate) HasScheduledOn() bool`

HasScheduledOn returns a boolean if a field has been set.

### GetSchemaId

`func (o *RequestCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *RequestCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *RequestCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


