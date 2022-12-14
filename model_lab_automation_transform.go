/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
	"time"
)

// LabAutomationTransform struct for LabAutomationTransform
type LabAutomationTransform struct {
	// The canonical url of the transform in the API.
	ApiURL *string `json:"apiURL,omitempty"`
	BlobId NullableString `json:"blobId,omitempty"`
	CustomTransformId NullableString `json:"customTransformId,omitempty"`
	Errors *LabAutomationBenchlingAppErrors `json:"errors,omitempty"`
	Id *string `json:"id,omitempty"`
	// DateTime the transform was last modified.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	OutputProcessorId *string `json:"outputProcessorId,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewLabAutomationTransform instantiates a new LabAutomationTransform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabAutomationTransform() *LabAutomationTransform {
	this := LabAutomationTransform{}
	return &this
}

// NewLabAutomationTransformWithDefaults instantiates a new LabAutomationTransform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabAutomationTransformWithDefaults() *LabAutomationTransform {
	this := LabAutomationTransform{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *LabAutomationTransform) GetApiURL() string {
	if o == nil || isNil(o.ApiURL) {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationTransform) GetApiURLOk() (*string, bool) {
	if o == nil || isNil(o.ApiURL) {
    return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasApiURL() bool {
	if o != nil && !isNil(o.ApiURL) {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *LabAutomationTransform) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetBlobId returns the BlobId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LabAutomationTransform) GetBlobId() string {
	if o == nil || isNil(o.BlobId.Get()) {
		var ret string
		return ret
	}
	return *o.BlobId.Get()
}

// GetBlobIdOk returns a tuple with the BlobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabAutomationTransform) GetBlobIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BlobId.Get(), o.BlobId.IsSet()
}

// HasBlobId returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasBlobId() bool {
	if o != nil && o.BlobId.IsSet() {
		return true
	}

	return false
}

// SetBlobId gets a reference to the given NullableString and assigns it to the BlobId field.
func (o *LabAutomationTransform) SetBlobId(v string) {
	o.BlobId.Set(&v)
}
// SetBlobIdNil sets the value for BlobId to be an explicit nil
func (o *LabAutomationTransform) SetBlobIdNil() {
	o.BlobId.Set(nil)
}

// UnsetBlobId ensures that no value is present for BlobId, not even an explicit nil
func (o *LabAutomationTransform) UnsetBlobId() {
	o.BlobId.Unset()
}

// GetCustomTransformId returns the CustomTransformId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LabAutomationTransform) GetCustomTransformId() string {
	if o == nil || isNil(o.CustomTransformId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomTransformId.Get()
}

// GetCustomTransformIdOk returns a tuple with the CustomTransformId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabAutomationTransform) GetCustomTransformIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CustomTransformId.Get(), o.CustomTransformId.IsSet()
}

// HasCustomTransformId returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasCustomTransformId() bool {
	if o != nil && o.CustomTransformId.IsSet() {
		return true
	}

	return false
}

// SetCustomTransformId gets a reference to the given NullableString and assigns it to the CustomTransformId field.
func (o *LabAutomationTransform) SetCustomTransformId(v string) {
	o.CustomTransformId.Set(&v)
}
// SetCustomTransformIdNil sets the value for CustomTransformId to be an explicit nil
func (o *LabAutomationTransform) SetCustomTransformIdNil() {
	o.CustomTransformId.Set(nil)
}

// UnsetCustomTransformId ensures that no value is present for CustomTransformId, not even an explicit nil
func (o *LabAutomationTransform) UnsetCustomTransformId() {
	o.CustomTransformId.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *LabAutomationTransform) GetErrors() LabAutomationBenchlingAppErrors {
	if o == nil || isNil(o.Errors) {
		var ret LabAutomationBenchlingAppErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationTransform) GetErrorsOk() (*LabAutomationBenchlingAppErrors, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given LabAutomationBenchlingAppErrors and assigns it to the Errors field.
func (o *LabAutomationTransform) SetErrors(v LabAutomationBenchlingAppErrors) {
	o.Errors = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LabAutomationTransform) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationTransform) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LabAutomationTransform) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *LabAutomationTransform) GetModifiedAt() time.Time {
	if o == nil || isNil(o.ModifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationTransform) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModifiedAt) {
    return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasModifiedAt() bool {
	if o != nil && !isNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *LabAutomationTransform) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetOutputProcessorId returns the OutputProcessorId field value if set, zero value otherwise.
func (o *LabAutomationTransform) GetOutputProcessorId() string {
	if o == nil || isNil(o.OutputProcessorId) {
		var ret string
		return ret
	}
	return *o.OutputProcessorId
}

// GetOutputProcessorIdOk returns a tuple with the OutputProcessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationTransform) GetOutputProcessorIdOk() (*string, bool) {
	if o == nil || isNil(o.OutputProcessorId) {
    return nil, false
	}
	return o.OutputProcessorId, true
}

// HasOutputProcessorId returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasOutputProcessorId() bool {
	if o != nil && !isNil(o.OutputProcessorId) {
		return true
	}

	return false
}

// SetOutputProcessorId gets a reference to the given string and assigns it to the OutputProcessorId field.
func (o *LabAutomationTransform) SetOutputProcessorId(v string) {
	o.OutputProcessorId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LabAutomationTransform) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabAutomationTransform) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LabAutomationTransform) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LabAutomationTransform) SetStatus(v string) {
	o.Status = &v
}

func (o LabAutomationTransform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiURL) {
		toSerialize["apiURL"] = o.ApiURL
	}
	if o.BlobId.IsSet() {
		toSerialize["blobId"] = o.BlobId.Get()
	}
	if o.CustomTransformId.IsSet() {
		toSerialize["customTransformId"] = o.CustomTransformId.Get()
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !isNil(o.OutputProcessorId) {
		toSerialize["outputProcessorId"] = o.OutputProcessorId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableLabAutomationTransform struct {
	value *LabAutomationTransform
	isSet bool
}

func (v NullableLabAutomationTransform) Get() *LabAutomationTransform {
	return v.value
}

func (v *NullableLabAutomationTransform) Set(val *LabAutomationTransform) {
	v.value = val
	v.isSet = true
}

func (v NullableLabAutomationTransform) IsSet() bool {
	return v.isSet
}

func (v *NullableLabAutomationTransform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabAutomationTransform(val *LabAutomationTransform) *NullableLabAutomationTransform {
	return &NullableLabAutomationTransform{value: val, isSet: true}
}

func (v NullableLabAutomationTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabAutomationTransform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


