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

// DnaSequenceBaseRequestForCreate struct for DnaSequenceBaseRequestForCreate
type DnaSequenceBaseRequestForCreate struct {
	// Aliases to add to the DNA sequence
	Aliases []string `json:"aliases,omitempty"`
	// Annotations to create on the DNA sequence. 
	Annotations []DnaAnnotation `json:"annotations,omitempty"`
	// IDs of users to set as the DNA sequence's authors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// Base pairs for the DNA sequence. 
	Bases string `json:"bases"`
	// Custom fields to add to the DNA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// Fields to set on the DNA sequence. Must correspond with the schema's field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder containing the DNA sequence. 
	FolderId *string `json:"folderId,omitempty"`
	// Whether the DNA sequence is circular or linear. 
	IsCircular bool `json:"isCircular"`
	// Name of the DNA sequence. 
	Name string `json:"name"`
	Primers []Primer `json:"primers,omitempty"`
	// ID of the DNA sequence's schema. 
	SchemaId *string `json:"schemaId,omitempty"`
	// Translations to create on the DNA sequence. Translations are specified by either a combination of 'start' and 'end' fields, or a list of regions. Both cannot be provided. 
	Translations []Translation `json:"translations,omitempty"`
}

// NewDnaSequenceBaseRequestForCreate instantiates a new DnaSequenceBaseRequestForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaSequenceBaseRequestForCreate(bases string, isCircular bool, name string) *DnaSequenceBaseRequestForCreate {
	this := DnaSequenceBaseRequestForCreate{}
	this.Bases = bases
	this.IsCircular = isCircular
	this.Name = name
	return &this
}

// NewDnaSequenceBaseRequestForCreateWithDefaults instantiates a new DnaSequenceBaseRequestForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaSequenceBaseRequestForCreateWithDefaults() *DnaSequenceBaseRequestForCreate {
	this := DnaSequenceBaseRequestForCreate{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *DnaSequenceBaseRequestForCreate) SetAliases(v []string) {
	o.Aliases = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetAnnotations() []DnaAnnotation {
	if o == nil || isNil(o.Annotations) {
		var ret []DnaAnnotation
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetAnnotationsOk() ([]DnaAnnotation, bool) {
	if o == nil || isNil(o.Annotations) {
    return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasAnnotations() bool {
	if o != nil && !isNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []DnaAnnotation and assigns it to the Annotations field.
func (o *DnaSequenceBaseRequestForCreate) SetAnnotations(v []DnaAnnotation) {
	o.Annotations = v
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetAuthorIds() []string {
	if o == nil || isNil(o.AuthorIds) {
		var ret []string
		return ret
	}
	return o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetAuthorIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given []string and assigns it to the AuthorIds field.
func (o *DnaSequenceBaseRequestForCreate) SetAuthorIds(v []string) {
	o.AuthorIds = v
}

// GetBases returns the Bases field value
func (o *DnaSequenceBaseRequestForCreate) GetBases() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bases
}

// GetBasesOk returns a tuple with the Bases field value
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetBasesOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bases, true
}

// SetBases sets field value
func (o *DnaSequenceBaseRequestForCreate) SetBases(v string) {
	o.Bases = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *DnaSequenceBaseRequestForCreate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *DnaSequenceBaseRequestForCreate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetFolderId() string {
	if o == nil || isNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetFolderIdOk() (*string, bool) {
	if o == nil || isNil(o.FolderId) {
    return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasFolderId() bool {
	if o != nil && !isNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *DnaSequenceBaseRequestForCreate) SetFolderId(v string) {
	o.FolderId = &v
}

// GetIsCircular returns the IsCircular field value
func (o *DnaSequenceBaseRequestForCreate) GetIsCircular() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCircular
}

// GetIsCircularOk returns a tuple with the IsCircular field value
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetIsCircularOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsCircular, true
}

// SetIsCircular sets field value
func (o *DnaSequenceBaseRequestForCreate) SetIsCircular(v bool) {
	o.IsCircular = v
}

// GetName returns the Name field value
func (o *DnaSequenceBaseRequestForCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DnaSequenceBaseRequestForCreate) SetName(v string) {
	o.Name = v
}

// GetPrimers returns the Primers field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetPrimers() []Primer {
	if o == nil || isNil(o.Primers) {
		var ret []Primer
		return ret
	}
	return o.Primers
}

// GetPrimersOk returns a tuple with the Primers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetPrimersOk() ([]Primer, bool) {
	if o == nil || isNil(o.Primers) {
    return nil, false
	}
	return o.Primers, true
}

// HasPrimers returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasPrimers() bool {
	if o != nil && !isNil(o.Primers) {
		return true
	}

	return false
}

// SetPrimers gets a reference to the given []Primer and assigns it to the Primers field.
func (o *DnaSequenceBaseRequestForCreate) SetPrimers(v []Primer) {
	o.Primers = v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *DnaSequenceBaseRequestForCreate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *DnaSequenceBaseRequestForCreate) GetTranslations() []Translation {
	if o == nil || isNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaSequenceBaseRequestForCreate) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || isNil(o.Translations) {
    return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *DnaSequenceBaseRequestForCreate) HasTranslations() bool {
	if o != nil && !isNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *DnaSequenceBaseRequestForCreate) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o DnaSequenceBaseRequestForCreate) MarshalJSON() ([]byte, error) {
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
	if true {
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
	if true {
		toSerialize["isCircular"] = o.IsCircular
	}
	if true {
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
	return json.Marshal(toSerialize)
}

type NullableDnaSequenceBaseRequestForCreate struct {
	value *DnaSequenceBaseRequestForCreate
	isSet bool
}

func (v NullableDnaSequenceBaseRequestForCreate) Get() *DnaSequenceBaseRequestForCreate {
	return v.value
}

func (v *NullableDnaSequenceBaseRequestForCreate) Set(val *DnaSequenceBaseRequestForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaSequenceBaseRequestForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaSequenceBaseRequestForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaSequenceBaseRequestForCreate(val *DnaSequenceBaseRequestForCreate) *NullableDnaSequenceBaseRequestForCreate {
	return &NullableDnaSequenceBaseRequestForCreate{value: val, isSet: true}
}

func (v NullableDnaSequenceBaseRequestForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaSequenceBaseRequestForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


