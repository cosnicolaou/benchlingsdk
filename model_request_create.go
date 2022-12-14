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

// RequestCreate struct for RequestCreate
type RequestCreate struct {
	// Array of assignees
	Assignees []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee `json:"assignees,omitempty"`
	// The request's fields
	Fields ModelMap `json:"fields"`
	// The ID of the project to which the request belongs.
	ProjectId string `json:"projectId"`
	// ID of the user making the request. If unspecified, the requestor is the request creator. 
	RequestorId NullableString `json:"requestorId,omitempty"`
	SampleGroups []RequestSampleGroupCreate `json:"sampleGroups"`
	// Date the request is scheduled to be executed on, in YYYY-MM-DD format.
	ScheduledOn *string `json:"scheduledOn,omitempty"`
	// ID of the request's schema.
	SchemaId string `json:"schemaId"`
}

// NewRequestCreate instantiates a new RequestCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreate(fields ModelMap, projectId string, sampleGroups []RequestSampleGroupCreate, schemaId string) *RequestCreate {
	this := RequestCreate{}
	this.Fields = fields
	this.ProjectId = projectId
	this.SampleGroups = sampleGroups
	this.SchemaId = schemaId
	return &this
}

// NewRequestCreateWithDefaults instantiates a new RequestCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreateWithDefaults() *RequestCreate {
	this := RequestCreate{}
	return &this
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *RequestCreate) GetAssignees() []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee {
	if o == nil || isNil(o.Assignees) {
		var ret []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreate) GetAssigneesOk() ([]OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee, bool) {
	if o == nil || isNil(o.Assignees) {
    return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *RequestCreate) HasAssignees() bool {
	if o != nil && !isNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee and assigns it to the Assignees field.
func (o *RequestCreate) SetAssignees(v []OneOfRequestWriteUserAssigneeRequestWriteTeamAssignee) {
	o.Assignees = v
}

// GetFields returns the Fields field value
func (o *RequestCreate) GetFields() ModelMap {
	if o == nil {
		var ret ModelMap
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *RequestCreate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *RequestCreate) SetFields(v ModelMap) {
	o.Fields = v
}

// GetProjectId returns the ProjectId field value
func (o *RequestCreate) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RequestCreate) GetProjectIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RequestCreate) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRequestorId returns the RequestorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestCreate) GetRequestorId() string {
	if o == nil || isNil(o.RequestorId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestorId.Get()
}

// GetRequestorIdOk returns a tuple with the RequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestCreate) GetRequestorIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RequestorId.Get(), o.RequestorId.IsSet()
}

// HasRequestorId returns a boolean if a field has been set.
func (o *RequestCreate) HasRequestorId() bool {
	if o != nil && o.RequestorId.IsSet() {
		return true
	}

	return false
}

// SetRequestorId gets a reference to the given NullableString and assigns it to the RequestorId field.
func (o *RequestCreate) SetRequestorId(v string) {
	o.RequestorId.Set(&v)
}
// SetRequestorIdNil sets the value for RequestorId to be an explicit nil
func (o *RequestCreate) SetRequestorIdNil() {
	o.RequestorId.Set(nil)
}

// UnsetRequestorId ensures that no value is present for RequestorId, not even an explicit nil
func (o *RequestCreate) UnsetRequestorId() {
	o.RequestorId.Unset()
}

// GetSampleGroups returns the SampleGroups field value
func (o *RequestCreate) GetSampleGroups() []RequestSampleGroupCreate {
	if o == nil {
		var ret []RequestSampleGroupCreate
		return ret
	}

	return o.SampleGroups
}

// GetSampleGroupsOk returns a tuple with the SampleGroups field value
// and a boolean to check if the value has been set.
func (o *RequestCreate) GetSampleGroupsOk() ([]RequestSampleGroupCreate, bool) {
	if o == nil {
    return nil, false
	}
	return o.SampleGroups, true
}

// SetSampleGroups sets field value
func (o *RequestCreate) SetSampleGroups(v []RequestSampleGroupCreate) {
	o.SampleGroups = v
}

// GetScheduledOn returns the ScheduledOn field value if set, zero value otherwise.
func (o *RequestCreate) GetScheduledOn() string {
	if o == nil || isNil(o.ScheduledOn) {
		var ret string
		return ret
	}
	return *o.ScheduledOn
}

// GetScheduledOnOk returns a tuple with the ScheduledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreate) GetScheduledOnOk() (*string, bool) {
	if o == nil || isNil(o.ScheduledOn) {
    return nil, false
	}
	return o.ScheduledOn, true
}

// HasScheduledOn returns a boolean if a field has been set.
func (o *RequestCreate) HasScheduledOn() bool {
	if o != nil && !isNil(o.ScheduledOn) {
		return true
	}

	return false
}

// SetScheduledOn gets a reference to the given string and assigns it to the ScheduledOn field.
func (o *RequestCreate) SetScheduledOn(v string) {
	o.ScheduledOn = &v
}

// GetSchemaId returns the SchemaId field value
func (o *RequestCreate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *RequestCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *RequestCreate) SetSchemaId(v string) {
	o.SchemaId = v
}

func (o RequestCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assignees) {
		toSerialize["assignees"] = o.Assignees
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.RequestorId.IsSet() {
		toSerialize["requestorId"] = o.RequestorId.Get()
	}
	if true {
		toSerialize["sampleGroups"] = o.SampleGroups
	}
	if !isNil(o.ScheduledOn) {
		toSerialize["scheduledOn"] = o.ScheduledOn
	}
	if true {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableRequestCreate struct {
	value *RequestCreate
	isSet bool
}

func (v NullableRequestCreate) Get() *RequestCreate {
	return v.value
}

func (v *NullableRequestCreate) Set(val *RequestCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreate(val *RequestCreate) *NullableRequestCreate {
	return &NullableRequestCreate{value: val, isSet: true}
}

func (v NullableRequestCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


