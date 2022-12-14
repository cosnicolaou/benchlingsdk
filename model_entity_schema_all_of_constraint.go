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

// EntitySchemaAllOfConstraint struct for EntitySchemaAllOfConstraint
type EntitySchemaAllOfConstraint struct {
	FieldDefinitionNames []string `json:"fieldDefinitionNames,omitempty"`
	HasUniqueResidues *bool `json:"hasUniqueResidues,omitempty"`
}

// NewEntitySchemaAllOfConstraint instantiates a new EntitySchemaAllOfConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySchemaAllOfConstraint() *EntitySchemaAllOfConstraint {
	this := EntitySchemaAllOfConstraint{}
	return &this
}

// NewEntitySchemaAllOfConstraintWithDefaults instantiates a new EntitySchemaAllOfConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySchemaAllOfConstraintWithDefaults() *EntitySchemaAllOfConstraint {
	this := EntitySchemaAllOfConstraint{}
	return &this
}

// GetFieldDefinitionNames returns the FieldDefinitionNames field value if set, zero value otherwise.
func (o *EntitySchemaAllOfConstraint) GetFieldDefinitionNames() []string {
	if o == nil || isNil(o.FieldDefinitionNames) {
		var ret []string
		return ret
	}
	return o.FieldDefinitionNames
}

// GetFieldDefinitionNamesOk returns a tuple with the FieldDefinitionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOfConstraint) GetFieldDefinitionNamesOk() ([]string, bool) {
	if o == nil || isNil(o.FieldDefinitionNames) {
    return nil, false
	}
	return o.FieldDefinitionNames, true
}

// HasFieldDefinitionNames returns a boolean if a field has been set.
func (o *EntitySchemaAllOfConstraint) HasFieldDefinitionNames() bool {
	if o != nil && !isNil(o.FieldDefinitionNames) {
		return true
	}

	return false
}

// SetFieldDefinitionNames gets a reference to the given []string and assigns it to the FieldDefinitionNames field.
func (o *EntitySchemaAllOfConstraint) SetFieldDefinitionNames(v []string) {
	o.FieldDefinitionNames = v
}

// GetHasUniqueResidues returns the HasUniqueResidues field value if set, zero value otherwise.
func (o *EntitySchemaAllOfConstraint) GetHasUniqueResidues() bool {
	if o == nil || isNil(o.HasUniqueResidues) {
		var ret bool
		return ret
	}
	return *o.HasUniqueResidues
}

// GetHasUniqueResiduesOk returns a tuple with the HasUniqueResidues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOfConstraint) GetHasUniqueResiduesOk() (*bool, bool) {
	if o == nil || isNil(o.HasUniqueResidues) {
    return nil, false
	}
	return o.HasUniqueResidues, true
}

// HasHasUniqueResidues returns a boolean if a field has been set.
func (o *EntitySchemaAllOfConstraint) HasHasUniqueResidues() bool {
	if o != nil && !isNil(o.HasUniqueResidues) {
		return true
	}

	return false
}

// SetHasUniqueResidues gets a reference to the given bool and assigns it to the HasUniqueResidues field.
func (o *EntitySchemaAllOfConstraint) SetHasUniqueResidues(v bool) {
	o.HasUniqueResidues = &v
}

func (o EntitySchemaAllOfConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FieldDefinitionNames) {
		toSerialize["fieldDefinitionNames"] = o.FieldDefinitionNames
	}
	if !isNil(o.HasUniqueResidues) {
		toSerialize["hasUniqueResidues"] = o.HasUniqueResidues
	}
	return json.Marshal(toSerialize)
}

type NullableEntitySchemaAllOfConstraint struct {
	value *EntitySchemaAllOfConstraint
	isSet bool
}

func (v NullableEntitySchemaAllOfConstraint) Get() *EntitySchemaAllOfConstraint {
	return v.value
}

func (v *NullableEntitySchemaAllOfConstraint) Set(val *EntitySchemaAllOfConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySchemaAllOfConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySchemaAllOfConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySchemaAllOfConstraint(val *EntitySchemaAllOfConstraint) *NullableEntitySchemaAllOfConstraint {
	return &NullableEntitySchemaAllOfConstraint{value: val, isSet: true}
}

func (v NullableEntitySchemaAllOfConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySchemaAllOfConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


