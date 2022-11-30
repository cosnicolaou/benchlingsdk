# LegacyWorkflowStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | DateTime at which the the legacy workflow stage was created | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the legacy workflow stage | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the legacy workflow stage | [optional] 

## Methods

### NewLegacyWorkflowStage

`func NewLegacyWorkflowStage() *LegacyWorkflowStage`

NewLegacyWorkflowStage instantiates a new LegacyWorkflowStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyWorkflowStageWithDefaults

`func NewLegacyWorkflowStageWithDefaults() *LegacyWorkflowStage`

NewLegacyWorkflowStageWithDefaults instantiates a new LegacyWorkflowStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LegacyWorkflowStage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacyWorkflowStage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacyWorkflowStage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacyWorkflowStage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LegacyWorkflowStage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacyWorkflowStage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacyWorkflowStage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegacyWorkflowStage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LegacyWorkflowStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacyWorkflowStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacyWorkflowStage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LegacyWorkflowStage) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


