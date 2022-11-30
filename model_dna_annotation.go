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

// DnaAnnotation struct for DnaAnnotation
type DnaAnnotation struct {
	// Hex color code used when displaying this feature in the UI.
	Color *string `json:"color,omitempty"`
	CustomFields []SequenceFeatureCustomField `json:"customFields,omitempty"`
	Name *string `json:"name,omitempty"`
	Notes *string `json:"notes,omitempty"`
	End *int32 `json:"end,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Strand *int32 `json:"strand,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDnaAnnotation instantiates a new DnaAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaAnnotation() *DnaAnnotation {
	this := DnaAnnotation{}
	return &this
}

// NewDnaAnnotationWithDefaults instantiates a new DnaAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaAnnotationWithDefaults() *DnaAnnotation {
	this := DnaAnnotation{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *DnaAnnotation) GetColor() string {
	if o == nil || isNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetColorOk() (*string, bool) {
	if o == nil || isNil(o.Color) {
    return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *DnaAnnotation) HasColor() bool {
	if o != nil && !isNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *DnaAnnotation) SetColor(v string) {
	o.Color = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DnaAnnotation) GetCustomFields() []SequenceFeatureCustomField {
	if o == nil || isNil(o.CustomFields) {
		var ret []SequenceFeatureCustomField
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetCustomFieldsOk() ([]SequenceFeatureCustomField, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DnaAnnotation) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []SequenceFeatureCustomField and assigns it to the CustomFields field.
func (o *DnaAnnotation) SetCustomFields(v []SequenceFeatureCustomField) {
	o.CustomFields = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnaAnnotation) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnaAnnotation) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnaAnnotation) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *DnaAnnotation) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *DnaAnnotation) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *DnaAnnotation) SetNotes(v string) {
	o.Notes = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *DnaAnnotation) GetEnd() int32 {
	if o == nil || isNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetEndOk() (*int32, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *DnaAnnotation) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *DnaAnnotation) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *DnaAnnotation) GetStart() int32 {
	if o == nil || isNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetStartOk() (*int32, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *DnaAnnotation) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *DnaAnnotation) SetStart(v int32) {
	o.Start = &v
}

// GetStrand returns the Strand field value if set, zero value otherwise.
func (o *DnaAnnotation) GetStrand() int32 {
	if o == nil || isNil(o.Strand) {
		var ret int32
		return ret
	}
	return *o.Strand
}

// GetStrandOk returns a tuple with the Strand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetStrandOk() (*int32, bool) {
	if o == nil || isNil(o.Strand) {
    return nil, false
	}
	return o.Strand, true
}

// HasStrand returns a boolean if a field has been set.
func (o *DnaAnnotation) HasStrand() bool {
	if o != nil && !isNil(o.Strand) {
		return true
	}

	return false
}

// SetStrand gets a reference to the given int32 and assigns it to the Strand field.
func (o *DnaAnnotation) SetStrand(v int32) {
	o.Strand = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DnaAnnotation) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAnnotation) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DnaAnnotation) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DnaAnnotation) SetType(v string) {
	o.Type = &v
}

func (o DnaAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !isNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.Strand) {
		toSerialize["strand"] = o.Strand
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDnaAnnotation struct {
	value *DnaAnnotation
	isSet bool
}

func (v NullableDnaAnnotation) Get() *DnaAnnotation {
	return v.value
}

func (v *NullableDnaAnnotation) Set(val *DnaAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaAnnotation(val *DnaAnnotation) *NullableDnaAnnotation {
	return &NullableDnaAnnotation{value: val, isSet: true}
}

func (v NullableDnaAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

