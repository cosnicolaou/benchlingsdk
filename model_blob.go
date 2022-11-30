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

// Blob struct for Blob
type Blob struct {
	// The universally unique identifier (UUID) for the blob.
	Id *string `json:"id,omitempty"`
	// eg. application/jpeg
	MimeType *string `json:"mimeType,omitempty"`
	// Name of the blob
	Name *string `json:"name,omitempty"`
	// One of RAW_FILE or VISUALIZATION. If VISUALIZATION, the blob may be displayed as an image preview. 
	Type *string `json:"type,omitempty"`
	UploadStatus *string `json:"uploadStatus,omitempty"`
}

// NewBlob instantiates a new Blob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlob() *Blob {
	this := Blob{}
	return &this
}

// NewBlobWithDefaults instantiates a new Blob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobWithDefaults() *Blob {
	this := Blob{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Blob) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Blob) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Blob) SetId(v string) {
	o.Id = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *Blob) GetMimeType() string {
	if o == nil || isNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetMimeTypeOk() (*string, bool) {
	if o == nil || isNil(o.MimeType) {
    return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *Blob) HasMimeType() bool {
	if o != nil && !isNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *Blob) SetMimeType(v string) {
	o.MimeType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Blob) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Blob) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Blob) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Blob) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Blob) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Blob) SetType(v string) {
	o.Type = &v
}

// GetUploadStatus returns the UploadStatus field value if set, zero value otherwise.
func (o *Blob) GetUploadStatus() string {
	if o == nil || isNil(o.UploadStatus) {
		var ret string
		return ret
	}
	return *o.UploadStatus
}

// GetUploadStatusOk returns a tuple with the UploadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetUploadStatusOk() (*string, bool) {
	if o == nil || isNil(o.UploadStatus) {
    return nil, false
	}
	return o.UploadStatus, true
}

// HasUploadStatus returns a boolean if a field has been set.
func (o *Blob) HasUploadStatus() bool {
	if o != nil && !isNil(o.UploadStatus) {
		return true
	}

	return false
}

// SetUploadStatus gets a reference to the given string and assigns it to the UploadStatus field.
func (o *Blob) SetUploadStatus(v string) {
	o.UploadStatus = &v
}

func (o Blob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UploadStatus) {
		toSerialize["uploadStatus"] = o.UploadStatus
	}
	return json.Marshal(toSerialize)
}

type NullableBlob struct {
	value *Blob
	isSet bool
}

func (v NullableBlob) Get() *Blob {
	return v.value
}

func (v *NullableBlob) Set(val *Blob) {
	v.value = val
	v.isSet = true
}

func (v NullableBlob) IsSet() bool {
	return v.isSet
}

func (v *NullableBlob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlob(val *Blob) *NullableBlob {
	return &NullableBlob{value: val, isSet: true}
}

func (v NullableBlob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


