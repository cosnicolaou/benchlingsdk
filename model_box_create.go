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

// BoxCreate struct for BoxCreate
type BoxCreate struct {
	Barcode *string `json:"barcode,omitempty"`
	Fields *map[string]Field `json:"fields,omitempty"`
	Name *string `json:"name,omitempty"`
	ParentStorageId *string `json:"parentStorageId,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	SchemaId string `json:"schemaId"`
}

// NewBoxCreate instantiates a new BoxCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoxCreate(schemaId string) *BoxCreate {
	this := BoxCreate{}
	this.SchemaId = schemaId
	return &this
}

// NewBoxCreateWithDefaults instantiates a new BoxCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxCreateWithDefaults() *BoxCreate {
	this := BoxCreate{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *BoxCreate) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxCreate) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *BoxCreate) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *BoxCreate) SetBarcode(v string) {
	o.Barcode = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *BoxCreate) GetFields() map[string]Field {
	if o == nil || isNil(o.Fields) {
		var ret map[string]Field
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxCreate) GetFieldsOk() (*map[string]Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *BoxCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]Field and assigns it to the Fields field.
func (o *BoxCreate) SetFields(v map[string]Field) {
	o.Fields = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BoxCreate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxCreate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BoxCreate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BoxCreate) SetName(v string) {
	o.Name = &v
}

// GetParentStorageId returns the ParentStorageId field value if set, zero value otherwise.
func (o *BoxCreate) GetParentStorageId() string {
	if o == nil || isNil(o.ParentStorageId) {
		var ret string
		return ret
	}
	return *o.ParentStorageId
}

// GetParentStorageIdOk returns a tuple with the ParentStorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxCreate) GetParentStorageIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentStorageId) {
    return nil, false
	}
	return o.ParentStorageId, true
}

// HasParentStorageId returns a boolean if a field has been set.
func (o *BoxCreate) HasParentStorageId() bool {
	if o != nil && !isNil(o.ParentStorageId) {
		return true
	}

	return false
}

// SetParentStorageId gets a reference to the given string and assigns it to the ParentStorageId field.
func (o *BoxCreate) SetParentStorageId(v string) {
	o.ParentStorageId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BoxCreate) GetProjectId() string {
	if o == nil || isNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxCreate) GetProjectIdOk() (*string, bool) {
	if o == nil || isNil(o.ProjectId) {
    return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BoxCreate) HasProjectId() bool {
	if o != nil && !isNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BoxCreate) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSchemaId returns the SchemaId field value
func (o *BoxCreate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *BoxCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *BoxCreate) SetSchemaId(v string) {
	o.SchemaId = v
}

func (o BoxCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ParentStorageId) {
		toSerialize["parentStorageId"] = o.ParentStorageId
	}
	if !isNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableBoxCreate struct {
	value *BoxCreate
	isSet bool
}

func (v NullableBoxCreate) Get() *BoxCreate {
	return v.value
}

func (v *NullableBoxCreate) Set(val *BoxCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxCreate(val *BoxCreate) *NullableBoxCreate {
	return &NullableBoxCreate{value: val, isSet: true}
}

func (v NullableBoxCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoxCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

