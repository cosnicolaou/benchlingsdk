# LegacyWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the legacy workflow was created | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description of the legacy workflow | [optional] 
**DisplayId** | Pointer to **string** | User-friendly ID of the legacy workflow | [optional] 
**Id** | Pointer to **string** | ID of the legacy workflow | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the legacy workflow | [optional] 
**ProjectId** | Pointer to **string** | ID of the project that contains the legacy workflow | [optional] 

## Methods

### NewLegacyWorkflow

`func NewLegacyWorkflow() *LegacyWorkflow`

NewLegacyWorkflow instantiates a new LegacyWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWorkflowWithDefaults

`func NewLegacyWorkflowWithDefaults() *LegacyWorkflow`

NewLegacyWorkflowWithDefaults instantiates a new LegacyWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LegacyWorkflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacyWorkflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacyWorkflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacyWorkflow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *LegacyWorkflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LegacyWorkflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LegacyWorkflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LegacyWorkflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LegacyWorkflow) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LegacyWorkflow) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayId

`func (o *LegacyWorkflow) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *LegacyWorkflow) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *LegacyWorkflow) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *LegacyWorkflow) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetId

`func (o *LegacyWorkflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyWorkflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyWorkflow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyWorkflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LegacyWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *LegacyWorkflow) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LegacyWorkflow) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LegacyWorkflow) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LegacyWorkflow) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


