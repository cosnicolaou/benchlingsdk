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

// AaSequenceArchiveRecord struct for AaSequenceArchiveRecord
type AaSequenceArchiveRecord struct {
	Reason *string `json:"reason,omitempty"`
}

// NewAaSequenceArchiveRecord instantiates a new AaSequenceArchiveRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaSequenceArchiveRecord() *AaSequenceArchiveRecord {
	this := AaSequenceArchiveRecord{}
	return &this
}

// NewAaSequenceArchiveRecordWithDefaults instantiates a new AaSequenceArchiveRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaSequenceArchiveRecordWithDefaults() *AaSequenceArchiveRecord {
	this := AaSequenceArchiveRecord{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AaSequenceArchiveRecord) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaSequenceArchiveRecord) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AaSequenceArchiveRecord) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AaSequenceArchiveRecord) SetReason(v string) {
	o.Reason = &v
}

func (o AaSequenceArchiveRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableAaSequenceArchiveRecord struct {
	value *AaSequenceArchiveRecord
	isSet bool
}

func (v NullableAaSequenceArchiveRecord) Get() *AaSequenceArchiveRecord {
	return v.value
}

func (v *NullableAaSequenceArchiveRecord) Set(val *AaSequenceArchiveRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAaSequenceArchiveRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAaSequenceArchiveRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaSequenceArchiveRecord(val *AaSequenceArchiveRecord) *NullableAaSequenceArchiveRecord {
	return &NullableAaSequenceArchiveRecord{value: val, isSet: true}
}

func (v NullableAaSequenceArchiveRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaSequenceArchiveRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


