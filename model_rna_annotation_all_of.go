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

// RnaAnnotationAllOf struct for RnaAnnotationAllOf
type RnaAnnotationAllOf struct {
	End *int32 `json:"end,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Strand *int32 `json:"strand,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewRnaAnnotationAllOf instantiates a new RnaAnnotationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaAnnotationAllOf() *RnaAnnotationAllOf {
	this := RnaAnnotationAllOf{}
	return &this
}

// NewRnaAnnotationAllOfWithDefaults instantiates a new RnaAnnotationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaAnnotationAllOfWithDefaults() *RnaAnnotationAllOf {
	this := RnaAnnotationAllOf{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *RnaAnnotationAllOf) GetEnd() int32 {
	if o == nil || isNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaAnnotationAllOf) GetEndOk() (*int32, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *RnaAnnotationAllOf) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *RnaAnnotationAllOf) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *RnaAnnotationAllOf) GetStart() int32 {
	if o == nil || isNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaAnnotationAllOf) GetStartOk() (*int32, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *RnaAnnotationAllOf) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *RnaAnnotationAllOf) SetStart(v int32) {
	o.Start = &v
}

// GetStrand returns the Strand field value if set, zero value otherwise.
func (o *RnaAnnotationAllOf) GetStrand() int32 {
	if o == nil || isNil(o.Strand) {
		var ret int32
		return ret
	}
	return *o.Strand
}

// GetStrandOk returns a tuple with the Strand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaAnnotationAllOf) GetStrandOk() (*int32, bool) {
	if o == nil || isNil(o.Strand) {
    return nil, false
	}
	return o.Strand, true
}

// HasStrand returns a boolean if a field has been set.
func (o *RnaAnnotationAllOf) HasStrand() bool {
	if o != nil && !isNil(o.Strand) {
		return true
	}

	return false
}

// SetStrand gets a reference to the given int32 and assigns it to the Strand field.
func (o *RnaAnnotationAllOf) SetStrand(v int32) {
	o.Strand = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RnaAnnotationAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaAnnotationAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RnaAnnotationAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RnaAnnotationAllOf) SetType(v string) {
	o.Type = &v
}

func (o RnaAnnotationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableRnaAnnotationAllOf struct {
	value *RnaAnnotationAllOf
	isSet bool
}

func (v NullableRnaAnnotationAllOf) Get() *RnaAnnotationAllOf {
	return v.value
}

func (v *NullableRnaAnnotationAllOf) Set(val *RnaAnnotationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaAnnotationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaAnnotationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaAnnotationAllOf(val *RnaAnnotationAllOf) *NullableRnaAnnotationAllOf {
	return &NullableRnaAnnotationAllOf{value: val, isSet: true}
}

func (v NullableRnaAnnotationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaAnnotationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


