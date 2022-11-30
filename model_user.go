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

// User struct for User
type User struct {
	Handle *string `json:"handle,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	IsSuspended *bool `json:"isSuspended,omitempty"`
	PasswordLastChangedAt NullableString `json:"passwordLastChangedAt,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *User) GetHandle() string {
	if o == nil || isNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetHandleOk() (*string, bool) {
	if o == nil || isNil(o.Handle) {
    return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *User) HasHandle() bool {
	if o != nil && !isNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *User) SetHandle(v string) {
	o.Handle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetIsSuspended returns the IsSuspended field value if set, zero value otherwise.
func (o *User) GetIsSuspended() bool {
	if o == nil || isNil(o.IsSuspended) {
		var ret bool
		return ret
	}
	return *o.IsSuspended
}

// GetIsSuspendedOk returns a tuple with the IsSuspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsSuspendedOk() (*bool, bool) {
	if o == nil || isNil(o.IsSuspended) {
    return nil, false
	}
	return o.IsSuspended, true
}

// HasIsSuspended returns a boolean if a field has been set.
func (o *User) HasIsSuspended() bool {
	if o != nil && !isNil(o.IsSuspended) {
		return true
	}

	return false
}

// SetIsSuspended gets a reference to the given bool and assigns it to the IsSuspended field.
func (o *User) SetIsSuspended(v bool) {
	o.IsSuspended = &v
}

// GetPasswordLastChangedAt returns the PasswordLastChangedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPasswordLastChangedAt() string {
	if o == nil || isNil(o.PasswordLastChangedAt.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordLastChangedAt.Get()
}

// GetPasswordLastChangedAtOk returns a tuple with the PasswordLastChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPasswordLastChangedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PasswordLastChangedAt.Get(), o.PasswordLastChangedAt.IsSet()
}

// HasPasswordLastChangedAt returns a boolean if a field has been set.
func (o *User) HasPasswordLastChangedAt() bool {
	if o != nil && o.PasswordLastChangedAt.IsSet() {
		return true
	}

	return false
}

// SetPasswordLastChangedAt gets a reference to the given NullableString and assigns it to the PasswordLastChangedAt field.
func (o *User) SetPasswordLastChangedAt(v string) {
	o.PasswordLastChangedAt.Set(&v)
}
// SetPasswordLastChangedAtNil sets the value for PasswordLastChangedAt to be an explicit nil
func (o *User) SetPasswordLastChangedAtNil() {
	o.PasswordLastChangedAt.Set(nil)
}

// UnsetPasswordLastChangedAt ensures that no value is present for PasswordLastChangedAt, not even an explicit nil
func (o *User) UnsetPasswordLastChangedAt() {
	o.PasswordLastChangedAt.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.IsSuspended) {
		toSerialize["isSuspended"] = o.IsSuspended
	}
	if o.PasswordLastChangedAt.IsSet() {
		toSerialize["passwordLastChangedAt"] = o.PasswordLastChangedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


