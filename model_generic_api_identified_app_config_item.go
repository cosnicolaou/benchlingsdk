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

// GenericApiIdentifiedAppConfigItem struct for GenericApiIdentifiedAppConfigItem
type GenericApiIdentifiedAppConfigItem struct {
	ApiURL *string `json:"apiURL,omitempty"`
	App *AppConfigItemApiMixinApp `json:"app,omitempty"`
	Id *string `json:"id,omitempty"`
	// DateTime the app config item was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// Array-based representation of config item's location in the tree in order from top to bottom.
	Path []string `json:"path,omitempty"`
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	RequiredConfig *bool `json:"requiredConfig,omitempty"`
	LinkedResource NullableLinkedAppConfigResourceMixinLinkedResource `json:"linkedResource,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewGenericApiIdentifiedAppConfigItem instantiates a new GenericApiIdentifiedAppConfigItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericApiIdentifiedAppConfigItem() *GenericApiIdentifiedAppConfigItem {
	this := GenericApiIdentifiedAppConfigItem{}
	return &this
}

// NewGenericApiIdentifiedAppConfigItemWithDefaults instantiates a new GenericApiIdentifiedAppConfigItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericApiIdentifiedAppConfigItemWithDefaults() *GenericApiIdentifiedAppConfigItem {
	this := GenericApiIdentifiedAppConfigItem{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *GenericApiIdentifiedAppConfigItem) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetApp() AppConfigItemApiMixinApp {
	if o == nil || isNil(o.App) {
		var ret AppConfigItemApiMixinApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetAppOk() (*AppConfigItemApiMixinApp, bool) {
	if o == nil || isNil(o.App) {
    return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasApp() bool {
	if o != nil && !isNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppConfigItemApiMixinApp and assigns it to the App field.
func (o *GenericApiIdentifiedAppConfigItem) SetApp(v AppConfigItemApiMixinApp) {
	o.App = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GenericApiIdentifiedAppConfigItem) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *GenericApiIdentifiedAppConfigItem) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetPath() []string {
	if o == nil || isNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetPathOk() ([]string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *GenericApiIdentifiedAppConfigItem) SetPath(v []string) {
	o.Path = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GenericApiIdentifiedAppConfigItem) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GenericApiIdentifiedAppConfigItem) SetDescription(v string) {
	o.Description = &v
}

// GetRequiredConfig returns the RequiredConfig field value if set, zero value otherwise.
func (o *GenericApiIdentifiedAppConfigItem) GetRequiredConfig() bool {
	if o == nil || isNil(o.RequiredConfig) {
		var ret bool
		return ret
	}
	return *o.RequiredConfig
}

// GetRequiredConfigOk returns a tuple with the RequiredConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericApiIdentifiedAppConfigItem) GetRequiredConfigOk() (*bool, bool) {
	if o == nil || isNil(o.RequiredConfig) {
    return nil, false
	}
	return o.RequiredConfig, true
}

// HasRequiredConfig returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasRequiredConfig() bool {
	if o != nil && !isNil(o.RequiredConfig) {
		return true
	}

	return false
}

// SetRequiredConfig gets a reference to the given bool and assigns it to the RequiredConfig field.
func (o *GenericApiIdentifiedAppConfigItem) SetRequiredConfig(v bool) {
	o.RequiredConfig = &v
}

// GetLinkedResource returns the LinkedResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericApiIdentifiedAppConfigItem) GetLinkedResource() LinkedAppConfigResourceMixinLinkedResource {
	if o == nil || isNil(o.LinkedResource.Get()) {
		var ret LinkedAppConfigResourceMixinLinkedResource
		return ret
	}
	return *o.LinkedResource.Get()
}

// GetLinkedResourceOk returns a tuple with the LinkedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericApiIdentifiedAppConfigItem) GetLinkedResourceOk() (*LinkedAppConfigResourceMixinLinkedResource, bool) {
	if o == nil {
    return nil, false
	}
	return o.LinkedResource.Get(), o.LinkedResource.IsSet()
}

// HasLinkedResource returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasLinkedResource() bool {
	if o != nil && o.LinkedResource.IsSet() {
		return true
	}

	return false
}

// SetLinkedResource gets a reference to the given NullableLinkedAppConfigResourceMixinLinkedResource and assigns it to the LinkedResource field.
func (o *GenericApiIdentifiedAppConfigItem) SetLinkedResource(v LinkedAppConfigResourceMixinLinkedResource) {
	o.LinkedResource.Set(&v)
}
// SetLinkedResourceNil sets the value for LinkedResource to be an explicit nil
func (o *GenericApiIdentifiedAppConfigItem) SetLinkedResourceNil() {
	o.LinkedResource.Set(nil)
}

// UnsetLinkedResource ensures that no value is present for LinkedResource, not even an explicit nil
func (o *GenericApiIdentifiedAppConfigItem) UnsetLinkedResource() {
	o.LinkedResource.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenericApiIdentifiedAppConfigItem) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericApiIdentifiedAppConfigItem) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *GenericApiIdentifiedAppConfigItem) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *GenericApiIdentifiedAppConfigItem) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *GenericApiIdentifiedAppConfigItem) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *GenericApiIdentifiedAppConfigItem) UnsetValue() {
	o.Value.Unset()
}

func (o GenericApiIdentifiedAppConfigItem) MarshalJSON() ([]byte, error) {
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
	if o.LinkedResource.IsSet() {
		toSerialize["linkedResource"] = o.LinkedResource.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGenericApiIdentifiedAppConfigItem struct {
	value *GenericApiIdentifiedAppConfigItem
	isSet bool
}

func (v NullableGenericApiIdentifiedAppConfigItem) Get() *GenericApiIdentifiedAppConfigItem {
	return v.value
}

func (v *NullableGenericApiIdentifiedAppConfigItem) Set(val *GenericApiIdentifiedAppConfigItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericApiIdentifiedAppConfigItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericApiIdentifiedAppConfigItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericApiIdentifiedAppConfigItem(val *GenericApiIdentifiedAppConfigItem) *NullableGenericApiIdentifiedAppConfigItem {
	return &NullableGenericApiIdentifiedAppConfigItem{value: val, isSet: true}
}

func (v NullableGenericApiIdentifiedAppConfigItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericApiIdentifiedAppConfigItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


