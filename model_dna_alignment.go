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

// DnaAlignment struct for DnaAlignment
type DnaAlignment struct {
	// The canonical url of the DNA Alignment in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	// DateTime the DNA Alignment was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id *string `json:"id,omitempty"`
	// DateTime the DNA Alignment was last modified
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	Name *string `json:"name,omitempty"`
	// The ID of the template or consensus DNA Sequence associated with the DNA Alignment
	ReferenceSequenceId *string `json:"referenceSequenceId,omitempty"`
	// The Benchling web UI url to view the DNA Alignment
	WebURL *string `json:"webURL,omitempty"`
	AlignedSequences []AlignedSequence `json:"alignedSequences,omitempty"`
}

// NewDnaAlignment instantiates a new DnaAlignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaAlignment() *DnaAlignment {
	this := DnaAlignment{}
	return &this
}

// NewDnaAlignmentWithDefaults instantiates a new DnaAlignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaAlignmentWithDefaults() *DnaAlignment {
	this := DnaAlignment{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *DnaAlignment) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *DnaAlignment) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *DnaAlignment) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DnaAlignment) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DnaAlignment) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DnaAlignment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DnaAlignment) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DnaAlignment) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DnaAlignment) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *DnaAlignment) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *DnaAlignment) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *DnaAlignment) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnaAlignment) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnaAlignment) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnaAlignment) SetName(v string) {
	o.Name = &v
}

// GetReferenceSequenceId returns the ReferenceSequenceId field value if set, zero value otherwise.
func (o *DnaAlignment) GetReferenceSequenceId() string {
	if o == nil || isNil(o.ReferenceSequenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceSequenceId
}

// GetReferenceSequenceIdOk returns a tuple with the ReferenceSequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetReferenceSequenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceSequenceId) {
    return nil, false
	}
	return o.ReferenceSequenceId, true
}

// HasReferenceSequenceId returns a boolean if a field has been set.
func (o *DnaAlignment) HasReferenceSequenceId() bool {
	if o != nil && !isNil(o.ReferenceSequenceId) {
		return true
	}

	return false
}

// SetReferenceSequenceId gets a reference to the given string and assigns it to the ReferenceSequenceId field.
func (o *DnaAlignment) SetReferenceSequenceId(v string) {
	o.ReferenceSequenceId = &v
}

// GetWebURL returns the WebURL field value if set, zero value otherwise.
func (o *DnaAlignment) GetWebURL() string {
	if o == nil || isNil(o.WebURL) {
		var ret string
		return ret
	}
	return *o.WebURL
}

// GetWebURLOk returns a tuple with the WebURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetWebURLOk() (*string, bool) {
	if o == nil || isNil(o.WebURL) {
    return nil, false
	}
	return o.WebURL, true
}

// HasWebURL returns a boolean if a field has been set.
func (o *DnaAlignment) HasWebURL() bool {
	if o != nil && !isNil(o.WebURL) {
		return true
	}

	return false
}

// SetWebURL gets a reference to the given string and assigns it to the WebURL field.
func (o *DnaAlignment) SetWebURL(v string) {
	o.WebURL = &v
}

// GetAlignedSequences returns the AlignedSequences field value if set, zero value otherwise.
func (o *DnaAlignment) GetAlignedSequences() []AlignedSequence {
	if o == nil || isNil(o.AlignedSequences) {
		var ret []AlignedSequence
		return ret
	}
	return o.AlignedSequences
}

// GetAlignedSequencesOk returns a tuple with the AlignedSequences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaAlignment) GetAlignedSequencesOk() ([]AlignedSequence, bool) {
	if o == nil || isNil(o.AlignedSequences) {
    return nil, false
	}
	return o.AlignedSequences, true
}

// HasAlignedSequences returns a boolean if a field has been set.
func (o *DnaAlignment) HasAlignedSequences() bool {
	if o != nil && !isNil(o.AlignedSequences) {
		return true
	}

	return false
}

// SetAlignedSequences gets a reference to the given []AlignedSequence and assigns it to the AlignedSequences field.
func (o *DnaAlignment) SetAlignedSequences(v []AlignedSequence) {
	o.AlignedSequences = v
}

func (o DnaAlignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ReferenceSequenceId) {
		toSerialize["referenceSequenceId"] = o.ReferenceSequenceId
	}
	if !isNil(o.WebURL) {
		toSerialize["webURL"] = o.WebURL
	}
	if !isNil(o.AlignedSequences) {
		toSerialize["alignedSequences"] = o.AlignedSequences
	}
	return json.Marshal(toSerialize)
}

type NullableDnaAlignment struct {
	value *DnaAlignment
	isSet bool
}

func (v NullableDnaAlignment) Get() *DnaAlignment {
	return v.value
}

func (v *NullableDnaAlignment) Set(val *DnaAlignment) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaAlignment) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaAlignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaAlignment(val *DnaAlignment) *NullableDnaAlignment {
	return &NullableDnaAlignment{value: val, isSet: true}
}

func (v NullableDnaAlignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaAlignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


