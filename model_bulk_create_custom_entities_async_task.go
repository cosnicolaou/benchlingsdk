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

// BulkCreateCustomEntitiesAsyncTask struct for BulkCreateCustomEntitiesAsyncTask
type BulkCreateCustomEntitiesAsyncTask struct {
	// Present only when status is FAILED for a bulk task. Contains information about the individual errors in the bulk task. 
	Errors map[string]interface{} `json:"errors,omitempty"`
	// Present only when status is FAILED. Contains information about the error.
	Message *string `json:"message,omitempty"`
	Response *BulkCreateCustomEntitiesAsyncTaskAllOfResponse `json:"response,omitempty"`
	// The current state of the task.
	Status string `json:"status"`
}

// NewBulkCreateCustomEntitiesAsyncTask instantiates a new BulkCreateCustomEntitiesAsyncTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCreateCustomEntitiesAsyncTask(status string) *BulkCreateCustomEntitiesAsyncTask {
	this := BulkCreateCustomEntitiesAsyncTask{}
	this.Status = status
	return &this
}

// NewBulkCreateCustomEntitiesAsyncTaskWithDefaults instantiates a new BulkCreateCustomEntitiesAsyncTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCreateCustomEntitiesAsyncTaskWithDefaults() *BulkCreateCustomEntitiesAsyncTask {
	this := BulkCreateCustomEntitiesAsyncTask{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BulkCreateCustomEntitiesAsyncTask) GetErrors() map[string]interface{} {
	if o == nil || isNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Errors) {
    return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *BulkCreateCustomEntitiesAsyncTask) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BulkCreateCustomEntitiesAsyncTask) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BulkCreateCustomEntitiesAsyncTask) SetMessage(v string) {
	o.Message = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *BulkCreateCustomEntitiesAsyncTask) GetResponse() BulkCreateCustomEntitiesAsyncTaskAllOfResponse {
	if o == nil || isNil(o.Response) {
		var ret BulkCreateCustomEntitiesAsyncTaskAllOfResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) GetResponseOk() (*BulkCreateCustomEntitiesAsyncTaskAllOfResponse, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) HasResponse() bool {
	if o != nil && !isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given BulkCreateCustomEntitiesAsyncTaskAllOfResponse and assigns it to the Response field.
func (o *BulkCreateCustomEntitiesAsyncTask) SetResponse(v BulkCreateCustomEntitiesAsyncTaskAllOfResponse) {
	o.Response = &v
}

// GetStatus returns the Status field value
func (o *BulkCreateCustomEntitiesAsyncTask) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkCreateCustomEntitiesAsyncTask) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkCreateCustomEntitiesAsyncTask) SetStatus(v string) {
	o.Status = v
}

func (o BulkCreateCustomEntitiesAsyncTask) MarshalJSON() ([]byte, error) {
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

type NullableBulkCreateCustomEntitiesAsyncTask struct {
	value *BulkCreateCustomEntitiesAsyncTask
	isSet bool
}

func (v NullableBulkCreateCustomEntitiesAsyncTask) Get() *BulkCreateCustomEntitiesAsyncTask {
	return v.value
}

func (v *NullableBulkCreateCustomEntitiesAsyncTask) Set(val *BulkCreateCustomEntitiesAsyncTask) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCreateCustomEntitiesAsyncTask) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCreateCustomEntitiesAsyncTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCreateCustomEntitiesAsyncTask(val *BulkCreateCustomEntitiesAsyncTask) *NullableBulkCreateCustomEntitiesAsyncTask {
	return &NullableBulkCreateCustomEntitiesAsyncTask{value: val, isSet: true}
}

func (v NullableBulkCreateCustomEntitiesAsyncTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCreateCustomEntitiesAsyncTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


