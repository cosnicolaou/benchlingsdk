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

// FloatFieldDefinitionAllOf struct for FloatFieldDefinitionAllOf
type FloatFieldDefinitionAllOf struct {
	DecimalPrecision NullableFloat32 `json:"decimalPrecision,omitempty"`
	LegalTextDropdownId NullableString `json:"legalTextDropdownId,omitempty"`
	NumericMax NullableFloat32 `json:"numericMax,omitempty"`
	NumericMin NullableFloat32 `json:"numericMin,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewFloatFieldDefinitionAllOf instantiates a new FloatFieldDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFloatFieldDefinitionAllOf() *FloatFieldDefinitionAllOf {
	this := FloatFieldDefinitionAllOf{}
	return &this
}

// NewFloatFieldDefinitionAllOfWithDefaults instantiates a new FloatFieldDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFloatFieldDefinitionAllOfWithDefaults() *FloatFieldDefinitionAllOf {
	this := FloatFieldDefinitionAllOf{}
	return &this
}

// GetDecimalPrecision returns the DecimalPrecision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FloatFieldDefinitionAllOf) GetDecimalPrecision() float32 {
	if o == nil || isNil(o.DecimalPrecision.Get()) {
		var ret float32
		return ret
	}
	return *o.DecimalPrecision.Get()
}

// GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FloatFieldDefinitionAllOf) GetDecimalPrecisionOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.DecimalPrecision.Get(), o.DecimalPrecision.IsSet()
}

// HasDecimalPrecision returns a boolean if a field has been set.
func (o *FloatFieldDefinitionAllOf) HasDecimalPrecision() bool {
	if o != nil && o.DecimalPrecision.IsSet() {
		return true
	}

	return false
}

// SetDecimalPrecision gets a reference to the given NullableFloat32 and assigns it to the DecimalPrecision field.
func (o *FloatFieldDefinitionAllOf) SetDecimalPrecision(v float32) {
	o.DecimalPrecision.Set(&v)
}
// SetDecimalPrecisionNil sets the value for DecimalPrecision to be an explicit nil
func (o *FloatFieldDefinitionAllOf) SetDecimalPrecisionNil() {
	o.DecimalPrecision.Set(nil)
}

// UnsetDecimalPrecision ensures that no value is present for DecimalPrecision, not even an explicit nil
func (o *FloatFieldDefinitionAllOf) UnsetDecimalPrecision() {
	o.DecimalPrecision.Unset()
}

// GetLegalTextDropdownId returns the LegalTextDropdownId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FloatFieldDefinitionAllOf) GetLegalTextDropdownId() string {
	if o == nil || isNil(o.LegalTextDropdownId.Get()) {
		var ret string
		return ret
	}
	return *o.LegalTextDropdownId.Get()
}

// GetLegalTextDropdownIdOk returns a tuple with the LegalTextDropdownId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FloatFieldDefinitionAllOf) GetLegalTextDropdownIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LegalTextDropdownId.Get(), o.LegalTextDropdownId.IsSet()
}

// HasLegalTextDropdownId returns a boolean if a field has been set.
func (o *FloatFieldDefinitionAllOf) HasLegalTextDropdownId() bool {
	if o != nil && o.LegalTextDropdownId.IsSet() {
		return true
	}

	return false
}

// SetLegalTextDropdownId gets a reference to the given NullableString and assigns it to the LegalTextDropdownId field.
func (o *FloatFieldDefinitionAllOf) SetLegalTextDropdownId(v string) {
	o.LegalTextDropdownId.Set(&v)
}
// SetLegalTextDropdownIdNil sets the value for LegalTextDropdownId to be an explicit nil
func (o *FloatFieldDefinitionAllOf) SetLegalTextDropdownIdNil() {
	o.LegalTextDropdownId.Set(nil)
}

// UnsetLegalTextDropdownId ensures that no value is present for LegalTextDropdownId, not even an explicit nil
func (o *FloatFieldDefinitionAllOf) UnsetLegalTextDropdownId() {
	o.LegalTextDropdownId.Unset()
}

// GetNumericMax returns the NumericMax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FloatFieldDefinitionAllOf) GetNumericMax() float32 {
	if o == nil || isNil(o.NumericMax.Get()) {
		var ret float32
		return ret
	}
	return *o.NumericMax.Get()
}

// GetNumericMaxOk returns a tuple with the NumericMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FloatFieldDefinitionAllOf) GetNumericMaxOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumericMax.Get(), o.NumericMax.IsSet()
}

// HasNumericMax returns a boolean if a field has been set.
func (o *FloatFieldDefinitionAllOf) HasNumericMax() bool {
	if o != nil && o.NumericMax.IsSet() {
		return true
	}

	return false
}

// SetNumericMax gets a reference to the given NullableFloat32 and assigns it to the NumericMax field.
func (o *FloatFieldDefinitionAllOf) SetNumericMax(v float32) {
	o.NumericMax.Set(&v)
}
// SetNumericMaxNil sets the value for NumericMax to be an explicit nil
func (o *FloatFieldDefinitionAllOf) SetNumericMaxNil() {
	o.NumericMax.Set(nil)
}

// UnsetNumericMax ensures that no value is present for NumericMax, not even an explicit nil
func (o *FloatFieldDefinitionAllOf) UnsetNumericMax() {
	o.NumericMax.Unset()
}

// GetNumericMin returns the NumericMin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FloatFieldDefinitionAllOf) GetNumericMin() float32 {
	if o == nil || isNil(o.NumericMin.Get()) {
		var ret float32
		return ret
	}
	return *o.NumericMin.Get()
}

// GetNumericMinOk returns a tuple with the NumericMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FloatFieldDefinitionAllOf) GetNumericMinOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumericMin.Get(), o.NumericMin.IsSet()
}

// HasNumericMin returns a boolean if a field has been set.
func (o *FloatFieldDefinitionAllOf) HasNumericMin() bool {
	if o != nil && o.NumericMin.IsSet() {
		return true
	}

	return false
}

// SetNumericMin gets a reference to the given NullableFloat32 and assigns it to the NumericMin field.
func (o *FloatFieldDefinitionAllOf) SetNumericMin(v float32) {
	o.NumericMin.Set(&v)
}
// SetNumericMinNil sets the value for NumericMin to be an explicit nil
func (o *FloatFieldDefinitionAllOf) SetNumericMinNil() {
	o.NumericMin.Set(nil)
}

// UnsetNumericMin ensures that no value is present for NumericMin, not even an explicit nil
func (o *FloatFieldDefinitionAllOf) UnsetNumericMin() {
	o.NumericMin.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FloatFieldDefinitionAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatFieldDefinitionAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FloatFieldDefinitionAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FloatFieldDefinitionAllOf) SetType(v string) {
	o.Type = &v
}

func (o FloatFieldDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DecimalPrecision.IsSet() {
		toSerialize["decimalPrecision"] = o.DecimalPrecision.Get()
	}
	if o.LegalTextDropdownId.IsSet() {
		toSerialize["legalTextDropdownId"] = o.LegalTextDropdownId.Get()
	}
	if o.NumericMax.IsSet() {
		toSerialize["numericMax"] = o.NumericMax.Get()
	}
	if o.NumericMin.IsSet() {
		toSerialize["numericMin"] = o.NumericMin.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFloatFieldDefinitionAllOf struct {
	value *FloatFieldDefinitionAllOf
	isSet bool
}

func (v NullableFloatFieldDefinitionAllOf) Get() *FloatFieldDefinitionAllOf {
	return v.value
}

func (v *NullableFloatFieldDefinitionAllOf) Set(val *FloatFieldDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFloatFieldDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFloatFieldDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloatFieldDefinitionAllOf(val *FloatFieldDefinitionAllOf) *NullableFloatFieldDefinitionAllOf {
	return &NullableFloatFieldDefinitionAllOf{value: val, isSet: true}
}

func (v NullableFloatFieldDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloatFieldDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

