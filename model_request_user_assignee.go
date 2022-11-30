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

// RequestUserAssignee struct for RequestUserAssignee
type RequestUserAssignee struct {
	User *UserSummary `json:"user,omitempty"`
}

// NewRequestUserAssignee instantiates a new RequestUserAssignee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestUserAssignee() *RequestUserAssignee {
	this := RequestUserAssignee{}
	return &this
}

// NewRequestUserAssigneeWithDefaults instantiates a new RequestUserAssignee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestUserAssigneeWithDefaults() *RequestUserAssignee {
	this := RequestUserAssignee{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RequestUserAssignee) GetUser() UserSummary {
	if o == nil || isNil(o.User) {
		var ret UserSummary
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUserAssignee) GetUserOk() (*UserSummary, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RequestUserAssignee) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserSummary and assigns it to the User field.
func (o *RequestUserAssignee) SetUser(v UserSummary) {
	o.User = &v
}

func (o RequestUserAssignee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableRequestUserAssignee struct {
	value *RequestUserAssignee
	isSet bool
}

func (v NullableRequestUserAssignee) Get() *RequestUserAssignee {
	return v.value
}

func (v *NullableRequestUserAssignee) Set(val *RequestUserAssignee) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestUserAssignee) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestUserAssignee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestUserAssignee(val *RequestUserAssignee) *NullableRequestUserAssignee {
	return &NullableRequestUserAssignee{value: val, isSet: true}
}

func (v NullableRequestUserAssignee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestUserAssignee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


