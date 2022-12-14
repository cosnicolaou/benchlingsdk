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

// RequestSampleGroupSamplesValue - struct for RequestSampleGroupSamplesValue
type RequestSampleGroupSamplesValue struct {
	RequestSampleWithBatch *RequestSampleWithBatch
	RequestSampleWithEntity *RequestSampleWithEntity
}

// RequestSampleWithBatchAsRequestSampleGroupSamplesValue is a convenience function that returns RequestSampleWithBatch wrapped in RequestSampleGroupSamplesValue
func RequestSampleWithBatchAsRequestSampleGroupSamplesValue(v *RequestSampleWithBatch) RequestSampleGroupSamplesValue {
	return RequestSampleGroupSamplesValue{
		RequestSampleWithBatch: v,
	}
}

// RequestSampleWithEntityAsRequestSampleGroupSamplesValue is a convenience function that returns RequestSampleWithEntity wrapped in RequestSampleGroupSamplesValue
func RequestSampleWithEntityAsRequestSampleGroupSamplesValue(v *RequestSampleWithEntity) RequestSampleGroupSamplesValue {
	return RequestSampleGroupSamplesValue{
		RequestSampleWithEntity: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestSampleGroupSamplesValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RequestSampleWithBatch
	err = newStrictDecoder(data).Decode(&dst.RequestSampleWithBatch)
	if err == nil {
		jsonRequestSampleWithBatch, _ := json.Marshal(dst.RequestSampleWithBatch)
		if string(jsonRequestSampleWithBatch) == "{}" { // empty struct
			dst.RequestSampleWithBatch = nil
		} else {
			match++
		}
	} else {
		dst.RequestSampleWithBatch = nil
	}

	// try to unmarshal data into RequestSampleWithEntity
	err = newStrictDecoder(data).Decode(&dst.RequestSampleWithEntity)
	if err == nil {
		jsonRequestSampleWithEntity, _ := json.Marshal(dst.RequestSampleWithEntity)
		if string(jsonRequestSampleWithEntity) == "{}" { // empty struct
			dst.RequestSampleWithEntity = nil
		} else {
			match++
		}
	} else {
		dst.RequestSampleWithEntity = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RequestSampleWithBatch = nil
		dst.RequestSampleWithEntity = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RequestSampleGroupSamplesValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestSampleGroupSamplesValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestSampleGroupSamplesValue) MarshalJSON() ([]byte, error) {
	if src.RequestSampleWithBatch != nil {
		return json.Marshal(&src.RequestSampleWithBatch)
	}

	if src.RequestSampleWithEntity != nil {
		return json.Marshal(&src.RequestSampleWithEntity)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestSampleGroupSamplesValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RequestSampleWithBatch != nil {
		return obj.RequestSampleWithBatch
	}

	if obj.RequestSampleWithEntity != nil {
		return obj.RequestSampleWithEntity
	}

	// all schemas are nil
	return nil
}

type NullableRequestSampleGroupSamplesValue struct {
	value *RequestSampleGroupSamplesValue
	isSet bool
}

func (v NullableRequestSampleGroupSamplesValue) Get() *RequestSampleGroupSamplesValue {
	return v.value
}

func (v *NullableRequestSampleGroupSamplesValue) Set(val *RequestSampleGroupSamplesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSampleGroupSamplesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSampleGroupSamplesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSampleGroupSamplesValue(val *RequestSampleGroupSamplesValue) *NullableRequestSampleGroupSamplesValue {
	return &NullableRequestSampleGroupSamplesValue{value: val, isSet: true}
}

func (v NullableRequestSampleGroupSamplesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSampleGroupSamplesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


