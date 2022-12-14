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

// ProjectOwner struct for ProjectOwner
type ProjectOwner struct {
	Organization *Organization
	UserSummary *UserSummary
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProjectOwner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Organization
	err = json.Unmarshal(data, &dst.Organization);
	if err == nil {
		jsonOrganization, _ := json.Marshal(dst.Organization)
		if string(jsonOrganization) == "{}" { // empty struct
			dst.Organization = nil
		} else {
			return nil // data stored in dst.Organization, return on the first match
		}
	} else {
		dst.Organization = nil
	}

	// try to unmarshal JSON data into UserSummary
	err = json.Unmarshal(data, &dst.UserSummary);
	if err == nil {
		jsonUserSummary, _ := json.Marshal(dst.UserSummary)
		if string(jsonUserSummary) == "{}" { // empty struct
			dst.UserSummary = nil
		} else {
			return nil // data stored in dst.UserSummary, return on the first match
		}
	} else {
		dst.UserSummary = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProjectOwner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProjectOwner) MarshalJSON() ([]byte, error) {
	if src.Organization != nil {
		return json.Marshal(&src.Organization)
	}

	if src.UserSummary != nil {
		return json.Marshal(&src.UserSummary)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProjectOwner struct {
	value *ProjectOwner
	isSet bool
}

func (v NullableProjectOwner) Get() *ProjectOwner {
	return v.value
}

func (v *NullableProjectOwner) Set(val *ProjectOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectOwner(val *ProjectOwner) *NullableProjectOwner {
	return &NullableProjectOwner{value: val, isSet: true}
}

func (v NullableProjectOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


