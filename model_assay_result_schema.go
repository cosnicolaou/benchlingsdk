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

// AssayResultSchema struct for AssayResultSchema
type AssayResultSchema struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	FieldDefinitions []SchemaFieldDefinitionsInner `json:"fieldDefinitions,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	DerivedFrom NullableString `json:"derivedFrom,omitempty"`
	Organization *BaseAssaySchemaAllOfOrganization `json:"organization,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	// DateTime the Assay Result Schema was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
}

// NewAssayResultSchema instantiates a new AssayResultSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResultSchema() *AssayResultSchema {
	this := AssayResultSchema{}
	return &this
}

// NewAssayResultSchemaWithDefaults instantiates a new AssayResultSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultSchemaWithDefaults() *AssayResultSchema {
	this := AssayResultSchema{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayResultSchema) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayResultSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *AssayResultSchema) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *AssayResultSchema) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *AssayResultSchema) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *AssayResultSchema) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetFieldDefinitions returns the FieldDefinitions field value if set, zero value otherwise.
func (o *AssayResultSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner {
	if o == nil || isNil(o.FieldDefinitions) {
		var ret []SchemaFieldDefinitionsInner
		return ret
	}
	return o.FieldDefinitions
}

// GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetFieldDefinitionsOk() ([]SchemaFieldDefinitionsInner, bool) {
	if o == nil || isNil(o.FieldDefinitions) {
    return nil, false
	}
	return o.FieldDefinitions, true
}

// HasFieldDefinitions returns a boolean if a field has been set.
func (o *AssayResultSchema) HasFieldDefinitions() bool {
	if o != nil && !isNil(o.FieldDefinitions) {
		return true
	}

	return false
}

// SetFieldDefinitions gets a reference to the given []SchemaFieldDefinitionsInner and assigns it to the FieldDefinitions field.
func (o *AssayResultSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner) {
	o.FieldDefinitions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssayResultSchema) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssayResultSchema) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssayResultSchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssayResultSchema) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssayResultSchema) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssayResultSchema) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssayResultSchema) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssayResultSchema) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssayResultSchema) SetType(v string) {
	o.Type = &v
}

// GetDerivedFrom returns the DerivedFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayResultSchema) GetDerivedFrom() string {
	if o == nil || isNil(o.DerivedFrom.Get()) {
		var ret string
		return ret
	}
	return *o.DerivedFrom.Get()
}

// GetDerivedFromOk returns a tuple with the DerivedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayResultSchema) GetDerivedFromOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DerivedFrom.Get(), o.DerivedFrom.IsSet()
}

// HasDerivedFrom returns a boolean if a field has been set.
func (o *AssayResultSchema) HasDerivedFrom() bool {
	if o != nil && o.DerivedFrom.IsSet() {
		return true
	}

	return false
}

// SetDerivedFrom gets a reference to the given NullableString and assigns it to the DerivedFrom field.
func (o *AssayResultSchema) SetDerivedFrom(v string) {
	o.DerivedFrom.Set(&v)
}
// SetDerivedFromNil sets the value for DerivedFrom to be an explicit nil
func (o *AssayResultSchema) SetDerivedFromNil() {
	o.DerivedFrom.Set(nil)
}

// UnsetDerivedFrom ensures that no value is present for DerivedFrom, not even an explicit nil
func (o *AssayResultSchema) UnsetDerivedFrom() {
	o.DerivedFrom.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AssayResultSchema) GetOrganization() BaseAssaySchemaAllOfOrganization {
	if o == nil || isNil(o.Organization) {
		var ret BaseAssaySchemaAllOfOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetOrganizationOk() (*BaseAssaySchemaAllOfOrganization, bool) {
	if o == nil || isNil(o.Organization) {
    return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AssayResultSchema) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given BaseAssaySchemaAllOfOrganization and assigns it to the Organization field.
func (o *AssayResultSchema) SetOrganization(v BaseAssaySchemaAllOfOrganization) {
	o.Organization = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *AssayResultSchema) GetSystemName() string {
	if o == nil || isNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetSystemNameOk() (*string, bool) {
	if o == nil || isNil(o.SystemName) {
    return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *AssayResultSchema) HasSystemName() bool {
	if o != nil && !isNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *AssayResultSchema) SetSystemName(v string) {
	o.SystemName = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AssayResultSchema) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultSchema) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AssayResultSchema) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AssayResultSchema) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o AssayResultSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.FieldDefinitions) {
		toSerialize["fieldDefinitions"] = o.FieldDefinitions
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.DerivedFrom.IsSet() {
		toSerialize["derivedFrom"] = o.DerivedFrom.Get()
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.SystemName) {
		toSerialize["systemName"] = o.SystemName
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResultSchema struct {
	value *AssayResultSchema
	isSet bool
}

func (v NullableAssayResultSchema) Get() *AssayResultSchema {
	return v.value
}

func (v *NullableAssayResultSchema) Set(val *AssayResultSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResultSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResultSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResultSchema(val *AssayResultSchema) *NullableAssayResultSchema {
	return &NullableAssayResultSchema{value: val, isSet: true}
}

func (v NullableAssayResultSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResultSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


