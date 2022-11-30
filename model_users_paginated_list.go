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

// UsersPaginatedList struct for UsersPaginatedList
type UsersPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	Users []User `json:"users,omitempty"`
}

// NewUsersPaginatedList instantiates a new UsersPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersPaginatedList() *UsersPaginatedList {
	this := UsersPaginatedList{}
	return &this
}

// NewUsersPaginatedListWithDefaults instantiates a new UsersPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersPaginatedListWithDefaults() *UsersPaginatedList {
	this := UsersPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *UsersPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *UsersPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *UsersPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UsersPaginatedList) GetUsers() []User {
	if o == nil || isNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPaginatedList) GetUsersOk() ([]User, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UsersPaginatedList) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *UsersPaginatedList) SetUsers(v []User) {
	o.Users = v
}

func (o UsersPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableUsersPaginatedList struct {
	value *UsersPaginatedList
	isSet bool
}

func (v NullableUsersPaginatedList) Get() *UsersPaginatedList {
	return v.value
}

func (v *NullableUsersPaginatedList) Set(val *UsersPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersPaginatedList(val *UsersPaginatedList) *NullableUsersPaginatedList {
	return &NullableUsersPaginatedList{value: val, isSet: true}
}

func (v NullableUsersPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


