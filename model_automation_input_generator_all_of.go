/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
	"time"
)

// AutomationInputGeneratorAllOf struct for AutomationInputGeneratorAllOf
type AutomationInputGeneratorAllOf struct {
	// The canonical url of the Automation Input Generator in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	// DateTime the Automation Input Generator was last modified
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id *string `json:"id,omitempty"`
	// DateTime the Automation Input Generator was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
}

// NewAutomationInputGeneratorAllOf instantiates a new AutomationInputGeneratorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationInputGeneratorAllOf() *AutomationInputGeneratorAllOf {
	this := AutomationInputGeneratorAllOf{}
	return &this
}

// NewAutomationInputGeneratorAllOfWithDefaults instantiates a new AutomationInputGeneratorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationInputGeneratorAllOfWithDefaults() *AutomationInputGeneratorAllOf {
	this := AutomationInputGeneratorAllOf{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *AutomationInputGeneratorAllOf) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorAllOf) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *AutomationInputGeneratorAllOf) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *AutomationInputGeneratorAllOf) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutomationInputGeneratorAllOf) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutomationInputGeneratorAllOf) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AutomationInputGeneratorAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationInputGeneratorAllOf) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorAllOf) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationInputGeneratorAllOf) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutomationInputGeneratorAllOf) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AutomationInputGeneratorAllOf) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorAllOf) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AutomationInputGeneratorAllOf) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AutomationInputGeneratorAllOf) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o AutomationInputGeneratorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAutomationInputGeneratorAllOf struct {
	value *AutomationInputGeneratorAllOf
	isSet bool
}

func (v NullableAutomationInputGeneratorAllOf) Get() *AutomationInputGeneratorAllOf {
	return v.value
}

func (v *NullableAutomationInputGeneratorAllOf) Set(val *AutomationInputGeneratorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationInputGeneratorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationInputGeneratorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationInputGeneratorAllOf(val *AutomationInputGeneratorAllOf) *NullableAutomationInputGeneratorAllOf {
	return &NullableAutomationInputGeneratorAllOf{value: val, isSet: true}
}

func (v NullableAutomationInputGeneratorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationInputGeneratorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

