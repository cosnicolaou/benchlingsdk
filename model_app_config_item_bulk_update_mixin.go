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

// AppConfigItemBulkUpdateMixin struct for AppConfigItemBulkUpdateMixin
type AppConfigItemBulkUpdateMixin struct {
	Id string `json:"id"`
}

// NewAppConfigItemBulkUpdateMixin instantiates a new AppConfigItemBulkUpdateMixin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemBulkUpdateMixin(id string) *AppConfigItemBulkUpdateMixin {
	this := AppConfigItemBulkUpdateMixin{}
	this.Id = id
	return &this
}

// NewAppConfigItemBulkUpdateMixinWithDefaults instantiates a new AppConfigItemBulkUpdateMixin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemBulkUpdateMixinWithDefaults() *AppConfigItemBulkUpdateMixin {
	this := AppConfigItemBulkUpdateMixin{}
	return &this
}

// GetId returns the Id field value
func (o *AppConfigItemBulkUpdateMixin) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemBulkUpdateMixin) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppConfigItemBulkUpdateMixin) SetId(v string) {
	o.Id = v
}

func (o AppConfigItemBulkUpdateMixin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemBulkUpdateMixin struct {
	value *AppConfigItemBulkUpdateMixin
	isSet bool
}

func (v NullableAppConfigItemBulkUpdateMixin) Get() *AppConfigItemBulkUpdateMixin {
	return v.value
}

func (v *NullableAppConfigItemBulkUpdateMixin) Set(val *AppConfigItemBulkUpdateMixin) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemBulkUpdateMixin) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemBulkUpdateMixin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemBulkUpdateMixin(val *AppConfigItemBulkUpdateMixin) *NullableAppConfigItemBulkUpdateMixin {
	return &NullableAppConfigItemBulkUpdateMixin{value: val, isSet: true}
}

func (v NullableAppConfigItemBulkUpdateMixin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemBulkUpdateMixin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


