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

// Well struct for Well
type Well struct {
	ArchiveRecord NullableContainerArchiveRecord `json:"archiveRecord,omitempty"`
	// Barcode of the well
	Barcode *string `json:"barcode,omitempty"`
	CheckoutRecord *CheckoutRecord `json:"checkoutRecord,omitempty"`
	// Array of well contents, each with a batch and concentration
	Contents []ContainerContent `json:"contents,omitempty"`
	// DateTime the well was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *UserSummary `json:"creator,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	// ID of the well
	Id *string `json:"id,omitempty"`
	// DateTime the well was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// Name of the well, defaults to barcode if name is not provided.
	Name *string `json:"name,omitempty"`
	// ID of containing parent storage, a plate well with a coordinate (e.g. plt_2bAks9dx:a2).
	ParentStorageId NullableString `json:"parentStorageId,omitempty"`
	ParentStorageSchema *SchemaSummary `json:"parentStorageSchema,omitempty"`
	// ID of the project if set
	ProjectId NullableString `json:"projectId,omitempty"`
	Quantity *ContainerQuantity `json:"quantity,omitempty"`
	Schema NullableAaSequenceSchema `json:"schema,omitempty"`
	Volume *DeprecatedContainerVolumeForResponse `json:"volume,omitempty"`
	WebURL *string `json:"webURL,omitempty"`
}

// NewWell instantiates a new Well object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWell() *Well {
	this := Well{}
	return &this
}

// NewWellWithDefaults instantiates a new Well object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellWithDefaults() *Well {
	this := Well{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Well) GetArchiveRecord() ContainerArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret ContainerArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Well) GetArchiveRecordOk() (*ContainerArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *Well) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableContainerArchiveRecord and assigns it to the ArchiveRecord field.
func (o *Well) SetArchiveRecord(v ContainerArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *Well) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *Well) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *Well) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *Well) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *Well) SetBarcode(v string) {
	o.Barcode = &v
}

// GetCheckoutRecord returns the CheckoutRecord field value if set, zero value otherwise.
func (o *Well) GetCheckoutRecord() CheckoutRecord {
	if o == nil || isNil(o.CheckoutRecord) {
		var ret CheckoutRecord
		return ret
	}
	return *o.CheckoutRecord
}

// GetCheckoutRecordOk returns a tuple with the CheckoutRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetCheckoutRecordOk() (*CheckoutRecord, bool) {
	if o == nil || isNil(o.CheckoutRecord) {
    return nil, false
	}
	return o.CheckoutRecord, true
}

// HasCheckoutRecord returns a boolean if a field has been set.
func (o *Well) HasCheckoutRecord() bool {
	if o != nil && !isNil(o.CheckoutRecord) {
		return true
	}

	return false
}

// SetCheckoutRecord gets a reference to the given CheckoutRecord and assigns it to the CheckoutRecord field.
func (o *Well) SetCheckoutRecord(v CheckoutRecord) {
	o.CheckoutRecord = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *Well) GetContents() []ContainerContent {
	if o == nil || isNil(o.Contents) {
		var ret []ContainerContent
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetContentsOk() ([]ContainerContent, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *Well) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []ContainerContent and assigns it to the Contents field.
func (o *Well) SetContents(v []ContainerContent) {
	o.Contents = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Well) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Well) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Well) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Well) GetCreator() UserSummary {
	if o == nil || isNil(o.Creator) {
		var ret UserSummary
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetCreatorOk() (*UserSummary, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Well) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserSummary and assigns it to the Creator field.
func (o *Well) SetCreator(v UserSummary) {
	o.Creator = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Well) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Well) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *Well) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Well) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Well) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Well) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Well) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Well) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Well) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Well) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Well) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Well) SetName(v string) {
	o.Name = &v
}

// GetParentStorageId returns the ParentStorageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Well) GetParentStorageId() string {
	if o == nil || isNil(o.ParentStorageId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentStorageId.Get()
}

// GetParentStorageIdOk returns a tuple with the ParentStorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Well) GetParentStorageIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentStorageId.Get(), o.ParentStorageId.IsSet()
}

// HasParentStorageId returns a boolean if a field has been set.
func (o *Well) HasParentStorageId() bool {
	if o != nil && o.ParentStorageId.IsSet() {
		return true
	}

	return false
}

// SetParentStorageId gets a reference to the given NullableString and assigns it to the ParentStorageId field.
func (o *Well) SetParentStorageId(v string) {
	o.ParentStorageId.Set(&v)
}
// SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil
func (o *Well) SetParentStorageIdNil() {
	o.ParentStorageId.Set(nil)
}

// UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
func (o *Well) UnsetParentStorageId() {
	o.ParentStorageId.Unset()
}

// GetParentStorageSchema returns the ParentStorageSchema field value if set, zero value otherwise.
func (o *Well) GetParentStorageSchema() SchemaSummary {
	if o == nil || isNil(o.ParentStorageSchema) {
		var ret SchemaSummary
		return ret
	}
	return *o.ParentStorageSchema
}

// GetParentStorageSchemaOk returns a tuple with the ParentStorageSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetParentStorageSchemaOk() (*SchemaSummary, bool) {
	if o == nil || isNil(o.ParentStorageSchema) {
    return nil, false
	}
	return o.ParentStorageSchema, true
}

// HasParentStorageSchema returns a boolean if a field has been set.
func (o *Well) HasParentStorageSchema() bool {
	if o != nil && !isNil(o.ParentStorageSchema) {
		return true
	}

	return false
}

// SetParentStorageSchema gets a reference to the given SchemaSummary and assigns it to the ParentStorageSchema field.
func (o *Well) SetParentStorageSchema(v SchemaSummary) {
	o.ParentStorageSchema = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Well) GetProjectId() string {
	if o == nil || isNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Well) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *Well) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *Well) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *Well) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *Well) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Well) GetQuantity() ContainerQuantity {
	if o == nil || isNil(o.Quantity) {
		var ret ContainerQuantity
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetQuantityOk() (*ContainerQuantity, bool) {
	if o == nil || isNil(o.Quantity) {
    return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Well) HasQuantity() bool {
	if o != nil && !isNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given ContainerQuantity and assigns it to the Quantity field.
func (o *Well) SetQuantity(v ContainerQuantity) {
	o.Quantity = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Well) GetSchema() AaSequenceSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret AaSequenceSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Well) GetSchemaOk() (*AaSequenceSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *Well) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableAaSequenceSchema and assigns it to the Schema field.
func (o *Well) SetSchema(v AaSequenceSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *Well) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *Well) UnsetSchema() {
	o.Schema.Unset()
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Well) GetVolume() DeprecatedContainerVolumeForResponse {
	if o == nil || isNil(o.Volume) {
		var ret DeprecatedContainerVolumeForResponse
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetVolumeOk() (*DeprecatedContainerVolumeForResponse, bool) {
	if o == nil || isNil(o.Volume) {
    return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Well) HasVolume() bool {
	if o != nil && !isNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given DeprecatedContainerVolumeForResponse and assigns it to the Volume field.
func (o *Well) SetVolume(v DeprecatedContainerVolumeForResponse) {
	o.Volume = &v
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *Well) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Well) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *Well) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *Well) SetWebURL(v string) {
	o.WebURL = &v
}

func (o Well) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !isNil(o.CheckoutRecord) {
		toSerialize["checkoutRecord"] = o.CheckoutRecord
	}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
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
	if !isNil(o.ParentStorageSchema) {
		toSerialize["parentStorageSchema"] = o.ParentStorageSchema
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if !isNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Schema.IsSet() {
		toSerialize["schema"] = o.Schema.Get()
	}
	if !isNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	return json.Marshal(toSerialize)
}

type NullableWell struct {
	value *Well
	isSet bool
}

func (v NullableWell) Get() *Well {
	return v.value
}

func (v *NullableWell) Set(val *Well) {
	v.value = val
	v.isSet = true
}

func (v NullableWell) IsSet() bool {
	return v.isSet
}

func (v *NullableWell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWell(val *Well) *NullableWell {
	return &NullableWell{value: val, isSet: true}
}

func (v NullableWell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


