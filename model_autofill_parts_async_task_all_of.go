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

// AutofillPartsAsyncTaskAllOf struct for AutofillPartsAsyncTaskAllOf
type AutofillPartsAsyncTaskAllOf struct {
	Response interface{} `json:"response,omitempty"`
}

// NewAutofillPartsAsyncTaskAllOf instantiates a new AutofillPartsAsyncTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutofillPartsAsyncTaskAllOf() *AutofillPartsAsyncTaskAllOf {
	this := AutofillPartsAsyncTaskAllOf{}
	return &this
}

// NewAutofillPartsAsyncTaskAllOfWithDefaults instantiates a new AutofillPartsAsyncTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutofillPartsAsyncTaskAllOfWithDefaults() *AutofillPartsAsyncTaskAllOf {
	this := AutofillPartsAsyncTaskAllOf{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutofillPartsAsyncTaskAllOf) GetResponse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutofillPartsAsyncTaskAllOf) GetResponseOk() (*interface{}, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return &o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AutofillPartsAsyncTaskAllOf) HasResponse() bool {
	if o != nil && isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given interface{} and assigns it to the Response field.
func (o *AutofillPartsAsyncTaskAllOf) SetResponse(v interface{}) {
	o.Response = v
}

func (o AutofillPartsAsyncTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableAutofillPartsAsyncTaskAllOf struct {
	value *AutofillPartsAsyncTaskAllOf
	isSet bool
}

func (v NullableAutofillPartsAsyncTaskAllOf) Get() *AutofillPartsAsyncTaskAllOf {
	return v.value
}

func (v *NullableAutofillPartsAsyncTaskAllOf) Set(val *AutofillPartsAsyncTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAutofillPartsAsyncTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAutofillPartsAsyncTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutofillPartsAsyncTaskAllOf(val *AutofillPartsAsyncTaskAllOf) *NullableAutofillPartsAsyncTaskAllOf {
	return &NullableAutofillPartsAsyncTaskAllOf{value: val, isSet: true}
}

func (v NullableAutofillPartsAsyncTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutofillPartsAsyncTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


