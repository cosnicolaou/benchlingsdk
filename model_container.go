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

// Container struct for Container
type Container struct {
	ArchiveRecord NullableContainerArchiveRecord `json:"archiveRecord,omitempty"`
	Barcode NullableString `json:"barcode,omitempty"`
	CheckoutRecord *CheckoutRecord `json:"checkoutRecord,omitempty"`
	Contents []ContainerContent `json:"contents,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator *UserSummary `json:"creator,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentStorageId NullableString `json:"parentStorageId,omitempty"`
	ParentStorageSchema NullableContainerParentStorageSchema `json:"parentStorageSchema,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	Quantity *ContainerQuantity `json:"quantity,omitempty"`
	Schema NullableContainerParentStorageSchema `json:"schema,omitempty"`
	Volume *DeprecatedContainerVolumeForResponse `json:"volume,omitempty"`
	WebURL *string `json:"webURL,omitempty"`
}

// NewContainer instantiates a new Container object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainer() *Container {
	this := Container{}
	return &this
}

// NewContainerWithDefaults instantiates a new Container object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerWithDefaults() *Container {
	this := Container{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetArchiveRecord() ContainerArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret ContainerArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetArchiveRecordOk() (*ContainerArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *Container) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableContainerArchiveRecord and assigns it to the ArchiveRecord field.
func (o *Container) SetArchiveRecord(v ContainerArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *Container) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *Container) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetBarcode returns the Barcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetBarcode() string {
	if o == nil || isNil(o.Barcode.Get()) {
		var ret string
		return ret
	}
	return *o.Barcode.Get()
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetBarcodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Barcode.Get(), o.Barcode.IsSet()
}

// HasBarcode returns a boolean if a field has been set.
func (o *Container) HasBarcode() bool {
	if o != nil && o.Barcode.IsSet() {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given NullableString and assigns it to the Barcode field.
func (o *Container) SetBarcode(v string) {
	o.Barcode.Set(&v)
}
// SetBarcodeNil sets the value for Barcode to be an explicit nil
func (o *Container) SetBarcodeNil() {
	o.Barcode.Set(nil)
}

// UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
func (o *Container) UnsetBarcode() {
	o.Barcode.Unset()
}

// GetCheckoutRecord returns the CheckoutRecord field value if set, zero value otherwise.
func (o *Container) GetCheckoutRecord() CheckoutRecord {
	if o == nil || isNil(o.CheckoutRecord) {
		var ret CheckoutRecord
		return ret
	}
	return *o.CheckoutRecord
}

// GetCheckoutRecordOk returns a tuple with the CheckoutRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetCheckoutRecordOk() (*CheckoutRecord, bool) {
	if o == nil || isNil(o.CheckoutRecord) {
    return nil, false
	}
	return o.CheckoutRecord, true
}

// HasCheckoutRecord returns a boolean if a field has been set.
func (o *Container) HasCheckoutRecord() bool {
	if o != nil && !isNil(o.CheckoutRecord) {
		return true
	}

	return false
}

// SetCheckoutRecord gets a reference to the given CheckoutRecord and assigns it to the CheckoutRecord field.
func (o *Container) SetCheckoutRecord(v CheckoutRecord) {
	o.CheckoutRecord = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *Container) GetContents() []ContainerContent {
	if o == nil || isNil(o.Contents) {
		var ret []ContainerContent
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetContentsOk() ([]ContainerContent, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *Container) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []ContainerContent and assigns it to the Contents field.
func (o *Container) SetContents(v []ContainerContent) {
	o.Contents = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Container) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Container) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Container) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Container) GetCreator() UserSummary {
	if o == nil || isNil(o.Creator) {
		var ret UserSummary
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetCreatorOk() (*UserSummary, bool) {
	if o == nil || isNil(o.Creator) {
    return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Container) HasCreator() bool {
	if o != nil && !isNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given UserSummary and assigns it to the Creator field.
func (o *Container) SetCreator(v UserSummary) {
	o.Creator = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Container) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Container) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *Container) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Container) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Container) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Container) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Container) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Container) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Container) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Container) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Container) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Container) SetName(v string) {
	o.Name = &v
}

// GetParentStorageId returns the ParentStorageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetParentStorageId() string {
	if o == nil || isNil(o.ParentStorageId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentStorageId.Get()
}

// GetParentStorageIdOk returns a tuple with the ParentStorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetParentStorageIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentStorageId.Get(), o.ParentStorageId.IsSet()
}

// HasParentStorageId returns a boolean if a field has been set.
func (o *Container) HasParentStorageId() bool {
	if o != nil && o.ParentStorageId.IsSet() {
		return true
	}

	return false
}

// SetParentStorageId gets a reference to the given NullableString and assigns it to the ParentStorageId field.
func (o *Container) SetParentStorageId(v string) {
	o.ParentStorageId.Set(&v)
}
// SetParentStorageIdNil sets the value for ParentStorageId to be an explicit nil
func (o *Container) SetParentStorageIdNil() {
	o.ParentStorageId.Set(nil)
}

// UnsetParentStorageId ensures that no value is present for ParentStorageId, not even an explicit nil
func (o *Container) UnsetParentStorageId() {
	o.ParentStorageId.Unset()
}

// GetParentStorageSchema returns the ParentStorageSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetParentStorageSchema() ContainerParentStorageSchema {
	if o == nil || isNil(o.ParentStorageSchema.Get()) {
		var ret ContainerParentStorageSchema
		return ret
	}
	return *o.ParentStorageSchema.Get()
}

// GetParentStorageSchemaOk returns a tuple with the ParentStorageSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetParentStorageSchemaOk() (*ContainerParentStorageSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.ParentStorageSchema.Get(), o.ParentStorageSchema.IsSet()
}

// HasParentStorageSchema returns a boolean if a field has been set.
func (o *Container) HasParentStorageSchema() bool {
	if o != nil && o.ParentStorageSchema.IsSet() {
		return true
	}

	return false
}

// SetParentStorageSchema gets a reference to the given NullableContainerParentStorageSchema and assigns it to the ParentStorageSchema field.
func (o *Container) SetParentStorageSchema(v ContainerParentStorageSchema) {
	o.ParentStorageSchema.Set(&v)
}
// SetParentStorageSchemaNil sets the value for ParentStorageSchema to be an explicit nil
func (o *Container) SetParentStorageSchemaNil() {
	o.ParentStorageSchema.Set(nil)
}

// UnsetParentStorageSchema ensures that no value is present for ParentStorageSchema, not even an explicit nil
func (o *Container) UnsetParentStorageSchema() {
	o.ParentStorageSchema.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetProjectId() string {
	if o == nil || isNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *Container) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *Container) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *Container) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *Container) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Container) GetQuantity() ContainerQuantity {
	if o == nil || isNil(o.Quantity) {
		var ret ContainerQuantity
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetQuantityOk() (*ContainerQuantity, bool) {
	if o == nil || isNil(o.Quantity) {
    return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Container) HasQuantity() bool {
	if o != nil && !isNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given ContainerQuantity and assigns it to the Quantity field.
func (o *Container) SetQuantity(v ContainerQuantity) {
	o.Quantity = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Container) GetSchema() ContainerParentStorageSchema {
	if o == nil || isNil(o.Schema.Get()) {
		var ret ContainerParentStorageSchema
		return ret
	}
	return *o.Schema.Get()
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Container) GetSchemaOk() (*ContainerParentStorageSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schema.Get(), o.Schema.IsSet()
}

// HasSchema returns a boolean if a field has been set.
func (o *Container) HasSchema() bool {
	if o != nil && o.Schema.IsSet() {
		return true
	}

	return false
}

// SetSchema gets a reference to the given NullableContainerParentStorageSchema and assigns it to the Schema field.
func (o *Container) SetSchema(v ContainerParentStorageSchema) {
	o.Schema.Set(&v)
}
// SetSchemaNil sets the value for Schema to be an explicit nil
func (o *Container) SetSchemaNil() {
	o.Schema.Set(nil)
}

// UnsetSchema ensures that no value is present for Schema, not even an explicit nil
func (o *Container) UnsetSchema() {
	o.Schema.Unset()
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *Container) GetVolume() DeprecatedContainerVolumeForResponse {
	if o == nil || isNil(o.Volume) {
		var ret DeprecatedContainerVolumeForResponse
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetVolumeOk() (*DeprecatedContainerVolumeForResponse, bool) {
	if o == nil || isNil(o.Volume) {
    return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *Container) HasVolume() bool {
	if o != nil && !isNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given DeprecatedContainerVolumeForResponse and assigns it to the Volume field.
func (o *Container) SetVolume(v DeprecatedContainerVolumeForResponse) {
	o.Volume = &v
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *Container) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Container) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *Container) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *Container) SetWebURL(v string) {
	o.WebURL = &v
}

func (o Container) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if o.Barcode.IsSet() {
		toSerialize["barcode"] = o.Barcode.Get()
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
	if o.ParentStorageSchema.IsSet() {
		toSerialize["parentStorageSchema"] = o.ParentStorageSchema.Get()
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

type NullableContainer struct {
	value *Container
	isSet bool
}

func (v NullableContainer) Get() *Container {
	return v.value
}

func (v *NullableContainer) Set(val *Container) {
	v.value = val
	v.isSet = true
}

func (v NullableContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainer(val *Container) *NullableContainer {
	return &NullableContainer{value: val, isSet: true}
}

func (v NullableContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


