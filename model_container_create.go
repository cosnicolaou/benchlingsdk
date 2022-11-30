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

// ContainerCreate struct for ContainerCreate
type ContainerCreate struct {
	Fields *map[string]Field `json:"fields,omitempty"`
	Name *string `json:"name,omitempty"`
	// ID of containing parent storage, can also specify a coordinate for plates and boxes (e.g. plt_2bAks9dx:a2).
	ParentStorageId *string `json:"parentStorageId,omitempty"`
	Barcode *string `json:"barcode,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	SchemaId string `json:"schemaId"`
}

// NewContainerCreate instantiates a new ContainerCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerCreate(schemaId string) *ContainerCreate {
	this := ContainerCreate{}
	this.SchemaId = schemaId
	return &this
}

// NewContainerCreateWithDefaults instantiates a new ContainerCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerCreateWithDefaults() *ContainerCreate {
	this := ContainerCreate{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ContainerCreate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ContainerCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *ContainerCreate) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerCreate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerCreate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerCreate) SetName(v string) {
	o.Name = &v
}

// GetParentStorageId returns the ParentStorageId field value if set, zero value otherwise.
func (o *ContainerCreate) GetParentStorageId() string {
	if o == nil || isNil(o.ParentStorageId) {
		var ret string
		return ret
	}
	return *o.ParentStorageId
}

// GetParentStorageIdOk returns a tuple with the ParentStorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreate) GetParentStorageIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentStorageId) {
    return nil, false
	}
	return o.ParentStorageId, true
}

// HasParentStorageId returns a boolean if a field has been set.
func (o *ContainerCreate) HasParentStorageId() bool {
	if o != nil && !isNil(o.ParentStorageId) {
		return true
	}

	return false
}

// SetParentStorageId gets a reference to the given string and assigns it to the ParentStorageId field.
func (o *ContainerCreate) SetParentStorageId(v string) {
	o.ParentStorageId = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *ContainerCreate) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreate) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *ContainerCreate) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *ContainerCreate) SetBarcode(v string) {
	o.Barcode = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerCreate) GetProjectId() string {
	if o == nil || isNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerCreate) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *ContainerCreate) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *ContainerCreate) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *ContainerCreate) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *ContainerCreate) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetSchemaId returns the SchemaId field value
func (o *ContainerCreate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *ContainerCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *ContainerCreate) SetSchemaId(v string) {
	o.SchemaId = v
}

func (o ContainerCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ParentStorageId) {
		toSerialize["parentStorageId"] = o.ParentStorageId
	}
	if !isNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableContainerCreate struct {
	value *ContainerCreate
	isSet bool
}

func (v NullableContainerCreate) Get() *ContainerCreate {
	return v.value
}

func (v *NullableContainerCreate) Set(val *ContainerCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerCreate(val *ContainerCreate) *NullableContainerCreate {
	return &NullableContainerCreate{value: val, isSet: true}
}

func (v NullableContainerCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


