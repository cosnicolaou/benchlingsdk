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

// LocationsArchive struct for LocationsArchive
type LocationsArchive struct {
	// Array of location IDs
	LocationIds []string `json:"locationIds"`
	// Reason that locations are being archived. 
	Reason string `json:"reason"`
	// Remove barcodes. Removing barcodes from archived storage that contain items will also remove barcodes from the contained items. 
	ShouldRemoveBarcodes *bool `json:"shouldRemoveBarcodes,omitempty"`
}

// NewLocationsArchive instantiates a new LocationsArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationsArchive(locationIds []string, reason string) *LocationsArchive {
	this := LocationsArchive{}
	this.LocationIds = locationIds
	this.Reason = reason
	return &this
}

// NewLocationsArchiveWithDefaults instantiates a new LocationsArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsArchiveWithDefaults() *LocationsArchive {
	this := LocationsArchive{}
	return &this
}

// GetLocationIds returns the LocationIds field value
func (o *LocationsArchive) GetLocationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value
// and a boolean to check if the value has been set.
func (o *LocationsArchive) GetLocationIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LocationIds, true
}

// SetLocationIds sets field value
func (o *LocationsArchive) SetLocationIds(v []string) {
	o.LocationIds = v
}

// GetReason returns the Reason field value
func (o *LocationsArchive) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *LocationsArchive) GetReasonOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *LocationsArchive) SetReason(v string) {
	o.Reason = v
}

// GetShouldRemoveBarcodes returns the ShouldRemoveBarcodes field value if set, zero value otherwise.
func (o *LocationsArchive) GetShouldRemoveBarcodes() bool {
	if o == nil || isNil(o.ShouldRemoveBarcodes) {
		var ret bool
		return ret
	}
	return *o.ShouldRemoveBarcodes
}

// GetShouldRemoveBarcodesOk returns a tuple with the ShouldRemoveBarcodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsArchive) GetShouldRemoveBarcodesOk() (*bool, bool) {
	if o == nil || isNil(o.ShouldRemoveBarcodes) {
    return nil, false
	}
	return o.ShouldRemoveBarcodes, true
}

// HasShouldRemoveBarcodes returns a boolean if a field has been set.
func (o *LocationsArchive) HasShouldRemoveBarcodes() bool {
	if o != nil && !isNil(o.ShouldRemoveBarcodes) {
		return true
	}

	return false
}

// SetShouldRemoveBarcodes gets a reference to the given bool and assigns it to the ShouldRemoveBarcodes field.
func (o *LocationsArchive) SetShouldRemoveBarcodes(v bool) {
	o.ShouldRemoveBarcodes = &v
}

func (o LocationsArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locationIds"] = o.LocationIds
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.ShouldRemoveBarcodes) {
		toSerialize["shouldRemoveBarcodes"] = o.ShouldRemoveBarcodes
	}
	return json.Marshal(toSerialize)
}

type NullableLocationsArchive struct {
	value *LocationsArchive
	isSet bool
}

func (v NullableLocationsArchive) Get() *LocationsArchive {
	return v.value
}

func (v *NullableLocationsArchive) Set(val *LocationsArchive) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationsArchive) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationsArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationsArchive(val *LocationsArchive) *NullableLocationsArchive {
	return &NullableLocationsArchive{value: val, isSet: true}
}

func (v NullableLocationsArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationsArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


