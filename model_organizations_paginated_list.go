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

// OrganizationsPaginatedList struct for OrganizationsPaginatedList
type OrganizationsPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
}

// NewOrganizationsPaginatedList instantiates a new OrganizationsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsPaginatedList() *OrganizationsPaginatedList {
	this := OrganizationsPaginatedList{}
	return &this
}

// NewOrganizationsPaginatedListWithDefaults instantiates a new OrganizationsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsPaginatedListWithDefaults() *OrganizationsPaginatedList {
	this := OrganizationsPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OrganizationsPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OrganizationsPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OrganizationsPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *OrganizationsPaginatedList) GetOrganizations() []Organization {
	if o == nil || isNil(o.Organizations) {
		var ret []Organization
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsPaginatedList) GetOrganizationsOk() ([]Organization, bool) {
	if o == nil || isNil(o.Organizations) {
    return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *OrganizationsPaginatedList) HasOrganizations() bool {
	if o != nil && !isNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []Organization and assigns it to the Organizations field.
func (o *OrganizationsPaginatedList) SetOrganizations(v []Organization) {
	o.Organizations = v
}

func (o OrganizationsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsPaginatedList struct {
	value *OrganizationsPaginatedList
	isSet bool
}

func (v NullableOrganizationsPaginatedList) Get() *OrganizationsPaginatedList {
	return v.value
}

func (v *NullableOrganizationsPaginatedList) Set(val *OrganizationsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsPaginatedList(val *OrganizationsPaginatedList) *NullableOrganizationsPaginatedList {
	return &NullableOrganizationsPaginatedList{value: val, isSet: true}
}

func (v NullableOrganizationsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


