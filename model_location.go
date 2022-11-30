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

// Location struct for Location
type Location struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	Barcode *string `json:"barcode,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Creator *UserSummary `json:"creator,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentStorageId NullableString `json:"parentStorageId,omitempty"`
	Schema NullableAaSequenceSchema `json:"schema,omitempty"`
	WebURL *string `json:"webURL,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *Location) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *Location) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *Location) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *Location) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *Location) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *Location) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *Location) SetBarcode(v string) {
	o.Barcode = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Location) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Location) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Location) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Location) GetCreator() UserSummary {
	if o == nil || isNil(o.Creator) {
		var ret UserSummary
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCreatorOk() (*UserSummary, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Location) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserSummary and assigns it to the Creator field.
func (o *Location) SetCreator(v UserSummary) {
	o.Creator = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Location) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Location) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *Location) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Location) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Location) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Location) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Location) GetModifiedAt() string {
	if o == nil || isNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetModifiedAtOk() (*string, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Location) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *Location) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Location) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Location) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Location) SetName(v string) {
	o.Name = &v
}

// GetParentStorageId returns the ParentStorageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetParentStorageId() string {
	if o == nil || isNil(o.ParentStorageId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentStorageId.Get()
}

// GetParentStorageIdOk returns a tuple with the ParentStorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetParentStorageIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentStorageId.Get(), o.ParentStorageId.IsSet()
}

// HasParentStorageId returns a boolean if a field has been set.
func (o *Location) HasParentStorageId() bool {
	if o != nil && o.ParentStorageId.IsSet() {
		return true
	}

	return false
}

// SetParentStorageId gets a reference to the given NullableString and assigns it to the ParentStorageId field.
func (o *Location) SetParentStorageId(v string) {
	o.ParentStorageId.Set(&v)
}
// SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil
func (o *Location) SetParentStorageIdNil() {
	o.ParentStorageId.Set(nil)
}

// UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
func (o *Location) UnsetParentStorageId() {
	o.ParentStorageId.Unset()
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Location) GetSchema() AaSequenceSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret AaSequenceSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Location) GetSchemaOk() (*AaSequenceSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *Location) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableAaSequenceSchema and assigns it to the Schema field.
func (o *Location) SetSchema(v AaSequenceSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *Location) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *Location) UnsetSchema() {
	o.Schema.Unset()
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *Location) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *Location) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *Location) SetWebURL(v string) {
	o.WebURL = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Creator) {
		toSerialize["creator"] = o.Creator
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
	if o.ParentStorageId.IsSet() {
		toSerialize["parentStorageId"] = o.ParentStorageId.Get()
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


