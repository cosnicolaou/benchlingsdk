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

// AppConfigurationPaginatedList struct for AppConfigurationPaginatedList
type AppConfigurationPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	AppConfigurationItems []AppConfigItem `json:"appConfigurationItems,omitempty"`
}

// NewAppConfigurationPaginatedList instantiates a new AppConfigurationPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigurationPaginatedList() *AppConfigurationPaginatedList {
	this := AppConfigurationPaginatedList{}
	return &this
}

// NewAppConfigurationPaginatedListWithDefaults instantiates a new AppConfigurationPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigurationPaginatedListWithDefaults() *AppConfigurationPaginatedList {
	this := AppConfigurationPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *AppConfigurationPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigurationPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *AppConfigurationPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *AppConfigurationPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAppConfigurationItems returns the AppConfigurationItems field value if set, zero value otherwise.
func (o *AppConfigurationPaginatedList) GetAppConfigurationItems() []AppConfigItem {
	if o == nil || isNil(o.AppConfigurationItems) {
		var ret []AppConfigItem
		return ret
	}
	return o.AppConfigurationItems
}

// GetAppConfigurationItemsOk returns a tuple with the AppConfigurationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigurationPaginatedList) GetAppConfigurationItemsOk() ([]AppConfigItem, bool) {
	if o == nil || isNil(o.AppConfigurationItems) {
    return nil, false
	}
	return o.AppConfigurationItems, true
}

// HasAppConfigurationItems returns a boolean if a field has been set.
func (o *AppConfigurationPaginatedList) HasAppConfigurationItems() bool {
	if o != nil && !isNil(o.AppConfigurationItems) {
		return true
	}

	return false
}

// SetAppConfigurationItems gets a reference to the given []AppConfigItem and assigns it to the AppConfigurationItems field.
func (o *AppConfigurationPaginatedList) SetAppConfigurationItems(v []AppConfigItem) {
	o.AppConfigurationItems = v
}

func (o AppConfigurationPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.AppConfigurationItems) {
		toSerialize["appConfigurationItems"] = o.AppConfigurationItems
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigurationPaginatedList struct {
	value *AppConfigurationPaginatedList
	isSet bool
}

func (v NullableAppConfigurationPaginatedList) Get() *AppConfigurationPaginatedList {
	return v.value
}

func (v *NullableAppConfigurationPaginatedList) Set(val *AppConfigurationPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigurationPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigurationPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigurationPaginatedList(val *AppConfigurationPaginatedList) *NullableAppConfigurationPaginatedList {
	return &NullableAppConfigurationPaginatedList{value: val, isSet: true}
}

func (v NullableAppConfigurationPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigurationPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


