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

// EntryTemplate Entry templates are templates that users can base new notebook entries off of. 
type EntryTemplate struct {
	// The canonical url of the Entry Template in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	// DateTime the template was created at
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *EntryTemplateCreator `json:"creator,omitempty"`
	CustomFields *map[string]CustomField `json:"customFields,omitempty"`
	// Array of day objects. Each day object has a day index (integer) and notes field (array of notes, expand further for details on note types). 
	Days []EntryTemplateDay `json:"days,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	// ID of the entry template
	Id *string `json:"id,omitempty"`
	// DateTime the template was last modified
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// Title of the template
	Name *string `json:"name,omitempty"`
	Schema NullableSchemaProperty4 `json:"schema,omitempty"`
	// ID of the collection that contains the template
	TemplateCollectionId *string `json:"templateCollectionId,omitempty"`
	// URL of the template
	WebURL *string `json:"webURL,omitempty"`
}

// NewEntryTemplate instantiates a new EntryTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryTemplate() *EntryTemplate {
	this := EntryTemplate{}
	return &this
}

// NewEntryTemplateWithDefaults instantiates a new EntryTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryTemplateWithDefaults() *EntryTemplate {
	this := EntryTemplate{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *EntryTemplate) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *EntryTemplate) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *EntryTemplate) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntryTemplate) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntryTemplate) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EntryTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *EntryTemplate) GetCreator() EntryTemplateCreator {
	if o == nil || isNil(o.Creator) {
		var ret EntryTemplateCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetCreatorOk() (*EntryTemplateCreator, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *EntryTemplate) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given EntryTemplateCreator and assigns it to the Creator field.
func (o *EntryTemplate) SetCreator(v EntryTemplateCreator) {
	o.Creator = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *EntryTemplate) GetCustomFields() map[string]CustomField {
	if o == nil || isNil(o.CustomFields) {
		var ret map[string]CustomField
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetCustomFieldsOk() (*map[string]CustomField, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EntryTemplate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]CustomField and assigns it to the CustomFields field.
func (o *EntryTemplate) SetCustomFields(v map[string]CustomField) {
	o.CustomFields = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *EntryTemplate) GetDays() []EntryTemplateDay {
	if o == nil || isNil(o.Days) {
		var ret []EntryTemplateDay
		return ret
	}
	return o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetDaysOk() ([]EntryTemplateDay, bool) {
	if o == nil || isNil(o.Days) {
    return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *EntryTemplate) HasDays() bool {
	if o != nil && !isNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given []EntryTemplateDay and assigns it to the Days field.
func (o *EntryTemplate) SetDays(v []EntryTemplateDay) {
	o.Days = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *EntryTemplate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *EntryTemplate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *EntryTemplate) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntryTemplate) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntryTemplate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntryTemplate) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *EntryTemplate) GetModifiedAt() string {
	if o == nil || isNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetModifiedAtOk() (*string, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *EntryTemplate) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *EntryTemplate) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntryTemplate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntryTemplate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntryTemplate) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntryTemplate) GetSchema() SchemaProperty4 {
	if o == nil || isNil(o.Schema.Get()) {
		var ret SchemaProperty4
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntryTemplate) GetSchemaOk() (*SchemaProperty4, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *EntryTemplate) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableSchemaProperty4 and assigns it to the Schema field.
func (o *EntryTemplate) SetSchema(v SchemaProperty4) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *EntryTemplate) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *EntryTemplate) UnsetSchema() {
	o.Schema.Unset()
}

// GetTemplateCollectionId returns the TemplateCollectionId field value if set, zero value otherwise.
func (o *EntryTemplate) GetTemplateCollectionId() string {
	if o == nil || isNil(o.TemplateCollectionId) {
		var ret string
		return ret
	}
	return *o.TemplateCollectionId
}

// GetTemplateCollectionIdOk returns a tuple with the TemplateCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetTemplateCollectionIdOk() (*string, bool) {
	if o == nil || isNil(o.TemplateCollectionId) {
    return nil, false
	}
	return o.TemplateCollectionId, true
}

// HasTemplateCollectionId returns a boolean if a field has been set.
func (o *EntryTemplate) HasTemplateCollectionId() bool {
	if o != nil && !isNil(o.TemplateCollectionId) {
		return true
	}

	return false
}

// SetTemplateCollectionId gets a reference to the given string and assigns it to the TemplateCollectionId field.
func (o *EntryTemplate) SetTemplateCollectionId(v string) {
	o.TemplateCollectionId = &v
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *EntryTemplate) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryTemplate) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *EntryTemplate) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *EntryTemplate) SetWebURL(v string) {
	o.WebURL = &v
}

func (o EntryTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
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
	if !isNil(o.Days) {
		toSerialize["days"] = o.Days
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
	if !isNil(o.TemplateCollectionId) {
		toSerialize["templateCollectionId"] = o.TemplateCollectionId
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	return json.Marshal(toSerialize)
}

type NullableEntryTemplate struct {
	value *EntryTemplate
	isSet bool
}

func (v NullableEntryTemplate) Get() *EntryTemplate {
	return v.value
}

func (v *NullableEntryTemplate) Set(val *EntryTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryTemplate(val *EntryTemplate) *NullableEntryTemplate {
	return &NullableEntryTemplate{value: val, isSet: true}
}

func (v NullableEntryTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

