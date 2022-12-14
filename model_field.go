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

// Field struct for Field
type Field struct {
	DisplayValue NullableString `json:"displayValue,omitempty"`
	IsMulti *bool `json:"isMulti,omitempty"`
	TextValue NullableString `json:"textValue,omitempty"`
	Type *FieldType `json:"type,omitempty"`
	Value NullableFieldValue `json:"value"`
}

// NewField instantiates a new Field object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewField(value NullableFieldValue) *Field {
	this := Field{}
	this.Value = value
	return &this
}

// NewFieldWithDefaults instantiates a new Field object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldWithDefaults() *Field {
	this := Field{}
	return &this
}

// GetDisplayValue returns the DisplayValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetDisplayValue() string {
	if o == nil || isNil(o.DisplayValue.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayValue.Get()
}

// GetDisplayValueOk returns a tuple with the DisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetDisplayValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DisplayValue.Get(), o.DisplayValue.IsSet()
}

// HasDisplayValue returns a boolean if a field has been set.
func (o *Field) HasDisplayValue() bool {
	if o != nil && o.DisplayValue.IsSet() {
		return true
	}

	return false
}

// SetDisplayValue gets a reference to the given NullableString and assigns it to the DisplayValue field.
func (o *Field) SetDisplayValue(v string) {
	o.DisplayValue.Set(&v)
}
// SetDisplayValueNil sets the value for DisplayValue to be an explicit nil
func (o *Field) SetDisplayValueNil() {
	o.DisplayValue.Set(nil)
}

// UnsetDisplayValue ensures that no value is present for DisplayValue, not even an explicit nil
func (o *Field) UnsetDisplayValue() {
	o.DisplayValue.Unset()
}

// GetIsMulti returns the IsMulti field value if set, zero value otherwise.
func (o *Field) GetIsMulti() bool {
	if o == nil || isNil(o.IsMulti) {
		var ret bool
		return ret
	}
	return *o.IsMulti
}

// GetIsMultiOk returns a tuple with the IsMulti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetIsMultiOk() (*bool, bool) {
	if o == nil || isNil(o.IsMulti) {
    return nil, false
	}
	return o.IsMulti, true
}

// HasIsMulti returns a boolean if a field has been set.
func (o *Field) HasIsMulti() bool {
	if o != nil && !isNil(o.IsMulti) {
		return true
	}

	return false
}

// SetIsMulti gets a reference to the given bool and assigns it to the IsMulti field.
func (o *Field) SetIsMulti(v bool) {
	o.IsMulti = &v
}

// GetTextValue returns the TextValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Field) GetTextValue() string {
	if o == nil || isNil(o.TextValue.Get()) {
		var ret string
		return ret
	}
	return *o.TextValue.Get()
}

// GetTextValueOk returns a tuple with the TextValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetTextValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TextValue.Get(), o.TextValue.IsSet()
}

// HasTextValue returns a boolean if a field has been set.
func (o *Field) HasTextValue() bool {
	if o != nil && o.TextValue.IsSet() {
		return true
	}

	return false
}

// SetTextValue gets a reference to the given NullableString and assigns it to the TextValue field.
func (o *Field) SetTextValue(v string) {
	o.TextValue.Set(&v)
}
// SetTextValueNil sets the value for TextValue to be an explicit nil
func (o *Field) SetTextValueNil() {
	o.TextValue.Set(nil)
}

// UnsetTextValue ensures that no value is present for TextValue, not even an explicit nil
func (o *Field) UnsetTextValue() {
	o.TextValue.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Field) GetType() FieldType {
	if o == nil || isNil(o.Type) {
		var ret FieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetTypeOk() (*FieldType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Field) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FieldType and assigns it to the Type field.
func (o *Field) SetType(v FieldType) {
	o.Type = &v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for FieldValue will be returned
func (o *Field) GetValue() FieldValue {
	if o == nil || o.Value.Get() == nil {
		var ret FieldValue
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Field) GetValueOk() (*FieldValue, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *Field) SetValue(v FieldValue) {
	o.Value.Set(&v)
}

func (o Field) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayValue.IsSet() {
		toSerialize["displayValue"] = o.DisplayValue.Get()
	}
	if !isNil(o.IsMulti) {
		toSerialize["isMulti"] = o.IsMulti
	}
	if o.TextValue.IsSet() {
		toSerialize["textValue"] = o.TextValue.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableField struct {
	value *Field
	isSet bool
}

func (v NullableField) Get() *Field {
	return v.value
}

func (v *NullableField) Set(val *Field) {
	v.value = val
	v.isSet = true
}

func (v NullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableField(val *Field) *NullableField {
	return &NullableField{value: val, isSet: true}
}

func (v NullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


