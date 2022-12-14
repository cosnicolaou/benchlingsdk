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

// EntitySchemaAllOf struct for EntitySchemaAllOf
type EntitySchemaAllOf struct {
	Constraint NullableEntitySchemaAllOfConstraint `json:"constraint,omitempty"`
	ContainableType *string `json:"containableType,omitempty"`
	// DateTime the Entity Schema was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewEntitySchemaAllOf instantiates a new EntitySchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySchemaAllOf() *EntitySchemaAllOf {
	this := EntitySchemaAllOf{}
	return &this
}

// NewEntitySchemaAllOfWithDefaults instantiates a new EntitySchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySchemaAllOfWithDefaults() *EntitySchemaAllOf {
	this := EntitySchemaAllOf{}
	return &this
}

// GetConstraint returns the Constraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitySchemaAllOf) GetConstraint() EntitySchemaAllOfConstraint {
	if o == nil || isNil(o.Constraint.Get()) {
		var ret EntitySchemaAllOfConstraint
		return ret
	}
	return *o.Constraint.Get()
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitySchemaAllOf) GetConstraintOk() (*EntitySchemaAllOfConstraint, bool) {
	if o == nil {
    return nil, false
	}
	return o.Constraint.Get(), o.Constraint.IsSet()
}

// HasConstraint returns a boolean if a field has been set.
func (o *EntitySchemaAllOf) HasConstraint() bool {
	if o != nil && o.Constraint.IsSet() {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given NullableEntitySchemaAllOfConstraint and assigns it to the Constraint field.
func (o *EntitySchemaAllOf) SetConstraint(v EntitySchemaAllOfConstraint) {
	o.Constraint.Set(&v)
}
// SetConstraintNil sets the value for Constraint to be an explicit nil
func (o *EntitySchemaAllOf) SetConstraintNil() {
	o.Constraint.Set(nil)
}

// UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
func (o *EntitySchemaAllOf) UnsetConstraint() {
	o.Constraint.Unset()
}

// GetContainableType returns the ContainableType field value if set, zero value otherwise.
func (o *EntitySchemaAllOf) GetContainableType() string {
	if o == nil || isNil(o.ContainableType) {
		var ret string
		return ret
	}
	return *o.ContainableType
}

// GetContainableTypeOk returns a tuple with the ContainableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOf) GetContainableTypeOk() (*string, bool) {
	if o == nil || isNil(o.ContainableType) {
    return nil, false
	}
	return o.ContainableType, true
}

// HasContainableType returns a boolean if a field has been set.
func (o *EntitySchemaAllOf) HasContainableType() bool {
	if o != nil && !isNil(o.ContainableType) {
		return true
	}

	return false
}

// SetContainableType gets a reference to the given string and assigns it to the ContainableType field.
func (o *EntitySchemaAllOf) SetContainableType(v string) {
	o.ContainableType = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *EntitySchemaAllOf) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOf) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *EntitySchemaAllOf) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *EntitySchemaAllOf) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitySchemaAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitySchemaAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntitySchemaAllOf) SetType(v string) {
	o.Type = &v
}

func (o EntitySchemaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraint.IsSet() {
		toSerialize["constraint"] = o.Constraint.Get()
	}
	if !isNil(o.ContainableType) {
		toSerialize["containableType"] = o.ContainableType
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEntitySchemaAllOf struct {
	value *EntitySchemaAllOf
	isSet bool
}

func (v NullableEntitySchemaAllOf) Get() *EntitySchemaAllOf {
	return v.value
}

func (v *NullableEntitySchemaAllOf) Set(val *EntitySchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySchemaAllOf(val *EntitySchemaAllOf) *NullableEntitySchemaAllOf {
	return &NullableEntitySchemaAllOf{value: val, isSet: true}
}

func (v NullableEntitySchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


