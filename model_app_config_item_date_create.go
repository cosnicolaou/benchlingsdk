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

// AppConfigItemDateCreate struct for AppConfigItemDateCreate
type AppConfigItemDateCreate struct {
	// App id to which this config item belongs.
	AppId string `json:"appId"`
	// Array-based representation of config item's location in the tree in order from top to bottom.
	Path []string `json:"path"`
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewAppConfigItemDateCreate instantiates a new AppConfigItemDateCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigItemDateCreate(appId string, path []string, type_ string, value string) *AppConfigItemDateCreate {
	this := AppConfigItemDateCreate{}
	this.AppId = appId
	this.Path = path
	this.Type = type_
	this.Value = value
	return &this
}

// NewAppConfigItemDateCreateWithDefaults instantiates a new AppConfigItemDateCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigItemDateCreateWithDefaults() *AppConfigItemDateCreate {
	this := AppConfigItemDateCreate{}
	return &this
}

// GetAppId returns the AppId field value
func (o *AppConfigItemDateCreate) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDateCreate) GetAppIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AppConfigItemDateCreate) SetAppId(v string) {
	o.AppId = v
}

// GetPath returns the Path field value
func (o *AppConfigItemDateCreate) GetPath() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDateCreate) GetPathOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Path, true
}

// SetPath sets field value
func (o *AppConfigItemDateCreate) SetPath(v []string) {
	o.Path = v
}

// GetType returns the Type field value
func (o *AppConfigItemDateCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDateCreate) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfigItemDateCreate) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AppConfigItemDateCreate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AppConfigItemDateCreate) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AppConfigItemDateCreate) SetValue(v string) {
	o.Value = v
}

func (o AppConfigItemDateCreate) MarshalJSON() ([]byte, error) {
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

type NullableAppConfigItemDateCreate struct {
	value *AppConfigItemDateCreate
	isSet bool
}

func (v NullableAppConfigItemDateCreate) Get() *AppConfigItemDateCreate {
	return v.value
}

func (v *NullableAppConfigItemDateCreate) Set(val *AppConfigItemDateCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemDateCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemDateCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemDateCreate(val *AppConfigItemDateCreate) *NullableAppConfigItemDateCreate {
	return &NullableAppConfigItemDateCreate{value: val, isSet: true}
}

func (v NullableAppConfigItemDateCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemDateCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


