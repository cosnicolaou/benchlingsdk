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

// AutomationInputGeneratorCompletedV2BetaEvent struct for AutomationInputGeneratorCompletedV2BetaEvent
type AutomationInputGeneratorCompletedV2BetaEvent struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	// These properties have been dropped from the payload due to size. 
	ExcludedProperties []string `json:"excludedProperties,omitempty"`
	Id *string `json:"id,omitempty"`
	Schema NullableEventBaseSchema `json:"schema,omitempty"`
	AutomationInputGenerator *AutomationFile `json:"automationInputGenerator,omitempty"`
	EventType *string `json:"eventType,omitempty"`
}

// NewAutomationInputGeneratorCompletedV2BetaEvent instantiates a new AutomationInputGeneratorCompletedV2BetaEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationInputGeneratorCompletedV2BetaEvent() *AutomationInputGeneratorCompletedV2BetaEvent {
	this := AutomationInputGeneratorCompletedV2BetaEvent{}
	return &this
}

// NewAutomationInputGeneratorCompletedV2BetaEventWithDefaults instantiates a new AutomationInputGeneratorCompletedV2BetaEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationInputGeneratorCompletedV2BetaEventWithDefaults() *AutomationInputGeneratorCompletedV2BetaEvent {
	this := AutomationInputGeneratorCompletedV2BetaEvent{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetDeprecated() bool {
	if o == nil || isNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetDeprecatedOk() (*bool, bool) {
	if o == nil || isNil(o.Deprecated) {
    return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasDeprecated() bool {
	if o != nil && !isNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetExcludedProperties returns the ExcludedProperties field value if set, zero value otherwise.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetExcludedProperties() []string {
	if o == nil || isNil(o.ExcludedProperties) {
		var ret []string
		return ret
	}
	return o.ExcludedProperties
}

// GetExcludedPropertiesOk returns a tuple with the ExcludedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetExcludedPropertiesOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedProperties) {
    return nil, false
	}
	return o.ExcludedProperties, true
}

// HasExcludedProperties returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasExcludedProperties() bool {
	if o != nil && !isNil(o.ExcludedProperties) {
		return true
	}

	return false
}

// SetExcludedProperties gets a reference to the given []string and assigns it to the ExcludedProperties field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetExcludedProperties(v []string) {
	o.ExcludedProperties = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetId(v string) {
	o.Id = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetSchema() EventBaseSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret EventBaseSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetSchemaOk() (*EventBaseSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableEventBaseSchema and assigns it to the Schema field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetSchema(v EventBaseSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *AutomationInputGeneratorCompletedV2BetaEvent) UnsetSchema() {
	o.Schema.Unset()
}

// GetAutomationInputGenerator returns the AutomationInputGenerator field value if set, zero value otherwise.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetAutomationInputGenerator() AutomationFile {
	if o == nil || isNil(o.AutomationInputGenerator) {
		var ret AutomationFile
		return ret
	}
	return *o.AutomationInputGenerator
}

// GetAutomationInputGeneratorOk returns a tuple with the AutomationInputGenerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetAutomationInputGeneratorOk() (*AutomationFile, bool) {
	if o == nil || isNil(o.AutomationInputGenerator) {
    return nil, false
	}
	return o.AutomationInputGenerator, true
}

// HasAutomationInputGenerator returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasAutomationInputGenerator() bool {
	if o != nil && !isNil(o.AutomationInputGenerator) {
		return true
	}

	return false
}

// SetAutomationInputGenerator gets a reference to the given AutomationFile and assigns it to the AutomationInputGenerator field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetAutomationInputGenerator(v AutomationFile) {
	o.AutomationInputGenerator = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetEventType() string {
	if o == nil || isNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *AutomationInputGeneratorCompletedV2BetaEvent) SetEventType(v string) {
	o.EventType = &v
}

func (o AutomationInputGeneratorCompletedV2BetaEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !isNil(o.ExcludedProperties) {
		toSerialize["excludedProperties"] = o.ExcludedProperties
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if !isNil(o.AutomationInputGenerator) {
		toSerialize["automationInputGenerator"] = o.AutomationInputGenerator
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return json.Marshal(toSerialize)
}

type NullableAutomationInputGeneratorCompletedV2BetaEvent struct {
	value *AutomationInputGeneratorCompletedV2BetaEvent
	isSet bool
}

func (v NullableAutomationInputGeneratorCompletedV2BetaEvent) Get() *AutomationInputGeneratorCompletedV2BetaEvent {
	return v.value
}

func (v *NullableAutomationInputGeneratorCompletedV2BetaEvent) Set(val *AutomationInputGeneratorCompletedV2BetaEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationInputGeneratorCompletedV2BetaEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationInputGeneratorCompletedV2BetaEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationInputGeneratorCompletedV2BetaEvent(val *AutomationInputGeneratorCompletedV2BetaEvent) *NullableAutomationInputGeneratorCompletedV2BetaEvent {
	return &NullableAutomationInputGeneratorCompletedV2BetaEvent{value: val, isSet: true}
}

func (v NullableAutomationInputGeneratorCompletedV2BetaEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationInputGeneratorCompletedV2BetaEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


