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

// RegistrySchema struct for RegistrySchema
type RegistrySchema struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	FieldDefinitions []SchemaFieldDefinitionsInner `json:"fieldDefinitions,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	RegistryId *string `json:"registryId,omitempty"`
}

// NewRegistrySchema instantiates a new RegistrySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrySchema() *RegistrySchema {
	this := RegistrySchema{}
	return &this
}

// NewRegistrySchemaWithDefaults instantiates a new RegistrySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrySchemaWithDefaults() *RegistrySchema {
	this := RegistrySchema{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegistrySchema) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistrySchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *RegistrySchema) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *RegistrySchema) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *RegistrySchema) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *RegistrySchema) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetFieldDefinitions returns the FieldDefinitions field value if set, zero value otherwise.
func (o *RegistrySchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner {
	if o == nil || isNil(o.FieldDefinitions) {
		var ret []SchemaFieldDefinitionsInner
		return ret
	}
	return o.FieldDefinitions
}

// GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySchema) GetFieldDefinitionsOk() ([]SchemaFieldDefinitionsInner, bool) {
	if o == nil || isNil(o.FieldDefinitions) {
    return nil, false
	}
	return o.FieldDefinitions, true
}

// HasFieldDefinitions returns a boolean if a field has been set.
func (o *RegistrySchema) HasFieldDefinitions() bool {
	if o != nil && !isNil(o.FieldDefinitions) {
		return true
	}

	return false
}

// SetFieldDefinitions gets a reference to the given []SchemaFieldDefinitionsInner and assigns it to the FieldDefinitions field.
func (o *RegistrySchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner) {
	o.FieldDefinitions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrySchema) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySchema) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrySchema) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrySchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistrySchema) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySchema) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistrySchema) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistrySchema) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegistrySchema) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySchema) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegistrySchema) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegistrySchema) SetType(v string) {
	o.Type = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *RegistrySchema) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySchema) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *RegistrySchema) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *RegistrySchema) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *RegistrySchema) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrySchema) GetRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.RegistryId) {
    return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *RegistrySchema) HasRegistryId() bool {
	if o != nil && !isNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *RegistrySchema) SetRegistryId(v string) {
	o.RegistryId = &v
}

func (o RegistrySchema) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.RegistryId) {
		toSerialize["registryId"] = o.RegistryId
	}
	return json.Marshal(toSerialize)
}

type NullableRegistrySchema struct {
	value *RegistrySchema
	isSet bool
}

func (v NullableRegistrySchema) Get() *RegistrySchema {
	return v.value
}

func (v *NullableRegistrySchema) Set(val *RegistrySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrySchema(val *RegistrySchema) *NullableRegistrySchema {
	return &NullableRegistrySchema{value: val, isSet: true}
}

func (v NullableRegistrySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

