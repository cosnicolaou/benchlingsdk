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

// BaseAppConfigItem struct for BaseAppConfigItem
type BaseAppConfigItem struct {
	ApiURL *string `json:"apiURL,omitempty"`
	App *AppConfigItemApiMixinApp `json:"app,omitempty"`
	Id *string `json:"id,omitempty"`
	// DateTime the app config item was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// Array-based representation of config item's location in the tree in order from top to bottom.
	Path []string `json:"path,omitempty"`
	// Type of the app config item
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	RequiredConfig *bool `json:"requiredConfig,omitempty"`
}

// NewBaseAppConfigItem instantiates a new BaseAppConfigItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseAppConfigItem() *BaseAppConfigItem {
	this := BaseAppConfigItem{}
	return &this
}

// NewBaseAppConfigItemWithDefaults instantiates a new BaseAppConfigItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseAppConfigItemWithDefaults() *BaseAppConfigItem {
	this := BaseAppConfigItem{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *BaseAppConfigItem) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetApp() AppConfigItemApiMixinApp {
	if o == nil || isNil(o.App) {
		var ret AppConfigItemApiMixinApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetAppOk() (*AppConfigItemApiMixinApp, bool) {
	if o == nil || isNil(o.App) {
    return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasApp() bool {
	if o != nil && !isNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppConfigItemApiMixinApp and assigns it to the App field.
func (o *BaseAppConfigItem) SetApp(v AppConfigItemApiMixinApp) {
	o.App = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseAppConfigItem) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *BaseAppConfigItem) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetPath() []string {
	if o == nil || isNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetPathOk() ([]string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BaseAppConfigItem) SetPath(v []string) {
	o.Path = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BaseAppConfigItem) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseAppConfigItem) SetDescription(v string) {
	o.Description = &v
}

// GetRequiredConfig returns the RequiredConfig field value if set, zero value otherwise.
func (o *BaseAppConfigItem) GetRequiredConfig() bool {
	if o == nil || isNil(o.RequiredConfig) {
		var ret bool
		return ret
	}
	return *o.RequiredConfig
}

// GetRequiredConfigOk returns a tuple with the RequiredConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseAppConfigItem) GetRequiredConfigOk() (*bool, bool) {
	if o == nil || isNil(o.RequiredConfig) {
    return nil, false
	}
	return o.RequiredConfig, true
}

// HasRequiredConfig returns a boolean if a field has been set.
func (o *BaseAppConfigItem) HasRequiredConfig() bool {
	if o != nil && !isNil(o.RequiredConfig) {
		return true
	}

	return false
}

// SetRequiredConfig gets a reference to the given bool and assigns it to the RequiredConfig field.
func (o *BaseAppConfigItem) SetRequiredConfig(v bool) {
	o.RequiredConfig = &v
}

func (o BaseAppConfigItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
	if !isNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.RequiredConfig) {
		toSerialize["requiredConfig"] = o.RequiredConfig
	}
	return json.Marshal(toSerialize)
}

type NullableBaseAppConfigItem struct {
	value *BaseAppConfigItem
	isSet bool
}

func (v NullableBaseAppConfigItem) Get() *BaseAppConfigItem {
	return v.value
}

func (v *NullableBaseAppConfigItem) Set(val *BaseAppConfigItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseAppConfigItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseAppConfigItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseAppConfigItem(val *BaseAppConfigItem) *NullableBaseAppConfigItem {
	return &NullableBaseAppConfigItem{value: val, isSet: true}
}

func (v NullableBaseAppConfigItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseAppConfigItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


