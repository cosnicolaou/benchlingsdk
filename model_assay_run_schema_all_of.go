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

// AssayRunSchemaAllOf struct for AssayRunSchemaAllOf
type AssayRunSchemaAllOf struct {
	AutomationInputFileConfigs []AssayRunSchemaAllOfAutomationInputFileConfigs `json:"automationInputFileConfigs,omitempty"`
	AutomationOutputFileConfigs []AssayRunSchemaAllOfAutomationInputFileConfigs `json:"automationOutputFileConfigs,omitempty"`
	// DateTime the Assay Run Schema was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAssayRunSchemaAllOf instantiates a new AssayRunSchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayRunSchemaAllOf() *AssayRunSchemaAllOf {
	this := AssayRunSchemaAllOf{}
	return &this
}

// NewAssayRunSchemaAllOfWithDefaults instantiates a new AssayRunSchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayRunSchemaAllOfWithDefaults() *AssayRunSchemaAllOf {
	this := AssayRunSchemaAllOf{}
	return &this
}

// GetAutomationInputFileConfigs returns the AutomationInputFileConfigs field value if set, zero value otherwise.
func (o *AssayRunSchemaAllOf) GetAutomationInputFileConfigs() []AssayRunSchemaAllOfAutomationInputFileConfigs {
	if o == nil || isNil(o.AutomationInputFileConfigs) {
		var ret []AssayRunSchemaAllOfAutomationInputFileConfigs
		return ret
	}
	return o.AutomationInputFileConfigs
}

// GetAutomationInputFileConfigsOk returns a tuple with the AutomationInputFileConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunSchemaAllOf) GetAutomationInputFileConfigsOk() ([]AssayRunSchemaAllOfAutomationInputFileConfigs, bool) {
	if o == nil || isNil(o.AutomationInputFileConfigs) {
    return nil, false
	}
	return o.AutomationInputFileConfigs, true
}

// HasAutomationInputFileConfigs returns a boolean if a field has been set.
func (o *AssayRunSchemaAllOf) HasAutomationInputFileConfigs() bool {
	if o != nil && !isNil(o.AutomationInputFileConfigs) {
		return true
	}

	return false
}

// SetAutomationInputFileConfigs gets a reference to the given []AssayRunSchemaAllOfAutomationInputFileConfigs and assigns it to the AutomationInputFileConfigs field.
func (o *AssayRunSchemaAllOf) SetAutomationInputFileConfigs(v []AssayRunSchemaAllOfAutomationInputFileConfigs) {
	o.AutomationInputFileConfigs = v
}

// GetAutomationOutputFileConfigs returns the AutomationOutputFileConfigs field value if set, zero value otherwise.
func (o *AssayRunSchemaAllOf) GetAutomationOutputFileConfigs() []AssayRunSchemaAllOfAutomationInputFileConfigs {
	if o == nil || isNil(o.AutomationOutputFileConfigs) {
		var ret []AssayRunSchemaAllOfAutomationInputFileConfigs
		return ret
	}
	return o.AutomationOutputFileConfigs
}

// GetAutomationOutputFileConfigsOk returns a tuple with the AutomationOutputFileConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunSchemaAllOf) GetAutomationOutputFileConfigsOk() ([]AssayRunSchemaAllOfAutomationInputFileConfigs, bool) {
	if o == nil || isNil(o.AutomationOutputFileConfigs) {
    return nil, false
	}
	return o.AutomationOutputFileConfigs, true
}

// HasAutomationOutputFileConfigs returns a boolean if a field has been set.
func (o *AssayRunSchemaAllOf) HasAutomationOutputFileConfigs() bool {
	if o != nil && !isNil(o.AutomationOutputFileConfigs) {
		return true
	}

	return false
}

// SetAutomationOutputFileConfigs gets a reference to the given []AssayRunSchemaAllOfAutomationInputFileConfigs and assigns it to the AutomationOutputFileConfigs field.
func (o *AssayRunSchemaAllOf) SetAutomationOutputFileConfigs(v []AssayRunSchemaAllOfAutomationInputFileConfigs) {
	o.AutomationOutputFileConfigs = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AssayRunSchemaAllOf) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunSchemaAllOf) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AssayRunSchemaAllOf) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AssayRunSchemaAllOf) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssayRunSchemaAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRunSchemaAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssayRunSchemaAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssayRunSchemaAllOf) SetType(v string) {
	o.Type = &v
}

func (o AssayRunSchemaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AutomationInputFileConfigs) {
		toSerialize["automationInputFileConfigs"] = o.AutomationInputFileConfigs
	}
	if !isNil(o.AutomationOutputFileConfigs) {
		toSerialize["automationOutputFileConfigs"] = o.AutomationOutputFileConfigs
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAssayRunSchemaAllOf struct {
	value *AssayRunSchemaAllOf
	isSet bool
}

func (v NullableAssayRunSchemaAllOf) Get() *AssayRunSchemaAllOf {
	return v.value
}

func (v *NullableAssayRunSchemaAllOf) Set(val *AssayRunSchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayRunSchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayRunSchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayRunSchemaAllOf(val *AssayRunSchemaAllOf) *NullableAssayRunSchemaAllOf {
	return &NullableAssayRunSchemaAllOf{value: val, isSet: true}
}

func (v NullableAssayRunSchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayRunSchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

