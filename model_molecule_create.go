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

// MoleculeCreate struct for MoleculeCreate
type MoleculeCreate struct {
	// Aliases to add to the Molecule.
	Aliases []string `json:"aliases,omitempty"`
	// IDs of users to set as the Molecule's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	ChemicalStructure MoleculeBaseRequestChemicalStructure `json:"chemicalStructure"`
	// Custom fields to add to the Molecule. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// Fields to set on the Molecule. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder containing the entity. Can be left empty when registryId is present.
	FolderId *string `json:"folderId,omitempty"`
	// Name of the Molecule. 
	Name string `json:"name"`
	// ID of the Molecule's schema. 
	SchemaId *string `json:"schemaId,omitempty"`
	// Entity registry ID to set for the registered entity. Cannot specify both entityRegistryId and namingStrategy at the same time. 
	EntityRegistryId *string `json:"entityRegistryId,omitempty"`
	NamingStrategy *NamingStrategy `json:"namingStrategy,omitempty"`
	// Registry ID into which entity should be registered. this is the ID of the registry which was configured for a particular organization To get available registryIds, use the [/registries endpoint](#/Registry/listRegistries)  Required in order for entities to be created directly in the registry. 
	RegistryId *string `json:"registryId,omitempty"`
}

// NewMoleculeCreate instantiates a new MoleculeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoleculeCreate(chemicalStructure MoleculeBaseRequestChemicalStructure, name string) *MoleculeCreate {
	this := MoleculeCreate{}
	this.ChemicalStructure = chemicalStructure
	this.Name = name
	return &this
}

// NewMoleculeCreateWithDefaults instantiates a new MoleculeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoleculeCreateWithDefaults() *MoleculeCreate {
	this := MoleculeCreate{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *MoleculeCreate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *MoleculeCreate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *MoleculeCreate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *MoleculeCreate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *MoleculeCreate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *MoleculeCreate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetChemicalStructure returns the ChemicalStructure field value
func (o *MoleculeCreate) GetChemicalStructure() MoleculeBaseRequestChemicalStructure {
	if o == nil {
		var ret MoleculeBaseRequestChemicalStructure
		return ret
	}

	return o.ChemicalStructure
}

// GetChemicalStructureOk returns a tuple with the ChemicalStructure field value
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetChemicalStructureOk() (*MoleculeBaseRequestChemicalStructure, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ChemicalStructure, true
}

// SetChemicalStructure sets field value
func (o *MoleculeCreate) SetChemicalStructure(v MoleculeBaseRequestChemicalStructure) {
	o.ChemicalStructure = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *MoleculeCreate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *MoleculeCreate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *MoleculeCreate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MoleculeCreate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MoleculeCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *MoleculeCreate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *MoleculeCreate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *MoleculeCreate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *MoleculeCreate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetName returns the Name field value
func (o *MoleculeCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MoleculeCreate) SetName(v string) {
	o.Name = v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *MoleculeCreate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *MoleculeCreate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *MoleculeCreate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise.
func (o *MoleculeCreate) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityRegistryId) {
    return nil, false
	}
	return o.EntityRegistryId, true
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *MoleculeCreate) HasEntityRegistryId() bool {
	if o != nil && !isNil(o.EntityRegistryId) {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given string and assigns it to the EntityRegistryId field.
func (o *MoleculeCreate) SetEntityRegistryId(v string) {
	o.EntityRegistryId = &v
}

// GetNamingStrategy returns the NamingStrategy field value if set, zero value otherwise.
func (o *MoleculeCreate) GetNamingStrategy() NamingStrategy {
	if o == nil || isNil(o.NamingStrategy) {
		var ret NamingStrategy
		return ret
	}
	return *o.NamingStrategy
}

// GetNamingStrategyOk returns a tuple with the NamingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetNamingStrategyOk() (*NamingStrategy, bool) {
	if o == nil || isNil(o.NamingStrategy) {
    return nil, false
	}
	return o.NamingStrategy, true
}

// HasNamingStrategy returns a boolean if a field has been set.
func (o *MoleculeCreate) HasNamingStrategy() bool {
	if o != nil && !isNil(o.NamingStrategy) {
		return true
	}

	return false
}

// SetNamingStrategy gets a reference to the given NamingStrategy and assigns it to the NamingStrategy field.
func (o *MoleculeCreate) SetNamingStrategy(v NamingStrategy) {
	o.NamingStrategy = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *MoleculeCreate) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoleculeCreate) GetRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.RegistryId) {
    return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *MoleculeCreate) HasRegistryId() bool {
	if o != nil && !isNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *MoleculeCreate) SetRegistryId(v string) {
	o.RegistryId = &v
}

func (o MoleculeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.AuthorIds) {
		toSerialize["authorIds"] = o.AuthorIds
	}
	if true {
		toSerialize["chemicalStructure"] = o.ChemicalStructure
	}
	if !isNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	if !isNil(o.EntityRegistryId) {
		toSerialize["entityRegistryId"] = o.EntityRegistryId
	}
	if !isNil(o.NamingStrategy) {
		toSerialize["namingStrategy"] = o.NamingStrategy
	}
	if !isNil(o.RegistryId) {
		toSerialize["registryId"] = o.RegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableMoleculeCreate struct {
	value *MoleculeCreate
	isSet bool
}

func (v NullableMoleculeCreate) Get() *MoleculeCreate {
	return v.value
}

func (v *NullableMoleculeCreate) Set(val *MoleculeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableMoleculeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableMoleculeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoleculeCreate(val *MoleculeCreate) *NullableMoleculeCreate {
	return &NullableMoleculeCreate{value: val, isSet: true}
}

func (v NullableMoleculeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoleculeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


