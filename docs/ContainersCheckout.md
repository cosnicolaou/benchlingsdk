# ContainersCheckout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeId** | **string** | User or Team API ID. | 
**Comment** | Pointer to **string** |  | [optional] 
**ContainerIds** | **[]string** | Array of container IDs. | 

## Methods

### NewContainersCheckout

`func NewContainersCheckout(assigneeId string, containerIds []string, ) *ContainersCheckout`

NewContainersCheckout instantiates a new ContainersCheckout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersCheckoutWithDefaults

`func NewContainersCheckoutWithDefaults() *ContainersCheckout`

NewContainersCheckoutWithDefaults instantiates a new ContainersCheckout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeId

`func (o *ContainersCheckout) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *ContainersCheckout) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *ContainersCheckout) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.


### GetComment

`func (o *ContainersCheckout) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ContainersCheckout) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ContainersCheckout) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ContainersCheckout) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainerIds

`func (o *ContainersCheckout) GetContainerIds() []string`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *ContainersCheckout) GetContainerIdsOk() (*[]string, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *ContainersCheckout) SetContainerIds(v []string)`

SetContainerIds sets ContainerIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


