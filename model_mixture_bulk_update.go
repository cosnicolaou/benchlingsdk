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

// MixtureBulkUpdate struct for MixtureBulkUpdate
type MixtureBulkUpdate struct {
	// Aliases to add to the mixture
	Aliases []string `json:"aliases,omitempty"`
	// The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with `units`. Supports scientific notation (1.23e4).
	Amount *string `json:"amount,omitempty"`
	// IDs of users to set as the mixture's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// Custom fields to add to the mixture. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	EntityRegistryId *string `json:"entityRegistryId,omitempty"`
	// Schema fields to set on the mixture. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. If you are setting the parent mixture field here, you must also specify `ingredients` 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder that the entity is moved into
	FolderId *string `json:"folderId,omitempty"`
	// Desired final state for the ingredients on this mixture. Each ingredient you specify will be matched with the existing ingredients on the mixture based on the component entity, and Benchling will create, update, or delete this mixture's ingredients so that the final state of this mixture's ingredients matches your request. Benchling will recognize that any ingredients you specify that match ingredients on the parent mixture (based on component entity) are inherited. This can be seen on the returned `ingredients[i].hasParent` attribute. 
	Ingredients []IngredientWriteParams `json:"ingredients,omitempty"`
	Name *string `json:"name,omitempty"`
	SchemaId *string `json:"schemaId,omitempty"`
	Units NullableMixtureMeasurementUnits `json:"units,omitempty"`
	Id string `json:"id"`
}

// NewMixtureBulkUpdate instantiates a new MixtureBulkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixtureBulkUpdate(id string) *MixtureBulkUpdate {
	this := MixtureBulkUpdate{}
	this.Id = id
	return &this
}

// NewMixtureBulkUpdateWithDefaults instantiates a new MixtureBulkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixtureBulkUpdateWithDefaults() *MixtureBulkUpdate {
	this := MixtureBulkUpdate{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *MixtureBulkUpdate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetAmount() string {
	if o == nil || isNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetAmountOk() (*string, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MixtureBulkUpdate) SetAmount(v string) {
	o.Amount = &v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *MixtureBulkUpdate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *MixtureBulkUpdate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityRegistryId) {
    return nil, false
	}
	return o.EntityRegistryId, true
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasEntityRegistryId() bool {
	if o != nil && !isNil(o.EntityRegistryId) {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given string and assigns it to the EntityRegistryId field.
func (o *MixtureBulkUpdate) SetEntityRegistryId(v string) {
	o.EntityRegistryId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *MixtureBulkUpdate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *MixtureBulkUpdate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetIngredients returns the Ingredients field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetIngredients() []IngredientWriteParams {
	if o == nil || isNil(o.Ingredients) {
		var ret []IngredientWriteParams
		return ret
	}
	return o.Ingredients
}

// GetIngredientsOk returns a tuple with the Ingredients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetIngredientsOk() ([]IngredientWriteParams, bool) {
	if o == nil || isNil(o.Ingredients) {
    return nil, false
	}
	return o.Ingredients, true
}

// HasIngredients returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasIngredients() bool {
	if o != nil && !isNil(o.Ingredients) {
		return true
	}

	return false
}

// SetIngredients gets a reference to the given []IngredientWriteParams and assigns it to the Ingredients field.
func (o *MixtureBulkUpdate) SetIngredients(v []IngredientWriteParams) {
	o.Ingredients = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MixtureBulkUpdate) SetName(v string) {
	o.Name = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *MixtureBulkUpdate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *MixtureBulkUpdate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureBulkUpdate) GetUnits() MixtureMeasurementUnits {
	if o == nil || isNil(o.Units.Get()) {
		var ret MixtureMeasurementUnits
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureBulkUpdate) GetUnitsOk() (*MixtureMeasurementUnits, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *MixtureBulkUpdate) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableMixtureMeasurementUnits and assigns it to the Units field.
func (o *MixtureBulkUpdate) SetUnits(v MixtureMeasurementUnits) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *MixtureBulkUpdate) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *MixtureBulkUpdate) UnsetUnits() {
	o.Units.Unset()
}

// GetId returns the Id field value
func (o *MixtureBulkUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MixtureBulkUpdate) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MixtureBulkUpdate) SetId(v string) {
	o.Id = v
}

func (o MixtureBulkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.AuthorIds) {
		toSerialize["authorIds"] = o.AuthorIds
	}
	if !isNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !isNil(o.EntityRegistryId) {
		toSerialize["entityRegistryId"] = o.EntityRegistryId
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !isNil(o.Ingredients) {
		toSerialize["ingredients"] = o.Ingredients
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	if o.Units.IsSet() {
		toSerialize["units"] = o.Units.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableMixtureBulkUpdate struct {
	value *MixtureBulkUpdate
	isSet bool
}

func (v NullableMixtureBulkUpdate) Get() *MixtureBulkUpdate {
	return v.value
}

func (v *NullableMixtureBulkUpdate) Set(val *MixtureBulkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMixtureBulkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMixtureBulkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixtureBulkUpdate(val *MixtureBulkUpdate) *NullableMixtureBulkUpdate {
	return &NullableMixtureBulkUpdate{value: val, isSet: true}
}

func (v NullableMixtureBulkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixtureBulkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


