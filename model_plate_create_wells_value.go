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

// PlateCreateWellsValue struct for PlateCreateWellsValue
type PlateCreateWellsValue struct {
	Barcode *string `json:"barcode,omitempty"`
}

// NewPlateCreateWellsValue instantiates a new PlateCreateWellsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlateCreateWellsValue() *PlateCreateWellsValue {
	this := PlateCreateWellsValue{}
	return &this
}

// NewPlateCreateWellsValueWithDefaults instantiates a new PlateCreateWellsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlateCreateWellsValueWithDefaults() *PlateCreateWellsValue {
	this := PlateCreateWellsValue{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *PlateCreateWellsValue) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateCreateWellsValue) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *PlateCreateWellsValue) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *PlateCreateWellsValue) SetBarcode(v string) {
	o.Barcode = &v
}

func (o PlateCreateWellsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	return json.Marshal(toSerialize)
}

type NullablePlateCreateWellsValue struct {
	value *PlateCreateWellsValue
	isSet bool
}

func (v NullablePlateCreateWellsValue) Get() *PlateCreateWellsValue {
	return v.value
}

func (v *NullablePlateCreateWellsValue) Set(val *PlateCreateWellsValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePlateCreateWellsValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePlateCreateWellsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlateCreateWellsValue(val *PlateCreateWellsValue) *NullablePlateCreateWellsValue {
	return &NullablePlateCreateWellsValue{value: val, isSet: true}
}

func (v NullablePlateCreateWellsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlateCreateWellsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

