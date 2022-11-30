# ContainersCheckin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** |  | [optional] 
**ContainerIds** | **[]string** | Array of container IDs. | 

## Methods

### NewContainersCheckin

`func NewContainersCheckin(containerIds []string, ) *ContainersCheckin`

NewContainersCheckin instantiates a new ContainersCheckin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersCheckinWithDefaults

`func NewContainersCheckinWithDefaults() *ContainersCheckin`

NewContainersCheckinWithDefaults instantiates a new ContainersCheckin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *ContainersCheckin) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ContainersCheckin) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ContainersCheckin) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ContainersCheckin) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetContainerIds

`func (o *ContainersCheckin) GetContainerIds() []string`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *ContainersCheckin) GetContainerIdsOk() (*[]string, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *ContainersCheckin) SetContainerIds(v []string)`

SetContainerIds sets ContainerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


