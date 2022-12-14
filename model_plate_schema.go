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

// PlateSchema struct for PlateSchema
type PlateSchema struct {
	ArchiveRecord NullableAaSequenceArchiveRecord `json:"archiveRecord,omitempty"`
	FieldDefinitions []SchemaFieldDefinitionsInner `json:"fieldDefinitions,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	RegistryId *string `json:"registryId,omitempty"`
	ContainerSchema NullableBoxSchemaAllOfContainerSchema `json:"containerSchema,omitempty"`
	Height *float32 `json:"height,omitempty"`
	PlateType *string `json:"plateType,omitempty"`
	Width *float32 `json:"width,omitempty"`
}

// NewPlateSchema instantiates a new PlateSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlateSchema() *PlateSchema {
	this := PlateSchema{}
	return &this
}

// NewPlateSchemaWithDefaults instantiates a new PlateSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlateSchemaWithDefaults() *PlateSchema {
	this := PlateSchema{}
	return &this
}

// GetArchiveRecord returns the ArchiveRecord field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlateSchema) GetArchiveRecord() AaSequenceArchiveRecord {
	if o == nil || isNil(o.ArchiveRecord.Get()) {
		var ret AaSequenceArchiveRecord
		return ret
	}
	return *o.ArchiveRecord.Get()
}

// GetArchiveRecordOk returns a tuple with the ArchiveRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlateSchema) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool) {
	if o == nil {
    return nil, false
	}
	return o.ArchiveRecord.Get(), o.ArchiveRecord.IsSet()
}

// HasArchiveRecord returns a boolean if a field has been set.
func (o *PlateSchema) HasArchiveRecord() bool {
	if o != nil && o.ArchiveRecord.IsSet() {
		return true
	}

	return false
}

// SetArchiveRecord gets a reference to the given NullableAaSequenceArchiveRecord and assigns it to the ArchiveRecord field.
func (o *PlateSchema) SetArchiveRecord(v AaSequenceArchiveRecord) {
	o.ArchiveRecord.Set(&v)
}
// SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil
func (o *PlateSchema) SetArchiveRecordNil() {
	o.ArchiveRecord.Set(nil)
}

// UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
func (o *PlateSchema) UnsetArchiveRecord() {
	o.ArchiveRecord.Unset()
}

// GetFieldDefinitions returns the FieldDefinitions field value if set, zero value otherwise.
func (o *PlateSchema) GetFieldDefinitions() []SchemaFieldDefinitionsInner {
	if o == nil || isNil(o.FieldDefinitions) {
		var ret []SchemaFieldDefinitionsInner
		return ret
	}
	return o.FieldDefinitions
}

// GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetFieldDefinitionsOk() ([]SchemaFieldDefinitionsInner, bool) {
	if o == nil || isNil(o.FieldDefinitions) {
    return nil, false
	}
	return o.FieldDefinitions, true
}

// HasFieldDefinitions returns a boolean if a field has been set.
func (o *PlateSchema) HasFieldDefinitions() bool {
	if o != nil && !isNil(o.FieldDefinitions) {
		return true
	}

	return false
}

// SetFieldDefinitions gets a reference to the given []SchemaFieldDefinitionsInner and assigns it to the FieldDefinitions field.
func (o *PlateSchema) SetFieldDefinitions(v []SchemaFieldDefinitionsInner) {
	o.FieldDefinitions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlateSchema) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlateSchema) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlateSchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlateSchema) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlateSchema) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlateSchema) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlateSchema) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlateSchema) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlateSchema) SetType(v string) {
	o.Type = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PlateSchema) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PlateSchema) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PlateSchema) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *PlateSchema) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.RegistryId) {
    return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *PlateSchema) HasRegistryId() bool {
	if o != nil && !isNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *PlateSchema) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetContainerSchema returns the ContainerSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlateSchema) GetContainerSchema() BoxSchemaAllOfContainerSchema {
	if o == nil || isNil(o.ContainerSchema.Get()) {
		var ret BoxSchemaAllOfContainerSchema
		return ret
	}
	return *o.ContainerSchema.Get()
}

// GetContainerSchemaOk returns a tuple with the ContainerSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlateSchema) GetContainerSchemaOk() (*BoxSchemaAllOfContainerSchema, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContainerSchema.Get(), o.ContainerSchema.IsSet()
}

// HasContainerSchema returns a boolean if a field has been set.
func (o *PlateSchema) HasContainerSchema() bool {
	if o != nil && o.ContainerSchema.IsSet() {
		return true
	}

	return false
}

// SetContainerSchema gets a reference to the given NullableBoxSchemaAllOfContainerSchema and assigns it to the ContainerSchema field.
func (o *PlateSchema) SetContainerSchema(v BoxSchemaAllOfContainerSchema) {
	o.ContainerSchema.Set(&v)
}
// SetContainerSchemaNil sets the value for ContainerSchema to be an explicit nil
func (o *PlateSchema) SetContainerSchemaNil() {
	o.ContainerSchema.Set(nil)
}

// UnsetContainerSchema ensures that no value is present for ContainerSchema, not even an explicit nil
func (o *PlateSchema) UnsetContainerSchema() {
	o.ContainerSchema.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *PlateSchema) GetHeight() float32 {
	if o == nil || isNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetHeightOk() (*float32, bool) {
	if o == nil || isNil(o.Height) {
    return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *PlateSchema) HasHeight() bool {
	if o != nil && !isNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *PlateSchema) SetHeight(v float32) {
	o.Height = &v
}

// GetPlateType returns the PlateType field value if set, zero value otherwise.
func (o *PlateSchema) GetPlateType() string {
	if o == nil || isNil(o.PlateType) {
		var ret string
		return ret
	}
	return *o.PlateType
}

// GetPlateTypeOk returns a tuple with the PlateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetPlateTypeOk() (*string, bool) {
	if o == nil || isNil(o.PlateType) {
    return nil, false
	}
	return o.PlateType, true
}

// HasPlateType returns a boolean if a field has been set.
func (o *PlateSchema) HasPlateType() bool {
	if o != nil && !isNil(o.PlateType) {
		return true
	}

	return false
}

// SetPlateType gets a reference to the given string and assigns it to the PlateType field.
func (o *PlateSchema) SetPlateType(v string) {
	o.PlateType = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *PlateSchema) GetWidth() float32 {
	if o == nil || isNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlateSchema) GetWidthOk() (*float32, bool) {
	if o == nil || isNil(o.Width) {
    return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *PlateSchema) HasWidth() bool {
	if o != nil && !isNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *PlateSchema) SetWidth(v float32) {
	o.Width = &v
}

func (o PlateSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveRecord.IsSet() {
		toSerialize["archiveRecord"] = o.ArchiveRecord.Get()
	}
	if !isNil(o.FieldDefinitions) {
		toSerialize["fieldDefinitions"] = o.FieldDefinitions
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.RegistryId) {
		toSerialize["registryId"] = o.RegistryId
	}
	if o.ContainerSchema.IsSet() {
		toSerialize["containerSchema"] = o.ContainerSchema.Get()
	}
	if !isNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !isNil(o.PlateType) {
		toSerialize["plateType"] = o.PlateType
	}
	if !isNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}

type NullablePlateSchema struct {
	value *PlateSchema
	isSet bool
}

func (v NullablePlateSchema) Get() *PlateSchema {
	return v.value
}

func (v *NullablePlateSchema) Set(val *PlateSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePlateSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePlateSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlateSchema(val *PlateSchema) *NullablePlateSchema {
	return &NullablePlateSchema{value: val, isSet: true}
}

func (v NullablePlateSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlateSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


