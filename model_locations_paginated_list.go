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

// LocationsPaginatedList struct for LocationsPaginatedList
type LocationsPaginatedList struct {
	Locations []Location `json:"locations,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// NewLocationsPaginatedList instantiates a new LocationsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationsPaginatedList() *LocationsPaginatedList {
	this := LocationsPaginatedList{}
	return &this
}

// NewLocationsPaginatedListWithDefaults instantiates a new LocationsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsPaginatedListWithDefaults() *LocationsPaginatedList {
	this := LocationsPaginatedList{}
	return &this
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *LocationsPaginatedList) GetLocations() []Location {
	if o == nil || isNil(o.Locations) {
		var ret []Location
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsPaginatedList) GetLocationsOk() ([]Location, bool) {
	if o == nil || isNil(o.Locations) {
    return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *LocationsPaginatedList) HasLocations() bool {
	if o != nil && !isNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []Location and assigns it to the Locations field.
func (o *LocationsPaginatedList) SetLocations(v []Location) {
	o.Locations = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *LocationsPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *LocationsPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *LocationsPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

func (o LocationsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableLocationsPaginatedList struct {
	value *LocationsPaginatedList
	isSet bool
}

func (v NullableLocationsPaginatedList) Get() *LocationsPaginatedList {
	return v.value
}

func (v *NullableLocationsPaginatedList) Set(val *LocationsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationsPaginatedList(val *LocationsPaginatedList) *NullableLocationsPaginatedList {
	return &NullableLocationsPaginatedList{value: val, isSet: true}
}

func (v NullableLocationsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


