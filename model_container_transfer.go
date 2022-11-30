/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
)

// ContainerTransfer struct for ContainerTransfer
type ContainerTransfer struct {
	// ID of the batch that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId. 
	SourceBatchId *string `json:"sourceBatchId,omitempty"`
	// ID of the container that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId. 
	SourceContainerId *string `json:"sourceContainerId,omitempty"`
	// ID of the entity that will be transferred in. Must specify one of sourceEntityId, sourceBatchId, or sourceContainerId. 
	SourceEntityId *string `json:"sourceEntityId,omitempty"`
	TransferQuantity *ContainerTransferBaseTransferQuantity `json:"transferQuantity,omitempty"`
	TransferVolume *ContainerTransferBaseTransferVolume `json:"transferVolume,omitempty"`
	// This represents what the contents of the destination container should look like post-transfer. 
	DestinationContents []ContainerTransferDestinationContentsItem `json:"destinationContents"`
	// This represents the desired final quantity of the destination container, post-dilution. If you don't want to dilute the destination container, you can omit this parameter. Supports mass, volume, and other quantities. 
	DestinationQuantity *ContainerQuantity `json:"destinationQuantity,omitempty"`
	// Deprecated - use destinationQuantity instead. 
	DestinationVolume *DeprecatedContainerVolumeForInput `json:"destinationVolume,omitempty"`
}

// NewContainerTransfer instantiates a new ContainerTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTransfer(destinationContents []ContainerTransferDestinationContentsItem) *ContainerTransfer {
	this := ContainerTransfer{}
	this.DestinationContents = destinationContents
	return &this
}

// NewContainerTransferWithDefaults instantiates a new ContainerTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTransferWithDefaults() *ContainerTransfer {
	this := ContainerTransfer{}
	return &this
}

// GetSourceBatchId returns the SourceBatchId field value if set, zero value otherwise.
func (o *ContainerTransfer) GetSourceBatchId() string {
	if o == nil || isNil(o.SourceBatchId) {
		var ret string
		return ret
	}
	return *o.SourceBatchId
}

// GetSourceBatchIdOk returns a tuple with the SourceBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetSourceBatchIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceBatchId) {
    return nil, false
	}
	return o.SourceBatchId, true
}

// HasSourceBatchId returns a boolean if a field has been set.
func (o *ContainerTransfer) HasSourceBatchId() bool {
	if o != nil && !isNil(o.SourceBatchId) {
		return true
	}

	return false
}

// SetSourceBatchId gets a reference to the given string and assigns it to the SourceBatchId field.
func (o *ContainerTransfer) SetSourceBatchId(v string) {
	o.SourceBatchId = &v
}

// GetSourceContainerId returns the SourceContainerId field value if set, zero value otherwise.
func (o *ContainerTransfer) GetSourceContainerId() string {
	if o == nil || isNil(o.SourceContainerId) {
		var ret string
		return ret
	}
	return *o.SourceContainerId
}

// GetSourceContainerIdOk returns a tuple with the SourceContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetSourceContainerIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceContainerId) {
    return nil, false
	}
	return o.SourceContainerId, true
}

// HasSourceContainerId returns a boolean if a field has been set.
func (o *ContainerTransfer) HasSourceContainerId() bool {
	if o != nil && !isNil(o.SourceContainerId) {
		return true
	}

	return false
}

// SetSourceContainerId gets a reference to the given string and assigns it to the SourceContainerId field.
func (o *ContainerTransfer) SetSourceContainerId(v string) {
	o.SourceContainerId = &v
}

// GetSourceEntityId returns the SourceEntityId field value if set, zero value otherwise.
func (o *ContainerTransfer) GetSourceEntityId() string {
	if o == nil || isNil(o.SourceEntityId) {
		var ret string
		return ret
	}
	return *o.SourceEntityId
}

// GetSourceEntityIdOk returns a tuple with the SourceEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetSourceEntityIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceEntityId) {
    return nil, false
	}
	return o.SourceEntityId, true
}

// HasSourceEntityId returns a boolean if a field has been set.
func (o *ContainerTransfer) HasSourceEntityId() bool {
	if o != nil && !isNil(o.SourceEntityId) {
		return true
	}

	return false
}

// SetSourceEntityId gets a reference to the given string and assigns it to the SourceEntityId field.
func (o *ContainerTransfer) SetSourceEntityId(v string) {
	o.SourceEntityId = &v
}

// GetTransferQuantity returns the TransferQuantity field value if set, zero value otherwise.
func (o *ContainerTransfer) GetTransferQuantity() ContainerTransferBaseTransferQuantity {
	if o == nil || isNil(o.TransferQuantity) {
		var ret ContainerTransferBaseTransferQuantity
		return ret
	}
	return *o.TransferQuantity
}

// GetTransferQuantityOk returns a tuple with the TransferQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetTransferQuantityOk() (*ContainerTransferBaseTransferQuantity, bool) {
	if o == nil || isNil(o.TransferQuantity) {
    return nil, false
	}
	return o.TransferQuantity, true
}

// HasTransferQuantity returns a boolean if a field has been set.
func (o *ContainerTransfer) HasTransferQuantity() bool {
	if o != nil && !isNil(o.TransferQuantity) {
		return true
	}

	return false
}

// SetTransferQuantity gets a reference to the given ContainerTransferBaseTransferQuantity and assigns it to the TransferQuantity field.
func (o *ContainerTransfer) SetTransferQuantity(v ContainerTransferBaseTransferQuantity) {
	o.TransferQuantity = &v
}

// GetTransferVolume returns the TransferVolume field value if set, zero value otherwise.
func (o *ContainerTransfer) GetTransferVolume() ContainerTransferBaseTransferVolume {
	if o == nil || isNil(o.TransferVolume) {
		var ret ContainerTransferBaseTransferVolume
		return ret
	}
	return *o.TransferVolume
}

// GetTransferVolumeOk returns a tuple with the TransferVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetTransferVolumeOk() (*ContainerTransferBaseTransferVolume, bool) {
	if o == nil || isNil(o.TransferVolume) {
    return nil, false
	}
	return o.TransferVolume, true
}

// HasTransferVolume returns a boolean if a field has been set.
func (o *ContainerTransfer) HasTransferVolume() bool {
	if o != nil && !isNil(o.TransferVolume) {
		return true
	}

	return false
}

// SetTransferVolume gets a reference to the given ContainerTransferBaseTransferVolume and assigns it to the TransferVolume field.
func (o *ContainerTransfer) SetTransferVolume(v ContainerTransferBaseTransferVolume) {
	o.TransferVolume = &v
}

// GetDestinationContents returns the DestinationContents field value
func (o *ContainerTransfer) GetDestinationContents() []ContainerTransferDestinationContentsItem {
	if o == nil {
		var ret []ContainerTransferDestinationContentsItem
		return ret
	}

	return o.DestinationContents
}

// GetDestinationContentsOk returns a tuple with the DestinationContents field value
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetDestinationContentsOk() ([]ContainerTransferDestinationContentsItem, bool) {
	if o == nil {
    return nil, false
	}
	return o.DestinationContents, true
}

// SetDestinationContents sets field value
func (o *ContainerTransfer) SetDestinationContents(v []ContainerTransferDestinationContentsItem) {
	o.DestinationContents = v
}

// GetDestinationQuantity returns the DestinationQuantity field value if set, zero value otherwise.
func (o *ContainerTransfer) GetDestinationQuantity() ContainerQuantity {
	if o == nil || isNil(o.DestinationQuantity) {
		var ret ContainerQuantity
		return ret
	}
	return *o.DestinationQuantity
}

// GetDestinationQuantityOk returns a tuple with the DestinationQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetDestinationQuantityOk() (*ContainerQuantity, bool) {
	if o == nil || isNil(o.DestinationQuantity) {
    return nil, false
	}
	return o.DestinationQuantity, true
}

// HasDestinationQuantity returns a boolean if a field has been set.
func (o *ContainerTransfer) HasDestinationQuantity() bool {
	if o != nil && !isNil(o.DestinationQuantity) {
		return true
	}

	return false
}

// SetDestinationQuantity gets a reference to the given ContainerQuantity and assigns it to the DestinationQuantity field.
func (o *ContainerTransfer) SetDestinationQuantity(v ContainerQuantity) {
	o.DestinationQuantity = &v
}

// GetDestinationVolume returns the DestinationVolume field value if set, zero value otherwise.
func (o *ContainerTransfer) GetDestinationVolume() DeprecatedContainerVolumeForInput {
	if o == nil || isNil(o.DestinationVolume) {
		var ret DeprecatedContainerVolumeForInput
		return ret
	}
	return *o.DestinationVolume
}

// GetDestinationVolumeOk returns a tuple with the DestinationVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTransfer) GetDestinationVolumeOk() (*DeprecatedContainerVolumeForInput, bool) {
	if o == nil || isNil(o.DestinationVolume) {
    return nil, false
	}
	return o.DestinationVolume, true
}

// HasDestinationVolume returns a boolean if a field has been set.
func (o *ContainerTransfer) HasDestinationVolume() bool {
	if o != nil && !isNil(o.DestinationVolume) {
		return true
	}

	return false
}

// SetDestinationVolume gets a reference to the given DeprecatedContainerVolumeForInput and assigns it to the DestinationVolume field.
func (o *ContainerTransfer) SetDestinationVolume(v DeprecatedContainerVolumeForInput) {
	o.DestinationVolume = &v
}

func (o ContainerTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceBatchId) {
		toSerialize["sourceBatchId"] = o.SourceBatchId
	}
	if !isNil(o.SourceContainerId) {
		toSerialize["sourceContainerId"] = o.SourceContainerId
	}
	if !isNil(o.SourceEntityId) {
		toSerialize["sourceEntityId"] = o.SourceEntityId
	}
	if !isNil(o.TransferQuantity) {
		toSerialize["transferQuantity"] = o.TransferQuantity
	}
	if !isNil(o.TransferVolume) {
		toSerialize["transferVolume"] = o.TransferVolume
	}
	if true {
		toSerialize["destinationContents"] = o.DestinationContents
	}
	if !isNil(o.DestinationQuantity) {
		toSerialize["destinationQuantity"] = o.DestinationQuantity
	}
	if !isNil(o.DestinationVolume) {
		toSerialize["destinationVolume"] = o.DestinationVolume
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTransfer struct {
	value *ContainerTransfer
	isSet bool
}

func (v NullableContainerTransfer) Get() *ContainerTransfer {
	return v.value
}

func (v *NullableContainerTransfer) Set(val *ContainerTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTransfer(val *ContainerTransfer) *NullableContainerTransfer {
	return &NullableContainerTransfer{value: val, isSet: true}
}

func (v NullableContainerTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

