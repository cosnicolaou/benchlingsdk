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

// DropdownSummariesPaginatedList struct for DropdownSummariesPaginatedList
type DropdownSummariesPaginatedList struct {
	Dropdowns []DropdownSummary `json:"dropdowns,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// NewDropdownSummariesPaginatedList instantiates a new DropdownSummariesPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropdownSummariesPaginatedList() *DropdownSummariesPaginatedList {
	this := DropdownSummariesPaginatedList{}
	return &this
}

// NewDropdownSummariesPaginatedListWithDefaults instantiates a new DropdownSummariesPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropdownSummariesPaginatedListWithDefaults() *DropdownSummariesPaginatedList {
	this := DropdownSummariesPaginatedList{}
	return &this
}

// GetDropdowns returns the Dropdowns field value if set, zero value otherwise.
func (o *DropdownSummariesPaginatedList) GetDropdowns() []DropdownSummary {
	if o == nil || isNil(o.Dropdowns) {
		var ret []DropdownSummary
		return ret
	}
	return o.Dropdowns
}

// GetDropdownsOk returns a tuple with the Dropdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropdownSummariesPaginatedList) GetDropdownsOk() ([]DropdownSummary, bool) {
	if o == nil || isNil(o.Dropdowns) {
    return nil, false
	}
	return o.Dropdowns, true
}

// HasDropdowns returns a boolean if a field has been set.
func (o *DropdownSummariesPaginatedList) HasDropdowns() bool {
	if o != nil && !isNil(o.Dropdowns) {
		return true
	}

	return false
}

// SetDropdowns gets a reference to the given []DropdownSummary and assigns it to the Dropdowns field.
func (o *DropdownSummariesPaginatedList) SetDropdowns(v []DropdownSummary) {
	o.Dropdowns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *DropdownSummariesPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropdownSummariesPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *DropdownSummariesPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *DropdownSummariesPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

func (o DropdownSummariesPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Dropdowns) {
		toSerialize["dropdowns"] = o.Dropdowns
	}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableDropdownSummariesPaginatedList struct {
	value *DropdownSummariesPaginatedList
	isSet bool
}

func (v NullableDropdownSummariesPaginatedList) Get() *DropdownSummariesPaginatedList {
	return v.value
}

func (v *NullableDropdownSummariesPaginatedList) Set(val *DropdownSummariesPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableDropdownSummariesPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableDropdownSummariesPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropdownSummariesPaginatedList(val *DropdownSummariesPaginatedList) *NullableDropdownSummariesPaginatedList {
	return &NullableDropdownSummariesPaginatedList{value: val, isSet: true}
}

func (v NullableDropdownSummariesPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropdownSummariesPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


