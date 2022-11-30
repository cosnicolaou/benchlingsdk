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

// DnaSequenceUpdate struct for DnaSequenceUpdate
type DnaSequenceUpdate struct {
	// Aliases to add to the DNA sequence
	Aliases []string `json:"aliases,omitempty"`
	// Annotations to create on the DNA sequence. 
	Annotations []DnaAnnotation `json:"annotations,omitempty"`
	// IDs of users to set as the DNA sequence's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// Base pairs for the DNA sequence. 
	Bases *string `json:"bases,omitempty"`
	// Custom fields to add to the DNA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// Fields to set on the DNA sequence. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder containing the DNA sequence. 
	FolderId *string `json:"folderId,omitempty"`
	// Whether the DNA sequence is circular or linear. 
	IsCircular *bool `json:"isCircular,omitempty"`
	// Name of the DNA sequence. 
	Name *string `json:"name,omitempty"`
	Primers []Primer `json:"primers,omitempty"`
	// ID of the DNA sequence's schema. 
	SchemaId *string `json:"schemaId,omitempty"`
	// Translations to create on the DNA sequence. Translations are specified by either a combination of 'start' and 'end' fields, or a list of regions. Both cannot be provided. 
	Translations []Translation `json:"translations,omitempty"`
	EntityRegistryId *string `json:"entityRegistryId,omitempty"`
}

// NewDnaSequenceUpdate instantiates a new DnaSequenceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaSequenceUpdate() *DnaSequenceUpdate {
	this := DnaSequenceUpdate{}
	return &this
}

// NewDnaSequenceUpdateWithDefaults instantiates a new DnaSequenceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaSequenceUpdateWithDefaults() *DnaSequenceUpdate {
	this := DnaSequenceUpdate{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *DnaSequenceUpdate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetAnnotations() []DnaAnnotation {
	if o == nil || isNil(o.Annotations) {
		var ret []DnaAnnotation
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetAnnotationsOk() ([]DnaAnnotation, bool) {
	if o == nil || isNil(o.Annotations) {
    return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []DnaAnnotation and assigns it to the Annotations field.
func (o *DnaSequenceUpdate) SetAnnotations(v []DnaAnnotation) {
	o.Annotations = v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *DnaSequenceUpdate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetBases returns the Bases field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetBases() string {
	if o == nil || isNil(o.Bases) {
		var ret string
		return ret
	}
	return *o.Bases
}

// GetBasesOk returns a tuple with the Bases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetBasesOk() (*string, bool) {
	if o == nil || isNil(o.Bases) {
    return nil, false
	}
	return o.Bases, true
}

// HasBases returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasBases() bool {
	if o != nil && !isNil(o.Bases) {
		return true
	}

	return false
}

// SetBases gets a reference to the given string and assigns it to the Bases field.
func (o *DnaSequenceUpdate) SetBases(v string) {
	o.Bases = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *DnaSequenceUpdate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *DnaSequenceUpdate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *DnaSequenceUpdate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetIsCircular returns the IsCircular field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetIsCircular() bool {
	if o == nil || isNil(o.IsCircular) {
		var ret bool
		return ret
	}
	return *o.IsCircular
}

// GetIsCircularOk returns a tuple with the IsCircular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetIsCircularOk() (*bool, bool) {
	if o == nil || isNil(o.IsCircular) {
    return nil, false
	}
	return o.IsCircular, true
}

// HasIsCircular returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasIsCircular() bool {
	if o != nil && !isNil(o.IsCircular) {
		return true
	}

	return false
}

// SetIsCircular gets a reference to the given bool and assigns it to the IsCircular field.
func (o *DnaSequenceUpdate) SetIsCircular(v bool) {
	o.IsCircular = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnaSequenceUpdate) SetName(v string) {
	o.Name = &v
}

// GetPrimers returns the Primers field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetPrimers() []Primer {
	if o == nil || isNil(o.Primers) {
		var ret []Primer
		return ret
	}
	return o.Primers
}

// GetPrimersOk returns a tuple with the Primers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetPrimersOk() ([]Primer, bool) {
	if o == nil || isNil(o.Primers) {
    return nil, false
	}
	return o.Primers, true
}

// HasPrimers returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasPrimers() bool {
	if o != nil && !isNil(o.Primers) {
		return true
	}

	return false
}

// SetPrimers gets a reference to the given []Primer and assigns it to the Primers field.
func (o *DnaSequenceUpdate) SetPrimers(v []Primer) {
	o.Primers = v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *DnaSequenceUpdate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetTranslations() []Translation {
	if o == nil || isNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || isNil(o.Translations) {
    return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasTranslations() bool {
	if o != nil && !isNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *DnaSequenceUpdate) SetTranslations(v []Translation) {
	o.Translations = v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise.
func (o *DnaSequenceUpdate) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceUpdate) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.EntityRegistryId) {
    return nil, false
	}
	return o.EntityRegistryId, true
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *DnaSequenceUpdate) HasEntityRegistryId() bool {
	if o != nil && !isNil(o.EntityRegistryId) {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given string and assigns it to the EntityRegistryId field.
func (o *DnaSequenceUpdate) SetEntityRegistryId(v string) {
	o.EntityRegistryId = &v
}

func (o DnaSequenceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
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
	if !isNil(o.IsCircular) {
		toSerialize["isCircular"] = o.IsCircular
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Primers) {
		toSerialize["primers"] = o.Primers
	}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	if !isNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	if !isNil(o.EntityRegistryId) {
		toSerialize["entityRegistryId"] = o.EntityRegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableDnaSequenceUpdate struct {
	value *DnaSequenceUpdate
	isSet bool
}

func (v NullableDnaSequenceUpdate) Get() *DnaSequenceUpdate {
	return v.value
}

func (v *NullableDnaSequenceUpdate) Set(val *DnaSequenceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaSequenceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaSequenceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaSequenceUpdate(val *DnaSequenceUpdate) *NullableDnaSequenceUpdate {
	return &NullableDnaSequenceUpdate{value: val, isSet: true}
}

func (v NullableDnaSequenceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaSequenceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

