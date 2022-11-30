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

// AssayRun struct for AssayRun
type AssayRun struct {
	// The canonical url of the Run in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Creator *UserSummary `json:"creator,omitempty"`
	EntryId NullableString `json:"entryId,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	Id *string `json:"id,omitempty"`
	IsReviewed *bool `json:"isReviewed,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	Schema NullableSchemaProperty1 `json:"schema,omitempty"`
	ValidationComment NullableString `json:"validationComment,omitempty"`
	ValidationStatus *AssayRunValidationStatus `json:"validationStatus,omitempty"`
}

// NewAssayRun instantiates a new AssayRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayRun() *AssayRun {
	this := AssayRun{}
	return &this
}

// NewAssayRunWithDefaults instantiates a new AssayRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayRunWithDefaults() *AssayRun {
	this := AssayRun{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *AssayRun) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *AssayRun) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *AssayRun) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayRun) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayRun) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *AssayRun) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *AssayRun) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *AssayRun) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *AssayRun) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssayRun) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssayRun) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AssayRun) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AssayRun) GetCreator() UserSummary {
	if o == nil || isNil(o.Creator) {
		var ret UserSummary
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetCreatorOk() (*UserSummary, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AssayRun) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserSummary and assigns it to the Creator field.
func (o *AssayRun) SetCreator(v UserSummary) {
	o.Creator = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayRun) GetEntryId() string {
	if o == nil || isNil(o.EntryId.Get()) {
		var ret string
		return ret
	}
	return *o.EntryId.Get()
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayRun) GetEntryIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EntryId.Get(), o.EntryId.IsSet()
}

// HasEntryId returns a boolean if a field has been set.
func (o *AssayRun) HasEntryId() bool {
	if o != nil && o.EntryId.IsSet() {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given NullableString and assigns it to the EntryId field.
func (o *AssayRun) SetEntryId(v string) {
	o.EntryId.Set(&v)
}
// SetEntryIdNil sets the value for EntryId to be an explicit nil
func (o *AssayRun) SetEntryIdNil() {
	o.EntryId.Set(nil)
}

// UnsetEntryId ensures that no value is present for EntryId, not even an explicit nil
func (o *AssayRun) UnsetEntryId() {
	o.EntryId.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *AssayRun) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *AssayRun) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *AssayRun) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssayRun) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssayRun) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssayRun) SetId(v string) {
	o.Id = &v
}

// GetIsReviewed returns the IsReviewed field value if set, zero value otherwise.
func (o *AssayRun) GetIsReviewed() bool {
	if o == nil || isNil(o.IsReviewed) {
		var ret bool
		return ret
	}
	return *o.IsReviewed
}

// GetIsReviewedOk returns a tuple with the IsReviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetIsReviewedOk() (*bool, bool) {
	if o == nil || isNil(o.IsReviewed) {
    return nil, false
	}
	return o.IsReviewed, true
}

// HasIsReviewed returns a boolean if a field has been set.
func (o *AssayRun) HasIsReviewed() bool {
	if o != nil && !isNil(o.IsReviewed) {
		return true
	}

	return false
}

// SetIsReviewed gets a reference to the given bool and assigns it to the IsReviewed field.
func (o *AssayRun) SetIsReviewed(v bool) {
	o.IsReviewed = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayRun) GetProjectId() string {
	if o == nil || isNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayRun) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *AssayRun) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *AssayRun) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *AssayRun) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *AssayRun) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayRun) GetSchema() SchemaProperty1 {
	if o == nil || isNil(o.Schema.Get()) {
		var ret SchemaProperty1
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayRun) GetSchemaOk() (*SchemaProperty1, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *AssayRun) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableSchemaProperty1 and assigns it to the Schema field.
func (o *AssayRun) SetSchema(v SchemaProperty1) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *AssayRun) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *AssayRun) UnsetSchema() {
	o.Schema.Unset()
}

// GetValidationComment returns the ValidationComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssayRun) GetValidationComment() string {
	if o == nil || isNil(o.ValidationComment.Get()) {
		var ret string
		return ret
	}
	return *o.ValidationComment.Get()
}

// GetValidationCommentOk returns a tuple with the ValidationComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssayRun) GetValidationCommentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ValidationComment.Get(), o.ValidationComment.IsSet()
}

// HasValidationComment returns a boolean if a field has been set.
func (o *AssayRun) HasValidationComment() bool {
	if o != nil && o.ValidationComment.IsSet() {
		return true
	}

	return false
}

// SetValidationComment gets a reference to the given NullableString and assigns it to the ValidationComment field.
func (o *AssayRun) SetValidationComment(v string) {
	o.ValidationComment.Set(&v)
}
// SetValidationCommentNil sets the value for ValidationComment to be an explicit nil
func (o *AssayRun) SetValidationCommentNil() {
	o.ValidationComment.Set(nil)
}

// UnsetValidationComment ensures that no value is present for ValidationComment, not even an explicit nil
func (o *AssayRun) UnsetValidationComment() {
	o.ValidationComment.Unset()
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *AssayRun) GetValidationStatus() AssayRunValidationStatus {
	if o == nil || isNil(o.ValidationStatus) {
		var ret AssayRunValidationStatus
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayRun) GetValidationStatusOk() (*AssayRunValidationStatus, bool) {
	if o == nil || isNil(o.ValidationStatus) {
    return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *AssayRun) HasValidationStatus() bool {
	if o != nil && !isNil(o.ValidationStatus) {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given AssayRunValidationStatus and assigns it to the ValidationStatus field.
func (o *AssayRun) SetValidationStatus(v AssayRunValidationStatus) {
	o.ValidationStatus = &v
}

func (o AssayRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
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
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if o.ValidationComment.IsSet() {
		toSerialize["validationComment"] = o.ValidationComment.Get()
	}
	if !isNil(o.ValidationStatus) {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableAssayRun struct {
	value *AssayRun
	isSet bool
}

func (v NullableAssayRun) Get() *AssayRun {
	return v.value
}

func (v *NullableAssayRun) Set(val *AssayRun) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayRun) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayRun(val *AssayRun) *NullableAssayRun {
	return &NullableAssayRun{value: val, isSet: true}
}

func (v NullableAssayRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

