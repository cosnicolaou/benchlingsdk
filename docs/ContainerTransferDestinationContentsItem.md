# ContainerTransferDestinationContentsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concentration** | Pointer to [**Measurement**](Measurement.md) |  | [optional] 
**EntityId** | **string** |  | 

## Methods

### NewContainerTransferDestinationContentsItem

`func NewContainerTransferDestinationContentsItem(entityId string, ) *ContainerTransferDestinationContentsItem`

NewContainerTransferDestinationContentsItem instantiates a new ContainerTransferDestinationContentsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTransferDestinationContentsItemWithDefaults

`func NewContainerTransferDestinationContentsItemWithDefaults() *ContainerTransferDestinationContentsItem`

NewContainerTransferDestinationContentsItemWithDefaults instantiates a new ContainerTransferDestinationContentsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcentration

`func (o *ContainerTransferDestinationContentsItem) GetConcentration() Measurement`

GetConcentration returns the Concentration field if non-nil, zero value otherwise.

### GetConcentrationOk

`func (o *ContainerTransferDestinationContentsItem) GetConcentrationOk() (*Measurement, bool)`

GetConcentrationOk returns a tuple with the Concentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcentration

`func (o *ContainerTransferDestinationContentsItem) SetConcentration(v Measurement)`

SetConcentration sets Concentration field to given value.

### HasConcentration

`func (o *ContainerTransferDestinationContentsItem) HasConcentration() bool`

HasConcentration returns a boolean if a field has been set.

### GetEntityId

`func (o *ContainerTransferDestinationContentsItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ContainerTransferDestinationContentsItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ContainerTransferDestinationContentsItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


