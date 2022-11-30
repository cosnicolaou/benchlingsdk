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

// ContainerLabels struct for ContainerLabels
type ContainerLabels struct {
	Barcode *string `json:"barcode,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewContainerLabels instantiates a new ContainerLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLabels() *ContainerLabels {
	this := ContainerLabels{}
	return &this
}

// NewContainerLabelsWithDefaults instantiates a new ContainerLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLabelsWithDefaults() *ContainerLabels {
	this := ContainerLabels{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *ContainerLabels) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerLabels) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *ContainerLabels) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *ContainerLabels) SetBarcode(v string) {
	o.Barcode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerLabels) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerLabels) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerLabels) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerLabels) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerLabels) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerLabels) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerLabels) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerLabels) SetName(v string) {
	o.Name = &v
}

func (o ContainerLabels) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableContainerLabels struct {
	value *ContainerLabels
	isSet bool
}

func (v NullableContainerLabels) Get() *ContainerLabels {
	return v.value
}

func (v *NullableContainerLabels) Set(val *ContainerLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLabels(val *ContainerLabels) *NullableContainerLabels {
	return &NullableContainerLabels{value: val, isSet: true}
}

func (v NullableContainerLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


