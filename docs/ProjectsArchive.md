# ProjectsArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | **[]string** | A list of project IDs to archive. | 
**Reason** | **string** | The reason for archiving the provided projects. Accepted reasons may differ based on tenant configuration.  | 

## Methods

### NewProjectsArchive

`func NewProjectsArchive(projectIds []string, reason string, ) *ProjectsArchive`

NewProjectsArchive instantiates a new ProjectsArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsArchiveWithDefaults

`func NewProjectsArchiveWithDefaults() *ProjectsArchive`

NewProjectsArchiveWithDefaults instantiates a new ProjectsArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *ProjectsArchive) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ProjectsArchive) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ProjectsArchive) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.


### GetReason

`func (o *ProjectsArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ProjectsArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ProjectsArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


