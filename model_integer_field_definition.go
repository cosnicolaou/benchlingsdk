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

// IntegerFieldDefinition struct for IntegerFieldDefinition
type IntegerFieldDefinition struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	Id *string `json:"id,omitempty"`
	IsMulti *bool `json:"isMulti,omitempty"`
	IsRequired *bool `json:"isRequired,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	NumericMax NullableFloat32 `json:"numericMax,omitempty"`
	NumericMin NullableFloat32 `json:"numericMin,omitempty"`
}

// NewIntegerFieldDefinition instantiates a new IntegerFieldDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerFieldDefinition() *IntegerFieldDefinition {
	this := IntegerFieldDefinition{}
	return &this
}

// NewIntegerFieldDefinitionWithDefaults instantiates a new IntegerFieldDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerFieldDefinitionWithDefaults() *IntegerFieldDefinition {
	this := IntegerFieldDefinition{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegerFieldDefinition) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegerFieldDefinition) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *IntegerFieldDefinition) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *IntegerFieldDefinition) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *IntegerFieldDefinition) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegerFieldDefinition) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerFieldDefinition) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegerFieldDefinition) SetId(v string) {
	o.Id = &v
}

// GetIsMulti returns the IsMulti field value if set, zero value otherwise.
func (o *IntegerFieldDefinition) GetIsMulti() bool {
	if o == nil || isNil(o.IsMulti) {
		var ret bool
		return ret
	}
	return *o.IsMulti
}

// GetIsMultiOk returns a tuple with the IsMulti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerFieldDefinition) GetIsMultiOk() (*bool, bool) {
	if o == nil || isNil(o.IsMulti) {
    return nil, false
	}
	return o.IsMulti, true
}

// HasIsMulti returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasIsMulti() bool {
	if o != nil && !isNil(o.IsMulti) {
		return true
	}

	return false
}

// SetIsMulti gets a reference to the given bool and assigns it to the IsMulti field.
func (o *IntegerFieldDefinition) SetIsMulti(v bool) {
	o.IsMulti = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *IntegerFieldDefinition) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerFieldDefinition) GetIsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequired) {
    return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasIsRequired() bool {
	if o != nil && !isNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *IntegerFieldDefinition) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegerFieldDefinition) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerFieldDefinition) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegerFieldDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegerFieldDefinition) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerFieldDefinition) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IntegerFieldDefinition) SetType(v string) {
	o.Type = &v
}

// GetNumericMax returns the NumericMax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegerFieldDefinition) GetNumericMax() float32 {
	if o == nil || isNil(o.NumericMax.Get()) {
		var ret float32
		return ret
	}
	return *o.NumericMax.Get()
}

// GetNumericMaxOk returns a tuple with the NumericMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegerFieldDefinition) GetNumericMaxOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumericMax.Get(), o.NumericMax.IsSet()
}

// HasNumericMax returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasNumericMax() bool {
	if o != nil && o.NumericMax.IsSet() {
		return true
	}

	return false
}

// SetNumericMax gets a reference to the given NullableFloat32 and assigns it to the NumericMax field.
func (o *IntegerFieldDefinition) SetNumericMax(v float32) {
	o.NumericMax.Set(&v)
}
// SetNumericMaxNil sets the value for NumericMax to be an explicit nil
func (o *IntegerFieldDefinition) SetNumericMaxNil() {
	o.NumericMax.Set(nil)
}

// UnsetNumericMax ensures that no value is present for NumericMax, not even an explicit nil
func (o *IntegerFieldDefinition) UnsetNumericMax() {
	o.NumericMax.Unset()
}

// GetNumericMin returns the NumericMin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegerFieldDefinition) GetNumericMin() float32 {
	if o == nil || isNil(o.NumericMin.Get()) {
		var ret float32
		return ret
	}
	return *o.NumericMin.Get()
}

// GetNumericMinOk returns a tuple with the NumericMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegerFieldDefinition) GetNumericMinOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumericMin.Get(), o.NumericMin.IsSet()
}

// HasNumericMin returns a boolean if a field has been set.
func (o *IntegerFieldDefinition) HasNumericMin() bool {
	if o != nil && o.NumericMin.IsSet() {
		return true
	}

	return false
}

// SetNumericMin gets a reference to the given NullableFloat32 and assigns it to the NumericMin field.
func (o *IntegerFieldDefinition) SetNumericMin(v float32) {
	o.NumericMin.Set(&v)
}
// SetNumericMinNil sets the value for NumericMin to be an explicit nil
func (o *IntegerFieldDefinition) SetNumericMinNil() {
	o.NumericMin.Set(nil)
}

// UnsetNumericMin ensures that no value is present for NumericMin, not even an explicit nil
func (o *IntegerFieldDefinition) UnsetNumericMin() {
	o.NumericMin.Unset()
}

func (o IntegerFieldDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsMulti) {
		toSerialize["isMulti"] = o.IsMulti
	}
	if !isNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.NumericMax.IsSet() {
		toSerialize["numericMax"] = o.NumericMax.Get()
	}
	if o.NumericMin.IsSet() {
		toSerialize["numericMin"] = o.NumericMin.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIntegerFieldDefinition struct {
	value *IntegerFieldDefinition
	isSet bool
}

func (v NullableIntegerFieldDefinition) Get() *IntegerFieldDefinition {
	return v.value
}

func (v *NullableIntegerFieldDefinition) Set(val *IntegerFieldDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerFieldDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerFieldDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerFieldDefinition(val *IntegerFieldDefinition) *NullableIntegerFieldDefinition {
	return &NullableIntegerFieldDefinition{value: val, isSet: true}
}

func (v NullableIntegerFieldDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegerFieldDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


