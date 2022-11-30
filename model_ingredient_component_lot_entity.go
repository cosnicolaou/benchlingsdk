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

// IngredientComponentLotEntity The entity representing the component lot for this ingredient. This is only present if the mixture schema supports component lots in \"storage\" format.
type IngredientComponentLotEntity struct {
	EntityRegistryId NullableString `json:"entityRegistryId,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewIngredientComponentLotEntity instantiates a new IngredientComponentLotEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngredientComponentLotEntity() *IngredientComponentLotEntity {
	this := IngredientComponentLotEntity{}
	return &this
}

// NewIngredientComponentLotEntityWithDefaults instantiates a new IngredientComponentLotEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngredientComponentLotEntityWithDefaults() *IngredientComponentLotEntity {
	this := IngredientComponentLotEntity{}
	return &this
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngredientComponentLotEntity) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId.Get()
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientComponentLotEntity) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntityRegistryId.Get(), o.EntityRegistryId.IsSet()
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *IngredientComponentLotEntity) HasEntityRegistryId() bool {
	if o != nil && o.EntityRegistryId.IsSet() {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given NullableString and assigns it to the EntityRegistryId field.
func (o *IngredientComponentLotEntity) SetEntityRegistryId(v string) {
	o.EntityRegistryId.Set(&v)
}
// SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil
func (o *IngredientComponentLotEntity) SetEntityRegistryIdNil() {
	o.EntityRegistryId.Set(nil)
}

// UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
func (o *IngredientComponentLotEntity) UnsetEntityRegistryId() {
	o.EntityRegistryId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IngredientComponentLotEntity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngredientComponentLotEntity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IngredientComponentLotEntity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IngredientComponentLotEntity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IngredientComponentLotEntity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngredientComponentLotEntity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IngredientComponentLotEntity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IngredientComponentLotEntity) SetName(v string) {
	o.Name = &v
}

func (o IngredientComponentLotEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityRegistryId.IsSet() {
		toSerialize["entityRegistryId"] = o.EntityRegistryId.Get()
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIngredientComponentLotEntity struct {
	value *IngredientComponentLotEntity
	isSet bool
}

func (v NullableIngredientComponentLotEntity) Get() *IngredientComponentLotEntity {
	return v.value
}

func (v *NullableIngredientComponentLotEntity) Set(val *IngredientComponentLotEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableIngredientComponentLotEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableIngredientComponentLotEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngredientComponentLotEntity(val *IngredientComponentLotEntity) *NullableIngredientComponentLotEntity {
	return &NullableIngredientComponentLotEntity{value: val, isSet: true}
}

func (v NullableIngredientComponentLotEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngredientComponentLotEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


