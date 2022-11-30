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

// AppConfigItemsBulkCreateRequest struct for AppConfigItemsBulkCreateRequest
type AppConfigItemsBulkCreateRequest struct {
	AppConfigurationItems []AppConfigItemCreate `json:"appConfigurationItems"`
}

// NewAppConfigItemsBulkCreateRequest instantiates a new AppConfigItemsBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemsBulkCreateRequest(appConfigurationItems []AppConfigItemCreate) *AppConfigItemsBulkCreateRequest {
	this := AppConfigItemsBulkCreateRequest{}
	this.AppConfigurationItems = appConfigurationItems
	return &this
}

// NewAppConfigItemsBulkCreateRequestWithDefaults instantiates a new AppConfigItemsBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemsBulkCreateRequestWithDefaults() *AppConfigItemsBulkCreateRequest {
	this := AppConfigItemsBulkCreateRequest{}
	return &this
}

// GetAppConfigurationItems returns the AppConfigurationItems field value
func (o *AppConfigItemsBulkCreateRequest) GetAppConfigurationItems() []AppConfigItemCreate {
	if o == nil {
		var ret []AppConfigItemCreate
		return ret
	}

	return o.AppConfigurationItems
}

// GetAppConfigurationItemsOk returns a tuple with the AppConfigurationItems field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemsBulkCreateRequest) GetAppConfigurationItemsOk() ([]AppConfigItemCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppConfigurationItems, true
}

// SetAppConfigurationItems sets field value
func (o *AppConfigItemsBulkCreateRequest) SetAppConfigurationItems(v []AppConfigItemCreate) {
	o.AppConfigurationItems = v
}

func (o AppConfigItemsBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appConfigurationItems"] = o.AppConfigurationItems
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemsBulkCreateRequest struct {
	value *AppConfigItemsBulkCreateRequest
	isSet bool
}

func (v NullableAppConfigItemsBulkCreateRequest) Get() *AppConfigItemsBulkCreateRequest {
	return v.value
}

func (v *NullableAppConfigItemsBulkCreateRequest) Set(val *AppConfigItemsBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemsBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemsBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemsBulkCreateRequest(val *AppConfigItemsBulkCreateRequest) *NullableAppConfigItemsBulkCreateRequest {
	return &NullableAppConfigItemsBulkCreateRequest{value: val, isSet: true}
}

func (v NullableAppConfigItemsBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemsBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


