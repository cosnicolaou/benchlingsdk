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

// BulkRegisterEntitiesAsyncTask struct for BulkRegisterEntitiesAsyncTask
type BulkRegisterEntitiesAsyncTask struct {
	// Present only when status is FAILED for a bulk task. Contains information about the individual errors in the bulk task. 
	Errors map[string]interface{} `json:"errors,omitempty"`
	// Present only when status is FAILED. Contains information about the error.
	Message *string `json:"message,omitempty"`
	Response interface{} `json:"response,omitempty"`
	// The current state of the task.
	Status string `json:"status"`
}

// NewBulkRegisterEntitiesAsyncTask instantiates a new BulkRegisterEntitiesAsyncTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRegisterEntitiesAsyncTask(status string) *BulkRegisterEntitiesAsyncTask {
	this := BulkRegisterEntitiesAsyncTask{}
	this.Status = status
	return &this
}

// NewBulkRegisterEntitiesAsyncTaskWithDefaults instantiates a new BulkRegisterEntitiesAsyncTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRegisterEntitiesAsyncTaskWithDefaults() *BulkRegisterEntitiesAsyncTask {
	this := BulkRegisterEntitiesAsyncTask{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BulkRegisterEntitiesAsyncTask) GetErrors() map[string]interface{} {
	if o == nil || isNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRegisterEntitiesAsyncTask) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Errors) {
    return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BulkRegisterEntitiesAsyncTask) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *BulkRegisterEntitiesAsyncTask) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BulkRegisterEntitiesAsyncTask) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRegisterEntitiesAsyncTask) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkRegisterEntitiesAsyncTask) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BulkRegisterEntitiesAsyncTask) SetMessage(v string) {
	o.Message = &v
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkRegisterEntitiesAsyncTask) GetResponse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkRegisterEntitiesAsyncTask) GetResponseOk() (*interface{}, bool) {
	if o == nil || isNil(o.Response) {
    return nil, false
	}
	return &o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *BulkRegisterEntitiesAsyncTask) HasResponse() bool {
	if o != nil && isNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given interface{} and assigns it to the Response field.
func (o *BulkRegisterEntitiesAsyncTask) SetResponse(v interface{}) {
	o.Response = v
}

// GetStatus returns the Status field value
func (o *BulkRegisterEntitiesAsyncTask) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkRegisterEntitiesAsyncTask) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkRegisterEntitiesAsyncTask) SetStatus(v string) {
	o.Status = v
}

func (o BulkRegisterEntitiesAsyncTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBulkRegisterEntitiesAsyncTask struct {
	value *BulkRegisterEntitiesAsyncTask
	isSet bool
}

func (v NullableBulkRegisterEntitiesAsyncTask) Get() *BulkRegisterEntitiesAsyncTask {
	return v.value
}

func (v *NullableBulkRegisterEntitiesAsyncTask) Set(val *BulkRegisterEntitiesAsyncTask) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRegisterEntitiesAsyncTask) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRegisterEntitiesAsyncTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRegisterEntitiesAsyncTask(val *BulkRegisterEntitiesAsyncTask) *NullableBulkRegisterEntitiesAsyncTask {
	return &NullableBulkRegisterEntitiesAsyncTask{value: val, isSet: true}
}

func (v NullableBulkRegisterEntitiesAsyncTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRegisterEntitiesAsyncTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


