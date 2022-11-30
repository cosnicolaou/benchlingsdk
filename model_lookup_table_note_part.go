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

// LookupTableNotePart struct for LookupTableNotePart
type LookupTableNotePart struct {
	ApiId *string `json:"apiId,omitempty"`
	Columns []StructuredTableColumnInfo `json:"columns,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewLookupTableNotePart instantiates a new LookupTableNotePart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupTableNotePart() *LookupTableNotePart {
	this := LookupTableNotePart{}
	return &this
}

// NewLookupTableNotePartWithDefaults instantiates a new LookupTableNotePart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupTableNotePartWithDefaults() *LookupTableNotePart {
	this := LookupTableNotePart{}
	return &this
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *LookupTableNotePart) GetApiId() string {
	if o == nil || isNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableNotePart) GetApiIdOk() (*string, bool) {
	if o == nil || isNil(o.ApiId) {
    return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *LookupTableNotePart) HasApiId() bool {
	if o != nil && !isNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *LookupTableNotePart) SetApiId(v string) {
	o.ApiId = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *LookupTableNotePart) GetColumns() []StructuredTableColumnInfo {
	if o == nil || isNil(o.Columns) {
		var ret []StructuredTableColumnInfo
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableNotePart) GetColumnsOk() ([]StructuredTableColumnInfo, bool) {
	if o == nil || isNil(o.Columns) {
    return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *LookupTableNotePart) HasColumns() bool {
	if o != nil && !isNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []StructuredTableColumnInfo and assigns it to the Columns field.
func (o *LookupTableNotePart) SetColumns(v []StructuredTableColumnInfo) {
	o.Columns = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LookupTableNotePart) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupTableNotePart) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LookupTableNotePart) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LookupTableNotePart) SetType(v string) {
	o.Type = &v
}

func (o LookupTableNotePart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiId) {
		toSerialize["apiId"] = o.ApiId
	}
	if !isNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLookupTableNotePart struct {
	value *LookupTableNotePart
	isSet bool
}

func (v NullableLookupTableNotePart) Get() *LookupTableNotePart {
	return v.value
}

func (v *NullableLookupTableNotePart) Set(val *LookupTableNotePart) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupTableNotePart) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupTableNotePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupTableNotePart(val *LookupTableNotePart) *NullableLookupTableNotePart {
	return &NullableLookupTableNotePart{value: val, isSet: true}
}

func (v NullableLookupTableNotePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupTableNotePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


