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

// SimpleNotePart Simple note parts include the following types: - 'text': plain text - 'code': preformatted code block - 'list_bullet': one \"line\" of a bulleted list - 'list_number': one \"line\" of a numbered list 
type SimpleNotePart struct {
	// All notes have an indentation level - the default is 0 for no indent. For lists, indentation gives notes hierarchy - a bulleted list with children is modeled as one note part with indentation 1 followed by note parts with indentation 2, for example.
	Indentation *int32 `json:"indentation,omitempty"`
	Type *string `json:"type,omitempty"`
	// Array of links referenced in text via an @-mention, hyperlink, or the drag-n-dropped preview attached to the note. 
	Links []EntryLink `json:"links,omitempty"`
	// The textual contents of the note.
	Text *string `json:"text,omitempty"`
}

// NewSimpleNotePart instantiates a new SimpleNotePart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleNotePart() *SimpleNotePart {
	this := SimpleNotePart{}
	var indentation int32 = 0
	this.Indentation = &indentation
	return &this
}

// NewSimpleNotePartWithDefaults instantiates a new SimpleNotePart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleNotePartWithDefaults() *SimpleNotePart {
	this := SimpleNotePart{}
	var indentation int32 = 0
	this.Indentation = &indentation
	return &this
}

// GetIndentation returns the Indentation field value if set, zero value otherwise.
func (o *SimpleNotePart) GetIndentation() int32 {
	if o == nil || isNil(o.Indentation) {
		var ret int32
		return ret
	}
	return *o.Indentation
}

// GetIndentationOk returns a tuple with the Indentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePart) GetIndentationOk() (*int32, bool) {
	if o == nil || isNil(o.Indentation) {
    return nil, false
	}
	return o.Indentation, true
}

// HasIndentation returns a boolean if a field has been set.
func (o *SimpleNotePart) HasIndentation() bool {
	if o != nil && !isNil(o.Indentation) {
		return true
	}

	return false
}

// SetIndentation gets a reference to the given int32 and assigns it to the Indentation field.
func (o *SimpleNotePart) SetIndentation(v int32) {
	o.Indentation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimpleNotePart) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePart) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimpleNotePart) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimpleNotePart) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SimpleNotePart) GetLinks() []EntryLink {
	if o == nil || isNil(o.Links) {
		var ret []EntryLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePart) GetLinksOk() ([]EntryLink, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SimpleNotePart) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []EntryLink and assigns it to the Links field.
func (o *SimpleNotePart) SetLinks(v []EntryLink) {
	o.Links = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SimpleNotePart) GetText() string {
	if o == nil || isNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleNotePart) GetTextOk() (*string, bool) {
	if o == nil || isNil(o.Text) {
    return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SimpleNotePart) HasText() bool {
	if o != nil && !isNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SimpleNotePart) SetText(v string) {
	o.Text = &v
}

func (o SimpleNotePart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Indentation) {
		toSerialize["indentation"] = o.Indentation
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !isNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleNotePart struct {
	value *SimpleNotePart
	isSet bool
}

func (v NullableSimpleNotePart) Get() *SimpleNotePart {
	return v.value
}

func (v *NullableSimpleNotePart) Set(val *SimpleNotePart) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleNotePart) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleNotePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleNotePart(val *SimpleNotePart) *NullableSimpleNotePart {
	return &NullableSimpleNotePart{value: val, isSet: true}
}

func (v NullableSimpleNotePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleNotePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


