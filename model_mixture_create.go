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

// MixtureCreate struct for MixtureCreate
type MixtureCreate struct {
	// Aliases to add to the mixture
	Aliases []string `json:"aliases,omitempty"`
	// The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with `units`. Supports scientific notation (1.23e4).
	Amount *string `json:"amount,omitempty"`
	// IDs of users to set as the mixture's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// Custom fields to add to the mixture. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// Entity registry ID to set for the registered entity. Cannot specify both entityRegistryId and namingStrategy at the same time. 
	EntityRegistryId *string `json:"entityRegistryId,omitempty"`
	// Schema fields to set on the mixture. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. If you are setting the parent mixture field here, you must also specify `ingredients` 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder containing the entity. Can be left empty when registryId is present.
	FolderId *string `json:"folderId,omitempty"`
	// Desired final state for the ingredients on this mixture. Each ingredient you specify will be matched with the existing ingredients on the mixture based on the component entity, and Benchling will create, update, or delete this mixture's ingredients so that the final state of this mixture's ingredients matches your request. Benchling will recognize that any ingredients you specify that match ingredients on the parent mixture (based on component entity) are inherited. This can be seen on the returned `ingredients[i].hasParent` attribute. 
	Ingredients []IngredientWriteParams `json:"ingredients"`
	Name string `json:"name"`
	SchemaId string `json:"schemaId"`
	Units NullableMixtureMeasurementUnits `json:"units"`
	NamingStrategy *NamingStrategy `json:"namingStrategy,omitempty"`
	// Registry ID into which entity should be registered. this is the ID of the registry which was configured for a particular organization To get available registryIds, use the [/registries endpoint](#/Registry/listRegistries)  Required in order for entities to be created directly in the registry. 
	RegistryId *string `json:"registryId,omitempty"`
}

// NewMixtureCreate instantiates a new MixtureCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixtureCreate(ingredients []IngredientWriteParams, name string, schemaId string, units NullableMixtureMeasurementUnits) *MixtureCreate {
	this := MixtureCreate{}
	this.Ingredients = ingredients
	this.Name = name
	this.SchemaId = schemaId
	this.Units = units
	return &this
}

// NewMixtureCreateWithDefaults instantiates a new MixtureCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixtureCreateWithDefaults() *MixtureCreate {
	this := MixtureCreate{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *MixtureCreate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *MixtureCreate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *MixtureCreate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MixtureCreate) GetAmount() string {
	if o == nil || isNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetAmountOk() (*string, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MixtureCreate) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MixtureCreate) SetAmount(v string) {
	o.Amount = &v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *MixtureCreate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *MixtureCreate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *MixtureCreate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *MixtureCreate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *MixtureCreate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *MixtureCreate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise.
func (o *MixtureCreate) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityRegistryId) {
    return nil, false
	}
	return o.EntityRegistryId, true
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *MixtureCreate) HasEntityRegistryId() bool {
	if o != nil && !isNil(o.EntityRegistryId) {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given string and assigns it to the EntityRegistryId field.
func (o *MixtureCreate) SetEntityRegistryId(v string) {
	o.EntityRegistryId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MixtureCreate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MixtureCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *MixtureCreate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *MixtureCreate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *MixtureCreate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *MixtureCreate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetIngredients returns the Ingredients field value
func (o *MixtureCreate) GetIngredients() []IngredientWriteParams {
	if o == nil {
		var ret []IngredientWriteParams
		return ret
	}

	return o.Ingredients
}

// GetIngredientsOk returns a tuple with the Ingredients field value
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetIngredientsOk() ([]IngredientWriteParams, bool) {
	if o == nil {
    return nil, false
	}
	return o.Ingredients, true
}

// SetIngredients sets field value
func (o *MixtureCreate) SetIngredients(v []IngredientWriteParams) {
	o.Ingredients = v
}

// GetName returns the Name field value
func (o *MixtureCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MixtureCreate) SetName(v string) {
	o.Name = v
}

// GetSchemaId returns the SchemaId field value
func (o *MixtureCreate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *MixtureCreate) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetUnits returns the Units field value
// If the value is explicit nil, the zero value for MixtureMeasurementUnits will be returned
func (o *MixtureCreate) GetUnits() MixtureMeasurementUnits {
	if o == nil || o.Units.Get() == nil {
		var ret MixtureMeasurementUnits
		return ret
	}

	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureCreate) GetUnitsOk() (*MixtureMeasurementUnits, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// SetUnits sets field value
func (o *MixtureCreate) SetUnits(v MixtureMeasurementUnits) {
	o.Units.Set(&v)
}

// GetNamingStrategy returns the NamingStrategy field value if set, zero value otherwise.
func (o *MixtureCreate) GetNamingStrategy() NamingStrategy {
	if o == nil || isNil(o.NamingStrategy) {
		var ret NamingStrategy
		return ret
	}
	return *o.NamingStrategy
}

// GetNamingStrategyOk returns a tuple with the NamingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetNamingStrategyOk() (*NamingStrategy, bool) {
	if o == nil || isNil(o.NamingStrategy) {
    return nil, false
	}
	return o.NamingStrategy, true
}

// HasNamingStrategy returns a boolean if a field has been set.
func (o *MixtureCreate) HasNamingStrategy() bool {
	if o != nil && !isNil(o.NamingStrategy) {
		return true
	}

	return false
}

// SetNamingStrategy gets a reference to the given NamingStrategy and assigns it to the NamingStrategy field.
func (o *MixtureCreate) SetNamingStrategy(v NamingStrategy) {
	o.NamingStrategy = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *MixtureCreate) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureCreate) GetRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.RegistryId) {
    return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *MixtureCreate) HasRegistryId() bool {
	if o != nil && !isNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *MixtureCreate) SetRegistryId(v string) {
	o.RegistryId = &v
}

func (o MixtureCreate) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["ingredients"] = o.Ingredients
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	if true {
		toSerialize["units"] = o.Units.Get()
	}
	if !isNil(o.NamingStrategy) {
		toSerialize["namingStrategy"] = o.NamingStrategy
	}
	if !isNil(o.RegistryId) {
		toSerialize["registryId"] = o.RegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableMixtureCreate struct {
	value *MixtureCreate
	isSet bool
}

func (v NullableMixtureCreate) Get() *MixtureCreate {
	return v.value
}

func (v *NullableMixtureCreate) Set(val *MixtureCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableMixtureCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableMixtureCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixtureCreate(val *MixtureCreate) *NullableMixtureCreate {
	return &NullableMixtureCreate{value: val, isSet: true}
}

func (v NullableMixtureCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixtureCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


