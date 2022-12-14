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

// PlatesArchive struct for PlatesArchive
type PlatesArchive struct {
	// Array of plate IDs
	PlateIds []string `json:"plateIds"`
	// Reason that plates are being archived. 
	Reason string `json:"reason"`
	// Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items. 
	ShouldRemoveBarcodes bool `json:"shouldRemoveBarcodes"`
}

// NewPlatesArchive instantiates a new PlatesArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatesArchive(plateIds []string, reason string, shouldRemoveBarcodes bool) *PlatesArchive {
	this := PlatesArchive{}
	this.PlateIds = plateIds
	this.Reason = reason
	this.ShouldRemoveBarcodes = shouldRemoveBarcodes
	return &this
}

// NewPlatesArchiveWithDefaults instantiates a new PlatesArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatesArchiveWithDefaults() *PlatesArchive {
	this := PlatesArchive{}
	return &this
}

// GetPlateIds returns the PlateIds field value
func (o *PlatesArchive) GetPlateIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PlateIds
}

// GetPlateIdsOk returns a tuple with the PlateIds field value
// and a boolean to check if the value has been set.
func (o *PlatesArchive) GetPlateIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PlateIds, true
}

// SetPlateIds sets field value
func (o *PlatesArchive) SetPlateIds(v []string) {
	o.PlateIds = v
}

// GetReason returns the Reason field value
func (o *PlatesArchive) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PlatesArchive) GetReasonOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PlatesArchive) SetReason(v string) {
	o.Reason = v
}

// GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field value
func (o *PlatesArchive) GetShouldRemoveBarcodes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldRemoveBarcodes
}

// GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field value
// and a boolean to check if the value has been set.
func (o *PlatesArchive) GetShouldRemoveBarcodesOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ShouldRemoveBarcodes, true
}

// SetShouldRemoveBarcodes sets field value
func (o *PlatesArchive) SetShouldRemoveBarcodes(v bool) {
	o.ShouldRemoveBarcodes = v
}

func (o PlatesArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plateIds"] = o.PlateIds
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["shouldRemoveBarcodes"] = o.ShouldRemoveBarcodes
	}
	return json.Marshal(toSerialize)
}

type NullablePlatesArchive struct {
	value *PlatesArchive
	isSet bool
}

func (v NullablePlatesArchive) Get() *PlatesArchive {
	return v.value
}

func (v *NullablePlatesArchive) Set(val *PlatesArchive) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatesArchive) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatesArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatesArchive(val *PlatesArchive) *NullablePlatesArchive {
	return &NullablePlatesArchive{value: val, isSet: true}
}

func (v NullablePlatesArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatesArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


