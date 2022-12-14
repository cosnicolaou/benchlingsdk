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

// SimpleNotePartAllOf struct for SimpleNotePartAllOf
type SimpleNotePartAllOf struct {
	// Array of links referenced in text via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note. 
	Links []EntryLink `json:"links,omitempty"`
	// The textual contents of the note.
	Text *string `json:"text,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewSimpleNotePartAllOf instantiates a new SimpleNotePartAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleNotePartAllOf() *SimpleNotePartAllOf {
	this := SimpleNotePartAllOf{}
	return &this
}

// NewSimpleNotePartAllOfWithDefaults instantiates a new SimpleNotePartAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleNotePartAllOfWithDefaults() *SimpleNotePartAllOf {
	this := SimpleNotePartAllOf{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SimpleNotePartAllOf) GetLinks() []EntryLink {
	if o == nil || isNil(o.Links) {
		var ret []EntryLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePartAllOf) GetLinksOk() ([]EntryLink, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SimpleNotePartAllOf) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []EntryLink and assigns it to the Links field.
func (o *SimpleNotePartAllOf) SetLinks(v []EntryLink) {
	o.Links = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SimpleNotePartAllOf) GetText() string {
	if o == nil || isNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePartAllOf) GetTextOk() (*string, bool) {
	if o == nil || isNil(o.Text) {
    return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SimpleNotePartAllOf) HasText() bool {
	if o != nil && !isNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SimpleNotePartAllOf) SetText(v string) {
	o.Text = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimpleNotePartAllOf) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePartAllOf) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimpleNotePartAllOf) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimpleNotePartAllOf) SetType(v string) {
	o.Type = &v
}

func (o SimpleNotePartAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleNotePartAllOf struct {
	value *SimpleNotePartAllOf
	isSet bool
}

func (v NullableSimpleNotePartAllOf) Get() *SimpleNotePartAllOf {
	return v.value
}

func (v *NullableSimpleNotePartAllOf) Set(val *SimpleNotePartAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleNotePartAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleNotePartAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleNotePartAllOf(val *SimpleNotePartAllOf) *NullableSimpleNotePartAllOf {
	return &NullableSimpleNotePartAllOf{value: val, isSet: true}
}

func (v NullableSimpleNotePartAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleNotePartAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


