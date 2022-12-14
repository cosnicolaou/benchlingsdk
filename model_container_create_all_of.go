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

// ContainerCreateAllOf struct for ContainerCreateAllOf
type ContainerCreateAllOf struct {
	Barcode *string `json:"barcode,omitempty"`
	ProjectId NullableString `json:"projectId,omitempty"`
	SchemaId string `json:"schemaId"`
}

// NewContainerCreateAllOf instantiates a new ContainerCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerCreateAllOf(schemaId string) *ContainerCreateAllOf {
	this := ContainerCreateAllOf{}
	this.SchemaId = schemaId
	return &this
}

// NewContainerCreateAllOfWithDefaults instantiates a new ContainerCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerCreateAllOfWithDefaults() *ContainerCreateAllOf {
	this := ContainerCreateAllOf{}
	return &this
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *ContainerCreateAllOf) GetBarcode() string {
	if o == nil || isNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerCreateAllOf) GetBarcodeOk() (*string, bool) {
	if o == nil || isNil(o.Barcode) {
    return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *ContainerCreateAllOf) HasBarcode() bool {
	if o != nil && !isNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *ContainerCreateAllOf) SetBarcode(v string) {
	o.Barcode = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerCreateAllOf) GetProjectId() string {
	if o == nil || isNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerCreateAllOf) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *ContainerCreateAllOf) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *ContainerCreateAllOf) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *ContainerCreateAllOf) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *ContainerCreateAllOf) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetSchemaId returns the SchemaId field value
func (o *ContainerCreateAllOf) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *ContainerCreateAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *ContainerCreateAllOf) SetSchemaId(v string) {
	o.SchemaId = v
}

func (o ContainerCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableContainerCreateAllOf struct {
	value *ContainerCreateAllOf
	isSet bool
}

func (v NullableContainerCreateAllOf) Get() *ContainerCreateAllOf {
	return v.value
}

func (v *NullableContainerCreateAllOf) Set(val *ContainerCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerCreateAllOf(val *ContainerCreateAllOf) *NullableContainerCreateAllOf {
	return &NullableContainerCreateAllOf{value: val, isSet: true}
}

func (v NullableContainerCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


