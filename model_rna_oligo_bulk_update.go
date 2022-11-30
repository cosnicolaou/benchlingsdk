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

// RnaOligoBulkUpdate struct for RnaOligoBulkUpdate
type RnaOligoBulkUpdate struct {
	Id *string `json:"id,omitempty"`
	// Aliases to add to the Oligo
	Aliases []string `json:"aliases,omitempty"`
	// IDs of users to set as the Oligo's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// Base pairs of the oligo. 
	Bases *string `json:"bases,omitempty"`
	// Custom fields to add to the Oligo. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// Fields to set on the Oligo. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder containing the Oligo. 
	FolderId *string `json:"folderId,omitempty"`
	// Name of the Oligo. 
	Name *string `json:"name,omitempty"`
	// ID of the oligo's schema. 
	SchemaId *string `json:"schemaId,omitempty"`
}

// NewRnaOligoBulkUpdate instantiates a new RnaOligoBulkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaOligoBulkUpdate() *RnaOligoBulkUpdate {
	this := RnaOligoBulkUpdate{}
	return &this
}

// NewRnaOligoBulkUpdateWithDefaults instantiates a new RnaOligoBulkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaOligoBulkUpdateWithDefaults() *RnaOligoBulkUpdate {
	this := RnaOligoBulkUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RnaOligoBulkUpdate) SetId(v string) {
	o.Id = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *RnaOligoBulkUpdate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *RnaOligoBulkUpdate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetBases returns the Bases field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetBases() string {
	if o == nil || isNil(o.Bases) {
		var ret string
		return ret
	}
	return *o.Bases
}

// GetBasesOk returns a tuple with the Bases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetBasesOk() (*string, bool) {
	if o == nil || isNil(o.Bases) {
    return nil, false
	}
	return o.Bases, true
}

// HasBases returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasBases() bool {
	if o != nil && !isNil(o.Bases) {
		return true
	}

	return false
}

// SetBases gets a reference to the given string and assigns it to the Bases field.
func (o *RnaOligoBulkUpdate) SetBases(v string) {
	o.Bases = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *RnaOligoBulkUpdate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *RnaOligoBulkUpdate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *RnaOligoBulkUpdate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RnaOligoBulkUpdate) SetName(v string) {
	o.Name = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *RnaOligoBulkUpdate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligoBulkUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *RnaOligoBulkUpdate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *RnaOligoBulkUpdate) SetSchemaId(v string) {
	o.SchemaId = &v
}

func (o RnaOligoBulkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.AuthorIds) {
		toSerialize["authorIds"] = o.AuthorIds
	}
	if !isNil(o.Bases) {
		toSerialize["bases"] = o.Bases
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
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableRnaOligoBulkUpdate struct {
	value *RnaOligoBulkUpdate
	isSet bool
}

func (v NullableRnaOligoBulkUpdate) Get() *RnaOligoBulkUpdate {
	return v.value
}

func (v *NullableRnaOligoBulkUpdate) Set(val *RnaOligoBulkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaOligoBulkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaOligoBulkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaOligoBulkUpdate(val *RnaOligoBulkUpdate) *NullableRnaOligoBulkUpdate {
	return &NullableRnaOligoBulkUpdate{value: val, isSet: true}
}

func (v NullableRnaOligoBulkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaOligoBulkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

