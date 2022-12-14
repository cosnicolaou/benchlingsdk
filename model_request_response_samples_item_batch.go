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

// RequestResponseSamplesItemBatch struct for RequestResponseSamplesItemBatch
type RequestResponseSamplesItemBatch struct {
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

// NewRequestResponseSamplesItemBatch instantiates a new RequestResponseSamplesItemBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestResponseSamplesItemBatch() *RequestResponseSamplesItemBatch {
	this := RequestResponseSamplesItemBatch{}
	return &this
}

// NewRequestResponseSamplesItemBatchWithDefaults instantiates a new RequestResponseSamplesItemBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestResponseSamplesItemBatchWithDefaults() *RequestResponseSamplesItemBatch {
	this := RequestResponseSamplesItemBatch{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestResponseSamplesItemBatch) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestResponseSamplesItemBatch) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *RequestResponseSamplesItemBatch) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *RequestResponseSamplesItemBatch) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *RequestResponseSamplesItemBatch) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RequestResponseSamplesItemBatch) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetCreator() UserSummary {
	if o == nil || isNil(o.Creator) {
		var ret UserSummary
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetCreatorOk() (*UserSummary, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserSummary and assigns it to the Creator field.
func (o *RequestResponseSamplesItemBatch) SetCreator(v UserSummary) {
	o.Creator = &v
}

// GetDefaultConcentration returns the DefaultConcentration field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetDefaultConcentration() Measurement {
	if o == nil || isNil(o.DefaultConcentration) {
		var ret Measurement
		return ret
	}
	return *o.DefaultConcentration
}

// GetDefaultConcentrationOk returns a tuple with the DefaultConcentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetDefaultConcentrationOk() (*Measurement, bool) {
	if o == nil || isNil(o.DefaultConcentration) {
    return nil, false
	}
	return o.DefaultConcentration, true
}

// HasDefaultConcentration returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasDefaultConcentration() bool {
	if o != nil && !isNil(o.DefaultConcentration) {
		return true
	}

	return false
}

// SetDefaultConcentration gets a reference to the given Measurement and assigns it to the DefaultConcentration field.
func (o *RequestResponseSamplesItemBatch) SetDefaultConcentration(v Measurement) {
	o.DefaultConcentration = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetEntity() BatchEntity {
	if o == nil || isNil(o.Entity) {
		var ret BatchEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetEntityOk() (*BatchEntity, bool) {
	if o == nil || isNil(o.Entity) {
    return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasEntity() bool {
	if o != nil && !isNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given BatchEntity and assigns it to the Entity field.
func (o *RequestResponseSamplesItemBatch) SetEntity(v BatchEntity) {
	o.Entity = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *RequestResponseSamplesItemBatch) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RequestResponseSamplesItemBatch) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *RequestResponseSamplesItemBatch) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestResponseSamplesItemBatch) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestResponseSamplesItemBatch) GetSchema() SchemaProperty2 {
	if o == nil || isNil(o.Schema.Get()) {
		var ret SchemaProperty2
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestResponseSamplesItemBatch) GetSchemaOk() (*SchemaProperty2, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableSchemaProperty2 and assigns it to the Schema field.
func (o *RequestResponseSamplesItemBatch) SetSchema(v SchemaProperty2) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *RequestResponseSamplesItemBatch) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *RequestResponseSamplesItemBatch) UnsetSchema() {
	o.Schema.Unset()
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *RequestResponseSamplesItemBatch) SetWebURL(v string) {
	o.WebURL = &v
}

// GetInaccessibleId returns the InaccessibleId field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetInaccessibleId() string {
	if o == nil || isNil(o.InaccessibleId) {
		var ret string
		return ret
	}
	return *o.InaccessibleId
}

// GetInaccessibleIdOk returns a tuple with the InaccessibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetInaccessibleIdOk() (*string, bool) {
	if o == nil || isNil(o.InaccessibleId) {
    return nil, false
	}
	return o.InaccessibleId, true
}

// HasInaccessibleId returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasInaccessibleId() bool {
	if o != nil && !isNil(o.InaccessibleId) {
		return true
	}

	return false
}

// SetInaccessibleId gets a reference to the given string and assigns it to the InaccessibleId field.
func (o *RequestResponseSamplesItemBatch) SetInaccessibleId(v string) {
	o.InaccessibleId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RequestResponseSamplesItemBatch) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestResponseSamplesItemBatch) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RequestResponseSamplesItemBatch) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RequestResponseSamplesItemBatch) SetType(v string) {
	o.Type = &v
}

func (o RequestResponseSamplesItemBatch) MarshalJSON() ([]byte, error) {
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

type NullableRequestResponseSamplesItemBatch struct {
	value *RequestResponseSamplesItemBatch
	isSet bool
}

func (v NullableRequestResponseSamplesItemBatch) Get() *RequestResponseSamplesItemBatch {
	return v.value
}

func (v *NullableRequestResponseSamplesItemBatch) Set(val *RequestResponseSamplesItemBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestResponseSamplesItemBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestResponseSamplesItemBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestResponseSamplesItemBatch(val *RequestResponseSamplesItemBatch) *NullableRequestResponseSamplesItemBatch {
	return &NullableRequestResponseSamplesItemBatch{value: val, isSet: true}
}

func (v NullableRequestResponseSamplesItemBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestResponseSamplesItemBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


