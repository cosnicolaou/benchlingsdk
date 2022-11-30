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

// AppConfigItemGenericCreate struct for AppConfigItemGenericCreate
type AppConfigItemGenericCreate struct {
	// App id to which this config item belongs.
	AppId string `json:"appId"`
	// Array-based representation of config item's location in the tree in order from top to bottom.
	Path []string `json:"path"`
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewAppConfigItemGenericCreate instantiates a new AppConfigItemGenericCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemGenericCreate(appId string, path []string, type_ string, value string) *AppConfigItemGenericCreate {
	this := AppConfigItemGenericCreate{}
	this.AppId = appId
	this.Path = path
	this.Type = type_
	this.Value = value
	return &this
}

// NewAppConfigItemGenericCreateWithDefaults instantiates a new AppConfigItemGenericCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemGenericCreateWithDefaults() *AppConfigItemGenericCreate {
	this := AppConfigItemGenericCreate{}
	return &this
}

// GetAppId returns the AppId field value
func (o *AppConfigItemGenericCreate) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemGenericCreate) GetAppIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AppConfigItemGenericCreate) SetAppId(v string) {
	o.AppId = v
}

// GetPath returns the Path field value
func (o *AppConfigItemGenericCreate) GetPath() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemGenericCreate) GetPathOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path, true
}

// SetPath sets field value
func (o *AppConfigItemGenericCreate) SetPath(v []string) {
	o.Path = v
}

// GetType returns the Type field value
func (o *AppConfigItemGenericCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemGenericCreate) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfigItemGenericCreate) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AppConfigItemGenericCreate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemGenericCreate) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AppConfigItemGenericCreate) SetValue(v string) {
	o.Value = v
}

func (o AppConfigItemGenericCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAppConfigItemGenericCreate struct {
	value *AppConfigItemGenericCreate
	isSet bool
}

func (v NullableAppConfigItemGenericCreate) Get() *AppConfigItemGenericCreate {
	return v.value
}

func (v *NullableAppConfigItemGenericCreate) Set(val *AppConfigItemGenericCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemGenericCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemGenericCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemGenericCreate(val *AppConfigItemGenericCreate) *NullableAppConfigItemGenericCreate {
	return &NullableAppConfigItemGenericCreate{value: val, isSet: true}
}

func (v NullableAppConfigItemGenericCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemGenericCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

