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

// ContainerContentBatch struct for ContainerContentBatch
type ContainerContentBatch struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	// DateTime at which the the result was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *UserSummary `json:"creator,omitempty"`
	DefaultConcentration *Measurement `json:"defaultConcentration,omitempty"`
	Entity *BatchEntity `json:"entity,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Name *string `json:"name,omitempty"`
	Schema NullableSchemaProperty2 `json:"schema,omitempty"`
	WebURL *string `json:"webURL,omitempty"`
	InaccessibleId *string `json:"inaccessibleId,omitempty"`
	// The type of this inaccessible item. Example values: \"custom_entity\", \"container\", \"plate\", \"dna_sequence\" 
	Type *string `json:"type,omitempty"`
}

// NewContainerContentBatch instantiates a new ContainerContentBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContentBatch() *ContainerContentBatch {
	this := ContainerContentBatch{}
	return &this
}

// NewContainerContentBatchWithDefaults instantiates a new ContainerContentBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContentBatchWithDefaults() *ContainerContentBatch {
	this := ContainerContentBatch{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContentBatch) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContentBatch) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *ContainerContentBatch) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *ContainerContentBatch) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *ContainerContentBatch) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ContainerContentBatch) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetCreator() UserSummary {
	if o == nil || isNil(o.Creator) {
		var ret UserSummary
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetCreatorOk() (*UserSummary, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserSummary and assigns it to the Creator field.
func (o *ContainerContentBatch) SetCreator(v UserSummary) {
	o.Creator = &v
}

// GetDefaultConcentration returns the DefaultConcentration field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetDefaultConcentration() Measurement {
	if o == nil || isNil(o.DefaultConcentration) {
		var ret Measurement
		return ret
	}
	return *o.DefaultConcentration
}

// GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetDefaultConcentrationOk() (*Measurement, bool) {
	if o == nil || isNil(o.DefaultConcentration) {
    return nil, false
	}
	return o.DefaultConcentration, true
}

// HasDefaultConcentration returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasDefaultConcentration() bool {
	if o != nil && !isNil(o.DefaultConcentration) {
		return true
	}

	return false
}

// SetDefaultConcentration gets a reference to the given Measurement and assigns it to the DefaultConcentration field.
func (o *ContainerContentBatch) SetDefaultConcentration(v Measurement) {
	o.DefaultConcentration = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetEntity() BatchEntity {
	if o == nil || isNil(o.Entity) {
		var ret BatchEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetEntityOk() (*BatchEntity, bool) {
	if o == nil || isNil(o.Entity) {
    return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasEntity() bool {
	if o != nil && !isNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given BatchEntity and assigns it to the Entity field.
func (o *ContainerContentBatch) SetEntity(v BatchEntity) {
	o.Entity = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *ContainerContentBatch) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerContentBatch) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ContainerContentBatch) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerContentBatch) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContentBatch) GetSchema() SchemaProperty2 {
	if o == nil || isNil(o.Schema.Get()) {
		var ret SchemaProperty2
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContentBatch) GetSchemaOk() (*SchemaProperty2, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableSchemaProperty2 and assigns it to the Schema field.
func (o *ContainerContentBatch) SetSchema(v SchemaProperty2) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *ContainerContentBatch) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *ContainerContentBatch) UnsetSchema() {
	o.Schema.Unset()
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *ContainerContentBatch) SetWebURL(v string) {
	o.WebURL = &v
}

// GetInaccessibleId returns the InaccessibleId field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetInaccessibleId() string {
	if o == nil || isNil(o.InaccessibleId) {
		var ret string
		return ret
	}
	return *o.InaccessibleId
}

// GetInaccessibleIdOk returns a tuple with the InaccessibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetInaccessibleIdOk() (*string, bool) {
	if o == nil || isNil(o.InaccessibleId) {
    return nil, false
	}
	return o.InaccessibleId, true
}

// HasInaccessibleId returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasInaccessibleId() bool {
	if o != nil && !isNil(o.InaccessibleId) {
		return true
	}

	return false
}

// SetInaccessibleId gets a reference to the given string and assigns it to the InaccessibleId field.
func (o *ContainerContentBatch) SetInaccessibleId(v string) {
	o.InaccessibleId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContainerContentBatch) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContentBatch) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContainerContentBatch) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContainerContentBatch) SetType(v string) {
	o.Type = &v
}

func (o ContainerContentBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !isNil(o.DefaultConcentration) {
		toSerialize["defaultConcentration"] = o.DefaultConcentration
	}
	if !isNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	if !isNil(o.InaccessibleId) {
		toSerialize["inaccessibleId"] = o.InaccessibleId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableContainerContentBatch struct {
	value *ContainerContentBatch
	isSet bool
}

func (v NullableContainerContentBatch) Get() *ContainerContentBatch {
	return v.value
}

func (v *NullableContainerContentBatch) Set(val *ContainerContentBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContentBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContentBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContentBatch(val *ContainerContentBatch) *NullableContainerContentBatch {
	return &NullableContainerContentBatch{value: val, isSet: true}
}

func (v NullableContainerContentBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContentBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


