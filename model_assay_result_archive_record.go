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

// AssayResultArchiveRecord ArchiveRecord Resource if the result is archived. This is null if the result is not archived. 
type AssayResultArchiveRecord struct {
	Reason *string `json:"reason,omitempty"`
}

// NewAssayResultArchiveRecord instantiates a new AssayResultArchiveRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssayResultArchiveRecord() *AssayResultArchiveRecord {
	this := AssayResultArchiveRecord{}
	return &this
}

// NewAssayResultArchiveRecordWithDefaults instantiates a new AssayResultArchiveRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssayResultArchiveRecordWithDefaults() *AssayResultArchiveRecord {
	this := AssayResultArchiveRecord{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AssayResultArchiveRecord) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssayResultArchiveRecord) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AssayResultArchiveRecord) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AssayResultArchiveRecord) SetReason(v string) {
	o.Reason = &v
}

func (o AssayResultArchiveRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableAssayResultArchiveRecord struct {
	value *AssayResultArchiveRecord
	isSet bool
}

func (v NullableAssayResultArchiveRecord) Get() *AssayResultArchiveRecord {
	return v.value
}

func (v *NullableAssayResultArchiveRecord) Set(val *AssayResultArchiveRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayResultArchiveRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayResultArchiveRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayResultArchiveRecord(val *AssayResultArchiveRecord) *NullableAssayResultArchiveRecord {
	return &NullableAssayResultArchiveRecord{value: val, isSet: true}
}

func (v NullableAssayResultArchiveRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayResultArchiveRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


