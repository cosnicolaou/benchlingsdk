# MultipleContainersTransferAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationContainerId** | **string** | ID of container that will be transferred into. | 
**FinalQuantity** | Pointer to [**ContainerQuantity**](ContainerQuantity.md) |  | [optional] 
**FinalVolume** | Pointer to [**DeprecatedContainerVolumeForInput**](DeprecatedContainerVolumeForInput.md) |  | [optional] 
**SourceConcentration** | Pointer to **map[string]interface{}** | Concentration at which to transfer entities or batches. Not applicable for container-to-container transfers (the concentration of the source container’s contents will be used).  | [optional] 

## Methods

### NewMultipleContainersTransferAllOf

`func NewMultipleContainersTransferAllOf(destinationContainerId string, ) *MultipleContainersTransferAllOf`

NewMultipleContainersTransferAllOf instantiates a new MultipleContainersTransferAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleContainersTransferAllOfWithDefaults

`func NewMultipleContainersTransferAllOfWithDefaults() *MultipleContainersTransferAllOf`

NewMultipleContainersTransferAllOfWithDefaults instantiates a new MultipleContainersTransferAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationContainerId

`func (o *MultipleContainersTransferAllOf) GetDestinationContainerId() string`

GetDestinationContainerId returns the DestinationContainerId field if non-nil, zero value otherwise.

### GetDestinationContainerIdOk

`func (o *MultipleContainersTransferAllOf) GetDestinationContainerIdOk() (*string, bool)`

GetDestinationContainerIdOk returns a tuple with the DestinationContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationContainerId

`func (o *MultipleContainersTransferAllOf) SetDestinationContainerId(v string)`

SetDestinationContainerId sets DestinationContainerId field to given value.


### GetFinalQuantity

`func (o *MultipleContainersTransferAllOf) GetFinalQuantity() ContainerQuantity`

GetFinalQuantity returns the FinalQuantity field if non-nil, zero value otherwise.

### GetFinalQuantityOk

`func (o *MultipleContainersTransferAllOf) GetFinalQuantityOk() (*ContainerQuantity, bool)`

GetFinalQuantityOk returns a tuple with the FinalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalQuantity

`func (o *MultipleContainersTransferAllOf) SetFinalQuantity(v ContainerQuantity)`

SetFinalQuantity sets FinalQuantity field to given value.

### HasFinalQuantity

`func (o *MultipleContainersTransferAllOf) HasFinalQuantity() bool`

HasFinalQuantity returns a boolean if a field has been set.

### GetFinalVolume

`func (o *MultipleContainersTransferAllOf) GetFinalVolume() DeprecatedContainerVolumeForInput`

GetFinalVolume returns the FinalVolume field if non-nil, zero value otherwise.

### GetFinalVolumeOk

`func (o *MultipleContainersTransferAllOf) GetFinalVolumeOk() (*DeprecatedContainerVolumeForInput, bool)`

GetFinalVolumeOk returns a tuple with the FinalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalVolume

`func (o *MultipleContainersTransferAllOf) SetFinalVolume(v DeprecatedContainerVolumeForInput)`

SetFinalVolume sets FinalVolume field to given value.

### HasFinalVolume

`func (o *MultipleContainersTransferAllOf) HasFinalVolume() bool`

HasFinalVolume returns a boolean if a field has been set.

### GetSourceConcentration

`func (o *MultipleContainersTransferAllOf) GetSourceConcentration() map[string]interface{}`

GetSourceConcentration returns the SourceConcentration field if non-nil, zero value otherwise.

### GetSourceConcentrationOk

`func (o *MultipleContainersTransferAllOf) GetSourceConcentrationOk() (*map[string]interface{}, bool)`

GetSourceConcentrationOk returns a tuple with the SourceConcentration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConcentration

`func (o *MultipleContainersTransferAllOf) SetSourceConcentration(v map[string]interface{})`

SetSourceConcentration sets SourceConcentration field to given value.

### HasSourceConcentration

`func (o *MultipleContainersTransferAllOf) HasSourceConcentration() bool`

HasSourceConcentration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


