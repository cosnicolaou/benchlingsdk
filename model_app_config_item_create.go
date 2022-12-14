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

// AppConfigItemCreate - struct for AppConfigItemCreate
type AppConfigItemCreate struct {
	AppConfigItemBooleanCreate *AppConfigItemBooleanCreate
	AppConfigItemDateCreate *AppConfigItemDateCreate
	AppConfigItemDatetimeCreate *AppConfigItemDatetimeCreate
	AppConfigItemFloatCreate *AppConfigItemFloatCreate
	AppConfigItemGenericCreate *AppConfigItemGenericCreate
	AppConfigItemIntegerCreate *AppConfigItemIntegerCreate
	AppConfigItemJsonCreate *AppConfigItemJsonCreate
}

// AppConfigItemBooleanCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemBooleanCreate wrapped in AppConfigItemCreate
func AppConfigItemBooleanCreateAsAppConfigItemCreate(v *AppConfigItemBooleanCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemBooleanCreate: v,
	}
}

// AppConfigItemDateCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemDateCreate wrapped in AppConfigItemCreate
func AppConfigItemDateCreateAsAppConfigItemCreate(v *AppConfigItemDateCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemDateCreate: v,
	}
}

// AppConfigItemDatetimeCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemDatetimeCreate wrapped in AppConfigItemCreate
func AppConfigItemDatetimeCreateAsAppConfigItemCreate(v *AppConfigItemDatetimeCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemDatetimeCreate: v,
	}
}

// AppConfigItemFloatCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemFloatCreate wrapped in AppConfigItemCreate
func AppConfigItemFloatCreateAsAppConfigItemCreate(v *AppConfigItemFloatCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemFloatCreate: v,
	}
}

// AppConfigItemGenericCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemGenericCreate wrapped in AppConfigItemCreate
func AppConfigItemGenericCreateAsAppConfigItemCreate(v *AppConfigItemGenericCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemGenericCreate: v,
	}
}

// AppConfigItemIntegerCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemIntegerCreate wrapped in AppConfigItemCreate
func AppConfigItemIntegerCreateAsAppConfigItemCreate(v *AppConfigItemIntegerCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemIntegerCreate: v,
	}
}

// AppConfigItemJsonCreateAsAppConfigItemCreate is a convenience function that returns AppConfigItemJsonCreate wrapped in AppConfigItemCreate
func AppConfigItemJsonCreateAsAppConfigItemCreate(v *AppConfigItemJsonCreate) AppConfigItemCreate {
	return AppConfigItemCreate{
		AppConfigItemJsonCreate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppConfigItemCreate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppConfigItemBooleanCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemBooleanCreate)
	if err == nil {
		jsonAppConfigItemBooleanCreate, _ := json.Marshal(dst.AppConfigItemBooleanCreate)
		if string(jsonAppConfigItemBooleanCreate) == "{}" { // empty struct
			dst.AppConfigItemBooleanCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemBooleanCreate = nil
	}

	// try to unmarshal data into AppConfigItemDateCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemDateCreate)
	if err == nil {
		jsonAppConfigItemDateCreate, _ := json.Marshal(dst.AppConfigItemDateCreate)
		if string(jsonAppConfigItemDateCreate) == "{}" { // empty struct
			dst.AppConfigItemDateCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemDateCreate = nil
	}

	// try to unmarshal data into AppConfigItemDatetimeCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemDatetimeCreate)
	if err == nil {
		jsonAppConfigItemDatetimeCreate, _ := json.Marshal(dst.AppConfigItemDatetimeCreate)
		if string(jsonAppConfigItemDatetimeCreate) == "{}" { // empty struct
			dst.AppConfigItemDatetimeCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemDatetimeCreate = nil
	}

	// try to unmarshal data into AppConfigItemFloatCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemFloatCreate)
	if err == nil {
		jsonAppConfigItemFloatCreate, _ := json.Marshal(dst.AppConfigItemFloatCreate)
		if string(jsonAppConfigItemFloatCreate) == "{}" { // empty struct
			dst.AppConfigItemFloatCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemFloatCreate = nil
	}

	// try to unmarshal data into AppConfigItemGenericCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemGenericCreate)
	if err == nil {
		jsonAppConfigItemGenericCreate, _ := json.Marshal(dst.AppConfigItemGenericCreate)
		if string(jsonAppConfigItemGenericCreate) == "{}" { // empty struct
			dst.AppConfigItemGenericCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemGenericCreate = nil
	}

	// try to unmarshal data into AppConfigItemIntegerCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemIntegerCreate)
	if err == nil {
		jsonAppConfigItemIntegerCreate, _ := json.Marshal(dst.AppConfigItemIntegerCreate)
		if string(jsonAppConfigItemIntegerCreate) == "{}" { // empty struct
			dst.AppConfigItemIntegerCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemIntegerCreate = nil
	}

	// try to unmarshal data into AppConfigItemJsonCreate
	err = newStrictDecoder(data).Decode(&dst.AppConfigItemJsonCreate)
	if err == nil {
		jsonAppConfigItemJsonCreate, _ := json.Marshal(dst.AppConfigItemJsonCreate)
		if string(jsonAppConfigItemJsonCreate) == "{}" { // empty struct
			dst.AppConfigItemJsonCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppConfigItemJsonCreate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppConfigItemBooleanCreate = nil
		dst.AppConfigItemDateCreate = nil
		dst.AppConfigItemDatetimeCreate = nil
		dst.AppConfigItemFloatCreate = nil
		dst.AppConfigItemGenericCreate = nil
		dst.AppConfigItemIntegerCreate = nil
		dst.AppConfigItemJsonCreate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppConfigItemCreate)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppConfigItemCreate)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppConfigItemCreate) MarshalJSON() ([]byte, error) {
	if src.AppConfigItemBooleanCreate != nil {
		return json.Marshal(&src.AppConfigItemBooleanCreate)
	}

	if src.AppConfigItemDateCreate != nil {
		return json.Marshal(&src.AppConfigItemDateCreate)
	}

	if src.AppConfigItemDatetimeCreate != nil {
		return json.Marshal(&src.AppConfigItemDatetimeCreate)
	}

	if src.AppConfigItemFloatCreate != nil {
		return json.Marshal(&src.AppConfigItemFloatCreate)
	}

	if src.AppConfigItemGenericCreate != nil {
		return json.Marshal(&src.AppConfigItemGenericCreate)
	}

	if src.AppConfigItemIntegerCreate != nil {
		return json.Marshal(&src.AppConfigItemIntegerCreate)
	}

	if src.AppConfigItemJsonCreate != nil {
		return json.Marshal(&src.AppConfigItemJsonCreate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppConfigItemCreate) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppConfigItemBooleanCreate != nil {
		return obj.AppConfigItemBooleanCreate
	}

	if obj.AppConfigItemDateCreate != nil {
		return obj.AppConfigItemDateCreate
	}

	if obj.AppConfigItemDatetimeCreate != nil {
		return obj.AppConfigItemDatetimeCreate
	}

	if obj.AppConfigItemFloatCreate != nil {
		return obj.AppConfigItemFloatCreate
	}

	if obj.AppConfigItemGenericCreate != nil {
		return obj.AppConfigItemGenericCreate
	}

	if obj.AppConfigItemIntegerCreate != nil {
		return obj.AppConfigItemIntegerCreate
	}

	if obj.AppConfigItemJsonCreate != nil {
		return obj.AppConfigItemJsonCreate
	}

	// all schemas are nil
	return nil
}

type NullableAppConfigItemCreate struct {
	value *AppConfigItemCreate
	isSet bool
}

func (v NullableAppConfigItemCreate) Get() *AppConfigItemCreate {
	return v.value
}

func (v *NullableAppConfigItemCreate) Set(val *AppConfigItemCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigItemCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigItemCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigItemCreate(val *AppConfigItemCreate) *NullableAppConfigItemCreate {
	return &NullableAppConfigItemCreate{value: val, isSet: true}
}

func (v NullableAppConfigItemCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigItemCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


