/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
	"fmt"
)

// SampleGroupStatus Status of a sample group
type SampleGroupStatus string

// List of SampleGroupStatus
const (
	IN_PROGRESS SampleGroupStatus = "IN_PROGRESS"
	COMPLETED SampleGroupStatus = "COMPLETED"
	FAILED SampleGroupStatus = "FAILED"
)

// All allowed values of SampleGroupStatus enum
var AllowedSampleGroupStatusEnumValues = []SampleGroupStatus{
	"IN_PROGRESS",
	"COMPLETED",
	"FAILED",
}

func (v *SampleGroupStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SampleGroupStatus(value)
	for _, existing := range AllowedSampleGroupStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SampleGroupStatus", value)
}

// NewSampleGroupStatusFromValue returns a pointer to a valid SampleGroupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSampleGroupStatusFromValue(v string) (*SampleGroupStatus, error) {
	ev := SampleGroupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SampleGroupStatus: valid values are %v", v, AllowedSampleGroupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SampleGroupStatus) IsValid() bool {
	for _, existing := range AllowedSampleGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SampleGroupStatus value
func (v SampleGroupStatus) Ptr() *SampleGroupStatus {
	return &v
}

type NullableSampleGroupStatus struct {
	value *SampleGroupStatus
	isSet bool
}

func (v NullableSampleGroupStatus) Get() *SampleGroupStatus {
	return v.value
}

func (v *NullableSampleGroupStatus) Set(val *SampleGroupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleGroupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleGroupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleGroupStatus(val *SampleGroupStatus) *NullableSampleGroupStatus {
	return &NullableSampleGroupStatus{value: val, isSet: true}
}

func (v NullableSampleGroupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleGroupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

