/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
	"time"
)

// CustomEntity struct for CustomEntity
type CustomEntity struct {
	Aliases []string `json:"aliases,omitempty"`
	// The canonical url of the Custom Entity in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	Authors []UserSummary `json:"authors,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *CustomEntityCreator `json:"creator,omitempty"`
	CustomFields *map[string]CustomField `json:"customFields,omitempty"`
	EntityRegistryId NullableString `json:"entityRegistryId,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	FolderId NullableString `json:"folderId,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Name *string `json:"name,omitempty"`
	RegistrationOrigin NullableAaSequenceRegistrationOrigin `json:"registrationOrigin,omitempty"`
	RegistryId NullableString `json:"registryId,omitempty"`
	Schema NullableCustomEntitySchema `json:"schema,omitempty"`
	WebURL *string `json:"webURL,omitempty"`
}

// NewCustomEntity instantiates a new CustomEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEntity() *CustomEntity {
	this := CustomEntity{}
	return &this
}

// NewCustomEntityWithDefaults instantiates a new CustomEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEntityWithDefaults() *CustomEntity {
	this := CustomEntity{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *CustomEntity) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *CustomEntity) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *CustomEntity) SetAliases(v []string) {
	o.Aliases = v
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *CustomEntity) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *CustomEntity) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *CustomEntity) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomEntity) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomEntity) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *CustomEntity) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *CustomEntity) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *CustomEntity) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *CustomEntity) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *CustomEntity) GetAuthors() []UserSummary {
	if o == nil || isNil(o.Authors) {
		var ret []UserSummary
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetAuthorsOk() ([]UserSummary, bool) {
	if o == nil || isNil(o.Authors) {
    return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *CustomEntity) HasAuthors() bool {
	if o != nil && !isNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []UserSummary and assigns it to the Authors field.
func (o *CustomEntity) SetAuthors(v []UserSummary) {
	o.Authors = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomEntity) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomEntity) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *CustomEntity) GetCreator() CustomEntityCreator {
	if o == nil || isNil(o.Creator) {
		var ret CustomEntityCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetCreatorOk() (*CustomEntityCreator, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *CustomEntity) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given CustomEntityCreator and assigns it to the Creator field.
func (o *CustomEntity) SetCreator(v CustomEntityCreator) {
	o.Creator = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CustomEntity) GetCustomFields() map[string]CustomField {
	if o == nil || isNil(o.CustomFields) {
		var ret map[string]CustomField
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetCustomFieldsOk() (*map[string]CustomField, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CustomEntity) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]CustomField and assigns it to the CustomFields field.
func (o *CustomEntity) SetCustomFields(v map[string]CustomField) {
	o.CustomFields = &v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomEntity) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId.Get()
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomEntity) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntityRegistryId.Get(), o.EntityRegistryId.IsSet()
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *CustomEntity) HasEntityRegistryId() bool {
	if o != nil && o.EntityRegistryId.IsSet() {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given NullableString and assigns it to the EntityRegistryId field.
func (o *CustomEntity) SetEntityRegistryId(v string) {
	o.EntityRegistryId.Set(&v)
}
// SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil
func (o *CustomEntity) SetEntityRegistryIdNil() {
	o.EntityRegistryId.Set(nil)
}

// UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
func (o *CustomEntity) UnsetEntityRegistryId() {
	o.EntityRegistryId.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CustomEntity) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CustomEntity) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *CustomEntity) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomEntity) GetFolderId() string {
	if o == nil || isNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomEntity) GetFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *CustomEntity) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *CustomEntity) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *CustomEntity) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *CustomEntity) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomEntity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomEntity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomEntity) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *CustomEntity) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *CustomEntity) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *CustomEntity) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomEntity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomEntity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomEntity) SetName(v string) {
	o.Name = &v
}

// GetRegistrationOrigin returns the RegistrationOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomEntity) GetRegistrationOrigin() AaSequenceRegistrationOrigin {
	if o == nil || isNil(o.RegistrationOrigin.Get()) {
		var ret AaSequenceRegistrationOrigin
		return ret
	}
	return *o.RegistrationOrigin.Get()
}

// GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomEntity) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool) {
	if o == nil {
    return nil, false
	}
	return o.RegistrationOrigin.Get(), o.RegistrationOrigin.IsSet()
}

// HasRegistrationOrigin returns a boolean if a field has been set.
func (o *CustomEntity) HasRegistrationOrigin() bool {
	if o != nil && o.RegistrationOrigin.IsSet() {
		return true
	}

	return false
}

// SetRegistrationOrigin gets a reference to the given NullableAaSequenceRegistrationOrigin and assigns it to the RegistrationOrigin field.
func (o *CustomEntity) SetRegistrationOrigin(v AaSequenceRegistrationOrigin) {
	o.RegistrationOrigin.Set(&v)
}
// SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil
func (o *CustomEntity) SetRegistrationOriginNil() {
	o.RegistrationOrigin.Set(nil)
}

// UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
func (o *CustomEntity) UnsetRegistrationOrigin() {
	o.RegistrationOrigin.Unset()
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomEntity) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId.Get()) {
		var ret string
		return ret
	}
	return *o.RegistryId.Get()
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomEntity) GetRegistryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RegistryId.Get(), o.RegistryId.IsSet()
}

// HasRegistryId returns a boolean if a field has been set.
func (o *CustomEntity) HasRegistryId() bool {
	if o != nil && o.RegistryId.IsSet() {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given NullableString and assigns it to the RegistryId field.
func (o *CustomEntity) SetRegistryId(v string) {
	o.RegistryId.Set(&v)
}
// SetRegistryIdNil sets the value for RegistryId to be an explicit nil
func (o *CustomEntity) SetRegistryIdNil() {
	o.RegistryId.Set(nil)
}

// UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
func (o *CustomEntity) UnsetRegistryId() {
	o.RegistryId.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomEntity) GetSchema() CustomEntitySchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret CustomEntitySchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomEntity) GetSchemaOk() (*CustomEntitySchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *CustomEntity) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableCustomEntitySchema and assigns it to the Schema field.
func (o *CustomEntity) SetSchema(v CustomEntitySchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *CustomEntity) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *CustomEntity) UnsetSchema() {
	o.Schema.Unset()
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *CustomEntity) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntity) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *CustomEntity) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *CustomEntity) SetWebURL(v string) {
	o.WebURL = &v
}

func (o CustomEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !isNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if o.EntityRegistryId.IsSet() {
		toSerialize["entityRegistryId"] = o.EntityRegistryId.Get()
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
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
	if o.RegistrationOrigin.IsSet() {
		toSerialize["registrationOrigin"] = o.RegistrationOrigin.Get()
	}
	if o.RegistryId.IsSet() {
		toSerialize["registryId"] = o.RegistryId.Get()
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEntity struct {
	value *CustomEntity
	isSet bool
}

func (v NullableCustomEntity) Get() *CustomEntity {
	return v.value
}

func (v *NullableCustomEntity) Set(val *CustomEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEntity(val *CustomEntity) *NullableCustomEntity {
	return &NullableCustomEntity{value: val, isSet: true}
}

func (v NullableCustomEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


