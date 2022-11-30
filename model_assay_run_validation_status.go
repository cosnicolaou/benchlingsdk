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

// AssayRunValidationStatus Must be either VALID or INVALID
type AssayRunValidationStatus string

// List of AssayRunValidationStatus
const (
	VALID AssayRunValidationStatus = "VALID"
	INVALID AssayRunValidationStatus = "INVALID"
)

// All allowed values of AssayRunValidationStatus enum
var AllowedAssayRunValidationStatusEnumValues = []AssayRunValidationStatus{
	"VALID",
	"INVALID",
}

func (v *AssayRunValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssayRunValidationStatus(value)
	for _, existing := range AllowedAssayRunValidationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssayRunValidationStatus", value)
}

// NewAssayRunValidationStatusFromValue returns a pointer to a valid AssayRunValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssayRunValidationStatusFromValue(v string) (*AssayRunValidationStatus, error) {
	ev := AssayRunValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssayRunValidationStatus: valid values are %v", v, AllowedAssayRunValidationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssayRunValidationStatus) IsValid() bool {
	for _, existing := range AllowedAssayRunValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssayRunValidationStatus value
func (v AssayRunValidationStatus) Ptr() *AssayRunValidationStatus {
	return &v
}

type NullableAssayRunValidationStatus struct {
	value *AssayRunValidationStatus
	isSet bool
}

func (v NullableAssayRunValidationStatus) Get() *AssayRunValidationStatus {
	return v.value
}

func (v *NullableAssayRunValidationStatus) Set(val *AssayRunValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAssayRunValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAssayRunValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssayRunValidationStatus(val *AssayRunValidationStatus) *NullableAssayRunValidationStatus {
	return &NullableAssayRunValidationStatus{value: val, isSet: true}
}

func (v NullableAssayRunValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssayRunValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

