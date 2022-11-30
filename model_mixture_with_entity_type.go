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

// MixtureWithEntityType struct for MixtureWithEntityType
type MixtureWithEntityType struct {
	Aliases []string `json:"aliases,omitempty"`
	// Derived from the mixture's schema.
	AllowMeasuredIngredients *bool `json:"allowMeasuredIngredients,omitempty"`
	// The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with `units`. Supports scientific notation (1.23e4).
	Amount *string `json:"amount,omitempty"`
	// The canonical url of the Mixture in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	Authors []UserSummary `json:"authors,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *CustomEntityCreator `json:"creator,omitempty"`
	CustomFields *map[string]CustomField `json:"customFields,omitempty"`
	EntityRegistryId NullableString `json:"entityRegistryId,omitempty"`
	// Mixtures can have up to one parent mixture field.
	Fields *ModelMap `json:"fields,omitempty"`
	FolderId NullableString `json:"folderId,omitempty"`
	Id *string `json:"id,omitempty"`
	// List of ingredients on this mixture.
	Ingredients []Ingredient `json:"ingredients,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Name *string `json:"name,omitempty"`
	RegistrationOrigin NullableAaSequenceRegistrationOrigin `json:"registrationOrigin,omitempty"`
	RegistryId NullableString `json:"registryId,omitempty"`
	Schema *MixtureSchema `json:"schema,omitempty"`
	Units NullableMixtureMeasurementUnits `json:"units,omitempty"`
	WebURL *string `json:"webURL,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
}

// NewMixtureWithEntityType instantiates a new MixtureWithEntityType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMixtureWithEntityType() *MixtureWithEntityType {
	this := MixtureWithEntityType{}
	return &this
}

// NewMixtureWithEntityTypeWithDefaults instantiates a new MixtureWithEntityType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMixtureWithEntityTypeWithDefaults() *MixtureWithEntityType {
	this := MixtureWithEntityType{}
	return &this
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetAliases() []string {
	if o == nil || isNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetAliasesOk() ([]string, bool) {
	if o == nil || isNil(o.Aliases) {
    return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasAliases() bool {
	if o != nil && !isNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *MixtureWithEntityType) SetAliases(v []string) {
	o.Aliases = v
}

// GetAllowMeasuredIngredients returns the AllowMeasuredIngredients field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetAllowMeasuredIngredients() bool {
	if o == nil || isNil(o.AllowMeasuredIngredients) {
		var ret bool
		return ret
	}
	return *o.AllowMeasuredIngredients
}

// GetAllowMeasuredIngredientsOk returns a tuple with the AllowMeasuredIngredients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetAllowMeasuredIngredientsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowMeasuredIngredients) {
    return nil, false
	}
	return o.AllowMeasuredIngredients, true
}

// HasAllowMeasuredIngredients returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasAllowMeasuredIngredients() bool {
	if o != nil && !isNil(o.AllowMeasuredIngredients) {
		return true
	}

	return false
}

// SetAllowMeasuredIngredients gets a reference to the given bool and assigns it to the AllowMeasuredIngredients field.
func (o *MixtureWithEntityType) SetAllowMeasuredIngredients(v bool) {
	o.AllowMeasuredIngredients = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetAmount() string {
	if o == nil || isNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetAmountOk() (*string, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MixtureWithEntityType) SetAmount(v string) {
	o.Amount = &v
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *MixtureWithEntityType) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureWithEntityType) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureWithEntityType) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *MixtureWithEntityType) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *MixtureWithEntityType) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *MixtureWithEntityType) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetAuthors() []UserSummary {
	if o == nil || isNil(o.Authors) {
		var ret []UserSummary
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetAuthorsOk() ([]UserSummary, bool) {
	if o == nil || isNil(o.Authors) {
    return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasAuthors() bool {
	if o != nil && !isNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []UserSummary and assigns it to the Authors field.
func (o *MixtureWithEntityType) SetAuthors(v []UserSummary) {
	o.Authors = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MixtureWithEntityType) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetCreator() CustomEntityCreator {
	if o == nil || isNil(o.Creator) {
		var ret CustomEntityCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetCreatorOk() (*CustomEntityCreator, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given CustomEntityCreator and assigns it to the Creator field.
func (o *MixtureWithEntityType) SetCreator(v CustomEntityCreator) {
	o.Creator = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetCustomFields() map[string]CustomField {
	if o == nil || isNil(o.CustomFields) {
		var ret map[string]CustomField
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetCustomFieldsOk() (*map[string]CustomField, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]CustomField and assigns it to the CustomFields field.
func (o *MixtureWithEntityType) SetCustomFields(v map[string]CustomField) {
	o.CustomFields = &v
}

// GetEntityRegistryId returns the EntityRegistryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureWithEntityType) GetEntityRegistryId() string {
	if o == nil || isNil(o.EntityRegistryId.Get()) {
		var ret string
		return ret
	}
	return *o.EntityRegistryId.Get()
}

// GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureWithEntityType) GetEntityRegistryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntityRegistryId.Get(), o.EntityRegistryId.IsSet()
}

// HasEntityRegistryId returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasEntityRegistryId() bool {
	if o != nil && o.EntityRegistryId.IsSet() {
		return true
	}

	return false
}

// SetEntityRegistryId gets a reference to the given NullableString and assigns it to the EntityRegistryId field.
func (o *MixtureWithEntityType) SetEntityRegistryId(v string) {
	o.EntityRegistryId.Set(&v)
}
// SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil
func (o *MixtureWithEntityType) SetEntityRegistryIdNil() {
	o.EntityRegistryId.Set(nil)
}

// UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
func (o *MixtureWithEntityType) UnsetEntityRegistryId() {
	o.EntityRegistryId.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *MixtureWithEntityType) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureWithEntityType) GetFolderId() string {
	if o == nil || isNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureWithEntityType) GetFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *MixtureWithEntityType) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *MixtureWithEntityType) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *MixtureWithEntityType) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MixtureWithEntityType) SetId(v string) {
	o.Id = &v
}

// GetIngredients returns the Ingredients field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetIngredients() []Ingredient {
	if o == nil || isNil(o.Ingredients) {
		var ret []Ingredient
		return ret
	}
	return o.Ingredients
}

// GetIngredientsOk returns a tuple with the Ingredients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetIngredientsOk() ([]Ingredient, bool) {
	if o == nil || isNil(o.Ingredients) {
    return nil, false
	}
	return o.Ingredients, true
}

// HasIngredients returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasIngredients() bool {
	if o != nil && !isNil(o.Ingredients) {
		return true
	}

	return false
}

// SetIngredients gets a reference to the given []Ingredient and assigns it to the Ingredients field.
func (o *MixtureWithEntityType) SetIngredients(v []Ingredient) {
	o.Ingredients = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *MixtureWithEntityType) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MixtureWithEntityType) SetName(v string) {
	o.Name = &v
}

// GetRegistrationOrigin returns the RegistrationOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureWithEntityType) GetRegistrationOrigin() AaSequenceRegistrationOrigin {
	if o == nil || isNil(o.RegistrationOrigin.Get()) {
		var ret AaSequenceRegistrationOrigin
		return ret
	}
	return *o.RegistrationOrigin.Get()
}

// GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureWithEntityType) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool) {
	if o == nil {
    return nil, false
	}
	return o.RegistrationOrigin.Get(), o.RegistrationOrigin.IsSet()
}

// HasRegistrationOrigin returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasRegistrationOrigin() bool {
	if o != nil && o.RegistrationOrigin.IsSet() {
		return true
	}

	return false
}

// SetRegistrationOrigin gets a reference to the given NullableAaSequenceRegistrationOrigin and assigns it to the RegistrationOrigin field.
func (o *MixtureWithEntityType) SetRegistrationOrigin(v AaSequenceRegistrationOrigin) {
	o.RegistrationOrigin.Set(&v)
}
// SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil
func (o *MixtureWithEntityType) SetRegistrationOriginNil() {
	o.RegistrationOrigin.Set(nil)
}

// UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
func (o *MixtureWithEntityType) UnsetRegistrationOrigin() {
	o.RegistrationOrigin.Unset()
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureWithEntityType) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId.Get()) {
		var ret string
		return ret
	}
	return *o.RegistryId.Get()
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureWithEntityType) GetRegistryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RegistryId.Get(), o.RegistryId.IsSet()
}

// HasRegistryId returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasRegistryId() bool {
	if o != nil && o.RegistryId.IsSet() {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given NullableString and assigns it to the RegistryId field.
func (o *MixtureWithEntityType) SetRegistryId(v string) {
	o.RegistryId.Set(&v)
}
// SetRegistryIdNil sets the value for RegistryId to be an explicit nil
func (o *MixtureWithEntityType) SetRegistryIdNil() {
	o.RegistryId.Set(nil)
}

// UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
func (o *MixtureWithEntityType) UnsetRegistryId() {
	o.RegistryId.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetSchema() MixtureSchema {
	if o == nil || isNil(o.Schema) {
		var ret MixtureSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetSchemaOk() (*MixtureSchema, bool) {
	if o == nil || isNil(o.Schema) {
    return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasSchema() bool {
	if o != nil && !isNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given MixtureSchema and assigns it to the Schema field.
func (o *MixtureWithEntityType) SetSchema(v MixtureSchema) {
	o.Schema = &v
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MixtureWithEntityType) GetUnits() MixtureMeasurementUnits {
	if o == nil || isNil(o.Units.Get()) {
		var ret MixtureMeasurementUnits
		return ret
	}
	return *o.Units.Get()
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MixtureWithEntityType) GetUnitsOk() (*MixtureMeasurementUnits, bool) {
	if o == nil {
    return nil, false
	}
	return o.Units.Get(), o.Units.IsSet()
}

// HasUnits returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasUnits() bool {
	if o != nil && o.Units.IsSet() {
		return true
	}

	return false
}

// SetUnits gets a reference to the given NullableMixtureMeasurementUnits and assigns it to the Units field.
func (o *MixtureWithEntityType) SetUnits(v MixtureMeasurementUnits) {
	o.Units.Set(&v)
}
// SetUnitsNil sets the value for Units to be an explicit nil
func (o *MixtureWithEntityType) SetUnitsNil() {
	o.Units.Set(nil)
}

// UnsetUnits ensures that no value is present for Units, not even an explicit nil
func (o *MixtureWithEntityType) UnsetUnits() {
	o.Units.Unset()
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *MixtureWithEntityType) SetWebURL(v string) {
	o.WebURL = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *MixtureWithEntityType) GetEntityType() string {
	if o == nil || isNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MixtureWithEntityType) GetEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.EntityType) {
    return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *MixtureWithEntityType) HasEntityType() bool {
	if o != nil && !isNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *MixtureWithEntityType) SetEntityType(v string) {
	o.EntityType = &v
}

func (o MixtureWithEntityType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !isNil(o.AllowMeasuredIngredients) {
		toSerialize["allowMeasuredIngredients"] = o.AllowMeasuredIngredients
	}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
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
	if !isNil(o.Ingredients) {
		toSerialize["ingredients"] = o.Ingredients
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
	if !isNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if o.Units.IsSet() {
		toSerialize["units"] = o.Units.Get()
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	if !isNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	return json.Marshal(toSerialize)
}

type NullableMixtureWithEntityType struct {
	value *MixtureWithEntityType
	isSet bool
}

func (v NullableMixtureWithEntityType) Get() *MixtureWithEntityType {
	return v.value
}

func (v *NullableMixtureWithEntityType) Set(val *MixtureWithEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableMixtureWithEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableMixtureWithEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMixtureWithEntityType(val *MixtureWithEntityType) *NullableMixtureWithEntityType {
	return &NullableMixtureWithEntityType{value: val, isSet: true}
}

func (v NullableMixtureWithEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMixtureWithEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


