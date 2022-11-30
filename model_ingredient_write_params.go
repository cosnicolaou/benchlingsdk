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

// IngredientWriteParams struct for IngredientWriteParams
type IngredientWriteParams struct {
	// The amount value of this ingredient in its mixture, in string format (to preserve full precision). Pair with `units`. Supports scientific notation (1.23e4). One ingredient on this mixture can have an amount value of `\"QS\"`. 
	Amount NullableString `json:"amount"`
	CatalogIdentifier NullableString `json:"catalogIdentifier"`
	// The entity that uniquely identifies this component.
	ComponentEntityId string `json:"componentEntityId"`
	// The container representing the component lot for this ingredient. This is only writable if the mixture schema supports component lots in \"storage\" format.
	ComponentLotContainerId NullableString `json:"componentLotContainerId"`
	// The entity representing the component lot for this ingredient. This is only writable if the mixture schema supports component lots in \"storage\" format.
	ComponentLotEntityId NullableString `json:"componentLotEntityId"`
	// Text representing the component lot for this ingredient. This is only writable if the mixture schema supports component lots in \"text\" format.
	ComponentLotText NullableString `json:"componentLotText"`
	Notes NullableString `json:"notes"`
	Units NullableIngredientMeasurementUnits `json:"units"`
}

// NewIngredientWriteParams instantiates a new IngredientWriteParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngredientWriteParams(amount NullableString, catalogIdentifier NullableString, componentEntityId string, componentLotContainerId NullableString, componentLotEntityId NullableString, componentLotText NullableString, notes NullableString, units NullableIngredientMeasurementUnits) *IngredientWriteParams {
	this := IngredientWriteParams{}
	this.Amount = amount
	this.CatalogIdentifier = catalogIdentifier
	this.ComponentEntityId = componentEntityId
	this.ComponentLotContainerId = componentLotContainerId
	this.ComponentLotEntityId = componentLotEntityId
	this.ComponentLotText = componentLotText
	this.Notes = notes
	this.Units = units
	return &this
}

// NewIngredientWriteParamsWithDefaults instantiates a new IngredientWriteParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngredientWriteParamsWithDefaults() *IngredientWriteParams {
	this := IngredientWriteParams{}
	return &this
}

// GetAmount returns the Amount field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IngredientWriteParams) GetAmount() string {
	if o == nil || o.Amount.Get() == nil {
		var ret string
		return ret
	}

	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetAmountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// SetAmount sets field value
func (o *IngredientWriteParams) SetAmount(v string) {
	o.Amount.Set(&v)
}

// GetCatalogIdentifier returns the CatalogIdentifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IngredientWriteParams) GetCatalogIdentifier() string {
	if o == nil || o.CatalogIdentifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.CatalogIdentifier.Get()
}

// GetCatalogIdentifierOk returns a tuple with the CatalogIdentifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetCatalogIdentifierOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CatalogIdentifier.Get(), o.CatalogIdentifier.IsSet()
}

// SetCatalogIdentifier sets field value
func (o *IngredientWriteParams) SetCatalogIdentifier(v string) {
	o.CatalogIdentifier.Set(&v)
}

// GetComponentEntityId returns the ComponentEntityId field value
func (o *IngredientWriteParams) GetComponentEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentEntityId
}

// GetComponentEntityIdOk returns a tuple with the ComponentEntityId field value
// and a boolean to check if the value has been set.
func (o *IngredientWriteParams) GetComponentEntityIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ComponentEntityId, true
}

// SetComponentEntityId sets field value
func (o *IngredientWriteParams) SetComponentEntityId(v string) {
	o.ComponentEntityId = v
}

// GetComponentLotContainerId returns the ComponentLotContainerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IngredientWriteParams) GetComponentLotContainerId() string {
	if o == nil || o.ComponentLotContainerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ComponentLotContainerId.Get()
}

// GetComponentLotContainerIdOk returns a tuple with the ComponentLotContainerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetComponentLotContainerIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentLotContainerId.Get(), o.ComponentLotContainerId.IsSet()
}

// SetComponentLotContainerId sets field value
func (o *IngredientWriteParams) SetComponentLotContainerId(v string) {
	o.ComponentLotContainerId.Set(&v)
}

// GetComponentLotEntityId returns the ComponentLotEntityId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IngredientWriteParams) GetComponentLotEntityId() string {
	if o == nil || o.ComponentLotEntityId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ComponentLotEntityId.Get()
}

// GetComponentLotEntityIdOk returns a tuple with the ComponentLotEntityId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetComponentLotEntityIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentLotEntityId.Get(), o.ComponentLotEntityId.IsSet()
}

// SetComponentLotEntityId sets field value
func (o *IngredientWriteParams) SetComponentLotEntityId(v string) {
	o.ComponentLotEntityId.Set(&v)
}

// GetComponentLotText returns the ComponentLotText field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IngredientWriteParams) GetComponentLotText() string {
	if o == nil || o.ComponentLotText.Get() == nil {
		var ret string
		return ret
	}

	return *o.ComponentLotText.Get()
}

// GetComponentLotTextOk returns a tuple with the ComponentLotText field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetComponentLotTextOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentLotText.Get(), o.ComponentLotText.IsSet()
}

// SetComponentLotText sets field value
func (o *IngredientWriteParams) SetComponentLotText(v string) {
	o.ComponentLotText.Set(&v)
}

// GetNotes returns the Notes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IngredientWriteParams) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetNotesOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// SetNotes sets field value
func (o *IngredientWriteParams) SetNotes(v string) {
	o.Notes.Set(&v)
}

// GetUnits returns the Units field value
// If the value is explicit nil, the zero value for IngredientMeasurementUnits will be returned
func (o *IngredientWriteParams) GetUnits() IngredientMeasurementUnits {
	if o == nil || o.Units.Get() == nil {
		var ret IngredientMeasurementUnits
		return ret
	}

	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngredientWriteParams) GetUnitsOk() (*IngredientMeasurementUnits, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// SetUnits sets field value
func (o *IngredientWriteParams) SetUnits(v IngredientMeasurementUnits) {
	o.Units.Set(&v)
}

func (o IngredientWriteParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount.Get()
	}
	if true {
		toSerialize["catalogIdentifier"] = o.CatalogIdentifier.Get()
	}
	if true {
		toSerialize["componentEntityId"] = o.ComponentEntityId
	}
	if true {
		toSerialize["componentLotContainerId"] = o.ComponentLotContainerId.Get()
	}
	if true {
		toSerialize["componentLotEntityId"] = o.ComponentLotEntityId.Get()
	}
	if true {
		toSerialize["componentLotText"] = o.ComponentLotText.Get()
	}
	if true {
		toSerialize["notes"] = o.Notes.Get()
	}
	if true {
		toSerialize["units"] = o.Units.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIngredientWriteParams struct {
	value *IngredientWriteParams
	isSet bool
}

func (v NullableIngredientWriteParams) Get() *IngredientWriteParams {
	return v.value
}

func (v *NullableIngredientWriteParams) Set(val *IngredientWriteParams) {
	v.value = val
	v.isSet = true
}

func (v NullableIngredientWriteParams) IsSet() bool {
	return v.isSet
}

func (v *NullableIngredientWriteParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngredientWriteParams(val *IngredientWriteParams) *NullableIngredientWriteParams {
	return &NullableIngredientWriteParams{value: val, isSet: true}
}

func (v NullableIngredientWriteParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngredientWriteParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


