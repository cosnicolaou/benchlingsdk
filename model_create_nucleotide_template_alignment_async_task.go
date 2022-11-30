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

// CreateNucleotideTemplateAlignmentAsyncTask struct for CreateNucleotideTemplateAlignmentAsyncTask
type CreateNucleotideTemplateAlignmentAsyncTask struct {
	// Present only when status is FAILED for a bulk task. Contains information about the individual errors in the bulk task. 
	Errors map[string]interface{} `json:"errors,omitempty"`
	// Present only when status is FAILED. Contains information about the error.
	Message *string `json:"message,omitempty"`
	Response *NucleotideAlignment `json:"response,omitempty"`
	// The current state of the task.
	Status string `json:"status"`
}

// NewCreateNucleotideTemplateAlignmentAsyncTask instantiates a new CreateNucleotideTemplateAlignmentAsyncTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNucleotideTemplateAlignmentAsyncTask(status string) *CreateNucleotideTemplateAlignmentAsyncTask {
	this := CreateNucleotideTemplateAlignmentAsyncTask{}
	this.Status = status
	return &this
}

// NewCreateNucleotideTemplateAlignmentAsyncTaskWithDefaults instantiates a new CreateNucleotideTemplateAlignmentAsyncTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNucleotideTemplateAlignmentAsyncTaskWithDefaults() *CreateNucleotideTemplateAlignmentAsyncTask {
	this := CreateNucleotideTemplateAlignmentAsyncTask{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetErrors() map[string]interface{} {
	if o == nil || isNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Errors) {
    return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetMessage(v string) {
	o.Message = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetResponse() NucleotideAlignment {
	if o == nil || isNil(o.Response) {
		var ret NucleotideAlignment
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetResponseOk() (*NucleotideAlignment, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given NucleotideAlignment and assigns it to the Response field.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetResponse(v NucleotideAlignment) {
	o.Response = &v
}

// GetStatus returns the Status field value
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetStatus(v string) {
	o.Status = v
}

func (o CreateNucleotideTemplateAlignmentAsyncTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNucleotideTemplateAlignmentAsyncTask struct {
	value *CreateNucleotideTemplateAlignmentAsyncTask
	isSet bool
}

func (v NullableCreateNucleotideTemplateAlignmentAsyncTask) Get() *CreateNucleotideTemplateAlignmentAsyncTask {
	return v.value
}

func (v *NullableCreateNucleotideTemplateAlignmentAsyncTask) Set(val *CreateNucleotideTemplateAlignmentAsyncTask) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNucleotideTemplateAlignmentAsyncTask) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNucleotideTemplateAlignmentAsyncTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNucleotideTemplateAlignmentAsyncTask(val *CreateNucleotideTemplateAlignmentAsyncTask) *NullableCreateNucleotideTemplateAlignmentAsyncTask {
	return &NullableCreateNucleotideTemplateAlignmentAsyncTask{value: val, isSet: true}
}

func (v NullableCreateNucleotideTemplateAlignmentAsyncTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNucleotideTemplateAlignmentAsyncTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


