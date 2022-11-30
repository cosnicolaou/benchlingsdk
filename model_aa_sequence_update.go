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

// AaSequenceUpdate struct for AaSequenceUpdate
type AaSequenceUpdate struct {
	// Aliases to add to the AA sequence
	Aliases []string `json:"aliases,omitempty"`
	// Amino acids for the AA sequence. 
	AminoAcids *string `json:"aminoAcids,omitempty"`
	// Annotations to create on the AA sequence. 
	Annotations []AaAnnotation `json:"annotations,omitempty"`
	// IDs of users to set as the AA sequence's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// Custom fields to add to the AA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// Fields to set on the AA sequence. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder containing the AA sequence. 
	FolderId *string `json:"folderId,omitempty"`
	// Name of the AA sequence. 
	Name *string `json:"name,omitempty"`
	// ID of the AA sequence's schema. 
	SchemaId *string `json:"schemaId,omitempty"`
	EntityRegistryId *string `json:"entityRegistryId,omitempty"`
}

// NewAaSequenceUpdate instantiates a new AaSequenceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaSequenceUpdate() *AaSequenceUpdate {
	this := AaSequenceUpdate{}
	return &this
}

// NewAaSequenceUpdateWithDefaults instantiates a new AaSequenceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaSequenceUpdateWithDefaults() *AaSequenceUpdate {
	this := AaSequenceUpdate{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *AaSequenceUpdate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAminoAcids returns the AminoAcids field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetAminoAcids() string {
	if o == nil || isNil(o.AminoAcids) {
		var ret string
		return ret
	}
	return *o.AminoAcids
}

// GetAminoAcidsOk returns a tuple with the AminoAcids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetAminoAcidsOk() (*string, bool) {
	if o == nil || isNil(o.AminoAcids) {
    return nil, false
	}
	return o.AminoAcids, true
}

// HasAminoAcids returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasAminoAcids() bool {
	if o != nil && !isNil(o.AminoAcids) {
		return true
	}

	return false
}

// SetAminoAcids gets a reference to the given string and assigns it to the AminoAcids field.
func (o *AaSequenceUpdate) SetAminoAcids(v string) {
	o.AminoAcids = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetAnnotations() []AaAnnotation {
	if o == nil || isNil(o.Annotations) {
		var ret []AaAnnotation
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetAnnotationsOk() ([]AaAnnotation, bool) {
	if o == nil || isNil(o.Annotations) {
    return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []AaAnnotation and assigns it to the Annotations field.
func (o *AaSequenceUpdate) SetAnnotations(v []AaAnnotation) {
	o.Annotations = v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *AaSequenceUpdate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *AaSequenceUpdate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *AaSequenceUpdate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *AaSequenceUpdate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AaSequenceUpdate) SetName(v string) {
	o.Name = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *AaSequenceUpdate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise.
func (o *AaSequenceUpdate) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceUpdate) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityRegistryId) {
    return nil, false
	}
	return o.EntityRegistryId, true
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *AaSequenceUpdate) HasEntityRegistryId() bool {
	if o != nil && !isNil(o.EntityRegistryId) {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given string and assigns it to the EntityRegistryId field.
func (o *AaSequenceUpdate) SetEntityRegistryId(v string) {
	o.EntityRegistryId = &v
}

func (o AaSequenceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.AminoAcids) {
		toSerialize["aminoAcids"] = o.AminoAcids
	}
	if !isNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !isNil(o.AuthorIds) {
		toSerialize["authorIds"] = o.AuthorIds
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
	if !isNil(o.EntityRegistryId) {
		toSerialize["entityRegistryId"] = o.EntityRegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableAaSequenceUpdate struct {
	value *AaSequenceUpdate
	isSet bool
}

func (v NullableAaSequenceUpdate) Get() *AaSequenceUpdate {
	return v.value
}

func (v *NullableAaSequenceUpdate) Set(val *AaSequenceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAaSequenceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAaSequenceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaSequenceUpdate(val *AaSequenceUpdate) *NullableAaSequenceUpdate {
	return &NullableAaSequenceUpdate{value: val, isSet: true}
}

func (v NullableAaSequenceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaSequenceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


