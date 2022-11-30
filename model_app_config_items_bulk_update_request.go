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

// AppConfigItemsBulkUpdateRequest struct for AppConfigItemsBulkUpdateRequest
type AppConfigItemsBulkUpdateRequest struct {
	AppConfigurationItems []AppConfigItemBulkUpdate `json:"appConfigurationItems"`
}

// NewAppConfigItemsBulkUpdateRequest instantiates a new AppConfigItemsBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemsBulkUpdateRequest(appConfigurationItems []AppConfigItemBulkUpdate) *AppConfigItemsBulkUpdateRequest {
	this := AppConfigItemsBulkUpdateRequest{}
	this.AppConfigurationItems = appConfigurationItems
	return &this
}

// NewAppConfigItemsBulkUpdateRequestWithDefaults instantiates a new AppConfigItemsBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemsBulkUpdateRequestWithDefaults() *AppConfigItemsBulkUpdateRequest {
	this := AppConfigItemsBulkUpdateRequest{}
	return &this
}

// GetAppConfigurationItems returns the AppConfigurationItems field value
func (o *AppConfigItemsBulkUpdateRequest) GetAppConfigurationItems() []AppConfigItemBulkUpdate {
	if o == nil {
		var ret []AppConfigItemBulkUpdate
		return ret
	}

	return o.AppConfigurationItems
}

// GetAppConfigurationItemsOk returns a tuple with the AppConfigurationItems field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemsBulkUpdateRequest) GetAppConfigurationItemsOk() ([]AppConfigItemBulkUpdate, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppConfigurationItems, true
}

// SetAppConfigurationItems sets field value
func (o *AppConfigItemsBulkUpdateRequest) SetAppConfigurationItems(v []AppConfigItemBulkUpdate) {
	o.AppConfigurationItems = v
}

func (o AppConfigItemsBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appConfigurationItems"] = o.AppConfigurationItems
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemsBulkUpdateRequest struct {
	value *AppConfigItemsBulkUpdateRequest
	isSet bool
}

func (v NullableAppConfigItemsBulkUpdateRequest) Get() *AppConfigItemsBulkUpdateRequest {
	return v.value
}

func (v *NullableAppConfigItemsBulkUpdateRequest) Set(val *AppConfigItemsBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemsBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemsBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemsBulkUpdateRequest(val *AppConfigItemsBulkUpdateRequest) *NullableAppConfigItemsBulkUpdateRequest {
	return &NullableAppConfigItemsBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableAppConfigItemsBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemsBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

