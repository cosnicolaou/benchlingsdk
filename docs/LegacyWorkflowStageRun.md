# LegacyWorkflowStageRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the legacy workflow stage run was created | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the legacy workflow stage run | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the legacy workflow stage run | [optional] 
**Status** | Pointer to **string** | Status of the legacy workflow stage run | [optional] 

## Methods

### NewLegacyWorkflowStageRun

`func NewLegacyWorkflowStageRun() *LegacyWorkflowStageRun`

NewLegacyWorkflowStageRun instantiates a new LegacyWorkflowStageRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWorkflowStageRunWithDefaults

`func NewLegacyWorkflowStageRunWithDefaults() *LegacyWorkflowStageRun`

NewLegacyWorkflowStageRunWithDefaults instantiates a new LegacyWorkflowStageRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LegacyWorkflowStageRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacyWorkflowStageRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacyWorkflowStageRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacyWorkflowStageRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LegacyWorkflowStageRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyWorkflowStageRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyWorkflowStageRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyWorkflowStageRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LegacyWorkflowStageRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyWorkflowStageRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyWorkflowStageRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyWorkflowStageRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *LegacyWorkflowStageRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegacyWorkflowStageRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegacyWorkflowStageRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegacyWorkflowStageRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


