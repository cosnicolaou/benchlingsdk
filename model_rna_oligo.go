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

// RnaOligo struct for RnaOligo
type RnaOligo struct {
	Oligo
	ApiURL *string `json:"apiURL,omitempty"`
	Bases *string `json:"bases,omitempty"`
	NucleotideType *string `json:"nucleotideType,omitempty"`
}

// NewRnaOligo instantiates a new RnaOligo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRnaOligo() *RnaOligo {
	this := RnaOligo{}
	return &this
}

// NewRnaOligoWithDefaults instantiates a new RnaOligo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRnaOligoWithDefaults() *RnaOligo {
	this := RnaOligo{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *RnaOligo) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligo) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *RnaOligo) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *RnaOligo) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetBases returns the Bases field value if set, zero value otherwise.
func (o *RnaOligo) GetBases() string {
	if o == nil || isNil(o.Bases) {
		var ret string
		return ret
	}
	return *o.Bases
}

// GetBasesOk returns a tuple with the Bases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligo) GetBasesOk() (*string, bool) {
	if o == nil || isNil(o.Bases) {
    return nil, false
	}
	return o.Bases, true
}

// HasBases returns a boolean if a field has been set.
func (o *RnaOligo) HasBases() bool {
	if o != nil && !isNil(o.Bases) {
		return true
	}

	return false
}

// SetBases gets a reference to the given string and assigns it to the Bases field.
func (o *RnaOligo) SetBases(v string) {
	o.Bases = &v
}

// GetNucleotideType returns the NucleotideType field value if set, zero value otherwise.
func (o *RnaOligo) GetNucleotideType() string {
	if o == nil || isNil(o.NucleotideType) {
		var ret string
		return ret
	}
	return *o.NucleotideType
}

// GetNucleotideTypeOk returns a tuple with the NucleotideType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RnaOligo) GetNucleotideTypeOk() (*string, bool) {
	if o == nil || isNil(o.NucleotideType) {
    return nil, false
	}
	return o.NucleotideType, true
}

// HasNucleotideType returns a boolean if a field has been set.
func (o *RnaOligo) HasNucleotideType() bool {
	if o != nil && !isNil(o.NucleotideType) {
		return true
	}

	return false
}

// SetNucleotideType gets a reference to the given string and assigns it to the NucleotideType field.
func (o *RnaOligo) SetNucleotideType(v string) {
	o.NucleotideType = &v
}

func (o RnaOligo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOligo, errOligo := json.Marshal(o.Oligo)
	if errOligo != nil {
		return []byte{}, errOligo
	}
	errOligo = json.Unmarshal([]byte(serializedOligo), &toSerialize)
	if errOligo != nil {
		return []byte{}, errOligo
	}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
	if !isNil(o.Bases) {
		toSerialize["bases"] = o.Bases
	}
	if !isNil(o.NucleotideType) {
		toSerialize["nucleotideType"] = o.NucleotideType
	}
	return json.Marshal(toSerialize)
}

type NullableRnaOligo struct {
	value *RnaOligo
	isSet bool
}

func (v NullableRnaOligo) Get() *RnaOligo {
	return v.value
}

func (v *NullableRnaOligo) Set(val *RnaOligo) {
	v.value = val
	v.isSet = true
}

func (v NullableRnaOligo) IsSet() bool {
	return v.isSet
}

func (v *NullableRnaOligo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRnaOligo(val *RnaOligo) *NullableRnaOligo {
	return &NullableRnaOligo{value: val, isSet: true}
}

func (v NullableRnaOligo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRnaOligo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


