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

// EntitySchemaAppConfigItemAllOf struct for EntitySchemaAppConfigItemAllOf
type EntitySchemaAppConfigItemAllOf struct {
	Subtype *SchemaDependencySubtypes `json:"subtype,omitempty"`
	Type *string `json:"type,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewEntitySchemaAppConfigItemAllOf instantiates a new EntitySchemaAppConfigItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySchemaAppConfigItemAllOf() *EntitySchemaAppConfigItemAllOf {
	this := EntitySchemaAppConfigItemAllOf{}
	return &this
}

// NewEntitySchemaAppConfigItemAllOfWithDefaults instantiates a new EntitySchemaAppConfigItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySchemaAppConfigItemAllOfWithDefaults() *EntitySchemaAppConfigItemAllOf {
	this := EntitySchemaAppConfigItemAllOf{}
	return &this
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *EntitySchemaAppConfigItemAllOf) GetSubtype() SchemaDependencySubtypes {
	if o == nil || isNil(o.Subtype) {
		var ret SchemaDependencySubtypes
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAppConfigItemAllOf) GetSubtypeOk() (*SchemaDependencySubtypes, bool) {
	if o == nil || isNil(o.Subtype) {
    return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *EntitySchemaAppConfigItemAllOf) HasSubtype() bool {
	if o != nil && !isNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given SchemaDependencySubtypes and assigns it to the Subtype field.
func (o *EntitySchemaAppConfigItemAllOf) SetSubtype(v SchemaDependencySubtypes) {
	o.Subtype = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitySchemaAppConfigItemAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAppConfigItemAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitySchemaAppConfigItemAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntitySchemaAppConfigItemAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitySchemaAppConfigItemAllOf) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitySchemaAppConfigItemAllOf) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EntitySchemaAppConfigItemAllOf) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EntitySchemaAppConfigItemAllOf) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *EntitySchemaAppConfigItemAllOf) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EntitySchemaAppConfigItemAllOf) UnsetValue() {
	o.Value.Unset()
}

func (o EntitySchemaAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEntitySchemaAppConfigItemAllOf struct {
	value *EntitySchemaAppConfigItemAllOf
	isSet bool
}

func (v NullableEntitySchemaAppConfigItemAllOf) Get() *EntitySchemaAppConfigItemAllOf {
	return v.value
}

func (v *NullableEntitySchemaAppConfigItemAllOf) Set(val *EntitySchemaAppConfigItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySchemaAppConfigItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySchemaAppConfigItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySchemaAppConfigItemAllOf(val *EntitySchemaAppConfigItemAllOf) *NullableEntitySchemaAppConfigItemAllOf {
	return &NullableEntitySchemaAppConfigItemAllOf{value: val, isSet: true}
}

func (v NullableEntitySchemaAppConfigItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySchemaAppConfigItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


