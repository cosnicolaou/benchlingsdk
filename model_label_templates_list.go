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

// LabelTemplatesList struct for LabelTemplatesList
type LabelTemplatesList struct {
	LabelTemplates []LabelTemplate `json:"labelTemplates,omitempty"`
}

// NewLabelTemplatesList instantiates a new LabelTemplatesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelTemplatesList() *LabelTemplatesList {
	this := LabelTemplatesList{}
	return &this
}

// NewLabelTemplatesListWithDefaults instantiates a new LabelTemplatesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelTemplatesListWithDefaults() *LabelTemplatesList {
	this := LabelTemplatesList{}
	return &this
}

// GetLabelTemplates returns the LabelTemplates field value if set, zero value otherwise.
func (o *LabelTemplatesList) GetLabelTemplates() []LabelTemplate {
	if o == nil || isNil(o.LabelTemplates) {
		var ret []LabelTemplate
		return ret
	}
	return o.LabelTemplates
}

// GetLabelTemplatesOk returns a tuple with the LabelTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelTemplatesList) GetLabelTemplatesOk() ([]LabelTemplate, bool) {
	if o == nil || isNil(o.LabelTemplates) {
    return nil, false
	}
	return o.LabelTemplates, true
}

// HasLabelTemplates returns a boolean if a field has been set.
func (o *LabelTemplatesList) HasLabelTemplates() bool {
	if o != nil && !isNil(o.LabelTemplates) {
		return true
	}

	return false
}

// SetLabelTemplates gets a reference to the given []LabelTemplate and assigns it to the LabelTemplates field.
func (o *LabelTemplatesList) SetLabelTemplates(v []LabelTemplate) {
	o.LabelTemplates = v
}

func (o LabelTemplatesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LabelTemplates) {
		toSerialize["labelTemplates"] = o.LabelTemplates
	}
	return json.Marshal(toSerialize)
}

type NullableLabelTemplatesList struct {
	value *LabelTemplatesList
	isSet bool
}

func (v NullableLabelTemplatesList) Get() *LabelTemplatesList {
	return v.value
}

func (v *NullableLabelTemplatesList) Set(val *LabelTemplatesList) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelTemplatesList) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelTemplatesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelTemplatesList(val *LabelTemplatesList) *NullableLabelTemplatesList {
	return &NullableLabelTemplatesList{value: val, isSet: true}
}

func (v NullableLabelTemplatesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelTemplatesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


