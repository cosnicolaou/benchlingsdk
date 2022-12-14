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

// AssayResult struct for AssayResult
type AssayResult struct {
	ArchiveRecord NullableAssayResultArchiveRecord `json:"archiveRecord,omitempty"`
	// DateTime at which the the result was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *AssayResultCreator `json:"creator,omitempty"`
	// ID of the entry that this result is attached to
	EntryId NullableString `json:"entryId,omitempty"`
	// Object mapping field names to a UserValidation Resource object for that field. To **set** validation for a result, you *must* use this object. 
	FieldValidation *map[string]UserValidation `json:"fieldValidation,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	// ID of the result
	Id *string `json:"id,omitempty"`
	// Whether or not this result is attached to an accepted entry
	IsReviewed *bool `json:"isReviewed,omitempty"`
	// ID of the project to insert the result into
	ProjectId NullableString `json:"projectId,omitempty"`
	Schema *SchemaProperty `json:"schema,omitempty"`
	ValidationComment *string `json:"validationComment,omitempty"`
	ValidationStatus *string `json:"validationStatus,omitempty"`
}

// NewAssayResult instantiates a new AssayResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResult() *AssayResult {
	this := AssayResult{}
	return &this
}

// NewAssayResultWithDefaults instantiates a new AssayResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultWithDefaults() *AssayResult {
	this := AssayResult{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayResult) GetArchiveRecord() AssayResultArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AssayResultArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayResult) GetArchiveRecordOk() (*AssayResultArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *AssayResult) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAssayResultArchiveRecord and assigns it to the ArchiveRecord field.
func (o *AssayResult) SetArchiveRecord(v AssayResultArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *AssayResult) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *AssayResult) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssayResult) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssayResult) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AssayResult) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AssayResult) GetCreator() AssayResultCreator {
	if o == nil || isNil(o.Creator) {
		var ret AssayResultCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetCreatorOk() (*AssayResultCreator, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AssayResult) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given AssayResultCreator and assigns it to the Creator field.
func (o *AssayResult) SetCreator(v AssayResultCreator) {
	o.Creator = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayResult) GetEntryId() string {
	if o == nil || isNil(o.EntryId.Get()) {
		var ret string
		return ret
	}
	return *o.EntryId.Get()
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayResult) GetEntryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntryId.Get(), o.EntryId.IsSet()
}

// HasEntryId returns a boolean if a field has been set.
func (o *AssayResult) HasEntryId() bool {
	if o != nil && o.EntryId.IsSet() {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given NullableString and assigns it to the EntryId field.
func (o *AssayResult) SetEntryId(v string) {
	o.EntryId.Set(&v)
}
// SetEntryIdNil sets the value for EntryId to be an explicit nil
func (o *AssayResult) SetEntryIdNil() {
	o.EntryId.Set(nil)
}

// UnsetEntryId ensures that no value is present for EntryId, not even an explicit nil
func (o *AssayResult) UnsetEntryId() {
	o.EntryId.Unset()
}

// GetFieldValidation returns the FieldValidation field value if set, zero value otherwise.
func (o *AssayResult) GetFieldValidation() map[string]UserValidation {
	if o == nil || isNil(o.FieldValidation) {
		var ret map[string]UserValidation
		return ret
	}
	return *o.FieldValidation
}

// GetFieldValidationOk returns a tuple with the FieldValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetFieldValidationOk() (*map[string]UserValidation, bool) {
	if o == nil || isNil(o.FieldValidation) {
    return nil, false
	}
	return o.FieldValidation, true
}

// HasFieldValidation returns a boolean if a field has been set.
func (o *AssayResult) HasFieldValidation() bool {
	if o != nil && !isNil(o.FieldValidation) {
		return true
	}

	return false
}

// SetFieldValidation gets a reference to the given map[string]UserValidation and assigns it to the FieldValidation field.
func (o *AssayResult) SetFieldValidation(v map[string]UserValidation) {
	o.FieldValidation = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AssayResult) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AssayResult) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *AssayResult) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssayResult) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssayResult) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssayResult) SetId(v string) {
	o.Id = &v
}

// GetIsReviewed returns the IsReviewed field value if set, zero value otherwise.
func (o *AssayResult) GetIsReviewed() bool {
	if o == nil || isNil(o.IsReviewed) {
		var ret bool
		return ret
	}
	return *o.IsReviewed
}

// GetIsReviewedOk returns a tuple with the IsReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetIsReviewedOk() (*bool, bool) {
	if o == nil || isNil(o.IsReviewed) {
    return nil, false
	}
	return o.IsReviewed, true
}

// HasIsReviewed returns a boolean if a field has been set.
func (o *AssayResult) HasIsReviewed() bool {
	if o != nil && !isNil(o.IsReviewed) {
		return true
	}

	return false
}

// SetIsReviewed gets a reference to the given bool and assigns it to the IsReviewed field.
func (o *AssayResult) SetIsReviewed(v bool) {
	o.IsReviewed = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayResult) GetProjectId() string {
	if o == nil || isNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayResult) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *AssayResult) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *AssayResult) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *AssayResult) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *AssayResult) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AssayResult) GetSchema() SchemaProperty {
	if o == nil || isNil(o.Schema) {
		var ret SchemaProperty
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetSchemaOk() (*SchemaProperty, bool) {
	if o == nil || isNil(o.Schema) {
    return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AssayResult) HasSchema() bool {
	if o != nil && !isNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given SchemaProperty and assigns it to the Schema field.
func (o *AssayResult) SetSchema(v SchemaProperty) {
	o.Schema = &v
}

// GetValidationComment returns the ValidationComment field value if set, zero value otherwise.
func (o *AssayResult) GetValidationComment() string {
	if o == nil || isNil(o.ValidationComment) {
		var ret string
		return ret
	}
	return *o.ValidationComment
}

// GetValidationCommentOk returns a tuple with the ValidationComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetValidationCommentOk() (*string, bool) {
	if o == nil || isNil(o.ValidationComment) {
    return nil, false
	}
	return o.ValidationComment, true
}

// HasValidationComment returns a boolean if a field has been set.
func (o *AssayResult) HasValidationComment() bool {
	if o != nil && !isNil(o.ValidationComment) {
		return true
	}

	return false
}

// SetValidationComment gets a reference to the given string and assigns it to the ValidationComment field.
func (o *AssayResult) SetValidationComment(v string) {
	o.ValidationComment = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *AssayResult) GetValidationStatus() string {
	if o == nil || isNil(o.ValidationStatus) {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResult) GetValidationStatusOk() (*string, bool) {
	if o == nil || isNil(o.ValidationStatus) {
    return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *AssayResult) HasValidationStatus() bool {
	if o != nil && !isNil(o.ValidationStatus) {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *AssayResult) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

func (o AssayResult) MarshalJSON() ([]byte, error) {
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
	if o.EntryId.IsSet() {
		toSerialize["entryId"] = o.EntryId.Get()
	}
	if !isNil(o.FieldValidation) {
		toSerialize["fieldValidation"] = o.FieldValidation
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsReviewed) {
		toSerialize["isReviewed"] = o.IsReviewed
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if !isNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !isNil(o.ValidationComment) {
		toSerialize["validationComment"] = o.ValidationComment
	}
	if !isNil(o.ValidationStatus) {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResult struct {
	value *AssayResult
	isSet bool
}

func (v NullableAssayResult) Get() *AssayResult {
	return v.value
}

func (v *NullableAssayResult) Set(val *AssayResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResult(val *AssayResult) *NullableAssayResult {
	return &NullableAssayResult{value: val, isSet: true}
}

func (v NullableAssayResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


